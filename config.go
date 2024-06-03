package netplan

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// loadConfig loads the netplan configuration from the specified file
func loadConfig(filepath string) (*NetPlan, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var config NetPlan
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// saveConfig saves the netplan configuration to the specified file
func saveConfig(filepath string, config *NetPlan) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
