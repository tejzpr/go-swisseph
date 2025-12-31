// Go Swiss Ephemeris - CGO Package
//
// Copyright (C) 2025-2026 Tejus Pratap <tejzpr@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package swisseph is a CGO package that compiles the Swiss Ephemeris C source files.
// This package exists solely to make CGO compile the C files in this directory.
// The C files are copied from the Swiss Ephemeris library and included directly in this repository.
package swisseph

/*
#cgo darwin CFLAGS: -I${SRCDIR} -std=c11 -Wno-unused-result -Wno-format
#cgo !darwin CFLAGS: -I${SRCDIR} -Wno-unused-result -Wno-format
#cgo LDFLAGS: -lm
*/
import "C"
