package main

import (
	"flag"
	"fmt"
	"hdatas/iscsi/iscsi"
	"hdatas/iscsi/util"
	"os"
)

const iscsiDefaultLun = 1

func main() {
	fs := flag.NewFlagSet("iscsi helper", flag.ExitOnError)
	config := util.InitConfig(fs)

	if config.Debug {
		fmt.Println("Enable debug logging")
		iscsi.EnableDebugLogging(os.Stdout)
	}

	showOptions()
}

func showOptions()  {
	option := util.FirstPageHint()
	switch option {
	case util.CommandLineOption_List_Sessions:
		fmt.Println()
		listIscsiSessions()
		break

	case util.CommandLineOption_ISCSI_Login:
		fmt.Println()
		iscsiLogin()
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
	showOptions()
}

func listIscsiSessions() {
	iscsiSessions, err := iscsi.GetCurrentSessions()
	if err != nil {
		fmt.Printf("failed to list the ISCSI sessions. error: %v\n", err)
		os.Exit(1)
	}

	if len(iscsiSessions) == 0 {
		fmt.Println("No session found.")
		showOptions()
	}

	fmt.Println("Current ISCSI session: ")
	for _, session := range iscsiSessions {
		fmt.Println("  " + session.ToString())
	}
}

func iscsiLogin()  {
	targetAddress := util.Hint("", "ISCSI IP Address: ", nil, "%s")

	targetIqnList, err := iscsi.Discovery(targetAddress, "default")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	targetIqn := util.Hint("Discovered targets:", "Please select the iqn want to login:", targetIqnList, "%d")

	connector := iscsi.Connector{
		TargetIqn:     targetIqn,
		TargetPortals: []string{targetAddress},
		Lun:           iscsiDefaultLun}
	device, err := connector.Connect()
	if err != nil {
		fmt.Printf("failed to connect the remote iqn: %s, error: %v\n", targetIqn, err)
		os.Exit(1)
	}
	fmt.Printf("Logging in the %s on device %s successfully.\n", targetIqn, device)
}