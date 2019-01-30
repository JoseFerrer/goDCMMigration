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
	TCOLLECTION = "transfer"
	HPROCESS    = "hprocess"
	QTYFIND     = "qualityf"
	QTYNOTFIND  = "qualitynf"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose the process to run (H - Only Harvest, T - Only Transfer, A - Analyse, HT - Harvest/Transfer, HTA - Harvest/Transfer/Analyse)")
	text, _ := reader.ReadString('\n')
	lenStr := len(text) - 1
	text = text[0:lenStr]

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
		fmt.Println("Start Harvest")
		harvestQuery(confFind, session, "Middle")
	}

	// Transfer
	fmt.Printf("\n************************************************************************** \n\t\t\t\t Start Transfer Process \n************************************************************************** \n")
	if (text == "T") || (text == "HT") {
		fmt.Println("Start Tranfer")
		transferQuery(confFind, session, "Middle")
	}

	// Quality and Assurance
	fmt.Printf("\n************************************************************************** \n\t\t\t\t Start Q&S Process \n************************************************************************** \n")
	if (text == "Q") || (text == "HTQ") {
		fmt.Println("Start Quality & Assurance")
		qualityQuery(confFind, session, "Middle")
		resp := bufio.NewReader(os.Stdin)
		fmt.Print("Do you want to move DICOMs not moved ? ( Y / N )")
		texto, _ := resp.ReadString('\n')
		lenSt := len(texto) - 1
		text = texto[0:lenSt]
		if texto == "Y" {
			fmt.Println("Start new Migration")
			newDCMMigration(confFind, session)
		}
	}
}
