// Go Swiss Ephemeris - Eclipses
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

package swisseph

// #cgo CFLAGS: -I${SRCDIR}/swisseph
// #include <stdlib.h>
// #include <string.h>
// #include "swephexp.h"
import "C"
import (
	"unsafe"
)

// SolEclipseWhenLoc finds the next solar eclipse for a given location
func SolEclipseWhenLoc(tjdStart float64, ifl int32, geopos [3]float64, backward bool) EclipseResult {
	var tret [10]C.double
	var attr [20]C.double
	var serr [asMaxch]C.char

	var geoposC [3]C.double
	for i := 0; i < 3; i++ {
		geoposC[i] = C.double(geopos[i])
	}

	bwd := C.int(0)
	if backward {
		bwd = 1
	}

	flag := C.swe_sol_eclipse_when_loc(
		C.double(tjdStart),
		C.int(ifl),
		&geoposC[0],
		&tret[0],
		&attr[0],
		bwd,
		&serr[0],
	)

	result := EclipseResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Attr:  make([]float64, 20),
	}

	if flag >= 0 {
		result.Maximum = float64(tret[0])
		result.Begin = float64(tret[1])
		result.End = float64(tret[2])
		result.Totality = float64(tret[3])

		for i := 0; i < 20; i++ {
			result.Attr[i] = float64(attr[i])
		}
	}

	return result
}

// SolEclipseWhenGlob finds the next solar eclipse globally
func SolEclipseWhenGlob(tjdStart float64, ifl int32, ifltype int32, backward bool) EclipseResult {
	var tret [10]C.double
	var serr [asMaxch]C.char

	bwd := C.int(0)
	if backward {
		bwd = 1
	}

	flag := C.swe_sol_eclipse_when_glob(
		C.double(tjdStart),
		C.int(ifl),
		C.int(ifltype),
		&tret[0],
		bwd,
		&serr[0],
	)

	result := EclipseResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
	}

	if flag >= 0 {
		result.Maximum = float64(tret[0])
		result.Begin = float64(tret[2])
		result.End = float64(tret[3])
	}

	return result
}

// SolEclipseHow calculates the attributes of a solar eclipse at a given location
func SolEclipseHow(tjdUt float64, ifl int32, geopos [3]float64) EclipseResult {
	var attr [20]C.double
	var serr [asMaxch]C.char

	var geoposC [3]C.double
	for i := 0; i < 3; i++ {
		geoposC[i] = C.double(geopos[i])
	}

	flag := C.swe_sol_eclipse_how(
		C.double(tjdUt),
		C.int(ifl),
		&geoposC[0],
		&attr[0],
		&serr[0],
	)

	result := EclipseResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Attr:  make([]float64, 20),
	}

	for i := 0; i < 20; i++ {
		result.Attr[i] = float64(attr[i])
	}

	return result
}

// SolEclipseWhere calculates where a solar eclipse is central or maximal
func SolEclipseWhere(tjdUt float64, ifl int32) EclipseWhereResult {
	var geopos [2]C.double
	var attr [20]C.double
	var serr [asMaxch]C.char

	flag := C.swe_sol_eclipse_where(
		C.double(tjdUt),
		C.int(ifl),
		&geopos[0],
		&attr[0],
		&serr[0],
	)

	result := EclipseWhereResult{
		Flag:      int32(flag),
		Longitude: float64(geopos[0]),
		Latitude:  float64(geopos[1]),
		Error:     C.GoString(&serr[0]),
		Attr:      make([]float64, 20),
	}

	for i := 0; i < 20; i++ {
		result.Attr[i] = float64(attr[i])
	}

	return result
}

// LunEclipseWhen finds the next lunar eclipse
func LunEclipseWhen(tjdStart float64, ifl int32, ifltype int32, backward bool) EclipseResult {
	var tret [10]C.double
	var serr [asMaxch]C.char

	bwd := C.int(0)
	if backward {
		bwd = 1
	}

	flag := C.swe_lun_eclipse_when(
		C.double(tjdStart),
		C.int(ifl),
		C.int(ifltype),
		&tret[0],
		bwd,
		&serr[0],
	)

	result := EclipseResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
	}

	if flag >= 0 {
		result.Maximum = float64(tret[0])
		result.Begin = float64(tret[2])
		result.End = float64(tret[3])
		result.Totality = float64(tret[4]) - float64(tret[5])
	}

	return result
}

// LunEclipseWhenLoc finds the next lunar eclipse for a given location
func LunEclipseWhenLoc(tjdStart float64, ifl int32, geopos [3]float64, backward bool) EclipseResult {
	var tret [10]C.double
	var attr [20]C.double
	var serr [asMaxch]C.char

	var geoposC [3]C.double
	for i := 0; i < 3; i++ {
		geoposC[i] = C.double(geopos[i])
	}

	bwd := C.int(0)
	if backward {
		bwd = 1
	}

	flag := C.swe_lun_eclipse_when_loc(
		C.double(tjdStart),
		C.int(ifl),
		&geoposC[0],
		&tret[0],
		&attr[0],
		bwd,
		&serr[0],
	)

	result := EclipseResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Attr:  make([]float64, 20),
	}

	if flag >= 0 {
		result.Maximum = float64(tret[0])
		result.Begin = float64(tret[2])
		result.End = float64(tret[3])

		for i := 0; i < 20; i++ {
			result.Attr[i] = float64(attr[i])
		}
	}

	return result
}

// LunEclipseHow calculates the attributes of a lunar eclipse
func LunEclipseHow(tjdUt float64, ifl int32, geopos [3]float64) EclipseResult {
	var attr [20]C.double
	var serr [asMaxch]C.char

	var geoposC [3]C.double
	for i := 0; i < 3; i++ {
		geoposC[i] = C.double(geopos[i])
	}

	flag := C.swe_lun_eclipse_how(
		C.double(tjdUt),
		C.int(ifl),
		&geoposC[0],
		&attr[0],
		&serr[0],
	)

	result := EclipseResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Attr:  make([]float64, 20),
	}

	for i := 0; i < 20; i++ {
		result.Attr[i] = float64(attr[i])
	}

	return result
}

// LunOccultWhenLoc finds the next lunar occultation for a given location
func LunOccultWhenLoc(tjdStart float64, ipl int32, starname string, ifl int32, geopos [3]float64, backward bool) EclipseResult {
	var tret [10]C.double
	var attr [20]C.double
	var serr [asMaxch]C.char

	var geoposC [3]C.double
	for i := 0; i < 3; i++ {
		geoposC[i] = C.double(geopos[i])
	}

	cStar := C.CString(starname)
	defer C.free(unsafe.Pointer(cStar))

	bwd := C.int(0)
	if backward {
		bwd = 1
	}

	flag := C.swe_lun_occult_when_loc(
		C.double(tjdStart),
		C.int(ipl),
		cStar,
		C.int(ifl),
		&geoposC[0],
		&tret[0],
		&attr[0],
		bwd,
		&serr[0],
	)

	result := EclipseResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Attr:  make([]float64, 20),
	}

	if flag >= 0 {
		result.Maximum = float64(tret[0])
		result.Begin = float64(tret[1])
		result.End = float64(tret[2])

		for i := 0; i < 20; i++ {
			result.Attr[i] = float64(attr[i])
		}
	}

	return result
}

// LunOccultWhenGlob finds the next lunar occultation globally
func LunOccultWhenGlob(tjdStart float64, ipl int32, starname string, ifl int32, ifltype int32, backward bool) EclipseResult {
	var tret [10]C.double
	var serr [asMaxch]C.char

	cStar := C.CString(starname)
	defer C.free(unsafe.Pointer(cStar))

	bwd := C.int(0)
	if backward {
		bwd = 1
	}

	flag := C.swe_lun_occult_when_glob(
		C.double(tjdStart),
		C.int(ipl),
		cStar,
		C.int(ifl),
		C.int(ifltype),
		&tret[0],
		bwd,
		&serr[0],
	)

	result := EclipseResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
	}

	if flag >= 0 {
		result.Maximum = float64(tret[0])
		result.Begin = float64(tret[2])
		result.End = float64(tret[3])
	}

	return result
}

// LunOccultWhere calculates where a lunar occultation is central or maximal
func LunOccultWhere(tjdUt float64, ipl int32, starname string, ifl int32) EclipseWhereResult {
	var geopos [2]C.double
	var attr [20]C.double
	var serr [asMaxch]C.char

	cStar := C.CString(starname)
	defer C.free(unsafe.Pointer(cStar))

	flag := C.swe_lun_occult_where(
		C.double(tjdUt),
		C.int(ipl),
		cStar,
		C.int(ifl),
		&geopos[0],
		&attr[0],
		&serr[0],
	)

	result := EclipseWhereResult{
		Flag:      int32(flag),
		Longitude: float64(geopos[0]),
		Latitude:  float64(geopos[1]),
		Error:     C.GoString(&serr[0]),
		Attr:      make([]float64, 20),
	}

	for i := 0; i < 20; i++ {
		result.Attr[i] = float64(attr[i])
	}

	return result
}
