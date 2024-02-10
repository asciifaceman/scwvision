/*
Scan USB devices for the correct eyetoy device
*/
package eyetoy

import (
	"fmt"

	"github.com/asciifaceman/hobocode"
	"github.com/davecgh/go-spew/spew"
	"github.com/google/gousb"
)

func Scan() error {
	hobocode.Info("Opening eyetoy device...")
	ctx := gousb.NewContext()
	//ctx.Debug(4)
	defer ctx.Close()

	dev, err := ctx.OpenDeviceWithVIDPID(gousb.ID(SonyEyeToyVendorID), gousb.ID(SonyEyeToyProductID))
	if err != nil {
		return err
	}
	defer dev.Close()

	dev.SetAutoDetach(true)

	spew.Dump(dev)
	spew.Dump(dev.Desc)

	hobocode.Info("Endpoint Description")

	intf, done, err := dev.DefaultInterface()
	defer done()
	if err != nil {
		return err
	}

	derp, _ := dev.Config(1)
	derp2, err := derp.Interface(0, 4)
	if err != nil {
		hobocode.Errorf("Derp err: %v", err)
	}
	defer derp2.Close()
	fmt.Println("==")
	fmt.Println(derp2.String())
	fmt.Println("==")

	fmt.Println(intf.Setting)
	ep, err := derp2.InEndpoint(1)
	if err != nil {
		return err
	}

	fmt.Println(ep.String())

	hobocode.Infof("Attempting test read - max size %d...", ep.Desc.MaxPacketSize)

	ctx.Debug(4)
	buf := make([]byte, ep.Desc.MaxPacketSize)
	readBytes, err := ep.Read(buf)
	ctx.Debug(3)
	if err != nil {
		return err
	}
	fmt.Printf("Read %d bytes", readBytes)

	return nil

	total := 0
	for i := 0; i < 10; i++ {
		readBytes, err := ep.Read(buf)
		if err != nil {
			fmt.Printf("Read [iter:%d] returned an err: %v\n", i, err)
		}
		if readBytes == 0 {
			fmt.Printf("Read [iter:%d] no data\n", i)
		}
		total += readBytes
	}
	fmt.Printf("Read %d bytes...\n", total)

	return nil
}
