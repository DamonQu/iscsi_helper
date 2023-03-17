package main

import (
	"flag"
	"fmt"
	"hdatas/iscsi/iscsi"
	"hdatas/iscsi/util"
	"os"
)



func main() {
	fs := flag.NewFlagSet("iscsi helper", flag.ExitOnError)
	config := util.InitConfig(fs)

	if config.Debug {
		fmt.Println("Enable debug logging")
		iscsi.EnableDebugLogging(os.Stdout)
	}

	option := util.FirstPageHint()
	switch option {
	case util.CommandLineOption_List_Sessions:
		fmt.Println("List sessions")
		break

	case util.CommandLineOption_ISCSI_Login:
		fmt.Println("ISCSI login")
		break

	case util.CommandLineOption_ISCSI_Logout:
		fmt.Println("ISCSI logout")
		break

	case util.CommandLineOption_MOUNT:
		fmt.Println("mount")
		break

	default:
		os.Exit(0)
	}
}


//func iscsiLogin(config *util.Config)  {
//	targetIqnList, err := iscsi.Discovery(config.TargetAddress, "default")
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//
//	selected := 0
//	for selected <= 0 || selected > len(targetIqnList) {
//		fmt.Println("Discovered targets:")
//		for index, targetIqn := range targetIqnList {
//			fmt.Printf("  [%d] %s\n", index + 1, targetIqn)
//		}
//		fmt.Printf("Please select the iqn want to login: ")
//		 _, _ = fmt.Scanf("%d", &selected)
//		if selected <= 0 || selected > len(targetIqnList) {
//			fmt.Printf("[Error] invalid index of the target iqn, it should be [1, %d].\n", len(targetIqnList))
//		}
//	}
//	config.TargetIqn = targetIqnList[selected - 1]
//
//	fmt.Println("try to login to the iscsi target: ", config.TargetIqn)
//	connector := iscsi.Connector{
//		TargetIqn: config.TargetIqn,
//		TargetPortals: []string{config.TargetAddress},
//		Lun: 1}
//	device, err := connector.Connect()
//	if err != nil {
//		fmt.Printf("failed to connect the remote iqn: %s, error: %v\n", config.TargetIqn, err)
//		os.Exit(1)
//	}
//	fmt.Printf("Logging in the %s on device %s successfully.\n", config.TargetIqn, device)
//}