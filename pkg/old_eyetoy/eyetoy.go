/*
Scan USB devices for the correct eyetoy device
*/
package old_eyetoy

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/asciifaceman/hobocode"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/gousb"
)

/*
New returns a new *EyeToy configured for signal-based interrupts via
*EyeToy.exeunt

It acquires the context. opens the connection, configures the signal notification
then returns the struct
*/
func New(timeout time.Duration) (*EyeToy, error) {
	e := &EyeToy{}
	e.exeunt = make(chan os.Signal, 1)
	e.GetContext()
	err := e.Open()
	if err != nil {
		return nil, fmt.Errorf("error opening: %v", err)
	}
	err = e.InitializeController()
	if err != nil {
		return nil, fmt.Errorf("error initializing the device: %v\nYou should consider plug-cycling the device now", err)
	}
	signal.Notify(e.exeunt, syscall.SIGINT, syscall.SIGTERM)
	return e, nil
}

// EyeToy encapsulates known interactions with the Sony EyeToy
type EyeToy struct {
	Context *gousb.Context
	Device  *gousb.Device
	Config  *gousb.Config
	exeunt  chan os.Signal
}

// GetContext acquires a new gousb context and injects it into *EyeToy
func (e *EyeToy) GetContext() {
	ctx := gousb.NewContext()
	e.Context = ctx
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
	hobocode.Debug("Opening eyetoy device...")
	e.GetContext()
	dev, err := e.Context.OpenDeviceWithVIDPID(gousb.ID(SonyEyeToyVendorID), gousb.ID(SonyEyeToyProductID))
	if err != nil {
		return err
	}
	dev.SetAutoDetach(true)

	hobocode.Debugf("Connection acquired with Eyetoy: %v", dev.Desc)
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

TODO: Find audio endpoint (IN:2)

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

/*
GetImage attempts to return a still image from the EyeToy camera
*/
func (e *EyeToy) GetImage() error {
	hobocode.Debugf("Reading image from interface endpoint...")

	_, done, ep, err := e.GetInterfaceEndpoint(4)
	if err != nil {
		return err
	}
	defer done()

	var buf []byte

	ctx, done := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer done()
	stream, err := ep.NewStream(ep.Desc.MaxPacketSize, 1)
	if err != nil {
		return err
	}
	var lastByte byte = 0x00

OUTER:
	for {
		select {
		case <-ctx.Done():
			hobocode.Warn("Timing out")
			break OUTER
		case <-e.exeunt:
			hobocode.Warn("Received signal, exiting")
			break OUTER
		default:
			packet := make([]byte, ep.Desc.MaxPacketSize)
			read, err := stream.Read(packet)

			if err != nil || read < 0 {
				hobocode.Errorf("Failed to read: %v", err)
				continue
			}

			if read > 0 {
				buf = append(buf, packet...)
			}

			for i, bit := range packet {
				if lastByte == 0xff {
					if bit == 0xd9 {
						break OUTER
					}
				}
				if bit == 0xff && i+1 < len(packet) {

					if packet[i+1] == 0xd9 {
						hobocode.Successf("Found EOF!")
						break OUTER
					}
				}
			}
			lastByte = packet[len(packet)-1]
		}
	}

	spew.Dump(buf)

	/*
		readBytes, err := ep.Read(buf)
		if err != nil {
			return err
		}
		if readBytes > 0 {
			spew.Dump(buf)

			image, err := jpeg.Decode(bytes.NewReader(buf))
			if err != nil {
				return err
			}
			spew.Dump(image)
		}
	*/
	return nil
}
