package prometheus

import (
	"PrometheusF6005/ont"
	"cmp"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"os"
	"strconv"
	"time"
)

// ONTCollector implements the prometheus.Collector interface
type ONTCollector struct {
	session *ont.Session
}

// NewONTCollector creates a new ONT metrics collector
func NewONTCollector(session *ont.Session) *ONTCollector {
	return &ONTCollector{
		session: session,
	}
}

// Describe implements prometheus.Collector
func (c *ONTCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- deviceInfoDesc
	ch <- cpuUsageDesc
	ch <- memoryUsageDesc
	ch <- uptimeDesc
	ch <- bytesDesc
	ch <- packetsDesc
	ch <- errorsDesc
	ch <- discardsDesc
	ch <- networkStatusDesc
	ch <- opticalSignalDesc
	ch <- opticalTempDesc
	ch <- opticalVoltageDesc
	ch <- opticalBiasCurrentDesc
	ch <- opticalStatusDesc
	ch <- rfPowerDesc
	ch <- opticalUptimeDesc
}

func sleepQuit(reaason string) {
	sleepTimeString := cmp.Or(os.Getenv("ONT_SLEEP_QUIT"), "60")
	sleepTime, _ := strconv.Atoi(sleepTimeString)

	log.Printf("[SleepQuit] %s, sleeping for %d seconds before quitting...\n", reaason, sleepTime)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	log.Println("[SleepQuit] Sleep time is over, exiting...")
	os.Exit(1)
}

// Collect implements prometheus.Collector
func (c *ONTCollector) Collect(ch chan<- prometheus.Metric) {
	// Collect Device Info
	deviceInfo, err := c.session.LoadDeviceInfo()
	if err != nil {
		sleepQuit(err.Error())
		return
	}

	ch <- prometheus.MustNewConstMetric(
		deviceInfoDesc,
		prometheus.GaugeValue,
		1,
		deviceInfo.Manufacturer,
		deviceInfo.ManufacturerOui,
		deviceInfo.VersionDate,
		deviceInfo.BootVersion,
		deviceInfo.SofwareVersion,
		deviceInfo.SoftwareVersionExtended,
		deviceInfo.SerialNumber,
		deviceInfo.Model,
		deviceInfo.HardwareVersion,
		deviceInfo.OnuAlias,
	)

	// CPU Usage metrics
	ch <- prometheus.MustNewConstMetric(
		cpuUsageDesc,
		prometheus.GaugeValue,
		float64(deviceInfo.CPUUsage1),
		"1",
	)
	ch <- prometheus.MustNewConstMetric(
		cpuUsageDesc,
		prometheus.GaugeValue,
		float64(deviceInfo.CPUUsage2),
		"2",
	)
	ch <- prometheus.MustNewConstMetric(
		cpuUsageDesc,
		prometheus.GaugeValue,
		float64(deviceInfo.CPUUsage3),
		"3",
	)
	ch <- prometheus.MustNewConstMetric(
		cpuUsageDesc,
		prometheus.GaugeValue,
		float64(deviceInfo.CPUUsage4),
		"4",
	)

	// Memory Usage metric
	ch <- prometheus.MustNewConstMetric(
		memoryUsageDesc,
		prometheus.GaugeValue,
		float64(deviceInfo.MemoryUsage),
	)

	// Uptime metric
	ch <- prometheus.MustNewConstMetric(
		uptimeDesc,
		prometheus.GaugeValue,
		float64(deviceInfo.Uptime),
	)

	// Collect LAN Info
	lanInfo, err := c.session.LoadLanInfo()
	if err != nil {
		sleepQuit(err.Error())
		return
	}
	// Network traffic metrics
	ch <- prometheus.MustNewConstMetric(
		bytesDesc,
		prometheus.CounterValue,
		float64(lanInfo.BytesIn),
		"in",
	)
	ch <- prometheus.MustNewConstMetric(
		bytesDesc,
		prometheus.CounterValue,
		float64(lanInfo.BytesOut),
		"out",
	)

	// Packet metrics
	ch <- prometheus.MustNewConstMetric(
		packetsDesc,
		prometheus.CounterValue,
		float64(lanInfo.PacketsUnicastIn),
		"in",
		"unicast",
	)
	ch <- prometheus.MustNewConstMetric(
		packetsDesc,
		prometheus.CounterValue,
		float64(lanInfo.PacketsUnicastOut),
		"out",
		"unicast",
	)
	ch <- prometheus.MustNewConstMetric(
		packetsDesc,
		prometheus.CounterValue,
		float64(lanInfo.PacketsMulticastIn),
		"in",
		"multicast",
	)
	ch <- prometheus.MustNewConstMetric(
		packetsDesc,
		prometheus.CounterValue,
		float64(lanInfo.PacketsMulticastOut),
		"out",
		"multicast",
	)

	// Error metrics
	ch <- prometheus.MustNewConstMetric(
		errorsDesc,
		prometheus.CounterValue,
		float64(lanInfo.PacketsErrorIn),
		"in",
	)
	ch <- prometheus.MustNewConstMetric(
		errorsDesc,
		prometheus.CounterValue,
		float64(lanInfo.PacketsErrorOut),
		"out",
	)

	// Discard metrics
	ch <- prometheus.MustNewConstMetric(
		discardsDesc,
		prometheus.CounterValue,
		float64(lanInfo.PacketsDiscardedIn),
		"in",
	)
	ch <- prometheus.MustNewConstMetric(
		discardsDesc,
		prometheus.CounterValue,
		float64(lanInfo.PacketsDiscardedOut),
		"out",
	)

	// Status metric
	ch <- prometheus.MustNewConstMetric(
		networkStatusDesc,
		prometheus.GaugeValue,
		float64(lanInfo.Status),
		strconv.Itoa(lanInfo.Speed),
		lanInfo.Duplex,
	)

	// Collect Optical Info
	opticalInfo, err := c.session.LoadOpticalData()
	if err != nil {
		sleepQuit(err.Error())
		return
	}
	// Signal power metrics
	ch <- prometheus.MustNewConstMetric(
		opticalSignalDesc,
		prometheus.GaugeValue,
		opticalInfo.TXPower,
		"tx",
	)
	ch <- prometheus.MustNewConstMetric(
		opticalSignalDesc,
		prometheus.GaugeValue,
		opticalInfo.RXPower,
		"rx",
	)

	// Temperature metric
	ch <- prometheus.MustNewConstMetric(
		opticalTempDesc,
		prometheus.GaugeValue,
		opticalInfo.OpticalModuleTemperature,
	)

	// Voltage metric
	ch <- prometheus.MustNewConstMetric(
		opticalVoltageDesc,
		prometheus.GaugeValue,
		float64(opticalInfo.OpticalModuleVoltage),
	)

	// Bias current metric
	ch <- prometheus.MustNewConstMetric(
		opticalBiasCurrentDesc,
		prometheus.GaugeValue,
		float64(opticalInfo.OpticalModuleBiasCurrent),
	)

	// Status metrics
	ch <- prometheus.MustNewConstMetric(
		opticalStatusDesc,
		prometheus.GaugeValue,
		float64(opticalInfo.LoS),
		"los",
	)
	ch <- prometheus.MustNewConstMetric(
		opticalStatusDesc,
		prometheus.GaugeValue,
		float64(opticalInfo.GPONRegistrationStatus),
		"gpon_registration",
	)
	ch <- prometheus.MustNewConstMetric(
		opticalStatusDesc,
		prometheus.GaugeValue,
		float64(opticalInfo.PONCatV),
		"catv",
	)

	// RF power metrics
	ch <- prometheus.MustNewConstMetric(
		rfPowerDesc,
		prometheus.GaugeValue,
		float64(opticalInfo.RFTXPower),
		"tx",
	)
	ch <- prometheus.MustNewConstMetric(
		rfPowerDesc,
		prometheus.GaugeValue,
		float64(opticalInfo.VideoRXPower),
		"rx",
	)
	ch <- prometheus.MustNewConstMetric(
		opticalUptimeDesc,
		prometheus.CounterValue,
		float64(opticalInfo.Uptime),
	)
}
