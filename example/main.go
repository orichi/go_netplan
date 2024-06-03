package main

import (
	"fmt"
	"log"

	netplan "github.com/orichi/go_netplan"
)

func main() {
	// 自动初始化网口名称
	interfaces, err := netplan.GetInterfaces()
	if err != nil {
		log.Fatalf("Failed to get network interfaces: %v", err)
	}
	if len(interfaces) == 0 {
		log.Fatalf("No network interfaces found")
	}
	fmt.Printf("Detected interfaces: %v\n", interfaces)

	// Load the netplan configuration
	config, err := netplan.LoadConfig("/etc/netplan/01-netcfg.yaml")
	if err != nil {
		log.Fatalf("Failed to load netplan configuration: %v", err)
	}

	// 使用第一个接口（假设存在）来演示
	iface := interfaces[0]

	// Add a virtual IP to iface
	err = config.AddVirtualIP(iface, "192.168.1.100/24")
	if err != nil {
		log.Fatalf("Failed to add virtual IP: %v", err)
	}

	// Modify the IP address for iface
	err = config.ModifyIP(iface, "192.168.1.100/24", "192.168.1.101/24")
	if err != nil {
		log.Fatalf("Failed to modify IP address: %v", err)
	}

	// Add a gateway to iface
	err = config.AddGateway(iface, "0.0.0.0/0", "192.168.1.1")
	if err != nil {
		log.Fatalf("Failed to add gateway: %v", err)
	}

	// Modify DNS settings for iface
	err = config.ModifyDNS(iface, []string{"8.8.8.8", "8.8.4.4"})
	if err != nil {
		log.Fatalf("Failed to modify DNS settings: %v", err)
	}

	// Save the updated configuration
	err = netplan.SaveConfig("/etc/netplan/01-netcfg.yaml", config)
	if err != nil {
		log.Fatalf("Failed to save netplan configuration: %v", err)
	}

	// Apply the new configuration
	err = netplan.ApplyConfig()
	if err != nil {
		log.Fatalf("Failed to apply netplan configuration: %v", err)
	}

	fmt.Println("Netplan configuration updated successfully")
}
