package gosensors

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os/exec"
	"strings"
)

// Sensors struct represents lm-sensors output.
// Content field contains string output.
// Chips field contains map[string]Entries.
// Example (JSON style):
// "coretemp-isa-0000": {
//	"CPU": "+60.0°C",
//	"GPU": "+48.0°C",
// }
type Sensors struct {
	Content string             `json:"-"`
	Chips   map[string]Entries `json:"chips"`
}

// Entries representing key, value pairs for chips.
// Example (JSON style):
// "GPU": "+56.0°C"
// "CPU": "+68.0°C"
type Entries map[string]string

func construct(content string) *Sensors {
	s := &Sensors{}
	s.Content = content
	s.Chips = map[string]Entries{}

	lines := strings.Split(s.Content, "\n")

	var chip string
	for _, line := range lines {
		if len(line) > 0 {
			if !strings.Contains(line, ":") {
				chip = line
				s.Chips[chip] = Entries{}
			} else if len(chip) > 0 {
				parts := strings.Split(line, ":")
				entry := parts[0]
				value := strings.TrimRight(strings.TrimLeft(parts[1], " "), " ")
				s.Chips[chip][entry] = value
			}
		}
	}

	return s
}

// NewFromSystem executes "sensors" system command and returns constructed Sensors struct.
// A successful call returns err == nil.
func NewFromSystem() (*Sensors, error) {
	out, err := exec.Command("sensors").Output()
	if err != nil {
		return &Sensors{}, errors.New("lm-sensors missing")
	}

	s := construct(string(out))

	return s, nil
}

// NewFromFile reads content from log file and returns constructed Sensors struct.
// A successful call returns err == nil.
func NewFromFile(path string) (*Sensors, error) {
	out, err := ioutil.ReadFile(path)
	if err != nil {
		return &Sensors{}, err
	}

	s := construct(string(out))
	return s, nil
}

// JSON returns JSON of Sensors.
func (s *Sensors) JSON() string {
	out, _ := json.Marshal(s)

	return string(out)
}

func (s *Sensors) String() string {
	return s.JSON()
}
