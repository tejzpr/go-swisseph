// Go Swiss Ephemeris - Utilities
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

// RiseTrans calculates rise, set, and transit times
func RiseTrans(tjdUt float64, ipl int32, starname string, epheflag int32, rsmi int32, geopos [3]float64, atpress float64, attemp float64) RiseTransResult {
	var tret C.double
	var serr [asMaxch]C.char

	var geoposC [3]C.double
	for i := 0; i < 3; i++ {
		geoposC[i] = C.double(geopos[i])
	}

	cStar := C.CString(starname)
	defer C.free(unsafe.Pointer(cStar))

	flag := C.swe_rise_trans(
		C.double(tjdUt),
		C.int(ipl),
		cStar,
		C.int(epheflag),
		C.int(rsmi),
		&geoposC[0],
		C.double(atpress),
		C.double(attemp),
		&tret,
		&serr[0],
	)

	return RiseTransResult{
		Flag:  int32(flag),
		Time:  float64(tret),
		Error: C.GoString(&serr[0]),
	}
}

// RiseTransTrueHor calculates rise, set, and transit times with true horizon
func RiseTransTrueHor(tjdUt float64, ipl int32, starname string, epheflag int32, rsmi int32, geopos [3]float64, atpress float64, attemp float64, horhgt float64) RiseTransResult {
	var tret C.double
	var serr [asMaxch]C.char

	var geoposC [3]C.double
	for i := 0; i < 3; i++ {
		geoposC[i] = C.double(geopos[i])
	}

	cStar := C.CString(starname)
	defer C.free(unsafe.Pointer(cStar))

	flag := C.swe_rise_trans_true_hor(
		C.double(tjdUt),
		C.int(ipl),
		cStar,
		C.int(epheflag),
		C.int(rsmi),
		&geoposC[0],
		C.double(atpress),
		C.double(attemp),
		C.double(horhgt),
		&tret,
		&serr[0],
	)

	return RiseTransResult{
		Flag:  int32(flag),
		Time:  float64(tret),
		Error: C.GoString(&serr[0]),
	}
}

// Pheno calculates planetary phenomena (ephemeris time)
func Pheno(tjdEt float64, ipl int32, iflag int32) CalcResult {
	var attr [20]C.double
	var serr [asMaxch]C.char

	flag := C.swe_pheno(
		C.double(tjdEt),
		C.int(ipl),
		C.int(iflag),
		&attr[0],
		&serr[0],
	)

	result := CalcResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Data:  make([]float64, 20),
	}

	for i := 0; i < 20; i++ {
		result.Data[i] = float64(attr[i])
	}

	return result
}

// PhenoUT calculates planetary phenomena (universal time)
func PhenoUT(tjdUt float64, ipl int32, iflag int32) CalcResult {
	var attr [20]C.double
	var serr [asMaxch]C.char

	flag := C.swe_pheno_ut(
		C.double(tjdUt),
		C.int(ipl),
		C.int(iflag),
		&attr[0],
		&serr[0],
	)

	result := CalcResult{
		Flag:  int32(flag),
		Error: C.GoString(&serr[0]),
		Data:  make([]float64, 20),
	}

	for i := 0; i < 20; i++ {
		result.Data[i] = float64(attr[i])
	}

	return result
}

// Azalt calculates azimuth and altitude from ecliptic or equatorial coordinates
func Azalt(tjdUt float64, calcflag int32, geopos [3]float64, atpress float64, attemp float64, xin [3]float64) AzaltResult {
	var xinC [3]C.double
	var xaz [3]C.double

	for i := 0; i < 3; i++ {
		xinC[i] = C.double(xin[i])
	}

	var geoposC [3]C.double
	for i := 0; i < 3; i++ {
		geoposC[i] = C.double(geopos[i])
	}

	C.swe_azalt(
		C.double(tjdUt),
		C.int(calcflag),
		&geoposC[0],
		C.double(atpress),
		C.double(attemp),
		&xinC[0],
		&xaz[0],
	)

	return AzaltResult{
		Azimuth:  float64(xaz[0]),
		Altitude: float64(xaz[1]),
		AppAlt:   float64(xaz[2]),
	}
}

// AzaltRev calculates ecliptic or equatorial coordinates from azimuth and altitude
func AzaltRev(tjdUt float64, calcflag int32, geopos [3]float64, xin [3]float64) [3]float64 {
	var xinC [3]C.double
	var xout [3]C.double

	for i := 0; i < 3; i++ {
		xinC[i] = C.double(xin[i])
	}

	var geoposC [3]C.double
	for i := 0; i < 3; i++ {
		geoposC[i] = C.double(geopos[i])
	}

	C.swe_azalt_rev(
		C.double(tjdUt),
		C.int(calcflag),
		&geoposC[0],
		&xinC[0],
		&xout[0],
	)

	var result [3]float64
	for i := 0; i < 3; i++ {
		result[i] = float64(xout[i])
	}

	return result
}

// Refrac calculates atmospheric refraction
func Refrac(inalt float64, atpress float64, attemp float64, calcflag int32) float64 {
	return float64(C.swe_refrac(
		C.double(inalt),
		C.double(atpress),
		C.double(attemp),
		C.int(calcflag),
	))
}

// RefracExtended calculates atmospheric refraction with extended parameters
func RefracExtended(inalt float64, geoalt float64, atpress float64, attemp float64, lapserate float64, calcflag int32) (float64, [4]float64) {
	var dret [4]C.double

	refr := C.swe_refrac_extended(
		C.double(inalt),
		C.double(geoalt),
		C.double(atpress),
		C.double(attemp),
		C.double(lapserate),
		C.int(calcflag),
		&dret[0],
	)

	var result [4]float64
	for i := 0; i < 4; i++ {
		result[i] = float64(dret[i])
	}

	return float64(refr), result
}

// Cotrans transforms coordinates
func Cotrans(xpo [3]float64, eps float64) [3]float64 {
	var xpoC [3]C.double
	var xpn [3]C.double

	for i := 0; i < 3; i++ {
		xpoC[i] = C.double(xpo[i])
	}

	C.swe_cotrans(&xpoC[0], &xpn[0], C.double(eps))

	var result [3]float64
	for i := 0; i < 3; i++ {
		result[i] = float64(xpn[i])
	}

	return result
}

// CotransSp transforms coordinates with speed
func CotransSp(xpo [6]float64, eps float64) [6]float64 {
	var xpoC [6]C.double
	var xpn [6]C.double

	for i := 0; i < 6; i++ {
		xpoC[i] = C.double(xpo[i])
	}

	C.swe_cotrans_sp(&xpoC[0], &xpn[0], C.double(eps))

	var result [6]float64
	for i := 0; i < 6; i++ {
		result[i] = float64(xpn[i])
	}

	return result
}

// Degnorm normalizes degrees to 0-360 range
func Degnorm(x float64) float64 {
	return float64(C.swe_degnorm(C.double(x)))
}

// Radnorm normalizes radians to 0-2π range
func Radnorm(x float64) float64 {
	return float64(C.swe_radnorm(C.double(x)))
}

// Csnorm normalizes centiseconds to 0-360° range
func Csnorm(p int32) int32 {
	return int32(C.swe_csnorm(C.int(p)))
}

// Difcsn calculates normalized difference in centiseconds
func Difcsn(p1 int32, p2 int32) int32 {
	return int32(C.swe_difcsn(C.int(p1), C.int(p2)))
}

// Difdegn calculates normalized difference in degrees
func Difdegn(p1 float64, p2 float64) float64 {
	return float64(C.swe_difdegn(C.double(p1), C.double(p2)))
}

// Difcs2n calculates normalized difference in centiseconds (2)
func Difcs2n(p1 int32, p2 int32) int32 {
	return int32(C.swe_difcs2n(C.int(p1), C.int(p2)))
}

// Difdeg2n calculates normalized difference in degrees (2)
func Difdeg2n(p1 float64, p2 float64) float64 {
	return float64(C.swe_difdeg2n(C.double(p1), C.double(p2)))
}

// Csroundsec rounds centiseconds
func Csroundsec(x int32) int32 {
	return int32(C.swe_csroundsec(C.int(x)))
}

// D2l converts double to int32
func D2l(x float64) int32 {
	return int32(C.swe_d2l(C.double(x)))
}

// SplitDeg splits degrees into components
func SplitDeg(ddeg float64, roundflag int32) SplitDegResult {
	var ideg, imin, isec C.int
	var dsecfr C.double
	var isgn C.int

	C.swe_split_deg(
		C.double(ddeg),
		C.int(roundflag),
		&ideg,
		&imin,
		&isec,
		&dsecfr,
		&isgn,
	)

	return SplitDegResult{
		Degree:     int32(ideg),
		Minute:     int32(imin),
		Second:     int32(isec),
		SecondFrac: float64(dsecfr),
		Sign:       int32(isgn),
	}
}

// Cs2timestr converts centiseconds to time string
func Cs2timestr(t int32, sep byte, suppressZero bool) string {
	var s [asMaxch]C.char

	suppr := C.int(0)
	if suppressZero {
		suppr = 1
	}

	C.swe_cs2timestr(C.int(t), C.int(sep), suppr, &s[0])
	return C.GoString(&s[0])
}

// Cs2lonlatstr converts centiseconds to longitude/latitude string
func Cs2lonlatstr(t int32, pchar byte, mchar byte) string {
	var s [asMaxch]C.char

	C.swe_cs2lonlatstr(C.int(t), C.char(pchar), C.char(mchar), &s[0])
	return C.GoString(&s[0])
}

// Cs2degstr converts centiseconds to degree string
func Cs2degstr(t int32) string {
	var s [asMaxch]C.char

	C.swe_cs2degstr(C.int(t), &s[0])
	return C.GoString(&s[0])
}

// NodAps calculates nodes and apsides (ephemeris time)
func NodAps(tjdEt float64, ipl int32, iflag int32, method int32) NodApsResult {
	var xnasc [6]C.double
	var xndsc [6]C.double
	var xperi [6]C.double
	var xaphe [6]C.double
	var serr [asMaxch]C.char

	flag := C.swe_nod_aps(
		C.double(tjdEt),
		C.int(ipl),
		C.int(iflag),
		C.int(method),
		&xnasc[0],
		&xndsc[0],
		&xperi[0],
		&xaphe[0],
		&serr[0],
	)

	result := NodApsResult{
		Flag:       int32(flag),
		Error:      C.GoString(&serr[0]),
		Ascending:  make([]float64, 6),
		Descending: make([]float64, 6),
		Perihelion: make([]float64, 6),
		Aphelion:   make([]float64, 6),
	}

	for i := 0; i < 6; i++ {
		result.Ascending[i] = float64(xnasc[i])
		result.Descending[i] = float64(xndsc[i])
		result.Perihelion[i] = float64(xperi[i])
		result.Aphelion[i] = float64(xaphe[i])
	}

	return result
}

// NodApsUT calculates nodes and apsides (universal time)
func NodApsUT(tjdUt float64, ipl int32, iflag int32, method int32) NodApsResult {
	var xnasc [6]C.double
	var xndsc [6]C.double
	var xperi [6]C.double
	var xaphe [6]C.double
	var serr [asMaxch]C.char

	flag := C.swe_nod_aps_ut(
		C.double(tjdUt),
		C.int(ipl),
		C.int(iflag),
		C.int(method),
		&xnasc[0],
		&xndsc[0],
		&xperi[0],
		&xaphe[0],
		&serr[0],
	)

	result := NodApsResult{
		Flag:       int32(flag),
		Error:      C.GoString(&serr[0]),
		Ascending:  make([]float64, 6),
		Descending: make([]float64, 6),
		Perihelion: make([]float64, 6),
		Aphelion:   make([]float64, 6),
	}

	for i := 0; i < 6; i++ {
		result.Ascending[i] = float64(xnasc[i])
		result.Descending[i] = float64(xndsc[i])
		result.Perihelion[i] = float64(xperi[i])
		result.Aphelion[i] = float64(xaphe[i])
	}

	return result
}

// GetOrbitalElements calculates orbital elements
func GetOrbitalElements(tjdEt float64, ipl int32, iflag int32) OrbitalElementsResult {
	var dret [50]C.double
	var serr [asMaxch]C.char

	flag := C.swe_get_orbital_elements(
		C.double(tjdEt),
		C.int(ipl),
		C.int(iflag),
		&dret[0],
		&serr[0],
	)

	result := OrbitalElementsResult{
		Flag:     int32(flag),
		Error:    C.GoString(&serr[0]),
		Elements: make([]float64, 50),
	}

	for i := 0; i < 50; i++ {
		result.Elements[i] = float64(dret[i])
	}

	return result
}

// OrbitMaxMinTrueDistance calculates maximum and minimum true distance
func OrbitMaxMinTrueDistance(tjdEt float64, ipl int32, iflag int32) (dmax float64, dmin float64, dtrue float64, err error) {
	var dmaxC, dminC, dtrueC C.double
	var serr [asMaxch]C.char

	result := C.swe_orbit_max_min_true_distance(
		C.double(tjdEt),
		C.int(ipl),
		C.int(iflag),
		&dmaxC,
		&dminC,
		&dtrueC,
		&serr[0],
	)

	if result == ERR {
		return 0, 0, 0, fmt.Errorf("%s", C.GoString(&serr[0]))
	}

	return float64(dmaxC), float64(dminC), float64(dtrueC), nil
}

// GetCurrentFileData gets information about currently used ephemeris files
func GetCurrentFileData(ifno int32) FileData {
	var tfstart, tfend C.double
	var denum C.int

	cPath := C.swe_get_current_file_data(
		C.int(ifno),
		&tfstart,
		&tfend,
		&denum,
	)

	return FileData{
		Path:      C.GoString(cPath),
		StartDate: float64(tfstart),
		EndDate:   float64(tfend),
		Denum:     int32(denum),
	}
}

// GauquelinSector calculates Gauquelin sector position
func GauquelinSector(tjdUt float64, ipl int32, starname string, iflag int32, imeth int32, geopos [3]float64, atpress float64, attemp float64) (float64, error) {
	var dgsect C.double
	var serr [asMaxch]C.char

	var geoposC [3]C.double
	for i := 0; i < 3; i++ {
		geoposC[i] = C.double(geopos[i])
	}

	cStar := C.CString(starname)
	defer C.free(unsafe.Pointer(cStar))

	result := C.swe_gauquelin_sector(
		C.double(tjdUt),
		C.int(ipl),
		cStar,
		C.int(iflag),
		C.int(imeth),
		&geoposC[0],
		C.double(atpress),
		C.double(attemp),
		&dgsect,
		&serr[0],
	)

	if result == ERR {
		return 0, fmt.Errorf("%s", C.GoString(&serr[0]))
	}

	return float64(dgsect), nil
}
