// Go Swiss Ephemeris - Basic Example
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
	fmt.Println("Swiss Ephemeris - Basic Example")
	fmt.Println("================================\n")

	// Optional: Set ephemeris file path
	// swisseph.SetEphePath("../swisseph/ephe")

	// Get library versions
	fmt.Printf("Go Bindings Version: %s\n", swisseph.GetPackageVersion())
	fmt.Printf("Swiss Ephemeris Version: %s\n\n", swisseph.Version())

	// Use current date/time
	now := time.Now().UTC()
	jd := swisseph.Julday(
		int32(now.Year()),
		int32(now.Month()),
		int32(now.Day()),
		float64(now.Hour())+float64(now.Minute())/60.0+float64(now.Second())/3600.0,
		swisseph.GregCal,
	)

	fmt.Printf("Current Date: %s\n", now.Format("2006-01-02 15:04:05 UTC"))
	fmt.Printf("Julian Day: %.6f\n\n", jd)

	// Calculate positions of all major planets
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
	}

	fmt.Println("Planetary Positions (Tropical Zodiac)")
	fmt.Println("--------------------------------------")

	for _, planet := range planets {
		result := swisseph.CalcUT(jd, planet.id, swisseph.FlagSwieph|swisseph.FlagSpeed)
		if result.Flag >= 0 {
			longitude := result.Data[0]
			latitude := result.Data[1]
			distance := result.Data[2]
			speedLon := result.Data[3]

			// Convert to zodiac sign
			sign := int(longitude / 30)
			degree := longitude - float64(sign)*30
			signs := []string{"Aries", "Taurus", "Gemini", "Cancer", "Leo", "Virgo",
				"Libra", "Scorpio", "Sagittarius", "Capricorn", "Aquarius", "Pisces"}

			fmt.Printf("%-10s: %6.2f° (%2d° %s) | Lat: %6.2f° | Dist: %8.4f AU | Speed: %+7.4f°/day\n",
				planet.name, longitude, int(degree), signs[sign], latitude, distance, speedLon)
		} else {
			fmt.Printf("%-10s: Error - %s\n", planet.name, result.Error)
		}
	}

	// Calculate house cusps for a location (example: New York City)
	fmt.Println("\nHouse Cusps (Placidus System)")
	fmt.Println("------------------------------")
	fmt.Println("Location: New York City (40.7128°N, 74.0060°W)")

	lat := 40.7128
	lon := -74.0060
	houses := swisseph.Houses(jd, lat, lon, 'P')

	if houses.Flag >= 0 {
		for i, cusp := range houses.Houses {
			sign := int(cusp / 30)
			degree := cusp - float64(sign)*30
			signs := []string{"Aries", "Taurus", "Gemini", "Cancer", "Leo", "Virgo",
				"Libra", "Scorpio", "Sagittarius", "Capricorn", "Aquarius", "Pisces"}
			fmt.Printf("House %2d: %6.2f° (%2d° %s)\n", i+1, cusp, int(degree), signs[sign])
		}

		fmt.Printf("\nAscendant: %6.2f°\n", houses.Points[swisseph.Asc])
		fmt.Printf("MC:        %6.2f°\n", houses.Points[swisseph.MC])
		fmt.Printf("ARMC:      %6.2f°\n", houses.Points[swisseph.ARMC])
		fmt.Printf("Vertex:    %6.2f°\n", houses.Points[swisseph.Vertex])
	}

	// Calculate Delta T
	fmt.Printf("\nDelta T (TT - UT): %.4f seconds\n", swisseph.Deltat(jd))

	// Calculate sidereal time
	fmt.Printf("Sidereal Time: %.6f hours\n", swisseph.Sidtime(jd))

	// Clean up
	swisseph.Close()
	fmt.Println("\nDone!")
}
