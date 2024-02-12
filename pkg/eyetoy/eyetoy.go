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
)

/*
EyeToy is the driver interface behind the eyetoy interactions
*/
type EyeToy struct {
	// SIGTERM or SIGINT signals
	term chan os.Signal
}

func New() (*EyeToy, error) {
	e := &EyeToy{}

	e.term = make(chan os.Signal, 1)
	signal.Notify(e.term, syscall.SIGINT, syscall.SIGTERM)

	return e, nil
}
