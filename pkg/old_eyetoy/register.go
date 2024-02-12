package old_eyetoy

import (
	"fmt"

	"github.com/asciifaceman/hobocode"
)

const (
	EN_CLK0 uint16 = 0x53
	EN_CLK1 uint16 = 0x54
)

// OVRegVal holds register and value combinations for
// sending to the device
type OVRegVal struct {
	Reg uint16
	Val uint8
}

// et_init defines reg/vals that initialize the sony eyetoy
// these are largely lifted from an OV519 implementation
// and all instructions might not actually be doing something
// important or be accurate
var et_init []*OVRegVal = []*OVRegVal{
	{OV519_YS_CTRL, 0x6d}, /* EnableSystem */
	{OV519_EN_CLK0, 0x9b},
	{OV519_EN_CLK1, 0xff}, /* set bit2 to enable jpeg */
	{0x5d, 0x03},
	{0x49, 0x01}, // UV[7] - I/O - UV Bit [7]
	{0x48, 0x00},
	/* Set LED pin to output mode. Bit 4 must be cleared or sensor
	 * detection will fail. This deserves further investigation. */
	{OV519_GPIO_IO_CTRL0, 0xee},
	{OV519_RESET1, 0x0f}, /* SetUsbInit */
	{OV519_RESET1, 0x00},
	{0x22, 0x00},
	/* windows reads 0x55 at this point*/
}

//

// ReadRegister reads the value from the given index and returns
// a uint8 of the byte set there
func (e *EyeToy) ReadRegister(index uint16) (uint8, int, error) {
	hobocode.Idebugf(1, "Reading register values from EyeToy: [%v]", index)

	var rtype uint8 = USB_DIR_IN | USB_TYPE_VENDOR | USB_RECIP_DEVICE
	buf := make([]byte, 1)
	ret, err := e.Device.Control(rtype, ReqIO519, 0, index, buf)
	return buf[0], ret, err
}

/*
WriteRegister writes the given uint8 value to the given uint16 register on
the Eyetoy
*/
func (e *EyeToy) WriteRegister(index uint16, value uint8) (int, error) {
	hobocode.Idebugf(1, "Writing register values to EyeToy: [%x] => [%x] (%v)", index, value, value)

	var rtype uint8 = USB_DIR_OUT | USB_TYPE_VENDOR | USB_RECIP_DEVICE
	data := []byte{value}

	ret, err := e.Device.Control(rtype, ReqIO519, 0, index, data)
	if err != nil {
		return ret, err
	}
	if ret < 0 {
		return ret, fmt.Errorf("WriteRegister [%x]:[%x] received code <0", index, value)
	}
	return ret, err
}

// InitializeController the device with a series of register writes over the wire to enable the system
// and perform initial configuration
func (e *EyeToy) InitializeController() error {
	hobocode.Info("Initializing...")
	for _, rv := range et_init {
		code, err := e.WriteRegister(rv.Reg, rv.Val)
		if err != nil {
			return err
		}
		if code < 0 {
			return fmt.Errorf("write to register [%v] with byte {%v} succeeded but device returned a <0 status", rv.Reg, rv.Val)
		}
	}
	hobocode.Success("Initialized.")
	return nil
}
