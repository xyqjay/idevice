/**
 * Author: Jay
 * File: idevice.go
 */

package lib

import (
	"encoding/json"
	"log"
	"os/exec"
	"strings"
)

func GetAllDevices() []Device {
	out, err := exec.Command("xcrun", "xctrace", "list", "devices").Output()
	if err != nil {
		log.Fatal(err)
	}
	x := string(out)
	start := false
	var tmpDevices = []Device{}

	for _, line := range strings.Split(strings.TrimSuffix(x, "\n"), "\n") {
		if strings.Contains(line, "Devices") {
			start = true
			continue
		}
		if strings.Contains(line, "Simulators") {
			start = false
			continue
		}
		if start && len(line) > 2 {
			var str = strings.Trim(line, " ")
			index := strings.LastIndex(str, " ")
			name := str[0:index]
			id := str[index : len(str)-1]
			id = strings.Replace(strings.Trim(id, " "), "(", "", 1)
			device := Device{id, string(BootStateBooted), strings.Trim(name, " "), true, Runtime{}, TrueDevice}

			tmpDevices = append(tmpDevices, device)
		}
	}

	return tmpDevices
}

func GetAllSimDevices() []Device {
	return getSimDevicesByState(BootStateAll)
}
func GetShutdownSimDevices() []Device {
	return getSimDevicesByState(BootStateShutdown)
}
func GetBotedSimDevices() []Device {
	return getSimDevicesByState(BootStateBooted)
}

func getSimDevicesByState(state BootState) []Device {
	out, err := exec.Command("xcrun", "simctl", "list", "--json").Output()
	if err != nil {
		log.Fatal(err)
	}
	var tmpMap map[string]interface{}
	err1 := json.Unmarshal(out, &tmpMap)
	if err1 != nil {
		panic(err1)
	}

	runtimes := tmpMap["runtimes"]
	runtimeResults := getSimRuntimes(runtimes)

	devices := tmpMap["devices"]

	devicesModel := getSimDevices(devices.(map[string]interface{}), runtimeResults, state)

	return devicesModel
}

func getSimRuntimes(runtimes interface{}) map[string]Runtime {
	// fmt.Printf("%T\n", runtimes)
	var tmp = map[string]Runtime{}
	switch v := runtimes.(type) {
	case []interface{}:
		for _, s := range v {
			switch ss := s.(type) {
			case map[string]interface{}:
				value := Runtime{ss["identifier"].(string), ss["name"].(string), ss["platform"].(string)}
				tmp[value.Identifier] = value
			default:
				// fmt.Printf("invalid type: %T\n", v)
				break
			}
		}
	default:
		// fmt.Printf("invalid type: %T\n", v)
		break
	}
	return tmp
}

func getSimDevices(devices map[string]interface{}, runtimes map[string]Runtime, state BootState) []Device {
	var tmpDevices = []Device{}
	for k, v := range devices {
		for _, d := range v.([]interface{}) {
			a := d.(map[string]interface{})
			udid := a["udid"].(string)
			_state := a["state"].(string)
			name := a["name"].(string)
			isAvailable := a["isAvailable"].(bool)
			if state == BootStateAll {
				runtime := runtimes[k]
				device := Device{udid, _state, name, isAvailable, runtime, Simulator}
				tmpDevices = append(tmpDevices, device)
			} else if _state == string(state) {
				runtime := runtimes[k]
				device := Device{udid, _state, name, isAvailable, runtime, Simulator}
				tmpDevices = append(tmpDevices, device)
			}
		}
	}
	return tmpDevices
}
