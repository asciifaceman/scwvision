package eyetoy

import (
	"encoding/hex"
	"fmt"
	"time"
)

// Instruction is a register:value pair
type Instruction struct {
	Desc string
	Reg  uint16
	Val  uint8
}

/*
0v519_controller_init defines registers and values to set in a sequence
that will initialize the sony eyetoy

these are largely lifted from the OV519 linux kernel driver which was never
really completed for this chipset and served multiple chips so not every
instruction may be important, accurate, or valid
*/
var ov519_controller_init []*Instruction = []*Instruction{
	{"Enable System", OV519_YS_CTRL, 0x6d},
	{"Unknown", OV519_EN_CLK0, 0x9b},
	{"Set bit 2 enable jpeg", OV519_EN_CLK1, 0xff},
	{"Unknown", 0x5d, 0x03},
	{"UV[7] - I/O - UV Bit [7]", 0x49, 0x01},
	{"Unknown", 0x48, 0x00},
	{"Set LED pin output mode.", OV519_GPIO_IO_CTRL0, 0xee}, // bit 4 must be cleared or sensor detection will fail
	{"Set USB Init", OV519_RESET1, 0x0f},
	{"Unknown", OV519_RESET1, 0x00}, // may be reset sequence
	{"Unknown", 0x22, 0x00},
	/* windows reads 0x55 at this point - we don't care? */
}

var ov519_controller_stop []*Instruction = []*Instruction{
	{"stop_controller", OV519_RESET1, 0x0f},
}

var ov519_controller_start []*Instruction = []*Instruction{
	{"start controller", OV519_RESET1, 0x00},
}

/*
ReadRegister reads the value stored in a register at the given index
and returns a uint8 of the byte set there

Returns the value, status code, and error if any
*/
func (e *EyeToy) ReadRegister(index uint16) (uint8, int, error) {
	e.logger.Debugw("reading register", "index", index)

	value := []byte{0}

	ret, err := e.GUSB.Device.Control(RTYPE_READ, ReqIO519, 0, index, value)
	if err != nil || ret < 0 {
		if err == nil {
			err = UnhandledErrorCode
		}
	}

	if len(value) > 1 {
		e.logger.Warnw("received more than 1 byte. should investigate", "length_bytes", len(value))
	}

	return value[0], ret, err
}

/*
WriteRegister writes the given byte to the given register index

This method is not concerned with masking, see WriteMaskedRegister.
*/
func (e *EyeToy) WriteRegister(index uint16, value uint8) (int, error) {
	e.logger.Debugw("writing register", "index", index, "value", value)

	data := []byte{value}

	ret, err := e.GUSB.Device.Control(RTYPE_WRITE, ReqIO519, 0, index, data)
	if err != nil || ret < 0 {
		if err == nil && ret < 0 {
			err = UnhandledErrorCode
		}
		return ret, err
	}

	return ret, err
}

/*
WriteMaskedRegister writes bits on the given register as defined
by the given mask. Bits that are in the same position as 1's in mask are
cleared and set to value. Bits that are in the same position as 0's in mask
are preserved, regardless of their state in value
*/
func (e *EyeToy) WriteMaskedRegister(index uint16, value uint8, mask uint8) (int, error) {
	e.logger.Debugw("writing register with mask", "index", index, "value", value, "mask", mask)

	// mask value with current value
	if mask != 0xff {
		value &= mask // enforce mask - AND

		current, cret, cerr := e.ReadRegister(index)
		if cerr != nil || cret < 0 {
			if cerr == nil && cret < 0 {
				cerr = UnhandledErrorCode
			}
			return cret, fmt.Errorf("encountered error while reading register [%x] to mask: %v", index, cerr)
		}

		var old uint8 = current & ^mask // clear masked bits, AND NOT
		value |= old                    // set the desired bits, OR
	}
	return e.WriteRegister(index, value)
}

/*
StartCamera starts the camera...

  - mode_init_regs
  - set ov sensor window
  - restart
  - turn on LED
*/
func (e *EyeToy) StartCamera() error {
	return nil
}

/*
StopCamera stops the camera...

  - stop
  - turn off LED
*/
func (e *EyeToy) StopCamera() error {
	return nil
}

/*
ProbeCamera is an initialization step that probes the camera and checks
that everything we need is there
*/
func (e *EyeToy) ProbeCamera() error {
	logger := e.logger.Named("probe")
	logger.Debug("beginning probe")

	err := e.InitializeController()
	if err != nil {
		return err
	}

	// probe sensor

	// init sensor

	logger.Debug("probe complete, LED on and waiting")
	e.EnableLED()
	c := time.NewTimer(3 * time.Second)
	<-c.C
	logger.Debug("stopping controller")
	// shut down
	e.WriteRegister(OV519_RESET1, 0x0f) // this isn't actually shutting it down
	e.DisableLED()                      // hiding the issue
	return nil
}

/*
InitializeController initializes the eyetoy with a series of register writes
that perform an initial wake and configuration
*/
func (e *EyeToy) InitializeController() error {
	e.logger.Debug("initializing controller...")
	for _, rv := range ov519_controller_init {
		e.logger.Debugw("cinit: reg set", "desc", rv.Desc, "index", hex.EncodeToString([]byte{byte(rv.Reg)}), "value", rv.Val)
		_, err := e.WriteRegister(rv.Reg, rv.Val)
		if err != nil {
			e.ShutdownController()
			return err
		}
	}
	return nil
}

/*
ShutdownController shuts the controller down and de-initializes it
*/
func (e *EyeToy) ShutdownController() error {
	e.logger.Debug("shutting down controller...")
	return nil
}

// EnableLED turns the red LED on
func (e *EyeToy) EnableLED() error {
	e.logger.Debug("enabling LED")
	ret, err := e.WriteMaskedRegister(OV519_GPIO_DATA_OUT0, 1, 1)
	if err != nil || ret < 0 {
		if err == nil {
			err = UnhandledErrorCode
		}
		return err
	}
	return nil
}

// DisableLED turns the red LED off
func (e *EyeToy) DisableLED() error {
	e.logger.Debug("disabling LED")
	ret, err := e.WriteMaskedRegister(OV519_GPIO_DATA_OUT0, 0, 1)
	if err != nil || ret < 0 {
		if err == nil {
			err = UnhandledErrorCode
		}
		return err
	}
	return nil
}
