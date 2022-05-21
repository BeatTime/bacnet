package local

import (
	"fmt"
	"github.com/NubeDev/bacnet/btypes"
	"github.com/NubeDev/bacnet/helpers/data"
	log "github.com/sirupsen/logrus"
)

//DeviceObjects get device objects
func (device *Device) DeviceObjects(deviceID btypes.ObjectInstance) (objectList []btypes.ObjectID, err error) {
	//get object list
	obj := &Object{
		ObjectID:   deviceID,
		ObjectType: btypes.DeviceType,
		Prop:       btypes.PropObjectList,
		ArrayIndex: btypes.ArrayAll, //btypes.ArrayAll

	}
	out, err := device.Read(obj)
	if err != nil { //this is a device that would have a low maxADPU
		if out.Object.Properties[0].Type == btypes.PropObjectList {
			log.Error("Note: PropObjectList reads may need to be broken up into multiple reads due to length. Read index 0 for array length")
		}
		return device.deviceObjectsBuilder(deviceID)

	}
	if len(out.Object.Properties) == 0 {
		fmt.Println("No value returned")
		return nil, nil
	}
	_, ids := data.ToArr(out)
	for _, id := range ids {
		objectID := id.(btypes.ObjectID)
		objectList = append(objectList, objectID)
	}
	return objectList, nil
}

//DeviceObjectsBuilder this is used when a device can't send the object list in the fully ArrayIndex
//it first reads the size of the object list and then loops the list to build an object list
func (device *Device) deviceObjectsBuilder(deviceID btypes.ObjectInstance) (objectList []btypes.ObjectID, err error) {
	//get object list
	obj := &Object{
		ObjectID:   deviceID,
		ObjectType: btypes.DeviceType,
		Prop:       btypes.PropObjectList,
		ArrayIndex: 0, //start at 0 and then loop through
	}
	out, err := device.Read(obj)
	if err != nil {
		return nil, err
	}
	_, o := data.ToUint32(out)
	var listLen = int(o)
	for i := 1; i <= listLen; i++ {
		obj.ArrayIndex = uint32(i)
		out, _ := device.Read(obj)
		objectID := out.Object.Properties[0].Data.(btypes.ObjectID)
		objectList = append(objectList, objectID)
	}
	return objectList, err

}
