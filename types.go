// Go Swiss Ephemeris - Types
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

// CalcResult represents the result of a planetary calculation
type CalcResult struct {
	Flag  int32     // Return flag
	Error string    // Error message if any
	Data  []float64 // Calculation results (longitude, latitude, distance, speed, etc.)
}

// HousesResult represents the result of house calculations
type HousesResult struct {
	Flag   int32     // Return flag
	Houses []float64 // House cusps
	Points []float64 // Ascendant, MC, ARMC, Vertex, etc.
}

// JulianDay represents a Julian day number
type JulianDay float64

// DateResult represents a calendar date
type DateResult struct {
	Year  int
	Month int
	Day   int
	Hour  float64
}

// UTCResult represents UTC date and time
type UTCResult struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second float64
}

// EclipseResult represents eclipse calculation results
type EclipseResult struct {
	Flag     int32     // Eclipse type flags
	Maximum  float64   // Time of maximum eclipse
	Begin    float64   // Begin time (if applicable)
	End      float64   // End time (if applicable)
	Totality float64   // Duration of totality (if applicable)
	Attr     []float64 // Eclipse attributes
	Error    string    // Error message if any
}

// EclipseWhereResult represents where an eclipse is visible
type EclipseWhereResult struct {
	Flag      int32     // Eclipse type flags
	Longitude float64   // Geographic longitude
	Latitude  float64   // Geographic latitude
	Attr      []float64 // Eclipse attributes
	Error     string    // Error message if any
}

// RiseTransResult represents rise/set/transit calculation results
type RiseTransResult struct {
	Flag  int32   // Return flag
	Time  float64 // Julian day of event
	Error string  // Error message if any
}

// FixstarResult represents fixed star calculation results
type FixstarResult struct {
	Flag     int32     // Return flag
	StarName string    // Actual star name used
	Data     []float64 // Position data
	Error    string    // Error message if any
}

// FixstarMagResult represents fixed star magnitude
type FixstarMagResult struct {
	Flag      int32   // Return flag
	StarName  string  // Actual star name used
	Magnitude float64 // Star magnitude
	Error     string  // Error message if any
}

// NodApsResult represents nodes and apsides calculation results
type NodApsResult struct {
	Flag        int32     // Return flag
	Ascending   []float64 // Ascending node data
	Descending  []float64 // Descending node data
	Perihelion  []float64 // Perihelion data
	Aphelion    []float64 // Aphelion data
	Error       string    // Error message if any
}

// OrbitalElementsResult represents orbital elements
type OrbitalElementsResult struct {
	Flag     int32     // Return flag
	Elements []float64 // Orbital elements
	Error    string    // Error message if any
}

// SplitDegResult represents split degree components
type SplitDegResult struct {
	Degree     int32   // Degree
	Minute     int32   // Minute
	Second     int32   // Second
	SecondFrac float64 // Fractional second
	Sign       int32   // Zodiac sign (if applicable)
}

// AzaltResult represents azimuth/altitude coordinates
type AzaltResult struct {
	Azimuth  float64 // Azimuth
	Altitude float64 // True altitude
	AppAlt   float64 // Apparent altitude (with refraction)
}

// HeliacalResult represents heliacal event calculation
type HeliacalResult struct {
	Flag  int32     // Return flag
	Time  []float64 // Event times
	Attr  []float64 // Event attributes
	Error string    // Error message if any
}

// FileData represents ephemeris file information
type FileData struct {
	Path      string  // File path
	StartDate float64 // Start date of file coverage
	EndDate   float64 // End date of file coverage
	Denum     int32   // DE number (for JPL files)
}

