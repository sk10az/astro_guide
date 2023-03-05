package resources

import (
	"embed"
	//"io/fs"
)

//go:embed img/planets/earth_logo.png
var EarthLogo []byte

//go:embed img
var Content embed.FS

// ===== PLANETS =====

//go:embed img/planets/mercury.png
var MercuryPlanets []byte

//go:embed img/planets/venus.png
var VenusPlanets []byte

//go:embed img/planets/earth.png
var EarthPlanets []byte

//go:embed img/planets/mars.png
var MarsPlanets []byte
