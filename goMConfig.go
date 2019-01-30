package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// error to string
type error interface {
	Error() string
}

// ConfigInit istruct
type ConfigInit struct {
	url       string
	DataBase1 string
	DataBase2 string
	DataBase3 string
	DataBase4 string
	ProcessH  string
}

//
type ConfigOut struct {
	DBState bool
	HState  bool
	TState  bool
	VState  bool
}

// ComQuery istruct
type ComQuery struct {
	State    string
	iRange   string
	oRange   string
	Tag      string
	TagValue string
}

// readJson fucntion to read json configuration
func readJson(path string) []byte {
	// Open JSON File Configuration
	jsonFile, err := os.Open(path)
	if err != nil {
		logFunction("JSON File Configuration is not found. " + err.Error())
		fmt.Println("JSON File Configuration is not found. ", err)
	}
	fmt.Println("Successfully Opened configuration.json")
	logFunction("Read JSON: Successfully Opened configuration.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func (o *ComQuery) getCommands(confFind ConfigJSON, value string) string {

	var (
		ConfOpt    string
		comando1   string
		modOpt     string
		ejecutable string
	)
	ejecutable = ""
	switch o.State {
	case "H":
		ejecutable = "./dcm4che5d13d2/bin/findscu "
	case "T":
		ejecutable = "./dcm4che5d13d2/bin/movescu "
	case "Q":
		ejecutable = "./dcm4che5d13d2/bin/findscu "
	}

	// PACS Configuration
	pacsaetitle := confFind.PacsAETitle
	ip := confFind.PacsIP
	port := confFind.PacsPort
	modPACS := "-c " + pacsaetitle + "@" + ip + ":" + port

	// Entity Configuration
	entityaetitle := confFind.ENTITYAETitle
	entityip := confFind.ENTITYIP
	entityport := confFind.ENTITYPort
	modAETitle := "-b " + entityaetitle + "@" + entityip + ":" + entityport

	// *********************************************************

	modOpt = ""
	switch o.State {
	case "H":
		// Elapsed date configuration
		flag, _ := strconv.Atoi(confFind.DateStart)
		ConfOpt = ""
		if flag != 0 {
			ConfOpt = "-m " + o.Tag + "=" + o.iRange + o.oRange + " -L STUDY "
			strs := []string{StudyID, AccessionNum, PatientID, StudyDesc, Modality, NStudyRS, NStudyRI}
			modOpt = strings.Join(strs, " ")
		}
		break
	case "T":
		ConfOpt = "-m " + o.Tag + "=" + value
		modOpt = " --dest " + confFind.BackUpTitle
	case "Q":
		// Search by Accession Number
		ConfOpt = "-m " + o.Tag + "=" + o.TagValue
		modOpt = ""
	}

	comando1 = ejecutable + " " + modAETitle + " " + modPACS + " " + ConfOpt + modOpt

	return comando1
}

func queryDCM(q string) string {
	// Execute FINDSCU
	cmd := exec.Command("/bin/bash", "-c", q)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd.Run() failedwith %s\n ", err)
		logFunction("cmd.Run() failedwith %s\n " + err.Error())
	}
	return string(out)
}

func logFunction(message string) {
	f, err := os.OpenFile("migration.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logFunction("error opening file: %v" + err.Error())
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(message)
}
