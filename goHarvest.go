package main

import (
	"fmt"
	"strconv"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Harvest struct
type Harvest struct {
	ID               string `bson:"id"`
	AccessionNumber  string `bson:"accessionnumber"`
	StudyID          string `bson:"studyid"`
	PatientID        string `bson:"patientid"`
	Modality         string `bson:"modality"`
	StudyDescription string `bson:"studydescription"`
	StudyDate        string `bson:"studydate"`
	StudyInstanceUID string `bson:"studyinstanceuid"`
	NSRS             string `bson:"nsrs"`
	NSRI             string `bson:"nsri"`
	SDateIn          string `bson:"sdatein"`
	SDateOut         string `bson:"sdateout"`
}

// HarvProcess struct for elapsed time search
type HarvProcess struct {
	ID      int    `bson:"id"`
	DateBra string `bson:"datebra"`
	DateKet string `bson:"dateket"`
}

func (h *Harvest) getValues(msn string, keys []string) []string {

	var bracketO int
	var bracketC int
	var message string
	value := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		fTag := strings.Index(msn, keys[i])
		if fTag == -1 {
			value[i] = "no value"
			break
		}
		iTagEnd := fTag + len(keys[i]) - 1
		message = msn[iTagEnd:len(msn)]
		bracketO = strings.Index(message, "[")
		bracketO = bracketO + iTagEnd + 1
		bracketC = strings.Index(message, "]")
		bracketC = bracketC + iTagEnd
		value[i] = msn[bracketO:bracketC]
	}
	return value
}

func (h *Harvest) dataHarvest(strquery string, ses *mgo.Session, id int, in string, out string) int {
	strSlice := strings.Split(strquery, respFromPacs)
	Nstudies := len(strSlice) - 1
	var (
		msn       string
		outStr    []string
		idStudies int
	)

	idStudies = id
	session := ses.Copy()
	defer session.Close()
	// Collection
	c := session.DB(DATABASE).C(HCOLLECTION)

	for i := 1; i <= Nstudies; i++ {
		msn = strSlice[i]
		keys := []string{"(0008,0050)", "(0008,0060)", "(0010,0020)", "(0020,0010)", "(0020,1206)", "(0020,1208)", "(0008,1030)", "(0020,000D)"}
		outStr = h.getValues(msn, keys)

		harvest := Harvest{
			ID:               strconv.Itoa(idStudies),
			AccessionNumber:  outStr[0],
			StudyID:          outStr[3],
			PatientID:        outStr[2],
			Modality:         outStr[1],
			StudyDescription: outStr[6],
			StudyDate:        in,
			StudyInstanceUID: outStr[7],
			NSRS:             outStr[4],
			NSRI:             outStr[5],
			SDateIn:          in,
			SDateOut:         out,
		}

		//fmt.Println(harvest)

		// Insert
		c.Insert(harvest)

		logFunction("Harvest: ID per day: " + strconv.Itoa(i))
		logFunction("Harvest: General ID per Migration: " + harvest.ID)
		idStudies++
	}
	return idStudies
}

func harvestQuery(confFind ConfigJSON, ses *mgo.Session, str string) {
	var (
		command   string
		counter   = 0
		query     string
		idPointer = 0
	)

	f := TimeHelper{confFind.DateStart, confFind.DateEnd}
	indLast := f.getdates(ses)
	fmt.Println(indLast)
	session := ses.Copy()
	defer session.Close()

	var datesEl HarvProcess
	// Collection
	c := session.DB(DATABASE).C(HPROCESS)
	for i := 0; i < indLast; i++ {
		sleepFunc(confFind, str)
		c.Find(bson.M{"id": i}).One(&datesEl)
		inicioD := datesEl.DateBra
		fmt.Println("Ver recoleccion de data", inicioD)
		finD := ""
		c := ComQuery{"H", inicioD, finD, "StudyDate", ""}
		command = c.getCommands(confFind, "")
		counter++
		logFunction("Harvest: End of query " + strconv.Itoa(counter))
		logFunction("Harvest: " + command)
		fmt.Println(command)
		qr := queryDCM(command)
		query = qr

		// Data Harvest: Data Collection and send DB
		h := new(Harvest)
		pointerInd := h.dataHarvest(query, session, idPointer, inicioD, finD)
		fmt.Println("Harvest DB saved ", pointerInd)
		idPointer = pointerInd
		logFunction("Harvest: Harvest DB saved")

	}
}
