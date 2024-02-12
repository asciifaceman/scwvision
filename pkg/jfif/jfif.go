/*
jfif simply provides some tools for inspecting a potentially-partial
stream of jpeg data - I promise I'm not trying to reinvent the
wheel here, but most jpeg libraries assume you're working with
valid JFIF out of the gate and I want to be able to pick apart
the file as it comes in over the read stream

I'm stupid what can I say
*/
package jfif

import "fmt"

type Segment struct {
	Raw    []byte
	Length int // bits 0, 1
}

type ApplicationHeader struct {
	Segment
	Identifier   string // bits 2:6
	VersionMajor int    // bits 7
	VersionMinor int    // bits 8
	Units        int    // bits 9
	DensityX     int    // bits 10:11
	DensityY     int    // bits 12:13
	ThumbnailX   int    // bits 14
	ThumbnailY   int    // bits 15
}

// FormatVersion returns nicely formatted version string
func (ah *ApplicationHeader) FormatVersion() string {
	return fmt.Sprintf("%d.%d", ah.VersionMajor, ah.VersionMinor)
}

// FormatDensity returns a formatted density string
func (ah *ApplicationHeader) FormatDensity() string {
	return fmt.Sprintf("%dx%d", ah.DensityX, ah.DensityY)
}

// FormatThumbnail returns a formatted thumbnail dimension string
func (ah *ApplicationHeader) FormatThumbnail() string {
	return fmt.Sprintf("%dx%d", ah.ThumbnailX, ah.ThumbnailY)
}

type QuantizationTable struct {
	Raw []byte // below bits begin after MARKER_QUANTABLE
}
