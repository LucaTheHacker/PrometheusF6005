package ont

import (
	"encoding/xml"
	"strconv"
	"time"
)

type OpticalInfo struct {
	LoS                      int
	GPONRegistrationStatus   int
	PONCatV                  int
	Uptime                   int
	OpticalModuleTemperature float64
	OpticalModuleVoltage     int
	OpticalModuleBiasCurrent int
	RFTXPower                int
	VideoRXPower             int
	TXPower                  float64
	RXPower                  float64
}

func (s *Session) LoadOpticalData() (*OpticalInfo, error) {
	resp, err := s.Get(s.Endpoint + "/?_type=menuData&_tag=optical_info_lua.lua&_=" + strconv.FormatInt(time.Now().Unix(), 10))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result OpticalDataResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Convert(), nil
}

type OpticalDataResponse struct {
	XMLName      xml.Name `xml:"ajax_response_xml_root"`
	Text         string   `xml:",chardata"`
	IFERRORPARAM string   `xml:"IF_ERRORPARAM"`
	IFERRORTYPE  string   `xml:"IF_ERRORTYPE"`
	IFERRORSTR   string   `xml:"IF_ERRORSTR"`
	IFERRORID    string   `xml:"IF_ERRORID"`
	OBJLOSINFOID struct {
		Text     string `xml:",chardata"`
		Instance struct {
			Text      string   `xml:",chardata"`
			ParaName  []string `xml:"ParaName"`
			ParaValue []string `xml:"ParaValue"`
		} `xml:"Instance"`
	} `xml:"OBJ_LOS_INFO_ID"`
	OBJGPONREGSTATUSID struct {
		Text     string `xml:",chardata"`
		Instance struct {
			Text      string   `xml:",chardata"`
			ParaName  []string `xml:"ParaName"`
			ParaValue []string `xml:"ParaValue"`
		} `xml:"Instance"`
	} `xml:"OBJ_GPONREGSTATUS_ID"`
	OBJPONCATVID struct {
		Text     string `xml:",chardata"`
		Instance struct {
			Text      string   `xml:",chardata"`
			ParaName  []string `xml:"ParaName"`
			ParaValue []string `xml:"ParaValue"`
		} `xml:"Instance"`
	} `xml:"OBJ_PON_CATV_ID"`
	OBJPONPOWERONTIMEID struct {
		Text     string `xml:",chardata"`
		Instance struct {
			Text      string   `xml:",chardata"`
			ParaName  []string `xml:"ParaName"`
			ParaValue []string `xml:"ParaValue"`
		} `xml:"Instance"`
	} `xml:"OBJ_PON_POWERONTIME_ID"`
	OBJPONOPTICALPARAID struct {
		Text     string `xml:",chardata"`
		Instance struct {
			Text      string   `xml:",chardata"`
			ParaName  []string `xml:"ParaName"`
			ParaValue []string `xml:"ParaValue"`
		} `xml:"Instance"`
	} `xml:"OBJ_PON_OPTICALPARA_ID"`
}

func (result OpticalDataResponse) Convert() *OpticalInfo {
	var opticalInfo OpticalInfo
	for i, name := range result.OBJLOSINFOID.Instance.ParaName {
		if i < len(result.OBJLOSINFOID.Instance.ParaValue) {
			value := result.OBJLOSINFOID.Instance.ParaValue[i]
			if name == "LOS" {
				los, err := strconv.Atoi(value)
				if err == nil {
					opticalInfo.LoS = los
				}
			}
		}
	}

	// Extract GPON registration status
	for i, name := range result.OBJGPONREGSTATUSID.Instance.ParaName {
		if i < len(result.OBJGPONREGSTATUSID.Instance.ParaValue) {
			value := result.OBJGPONREGSTATUSID.Instance.ParaValue[i]
			if name == "RegStatus" {
				status, err := strconv.Atoi(value)
				if err == nil {
					opticalInfo.GPONRegistrationStatus = status
				}
			}
		}
	}

	// Extract PON CATV information
	for i, name := range result.OBJPONCATVID.Instance.ParaName {
		if i < len(result.OBJPONCATVID.Instance.ParaValue) {
			value := result.OBJPONCATVID.Instance.ParaValue[i]
			if name == "CatvEnable" {
				catv, err := strconv.Atoi(value)
				if err == nil {
					opticalInfo.PONCatV = catv
				}
			}
		}
	}

	// Extract power-on time (uptime)
	for i, name := range result.OBJPONPOWERONTIMEID.Instance.ParaName {
		if i < len(result.OBJPONPOWERONTIMEID.Instance.ParaValue) {
			value := result.OBJPONPOWERONTIMEID.Instance.ParaValue[i]
			if name == "PONOnTime" {
				uptime, err := strconv.Atoi(value)
				if err == nil {
					opticalInfo.Uptime = uptime
				}
			}
		}
	}

	// Extract optical parameters
	for i, name := range result.OBJPONOPTICALPARAID.Instance.ParaName {
		if i < len(result.OBJPONOPTICALPARAID.Instance.ParaValue) {
			value := result.OBJPONOPTICALPARAID.Instance.ParaValue[i]

			switch name {
			case "Current":
				current, err := strconv.Atoi(value)
				if err == nil {
					opticalInfo.OpticalModuleBiasCurrent = current
				}
			case "Temp":
				temp, err := strconv.ParseFloat(value, 64)
				if err == nil {
					opticalInfo.OpticalModuleTemperature = temp
				}
			case "RFTxPower":
				power, err := strconv.Atoi(value)
				if err == nil {
					opticalInfo.RFTXPower = power
				}
			case "VideoRxPower":
				power, err := strconv.Atoi(value)
				if err == nil {
					opticalInfo.VideoRXPower = power
				}
			case "TxPower":
				power, err := strconv.ParseFloat(value, 64)
				if err == nil {
					opticalInfo.TXPower = power
				}
			case "RxPower":
				power, err := strconv.ParseFloat(value, 64)
				if err == nil {
					opticalInfo.RXPower = power
				}
			case "Volt":
				voltage, err := strconv.Atoi(value)
				if err == nil {
					opticalInfo.OpticalModuleVoltage = voltage
				}
			}
		}
	}

	return &opticalInfo
}
