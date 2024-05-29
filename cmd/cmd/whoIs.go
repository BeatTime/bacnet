package cmd

import (
	"fmt"
	"github.com/BeatTime/bacnet"
	"github.com/BeatTime/bacnet/btypes"
	pprint "github.com/BeatTime/bacnet/helpers/print"
	"github.com/spf13/cobra"
	"time"
)

// Flags
var startRange int
var endRange int

var outputFilename string

// whoIsCmd represents the whoIs command
var whoIsCmd = &cobra.Command{
	Use:   "whois",
	Short: "BACnet device discovery",
	Long: `whoIs does a bacnet network discovery to find devices in the network
 given the provided range.`,
	Run: main,
}

func main(cmd *cobra.Command, args []string) {

	client, err := bacnet.NewClient(&bacnet.ClientBuilder{
		Interface:  Interface,
		Ip:         "10.245.3.254",
		Port:       Port,
		SubnetCIDR: 24,
	})
	defer client.ClientClose(false)
	go client.ClientRun()

	//networkInstance, err := network.New(&network.Network{Interface: Interface, Port: Port})
	//if err != nil {
	//	fmt.Println("ERR-networkInstance", err)
	//	return
	//}
	//defer networkInstance.NetworkClose(false)
	//go networkInstance.NetworkRun()

	if runDiscover {
		//device, err := network.NewDevice(networkInstance, &network.Device{Ip: deviceIP, Port: Port})
		//err = device.DeviceDiscover()
		fmt.Println(err)
		return
	}

	wi := &bacnet.WhoIsOpts{
		High:            endRange,
		Low:             startRange,
		GlobalBroadcast: true,
		NetworkNumber:   uint16(networkNumber),
	}

	whoIs, err := client.WhoIs(wi)
	if err != nil {
		fmt.Println("ERR-whoIs", err)
		return
	}

	pprint.PrintJOSN(whoIs)

	whoIs, err = client.WhoIs(wi)
	if err != nil {
		fmt.Println("ERR-whoIs", err)
		return
	}
	fmt.Println("whois 2nd")
	pprint.PrintJOSN(whoIs)

	tmp := btypes.PropertyData{
		Object: btypes.Object{
			ID: btypes.ObjectID{
				Type:     btypes.ObjectType(btypes.TypeAnalogInput),
				Instance: btypes.ObjectInstance(0),
			},
			Properties: []btypes.Property{
				{
					Type:       btypes.PropPresentValue, // Present value
					ArrayIndex: btypes.ArrayAll,
				},
			},
		},
	}
	timer := time.NewTicker(5 * time.Second)
	for {
		<-timer.C
		property, err := client.ReadProperty(whoIs[1], tmp)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(property.Object.Properties[0].Data)
		}
	}
}

func init() {
	RootCmd.AddCommand(whoIsCmd)
	whoIsCmd.Flags().BoolVar(&runDiscover, "discover", false, "run network discover")
	whoIsCmd.Flags().IntVarP(&startRange, "start", "s", -1, "Start range of discovery")
	whoIsCmd.Flags().IntVarP(&endRange, "end", "e", int(0xBAC0), "End range of discovery")
	whoIsCmd.Flags().IntVarP(&networkNumber, "network", "", 0, "network number")
	whoIsCmd.Flags().StringVarP(&outputFilename, "out", "o", "", "Output results into the given filename in json structure.")
}
