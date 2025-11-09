// Go Swiss Ephemeris - Fixed Stars
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

// Fixstar calculates fixed star positions (ephemeris time)
func Fixstar(star string, tjdEt float64, iflag int32) FixstarResult {
	var xx [6]C.double
	var serr [asMaxch]C.char

	// Allocate buffer for star name (input/output)
	starBuf := make([]byte, MaxStname)
	copy(starBuf, []byte(star))
	cStar := (*C.char)(unsafe.Pointer(&starBuf[0]))

	flag := C.swe_fixstar(
		cStar,
		C.double(tjdEt),
		C.int(iflag),
		&xx[0],
		&serr[0],
	)

	result := FixstarResult{
		Flag:     int32(flag),
		StarName: C.GoString(cStar),
		Error:    C.GoString(&serr[0]),
		Data:     make([]float64, 6),
	}

	for i := 0; i < 6; i++ {
		result.Data[i] = float64(xx[i])
	}

	return result
}

// FixstarUT calculates fixed star positions (universal time)
func FixstarUT(star string, tjdUt float64, iflag int32) FixstarResult {
	var xx [6]C.double
	var serr [asMaxch]C.char

	starBuf := make([]byte, MaxStname)
	copy(starBuf, []byte(star))
	cStar := (*C.char)(unsafe.Pointer(&starBuf[0]))

	flag := C.swe_fixstar_ut(
		cStar,
		C.double(tjdUt),
		C.int(iflag),
		&xx[0],
		&serr[0],
	)

	result := FixstarResult{
		Flag:     int32(flag),
		StarName: C.GoString(cStar),
		Error:    C.GoString(&serr[0]),
		Data:     make([]float64, 6),
	}

	for i := 0; i < 6; i++ {
		result.Data[i] = float64(xx[i])
	}

	return result
}

// FixstarMag calculates fixed star magnitude
func FixstarMag(star string) FixstarMagResult {
	var mag C.double
	var serr [asMaxch]C.char

	starBuf := make([]byte, MaxStname)
	copy(starBuf, []byte(star))
	cStar := (*C.char)(unsafe.Pointer(&starBuf[0]))

	flag := C.swe_fixstar_mag(
		cStar,
		&mag,
		&serr[0],
	)

	return FixstarMagResult{
		Flag:      int32(flag),
		StarName:  C.GoString(cStar),
		Magnitude: float64(mag),
		Error:     C.GoString(&serr[0]),
	}
}

// Fixstar2 calculates fixed star positions using new star file format (ephemeris time)
func Fixstar2(star string, tjdEt float64, iflag int32) FixstarResult {
	var xx [6]C.double
	var serr [asMaxch]C.char

	starBuf := make([]byte, MaxStname)
	copy(starBuf, []byte(star))
	cStar := (*C.char)(unsafe.Pointer(&starBuf[0]))

	flag := C.swe_fixstar2(
		cStar,
		C.double(tjdEt),
		C.int(iflag),
		&xx[0],
		&serr[0],
	)

	result := FixstarResult{
		Flag:     int32(flag),
		StarName: C.GoString(cStar),
		Error:    C.GoString(&serr[0]),
		Data:     make([]float64, 6),
	}

	for i := 0; i < 6; i++ {
		result.Data[i] = float64(xx[i])
	}

	return result
}

// Fixstar2UT calculates fixed star positions using new star file format (universal time)
func Fixstar2UT(star string, tjdUt float64, iflag int32) FixstarResult {
	var xx [6]C.double
	var serr [asMaxch]C.char

	starBuf := make([]byte, MaxStname)
	copy(starBuf, []byte(star))
	cStar := (*C.char)(unsafe.Pointer(&starBuf[0]))

	flag := C.swe_fixstar2_ut(
		cStar,
		C.double(tjdUt),
		C.int(iflag),
		&xx[0],
		&serr[0],
	)

	result := FixstarResult{
		Flag:     int32(flag),
		StarName: C.GoString(cStar),
		Error:    C.GoString(&serr[0]),
		Data:     make([]float64, 6),
	}

	for i := 0; i < 6; i++ {
		result.Data[i] = float64(xx[i])
	}

	return result
}

// Fixstar2Mag calculates fixed star magnitude using new star file format
func Fixstar2Mag(star string) FixstarMagResult {
	var mag C.double
	var serr [asMaxch]C.char

	starBuf := make([]byte, MaxStname)
	copy(starBuf, []byte(star))
	cStar := (*C.char)(unsafe.Pointer(&starBuf[0]))

	flag := C.swe_fixstar2_mag(
		cStar,
		&mag,
		&serr[0],
	)

	return FixstarMagResult{
		Flag:      int32(flag),
		StarName:  C.GoString(cStar),
		Magnitude: float64(mag),
		Error:     C.GoString(&serr[0]),
	}
}
