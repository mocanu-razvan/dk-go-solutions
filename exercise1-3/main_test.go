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

import "testing"

func benchmarkEcho1(b *testing.B, size int) {
	args := make([]string, size)
	for i := 0; i < size; i++ {
		args[i] = "-x"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		echo1(args)
	}
}
func benchmarkEcho2(b *testing.B, size int) {
	args := make([]string, size)
	for i := 0; i < size; i++ {
		args[i] = "-x"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		echo2(args)
	}
}
func benchmarkEcho3(b *testing.B, size int) {
	args := make([]string, size)
	for i := 0; i < size; i++ {
		args[i] = "-x"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		echo3(args)
	}
}
func BenchmarkEcho1e1(b *testing.B) { benchmarkEcho1(b, 1e1) }
func BenchmarkEcho1e3(b *testing.B) { benchmarkEcho1(b, 1e3) }
func BenchmarkEcho1e5(b *testing.B) { benchmarkEcho1(b, 1e5) }
func BenchmarkEcho2e1(b *testing.B) { benchmarkEcho2(b, 1e1) }
func BenchmarkEcho2e3(b *testing.B) { benchmarkEcho2(b, 1e3) }
func BenchmarkEcho2e5(b *testing.B) { benchmarkEcho2(b, 1e5) }
func BenchmarkEcho3e1(b *testing.B) { benchmarkEcho3(b, 1e1) }
func BenchmarkEcho3e3(b *testing.B) { benchmarkEcho3(b, 1e3) }
func BenchmarkEcho3e5(b *testing.B) { benchmarkEcho3(b, 1e5) }
