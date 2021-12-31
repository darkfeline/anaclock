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

package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConvert(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p    timepoint
		want string
	}{
		{timepoint{12, 00}, " :  .  12 .  : "},
		{timepoint{12, 01}, "12| .  :  .  13"},
		{timepoint{00, 02}, "00 |.  :  .  01"},
		{timepoint{23, 10}, "23  .  :  .| 00"},
		{timepoint{23, 11}, "23  .  :  . |00"},
		{timepoint{23, 06}, "23  .  |  .  00"},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%#v", c.p), func(t *testing.T) {
			t.Parallel()
			got := convert(c.p)
			if diff := diff(c.want, string(got)); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestTimepoint_Add(t *testing.T) {
	t.Parallel()
	cases := []struct {
		p    timepoint
		i    int
		want timepoint
	}{
		{timepoint{1, 1}, 1, timepoint{1, 2}},
		{timepoint{1, 1}, -1, timepoint{1, 0}},
		{timepoint{5, 6}, 6, timepoint{6, 0}},
		{timepoint{5, 6}, -6, timepoint{5, 0}},
		{timepoint{23, intervals - 1}, 1, timepoint{0, 0}},
		{timepoint{0, 0}, -1, timepoint{23, intervals - 1}},
	}
	for _, c := range cases {
		c := c
		t.Run(fmt.Sprintf("%#v + %d", c.p, c.i), func(t *testing.T) {
			t.Parallel()
			got := c.p.add(c.i)
			if diff := diff(c.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func diff(a, b interface{}) string {
	return cmp.Diff(a, b, cmp.AllowUnexported(timepoint{}))
}
