/**
 * Author: Jay
 * File: idevice_test.go
 */

package lib

import (
	"fmt"
	"testing"
)

func TestGetAllDevices(t *testing.T) {
	res := GetAllDevices()
	for i := 0; i < len(res); i++ {
		d := res[i]
		fmt.Println(d.Description())
	}
}
func TestGetSimulatersBotedDevices(t *testing.T) {
	res := GetBotedSimDevices()
	for i := 0; i < len(res); i++ {
		d := res[i]
		fmt.Println(d.Description())
	}
}

func TestGetSimulatersAllDevices(t *testing.T) {
	res := GetAllSimDevices()
	for i := 0; i < len(res); i++ {
		d := res[i]
		fmt.Println(d.Description())
	}
}

func TestGetShutdownDevices(t *testing.T) {
	res := GetShutdownSimDevices()
	for i := 0; i < len(res); i++ {
		d := res[i]
		fmt.Println(d.Description())
	}
}
