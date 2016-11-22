package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"encoding/json"
	"strings"
	"strconv"
)

// Looks like:
// {"timestamp":"1479847503.794862747","source":"atc","message":"atc.baggage-collector.could-not-locate-worker","log_level":1,"data":{"session":"16","worker-id":"fe6073d0-5d6e-4076-9387-f8173b015191"}}
type Line struct {
	Timestamp EpochTime    `json:"timestamp"`
	Source    string       `json:"source"`
	Message   string       `json:"message"`
	Log_level int          `json:"log_level"`
	Data      interface{}  `json:"data"`
}

type EpochTime struct {
	Time string
}

var nullTime = "null"
var timeFormat = "2006/01/02 15:04:05.999999 -0700"

func (ct *EpochTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = s
		return
	}
	parts := strings.Split(s, ".")
	epochSec, err := strconv.ParseInt(parts[0], 10, 32)
	epochNsec, err := strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		log.Fatalf("Unable to convert time %s - %v", string(b), err)
	}
	ct.Time = time.Unix(epochSec, epochNsec).Format(timeFormat)
	return
}

func (ct *EpochTime) MarshalJSON() ([]byte, error) {
	if ct.Time == nullTime {
		return []byte(nullTime), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time)), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := &Line{}

		err := json.Unmarshal(scanner.Bytes(), line)
		if err != nil {
			log.Fatalf("Unable to process line %s - %v", scanner.Text(), err)
		}

		raw, err := json.Marshal(line)
		if err != nil {
			log.Fatalf("Unable to unmarshal line %s - %v", scanner.Text(), err)
		}

		fmt.Println(string(raw))
		//fmt.Printf("%v\n", line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
