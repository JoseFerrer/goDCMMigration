package main

import (
	"fmt"
	"strconv"
	"time"

	mgo "gopkg.in/mgo.v2"
)

type ConfigJSON struct {
	PacsAETitle   string
	PacsIP        string
	PacsPort      string
	ENTITYAETitle string
	ENTITYIP      string
	ENTITYPort    string
	BackUpTitle   string
	NroTag        string
	DateStart     string
	DateEnd       string
	Monday        struct {
		StartHour    int
		StartMinutes int
		EndHour      int
		EndMinutes   int
	}
	Tuesday struct {
		StartHour    int
		StartMinutes int
		EndHour      int
		EndMinutes   int
	}
	Wednesday struct {
		StartHour    int
		StartMinutes int
		EndHour      int
		EndMinutes   int
	}
	Thursday struct {
		StartHour    int
		StartMinutes int
		EndHour      int
		EndMinutes   int
	}
	Friday struct {
		StartHour    int
		StartMinutes int
		EndHour      int
		EndMinutes   int
	}
	Saturday struct {
		StartHour    int
		StartMinutes int
		EndHour      int
		EndMinutes   int
	}
	Sunday struct {
		StartHour    int
		StartMinutes int
		EndHour      int
		EndMinutes   int
	}
}

// TimeHelper struct for tu use
type TimeHelper struct {
	StartD string
	StopD  string
}

func sleepFunc(confFind ConfigJSON, stage string) {
	var strnotification = ""

	dInit := confFind.getDuration(stage)
	fmt.Println(dInit.Minutes())
	if dInit.Minutes() > 0 {
		fmt.Printf("Sleep %f minutes before beginning the migration.\n", dInit.Minutes())
		logFunction("Sleep: Sleep program.")
		time.Sleep(dInit)
	}
	fmt.Println("Final sleep for harvest")
	logFunction("Sleep: Final sleep")

	loc, _ := time.LoadLocation("America/Lima")
	nowN := time.Now().In(loc)
	fmt.Printf("Start Migration at %d-%02d-%02dT%02d:%02d:%02d-00:00\n", nowN.Year(), nowN.Month(), nowN.Day(),
		nowN.Hour(), nowN.Minute(), nowN.Second())
	strnotification = strconv.Itoa(nowN.Hour()) + ":" + strconv.Itoa(nowN.Minute()) + ":" + strconv.Itoa(nowN.Second())
	logFunction("Start Migration at " + strnotification)
}

// Debe estar en el principal
func (m *ConfigJSON) getDuration(state string) (etime time.Duration) {
	var (
		/* mins         bool
		hours        bool */
		startMin   int
		startHour  int
		endHour    int
		endMinutes int
		/* nday         int
		compAMin     int
		compAHour    int
		compBMin     int
		compBHour    int */
		eTime1       time.Duration
		eTime2       time.Duration
		eTime3       time.Duration
		durationNext time.Duration
	)

	loc, _ := time.LoadLocation("America/Lima")
	now := time.Now().In(loc)
	fmt.Println("Time now: ", now)
	switch now.Weekday() {
	case time.Monday:
		startHour = m.Monday.StartHour
		startMin = m.Monday.StartMinutes
		endHour = m.Monday.EndHour
		endMinutes = m.Monday.EndMinutes
		dayin := time.Date(now.Year(), now.Month(), now.Day(), startHour, startMin, 0, 0, loc)
		dayout := time.Date(now.Year(), now.Month(), now.Day(), endHour, endMinutes, 0, 0, loc)
		nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, m.Tuesday.StartHour, m.Tuesday.StartMinutes, 0, 0, loc)
		eTime1 = dayin.Sub(now)
		eTime2 = dayout.Sub(now)
		eTime3 = nextDay.Sub(now)
		if eTime1 > 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = eTime1
		}
		if eTime1 < 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = eTime1
		}
		if eTime1 < 0 && eTime2 < 0 && eTime3 > 0 {
			durationNext = eTime3
		}
	case time.Tuesday:
		startHour = m.Tuesday.StartHour
		startMin = m.Tuesday.StartMinutes
		endHour = m.Tuesday.EndHour
		endMinutes = m.Tuesday.EndMinutes
		dayin := time.Date(now.Year(), now.Month(), now.Day(), startHour, startMin, 0, 0, loc)
		dayout := time.Date(now.Year(), now.Month(), now.Day(), endHour, endMinutes, 0, 0, loc)
		nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, m.Wednesday.StartHour, m.Wednesday.StartMinutes, 0, 0, loc)
		eTime1 = dayin.Sub(now)
		eTime2 = dayout.Sub(now)
		eTime3 = nextDay.Sub(now)
		if eTime1 > 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = eTime1
		}
		if eTime1 < 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = eTime2
		}
		if eTime1 < 0 && eTime2 < 0 && eTime3 > 0 {
			durationNext = eTime3
		}
	case time.Wednesday:
		startHour = m.Wednesday.StartHour
		startMin = m.Wednesday.StartMinutes
		endHour = m.Wednesday.EndHour
		endMinutes = m.Wednesday.EndMinutes
		dayin := time.Date(now.Year(), now.Month(), now.Day(), startHour, startMin, 0, 0, loc)
		dayout := time.Date(now.Year(), now.Month(), now.Day(), endHour, endMinutes, 0, 0, loc)
		nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, m.Thursday.StartHour, m.Thursday.StartMinutes, 0, 0, loc)
		eTime1 = dayin.Sub(now)
		eTime2 = dayout.Sub(now)
		eTime3 = nextDay.Sub(now)
		if eTime1 > 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = eTime1
			fmt.Println("Ingreso antes del tiempo")
		}
		if eTime1 < 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = -eTime2
			fmt.Println("Ingreso durante el tiempo")
		}
		if eTime1 < 0 && eTime2 < 0 && eTime3 > 0 {
			durationNext = eTime3
			fmt.Println("Ingreso despues del tiempo")
		}
	case time.Thursday:
		startHour = m.Thursday.StartHour
		startMin = m.Thursday.StartMinutes
		endHour = m.Thursday.EndHour
		endMinutes = m.Thursday.EndMinutes
		dayin := time.Date(now.Year(), now.Month(), now.Day(), startHour, startMin, 0, 0, loc)
		dayout := time.Date(now.Year(), now.Month(), now.Day(), endHour, endMinutes, 0, 0, loc)
		nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, m.Friday.StartHour, m.Friday.StartMinutes, 0, 0, loc)
		eTime1 = dayin.Sub(now)
		eTime2 = dayout.Sub(now)
		eTime3 = nextDay.Sub(now)
		if eTime1 > 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = eTime1
		}
		if eTime1 < 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = -eTime2
		}
		if eTime1 < 0 && eTime2 < 0 && eTime3 > 0 {
			durationNext = eTime3
		}
	case time.Friday:
		startHour = m.Friday.StartHour
		startMin = m.Friday.StartMinutes
		endHour = m.Friday.EndHour
		endMinutes = m.Friday.EndMinutes
		dayin := time.Date(now.Year(), now.Month(), now.Day(), startHour, startMin, 0, 0, loc)
		dayout := time.Date(now.Year(), now.Month(), now.Day(), endHour, endMinutes, 0, 0, loc)
		nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, m.Saturday.StartHour, m.Saturday.StartMinutes, 0, 0, loc)
		eTime1 = dayin.Sub(now)
		eTime2 = dayout.Sub(now)
		eTime3 = nextDay.Sub(now)
		if eTime1 > 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = eTime1
		}
		if eTime1 < 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = -eTime2
		}
		if eTime1 < 0 && eTime2 < 0 && eTime3 > 0 {
			durationNext = eTime3
		}
	case time.Saturday:
		startHour = m.Saturday.StartHour
		startMin = m.Saturday.StartMinutes
		endHour = m.Saturday.EndHour
		endMinutes = m.Saturday.EndMinutes
		dayin := time.Date(now.Year(), now.Month(), now.Day(), startHour, startMin, 0, 0, loc)
		dayout := time.Date(now.Year(), now.Month(), now.Day(), endHour, endMinutes, 0, 0, loc)
		nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, m.Sunday.StartHour, m.Sunday.StartMinutes, 0, 0, loc)
		eTime1 = dayin.Sub(now)
		eTime2 = dayout.Sub(now)
		eTime3 = nextDay.Sub(now)
		if eTime1 > 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = eTime1
		}
		if eTime1 < 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = -eTime2
		}
		if eTime1 < 0 && eTime2 < 0 && eTime3 > 0 {
			durationNext = eTime3
		}
	case time.Sunday:
		startHour = m.Sunday.StartHour
		startMin = m.Sunday.StartMinutes
		endHour = m.Sunday.EndHour
		endMinutes = m.Sunday.EndMinutes
		dayin := time.Date(now.Year(), now.Month(), now.Day(), startHour, startMin, 0, 0, loc)
		dayout := time.Date(now.Year(), now.Month(), now.Day(), endHour, endMinutes, 0, 0, loc)
		nextDay := time.Date(now.Year(), now.Month(), now.Day()+1, m.Monday.StartHour, m.Monday.StartMinutes, 0, 0, loc)
		eTime1 = dayin.Sub(now)
		eTime2 = dayout.Sub(now)
		eTime3 = nextDay.Sub(now)
		if eTime1 > 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = eTime1
		}
		if eTime1 < 0 && eTime2 > 0 && eTime3 > 0 {
			durationNext = -eTime2
		}
		if eTime1 < 0 && eTime2 < 0 && eTime3 > 0 {
			durationNext = eTime3
		}
	}

	/* mins = false
	hours = false

	compAMin = now.Minute()
	compAHour = now.Hour()
	switch state {
	case "Initial":
		compBMin = startMin
		compBHour = startHour
	case "Middle":
		compBMin = endMinutes
		compBHour = endHour
	case "NextD":
		return durationNext
	}

	if compAMin+compBMin >= 60 {
		mins = true
	}
	if mins == true {
		if compAHour+compBHour+1 >= 24 {
			hours = true
		}
	} else {
		if compAHour+compBHour >= 24 {
			hours = true
		}
	}
	if hours == true {
		nday = now.Day() + 1
	} else {
		nday = now.Day()
	}
	s := time.Date(now.Year(), now.Month(), nday, compAHour, compAMin, 0, 0, loc)
	e := time.Date(now.Year(), now.Month(), nday, compBHour, compBMin, 0, 0, loc)

	duration := e.Sub(s)
	return duration */

	return durationNext
}

func time2String(newtime time.Time) string {
	var (
		fecha = ""
		mes   = ""
	)

	mes = strconv.Itoa(int(newtime.Month()))
	if len(mes) == 1 {
		mes = "0" + mes
	}

	day := strconv.Itoa(newtime.Day())

	if len(day) > 1 {
		fecha = strconv.Itoa(newtime.Year()) + mes + day
	} else {
		fecha = strconv.Itoa(newtime.Year()) + mes + "0" + day
	}
	return fecha
}

func str2Time(d2Time string) time.Time {
	loc, _ := time.LoadLocation("America/Lima")
	anho, _ := strconv.Atoi(d2Time[0:4])
	m, _ := strconv.Atoi(d2Time[4:6])
	dia, _ := strconv.Atoi(d2Time[6:8])

	return time.Date(anho, time.Month(m), dia, 0, 0, 0, 0, loc)
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start.Add(-24*time.Hour)) && check.Before(end.Add(24*time.Hour))
}

// getdates extract the elapsed dates for query c-find
func (f TimeHelper) getdates(ses *mgo.Session) int {

	var varDay time.Time

	start := str2Time(f.StartD)
	end := str2Time(f.StopD)
	diff := end.Sub(start)

	varDay = start
	session := ses.Copy()
	defer session.Close()
	c := session.DB(DATABASE).C(HPROCESS)
	for i := 0; i <= int(diff.Hours()/24); i++ {
		timeDay := time2String(varDay)
		varDay = varDay.Add(24 * time.Hour)
		elap := HarvProcess{
			ID:      i,
			DateBra: timeDay,
			DateKet: "",
		}
		c.Insert(elap)
		fmt.Println(elap)
	}
	return int(diff.Hours() / 24)
}
