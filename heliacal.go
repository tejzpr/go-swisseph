// Go Swiss Ephemeris - Heliacal Events
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
	"fmt"
	"unsafe"
)

// HeliacalUT calculates heliacal risings and settings (universal time)
func HeliacalUT(tjdstart float64, geopos [3]float64, datm [4]float64, dobs [6]float64, objectname string, eventType int32, helflag int32) HeliacalResult {
	var dret [50]C.double
	var serr [asMaxch]C.char

	var geoposC [3]C.double
	for i := 0; i < 3; i++ {
		geoposC[i] = C.double(geopos[i])
	}

	var datmC [4]C.double
	for i := 0; i < 4; i++ {
		datmC[i] = C.double(datm[i])
	}

	var dobsC [6]C.double
	for i := 0; i < 6; i++ {
		dobsC[i] = C.double(dobs[i])
	}

	cObj := C.CString(objectname)
	defer C.free(unsafe.Pointer(cObj))

	flag := C.swe_heliacal_ut(
		C.double(tjdstart),
		&geoposC[0],
		&datmC[0],
		&dobsC[0],
		cObj,
		C.int(eventType),
		C.int(helflag),
		&dret[0],
		&serr[0],
	)

	result := HeliacalResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Time:  make([]float64, 3),
		Attr:  make([]float64, 47),
	}

	if flag >= 0 {
		for i := 0; i < 3; i++ {
			result.Time[i] = float64(dret[i])
		}
		for i := 3; i < 50; i++ {
			result.Attr[i-3] = float64(dret[i])
		}
	}

	return result
}

// HeliacalPhenoUT calculates heliacal phenomena (universal time)
func HeliacalPhenoUT(tjdUt float64, geopos [3]float64, datm [4]float64, dobs [6]float64, objectname string, eventType int32, helflag int32) HeliacalResult {
	var darr [50]C.double
	var serr [asMaxch]C.char

	var geoposC [3]C.double
	for i := 0; i < 3; i++ {
		geoposC[i] = C.double(geopos[i])
	}

	var datmC [4]C.double
	for i := 0; i < 4; i++ {
		datmC[i] = C.double(datm[i])
	}

	var dobsC [6]C.double
	for i := 0; i < 6; i++ {
		dobsC[i] = C.double(dobs[i])
	}

	cObj := C.CString(objectname)
	defer C.free(unsafe.Pointer(cObj))

	flag := C.swe_heliacal_pheno_ut(
		C.double(tjdUt),
		&geoposC[0],
		&datmC[0],
		&dobsC[0],
		cObj,
		C.int(eventType),
		C.int(helflag),
		&darr[0],
		&serr[0],
	)

	result := HeliacalResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Attr:  make([]float64, 50),
	}

	for i := 0; i < 50; i++ {
		result.Attr[i] = float64(darr[i])
	}

	return result
}

// VisLimitMag calculates the limiting visual magnitude
func VisLimitMag(tjdUt float64, geopos [3]float64, datm [4]float64, dobs [6]float64, objectname string, helflag int32) HeliacalResult {
	var dret [50]C.double
	var serr [asMaxch]C.char

	var geoposC [3]C.double
	for i := 0; i < 3; i++ {
		geoposC[i] = C.double(geopos[i])
	}

	var datmC [4]C.double
	for i := 0; i < 4; i++ {
		datmC[i] = C.double(datm[i])
	}

	var dobsC [6]C.double
	for i := 0; i < 6; i++ {
		dobsC[i] = C.double(dobs[i])
	}

	cObj := C.CString(objectname)
	defer C.free(unsafe.Pointer(cObj))

	flag := C.swe_vis_limit_mag(
		C.double(tjdUt),
		&geoposC[0],
		&datmC[0],
		&dobsC[0],
		cObj,
		C.int(helflag),
		&dret[0],
		&serr[0],
	)

	result := HeliacalResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Attr:  make([]float64, 50),
	}

	for i := 0; i < 50; i++ {
		result.Attr[i] = float64(dret[i])
	}

	return result
}

// Solcross calculates when the Sun crosses a specific longitude (ephemeris time)
func Solcross(x2cross float64, tjdEt float64, flag int32) (float64, error) {
	var serr [asMaxch]C.char

	jdcross := C.swe_solcross(
		C.double(x2cross),
		C.double(tjdEt),
		C.int(flag),
		&serr[0],
	)

	errStr := C.GoString(&serr[0])
	if errStr != "" {
		return float64(jdcross), fmt.Errorf("%s", errStr)
	}

	return float64(jdcross), nil
}

// SolcrossUT calculates when the Sun crosses a specific longitude (universal time)
func SolcrossUT(x2cross float64, tjdUt float64, flag int32) (float64, error) {
	var serr [asMaxch]C.char

	jdcross := C.swe_solcross_ut(
		C.double(x2cross),
		C.double(tjdUt),
		C.int(flag),
		&serr[0],
	)

	errStr := C.GoString(&serr[0])
	if errStr != "" {
		return float64(jdcross), fmt.Errorf("%s", errStr)
	}

	return float64(jdcross), nil
}

// Mooncross calculates when the Moon crosses a specific longitude (ephemeris time)
func Mooncross(x2cross float64, tjdEt float64, flag int32) (float64, error) {
	var serr [asMaxch]C.char

	jdcross := C.swe_mooncross(
		C.double(x2cross),
		C.double(tjdEt),
		C.int(flag),
		&serr[0],
	)

	errStr := C.GoString(&serr[0])
	if errStr != "" {
		return float64(jdcross), fmt.Errorf("%s", errStr)
	}

	return float64(jdcross), nil
}

// MooncrossUT calculates when the Moon crosses a specific longitude (universal time)
func MooncrossUT(x2cross float64, tjdUt float64, flag int32) (float64, error) {
	var serr [asMaxch]C.char

	jdcross := C.swe_mooncross_ut(
		C.double(x2cross),
		C.double(tjdUt),
		C.int(flag),
		&serr[0],
	)

	errStr := C.GoString(&serr[0])
	if errStr != "" {
		return float64(jdcross), fmt.Errorf("%s", errStr)
	}

	return float64(jdcross), nil
}

// MooncrossNode calculates when the Moon crosses its node (ephemeris time)
func MooncrossNode(tjdEt float64, flag int32) (float64, float64, error) {
	var xlon, xlat C.double
	var serr [asMaxch]C.char

	jdcross := C.swe_mooncross_node(
		C.double(tjdEt),
		C.int(flag),
		&xlon,
		&xlat,
		&serr[0],
	)

	errStr := C.GoString(&serr[0])
	if errStr != "" {
		return float64(jdcross), float64(xlon), fmt.Errorf("%s", errStr)
	}

	return float64(jdcross), float64(xlon), nil
}

// MooncrossNodeUT calculates when the Moon crosses its node (universal time)
func MooncrossNodeUT(tjdUt float64, flag int32) (float64, float64, error) {
	var xlon, xlat C.double
	var serr [asMaxch]C.char

	jdcross := C.swe_mooncross_node_ut(
		C.double(tjdUt),
		C.int(flag),
		&xlon,
		&xlat,
		&serr[0],
	)

	errStr := C.GoString(&serr[0])
	if errStr != "" {
		return float64(jdcross), float64(xlon), fmt.Errorf("%s", errStr)
	}

	return float64(jdcross), float64(xlon), nil
}

// HelioCross calculates when a planet crosses a specific heliocentric longitude (ephemeris time)
func HelioCross(ipl int32, x2cross float64, tjdEt float64, iflag int32, dir int32) (float64, error) {
	var jdcross C.double
	var serr [asMaxch]C.char

	result := C.swe_helio_cross(
		C.int(ipl),
		C.double(x2cross),
		C.double(tjdEt),
		C.int(iflag),
		C.int(dir),
		&jdcross,
		&serr[0],
	)

	if result == ERR {
		return 0, fmt.Errorf("%s", C.GoString(&serr[0]))
	}

	return float64(jdcross), nil
}

// HelioCrossUT calculates when a planet crosses a specific heliocentric longitude (universal time)
func HelioCrossUT(ipl int32, x2cross float64, tjdUt float64, iflag int32, dir int32) (float64, error) {
	var jdcross C.double
	var serr [asMaxch]C.char

	result := C.swe_helio_cross_ut(
		C.int(ipl),
		C.double(x2cross),
		C.double(tjdUt),
		C.int(iflag),
		C.int(dir),
		&jdcross,
		&serr[0],
	)

	if result == ERR {
		return 0, fmt.Errorf("%s", C.GoString(&serr[0]))
	}

	return float64(jdcross), nil
}
