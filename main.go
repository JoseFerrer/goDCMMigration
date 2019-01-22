package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

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

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose the process to run (H - Only Harvest, T - Only Transfer, A - Analyse, HT - Harvest/Transfer):")
	text, _ := reader.ReadString('\n')
	lenStr := len(text) - 1
	if text == "H\n" {
		fmt.Println(text[0:lenStr])
	}

	return

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
	if (text == "H") || (text == "HT") {
		harvestQuery(confFind, session, "Middle")
	}

	// Transfer
	fmt.Printf("\n************************************************************************** \n\t\t\t\t Start Transfer Process \n************************************************************************** \n")
	if (text == "T") || (text == "HT") {
		transferQuery(confFind, session, "Middle")
	}

}
