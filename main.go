package main

import (
	"PrometheusF6005/ont"
	"encoding/json"
	"fmt"
)

func main() {
	session, err := ont.Login("http://192.168.1.1", "admin", "admin")
	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}

	data, _ := session.LoadOpticalData()
	pl, _ := json.Marshal(data)
	fmt.Println(string(pl))
	data2, err := session.LoadLanInfo()
	pl, _ = json.Marshal(data2)
	fmt.Println(string(pl))
	data3, _ := session.LoadDeviceInfo()
	pl, _ = json.Marshal(data3)
	fmt.Println(string(pl))
	return
}
