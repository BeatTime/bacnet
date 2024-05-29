package cmd

import (
	"fmt"
	"github.com/BeatTime/bacnet"
	"github.com/spf13/cobra"
	"time"
)

var (
	instanceId uint32
)

var iam = &cobra.Command{
	Use:   "iam",
	Short: "BACnet device iam",
	Long: `whoIs does a bacnet network discovery to find devices in the network
 given the provided range.`,
	Run: iamFunc,
}

func iamFunc(cmd *cobra.Command, args []string) {
	client, err := bacnet.NewClient(&bacnet.ClientBuilder{
		//Interface:  Interface,
		Ip:               "0.0.0.0",
		Port:             47808,
		SubnetCIDR:       24,
		DeviceInstanceId: instanceId,
	})
	if err != nil {
		return
	}
	defer client.ClientClose(false)
	go client.ClientRun()

	ticker := time.NewTicker(1 * time.Second)
	for {
		<-ticker.C
		fmt.Println("tick")
	}
}

func init() {
	RootCmd.AddCommand(iam)
	//iam.Flags().IntVarP(&endRange, "instance", "s", 718, "ID")
	iam.PersistentFlags().Uint32VarP(&instanceId, "device", "d", 718, "device id")
}
