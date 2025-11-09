// Go Swiss Ephemeris
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

import (
	"fmt"
	"unsafe"

	_ "github.com/tejzpr/go-swisseph/swisseph"
)

/*
#cgo CFLAGS: -I${SRCDIR}/swisseph -Wno-unused-result -Wno-format
#include <stdlib.h>
#include <string.h>
#include "swephexp.h"
*/
import "C"

const (
	// AS_MAXCH is the maximum length for error messages
	asMaxch = 256
)

// PackageVersion is the version of the Go Swiss Ephemeris bindings
const PackageVersion = "v1.0.0"

// Version returns the Swiss Ephemeris version string
func Version() string {
	cVersion := C.swe_version((*C.char)(C.malloc(C.size_t(asMaxch))))
	defer C.free(unsafe.Pointer(cVersion))
	return C.GoString(cVersion)
}

// GetPackageVersion returns the version of the Go Swiss Ephemeris bindings package
func GetPackageVersion() string {
	return PackageVersion
}

// Close closes the Swiss Ephemeris and frees resources
func Close() {
	C.swe_close()
}

// SetEphePath sets the directory path for ephemeris files
func SetEphePath(path string) {
	if path == "" {
		C.swe_set_ephe_path(nil)
		return
	}
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	C.swe_set_ephe_path(cPath)
}

// SetJplFile sets the JPL ephemeris file name
func SetJplFile(fname string) {
	cFname := C.CString(fname)
	defer C.free(unsafe.Pointer(cFname))
	C.swe_set_jpl_file(cFname)
}

// GetLibraryPath returns the path where the library looks for ephemeris files
func GetLibraryPath() string {
	cPath := (*C.char)(C.malloc(C.size_t(asMaxch)))
	defer C.free(unsafe.Pointer(cPath))
	C.swe_get_library_path(cPath)
	return C.GoString(cPath)
}

// SetSidMode sets the sidereal mode for calculations
func SetSidMode(sidMode int32, t0 float64, ayanT0 float64) {
	C.swe_set_sid_mode(C.int(sidMode), C.double(t0), C.double(ayanT0))
}

// SetTopo sets the geographic location for topocentric calculations
func SetTopo(geoLon, geoLat, altitude float64) {
	C.swe_set_topo(C.double(geoLon), C.double(geoLat), C.double(altitude))
}

// SetTidAcc sets the tidal acceleration value
func SetTidAcc(tidAcc float64) {
	C.swe_set_tid_acc(C.double(tidAcc))
}

// GetTidAcc returns the current tidal acceleration value
func GetTidAcc() float64 {
	return float64(C.swe_get_tid_acc())
}

// SetDeltaTUserdef sets a user-defined Delta T value
func SetDeltaTUserdef(dt float64) {
	C.swe_set_delta_t_userdef(C.double(dt))
}

// Calc calculates planetary positions for a given Julian day (ephemeris time)
func Calc(tjdEt float64, ipl int32, iflag int32) CalcResult {
	var xx [6]C.double
	var serr [asMaxch]C.char

	flag := C.swe_calc(
		C.double(tjdEt),
		C.int(ipl),
		C.int(iflag),
		&xx[0],
		&serr[0],
	)

	result := CalcResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Data:  make([]float64, 6),
	}

	for i := 0; i < 6; i++ {
		result.Data[i] = float64(xx[i])
	}

	return result
}

// CalcUT calculates planetary positions for a given Julian day (universal time)
func CalcUT(tjdUt float64, ipl int32, iflag int32) CalcResult {
	var xx [6]C.double
	var serr [asMaxch]C.char

	flag := C.swe_calc_ut(
		C.double(tjdUt),
		C.int(ipl),
		C.int(iflag),
		&xx[0],
		&serr[0],
	)

	result := CalcResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Data:  make([]float64, 6),
	}

	for i := 0; i < 6; i++ {
		result.Data[i] = float64(xx[i])
	}

	return result
}

// CalcPctr calculates planetocentric positions
func CalcPctr(tjdEt float64, ipl int32, iplctr int32, iflag int32) CalcResult {
	var xx [6]C.double
	var serr [asMaxch]C.char

	flag := C.swe_calc_pctr(
		C.double(tjdEt),
		C.int(ipl),
		C.int(iplctr),
		C.int(iflag),
		&xx[0],
		&serr[0],
	)

	result := CalcResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Data:  make([]float64, 6),
	}

	for i := 0; i < 6; i++ {
		result.Data[i] = float64(xx[i])
	}

	return result
}

// Julday calculates the Julian day number from a calendar date
func Julday(year, month, day int32, hour float64, gregflag int32) float64 {
	return float64(C.swe_julday(
		C.int(year),
		C.int(month),
		C.int(day),
		C.double(hour),
		C.int(gregflag),
	))
}

// Revjul converts a Julian day number to a calendar date
func Revjul(jd float64, gregflag int32) DateResult {
	var year, month, day C.int
	var hour C.double

	C.swe_revjul(
		C.double(jd),
		C.int(gregflag),
		&year,
		&month,
		&day,
		&hour,
	)

	return DateResult{
		Year:  int(year),
		Month: int(month),
		Day:   int(day),
		Hour:  float64(hour),
	}
}

// DateConversion converts between Julian and Gregorian calendar
func DateConversion(year, month, day int32, hour float64, cal byte) (float64, error) {
	var tjd C.double

	result := C.swe_date_conversion(
		C.int(year),
		C.int(month),
		C.int(day),
		C.double(hour),
		C.char(cal),
		&tjd,
	)

	if result == ERR {
		return 0, fmt.Errorf("invalid date")
	}

	return float64(tjd), nil
}

// UtcToJd converts UTC to Julian day
func UtcToJd(year, month, day, hour, min int32, sec float64, gregflag int32) (dret [2]float64, err error) {
	var dretC [2]C.double
	var serr [asMaxch]C.char

	result := C.swe_utc_to_jd(
		C.int(year),
		C.int(month),
		C.int(day),
		C.int(hour),
		C.int(min),
		C.double(sec),
		C.int(gregflag),
		&dretC[0],
		&serr[0],
	)

	if result == ERR {
		return dret, fmt.Errorf("%s", C.GoString(&serr[0]))
	}

	dret[0] = float64(dretC[0])
	dret[1] = float64(dretC[1])

	return dret, nil
}

// JdetToUtc converts Julian day (ET) to UTC
func JdetToUtc(tjdEt float64, gregflag int32) UTCResult {
	var year, month, day, hour, min C.int
	var sec C.double

	C.swe_jdet_to_utc(
		C.double(tjdEt),
		C.int(gregflag),
		&year,
		&month,
		&day,
		&hour,
		&min,
		&sec,
	)

	return UTCResult{
		Year:   int(year),
		Month:  int(month),
		Day:    int(day),
		Hour:   int(hour),
		Minute: int(min),
		Second: float64(sec),
	}
}

// Jdut1ToUtc converts Julian day (UT1) to UTC
func Jdut1ToUtc(tjdUt float64, gregflag int32) UTCResult {
	var year, month, day, hour, min C.int
	var sec C.double

	C.swe_jdut1_to_utc(
		C.double(tjdUt),
		C.int(gregflag),
		&year,
		&month,
		&day,
		&hour,
		&min,
		&sec,
	)

	return UTCResult{
		Year:   int(year),
		Month:  int(month),
		Day:    int(day),
		Hour:   int(hour),
		Minute: int(min),
		Second: float64(sec),
	}
}

// UtcTimeZone converts local time to UTC or vice versa
func UtcTimeZone(year, month, day, hour, min int32, sec float64, timezone float64) UTCResult {
	var yearOut, monthOut, dayOut, hourOut, minOut C.int
	var secOut C.double

	C.swe_utc_time_zone(
		C.int(year),
		C.int(month),
		C.int(day),
		C.int(hour),
		C.int(min),
		C.double(sec),
		C.double(timezone),
		&yearOut,
		&monthOut,
		&dayOut,
		&hourOut,
		&minOut,
		&secOut,
	)

	return UTCResult{
		Year:   int(yearOut),
		Month:  int(monthOut),
		Day:    int(dayOut),
		Hour:   int(hourOut),
		Minute: int(minOut),
		Second: float64(secOut),
	}
}

// Houses calculates house cusps and other points
func Houses(tjdUt float64, geolat, geolon float64, hsys byte) HousesResult {
	var cusps [37]C.double
	var ascmc [10]C.double

	flag := C.swe_houses(
		C.double(tjdUt),
		C.double(geolat),
		C.double(geolon),
		C.int(hsys),
		&cusps[0],
		&ascmc[0],
	)

	// Determine number of houses based on system
	numHouses := 12
	if hsys == 'G' && flag == OK {
		numHouses = 36
	}

	result := HousesResult{
		Flag:   int32(flag),
		Houses: make([]float64, numHouses),
		Points: make([]float64, 8),
	}

	// Copy houses (skip index 0)
	for i := 0; i < numHouses; i++ {
		result.Houses[i] = float64(cusps[i+1])
	}

	// Copy points
	for i := 0; i < 8; i++ {
		result.Points[i] = float64(ascmc[i])
	}

	return result
}

// HousesEx calculates house cusps with extended options
func HousesEx(tjdUt float64, iflag int32, geolat, geolon, hsys byte) HousesResult {
	var cusps [37]C.double
	var ascmc [10]C.double

	flag := C.swe_houses_ex(
		C.double(tjdUt),
		C.int(iflag),
		C.double(geolat),
		C.double(geolon),
		C.int(hsys),
		&cusps[0],
		&ascmc[0],
	)

	numHouses := 12
	if hsys == 'G' && flag == OK {
		numHouses = 36
	}

	result := HousesResult{
		Flag:   int32(flag),
		Houses: make([]float64, numHouses),
		Points: make([]float64, 8),
	}

	for i := 0; i < numHouses; i++ {
		result.Houses[i] = float64(cusps[i+1])
	}

	for i := 0; i < 8; i++ {
		result.Points[i] = float64(ascmc[i])
	}

	return result
}

// HousesEx2 calculates house cusps with extended options (version 2)
func HousesEx2(tjdUt float64, iflag int32, geolat, geolon, hsys byte) HousesResult {
	var cusps [37]C.double
	var ascmc [10]C.double
	var cuspSpeed [37]C.double
	var ascmcSpeed [10]C.double
	var serr [asMaxch]C.char

	flag := C.swe_houses_ex2(
		C.double(tjdUt),
		C.int(iflag),
		C.double(geolat),
		C.double(geolon),
		C.int(hsys),
		&cusps[0],
		&ascmc[0],
		&cuspSpeed[0],
		&ascmcSpeed[0],
		&serr[0],
	)

	numHouses := 12
	if hsys == 'G' && flag == OK {
		numHouses = 36
	}

	result := HousesResult{
		Flag:   int32(flag),
		Houses: make([]float64, numHouses),
		Points: make([]float64, 8),
	}

	for i := 0; i < numHouses; i++ {
		result.Houses[i] = float64(cusps[i+1])
	}

	for i := 0; i < 8; i++ {
		result.Points[i] = float64(ascmc[i])
	}

	return result
}

// HousesArmc calculates house cusps from ARMC
func HousesArmc(armc float64, geolat float64, eps float64, hsys byte) HousesResult {
	var cusps [37]C.double
	var ascmc [10]C.double

	flag := C.swe_houses_armc(
		C.double(armc),
		C.double(geolat),
		C.double(eps),
		C.int(hsys),
		&cusps[0],
		&ascmc[0],
	)

	numHouses := 12
	if hsys == 'G' && flag == OK {
		numHouses = 36
	}

	result := HousesResult{
		Flag:   int32(flag),
		Houses: make([]float64, numHouses),
		Points: make([]float64, 8),
	}

	for i := 0; i < numHouses; i++ {
		result.Houses[i] = float64(cusps[i+1])
	}

	for i := 0; i < 8; i++ {
		result.Points[i] = float64(ascmc[i])
	}

	return result
}

// HousesArmcEx2 calculates house cusps from ARMC with extended options
func HousesArmcEx2(armc float64, geolat float64, eps float64, hsys byte) HousesResult {
	var cusps [37]C.double
	var ascmc [10]C.double
	var cuspSpeed [37]C.double
	var ascmcSpeed [10]C.double
	var serr [asMaxch]C.char

	flag := C.swe_houses_armc_ex2(
		C.double(armc),
		C.double(geolat),
		C.double(eps),
		C.int(hsys),
		&cusps[0],
		&ascmc[0],
		&cuspSpeed[0],
		&ascmcSpeed[0],
		&serr[0],
	)

	numHouses := 12
	if hsys == 'G' && flag == OK {
		numHouses = 36
	}

	result := HousesResult{
		Flag:   int32(flag),
		Houses: make([]float64, numHouses),
		Points: make([]float64, 8),
	}

	for i := 0; i < numHouses; i++ {
		result.Houses[i] = float64(cusps[i+1])
	}

	for i := 0; i < 8; i++ {
		result.Points[i] = float64(ascmc[i])
	}

	return result
}

// HousePos calculates the house position of a celestial point
func HousePos(armc float64, geolat float64, eps float64, hsys byte, lon float64, lat float64) (float64, error) {
	var xpin [2]C.double
	var serr [asMaxch]C.char

	xpin[0] = C.double(lon)
	xpin[1] = C.double(lat)

	pos := C.swe_house_pos(
		C.double(armc),
		C.double(geolat),
		C.double(eps),
		C.int(hsys),
		&xpin[0],
		&serr[0],
	)

	if pos < 0 {
		return 0, fmt.Errorf("%s", C.GoString(&serr[0]))
	}

	return float64(pos), nil
}

// HouseName returns the name of a house system
func HouseName(hsys byte) string {
	cName := C.swe_house_name(C.int(hsys))
	return C.GoString(cName)
}

// GetPlanetName returns the name of a planet
func GetPlanetName(ipl int32) string {
	var name [asMaxch]C.char
	C.swe_get_planet_name(C.int(ipl), &name[0])
	return C.GoString(&name[0])
}

// GetAyanamsaName returns the name of an ayanamsa
func GetAyanamsaName(sidMode int32) string {
	cName := C.swe_get_ayanamsa_name(C.int(sidMode))
	return C.GoString(cName)
}

// GetAyanamsa calculates the ayanamsa for a given Julian day (ephemeris time)
func GetAyanamsa(tjdEt float64) float64 {
	return float64(C.swe_get_ayanamsa(C.double(tjdEt)))
}

// GetAyanamsaUT calculates the ayanamsa for a given Julian day (universal time)
func GetAyanamsaUT(tjdUt float64) float64 {
	return float64(C.swe_get_ayanamsa_ut(C.double(tjdUt)))
}

// GetAyanamsaEx calculates the ayanamsa with extended information (ephemeris time)
func GetAyanamsaEx(tjdEt float64, iflag int32) CalcResult {
	var daya C.double
	var serr [asMaxch]C.char

	flag := C.swe_get_ayanamsa_ex(
		C.double(tjdEt),
		C.int(iflag),
		&daya,
		&serr[0],
	)

	return CalcResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Data:  []float64{float64(daya)},
	}
}

// GetAyanamsaExUT calculates the ayanamsa with extended information (universal time)
func GetAyanamsaExUT(tjdUt float64, iflag int32) CalcResult {
	var daya C.double
	var serr [asMaxch]C.char

	flag := C.swe_get_ayanamsa_ex_ut(
		C.double(tjdUt),
		C.int(iflag),
		&daya,
		&serr[0],
	)

	return CalcResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Data:  []float64{float64(daya)},
	}
}

// Sidtime calculates sidereal time
func Sidtime(tjdUt float64) float64 {
	return float64(C.swe_sidtime(C.double(tjdUt)))
}

// Sidtime0 calculates sidereal time at longitude 0
func Sidtime0(tjdUt float64, eps float64, nut float64) float64 {
	return float64(C.swe_sidtime0(C.double(tjdUt), C.double(eps), C.double(nut)))
}

// Deltat calculates Delta T (TT - UT) for a given Julian day
func Deltat(tjd float64) float64 {
	return float64(C.swe_deltat(C.double(tjd)))
}

// DeltatEx calculates Delta T with extended options
func DeltatEx(tjd float64, iflag int32) (float64, error) {
	var serr [asMaxch]C.char

	dt := C.swe_deltat_ex(
		C.double(tjd),
		C.int(iflag),
		&serr[0],
	)

	errStr := C.GoString(&serr[0])
	if errStr != "" {
		return float64(dt), fmt.Errorf("%s", errStr)
	}

	return float64(dt), nil
}

// TimeEqu calculates the equation of time
func TimeEqu(tjd float64) (float64, error) {
	var e C.double
	var serr [asMaxch]C.char

	result := C.swe_time_equ(
		C.double(tjd),
		&e,
		&serr[0],
	)

	if result == ERR {
		return 0, fmt.Errorf("%s", C.GoString(&serr[0]))
	}

	return float64(e), nil
}

// LmtToLat converts local mean time to local apparent time
func LmtToLat(tjdLmt float64, geolon float64) (float64, error) {
	var tjdLat C.double
	var serr [asMaxch]C.char

	result := C.swe_lmt_to_lat(
		C.double(tjdLmt),
		C.double(geolon),
		&tjdLat,
		&serr[0],
	)

	if result == ERR {
		return 0, fmt.Errorf("%s", C.GoString(&serr[0]))
	}

	return float64(tjdLat), nil
}

// LatToLmt converts local apparent time to local mean time
func LatToLmt(tjdLat float64, geolon float64) (float64, error) {
	var tjdLmt C.double
	var serr [asMaxch]C.char

	result := C.swe_lat_to_lmt(
		C.double(tjdLat),
		C.double(geolon),
		&tjdLmt,
		&serr[0],
	)

	if result == ERR {
		return 0, fmt.Errorf("%s", C.GoString(&serr[0]))
	}

	return float64(tjdLmt), nil
}

// DayOfWeek returns the day of week for a given Julian day (0=Monday, 6=Sunday)
func DayOfWeek(jd float64) int {
	return int(C.swe_day_of_week(C.double(jd)))
}
