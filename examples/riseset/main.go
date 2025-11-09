// Go Swiss Ephemeris - Rise/Set Example
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

package main

import (
	"fmt"
	"time"

	swisseph "github.com/tejzpr/go-swisseph"
)

func main() {
	fmt.Println("Swiss Ephemeris - Rise/Set Times Calculator")
	fmt.Println("===========================================\n")

	// Location: San Francisco
	lat := 37.7749
	lon := -122.4194
	altitude := 0.0
	locationName := "San Francisco, CA"

	fmt.Printf("Location: %s\n", locationName)
	fmt.Printf("Coordinates: %.4f°N, %.4f°W\n", lat, -lon)
	fmt.Printf("Altitude: %.0f meters\n\n", altitude)

	// Set topocentric location
	swisseph.SetTopo(lon, lat, altitude)

	// Use current date
	now := time.Now().UTC()
	jd := swisseph.Julday(
		int32(now.Year()),
		int32(now.Month()),
		int32(now.Day()),
		0.0, // Start of day
		swisseph.GregCal,
	)

	fmt.Printf("Date: %s\n\n", now.Format("2006-01-02"))

	// Atmospheric conditions
	atpress := 1013.25 // Standard atmospheric pressure in mbar
	attemp := 15.0     // Temperature in Celsius

	geopos := [3]float64{lon, lat, altitude}

	// Calculate for Sun
	fmt.Println("Sun")
	fmt.Println("---")
	calculateRiseSetTransit(jd, swisseph.Sun, "", geopos, atpress, attemp)

	// Calculate for Moon
	fmt.Println("\nMoon")
	fmt.Println("----")
	calculateRiseSetTransit(jd, swisseph.Moon, "", geopos, atpress, attemp)

	// Calculate for planets
	planets := []struct {
		id   int32
		name string
	}{
		{swisseph.Mercury, "Mercury"},
		{swisseph.Venus, "Venus"},
		{swisseph.Mars, "Mars"},
		{swisseph.Jupiter, "Jupiter"},
		{swisseph.Saturn, "Saturn"},
	}

	for _, planet := range planets {
		fmt.Printf("\n%s\n", planet.name)
		fmt.Println("--------")
		calculateRiseSetTransit(jd, planet.id, "", geopos, atpress, attemp)
	}

	// Calculate twilight times
	fmt.Println("\nTwilight Times")
	fmt.Println("--------------")
	calculateTwilight(jd, geopos, atpress, attemp)

	// Calculate for a bright star
	fmt.Println("\nSirius (Brightest Star)")
	fmt.Println("-----------------------")
	calculateRiseSetTransit(jd, swisseph.FixstarFlag, "Sirius", geopos, atpress, attemp)

	swisseph.Close()
	fmt.Println("\nDone!")
}

func calculateRiseSetTransit(jd float64, ipl int32, starname string, geopos [3]float64, atpress float64, attemp float64) {
	// Calculate rise
	rise := swisseph.RiseTrans(jd, ipl, starname, swisseph.FlagSwieph,
		swisseph.CalcRise, geopos, atpress, attemp)

	if rise.Flag >= 0 {
		riseTime := swisseph.Revjul(rise.Time, swisseph.GregCal)
		fmt.Printf("Rise:    %02.0f:%02.0f UTC\n",
			riseTime.Hour, (riseTime.Hour-float64(int(riseTime.Hour)))*60)
	} else if rise.Error != "" {
		fmt.Printf("Rise:    %s\n", rise.Error)
	}

	// Calculate transit (culmination)
	transit := swisseph.RiseTrans(jd, ipl, starname, swisseph.FlagSwieph,
		swisseph.CalcMtransit, geopos, atpress, attemp)

	if transit.Flag >= 0 {
		transitTime := swisseph.Revjul(transit.Time, swisseph.GregCal)
		fmt.Printf("Transit: %02.0f:%02.0f UTC\n",
			transitTime.Hour, (transitTime.Hour-float64(int(transitTime.Hour)))*60)

		// Calculate altitude at transit
		result := swisseph.CalcUT(transit.Time, ipl, swisseph.FlagSwieph|swisseph.FlagEquatorial)
		if result.Flag >= 0 {
			// Convert to horizontal coordinates
			xin := [3]float64{result.Data[0], result.Data[1], result.Data[2]}
			azalt := swisseph.Azalt(transit.Time, swisseph.Equ2Hor, geopos, atpress, attemp, xin)
			fmt.Printf("         (Altitude: %.2f°)\n", azalt.Altitude)
		}
	} else if transit.Error != "" {
		fmt.Printf("Transit: %s\n", transit.Error)
	}

	// Calculate set
	set := swisseph.RiseTrans(jd, ipl, starname, swisseph.FlagSwieph,
		swisseph.CalcSet, geopos, atpress, attemp)

	if set.Flag >= 0 {
		setTime := swisseph.Revjul(set.Time, swisseph.GregCal)
		fmt.Printf("Set:     %02.0f:%02.0f UTC\n",
			setTime.Hour, (setTime.Hour-float64(int(setTime.Hour)))*60)
	} else if set.Error != "" {
		fmt.Printf("Set:     %s\n", set.Error)
	}

	// Calculate day length for Sun and Moon
	if ipl == swisseph.Sun || ipl == swisseph.Moon {
		if rise.Flag >= 0 && set.Flag >= 0 {
			dayLength := (set.Time - rise.Time) * 24 // Convert to hours
			hours := int(dayLength)
			minutes := int((dayLength - float64(hours)) * 60)
			fmt.Printf("Length:  %dh %02dm\n", hours, minutes)
		}
	}
}

func calculateTwilight(jd float64, geopos [3]float64, atpress float64, attemp float64) {
	// Civil twilight (Sun 6° below horizon)
	civilStart := swisseph.RiseTrans(jd, swisseph.Sun, "", swisseph.FlagSwieph,
		swisseph.CalcRise|swisseph.BitCivilTwilight, geopos, atpress, attemp)

	if civilStart.Flag >= 0 {
		time := swisseph.Revjul(civilStart.Time, swisseph.GregCal)
		fmt.Printf("Civil Dawn:      %02.0f:%02.0f UTC\n",
			time.Hour, (time.Hour-float64(int(time.Hour)))*60)
	}

	// Nautical twilight (Sun 12° below horizon)
	nauticalStart := swisseph.RiseTrans(jd, swisseph.Sun, "", swisseph.FlagSwieph,
		swisseph.CalcRise|swisseph.BitNauticTwilight, geopos, atpress, attemp)

	if nauticalStart.Flag >= 0 {
		time := swisseph.Revjul(nauticalStart.Time, swisseph.GregCal)
		fmt.Printf("Nautical Dawn:   %02.0f:%02.0f UTC\n",
			time.Hour, (time.Hour-float64(int(time.Hour)))*60)
	}

	// Astronomical twilight (Sun 18° below horizon)
	astroStart := swisseph.RiseTrans(jd, swisseph.Sun, "", swisseph.FlagSwieph,
		swisseph.CalcRise|swisseph.BitAstroTwilight, geopos, atpress, attemp)

	if astroStart.Flag >= 0 {
		time := swisseph.Revjul(astroStart.Time, swisseph.GregCal)
		fmt.Printf("Astronomical Dawn: %02.0f:%02.0f UTC\n",
			time.Hour, (time.Hour-float64(int(time.Hour)))*60)
	}

	// Evening twilights
	civilEnd := swisseph.RiseTrans(jd, swisseph.Sun, "", swisseph.FlagSwieph,
		swisseph.CalcSet|swisseph.BitCivilTwilight, geopos, atpress, attemp)

	if civilEnd.Flag >= 0 {
		time := swisseph.Revjul(civilEnd.Time, swisseph.GregCal)
		fmt.Printf("Civil Dusk:      %02.0f:%02.0f UTC\n",
			time.Hour, (time.Hour-float64(int(time.Hour)))*60)
	}

	nauticalEnd := swisseph.RiseTrans(jd, swisseph.Sun, "", swisseph.FlagSwieph,
		swisseph.CalcSet|swisseph.BitNauticTwilight, geopos, atpress, attemp)

	if nauticalEnd.Flag >= 0 {
		time := swisseph.Revjul(nauticalEnd.Time, swisseph.GregCal)
		fmt.Printf("Nautical Dusk:   %02.0f:%02.0f UTC\n",
			time.Hour, (time.Hour-float64(int(time.Hour)))*60)
	}

	astroEnd := swisseph.RiseTrans(jd, swisseph.Sun, "", swisseph.FlagSwieph,
		swisseph.CalcSet|swisseph.BitAstroTwilight, geopos, atpress, attemp)

	if astroEnd.Flag >= 0 {
		time := swisseph.Revjul(astroEnd.Time, swisseph.GregCal)
		fmt.Printf("Astronomical Dusk: %02.0f:%02.0f UTC\n",
			time.Hour, (time.Hour-float64(int(time.Hour)))*60)
	}
}
