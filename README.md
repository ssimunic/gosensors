# gosensors
[![GoDoc](https://godoc.org/github.com/ssimunic/gosensors?status.png)](http://godoc.org/github.com/ssimunic/gosensors)

Hardware sensors using Go and lm-sensors

## Setup
* Install `lm-sensors` with `sudo apt get install lm-sensors`
* `go get github.com/ssimunic/gosensors` in your working directory

## Example

### Program

```go
package main

import (
	"fmt"
	"github.com/ssimunic/gosensors"
)

func main() {
	sensors, err := gosensors.NewFromSystem()
	// sensors, err := gosensors.NewFromFile("/path/to/log.txt")

	if err != nil {
		panic(err)
	}

	// Sensors implements Stringer interface,
	// so code below will print out JSON
	fmt.Println(sensors)

	// Also valid
	// fmt.Println("JSON:", sensors.JSON())

	// Iterate over chips
	for chip := range sensors.Chips {
		// Iterate over entries
		for key, value := range sensors.Chips[chip] {
			// If CPU or GPU, print out
			if key == "CPU" || key == "GPU" {
				fmt.Println(key, value)
			}
		}
	}
}
```

### Output:
```
{"chips":{"acpitz-virtual-0":{"Adapter":"Virtual device","temp1":"+25.0°C  (crit = +107.0°C)"},"coretemp-isa-0000":{"Adapter":"ISA adapter","Core 0":"+76.0°C  (high = +105.0°C, crit = +105.0°C)","Core 1":"+74.0°C  (high = +105.0°C, crit = +105.0°C)","Physical id 0":"+76.0°C  (high = +105.0°C, crit = +105.0°C)"},"dell_smm-virtual-0":{"Adapter":"Virtual device","CPU":"+65.0°C","GPU":"+50.0°C","Other":"+50.0°C","Processor Fan":"2200 RPM","SODIMM":"+39.0°C","fan2":"2200 RPM"}}}
CPU +65.0°C
GPU +50.0°C
```
