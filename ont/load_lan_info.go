package ont

import (
	"encoding/xml"
	"strconv"
	"time"
)

type LanInfo struct {
	PacketsDiscardedIn  int
	PacketsDiscardedOut int

	PacketsErrorIn  int
	PacketsErrorOut int

	PacketsMulticastIn  int
	PacketsMulticastOut int

	PacketsUnicastIn  int
	PacketsUnicastOut int

	BytesIn  int
	BytesOut int

	PacketsIn  int
	PacketsOut int

	Status int
	Duplex string
	Speed  int
}

func (s *Session) LoadLanInfo() (*LanInfo, error) {
	_, _ = s.Get(s.Endpoint + "/?_type=menuView&_tag=localNetStatus&Menu3Location=0&_" + strconv.FormatInt(time.Now().Unix(), 10))
	resp, err := s.Get(s.Endpoint + "/?_type=menuData&_tag=status_lan_info_lua.lua&_=" + strconv.FormatInt(time.Now().Unix(), 10))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result LanInfoResponse
	if err := xml.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Convert(), nil
}

type LanInfoResponse struct {
	XMLName                 xml.Name `xml:"ajax_response_xml_root"`
	Text                    string   `xml:",chardata"`
	IFERRORPARAM            string   `xml:"IF_ERRORPARAM"`
	IFERRORTYPE             string   `xml:"IF_ERRORTYPE"`
	IFERRORSTR              string   `xml:"IF_ERRORSTR"`
	IFERRORID               string   `xml:"IF_ERRORID"`
	OBJPONPORTBASICSTATUSID struct {
		Text     string `xml:",chardata"`
		Instance struct {
			Text      string   `xml:",chardata"`
			ParaName  []string `xml:"ParaName"`
			ParaValue []string `xml:"ParaValue"`
		} `xml:"Instance"`
	} `xml:"OBJ_PON_PORT_BASIC_STATUS_ID"`
}

func (result LanInfoResponse) Convert() *LanInfo {
	var lanInfo LanInfo

	for i, name := range result.OBJPONPORTBASICSTATUSID.Instance.ParaName {
		if i < len(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue) {
			switch name {
			case "InDiscard":
				lanInfo.PacketsDiscardedIn, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			case "OutDiscard":
				lanInfo.PacketsDiscardedOut, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			case "InError":
				lanInfo.PacketsErrorIn, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			case "OutError":
				lanInfo.PacketsErrorOut, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			case "InMulticast":
				lanInfo.PacketsMulticastIn, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			case "OutMulticast":
				lanInfo.PacketsMulticastOut, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			case "InUnicast":
				lanInfo.PacketsUnicastIn, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			case "OutUnicast":
				lanInfo.PacketsUnicastOut, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			case "InBytes":
				lanInfo.BytesIn, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			case "OutBytes":
				lanInfo.BytesOut, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			case "InPkts":
				lanInfo.PacketsIn, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			case "OutPkts":
				lanInfo.PacketsOut, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			case "Status":
				lanInfo.Status, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			case "Duplex":
				lanInfo.Duplex = result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i]
			case "Speed":
				lanInfo.Speed, _ = strconv.Atoi(result.OBJPONPORTBASICSTATUSID.Instance.ParaValue[i])
			}
		}
	}

	return &lanInfo
}
