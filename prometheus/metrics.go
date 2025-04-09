package prometheus

import "github.com/prometheus/client_golang/prometheus"

var (
	// Device Info metrics
	deviceInfoDesc = prometheus.NewDesc(
		"device_info",
		"Device information",
		[]string{"manufacturer", "manufacturer_oui", "version_date", "boot_version",
			"software_version", "software_version_extended", "serial_number",
			"model", "hardware_version", "onu_alias"},
		nil,
	)
	cpuUsageDesc = prometheus.NewDesc(
		"usage_cpu",
		"CPU usage percentage per core",
		[]string{"core"},
		nil,
	)
	memoryUsageDesc = prometheus.NewDesc(
		"usage_memory",
		"Memory usage percentage",
		nil,
		nil,
	)
	uptimeDesc = prometheus.NewDesc(
		"device_uptime",
		"Device uptime in seconds",
		nil,
		nil,
	)

	// LAN Info metrics
	bytesDesc = prometheus.NewDesc(
		"octets_total",
		"Number of bytes transmitted/received",
		[]string{"direction"},
		nil,
	)
	packetsDesc = prometheus.NewDesc(
		"packets_total",
		"Number of packets transmitted/received",
		[]string{"direction", "type"},
		nil,
	)
	errorsDesc = prometheus.NewDesc(
		"packets_errors_total",
		"Number of network errors",
		[]string{"direction"},
		nil,
	)
	discardsDesc = prometheus.NewDesc(
		"packets_discards_total",
		"Number of discarded packets",
		[]string{"direction"},
		nil,
	)
	networkStatusDesc = prometheus.NewDesc(
		"ethernet_status",
		"Ethernet interface status (1 = up, 0 = down)",
		[]string{"speed", "duplex"},
		nil,
	)

	// Optical Info metrics
	opticalSignalDesc = prometheus.NewDesc(
		"optical_power",
		"Optical signal power (dBm)",
		[]string{"direction"},
		nil,
	)
	opticalTempDesc = prometheus.NewDesc(
		"optical_temperature",
		"Optical module temperature in Celsius",
		nil,
		nil,
	)
	opticalVoltageDesc = prometheus.NewDesc(
		"optical_voltage",
		"Optical module supply voltage",
		nil,
		nil,
	)
	opticalBiasCurrentDesc = prometheus.NewDesc(
		"optical_bias_current",
		"Optical transmitter bias current in mA",
		nil,
		nil,
	)
	opticalStatusDesc = prometheus.NewDesc(
		"optical_status",
		"GPON status indicators",
		[]string{"type"},
		nil,
	)
	rfPowerDesc = prometheus.NewDesc(
		"optical_rf_power",
		"RF power levels in dBm",
		[]string{"type"},
		nil,
	)
)
