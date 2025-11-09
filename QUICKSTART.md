# Quick Start Guide

Get up and running with Go Swiss Ephemeris in 5 minutes!

## Installation

```bash
# Clone the repository
git clone https://github.com/tejzpr/go-swisseph.git
cd go-swisseph

# IMPORTANT: Initialize the Swiss Ephemeris submodule
git submodule update --init --recursive

# Verify installation
go test -v -run TestVersion
```

**Note**: The Swiss Ephemeris C library is included as a Git submodule. You must run `git submodule update --init --recursive` before building!

## Your First Program

Create a file named `main.go`:

```go
package main

import (
    "fmt"
    swisseph "github.com/tejzpr/go-swisseph"
)

func main() {
    // Calculate Julian day for January 1, 2024, 12:00 UTC
    jd := swisseph.Julday(2024, 1, 1, 12.0, swisseph.GregCal)
    
    // Calculate Sun position
    result := swisseph.CalcUT(jd, swisseph.Sun, swisseph.FlagSwieph)
    
    if result.Flag >= 0 {
        fmt.Printf("Sun longitude: %.6f°\n", result.Data[0])
    } else {
        fmt.Printf("Error: %s\n", result.Error)
    }
    
    // Clean up
    swisseph.Close()
}
```

Run it:

```bash
go run main.go
```

Output:
```
Sun longitude: 280.234567°
```

## Common Tasks

### Calculate All Planets

```go
planets := []int32{
    swisseph.Sun, swisseph.Moon, swisseph.Mercury, swisseph.Venus,
    swisseph.Mars, swisseph.Jupiter, swisseph.Saturn, swisseph.Uranus,
    swisseph.Neptune, swisseph.Pluto,
}

jd := swisseph.Julday(2024, 1, 1, 12.0, swisseph.GregCal)

for _, planet := range planets {
    name := swisseph.GetPlanetName(planet)
    result := swisseph.CalcUT(jd, planet, swisseph.FlagSwieph)
    
    if result.Flag >= 0 {
        fmt.Printf("%s: %.2f°\n", name, result.Data[0])
    }
}
```

### Calculate Houses

```go
jd := swisseph.Julday(2024, 1, 1, 12.0, swisseph.GregCal)
lat := 51.5074  // London latitude
lon := -0.1278  // London longitude

houses := swisseph.Houses(jd, lat, lon, 'P') // Placidus system

if houses.Flag >= 0 {
    fmt.Printf("Ascendant: %.2f°\n", houses.Points[swisseph.Asc])
    fmt.Printf("MC: %.2f°\n", houses.Points[swisseph.MC])
    
    for i, cusp := range houses.Houses {
        fmt.Printf("House %d: %.2f°\n", i+1, cusp)
    }
}
```

### Calculate Sunrise/Sunset

```go
jd := swisseph.Julday(2024, 1, 1, 0.0, swisseph.GregCal)
geopos := [3]float64{-0.1278, 51.5074, 0} // London

// Sunrise
sunrise := swisseph.RiseTrans(jd, swisseph.Sun, "", 
    swisseph.FlagSwieph, swisseph.CalcRise, geopos, 1013.25, 15.0)

if sunrise.Flag >= 0 {
    time := swisseph.Revjul(sunrise.Time, swisseph.GregCal)
    fmt.Printf("Sunrise: %02.0f:%02.0f UTC\n", 
        time.Hour, (time.Hour-float64(int(time.Hour)))*60)
}

// Sunset
sunset := swisseph.RiseTrans(jd, swisseph.Sun, "", 
    swisseph.FlagSwieph, swisseph.CalcSet, geopos, 1013.25, 15.0)

if sunset.Flag >= 0 {
    time := swisseph.Revjul(sunset.Time, swisseph.GregCal)
    fmt.Printf("Sunset: %02.0f:%02.0f UTC\n", 
        time.Hour, (time.Hour-float64(int(time.Hour)))*60)
}
```

### Find Next Eclipse

```go
jd := swisseph.Julday(2024, 1, 1, 12.0, swisseph.GregCal)

eclipse := swisseph.SolEclipseWhenGlob(jd, swisseph.FlagSwieph, 
    swisseph.EclAlltypesSolar, false)

if eclipse.Flag >= 0 {
    date := swisseph.Revjul(eclipse.Maximum, swisseph.GregCal)
    fmt.Printf("Next solar eclipse: %04d-%02d-%02d\n", 
        date.Year, date.Month, date.Day)
    
    if eclipse.Flag & swisseph.EclTotal != 0 {
        fmt.Println("Type: Total")
    } else if eclipse.Flag & swisseph.EclAnnular != 0 {
        fmt.Println("Type: Annular")
    }
}
```

### Use Sidereal Zodiac

```go
// Set sidereal mode (Lahiri ayanamsa)
swisseph.SetSidMode(swisseph.SidmLahiri, 0, 0)

jd := swisseph.Julday(2024, 1, 1, 12.0, swisseph.GregCal)

// Calculate with sidereal flag
result := swisseph.CalcUT(jd, swisseph.Sun, 
    swisseph.FlagSwieph | swisseph.FlagSidereal)

fmt.Printf("Sidereal Sun: %.2f°\n", result.Data[0])

// Get ayanamsa
ayanamsa := swisseph.GetAyanamsaUT(jd)
fmt.Printf("Ayanamsa: %.2f°\n", ayanamsa)
```

## Important Constants

### Planets
```go
swisseph.Sun      // 0
swisseph.Moon     // 1
swisseph.Mercury  // 2
swisseph.Venus    // 3
swisseph.Mars     // 4
swisseph.Jupiter  // 5
swisseph.Saturn   // 6
swisseph.Uranus   // 7
swisseph.Neptune  // 8
swisseph.Pluto    // 9
```

### Calculation Flags
```go
swisseph.FlagSwieph    // Use Swiss Ephemeris (default)
swisseph.FlagSpeed     // Calculate speed
swisseph.FlagSidereal  // Sidereal zodiac
swisseph.FlagTopoctr   // Topocentric
```

### House Systems
```go
'P' // Placidus (most common)
'K' // Koch
'O' // Porphyrius
'R' // Regiomontanus
'C' // Campanus
'E' // Equal (from Asc)
'W' // Whole sign
```

## Tips

1. **Always call `swisseph.Close()`** at the end of your program to free resources
2. **Check return flags** - negative flags indicate errors
3. **Use ephemeris files** for high precision (download separately)
4. **Set ephemeris path** if using external files: `swisseph.SetEphePath("./ephe")`
5. **Check the examples** directory for more complete programs

## Next Steps

- Read the [full README](README.md) for detailed documentation
- Check out the [examples](examples/) directory
- Review the [API documentation](https://pkg.go.dev/github.com/tejzpr/go-swisseph)
- Learn about [Swiss Ephemeris](https://www.astro.com/swisseph/)

## Need Help?

- Check the [examples](examples/) directory
- Read the [official Swiss Ephemeris documentation](https://www.astro.com/swisseph/swephprg.htm)
- Open an issue on GitHub
- Review the test files for more usage examples

Happy calculating! 🌟

