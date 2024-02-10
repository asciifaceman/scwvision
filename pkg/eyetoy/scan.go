/*
Scan USB devices for the correct eyetoy device
*/
package eyetoy

import (
	"fmt"

	"github.com/asciifaceman/hobocode"
	"github.com/google/gousb"
)

// EyeToy encapsulates known interactions with the Sony EyeToy
type EyeToy struct {
	Context *gousb.Context
	Device  *gousb.Device
	Config  *gousb.Config
}

// Close closes out the device connection and context
func (e *EyeToy) Close() {
	err := e.Config.Close()
	if err != nil {
		hobocode.Errorf("Error closing config: %v", err)
	}
	err = e.Device.Close()
	if err != nil {
		hobocode.Errorf("Error closing device: %v", err)
	}
	err = e.Context.Close()
	if err != nil {
		hobocode.Errorf("Error closing context: %v", err)
	}

}

// GetContext acquires a new gousb context and injects it into *EyeToy
func (e *EyeToy) GetContext() {
	ctx := gousb.NewContext()
	e.Context = ctx
}

/*
Open opens the device with our known VID and PID. It sets autodetach to prevent
kernel module interference, then acquires the only known working config for
this device and places both inside *EyeToy

	e := &EyeToy{}
	e.GetContext()
	err := e.Open()
	if err != nil {
		// handle error
	}
*/
func (e *EyeToy) Open() error {
	hobocode.HeaderLeft("Connection")
	hobocode.Info("Opening eyetoy device...")
	e.GetContext()
	dev, err := e.Context.OpenDeviceWithVIDPID(gousb.ID(SonyEyeToyVendorID), gousb.ID(SonyEyeToyProductID))
	if err != nil {
		return err
	}
	dev.SetAutoDetach(true)

	hobocode.Infof("Connection acquired with Eyetoy: %v", dev.Desc)
	e.Device = dev
	c, err := dev.Config(EyeToyPrimaryConfig)
	if err != nil {
		return err
	}
	e.Config = c
	return nil
}

/*
GetInterfaceEndpoint returns an interface of the given Alternate with a
done/Close function as well as the only endpoint available (0x81 1:IN)

	iface, done, endpoint, err := eyetoy.GetInterfaceEndpoint(1)
	if err != nil {
		// handle error
	}
	buf := make([]byte, endpoint.Desc.MaxPacketSize)
	readBytes, err := endpoint.Read(buf)
	done()
*/
func (e *EyeToy) GetInterfaceEndpoint(alt int) (*gousb.Interface, func(), *gousb.InEndpoint, error) {
	if alt >= 5 {
		return nil, nil, nil, fmt.Errorf("provided alt [%d] not in supported alternates: [0, 1, 2, 3, 4]", alt)
	}
	iface, err := e.Config.Interface(EyeToyPrimaryInterface, alt)
	if err != nil {
		return nil, nil, nil, err
	}

	ep, err := iface.InEndpoint(EyeToyPrimaryEndpoint)
	if err != nil {
		return nil, nil, nil, err
	}

	return iface, iface.Close, ep, nil
}

// ReadEndpoint reads one Packet from the given endpoint
func (e *EyeToy) ReadEndpoint(ep *gousb.InEndpoint) (int, []byte, error) {
	hobocode.HeaderLeft("Communicate")
	hobocode.Infof("Reading from endpoint [%s]", ep.String())
	buf := make([]byte, ep.Desc.MaxPacketSize)

	readBytes, err := ep.Read(buf)
	if err != nil {
		return 0, nil, err
	}

	return readBytes, buf, nil
}
