package btypes

import "fmt"

/*
Device Services Supported
Type bitSting
This property indicates which standardized protocol services are supported by this device's protocol implementation.
*/

type ServicesSupported struct {
	Name   string
	Number uint16
}

var acknowledgeAlarm = ServicesSupported{
	Name:   "acknowledgeAlarm",
	Number: 0,
}

var confirmedCOVNotification = ServicesSupported{
	Name:   "confirmedCOVNotification",
	Number: 1,
}

var supportedList = map[ServicesSupported]string{
	acknowledgeAlarm:         acknowledgeAlarm.Name,
	confirmedCOVNotification: confirmedCOVNotification.Name,
}

func (support ServicesSupported) ListAll() map[ServicesSupported]string {
	return supportedList
}

func (support ServicesSupported) GetType(s string) *ServicesSupported {
	for typ, str := range supportedList {
		if s == str {
			return &typ
		}
	}
	return nil

}

func (support ServicesSupported) GetString(t ServicesSupported) string {
	s, ok := supportedList[t]
	if !ok {
		return fmt.Sprintf("Unknown (%s)", t.Name)
	}
	return fmt.Sprintf("%s", s)
}

//protocolServicesSupported	97
//bitString
const (
	//acknowledgeAlarm           = 0
	//confirmedCOVNotification   = 1
	confirmedEventNotification = 2
	getAlarmSummary            = 3
	getEnrollmentSummary       = 4
	subscribeCOV               = 5
	atomicReadFile             = 6
	atomicWriteFile            = 7
	addListElement             = 8
	removeListElement          = 9
	createObject               = 10
	deleteObject               = 11
	readProperty               = 12
	//readPropertyConditional':13      # removed in version 1 revision 12
	readPropertyMultiple       = 14
	writeProperty              = 15
	writePropertyMultiple      = 16
	deviceCommunicationControl = 17
	confirmedPrivateTransfer   = 18
	confirmedTextMessage       = 19
	reinitializeDevice         = 20
	vtOpen                     = 21
	vtClose                    = 22
	vtData                     = 23
	//# , 'authenticate':24                 # removed in version 1 revision 11
	//# , 'requestKey':25                   # removed in version 1 revision 11
	iAm                          = 26
	iHave                        = 27
	unconfirmedCOVNotification   = 28
	unconfirmedEventNotification = 29
	unconfirmedPrivateTransfer   = 30
	unconfirmedTextMessage       = 31
	timeSynchronization          = 32
	whoHas                       = 33
	whoIs                        = 34
	readRange                    = 35
	utcTimeSynchronization       = 36
	lifeSafetyOperation          = 37
	subscribeCOVProperty         = 38
	getEventInformation          = 39
	writeGroup                   = 40
)