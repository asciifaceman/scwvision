package eyetoy

import "github.com/google/gousb"

/*
Application Constants
*/
const (
	SonyEyeToyVendorID  uint16 = 0x054c // 1356 Sony Corp.
	SonyEyeToyProductID uint16 = 0x0154 // 340 Eyetoy Device
)

// findEyetoy is a convenience wrapper for discovering EyeToy devices
func findEyetoy() func(desc *gousb.DeviceDesc) bool {
	return func(desc *gousb.DeviceDesc) bool {
		return desc.Product == gousb.ID(SonyEyeToyProductID) && desc.Vendor == gousb.ID(SonyEyeToyVendorID)
	}
}

/*
Eyetoy USB Interface Constants
*/
const (
	EyeToyPrimaryConfig    int = 1 // only one configuration available
	EyeToyPrimaryInterface int = 0 // only one available interface
	EyeToyPrimaryEndpoint  int = 1
)

/*
OV519 Generic Constants
*/
const (
	RTYPE_READ                   = USB_DIR_IN | USB_TYPE_VENDOR | USB_RECIP_DEVICE
	RTYPE_WRITE                  = USB_DIR_OUT | USB_TYPE_VENDOR | USB_RECIP_DEVICE
	ReqIO519               uint8 = 1
	USB_TYPE_VENDOR        uint8 = (0x02 << 5) // 0x00000040 == 64
	USB_DIR_IN             uint8 = 0x00000080
	USB_DIR_OUT            uint8 = 0x00000000
	USB_RECIP_DEVICE       uint8 = 0x00
	OV519_ENDPOINT_ADDRESS int   = 1 /* Isoc endpoint number */
)

/*
OV519 Registers (controller)
*/
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

	OV519_CAM_H_SIZE    uint16 = 0x10
	OV519_CAM_V_SIZE    uint16 = 0x11
	OV519_CAM_X_OFFSETL uint16 = 0x12
	OV519_CAM_X_OFFSETH uint16 = 0x13
	OV519_CAM_Y_OFFSETL uint16 = 0x14
	OV519_CAM_Y_OFFSETH uint16 = 0x15
	OV519_CAM_DIVIDER   uint16 = 0x16
	OV519_CAM_DFR       uint16 = 0x20
	OV519_CAM_FORMAT    uint16 = 0x25
)

/*
OV7648 Registers (sensor)
*/
