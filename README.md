# github.com/xyqjay/idevice
Used to get available iPhone (simulator) / Mac devices using golang.

## ！Depends on command line `xcrun`

The principle is analysis command line `xcrun simctl list --json`'s outputs.

# Example
Code:
```
	res := GetBotedSimDevices()
	for i := 0; i < len(res); i++ {
		d := res[i]
		fmt.Println(d.Description())
	}
```
OutPut:
```
4914C636-171A-41A2-9EB4-DC4CBC76C7B9 iPhone 15 (Simulator) iOS 17.2 Booted Available
```