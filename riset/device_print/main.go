package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"

	"gopkg.in/routeros.v2"
)

type DeviceInfo struct {
	Platform          string `json:"platform"`
	Board             string `json:"board-name"`
	Uptime            string `json:"uptime"`
	Version           string `json:"version"`
	BuildTime         string `json:"build-time"`
	FreeMemory        string `json:"free-memory"`
	TotalMemory       string `json:"total-memory"`
	Cpu               string `json:"cpu"`
	CpuCount          string `json:"cpu-count"`
	CpuFrequency      string `json:"cpu-frequency"`
	CpuLoad           string `json:"cpu-load"`
	FreeHddSpace      string `json:"free-hdd-space"`
	TotalHddSpace     string `json:"total-hdd-space"`
	BadBlocks         string `json:"bad-blocks"`
	ArchitectureNames string `json:"architecture-name"`
}

// [{`uptime` `3d15h22m9s`} {`version` `6.47.1 (stable)`} {`build-time` `Jul/08/2020 12:34:22`} {`free-memory` `228204544`} {`total-memory` `268435456`} {`cpu` `MIPS 24Kc V7.4`} {`cpu-count` `1`} {`cpu-frequency` `680`} {`cpu-load` `32`} {`free-hdd-space` `510930944`} {`total-hdd-space` `536870912`} {`write-sect-since-reboot` `4920`} {`write-sect-total` `10856`} {`bad-blocks` `0.4`} {`architecture-name` `mipsbe`} {`board-name` `RB450G`} {`platform` `MikroTik`}]
var (
	command  = "/system/resource/print"
	address  = "192.168.4.1:8728"
	username = "admin"
	password = "segopecel12"
)

func main() {
	flag.Parse()

	c, err := routeros.Dial(address, username, password)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	r, err := c.RunArgs(strings.Split(command, " "))
	if err != nil {
		log.Fatal(err)
	}

	jsonbody, err := json.Marshal(r.Re[0].Map)
	if err != nil {
		// do error check
		fmt.Println(err)
		return
	}
	deviceInfo := DeviceInfo{}
	if err := json.Unmarshal(jsonbody, &deviceInfo); err != nil {
		// do error check
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", deviceInfo)
}
