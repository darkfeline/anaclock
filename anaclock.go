// Copyright (C) 2021 Allen Li
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Command anaclock prints a simple analog clock as a line of text.
package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	now := time.Now()
	os.Stdout.Write(convert(newTimepoint(now)))
}

const intervals = 12
const timeline = " :  .  :  .  : "

func convert(now timepoint) []byte {
	var out [intervals + 3]byte
	copy(out[:], []byte(timeline))

	if now.interval == 0 {
		setHour(out[7:], now.hour)
		return out[:]
	}
	setHour(out[:], now.hour)
	setHour(out[len(out)-2:], (now.hour+1)%24)
	// Mark now
	out[now.interval+1] = '|'
	return out[:]
}

func setHour(out []byte, hour int) {
	copy(out[:2], []byte(fmt.Sprintf("%02d", hour)))
}

type timepoint struct {
	hour     int
	interval int
}

func newTimepoint(t time.Time) timepoint {
	progress := float64(t.Minute()) / 60
	return timepoint{
		hour:     t.Hour(),
		interval: int(math.Round(progress * intervals)),
	}
}

// Add intervals
func (p timepoint) add(i int) timepoint {
	total := p.hour*intervals + p.interval + i
	total += intervals * 24 // handle negative total
	p.hour = total / intervals % 24
	p.interval = total % intervals
	return p
}
