# Go Swiss Ephemeris

A comprehensive Go binding for the Swiss Ephemeris library - the definitive astronomical calculation library for astrology and astronomy applications.

**Version:** v1.0.0

## Features

- **100% API Coverage** - All Swiss Ephemeris functions are available
- **Type-Safe** - Idiomatic Go interfaces with proper error handling
- **High Performance** - Direct CGO bindings to the C library
- **Well Documented** - Comprehensive documentation and examples
- **Version Matched** - Based on Swiss Ephemeris 2.10.03b

## Installation

### Prerequisites

- Go 1.21 or later
- A C compiler (gcc, clang, or MSVC)

### Install from Source

```bash
# Clone the repository
git clone https://github.com/tejzpr/go-swisseph.git
cd go-swisseph

# Now you can use the library
go test -v
```

### Install via go get

```bash
go get github.com/tejzpr/go-swisseph
```

### Important: Swiss Ephemeris Source Code

This library includes the Swiss Ephemeris C library source code directly in the `swisseph/` directory. The C source files are compiled automatically by CGO when you build the library. No additional setup is required.

The `swisseph` package is a separate Go package that exists solely to make CGO compile the C files in that directory. CGO automatically compiles all `.c` files in a package directory when that package has `import "C"`. The main `swisseph` package imports `swisseph` to ensure the C files are compiled and linked.

## Licensing

### Go Bindings License

This Go binding library (`github.com/tejzpr/go-swisseph`) is licensed under the **GNU Affero General Public License version 3 or later (AGPL-3.0-or-later)**.

The Go wrapper code, bindings, and this repository are licensed under AGPL-3.0-or-later, which ensures full compatibility with Swiss Ephemeris licensing requirements.

### Swiss Ephemeris Licensing Requirements

**IMPORTANT**: This library includes the Swiss Ephemeris C library source code directly in the `swisseph/` directory, licensed under AGPL-3.0-or-later.

**Users of this library MUST comply with Swiss Ephemeris licensing terms**. The Swiss Ephemeris source code included in this repository is available under two license options:

1. **GNU Affero General Public License (AGPL-3.0-or-later)**
   - For open source projects
   - When using the AGPL version, the entire project (Go bindings + Swiss Ephemeris) is AGPL-3.0-or-later
   - This is the recommended approach for open source projects

2. **Swiss Ephemeris Professional License**
   - Commercial license available from Astrodienst AG
   - Allows use in proprietary/commercial applications
   - Requires purchasing a license from Astrodienst AG

For more information about Swiss Ephemeris licensing, visit:
- https://www.astro.com/swisseph/
- Contact: Astrodienst AG, Switzerland

## Quick Start

```go
package main

import (
    "fmt"
    swisseph "github.com/tejzpr/go-swisseph"
)

func main() {
    // Set the path to ephemeris files (optional, uses default if not set)
    swisseph.SetEphePath("./ephe")
    
    // Calculate Julian day for a date
    jd := swisseph.Julday(2024, 1, 1, 12.0, swisseph.GregCal)
    fmt.Printf("Julian Day: %.6f\n", jd)
    
    // Calculate Sun position
    result := swisseph.CalcUT(jd, swisseph.Sun, swisseph.FlagSwieph|swisseph.FlagSpeed)
    if result.Flag >= 0 {
        fmt.Printf("Sun Longitude: %.6f°\n", result.Data[0])
        fmt.Printf("Sun Latitude: %.6f°\n", result.Data[1])
        fmt.Printf("Sun Distance: %.6f AU\n", result.Data[2])
    } else {
        fmt.Printf("Error: %s\n", result.Error)
    }
    
    // Clean up
    swisseph.Close()
}
```

## Core Functions

### Date and Time Conversions

```go
// Calculate Julian day from calendar date
jd := swisseph.Julday(2024, 1, 1, 12.0, swisseph.GregCal)

// Convert Julian day back to calendar date
date := swisseph.Revjul(jd, swisseph.GregCal)
fmt.Printf("Year: %d, Month: %d, Day: %d, Hour: %.2f\n", 
    date.Year, date.Month, date.Day, date.Hour)

// Convert UTC to Julian day
jdArr, err := swisseph.UtcToJd(2024, 1, 1, 12, 0, 0.0, swisseph.GregCal)
if err == nil {
    fmt.Printf("JD ET: %.6f, JD UT: %.6f\n", jdArr[0], jdArr[1])
}

// Calculate Delta T (difference between TT and UT)
dt := swisseph.Deltat(jd)
fmt.Printf("Delta T: %.6f seconds\n", dt)
```

### Planetary Calculations

```go
// Calculate planet positions (Universal Time)
result := swisseph.CalcUT(jd, swisseph.Moon, swisseph.FlagSwieph|swisseph.FlagSpeed)
if result.Flag >= 0 {
    fmt.Printf("Longitude: %.6f°\n", result.Data[0])
    fmt.Printf("Latitude: %.6f°\n", result.Data[1])
    fmt.Printf("Distance: %.6f AU\n", result.Data[2])
    fmt.Printf("Speed in Long: %.6f°/day\n", result.Data[3])
    fmt.Printf("Speed in Lat: %.6f°/day\n", result.Data[4])
    fmt.Printf("Speed in Dist: %.6f AU/day\n", result.Data[5])
}

// Available planets
planets := []int32{
    swisseph.Sun, swisseph.Moon, swisseph.Mercury, swisseph.Venus,
    swisseph.Mars, swisseph.Jupiter, swisseph.Saturn, swisseph.Uranus,
    swisseph.Neptune, swisseph.Pluto,
}

for _, planet := range planets {
    name := swisseph.GetPlanetName(planet)
    result := swisseph.CalcUT(jd, planet, swisseph.FlagSwieph)
    fmt.Printf("%s: %.6f°\n", name, result.Data[0])
}
```

### House Systems

```go
// Calculate houses for a location
lat := 51.5074  // London latitude
lon := -0.1278  // London longitude

houses := swisseph.Houses(jd, lat, lon, 'P') // Placidus system
if houses.Flag >= 0 {
    fmt.Println("House Cusps:")
    for i, cusp := range houses.Houses {
        fmt.Printf("House %d: %.6f°\n", i+1, cusp)
    }
    
    fmt.Printf("Ascendant: %.6f°\n", houses.Points[swisseph.Asc])
    fmt.Printf("MC: %.6f°\n", houses.Points[swisseph.MC])
    fmt.Printf("ARMC: %.6f°\n", houses.Points[swisseph.ARMC])
    fmt.Printf("Vertex: %.6f°\n", houses.Points[swisseph.Vertex])
}

// Available house systems
// 'P' = Placidus
// 'K' = Koch
// 'O' = Porphyrius
// 'R' = Regiomontanus
// 'C' = Campanus
// 'E' = Equal (from Asc)
// 'W' = Whole sign
// And many more...
```

### Fixed Stars

```go
// Calculate position of a fixed star
star := swisseph.FixstarUT("Aldebaran", jd, swisseph.FlagSwieph)
if star.Flag >= 0 {
    fmt.Printf("Star: %s\n", star.StarName)
    fmt.Printf("Longitude: %.6f°\n", star.Data[0])
    fmt.Printf("Latitude: %.6f°\n", star.Data[1])
}

// Get star magnitude
mag := swisseph.FixstarMag("Sirius")
if mag.Flag >= 0 {
    fmt.Printf("Magnitude of %s: %.2f\n", mag.StarName, mag.Magnitude)
}
```

### Eclipse Calculations

```go
// Find next solar eclipse globally
eclipse := swisseph.SolEclipseWhenGlob(jd, swisseph.FlagSwieph, 
    swisseph.EclAlltypesSolar, false)
if eclipse.Flag >= 0 {
    fmt.Printf("Next solar eclipse: JD %.6f\n", eclipse.Maximum)
    
    // Get eclipse type
    if eclipse.Flag&swisseph.EclTotal != 0 {
        fmt.Println("Type: Total")
    } else if eclipse.Flag&swisseph.EclAnnular != 0 {
        fmt.Println("Type: Annular")
    } else if eclipse.Flag&swisseph.EclPartial != 0 {
        fmt.Println("Type: Partial")
    }
}

// Find next lunar eclipse
lunEclipse := swisseph.LunEclipseWhen(jd, swisseph.FlagSwieph, 
    swisseph.EclAlltypesLunar, false)
if lunEclipse.Flag >= 0 {
    fmt.Printf("Next lunar eclipse: JD %.6f\n", lunEclipse.Maximum)
}

// Calculate eclipse for a specific location
geopos := [3]float64{-0.1278, 51.5074, 0} // London: lon, lat, altitude
localEclipse := swisseph.SolEclipseWhenLoc(jd, swisseph.FlagSwieph, 
    geopos, false)
if localEclipse.Flag >= 0 {
    fmt.Printf("Eclipse visible from location at JD %.6f\n", 
        localEclipse.Maximum)
}
```

### Rise, Set, and Transit Times

```go
// Calculate sunrise
geopos := [3]float64{-0.1278, 51.5074, 0} // London
sunrise := swisseph.RiseTrans(jd, swisseph.Sun, "", 
    swisseph.FlagSwieph, swisseph.CalcRise, geopos, 1013.25, 15.0)
if sunrise.Flag >= 0 {
    fmt.Printf("Sunrise: JD %.6f\n", sunrise.Time)
    
    // Convert to readable time
    utc := swisseph.JdetToUtc(sunrise.Time, swisseph.GregCal)
    fmt.Printf("Sunrise: %04d-%02d-%02d %02d:%02d:%.0f UTC\n",
        utc.Year, utc.Month, utc.Day, utc.Hour, utc.Minute, utc.Second)
}

// Calculate sunset
sunset := swisseph.RiseTrans(jd, swisseph.Sun, "", 
    swisseph.FlagSwieph, swisseph.CalcSet, geopos, 1013.25, 15.0)

// Calculate transit (culmination)
transit := swisseph.RiseTrans(jd, swisseph.Sun, "", 
    swisseph.FlagSwieph, swisseph.CalcMtransit, geopos, 1013.25, 15.0)
```

### Sidereal Calculations

```go
// Set sidereal mode
swisseph.SetSidMode(swisseph.SidmLahiri, 0, 0)

// Calculate with sidereal zodiac
result := swisseph.CalcUT(jd, swisseph.Sun, 
    swisseph.FlagSwieph|swisseph.FlagSidereal)
fmt.Printf("Sidereal Sun: %.6f°\n", result.Data[0])

// Get ayanamsa
ayanamsa := swisseph.GetAyanamsaUT(jd)
fmt.Printf("Ayanamsa: %.6f°\n", ayanamsa)

// Get ayanamsa name
name := swisseph.GetAyanamsaName(swisseph.SidmLahiri)
fmt.Printf("Ayanamsa system: %s\n", name)
```

### Topocentric Calculations

```go
// Set geographic location for topocentric calculations
swisseph.SetTopo(-0.1278, 51.5074, 0) // London

// Calculate with topocentric flag
result := swisseph.CalcUT(jd, swisseph.Moon, 
    swisseph.FlagSwieph|swisseph.FlagTopoctr)
fmt.Printf("Topocentric Moon: %.6f°\n", result.Data[0])
```

### Coordinate Transformations

```go
// Convert ecliptic to horizontal coordinates
geopos := [3]float64{-0.1278, 51.5074, 0}
ecliptic := [3]float64{120.0, 5.0, 1.0} // lon, lat, distance

azalt := swisseph.Azalt(jd, swisseph.Ecl2Hor, geopos, 1013.25, 15.0, ecliptic)
fmt.Printf("Azimuth: %.6f°\n", azalt.Azimuth)
fmt.Printf("Altitude: %.6f°\n", azalt.Altitude)
fmt.Printf("Apparent Altitude: %.6f°\n", azalt.AppAlt)

// Convert back
horizontal := [3]float64{azalt.Azimuth, azalt.Altitude, 1.0}
eclBack := swisseph.AzaltRev(jd, swisseph.Hor2Ecl, geopos, horizontal)
fmt.Printf("Back to ecliptic: %.6f°, %.6f°\n", eclBack[0], eclBack[1])
```

### Utility Functions

```go
// Normalize degrees
normalized := swisseph.Degnorm(375.5) // Returns 15.5

// Calculate difference between two positions
diff := swisseph.Difdegn(350.0, 10.0) // Returns -20.0 (shortest arc)

// Split degrees into components
split := swisseph.SplitDeg(123.456789, swisseph.SplitDegZodiacal)
fmt.Printf("Sign: %d, Degree: %d, Minute: %d, Second: %d\n",
    split.Sign, split.Degree, split.Minute, split.Second)

// Get day of week (0=Monday, 6=Sunday)
dow := swisseph.DayOfWeek(jd)
days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", 
    "Friday", "Saturday", "Sunday"}
fmt.Printf("Day of week: %s\n", days[dow])
```

## Ephemeris Files

Ephemeris files are required to enable high-precision calculations for planets and asteroids. This library does not include any ephemeris files by default, but you can download them from the official sources:

### Download Sources

- **Swiss Ephemeris GitHub**: [https://github.com/aloistr/swisseph/tree/master/ephe](https://github.com/aloistr/swisseph/tree/master/ephe) (main planets only)
- **Astrodienst Dropbox**: [https://www.dropbox.com/scl/fo/y3naz62gy6f6qfrhquu7u/h?rlkey=ejltdhb262zglm7eo6yfj2940&dl=0](https://www.dropbox.com/scl/fo/y3naz62gy6f6qfrhquu7u/h?rlkey=ejltdhb262zglm7eo6yfj2940&dl=0) (all files)

### File Types and Coverage

Each main ephemeris file covers a range of 600 years starting from the century indicated in its name. For example, the file `sepl_18.se1` is valid from year 1800 until year 2400.

#### Main Files
- **sepl files** - Planets (600-year range)
- **seplm files** - Planets for BC dates (600-year range)
- **semo files** - Moon (600-year range)
- **semom files** - Moon for BC dates (600-year range)
- **seas files** - Main asteroids (600-year range)
- **seasm files** - Main asteroids for BC dates (600-year range)

#### Advanced Files (from Dropbox)
- **all_ast folder** - Files for all asteroids (600-year range)
- **long_ast folder** - Files for named asteroids (6000-year range)
- **JPL binary files** - Files for NASA's JPL ephemerides (`de*.eph`)
- **ephe/sat folder** - Files for planetary moons

### Installation Steps

1. **Download the files** you need from one of the sources above
2. **Create a directory** for the ephemeris files (e.g., `./ephe`)
3. **Place the files** in the directory:
   - Main planet/moon files go in the root of the directory
   - Asteroid files should be placed in `astxxx` subfolders (same structure as in Dropbox)
   - Planetary moon files should be placed in a `sat` subfolder
4. **Set the path** in your Go code:

```go
swisseph.SetEphePath("./ephe")
```

### Example Directory Structure

```
./ephe/
├── sepl_18.se1          # Planets 1800-2400
├── semo_18.se1          # Moon 1800-2400
├── seas_18.se1          # Main asteroids 1800-2400
├── de431.eph            # JPL ephemeris (optional)
├── sat/                 # Planetary moons (optional)
│   └── ...
└── ast000/              # Asteroids (optional)
    ├── ast001.se1
    └── ...
```

### Usage Example

```go
package main

import (
    "fmt"
    swisseph "github.com/tejzpr/go-swisseph"
)

func main() {
    // Set ephemeris path before calculations
    swisseph.SetEphePath("./ephe")
    
    // Now calculations will use high-precision ephemeris data
    jd := swisseph.Julday(2024, 1, 1, 12.0, swisseph.GregCal)
    result := swisseph.CalcUT(jd, swisseph.Mars, swisseph.FlagSwieph)
    
    if result.Flag >= 0 {
        fmt.Printf("Mars: %.6f°\n", result.Data[0])
    }
    
    swisseph.Close()
}
```

### Notes

- If no ephemeris path is set, the library will use built-in Moshier ephemeris (lower precision)
- For most astrological applications, downloading `sepl_18.se1` and `semo_18.se1` is sufficient
- JPL ephemeris files provide the highest precision but are much larger
- More information can be found in the [Swiss Ephemeris documentation](https://www.astro.com/ftp/swisseph/doc/swisseph.htm)

## Planet Numbers

```go
const (
    Sun      = 0
    Moon     = 1
    Mercury  = 2
    Venus    = 3
    Mars     = 4
    Jupiter  = 5
    Saturn   = 6
    Uranus   = 7
    Neptune  = 8
    Pluto    = 9
    MeanNode = 10
    TrueNode = 11
    Chiron   = 15
    // ... and many more
)
```

## Calculation Flags

```go
const (
    FlagSwieph    = 2      // Use Swiss Ephemeris
    FlagSpeed     = 256    // Calculate speed
    FlagSidereal  = 65536  // Sidereal zodiac
    FlagTopoctr   = 32768  // Topocentric
    FlagEquatorial = 2048  // Equatorial coordinates
    FlagRadians   = 8192   // Return radians instead of degrees
    // ... and many more
)
```

## Documentation

For detailed information about the Swiss Ephemeris, see:

- [Official Programmer's Documentation](https://www.astro.com/swisseph/swephprg.htm)
- [Official User Guide](https://www.astro.com/ftp/swisseph/doc/swisseph.htm)

## Troubleshooting

### Compilation Errors: "swephexp.h: No such file or directory"

This should not occur as the header files are included in the repository. If you see this error, ensure you have the latest version of the repository with the `swisseph/` directory.

### Tests Fail with "no Go files"

Make sure you're in the correct directory:

```bash
cd /path/to/go-swisseph
go test -v
```

### Updating Swiss Ephemeris Version

To update to a newer version of Swiss Ephemeris, you would need to manually copy the updated C source files from the Swiss Ephemeris repository to the `swisseph/` directory. However, this is typically not necessary as the included version is stable and well-tested.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

### Go Bindings

This Go binding library is licensed under the **GNU Affero General Public License version 3 or later (AGPL-3.0-or-later)**.

See [LICENSE](LICENSE) file for the full AGPL-3.0 license text.

### Swiss Ephemeris

Swiss Ephemeris: Copyright © 1997-2021 Astrodienst AG, Switzerland

The Swiss Ephemeris library source code is included directly in this repository in the `swisseph/` directory, licensed under AGPL-3.0-or-later. Users must comply with Swiss Ephemeris licensing terms (AGPL or Professional License).

### Go Bindings Copyright

Go Bindings: Copyright © 2024-2025 Tejus Pratap
