package network

import (
	"fmt"
	"github.com/NubeDev/bacnet/btypes"
	"testing"
)

func TestDevice_WritePointName(t *testing.T) {

	localDevice, err := New(&Network{Interface: iface, Port: 47809})
	if err != nil {
		fmt.Println("ERR-client", err)
		return
	}
	defer localDevice.NetworkClose()
	go localDevice.NetworkRun()

	device, err := NewDevice(localDevice, &Device{Ip: deviceIP, DeviceID: deviceID})
	if err != nil {
		return
	}

	pnt := &Point{
		ObjectID:   1,
		ObjectType: btypes.AnalogInput,
	}

	err = device.WritePointName(pnt, "new-name")
	fmt.Println(err)
	if err != nil {
		//return
	}

	read, err := device.ReadPointName(pnt)
	fmt.Println(err)
	if err != nil {
	}
	fmt.Println(read, err)

	//re

}