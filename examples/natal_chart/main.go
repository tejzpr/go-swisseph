// Go Swiss Ephemeris - Natal Chart Example
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

	swisseph "github.com/tejzpr/go-swisseph"
)

func main() {
	fmt.Println("Swiss Ephemeris - Natal Chart Example")
	fmt.Println("======================================\n")

	// Birth data
	year := int32(1990)
	month := int32(7)
	day := int32(15)
	hour := 14.5 // 14:30

	// Birth location (example: London)
	lat := 51.5074
	lon := -0.1278
	locationName := "London, UK"

	fmt.Printf("Birth Date: %04d-%02d-%02d %02.0f:%02.0f\n", year, month, day, hour, (hour-float64(int(hour)))*60)
	fmt.Printf("Birth Place: %s (%.4f°N, %.4f°E)\n\n", locationName, lat, lon)

	// Calculate Julian day
	jd := swisseph.Julday(year, month, day, hour, swisseph.GregCal)
	fmt.Printf("Julian Day: %.6f\n\n", jd)

	// Set topocentric location
	swisseph.SetTopo(lon, lat, 0)

	// Calculate house cusps
	fmt.Println("House Cusps (Placidus)")
	fmt.Println("----------------------")
	houses := swisseph.Houses(jd, lat, lon, 'P')

	if houses.Flag >= 0 {
		for i, cusp := range houses.Houses {
			fmt.Printf("House %2d: %s\n", i+1, formatDegree(cusp))
		}

		fmt.Printf("\nAngles:\n")
		fmt.Printf("Ascendant (ASC): %s\n", formatDegree(houses.Points[swisseph.Asc]))
		fmt.Printf("Midheaven (MC):  %s\n", formatDegree(houses.Points[swisseph.MC]))
		fmt.Printf("Descendant (DSC): %s\n", formatDegree(swisseph.Degnorm(houses.Points[swisseph.Asc]+180)))
		fmt.Printf("Imum Coeli (IC):  %s\n", formatDegree(swisseph.Degnorm(houses.Points[swisseph.MC]+180)))
	}

	// Calculate planetary positions
	fmt.Println("\nPlanetary Positions")
	fmt.Println("-------------------")

	planets := []struct {
		id   int32
		name string
	}{
		{swisseph.Sun, "Sun"},
		{swisseph.Moon, "Moon"},
		{swisseph.Mercury, "Mercury"},
		{swisseph.Venus, "Venus"},
		{swisseph.Mars, "Mars"},
		{swisseph.Jupiter, "Jupiter"},
		{swisseph.Saturn, "Saturn"},
		{swisseph.Uranus, "Uranus"},
		{swisseph.Neptune, "Neptune"},
		{swisseph.Pluto, "Pluto"},
		{swisseph.TrueNode, "North Node"},
		{swisseph.Chiron, "Chiron"},
	}

	for _, planet := range planets {
		result := swisseph.CalcUT(jd, planet.id, swisseph.FlagSwieph|swisseph.FlagSpeed)
		if result.Flag >= 0 {
			longitude := result.Data[0]
			speedLon := result.Data[3]

			// Determine if retrograde
			retrograde := ""
			if speedLon < 0 {
				retrograde = " (R)"
			}

			// Find house placement
			house := findHouse(longitude, houses.Houses)

			fmt.Printf("%-12s: %s%s in House %d\n",
				planet.name, formatDegree(longitude), retrograde, house)
		}
	}

	// Calculate aspects between planets
	fmt.Println("\nMajor Aspects")
	fmt.Println("-------------")

	// Store planet positions
	positions := make(map[string]float64)
	for _, planet := range planets[:10] { // Only major planets for aspects
		result := swisseph.CalcUT(jd, planet.id, swisseph.FlagSwieph)
		if result.Flag >= 0 {
			positions[planet.name] = result.Data[0]
		}
	}

	// Calculate aspects
	aspects := []struct {
		name  string
		angle float64
		orb   float64
	}{
		{"Conjunction", 0, 8},
		{"Sextile", 60, 6},
		{"Square", 90, 8},
		{"Trine", 120, 8},
		{"Opposition", 180, 8},
	}

	planetNames := []string{"Sun", "Moon", "Mercury", "Venus", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune", "Pluto"}

	for i := 0; i < len(planetNames); i++ {
		for j := i + 1; j < len(planetNames); j++ {
			p1 := planetNames[i]
			p2 := planetNames[j]
			pos1 := positions[p1]
			pos2 := positions[p2]

			diff := swisseph.Difdegn(pos1, pos2)
			if diff < 0 {
				diff = -diff
			}

			for _, aspect := range aspects {
				aspectDiff := diff - aspect.angle
				if aspectDiff < 0 {
					aspectDiff = -aspectDiff
				}

				if aspectDiff <= aspect.orb {
					fmt.Printf("%-10s %-12s %-10s (orb: %.2f°)\n",
						p1, aspect.name, p2, aspectDiff)
				}
			}
		}
	}

	// Calculate lunar phase
	sunResult := swisseph.CalcUT(jd, swisseph.Sun, swisseph.FlagSwieph)
	moonResult := swisseph.CalcUT(jd, swisseph.Moon, swisseph.FlagSwieph)
	if sunResult.Flag >= 0 && moonResult.Flag >= 0 {
		phase := swisseph.Degnorm(moonResult.Data[0] - sunResult.Data[0])
		phaseName := getLunarPhaseName(phase)
		fmt.Printf("\nLunar Phase: %s (%.2f°)\n", phaseName, phase)
	}

	swisseph.Close()
	fmt.Println("\nDone!")
}

func formatDegree(deg float64) string {
	signs := []string{"Aries", "Taurus", "Gemini", "Cancer", "Leo", "Virgo",
		"Libra", "Scorpio", "Sagittarius", "Capricorn", "Aquarius", "Pisces"}
	symbols := []string{"♈", "♉", "♊", "♋", "♌", "♍", "♎", "♏", "♐", "♑", "♒", "♓"}

	sign := int(deg / 30)
	degree := deg - float64(sign)*30
	split := swisseph.SplitDeg(degree, 0)

	return fmt.Sprintf("%2d° %02d' %02d\" %s %s",
		split.Degree, split.Minute, split.Second, signs[sign], symbols[sign])
}

func findHouse(longitude float64, houses []float64) int {
	for i := 0; i < len(houses); i++ {
		nextHouse := houses[(i+1)%len(houses)]
		currentHouse := houses[i]

		// Handle wrap-around at 360°
		if nextHouse < currentHouse {
			if longitude >= currentHouse || longitude < nextHouse {
				return i + 1
			}
		} else {
			if longitude >= currentHouse && longitude < nextHouse {
				return i + 1
			}
		}
	}
	return 1
}

func getLunarPhaseName(phase float64) string {
	if phase < 45 {
		return "New Moon"
	} else if phase < 90 {
		return "Waxing Crescent"
	} else if phase < 135 {
		return "First Quarter"
	} else if phase < 180 {
		return "Waxing Gibbous"
	} else if phase < 225 {
		return "Full Moon"
	} else if phase < 270 {
		return "Waning Gibbous"
	} else if phase < 315 {
		return "Last Quarter"
	} else {
		return "Waning Crescent"
	}
}

