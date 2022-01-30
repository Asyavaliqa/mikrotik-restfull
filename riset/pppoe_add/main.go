package main

import (
	"flag"
	"fmt"
	"gopkg.in/routeros.v2"
	"log"
	"strings"
)

var (
	command    = fmt.Sprintf(`/ppp/secret/set =name=xx@gmail.com =comment=xx =.id=*15`)
	address    = "127.0.0.1:28728"
	username   = "admin"
	password   = ""
	properties = "uptime,version"
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

	//for _, re := range r.Re {
	//	for _, p := range strings.Split(*properties, ",") {
	//		fmt.Print(re.Map[p], "\t")
	//	}
	//	fmt.Print("\n")
	//}
	log.Println(r.Done.String())
	//log.Print(r.Re[0].Map["version"])
}
