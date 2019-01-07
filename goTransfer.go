package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Moved struct
type Moved struct {
	ID               string
	IDH              string
	AccNum           string
	StudyID          string
	StudyInstanceUID string
}

func transferQuery(conf ConfigJSON, sess *mgo.Session, str string) {
	var (
		command  string
		query    string
		countIn  = 0
		initialI = 1
	)

	logFunction("Transfer: Start transfer process")
	session := sess.Copy()
	defer session.Close()
	// Collection HCOLLECTION
	col := session.DB(DATABASE).C(HCOLLECTION)
	dbSize, _ := col.Count()

	// Collection TPROCESS
	colT := session.DB(DATABASE).C(TCOLLECTIOM)
	tdbSize, _ := colT.Count()
	fmt.Printf("Size of transfer %d /n", tdbSize)

	if tdbSize != 0 {
		TSearchid := Moved{}
		colT.Find(bson.M{"id": strconv.Itoa(tdbSize - 1)}).One(&TSearchid)
		initialI, _ = strconv.Atoi(TSearchid.IDH)
		countIn, _ = strconv.Atoi(TSearchid.ID)
		countIn++
		initialI++
	}

	//fmt.Println("The size of studies to move are ", dbSize, ".")
	//logFunction("Transfer: The size of studies to move are " + strconv.Itoa(dbSize))

	fmt.Println("The last id is: ", initialI, ".")
	logFunction("Transfer: The last  " + strconv.Itoa(initialI))

	for i := initialI; i <= dbSize; i++ {
		fromHarvest := Harvest{}
		col.Find(bson.M{"id": strconv.Itoa(i)}).One(&fromHarvest)
		Acc := fromHarvest.AccessionNumber
		//fmt.Println("The Accession Number is: ", Acc)

		c := ComQuery{"T", "", "", conf.NroTag, Acc}
		if Acc != "" {
			command = c.getCommands(conf, Acc)
			logFunction("Transfer: command for terminal " + command)
			fmt.Println(command)

			// Se puso a proposito
			// fmt.Println("Start Sleep")
			// time.Sleep(10 * time.Minute)
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Do you want to continue (Yes/No) (Y/N): ")
			text, _ := reader.ReadString('\n')
			if text == "N" {
				return
			}

			query = queryDCM(command)
			fmt.Println(query)
			colTrans := session.DB(DATABASE).C(TCOLLECTIOM)
			TransData := Moved{
				ID:               strconv.Itoa(countIn),
				IDH:              fromHarvest.ID,
				AccNum:           Acc,
				StudyID:          fromHarvest.StudyID,
				StudyInstanceUID: fromHarvest.StudyInstanceUID,
			}
			colTrans.Insert(TransData)
			countIn++
		}
	}
}
