// Exercise 1.3
// Copyright (c) 2019 Răzvan Mocanu

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join. (Section 1.6
// illustrates part of the time package, and section 11.4 shows how to write
// benchmark tests for systematic performance evaluation.)
package main

import (
	"strings"
)

func echo1(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
}

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
}

func echo3(args []string) {
	strings.Join(args[1:], " ")
}
