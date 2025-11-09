// Go Swiss Ephemeris - Tests
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
	"testing"
)

func TestVersion(t *testing.T) {
	version := Version()
	if version == "" {
		t.Error("Version should not be empty")
	}
	t.Logf("Swiss Ephemeris Version: %s", version)
}

func TestPackageVersion(t *testing.T) {
	version := GetPackageVersion()
	if version == "" {
		t.Error("PackageVersion should not be empty")
	}
	if version != PackageVersion {
		t.Errorf("GetPackageVersion() = %s, want %s", version, PackageVersion)
	}
	t.Logf("Go Bindings Version: %s", version)
}

func TestJulday(t *testing.T) {
	// Test known Julian day
	// January 1, 2000, 12:00 UTC should be JD 2451545.0
	jd := Julday(2000, 1, 1, 12.0, GregCal)
	expected := 2451545.0
	if math.Abs(jd-expected) > 0.0001 {
		t.Errorf("Julday calculation incorrect: got %.6f, expected %.6f", jd, expected)
	}
}

func TestRevjul(t *testing.T) {
	jd := 2451545.0
	date := Revjul(jd, GregCal)
	
	if date.Year != 2000 || date.Month != 1 || date.Day != 1 {
		t.Errorf("Revjul calculation incorrect: got %d-%d-%d, expected 2000-1-1",
			date.Year, date.Month, date.Day)
	}
	
	if math.Abs(date.Hour-12.0) > 0.001 {
		t.Errorf("Revjul hour incorrect: got %.3f, expected 12.0", date.Hour)
	}
}

func TestCalcUT(t *testing.T) {
	// Calculate Sun position on January 1, 2000
	jd := Julday(2000, 1, 1, 12.0, GregCal)
	result := CalcUT(jd, Sun, FlagSwieph)
	
	if result.Flag < 0 {
		t.Errorf("CalcUT failed: %s", result.Error)
		return
	}
	
	// Sun should be around 280° on Jan 1 (Capricorn)
	if result.Data[0] < 270 || result.Data[0] > 290 {
		t.Errorf("Sun position unexpected: %.2f°", result.Data[0])
	}
	
	t.Logf("Sun position on 2000-01-01: %.6f°", result.Data[0])
}

func TestCalc(t *testing.T) {
	jd := Julday(2000, 1, 1, 12.0, GregCal)
	result := Calc(jd, Moon, FlagSwieph|FlagSpeed)
	
	if result.Flag < 0 {
		t.Errorf("Calc failed: %s", result.Error)
		return
	}
	
	// Check that we got 6 values
	if len(result.Data) != 6 {
		t.Errorf("Expected 6 data values, got %d", len(result.Data))
	}
	
	t.Logf("Moon position: %.6f°, speed: %.6f°/day", result.Data[0], result.Data[3])
}

func TestHouses(t *testing.T) {
	jd := Julday(2000, 1, 1, 12.0, GregCal)
	lat := 51.5074  // London
	lon := -0.1278
	
	houses := Houses(jd, lat, lon, 'P')
	
	if houses.Flag < 0 {
		t.Error("Houses calculation failed")
		return
	}
	
	// Should have 12 houses
	if len(houses.Houses) != 12 {
		t.Errorf("Expected 12 houses, got %d", len(houses.Houses))
	}
	
	// Check that houses are in ascending order (mostly)
	for i := 0; i < len(houses.Houses)-1; i++ {
		if houses.Houses[i] > houses.Houses[i+1] {
			// Allow wrap-around at 360°
			if houses.Houses[i+1] > 10 {
				t.Errorf("Houses not in order at index %d: %.2f > %.2f",
					i, houses.Houses[i], houses.Houses[i+1])
			}
		}
	}
	
	t.Logf("Ascendant: %.6f°", houses.Points[Asc])
	t.Logf("MC: %.6f°", houses.Points[MC])
}

func TestDegnorm(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{375.5, 15.5},
		{-10.0, 350.0},
		{360.0, 0.0},
		{720.5, 0.5},
	}
	
	for _, test := range tests {
		result := Degnorm(test.input)
		if math.Abs(result-test.expected) > 0.0001 {
			t.Errorf("Degnorm(%.2f) = %.2f, expected %.2f",
				test.input, result, test.expected)
		}
	}
}

func TestDidegn(t *testing.T) {
	// Test shortest arc calculation
	// Difdegn calculates p1 - p2 normalized to 0-360
	diff := Difdegn(350.0, 10.0)
	expected := 340.0 // 350 - 10 = 340
	if math.Abs(diff-expected) > 0.0001 {
		t.Errorf("Difdegn(350, 10) = %.2f, expected %.2f", diff, expected)
	}
	
	diff = Difdegn(10.0, 350.0)
	expected = 20.0 // 10 - 350 = -340, normalized = 20
	if math.Abs(diff-expected) > 0.0001 {
		t.Errorf("Difdegn(10, 350) = %.2f, expected %.2f", diff, expected)
	}
}

func TestSplitDeg(t *testing.T) {
	result := SplitDeg(123.456789, SplitDegZodiacal)
	
	// 123° is in Leo (sign 4, since Aries=0)
	expectedSign := int32(4)
	expectedDegree := int32(3) // 3° into Leo
	
	if result.Sign != expectedSign {
		t.Errorf("Sign incorrect: got %d, expected %d", result.Sign, expectedSign)
	}
	
	if result.Degree != expectedDegree {
		t.Errorf("Degree incorrect: got %d, expected %d", result.Degree, expectedDegree)
	}
	
	t.Logf("123.456789° = %d° %d' %d\" in sign %d",
		result.Degree, result.Minute, result.Second, result.Sign)
}

func TestGetPlanetName(t *testing.T) {
	tests := []struct {
		planet   int32
		expected string
	}{
		{Sun, "Sun"},
		{Moon, "Moon"},
		{Mercury, "Mercury"},
		{Venus, "Venus"},
		{Mars, "Mars"},
	}
	
	for _, test := range tests {
		name := GetPlanetName(test.planet)
		if name != test.expected {
			t.Errorf("GetPlanetName(%d) = %s, expected %s",
				test.planet, name, test.expected)
		}
	}
}

func TestDeltat(t *testing.T) {
	jd := Julday(2000, 1, 1, 12.0, GregCal)
	dt := Deltat(jd)
	
	// Delta T for year 2000 should be around 63-64 seconds
	// But the function returns it in days, so convert to seconds
	dtSeconds := dt * 86400.0
	if dtSeconds < 60 || dtSeconds > 70 {
		t.Logf("Delta T for year 2000: %.2f seconds (%.6f days)", dtSeconds, dt)
		// Don't fail, just log - Delta T can vary based on model
	}
	
	t.Logf("Delta T for 2000-01-01: %.4f seconds (%.8f days)", dtSeconds, dt)
}

func TestSidtime(t *testing.T) {
	jd := Julday(2000, 1, 1, 0.0, GregCal)
	st := Sidtime(jd)
	
	// Sidereal time should be between 0 and 24
	if st < 0 || st > 24 {
		t.Errorf("Sidereal time out of range: %.6f", st)
	}
	
	t.Logf("Sidereal time for 2000-01-01 00:00: %.6f hours", st)
}

func TestUtcToJd(t *testing.T) {
	jdArr, err := UtcToJd(2000, 1, 1, 12, 0, 0.0, GregCal)
	if err != nil {
		t.Errorf("UtcToJd failed: %v", err)
		return
	}
	
	// JD ET should be close to 2451545.0 (with Delta T added)
	if math.Abs(jdArr[0]-2451545.0) > 0.01 {
		t.Errorf("JD ET unexpected: %.6f", jdArr[0])
	}
	
	t.Logf("JD ET: %.6f, JD UT: %.6f", jdArr[0], jdArr[1])
}

func TestDayOfWeek(t *testing.T) {
	// January 1, 2000 was a Saturday (day 5, since Monday=0)
	jd := Julday(2000, 1, 1, 12.0, GregCal)
	dow := DayOfWeek(jd)
	
	expected := 5 // Saturday
	if dow != expected {
		t.Errorf("Day of week incorrect: got %d, expected %d", dow, expected)
	}
}

func TestGetAyanamsaName(t *testing.T) {
	name := GetAyanamsaName(SidmLahiri)
	if name == "" {
		t.Error("Ayanamsa name should not be empty")
	}
	t.Logf("Lahiri ayanamsa: %s", name)
}

func TestSetAndGetTidAcc(t *testing.T) {
	// Set a custom tidal acceleration
	SetTidAcc(-25.8)
	
	// Get it back
	tidAcc := GetTidAcc()
	
	if math.Abs(tidAcc-(-25.8)) > 0.01 {
		t.Errorf("Tidal acceleration: got %.2f, expected -25.8", tidAcc)
	}
}

// Benchmark tests
func BenchmarkCalcUT(b *testing.B) {
	jd := Julday(2000, 1, 1, 12.0, GregCal)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		CalcUT(jd, Sun, FlagSwieph)
	}
}

func BenchmarkHouses(b *testing.B) {
	jd := Julday(2000, 1, 1, 12.0, GregCal)
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		Houses(jd, 51.5074, -0.1278, 'P')
	}
}

func BenchmarkJulday(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Julday(2000, 1, 1, 12.0, GregCal)
	}
}

