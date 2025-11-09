// Go Swiss Ephemeris - Ephemeris Tests
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
	"math"
	"os"
	"path/filepath"
	"testing"
)

// getEphePath returns the path to ephemeris files for testing
func getEphePath() string {
	// Try to find the ephemeris directory
	paths := []string{
		"../swisseph/ephe",
		"../../swisseph/ephe",
		"./swisseph/ephe",
		"./swisseph_/ephe",
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			absPath, _ := filepath.Abs(path)
			return absPath
		}
	}

	return ""
}

// skipIfNoEphe skips the test if ephemeris files are not available
func skipIfNoEphe(t *testing.T) string {
	ephePath := getEphePath()
	if ephePath == "" {
		t.Skip("Ephemeris files not found, skipping test")
	}
	return ephePath
}

func TestWithEphemeris_SunPosition(t *testing.T) {
	ephePath := skipIfNoEphe(t)
	SetEphePath(ephePath)
	defer Close()

	// Test Sun position on January 1, 2000, 12:00 UTC
	jd := Julday(2000, 1, 1, 12.0, GregCal)
	result := CalcUT(jd, Sun, FlagSwieph|FlagSpeed)

	if result.Flag < 0 {
		t.Fatalf("CalcUT failed: %s", result.Error)
	}

	// Sun should be around 280° (Capricorn) on Jan 1, 2000
	expectedLon := 280.0
	tolerance := 2.0

	if math.Abs(result.Data[0]-expectedLon) > tolerance {
		t.Errorf("Sun longitude unexpected: got %.6f°, expected ~%.0f°", result.Data[0], expectedLon)
	}

	// Check that we have speed data
	if result.Data[3] == 0 {
		t.Error("Expected non-zero speed for Sun")
	}

	t.Logf("Sun on 2000-01-01: Lon=%.6f°, Lat=%.6f°, Dist=%.6f AU, Speed=%.6f°/day",
		result.Data[0], result.Data[1], result.Data[2], result.Data[3])
}

func TestWithEphemeris_MoonPosition(t *testing.T) {
	ephePath := skipIfNoEphe(t)
	SetEphePath(ephePath)
	defer Close()

	// Test Moon position on January 1, 2000, 12:00 UTC
	jd := Julday(2000, 1, 1, 12.0, GregCal)
	result := CalcUT(jd, Moon, FlagSwieph|FlagSpeed)

	if result.Flag < 0 {
		t.Fatalf("CalcUT failed: %s", result.Error)
	}

	// Moon moves fast, should have significant speed
	if math.Abs(result.Data[3]) < 10.0 {
		t.Errorf("Moon speed seems too low: %.6f°/day", result.Data[3])
	}

	// Moon distance should be around 1 AU / 400 (roughly)
	if result.Data[2] < 0.0020 || result.Data[2] > 0.0030 {
		t.Errorf("Moon distance unexpected: %.6f AU", result.Data[2])
	}

	t.Logf("Moon on 2000-01-01: Lon=%.6f°, Lat=%.6f°, Dist=%.6f AU, Speed=%.6f°/day",
		result.Data[0], result.Data[1], result.Data[2], result.Data[3])
}

func TestWithEphemeris_AllPlanets(t *testing.T) {
	ephePath := skipIfNoEphe(t)
	SetEphePath(ephePath)
	defer Close()

	jd := Julday(2024, 1, 1, 12.0, GregCal)

	planets := []struct {
		id   int32
		name string
	}{
		{Sun, "Sun"},
		{Moon, "Moon"},
		{Mercury, "Mercury"},
		{Venus, "Venus"},
		{Mars, "Mars"},
		{Jupiter, "Jupiter"},
		{Saturn, "Saturn"},
		{Uranus, "Uranus"},
		{Neptune, "Neptune"},
		{Pluto, "Pluto"},
	}

	for _, planet := range planets {
		result := CalcUT(jd, planet.id, FlagSwieph)

		if result.Flag < 0 {
			t.Errorf("%s calculation failed: %s", planet.name, result.Error)
			continue
		}

		// Check longitude is in valid range
		if result.Data[0] < 0 || result.Data[0] >= 360 {
			t.Errorf("%s longitude out of range: %.6f°", planet.name, result.Data[0])
		}

		// Check latitude is reasonable (except Sun which should be near 0)
		if planet.id == Sun && math.Abs(result.Data[1]) > 0.01 {
			t.Errorf("Sun latitude should be near 0: %.6f°", result.Data[1])
		}

		t.Logf("%s: Lon=%.6f°, Lat=%.6f°, Dist=%.6f AU",
			planet.name, result.Data[0], result.Data[1], result.Data[2])
	}
}

func TestWithEphemeris_HistoricalDate(t *testing.T) {
	ephePath := skipIfNoEphe(t)
	SetEphePath(ephePath)
	defer Close()

	// Test a historical date: July 20, 1969 (Moon landing)
	jd := Julday(1969, 7, 20, 20.0, GregCal)
	result := CalcUT(jd, Sun, FlagSwieph)

	if result.Flag < 0 {
		t.Fatalf("Historical date calculation failed: %s", result.Error)
	}

	// Sun should be in Cancer/Leo (around 120-150°) in July
	if result.Data[0] < 90 || result.Data[0] > 150 {
		t.Errorf("Sun position unexpected for July 1969: %.6f°", result.Data[0])
	}

	t.Logf("Sun on 1969-07-20: %.6f°", result.Data[0])
}

func TestWithEphemeris_FutureDate(t *testing.T) {
	ephePath := skipIfNoEphe(t)
	SetEphePath(ephePath)
	defer Close()

	// Test a future date: January 1, 2100
	jd := Julday(2100, 1, 1, 12.0, GregCal)
	result := CalcUT(jd, Mars, FlagSwieph)

	if result.Flag < 0 {
		t.Fatalf("Future date calculation failed: %s", result.Error)
	}

	// Just verify we get a valid result
	if result.Data[0] < 0 || result.Data[0] >= 360 {
		t.Errorf("Mars longitude out of range: %.6f°", result.Data[0])
	}

	t.Logf("Mars on 2100-01-01: %.6f°", result.Data[0])
}

func TestWithEphemeris_BCDate(t *testing.T) {
	ephePath := skipIfNoEphe(t)
	SetEphePath(ephePath)
	defer Close()

	// Test BC date: 100 BC (year -99 in astronomical year numbering)
	jd := Julday(-99, 1, 1, 12.0, GregCal)
	result := CalcUT(jd, Sun, FlagSwieph)

	if result.Flag < 0 {
		t.Logf("BC date calculation: %s (may be expected if BC files not present)", result.Error)
		return
	}

	// Just verify we get a valid result
	if result.Data[0] < 0 || result.Data[0] >= 360 {
		t.Errorf("Sun longitude out of range for BC date: %.6f°", result.Data[0])
	}

	t.Logf("Sun on 100 BC: %.6f°", result.Data[0])
}

func TestWithEphemeris_SpeedCalculation(t *testing.T) {
	ephePath := skipIfNoEphe(t)
	SetEphePath(ephePath)
	defer Close()

	jd := Julday(2024, 1, 1, 12.0, GregCal)

	// Test with speed flag
	resultWithSpeed := CalcUT(jd, Venus, FlagSwieph|FlagSpeed)
	if resultWithSpeed.Flag < 0 {
		t.Fatalf("Calculation with speed failed: %s", resultWithSpeed.Error)
	}

	// Test without speed flag
	resultNoSpeed := CalcUT(jd, Venus, FlagSwieph)
	if resultNoSpeed.Flag < 0 {
		t.Fatalf("Calculation without speed failed: %s", resultNoSpeed.Error)
	}

	// With speed flag, speed values should be non-zero
	if resultWithSpeed.Data[3] == 0 {
		t.Error("Expected non-zero speed with FlagSpeed")
	}

	// Position should be the same regardless of speed flag
	if math.Abs(resultWithSpeed.Data[0]-resultNoSpeed.Data[0]) > 0.0001 {
		t.Errorf("Position differs with/without speed flag: %.6f vs %.6f",
			resultWithSpeed.Data[0], resultNoSpeed.Data[0])
	}

	t.Logf("Venus speed: %.6f°/day", resultWithSpeed.Data[3])
}

func TestWithEphemeris_Retrograde(t *testing.T) {
	ephePath := skipIfNoEphe(t)
	SetEphePath(ephePath)
	defer Close()

	// Find a date when Mercury is retrograde (approximate)
	// Mercury retrogrades roughly 3 times a year
	jd := Julday(2024, 4, 10, 12.0, GregCal)

	result := CalcUT(jd, Mercury, FlagSwieph|FlagSpeed)
	if result.Flag < 0 {
		t.Fatalf("Calculation failed: %s", result.Error)
	}

	// Check if retrograde (negative speed)
	if result.Data[3] < 0 {
		t.Logf("Mercury is retrograde on 2024-04-10: speed=%.6f°/day", result.Data[3])
	} else {
		t.Logf("Mercury is direct on 2024-04-10: speed=%.6f°/day", result.Data[3])
	}
}

func TestWithEphemeris_Precision(t *testing.T) {
	ephePath := skipIfNoEphe(t)
	SetEphePath(ephePath)
	defer Close()

	jd := Julday(2000, 1, 1, 12.0, GregCal)

	// Calculate same position twice
	result1 := CalcUT(jd, Jupiter, FlagSwieph)
	result2 := CalcUT(jd, Jupiter, FlagSwieph)

	if result1.Flag < 0 || result2.Flag < 0 {
		t.Fatal("Calculation failed")
	}

	// Results should be identical
	for i := 0; i < 6; i++ {
		if result1.Data[i] != result2.Data[i] {
			t.Errorf("Results differ on repeated calculation: index %d: %.15f vs %.15f",
				i, result1.Data[i], result2.Data[i])
		}
	}

	t.Logf("Precision test passed: Jupiter position consistent")
}

func TestWithEphemeris_MultipleFiles(t *testing.T) {
	ephePath := skipIfNoEphe(t)
	SetEphePath(ephePath)
	defer Close()

	// Test dates that span multiple ephemeris files
	dates := []struct {
		year  int32
		month int32
		day   int32
	}{
		{1800, 1, 1},   // sepl_18.se1
		{1950, 6, 15},  // sepl_18.se1
		{2000, 1, 1},   // sepl_18.se1
		{2100, 12, 31}, // sepl_18.se1
		{2300, 7, 4},   // sepl_24.se1
	}

	for _, date := range dates {
		jd := Julday(date.year, date.month, date.day, 12.0, GregCal)
		result := CalcUT(jd, Saturn, FlagSwieph)

		if result.Flag < 0 {
			t.Logf("Date %d-%02d-%02d: %s (file may not be present)",
				date.year, date.month, date.day, result.Error)
			continue
		}

		if result.Data[0] < 0 || result.Data[0] >= 360 {
			t.Errorf("Saturn longitude out of range for %d-%02d-%02d: %.6f°",
				date.year, date.month, date.day, result.Data[0])
		}

		t.Logf("Saturn on %d-%02d-%02d: %.6f°",
			date.year, date.month, date.day, result.Data[0])
	}
}

func TestWithEphemeris_GetCurrentFileData(t *testing.T) {
	ephePath := skipIfNoEphe(t)
	SetEphePath(ephePath)
	defer Close()

	// Calculate something to load a file
	jd := Julday(2000, 1, 1, 12.0, GregCal)
	result := CalcUT(jd, Sun, FlagSwieph)

	if result.Flag < 0 {
		t.Fatalf("Calculation failed: %s", result.Error)
	}

	// Get file data for the Sun
	fileData := GetCurrentFileData(0)

	if fileData.Path != "" {
		t.Logf("Ephemeris file: %s", fileData.Path)
		t.Logf("Coverage: JD %.2f to %.2f", fileData.StartDate, fileData.EndDate)
		t.Logf("DE number: %d", fileData.Denum)

		// Verify the file covers our test date
		if jd < fileData.StartDate || jd > fileData.EndDate {
			t.Errorf("File doesn't cover test date: JD %.2f not in [%.2f, %.2f]",
				jd, fileData.StartDate, fileData.EndDate)
		}
	}
}

func TestWithEphemeris_CompareWithMoshier(t *testing.T) {
	ephePath := skipIfNoEphe(t)

	jd := Julday(2000, 1, 1, 12.0, GregCal)

	// Calculate with Swiss Ephemeris files
	SetEphePath(ephePath)
	resultSwiss := CalcUT(jd, Mars, FlagSwieph)
	Close()

	if resultSwiss.Flag < 0 {
		t.Fatalf("Swiss Ephemeris calculation failed: %s", resultSwiss.Error)
	}

	// Calculate with Moshier (no ephemeris path)
	SetEphePath("")
	resultMoshier := CalcUT(jd, Mars, FlagMoseph)
	Close()

	if resultMoshier.Flag < 0 {
		t.Fatalf("Moshier calculation failed: %s", resultMoshier.Error)
	}

	// Results should be close but not identical
	diff := math.Abs(resultSwiss.Data[0] - resultMoshier.Data[0])

	t.Logf("Mars longitude on 2000-01-01:")
	t.Logf("  Swiss Ephemeris: %.6f°", resultSwiss.Data[0])
	t.Logf("  Moshier:         %.6f°", resultMoshier.Data[0])
	t.Logf("  Difference:      %.6f°", diff)

	// Difference should be small (within a few arc-minutes)
	if diff > 0.1 {
		t.Logf("Warning: Large difference between Swiss Ephemeris and Moshier: %.6f°", diff)
	}
}

func TestWithEphemeris_AllPlanetsSanJose1994(t *testing.T) {
	ephePath := skipIfNoEphe(t)
	SetEphePath(ephePath)
	defer Close()

	// San Jose, CA, USA coordinates
	// Coordinates: 121W53'00, 37N20'00
	// Convert to decimal: 121° 53' 0" W = -121.883333°, 37° 20' 0" N = 37.333333°
	sanJoseLat := 37.0 + 20.0/60.0     // 37° 20' = 37.333333°
	sanJoseLon := -(121.0 + 53.0/60.0) // 121° 53' W = -121.883333°
	sanJoseAlt := 0.0

	// Set topocentric location for San Jose
	SetTopo(sanJoseLon, sanJoseLat, sanJoseAlt)

	// Set Lahiri ayanamsa for sidereal calculations
	SetSidMode(SidmLahiri, 0, 0)

	// Date: Wed Aug 17, 1994, 08:51:00
	// Time zone: 08:00:00 with DST 01:00:00
	// TZ 08:00 means UTC-8:00 (Pacific Standard Time)
	// DST 01:00 means add 1 hour (Pacific Daylight Time)
	// So effective offset is UTC-7:00 (PDT)
	// Local time 08:51 PDT = UTC 15:51
	year := int32(1994)
	month := int32(8)
	day := int32(17)
	hour := 8.0 + 51.0/60.0 + 0.0/3600.0 // 08:51:00 = 8.85 hours

	// Convert from PDT (UTC-7:00) to UTC
	// PDT 08:51 + 7:00 = UTC 15:51
	pdtOffset := -7.0           // PDT is UTC-7:00 (7 hours behind)
	hourUTC := hour - pdtOffset // Subtract negative = add

	// If time goes over 24, adjust the day
	if hourUTC >= 24 {
		hourUTC -= 24
		day += 1
	}

	// Calculate Julian day using UTC time
	jdUTC := Julday(year, month, day, hourUTC, GregCal)

	// Also calculate with local time for comparison
	jdLocal := Julday(year, month, day, hour, GregCal)

	t.Logf("Test Date: Wednesday, August 17, 1994, 08:51:00 PDT (UTC-7:00)")
	t.Logf("Converted to UTC: %04d-%02d-%02d %02.0f:%02.0f:00",
		year, month, day, hourUTC, (hourUTC-float64(int(hourUTC)))*60)
	t.Logf("Location: San Jose, CA, USA (%.4f°N, %.4f°W)", sanJoseLat, -sanJoseLon)
	t.Logf("Julian Day (UTC): %.6f", jdUTC)
	t.Logf("Julian Day (Local): %.6f", jdLocal)

	// Use UTC time for calculations (standard practice)
	jd := jdUTC

	// Get the ayanamsa value
	ayanamsa := GetAyanamsaUT(jd)

	signs := []string{"Aries", "Taurus", "Gemini", "Cancer", "Leo", "Virgo",
		"Libra", "Scorpio", "Sagittarius", "Capricorn", "Aquarius", "Pisces"}

	// Calculate all planets
	planets := []struct {
		id   int32
		name string
	}{
		{Sun, "Sun"},
		{Moon, "Moon"},
		{Mercury, "Mercury"},
		{Venus, "Venus"},
		{Mars, "Mars"},
		{Jupiter, "Jupiter"},
		{Saturn, "Saturn"},
		{Uranus, "Uranus"},
		{Neptune, "Neptune"},
		{Pluto, "Pluto"},
		{TrueNode, "Rahu (True Node)"},
		{MeanNode, "Mean Node"},
	}

	t.Logf("\n=== TROPICAL ZODIAC (WESTERN) ===")
	for _, planet := range planets {
		result := CalcUT(jd, planet.id, FlagSwieph|FlagSpeed)
		if result.Flag < 0 {
			t.Logf("%s: Error - %s", planet.name, result.Error)
			continue
		}

		sign := int(result.Data[0] / 30)
		degree := result.Data[0] - float64(sign)*30
		split := SplitDeg(degree, 0)

		retrograde := ""
		if result.Data[3] < 0 {
			retrograde = " (R)"
		}

		t.Logf("%-18s: %3d° %2d' %2d\" %-12s (%.6f°)%s",
			planet.name,
			split.Degree, split.Minute, split.Second,
			signs[sign],
			result.Data[0],
			retrograde)
	}

	t.Logf("\n=== SIDEREAL ZODIAC (LAHIRI - VEDIC) ===")
	t.Logf("Ayanamsa (Lahiri): %.6f°", ayanamsa)

	for _, planet := range planets {
		result := CalcUT(jd, planet.id, FlagSwieph|FlagSpeed|FlagSidereal)
		if result.Flag < 0 {
			t.Logf("%s: Error - %s", planet.name, result.Error)
			continue
		}

		sign := int(result.Data[0] / 30)
		degree := result.Data[0] - float64(sign)*30
		split := SplitDeg(degree, 0)

		retrograde := ""
		if result.Data[3] < 0 {
			retrograde = " (R)"
		}

		t.Logf("%-18s: %3d° %2d' %2d\" %-12s (%.6f°)%s",
			planet.name,
			split.Degree, split.Minute, split.Second,
			signs[sign],
			result.Data[0],
			retrograde)
	}

	// Calculate house positions for San Jose
	// For Vedic astrology, use Whole Sign houses ('W')
	// But also calculate Placidus for comparison
	housesWhole := Houses(jd, sanJoseLat, sanJoseLon, 'W')    // Whole Sign (Vedic)
	housesPlacidus := Houses(jd, sanJoseLat, sanJoseLon, 'P') // Placidus (Western)

	if housesWhole.Flag >= 0 {
		// For Vedic astrology, subtract ayanamsa from tropical Ascendant to get sidereal Ascendant
		tropicalAsc := housesWhole.Points[Asc]
		siderealAsc := Degnorm(tropicalAsc - ayanamsa)

		ascSign := int(siderealAsc / 30)
		ascDegree := siderealAsc - float64(ascSign)*30
		ascSplit := SplitDeg(ascDegree, 0)

		tropicalMC := housesWhole.Points[MC]
		siderealMC := Degnorm(tropicalMC - ayanamsa)

		mcSign := int(siderealMC / 30)
		mcDegree := siderealMC - float64(mcSign)*30
		mcSplit := SplitDeg(mcDegree, 0)

		t.Logf("\n=== HOUSE POSITIONS (WHOLE SIGN - VEDIC/SIDEREAL) ===")
		t.Logf("Ascendant (Tropical): %.6f° = %s", tropicalAsc, signs[int(tropicalAsc/30)])
		t.Logf("Ascendant (Sidereal): %.6f° = %d° %d' %d\" %s",
			siderealAsc,
			ascSplit.Degree, ascSplit.Minute, ascSplit.Second,
			signs[ascSign])
		t.Logf("MC (Sidereal): %.6f° = %d° %d' %d\" %s",
			siderealMC,
			mcSplit.Degree, mcSplit.Minute, mcSplit.Second,
			signs[mcSign])

		t.Logf("\n=== PLANET HOUSE PLACEMENTS (WHOLE SIGN - VEDIC) ===")

		// Calculate house placement for each planet
		for _, planet := range planets {
			result := CalcUT(jd, planet.id, FlagSwieph|FlagSidereal)
			if result.Flag < 0 {
				continue
			}

			planetSign := int(result.Data[0] / 30)
			planetHouse := ((planetSign - ascSign + 12) % 12) + 1

			t.Logf("%-18s: House %2d (%s)", planet.name, planetHouse, signs[planetSign])
		}
	}

	// Also show Placidus for comparison
	if housesPlacidus.Flag >= 0 {
		ascSignPlac := int(housesPlacidus.Points[Asc] / 30)
		ascDegreePlac := housesPlacidus.Points[Asc] - float64(ascSignPlac)*30
		ascSplitPlac := SplitDeg(ascDegreePlac, 0)

		t.Logf("\n=== HOUSE POSITIONS (PLACIDUS - WESTERN/TROPICAL) ===")
		t.Logf("Ascendant: %.6f° = %d° %d' %d\" %s",
			housesPlacidus.Points[Asc],
			ascSplitPlac.Degree, ascSplitPlac.Minute, ascSplitPlac.Second,
			signs[ascSignPlac])

		t.Logf("\n=== PLANET HOUSE PLACEMENTS (PLACIDUS - WESTERN) ===")

		// Calculate house placement for each planet using Placidus
		for _, planet := range planets {
			result := CalcUT(jd, planet.id, FlagSwieph)
			if result.Flag < 0 {
				continue
			}

			planetLon := result.Data[0]
			planetHouse := 1

			// Find which Placidus house the planet is in
			for i := 0; i < len(housesPlacidus.Houses); i++ {
				nextHouse := housesPlacidus.Houses[(i+1)%len(housesPlacidus.Houses)]
				currentHouse := housesPlacidus.Houses[i]

				// Handle wrap-around at 360°
				if nextHouse < currentHouse {
					if planetLon >= currentHouse || planetLon < nextHouse {
						planetHouse = i + 1
						break
					}
				} else {
					if planetLon >= currentHouse && planetLon < nextHouse {
						planetHouse = i + 1
						break
					}
				}
			}

			planetSign := int(planetLon / 30)
			t.Logf("%-18s: House %2d (%s)", planet.name, planetHouse, signs[planetSign])
		}
	}
}

func BenchmarkWithEphemeris_CalcUT(b *testing.B) {
	ephePath := getEphePath()
	if ephePath == "" {
		b.Skip("Ephemeris files not found")
	}

	SetEphePath(ephePath)
	defer Close()

	jd := Julday(2024, 1, 1, 12.0, GregCal)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CalcUT(jd, Sun, FlagSwieph)
	}
}

func BenchmarkWithEphemeris_AllPlanets(b *testing.B) {
	ephePath := getEphePath()
	if ephePath == "" {
		b.Skip("Ephemeris files not found")
	}

	SetEphePath(ephePath)
	defer Close()

	jd := Julday(2024, 1, 1, 12.0, GregCal)
	planets := []int32{Sun, Moon, Mercury, Venus, Mars, Jupiter, Saturn, Uranus, Neptune, Pluto}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, planet := range planets {
			CalcUT(jd, planet, FlagSwieph)
		}
	}
}
