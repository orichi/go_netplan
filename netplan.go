package netplan

import (
	"fmt"
	"os/exec"
	"strings"
)

// NetPlan represents the netplan configuration
type NetPlan struct {
	Network struct {
		Version   int                 `yaml:"version"`
		Ethernets map[string]Ethernet `yaml:"ethernets"`
	} `yaml:"network"`
}

// Ethernet represents an ethernet interface configuration
type Ethernet struct {
	Addresses   []string   `yaml:"addresses,omitempty"`
	Routes      []Route    `yaml:"routes,omitempty"`
	Nameservers Nameserver `yaml:"nameservers,omitempty"`
}

// Route represents a network route configuration
type Route struct {
	To  string `yaml:"to"`
	Via string `yaml:"via"`
}

// Nameserver represents DNS configuration
type Nameserver struct {
	Addresses []string `yaml:"addresses,omitempty"`
}

// LoadConfig loads the netplan configuration from the specified file
func LoadConfig(filepath string) (*NetPlan, error) {
	return loadConfig(filepath)
}

// SaveConfig saves the netplan configuration to the specified file
func SaveConfig(filepath string, config *NetPlan) error {
	return saveConfig(filepath, config)
}

// ApplyConfig applies the netplan configuration using 'netplan apply'
func ApplyConfig() error {
	cmd := exec.Command("sudo", "netplan", "apply")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to apply netplan configuration: %s", string(output))
	}
	return nil
}

// GetInterfaces returns the names of network interfaces available on the system
func GetInterfaces() ([]string, error) {
	cmd := exec.Command("ls", "/sys/class/net")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to list network interfaces: %s", string(output))
	}
	interfaces := strings.Fields(string(output))
	return interfaces, nil
}
