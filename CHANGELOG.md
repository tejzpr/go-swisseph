# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2024-11-08

### Added

#### Core Functionality
- Complete CGO bindings to Swiss Ephemeris 2.10.03b C library
- Full API coverage with 100+ functions
- Type-safe Go interfaces with proper error handling
- Comprehensive constant definitions

#### Calculation Functions
- Planetary position calculations (`Calc`, `CalcUT`, `CalcPctr`)
- Fixed star calculations (`Fixstar`, `FixstarUT`, `Fixstar2`, etc.)
- House system calculations (Placidus, Koch, and 20+ other systems)
- Eclipse calculations (solar and lunar)
- Occultation calculations
- Rise/set/transit calculations
- Heliacal rising and setting calculations
- Nodes and apsides calculations
- Orbital elements calculations

#### Date and Time Functions
- Julian day conversions (`Julday`, `Revjul`)
- UTC to Julian day conversions (`UtcToJd`)
- Julian day to UTC conversions (`JdetToUtc`, `Jdut1ToUtc`)
- Calendar conversions (`DateConversion`)
- Time zone conversions (`UtcTimeZone`)
- Delta T calculations (`Deltat`, `DeltatEx`)
- Sidereal time calculations (`Sidtime`, `Sidtime0`)
- Equation of time (`TimeEqu`)
- Local mean/apparent time conversions (`LmtToLat`, `LatToLmt`)

#### Coordinate Transformations
- Azimuth/altitude calculations (`Azalt`, `AzaltRev`)
- Coordinate transformations (`Cotrans`, `CotransSp`)
- Atmospheric refraction (`Refrac`, `RefracExtended`)

#### Utility Functions
- Degree normalization (`Degnorm`, `Radnorm`)
- Angle difference calculations (`Difdegn`, `Difdeg2n`)
- Degree splitting (`SplitDeg`)
- String formatting functions (`Cs2timestr`, `Cs2lonlatstr`, `Cs2degstr`)
- Day of week calculation (`DayOfWeek`)

#### Configuration Functions
- Ephemeris path setting (`SetEphePath`)
- JPL file setting (`SetJplFile`)
- Sidereal mode setting (`SetSidMode`)
- Topocentric location setting (`SetTopo`)
- Tidal acceleration setting (`SetTidAcc`)
- Delta T user definition (`SetDeltaTUserdef`)

#### Information Functions
- Library version (`Version`)
- Library path retrieval (`GetLibraryPath`)
- Planet name retrieval (`GetPlanetName`)
- Ayanamsa name retrieval (`GetAyanamsaName`)
- Current file data (`GetCurrentFileData`)

#### Examples
- Basic planetary calculations example
- Natal chart calculation example
- Eclipse finder example
- Rise/set times calculator example

#### Documentation
- Comprehensive README with usage examples
- API documentation with code examples
- Contributing guidelines
- Detailed LICENSE information
- CHANGELOG

#### Testing
- Unit tests for core functions
- Benchmark tests for performance measurement
- Test coverage reporting

#### Build System
- Makefile with common tasks
- Go module configuration
- .gitignore for clean repository

### Technical Details

#### Supported Platforms
- Linux (x86_64, ARM64)
- macOS (Intel, Apple Silicon)
- Windows (x86_64)

#### Dependencies
- Go 1.21 or later
- C compiler (gcc, clang, or MSVC)
- Swiss Ephemeris C library (included)

#### Performance
- Direct CGO bindings for optimal performance
- Minimal overhead over C library
- Efficient memory management

### Notes

This is the initial release of Go Swiss Ephemeris, providing complete bindings
to the Swiss Ephemeris library version 2.10.03b. All major functionality from
the original C library is available through idiomatic Go interfaces.

The library follows the Swiss Ephemeris licensing model, offering dual licensing
under AGPL-3.0-or-later for open source projects and LGPL-3.0-or-later for
commercial projects with a professional license.

[1.0.0]: https://github.com/tejzpr/go-swisseph/releases/tag/v1.0.0

