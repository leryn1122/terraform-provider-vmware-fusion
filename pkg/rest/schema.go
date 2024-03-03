package rest

const (
	VMWareAPI_GetAllVirtualNetworks               = "/vmnet"
	VMWareAPI_GetAllMacToIpSettingsForDhcpService = "/vmnet/{vmnet}/mactoip"
	VMWareAPI_GetAllPortForwardings               = "/vmnet/{vmnet}/portforward"
	VMWareAPI_UpdateMacToIpBinding                = "/vmnet/{vmnet}/mactoip/{mac}"
	VMWareAPI_UpdatePortForwarding                = "/vmnet/{vmnet}/portforward/{protocol}/{port}"
	VMWareAPI_CreateVirtualNetwork                = "/vmnets"
	VMWareAPI_DeletePortForwarding                = "/vmnet/{vmnet}/portforward/{protocol}/{port}"

	VMWareAPI_GetVmList                    = "/vms"
	VMWareAPI_GetVm                        = "/vms/{id}"
	VMWareAPI_GetVmConfigParams            = "/vms/{id}/params/{name}"
	VMWareAPI_GetVmRestrictionsInformation = "/vms/{id}/restrictions"
	VMWareAPI_UpdateVmSetting              = "/vms/{id}"
	VMWareAPI_UpdateVmConfigParams         = "/vms/{id}/params"
	VMWareAPI_CreateVmCopy                 = "/vms"
	VMWareAPI_RegisterVmToVmLibrary        = "/vms/registration"
	VMWareAPI_DeleteVm                     = "/vms/{id}"

	VMWareAPI_GetVmIpAddress                   = "/vms/{id}/ip"
	VMWareAPI_GetAllVmNetworkAdapters          = "/vms/{id}/nic"
	VMWareAPI_GetIpStackConfigurationOfAllNics = "/vms/{id}/nicips"
	VMWareAPI_UpdateVmNetworkAdapter           = "/vms/{id}/nic/{index}"
	VMWareAPI_CreateVmNetworkAdapter           = "/vms/{id}/nic"
	VMWareAPI_DeleteVmNetworkAdapter           = "/vms/{id}/nic/{index}"

	VMWareAPI_GetVmPowerState    = "/vms/{id}/power"
	VMWareAPI_ChangeVmPowerState = "/vms/{id}/power"

	VMWareAPI_GetVmAllSharedFoldersMounted = "/vms/{id}/sharedfolders"
	VMWareAPI_UpdateVmSharedFolderMounted  = "/vms/{id}/sharedfolders/{folder_id}"
	VMWareAPI_MountVmNewSharedFolder       = "/vms/{id}/sharedfolders"
	VMWareAPI_DeleteSharedFolder           = "/vms/{id}/sharedfolders/{folder_id}"
)

type GetAllVirtualNetworksResponse struct {
	Num    int `json:"num"`
	VmNets []struct {
		Name   string `json:"name"`
		Type   string `json:"bridged"`
		DHCP   string `json:"dhcp"`
		Subnet string `json:"subnet"`
		Mask   string `json:"mask"`
	} `json:"vmnets"`
}

type GetAllMacToIpSettingsForDhcpServiceResponse struct {
	Num      int `json:"num"`
	MacToIps []struct {
		Vmnet string `json:"vmnet"`
		Mac   string `json:"mac"`
		Ip    string `json:"ip"`
	} `json:"mactoips"`
}

type GetAllPortForwardingsResponse struct {
	Num             int `json:"num"`
	PortForwardings []struct {
		Port     int    `json:"port"`
		Protocol string `json:"protocol"`
		Desc     string `json:"desc"`
		Guest    struct {
			Ip   string `json:"ip"`
			Port int    `json:"port"`
		} `json:"guest"`
	} `json:"port_forwardings"`
}

type UpdateMacToIpBindingRequest struct {
	IP string `json:"IP"`
}

type UpdateMacToIpBindingResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UpdatePortForwardingRequest struct {
	GuestIp   string `json:"guestIp"`
	GuestPort int    `json:"guestPort"`
	Desc      string `json:"desc"`
}

type UpdatePortForwardingResponse = VmwareFusionError

type CreateVirtualNetworkRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type CreateVirtualNetworkResponse struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Dhcp   string `json:"dhcp"`
	Subnet string `json:"subnet"`
	Mask   string `json:"mask"`
}

type DeletePortForwardingRequest struct {
}

type DeletePortForwardingResponse struct {
}

type GetVmListRequest struct {
}

type GetVmListResponse = []struct {
	Id   string `json:"id"`
	Path string `json:"path"`
}

type GetVmSettingInformationRequest struct {
}

type GetVmSettingInformationResponse struct {
	Id  string `json:"id"`
	Cpu struct {
		Processors int `json:"processors"`
	} `json:"cpu"`
	Memory int `json:"memory"`
}

type GetVmConfigParamsResponse struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GetVmRestrictionsInformationResponse struct {
	Id                  string `json:"id"`
	ManagedOrg          string `json:"managedOrg"`
	IntegrityConstraint string `json:"integrityconstraint"`
	Cpu                 struct {
		Processors int `json:"processors"`
	} `json:"cpu"`
	Memory        int `json:"memory"`
	ApplianceView struct {
		Author        string `json:"author"`
		Version       string `json:"version"`
		Port          int    `json:"port"`
		ShowAtPowerOn string `json:"showAtPowerOn"`
	} `json:"applianceView"`
	CddvdList struct {
		Num     int `json:"num"`
		Devices []struct {
			Index            int    `json:"index"`
			StartConnected   string `json:"startConnected"`
			ConnectionStatus int    `json:"connectionStatus"`
			DevicePath       string `json:"devicePath"`
		} `json:"devices"`
	} `json:"cddvdList"`
	FloopyList struct {
		Num     int `json:"num"`
		Devices []struct {
			Index            int    `json:"index"`
			StartConnected   string `json:"startConnected"`
			ConnectionStatus int    `json:"connectionStatus"`
			DevicePath       string `json:"devicePath"`
		} `json:"devices"`
	} `json:"floopyList"`
	FirewareType   int `json:"firewareType"`
	GuestIsolation struct {
		CopyDisabled  string `json:"copyDisabled"`
		DndDisabled   string `json:"dndDisabled"`
		HgfsDisabled  string `json:"hgfsDisabled"`
		PasteDisabled string `json:"pasteDisabled"`
	} `json:"guestIsolation"`
	NicList struct {
		Num  int `json:"num"`
		Nics []struct {
			Index      int    `json:"index"`
			Type       string `json:"type"`
			VmNet      string `json:"vmnet"`
			MacAddress string `json:"macAddress"`
		} `json:"nics"`
	} `json:"niclist"`
	ParallelPortList struct {
		Num     int `json:"num"`
		Devices []struct {
			Index            int    `json:"index"`
			StartConnected   string `json:"startConnected"`
			ConnectionStatus int    `json:"connectionStatus"`
			DevicePath       string `json:"devicePath"`
		} `json:"devices"`
	} `json:"parallelPortList"`
	SerialPortList struct {
		Num     int `json:"num"`
		Devices []struct {
			Index            int    `json:"index"`
			StartConnected   string `json:"startConnected"`
			ConnectionStatus int    `json:"connectionStatus"`
			DevicePath       string `json:"devicePath"`
		} `json:"devices"`
	} `json:"serialPortList"`
	UsbList struct {
		Num        int `json:"num"`
		UsbDevices []struct {
			Index       int    `json:"index"`
			Connected   string `json:"connected"`
			BackingInfo string `json:"backingInfo"`
			BackingType int    `json:"BackingType"`
		} `json:"usbDevices"`
	} `json:"usbList"`
	RemoteVNC struct {
		VNCEnabled string `json:"VNCEnabled"`
		VNCPort    int    `json:"VNCPort"`
	} `json:"remoteVNC"`
}

type UpdateVmSettingRequest struct {
	Processors int `json:"processors"`
	Memory     int `json:"memory"`
}

type UpdateVmSettingResponse struct {
	Id  string `json:"id"`
	Cpu struct {
		Processors int `json:"processors"`
	} `json:"cpu"`
	Memory int `json:"memory"`
}

type UpdateVmConfigParamsRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type UpdateVmConfigParamsResponse = VmwareFusionError

type CreateVmCopyRequest struct {
	Name     string `json:"name"`
	ParentId string `json:"parentId"`
}

type CreateVmCopyResponse struct {
	Id  string `json:"id"`
	Cpu struct {
		Processors int `json:"processors"`
	} `json:"cpu"`
	Memory int `json:"memory"`
}

type RegisterVmToVmLibraryRequest struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type RegisterVmToVmLibraryResponse struct {
	Id   string `json:"id"`
	Path string `json:"path"`
}

type DeleteVmRequest struct {
}

type DeleteVmResponse struct {
}

type GetVmIpAddressResponse struct {
	Ip string `json:"ip"`
}

type GetAllVmNetworkAdaptersResponse struct {
	Num  int `json:"num"`
	Nics []struct {
		Index      int    `json:"index"`
		Type       string `json:"type"`
		Vmnet      string `json:"vmnet"`
		MacAddress string `json:"macAddress"`
	} `json:"nics"`
}

type GetIpStackConfigurationOfAllNics struct {
	Nics struct {
		Mac string   `json:"mac"`
		Ip  []string `json:"ip"`
		Dns struct {
			Hostname   string   `json:"hostname"`
			Domainname string   `json:"domainname"`
			Server     []string `json:"server"`
			Search     []string `json:"search"`
		} `json:"dns"`
		Wins struct {
			Primary   string `json:"primary"`
			Secondary string `json:"secondary"`
		} `json:"wins"`
		Dhcp4 struct {
			Enabled bool   `json:"enabled"`
			Setting string `json:"setting"`
		} `json:"dhcp4"`
		Dhcp6 struct {
			Enabled bool   `json:"enabled"`
			Setting string `json:"setting"`
		} `json:"dhcp6"`
	} `json:"nics"`
	Routes []struct {
		Dest      string `json:"dest"`
		Prefix    int    `json:"prefix"`
		Nexthop   string `json:"nexthop"`
		Interface int    `json:"interface"`
		Type      int    `json:"type"`
		Metric    int    `json:"metric"`
	} `json:"routes"`
	Dns struct {
		Hostname   string   `json:"hostname"`
		Domainname string   `json:"domainname"`
		Server     []string `json:"server"`
		Search     []string `json:"search"`
	} `json:"dns"`
	Wins struct {
		Primary   string `json:"primary"`
		Secondary string `json:"secondary"`
	} `json:"wins"`
	DhcpV4 struct {
		Enabled bool   `json:"enabled"`
		Setting string `json:"setting"`
	} `json:"dhcpv4"`
	DhcpV6 struct {
		Enabled bool   `json:"enabled"`
		Setting string `json:"setting"`
	} `json:"dhcpv6"`
}

type UpdateVmNetworkAdapterRequest struct {
	Type  string `json:"type"`
	VmNet string `json:"vmnet"`
}

type UpdateVmNetworkAdapterResponse struct {
	Index      int    `json:"index"`
	Type       string `json:"type"`
	VmNet      string `json:"vmnet"`
	MacAddress string `json:"macAddress"`
}

type CreateVmNetworkAdapterRequest struct {
	Type  string `json:"type"`
	Vmnet string `json:"vmnet"`
}

type CreateVmNetworkAdapterResponse struct {
	Index      int    `json:"index"`
	Type       string `json:"type"`
	Vmnet      string `json:"vmnet"`
	MacAddress string `json:"macAddress"`
}

type DeleteVmNetworkAdapterRequest struct {
}

type DeleteVmNetworkAdapterResponse struct {
}

type GetVmPowerStateResponse struct {
	PowerState string `json:"power_state"`
}

type ChangeVmPowerStateRequest struct {
}

type ChangeVmPowerStateResponse struct {
	PowerState string `json:"power_state"`
}

type GetVmAllSharedFoldersMountedResponse = []struct {
	FolderId string `json:"folder_id"`
	HostPath string `json:"host_path"`
	Flags    int    `json:"flags"`
}

type UpdateVmSharedFolderMountedRequest struct {
	HostPath string `json:"host_path"`
	Flags    int    `json:"flags"`
}

type UpdateVmSharedFolderMountedResponse = []struct {
	FolderId string `json:"folder_id"`
	HostPath string `json:"host_path"`
	Flags    int    `json:"flags"`
}

type MountVmNewSharedFolderRequest struct {
	FolderId string `json:"folder_id"`
	HostPath string `json:"host_path"`
	Flags    int    `json:"flags"`
}

type MountVmNewSharedFolderResponse = []struct {
	FolderId string `json:"folder_id"`
	HostPath string `json:"host_path"`
	Flags    int    `json:"flags"`
}

type DeleteSharedFolderRequest struct {
}

type DeleteSharedFolderResponse struct {
}
