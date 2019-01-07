package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	mgo "gopkg.in/mgo.v2"
)

// Moved struct
type Moved struct {
	ID      string `bson:"id"`
	IDH     string `bson:"idh"`
	AccNum  string `bson:"accnum"`
	StudyID string `bson:"studyid"`
}

func queryDCM(q string) string {
	// Execute FINDSCU
	cmd := exec.Command("/bin/bash", "-c", q)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd.Run() failedwith %s\n ", err)
	}
	return string(out)
}

func main() {

	var (
		url         = "localhost"
		DATABASE    = "migration"
		TCOLLECTIOM = "transfer"
	)

	// MongoDB initialization
	session, _ := mgo.Dial(url)
	defer session.Close()

	colTrans := session.DB(DATABASE).C(TCOLLECTIOM)

	path := "./dcm4che5d13d2/bin/dcmdump "
	fileDir := "--directory=results/"

	files, _ := ioutil.ReadDir(fileDir)
	for _, file := range files {
		fmt.Println(file.Name())
		command := path + fileDir + file.Name()
		query := queryDCM(command)
		nro := strings.Count(query, "ONLNE")
	}

	TransData := Moved{}
	dbSize, _ := colTrans.Count()

}
