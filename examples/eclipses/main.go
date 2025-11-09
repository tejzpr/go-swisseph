// Go Swiss Ephemeris - Eclipses Example
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
	fmt.Println("Swiss Ephemeris - Eclipse Calculator")
	fmt.Println("====================================\n")

	// Start from current date
	now := time.Now().UTC()
	jd := swisseph.Julday(
		int32(now.Year()),
		int32(now.Month()),
		int32(now.Day()),
		12.0,
		swisseph.GregCal,
	)

	fmt.Printf("Searching for eclipses from: %s\n\n", now.Format("2006-01-02"))

	// Find next 5 solar eclipses
	fmt.Println("Next Solar Eclipses")
	fmt.Println("-------------------")

	searchJd := jd
	for i := 0; i < 5; i++ {
		eclipse := swisseph.SolEclipseWhenGlob(searchJd, swisseph.FlagSwieph,
			swisseph.EclAlltypesSolar, false)

		if eclipse.Flag >= 0 {
			date := swisseph.Revjul(eclipse.Maximum, swisseph.GregCal)

			eclipseType := getEclipseType(eclipse.Flag)
			fmt.Printf("%d. %04d-%02d-%02d %02.0f:%02.0f UTC - %s\n",
				i+1, date.Year, date.Month, date.Day,
				date.Hour, (date.Hour-float64(int(date.Hour)))*60,
				eclipseType)

			// Get location where eclipse is central/maximal
			where := swisseph.SolEclipseWhere(eclipse.Maximum, swisseph.FlagSwieph)
			if where.Flag >= 0 {
				fmt.Printf("   Central at: %.4f°N, %.4f°E\n",
					where.Latitude, where.Longitude)
			}

			searchJd = eclipse.Maximum + 1 // Search for next eclipse
		} else {
			fmt.Printf("Error: %s\n", eclipse.Error)
			break
		}
	}

	// Find next 5 lunar eclipses
	fmt.Println("\nNext Lunar Eclipses")
	fmt.Println("-------------------")

	searchJd = jd
	for i := 0; i < 5; i++ {
		eclipse := swisseph.LunEclipseWhen(searchJd, swisseph.FlagSwieph,
			swisseph.EclAlltypesLunar, false)

		if eclipse.Flag >= 0 {
			date := swisseph.Revjul(eclipse.Maximum, swisseph.GregCal)

			eclipseType := getLunarEclipseType(eclipse.Flag)
			fmt.Printf("%d. %04d-%02d-%02d %02.0f:%02.0f UTC - %s\n",
				i+1, date.Year, date.Month, date.Day,
				date.Hour, (date.Hour-float64(int(date.Hour)))*60,
				eclipseType)

			searchJd = eclipse.Maximum + 1
		} else {
			fmt.Printf("Error: %s\n", eclipse.Error)
			break
		}
	}

	// Check eclipse visibility for a specific location
	fmt.Println("\nEclipse Visibility Check")
	fmt.Println("------------------------")
	fmt.Println("Location: New York City (40.7128°N, 74.0060°W)")

	geopos := [3]float64{-74.0060, 40.7128, 0}
	localEclipse := swisseph.SolEclipseWhenLoc(jd, swisseph.FlagSwieph,
		geopos, false)

	if localEclipse.Flag >= 0 {
		date := swisseph.Revjul(localEclipse.Maximum, swisseph.GregCal)
		eclipseType := getEclipseType(localEclipse.Flag)

		fmt.Printf("\nNext visible solar eclipse:\n")
		fmt.Printf("Date: %04d-%02d-%02d\n", date.Year, date.Month, date.Day)
		fmt.Printf("Type: %s\n", eclipseType)

		// Eclipse times
		if localEclipse.Begin > 0 {
			beginDate := swisseph.Revjul(localEclipse.Begin, swisseph.GregCal)
			fmt.Printf("Begin: %02.0f:%02.0f UTC\n",
				beginDate.Hour, (beginDate.Hour-float64(int(beginDate.Hour)))*60)
		}

		maxDate := swisseph.Revjul(localEclipse.Maximum, swisseph.GregCal)
		fmt.Printf("Maximum: %02.0f:%02.0f UTC\n",
			maxDate.Hour, (maxDate.Hour-float64(int(maxDate.Hour)))*60)

		if localEclipse.End > 0 {
			endDate := swisseph.Revjul(localEclipse.End, swisseph.GregCal)
			fmt.Printf("End: %02.0f:%02.0f UTC\n",
				endDate.Hour, (endDate.Hour-float64(int(endDate.Hour)))*60)
		}

		// Eclipse magnitude
		if len(localEclipse.Attr) > 0 {
			fmt.Printf("Magnitude: %.4f\n", localEclipse.Attr[0])
		}
	}

	swisseph.Close()
	fmt.Println("\nDone!")
}

func getEclipseType(flag int32) string {
	if flag&swisseph.EclTotal != 0 {
		return "Total Solar Eclipse"
	} else if flag&swisseph.EclAnnular != 0 {
		return "Annular Solar Eclipse"
	} else if flag&swisseph.EclAnnularTotal != 0 {
		return "Hybrid Solar Eclipse"
	} else if flag&swisseph.EclPartial != 0 {
		return "Partial Solar Eclipse"
	}
	return "Solar Eclipse"
}

func getLunarEclipseType(flag int32) string {
	if flag&swisseph.EclTotal != 0 {
		return "Total Lunar Eclipse"
	} else if flag&swisseph.EclPartial != 0 {
		return "Partial Lunar Eclipse"
	} else if flag&swisseph.EclPenumbral != 0 {
		return "Penumbral Lunar Eclipse"
	}
	return "Lunar Eclipse"
}

