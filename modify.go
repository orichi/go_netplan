package netplan

import "fmt"

// AddVirtualIP adds a virtual IP to the specified interface
func (np *NetPlan) AddVirtualIP(iface string, ip string) error {
	if eth, exists := np.Network.Ethernets[iface]; exists {
		eth.Addresses = append(eth.Addresses, ip)
		np.Network.Ethernets[iface] = eth
		return nil
	}
	return fmt.Errorf("interface %s not found", iface)
}

// RemoveVirtualIP removes a virtual IP from the specified interface
func (np *NetPlan) RemoveVirtualIP(iface string, ip string) error {
	if eth, exists := np.Network.Ethernets[iface]; exists {
		addresses := eth.Addresses
		for i, addr := range addresses {
			if addr == ip {
				eth.Addresses = append(addresses[:i], addresses[i+1:]...)
				np.Network.Ethernets[iface] = eth
				return nil
			}
		}
		return fmt.Errorf("IP address %s not found on interface %s", ip, iface)
	}
	return fmt.Errorf("interface %s not found", iface)
}

// AddGateway adds a gateway to the specified interface
func (np *NetPlan) AddGateway(iface string, to string, via string) error {
	if eth, exists := np.Network.Ethernets[iface]; exists {
		route := Route{
			To:  to,
			Via: via,
		}
		eth.Routes = append(eth.Routes, route)
		np.Network.Ethernets[iface] = eth
		return nil
	}
	return fmt.Errorf("interface %s not found", iface)
}

// ModifyDNS modifies the DNS settings for the specified interface
func (np *NetPlan) ModifyDNS(iface string, dns []string) error {
	if eth, exists := np.Network.Ethernets[iface]; exists {
		eth.Nameservers.Addresses = dns
		np.Network.Ethernets[iface] = eth
		return nil
	}
	return fmt.Errorf("interface %s not found", iface)
}

// ModifyIP modifies the IP address for the specified interface
func (np *NetPlan) ModifyIP(iface string, oldIP string, newIP string) error {
	if eth, exists := np.Network.Ethernets[iface]; exists {
		for i, addr := range eth.Addresses {
			if addr == oldIP {
				eth.Addresses[i] = newIP
				np.Network.Ethernets[iface] = eth
				return nil
			}
		}
		return fmt.Errorf("IP address %s not found on interface %s", oldIP, iface)
	}
	return fmt.Errorf("interface %s not found", iface)
}
