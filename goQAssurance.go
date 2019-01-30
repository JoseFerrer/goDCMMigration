package main

import (
	"fmt"
	"strconv"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// QualityS struct
type QualityS struct {
	ID     int    `bson:"id"`
	Access string `bson:"accessionnumber"`
}

// findData struct
type findData struct {
	IDH             string
	IDT             string
	AccessionNumber string
}

func (h *Harvest) getValuesQ(msn string, keys []string) []string {

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

func (h *Harvest) dataQuality(strquery string, ses *mgo.Session, id int) int {
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
	c := session.DB(DATABASE).C(QTYFIND)

	for i := 1; i <= Nstudies; i++ {
		msn = strSlice[i]
		keys := []string{"(0008,0050)", "(0008,0060)", "(0010,0020)", "(0020,0010)", "(0020,1206)", "(0020,1208)", "(0008,1030)", "(0020,000D)", "(0008,0020)"}
		outStr = h.getValuesQ(msn, keys)

		harvest := Harvest{
			ID:               strconv.Itoa(idStudies),
			AccessionNumber:  outStr[0],
			StudyID:          outStr[3],
			PatientID:        outStr[2],
			Modality:         outStr[1],
			StudyDescription: outStr[6],
			StudyDate:        outStr[8],
			StudyInstanceUID: outStr[7],
			NSRS:             outStr[4],
			NSRI:             outStr[5],
			SDateIn:          outStr[8],
			SDateOut:         "",
		}

		//fmt.Println(harvest)

		// Insert
		c.Insert(harvest)

		logFunction("Quality: ID per day: " + strconv.Itoa(i))
		logFunction("Quality: General ID per Migration: " + harvest.ID)
		idStudies++
	}
	return idStudies
}

func qualityQuery(confFind ConfigJSON, ses *mgo.Session, str string) {
	var (
		command     string
		counterMig  = 0
		counterNMig = 0
		query       string
		idPointer   = 0
	)

	session := ses.Copy()
	defer session.Close()
	// Collection TCOLLECTION
	col := session.DB(DATABASE).C(TCOLLECTION)
	dbSize, _ := col.Count()
	fmt.Println("The number of studies to verify are : ", dbSize)

	var trData Moved
	// Collection
	d := session.DB(DATABASE).C(QTYNOTFIND)
	for i := 0; i < dbSize; i++ {
		sleepFunc(confFind, str)
		col.Find(bson.M{"id": i}).One(&trData)
		aNumber := trData.AccNum
		c := ComQuery{"Q", "", "", "00080050", aNumber}
		command = c.getCommands(confFind, "")
		fmt.Println(command)
		return

		qr := queryDCM(command)
		if !(strings.Contains(qr, "status=ff00H")) {
			fmt.Println("Study with Accession Number: " + aNumber + " not found in Back Up")
			logFunction("Quality: Study with Accession Number: " + aNumber + " not found in Back Up")
			notfind := findData{
				IDH:             trData.IDH,
				IDT:             trData.ID,
				AccessionNumber: aNumber,
			}
			d.Insert(notfind)
			counterNMig++
		} else {
			query = qr
			h := new(Harvest)
			pointerInd := h.dataQuality(query, session, idPointer)
			fmt.Println("Quality DB saved ", pointerInd)
			counterMig++
		}

	}
}
