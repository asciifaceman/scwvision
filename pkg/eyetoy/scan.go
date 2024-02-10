/*
	Scan USB devices for the correct eyetoy device
*/
package eyetoy

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/asciifaceman/hobocode"
	"github.com/karalabe/usb"
)

func Scan() ([]usb.DeviceInfo, error) {
	if !usb.Supported() {
		return nil, fmt.Errorf("platform not supported for USB operations")
	}

	devices, err := usb.EnumerateHid(SonyEyeToyVendorID, SonyEyeToyProductID)
	if err != nil {
		return nil, err
	}

	//for _, device := range devices {
	//	d, e := device.Open()
	//	if e != nil {
	//		hobocode.Errorf("Failed to open device %d:%d (%v)", device.VendorID, device.ProductID, e)
	//		continue
	//	}
	//	err := d.Close()
	//	if err != nil {
	//		hobocode.Errorf("Error closing device %d:%d (%v)", device.VendorID, device.ProductID, err)
	//	}
	//}

	return devices, nil
}

func Display(devices []usb.DeviceInfo) {
	for _, device := range devices {
		splitPath := strings.Split(device.Path, ":") // bus, device, id
		bus := new(big.Int)
		bus.SetString(splitPath[0], 16)
		dev := new(big.Int)
		dev.SetString(splitPath[1], 16)

		hobocode.HeaderLeft(fmt.Sprintf("Device [%v]:[%v]:[%v]", bus, dev, splitPath[2]))
		hobocode.Infof("VendorID: %d", device.VendorID)
		hobocode.Infof("ProductID: %d", device.ProductID)
	}
}
