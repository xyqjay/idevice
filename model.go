package lib

import "fmt"

type BootState string
type DeviceType string

const (
	BootStateBooted   BootState = "Booted"
	BootStateShutdown BootState = "Shutdown"
	BootStateAll      BootState = "All"
)

const (
	Simulator  DeviceType = "Simulator"
	TrueDevice DeviceType = "TrueDevice"

	IPhone            DeviceType = "IPhone"
	IPad              DeviceType = "IPad"
	MacBookPro        DeviceType = "MacBookPro"
	MacBookAir        DeviceType = "MacBookAir"
	IMac              DeviceType = "IMac"
	MacStudio         DeviceType = "MacStudio"
	MacMini           DeviceType = "MacMini"
	DeviceTypeUnknown DeviceType = "DeviceTypeUnknown"
)

type Runtime struct {
	Identifier string // eg.: com.apple.CoreSimulator.SimRuntime.iOS-13-1
	Name       string // eg.: iOS 13.1
	Platform   string // eg.: iOS
}

type Device struct {
	Udid        string //
	State       string // Shutdown / Booted
	Name        string //
	IsAvailable bool
	Runtime     Runtime
	Type        DeviceType
}

func (f Device) availableStatus() string {
	if f.IsAvailable {
		return "Available"
	}
	return "UnAvailable"
}

func (f Device) Description() string {
	str := fmt.Sprintf("%s %s (%s) %s %s %s", f.Udid, f.Name, f.Type, f.Runtime.Name, f.State, f.availableStatus())
	return str
}
