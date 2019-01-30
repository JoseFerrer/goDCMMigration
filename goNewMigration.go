package main

import (
	"fmt"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func newDCMMigration(confFind ConfigJSON, ses *mgo.Session) {

	var (
		command string
		query   string
	)

	logFunction("Quality: Start new move DICOM process")
	session := ses.Copy()
	defer session.Close()
	// Collection TCOLLECTION
	col := session.DB(DATABASE).C(TCOLLECTION)
	dbSize, _ := col.Count()

	for i := 0; i <= dbSize; i++ {
		sleepFunc(confFind, "Middle")
		fromTrasnfer := Moved{}
		col.Find(bson.M{"id": strconv.Itoa(i)}).One(&fromTrasnfer)
		Acc := fromTrasnfer.AccNum

		c := ComQuery{"T", "", "", confFind.NroTag, Acc}
		command = c.getCommands(confFind, Acc)
		logFunction("Quality: command for terminal " + command)
		fmt.Println(command)
		query = queryDCM(command)
		fmt.Println(query)
	}

}
