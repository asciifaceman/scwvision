package eyetoy

import (
	"fmt"

	"github.com/google/gousb"
	"go.uber.org/zap"
)

/*
GUSB wraps the gousb stuff in one place
*/
type GUSB struct {
	logger  *zap.SugaredLogger
	Context *gousb.Context
	Device  *gousb.Device
	Config  *gousb.Config
}

// NewGUSB returns a *GUSB connected to the target device
func NewGUSB(l *zap.SugaredLogger) *GUSB {
	g := &GUSB{
		logger: l.Named("gousb"),
	}
	g.GetContext()

	return g
}

// GetContext acquires a gousb context and adds it to the struct
func (g *GUSB) GetContext() {
	ctx := gousb.NewContext()
	g.Context = ctx
}

/*
Open acquires a connection to the Sony EyeToy device
this will fail on any other device and may not work on
all Eyetoy models

this can only handle one connected device at the moment
*/
func (g *GUSB) Open() error {
	g.logger.Debug("attempting to open connection to sony eyetoy device")

	devs, err := g.Context.OpenDevices(findEyetoy())
	if err != nil {
		return err
	}

	if len(devs) > 1 {
		g.logger.Warnw("found more than one eyetoy device, defaulting to the first found",
			"found", len(devs),
			"using", devs[0].Desc.String(),
		)
	}

	g.Device = devs[0]

	for index, device := range devs {
		// close unused device connections
		if index == 0 {
			continue
		}
		device.Close()
	}

	g.logger.Debugf("acquiring primary config from eyetoy device",
		"device", g.Device.Desc.String(),
		"config", EyeToyPrimaryConfig,
	)

	c, err := g.Device.Config(EyeToyPrimaryConfig)
	if err != nil {
		return err
	}
	g.Config = c

	return nil
}

// Close closes everything down in the correct order
func (g *GUSB) Close() {
	g.logger.Debug("Closing down connections...")
	err := g.Config.Close()
	if err != nil {
		g.logger.Errorw("failed to close config", "error", err)
	}

	err = g.Context.Close()
	if err != nil {
		g.logger.Errorw("failed to close context", "error", err)
	}

	err = g.Device.Close()
	if err != nil {
		g.logger.Errorw("failed to close device", "error", err)
	}
}

/*
GetInterfaceEndpoint acquires an interface:alt with the primary known
interface and the only known primary IN endpoint for the given interface:alt
combination

this may change as I now believe there are more endpoints based on the
datasheet but more discovery is needed

returned function is a done closer

interface should be closed when operation is done
*/
func (g *GUSB) GetInterfaceEndpoint(alt int) (*gousb.Interface, func(), *gousb.InEndpoint, error) {
	if alt >= 5 {
		return nil, nil, nil, fmt.Errorf("provided alt [%d] not in supported alternates. supported: [0, 1, 2, 3, 4]", alt)
	}

	iface, err := g.Config.Interface(EyeToyPrimaryInterface, alt)
	if err != nil {
		return nil, nil, nil, err
	}

	ep, err := iface.InEndpoint(EyeToyPrimaryEndpoint)
	if err != nil {
		return nil, nil, nil, err
	}

	return iface, iface.Close, ep, nil
}
