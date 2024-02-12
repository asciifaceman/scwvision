/*
Constants for eyetoy package
*/
package old_eyetoy

const (
	OFF                    uint16 = 0
	ON                     uint16 = 1
	SonyEyeToyVendorID     uint16 = 1356 //0x054c Sony Corp.
	SonyEyeToyProductID    uint16 = 340  // 0x0154 Eyetoy Audio Device
	EyeToyPrimaryConfig    int    = 1    // only one configuration
	EyeToyPrimaryInterface int    = 0    // only one interface
	EyeToyPrimaryEndpoint  int    = 1
	ReqIO519               uint8  = 1
	USB_TYPE_VENDOR               = (0x02 << 5) // 0x00000040 >> 64
	USB_DIR_IN                    = 0x00000080
	USB_DIR_OUT                   = 0x00000000
	USB_RECIP_DEVICE              = 0x00
)

// Registers
const (
	OV519_RESET0         uint16 = 0x50
	OV519_RESET1         uint16 = 0x51
	OV519_EN_CLK0        uint16 = 0x53
	OV519_EN_CLK1        uint16 = 0x54
	OV519_AUDIO_CLK      uint16 = 0x55
	OV519_SNAPSHOT       uint16 = 0x57
	OV519_PONOFF         uint16 = 0x58
	OV519_CAMERA_CLOCK   uint16 = 0x59
	OV519_YS_CTRL        uint16 = 0x5A
	OV519_PWDN           uint16 = 0x5D
	OV519_GPIO_DATA_OUT0 uint16 = 0x71
	OV519_GPIO_IO_CTRL0  uint16 = 0x72
)

// Common Values
