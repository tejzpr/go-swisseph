// Package swisseph provides Go bindings for the Swiss Ephemeris library.
// This is a comprehensive astronomical calculation library for astrology and astronomy.
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
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package swisseph

// Return codes
const (
	OK  = 0
	ERR = -1
)

// Astronomical unit conversions
const (
	AunitToKm        = 149597870.700
	AunitToLightyear = 1.0 / 63241.07708427
	AunitToParsec    = 1.0 / 206264.8062471
)

// Calendar types
const (
	JulCal  = 0 // Julian calendar
	GregCal = 1 // Gregorian calendar
)

// Planet numbers for calculations
const (
	EclNut   = -1 // Eclipse/nutation
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
	MeanApog = 12
	OscuApog = 13
	Earth    = 14
	Chiron   = 15
	Pholus   = 16
	Ceres    = 17
	Pallas   = 18
	Juno     = 19
	Vesta    = 20
	IntpApog = 21
	IntpPerg = 22
)

// Number of planets and offsets
const (
	NPlanets      = 23
	PlmoonOffset  = 9000
	AstOffset     = 10000
	Varuna        = AstOffset + 20000
	FictOffset    = 40
	FictOffset1   = 39
	FictMax       = 999
	NFictElem     = 15
	CometOffset   = 1000
	NAllNatPoints = NPlanets + NFictElem
)

// Hamburger/Uranian planets and fictitious bodies
const (
	Cupido           = 40
	Hades            = 41
	Zeus             = 42
	Kronos           = 43
	Apollon          = 44
	Admetos          = 45
	Vulkanus         = 46
	Poseidon         = 47
	Isis             = 48
	Nibiru           = 49
	Harrington       = 50
	NeptuneLeverrier = 51
	NeptuneAdams     = 52
	PlutoLowell      = 53
	PlutoPickering   = 54
	Vulcan           = 55
	WhiteMoon        = 56
	Proserpina       = 57
	Waldemath        = 58
)

// Fixed stars
const (
	FixstarFlag = -10
)

// House cusps and angles
const (
	Asc    = 0 // Ascendant
	MC     = 1 // Medium Coeli (Midheaven)
	ARMC   = 2 // Sidereal time
	Vertex = 3
	Equasc = 4 // Equatorial ascendant
	Coasc1 = 5 // Co-ascendant (W. Koch)
	Coasc2 = 6 // Co-ascendant (M. Munkasey)
	Polasc = 7 // Polar ascendant (M. Munkasey)
	Nascmc = 8
)

// Calculation flags
const (
	FlagJpleph       = 1    // Use JPL ephemeris
	FlagSwieph       = 2    // Use Swiss Ephemeris
	FlagMoseph       = 4    // Use Moshier ephemeris
	FlagHelctr       = 8    // Heliocentric position
	FlagTruepos      = 16   // True/geometric position
	FlagJ2000        = 32   // J2000 equinox
	FlagNonut        = 64   // No nutation
	FlagSpeed3       = 128  // Speed from 3 positions (deprecated)
	FlagSpeed        = 256  // High precision speed
	FlagNogdefl      = 512  // No gravitational deflection
	FlagNoaberr      = 1024 // No aberration
	FlagAstrometric  = FlagNoaberr | FlagNogdefl
	FlagEquatorial   = 2048  // Equatorial coordinates
	FlagXYZ          = 4096  // Cartesian coordinates
	FlagRadians      = 8192  // Radians instead of degrees
	FlagBaryctr      = 16384 // Barycentric position
	FlagTopoctr      = 32768 // Topocentric position
	FlagOrbelAA      = FlagTopoctr
	FlagTropical     = 0      // Tropical zodiac (default)
	FlagSidereal     = 65536  // Sidereal zodiac
	FlagICRS         = 131072 // ICRS reference frame
	FlagDpsideps1980 = 262144 // IAU 1980 nutation
	FlagJplhor       = FlagDpsideps1980
	FlagJplhorApprox = 524288  // JPL Horizons approximation
	FlagCenterBody   = 1048576 // Center body
	FlagTestPlmoon   = 2097152 | FlagJ2000 | FlagICRS | FlagHelctr | FlagTruepos
)

// Sidereal mode bits
const (
	Sidbits            = 256
	SidbitEclT0        = 256
	SidbitSSYPlane     = 512
	SidbitUserUT       = 1024
	SidbitEclDate      = 2048
	SidbitNoPrecOffset = 4096
	SidbitPrecOrig     = 8192
)

// Sidereal modes
const (
	SidmFaganBradley       = 0
	SidmLahiri             = 1
	SidmDeluce             = 2
	SidmRaman              = 3
	SidmUshashashi         = 4
	SidmKrishnamurti       = 5
	SidmDjwhalKhul         = 6
	SidmYukteshwar         = 7
	SidmJNBhasin           = 8
	SidmBabylKugler1       = 9
	SidmBabylKugler2       = 10
	SidmBabylKugler3       = 11
	SidmBabylHuber         = 12
	SidmBabylEtpsc         = 13
	SidmAldebaran15Tau     = 14
	SidmHipparchos         = 15
	SidmSassanian          = 16
	SidmGalcent0Sag        = 17
	SidmJ2000              = 18
	SidmJ1900              = 19
	SidmB1950              = 20
	SidmSuryasiddhanta     = 21
	SidmSuryasiddhantaMsun = 22
	SidmAryabhata          = 23
	SidmAryabhataMsun      = 24
	SidmSSRevati           = 25
	SidmSSCitra            = 26
	SidmTrueCitra          = 27
	SidmTrueRevati         = 28
	SidmTruePushya         = 29
	SidmGalcentRgilbrand   = 30
	SidmGalequIAU1958      = 31
	SidmGalequTrue         = 32
	SidmGalequMula         = 33
	SidmGalalignMardyks    = 34
	SidmTrueMula           = 35
	SidmGalcentMulaWilhelm = 36
	SidmAryabhata522       = 37
	SidmBabylBritton       = 38
	SidmTrueSheoran        = 39
	SidmGalcentCochrane    = 40
	SidmGalequFiorenza     = 41
	SidmValensMoon         = 42
	SidmLahiri1940         = 43
	SidmLahiriVP285        = 44
	SidmKrishnamurtiVP291  = 45
	SidmLahiriICRC         = 46
	SidmUser               = 255
	NsidmPredef            = 47
)

// Node and apsides bits
const (
	NodbitMean    = 1
	NodbitOscu    = 2
	NodbitOscuBar = 4
	NodbitFopoint = 256
)

// Default ephemeris
const (
	FlagDefaulteph = FlagSwieph
)

// Maximum star name length
const (
	MaxStname = 256
)

// Eclipse types
const (
	EclCentral          = 1
	EclNoncentral       = 2
	EclTotal            = 4
	EclAnnular          = 8
	EclPartial          = 16
	EclAnnularTotal     = 32
	EclHybrid           = 32
	EclPenumbral        = 64
	EclAlltypesSolar    = EclCentral | EclNoncentral | EclTotal | EclAnnular | EclPartial | EclAnnularTotal
	EclAlltypesLunar    = EclTotal | EclPartial | EclPenumbral
	EclVisible          = 128
	EclMaxVisible       = 256
	Ecl1stVisible       = 512
	EclPartbegVisible   = 512
	Ecl2ndVisible       = 1024
	EclTotbegVisible    = 1024
	Ecl3rdVisible       = 2048
	EclTotendVisible    = 2048
	Ecl4thVisible       = 4096
	EclPartendVisible   = 4096
	EclPenumbbegVisible = 8192
	EclPenumbendVisible = 16384
	EclOccBegDaylight   = 8192
	EclOccEndDaylight   = 16384
	EclOneTry           = 32768
)

// Rise/set calculation flags
const (
	CalcRise           = 1
	CalcSet            = 2
	CalcMtransit       = 4
	CalcItransit       = 8
	BitDiscCenter      = 256
	BitDiscBottom      = 8192
	BitGeoctrNoEclLat  = 128
	BitNoRefraction    = 512
	BitCivilTwilight   = 1024
	BitNauticTwilight  = 2048
	BitAstroTwilight   = 4096
	BitFixedDiscSize   = 16384
	BitForceSlowMethod = 32768
	BitHinduRising     = BitDiscCenter | BitNoRefraction | BitGeoctrNoEclLat
)

// Coordinate transformation
const (
	Ecl2Hor = 0
	Equ2Hor = 1
	Hor2Ecl = 0
	Hor2Equ = 1
)

// Refraction
const (
	TrueToApp = 0
	AppToTrue = 1
)

// JPL ephemeris files
const (
	DeNumber    = 431
	FnameDe200  = "de200.eph"
	FnameDe403  = "de403.eph"
	FnameDe404  = "de404.eph"
	FnameDe405  = "de405.eph"
	FnameDe406  = "de406.eph"
	FnameDe431  = "de431.eph"
	FnameDft    = FnameDe431
	FnameDft2   = FnameDe406
	StarfileOld = "fixstars.cat"
	Starfile    = "sefstars.txt"
	Astnamfile  = "seasnam.txt"
	Fictfile    = "seorbel.txt"
)

// Split degree flags
const (
	SplitDegRoundSec  = 1
	SplitDegRoundMin  = 2
	SplitDegRoundDeg  = 4
	SplitDegZodiacal  = 8
	SplitDegNakshatra = 1024
	SplitDegKeepSign  = 16
	SplitDegKeepDeg   = 32
)

// Heliacal events
const (
	HeliacalRising    = 1
	HeliacalSetting   = 2
	MorningFirst      = HeliacalRising
	EveningLast       = HeliacalSetting
	EveningFirst      = 3
	MorningLast       = 4
	AcronychalRising  = 5
	AcronychalSetting = 6
	CosmicalSetting   = AcronychalSetting
)

// Heliacal flags
const (
	HelflagLongSearch     = 128
	HelflagHighPrecision  = 256
	HelflagOpticalParams  = 512
	HelflagNoDetails      = 1024
	HelflagSearch1Period  = 2048
	HelflagVislimDark     = 4096
	HelflagVislimNomoon   = 8192
	HelflagVislimPhotopic = 16384
	HelflagVislimScotopic = 32768
	HelflagAV             = 65536
	HelflagAvkindVR       = 65536
	HelflagAvkindPTO      = 131072
	HelflagAvkindMin7     = 262144
	HelflagAvkindMin9     = 524288
	HelflagAvkind         = HelflagAvkindVR | HelflagAvkindPTO | HelflagAvkindMin7 | HelflagAvkindMin9
	TjdInvalid            = 99999999.0
	SimulateVictorvb      = 1
)

// Photopic flags
const (
	PhotopicFlag  = 0
	ScotopicFlag  = 1
	MixedopicFlag = 2
)

// Tidal acceleration values
const (
	TidalDe200          = -23.8946
	TidalDe403          = -25.580
	TidalDe404          = -25.580
	TidalDe405          = -25.826
	TidalDe406          = -25.826
	TidalDe421          = -25.85
	TidalDe422          = -25.85
	TidalDe430          = -25.82
	TidalDe431          = -25.80
	Tidal26             = -26.0
	TidalStephenson2016 = -25.85
	TidalDefault        = TidalDe431
	TidalAutomatic      = 999999
	TidalMoseph         = TidalDe404
	TidalSwieph         = TidalDefault
	TidalJpleph         = TidalDefault
)

// Delta T
const (
	DeltatAutomatic = -1e-10
)

// Models
const (
	ModelDeltat        = 0
	ModelPrecLongterm  = 1
	ModelPrecShortterm = 2
	ModelNut           = 3
	ModelBias          = 4
	ModelJplhorMode    = 5
	ModelJplhoraMode   = 6
	ModelSidt          = 7
	NseModels          = 8
)

// Precession models
const (
	ModNprec             = 11
	ModPrecIAU1976       = 1
	ModPrecLaskar1986    = 2
	ModPrecWillEpsLask   = 3
	ModPrecWilliams1994  = 4
	ModPrecSimon1994     = 5
	ModPrecIAU2000       = 6
	ModPrecBretagnon2003 = 7
	ModPrecIAU2006       = 8
	ModPrecVondrak2011   = 9
	ModPrecOwen1990      = 10
	ModPrecNewcomb       = 11
	ModPrecDefault       = ModPrecVondrak2011
	ModPrecDefaultShort  = ModPrecVondrak2011
)

// Nutation models
const (
	ModNnut           = 5
	ModNutIAU1980     = 1
	ModNutIAUCorr1987 = 2
	ModNutIAU2000A    = 3
	ModNutIAU2000B    = 4
	ModNutWoolard     = 5
	ModNutDefault     = ModNutIAU2000B
)

// Sidereal time models
const (
	ModNsidt            = 4
	ModSidtIAU1976      = 1
	ModSidtIAU2006      = 2
	ModSidtIERSConv2010 = 3
	ModSidtLongterm     = 4
	ModSidtDefault      = ModSidtLongterm
)

// Bias models
const (
	ModNbias       = 3
	ModBiasNone    = 1
	ModBiasIAU2000 = 2
	ModBiasIAU2006 = 3
	ModBiasDefault = ModBiasIAU2006
)

// JPL Horizons models
const (
	ModNjplhor             = 2
	ModJplhorLongAgreement = 1
	ModJplhorDefault       = ModJplhorLongAgreement
)

// JPL Horizons approximation models
const (
	ModNjplhora       = 3
	ModJplhora1       = 1
	ModJplhora2       = 2
	ModJplhora3       = 3
	ModJplhoraDefault = ModJplhora3
)

// Delta T models
const (
	ModNdeltat                      = 5
	ModDeltatStephensonMorrison1984 = 1
	ModDeltatStephenson1997         = 2
	ModDeltatStephensonMorrison2004 = 3
	ModDeltatEspenakMeeus2006       = 4
	ModDeltatStephensonEtc2016      = 5
	ModDeltatDefault                = ModDeltatStephensonEtc2016
)
