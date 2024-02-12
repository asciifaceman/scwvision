/*
The EyeToy package is a pseudo-lift of significant portions of the OV519 driver
from the linux kernel but only the relevant bits for the specific OV519 controller
and OV7648 sensor combination used in the Sony EyeToy. As of this writing this
was written for the SLEH-00030 model - other models may have different hardware
but I did not have one to test.
*/
package eyetoy

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

/*
EyeToy is the driver interface behind the eyetoy interactions
*/
type EyeToy struct {
	logger *zap.SugaredLogger

	// SIGTERM or SIGINT signals
	term chan os.Signal

	GUSB *GUSB
}

/*
New returns an *EyeToy configured for signal-based interrupts

It acquires a gousb context, opens the device connection, configures
the signal notifys and returns
*/
func New() (*EyeToy, error) {
	l, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	e := &EyeToy{
		logger: l.Named("eyetoy").Sugar(),
	}

	e.term = make(chan os.Signal, 1)
	signal.Notify(e.term, syscall.SIGINT, syscall.SIGTERM)

	return e, nil
}

/*
Eyetoy test sequence used by the test entrypoint
*/
func (e *EyeToy) Test(blinks int) error {
	g := NewGUSB(e.logger)
	err := g.Open()
	if err != nil {
		return err
	}
	return nil
}
