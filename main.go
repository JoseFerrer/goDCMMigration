package main

import (
	"encoding/json"
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

const (
	url         = "localhost"
	DATABASE    = "migration"
	HCOLLECTION = "harvest"
	TCOLLECTIOM = "transfer"
	HPROCESS    = "hprocess"
)

func main() {

	// MongoDB initialization
	session, _ := mgo.Dial(url)
	defer session.Close()

	// Initial state
	byteValue := readJson("config.json")
	var confFind ConfigJSON
	json.Unmarshal(byteValue, &confFind)

	// Check time to sleep
	sleepFunc(confFind, "Initial")

	// Harvest
	fmt.Printf("\n************************************************************************** \n\t\t\t\t Start Harvest Process \n************************************************************************** \n")
	//harvestQuery(confFind, session, "Middle")

	// Transfer
	fmt.Printf("\n************************************************************************** \n\t\t\t\t Start Transfer Process \n************************************************************************** \n")
	transferQuery(confFind, session, "Middle")

}
