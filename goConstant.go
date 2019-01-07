package main

const (

	// StudyID : tag dicom of study identification
	StudyID = "-r 00200010"
	// AccessionNum : tag dicom of accession number
	AccessionNum = "-r 00080050"
	// PatientID : tag dicom of patient identification
	PatientID = "-r 00100020"
	// StudyDesc :  tag dicom of study description
	StudyDesc = "-r 00081030"
	// Modality : tag dicom of modality
	Modality = "-r 00080060"
	// StudyDate :  tag dicom of study date
	StudyDate = "-r 00080020"
	// StudyTime : tag dicom of study time
	StudyTime = "-r 00080030"
	// NStudyRS : tag dicom of number of studies related series
	NStudyRS = "-r 00201206"
	// NStudyRI : tag dicom of number of studies related instances
	NStudyRI = "-r 00201208"
	// respFromPacs Positive response from PACS
	respFromPacs = "status=ff00H"
)
