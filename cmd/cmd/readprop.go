package cmd

import (
	"fmt"
	"github.com/NubeDev/bacnet"
	"github.com/NubeDev/bacnet/btypes"
	pprint "github.com/NubeDev/bacnet/helpers/print"
	"github.com/NubeDev/bacnet/local"
	"github.com/spf13/cobra"
	"strconv"
)

// Flags
var (
	networkNumber     int
	deviceID          int
	deviceIP          string
	devicePort        int
	deviceHardwareMac int
	objectID          int
	objectType        int
	arrayIndex        uint32
	propertyType      string
	listProperties    bool
	segmentation      int
	maxADPU           int
	getDevicePoints   bool
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Prints out a device's object's property",
	Long: `
 Given a device's object instance and selected property, we print the value
 stored there. There are some autocomplete features to try and minimize the
 amount of arguments that need to be passed, but do take into consideration
 this discovery process may cause longer reads.
	`,
	Run: readProp,
}

func readProp(cmd *cobra.Command, args []string) {

	localDevice, err := local.New(&local.Local{Interface: Interface, Port: Port})
	if err != nil {
		fmt.Println("ERR-client", err)
		return
	}
	defer localDevice.ClientClose()
	go localDevice.ClientRun()

	device, err := local.NewDevice(localDevice, &local.Device{Ip: deviceIP, DeviceID: deviceID, NetworkNumber: networkNumber, MacMSTP: deviceHardwareMac, MaxApdu: uint32(maxADPU), Segmentation: uint32(segmentation)})
	if err != nil {
		return
	}

	if getDevicePoints {
		points, err := device.GetDevicePoints(btypes.ObjectInstance(deviceID))
		if err != nil {
			return
		}
		pprint.PrintJOSN(points)
		return
	}

	var propInt btypes.PropertyType
	// Check to see if an int was passed
	if i, err := strconv.Atoi(propertyType); err == nil {
		propInt = btypes.PropertyType(uint32(i))
	} else {
		propInt, err = btypes.Get(propertyType)
	}

	obj := &local.Object{
		ObjectID:   btypes.ObjectInstance(objectID),
		ObjectType: btypes.ObjectType(objectType),
		Prop:       propInt,
		ArrayIndex: arrayIndex, //btypes.ArrayAll

	}
	read, err := device.Read(obj)
	fmt.Println(err)
	fmt.Println(read)
}
func init() {
	// Descriptions are kept separate for legibility purposes.
	propertyTypeDescr := `type of read that will be done. Support both the
	property type as an integer or as a string. e.g. PropObjectName or 77 are both
	support. Run --list to see available properties.`
	listPropertiesDescr := `list all string versions of properties that are
	support by property flag`

	RootCmd.AddCommand(readCmd)

	// Pass flags to children
	readCmd.PersistentFlags().IntVarP(&deviceID, "device", "", 202, "device id")
	readCmd.Flags().StringVarP(&deviceIP, "address", "", "192.168.15.202", "device ip")
	readCmd.Flags().IntVarP(&devicePort, "dport", "", 47808, "device port")
	readCmd.Flags().IntVarP(&networkNumber, "network", "", 0, "bacnet network number")
	readCmd.Flags().IntVarP(&deviceHardwareMac, "mstp", "", 0, "device hardware mstp addr")
	readCmd.Flags().IntVarP(&maxADPU, "adpu", "", 0, "device max adpu")
	readCmd.Flags().IntVarP(&segmentation, "seg", "", 0, "device segmentation")
	readCmd.Flags().IntVarP(&objectID, "objectID", "", 202, "object ID")
	readCmd.Flags().IntVarP(&objectType, "objectType", "", 8, "object type")
	readCmd.Flags().StringVarP(&propertyType, "property", "", btypes.ObjectNameStr, propertyTypeDescr)
	readCmd.Flags().Uint32Var(&arrayIndex, "index", bacnet.ArrayAll, "Which position to return.")

	readCmd.PersistentFlags().BoolVarP(&listProperties, "list", "l", false, listPropertiesDescr)

	readCmd.PersistentFlags().BoolVarP(&getDevicePoints, "device-points", "", false, "get device points list")
}
