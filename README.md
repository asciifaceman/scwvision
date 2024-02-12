# SCWVision
This project is just a tinker-toy where I attempt to reimplement portions of the OV519 linux driver with only the relevant portions for the Sony EyeToy in order to play with it a bit.

You might ask yourself, "why?".

> why not?

You might ask yourself, "why not just use a webcam library"

> And you're right, that would probably work out of the box, but what would I learn?

First working prototype screenshot
![img](static/scwvision_prototype.png)

Developed and Tested for the `SLEH-00030` although all models should utilize the same base chip.

| Model      | Controller | Sensor               |
|------------|------------|----------------------|
| SLEH-00030 | OV519      | OV7648 (764X family) |

# Sony EyeToy
The Sony EyeToy was an interesting attempt by Sony to develop a peripheral for the PlayStation 2 to add new ways of interacting with games. It was used to process gestures and other inputs, and could be considered the ancestor of the PSVR/PSVR2 given the EyeToy beget the Move controllers beget the later cameras, tracking technology, etc which led to the present.

While largely forgotten, the EyeToy was a fascinating and important part of the development of modern PlayStation peripherals.

The sony eyetoy, for all intents and purposes, "is essentially just a normal USB 1.1 camera in a case that matches the PS2's design language" - PS Dev Wiki

The EyeToy was capable of 320x240 resolution images on the PS2, but can be pushed to 640x480 with custom drivers (one goal of this project).

### EyeToy Identification
The SonyEyetoy used the same generic internal hardware, although it was manufactured by several companies throughout its lifecycle.

| Model      | Manufacturer   | Notes    |
|------------|----------------|----------|
| SLEH-00030 | Logitech       | Original |
| SLEH-00031 | Namtai         |          |
| SCEH-0004  | Namtai/Chicony | Silver redesign |

The EyeToy vendor and product IDs are defined as

```go
const (
  SonyEyeToyVendorID  uint16 = 0x054c // 1356 Sony Corp.
  SonyEyeToyProductID uint16 = 0x0154 // 340 EyeToy Device
)
```

Although how an OS interprets these IDs may change, with `0x054c` occasionally appearing to be `Logitech` and `0x0154` alternating between `Eyetoy Audio Device` and `Logitech EyeToy USB Camera` or others, it's definitely a Sony VendorID (with many other devices under its ID).

### USB Configuration
The EyeToy appears to present a single valid configuration (`idx:1`), posessing only a single valid Interface (`idx: 0`), however that interface contains several alternates. The datasheet for the `OV519` does indicate that there should be other endpoints such as `IN:2 Audio/Midi 1.0 1x40 bytes, 1frame` however I have not completed my R&D in this space at the time of this writing and do not know more.

| Configuration | Interface | Alternate | Endpoint | Packet Size |
|---------------|-----------|-----------|----------|-------------|
| 1             | 0         | 0         | 0x81 1:IN| 0 bytes     |
| 1             | 0         | 1         | 0x81 1:IN| 384 bytes   |
| 1             | 0         | 2         | 0x81 1:IN| 512 bytes   |
| 1             | 0         | 3         | 0x81 1:IN| 768 bytes   |
| 1             | 0         | 4         | 0x81 1:IN| 896 bytes   |

Performing a buffered read on `1:0:0:IN` will always fail with a `packet size too large` error since it's packet size is 0.

However buffered reads on the other 4 endpoints will perform a roundtrip transaction, 

```
[ 0.133531] [0002fea7] libusb: debug [libusb_handle_events_timeout_completed] doing our own event handling
[ 0.133572] [0002fea6] libusb: debug [libusb_submit_transfer] transfer 0x7faab4001100
[ 0.133587] [0002fea6] libusb: debug [submit_iso_transfer] need 1 urbs for new transfer with length 3840
[ 0.133580] [0002fea7] libusb: debug [usbi_wait_for_events] poll() 3 fds with timeout in 100ms
[ 0.145376] [0002fea7] libusb: debug [usbi_wait_for_events] poll() returned 1
[ 0.145418] [0002fea7] libusb: debug [reap_for_handle] urb type=0 status=0 transferred=0
[ 0.145422] [0002fea7] libusb: debug [handle_iso_completion] handling completion status 0 of iso urb 1/1
[ 0.145424] [0002fea7] libusb: debug [handle_iso_completion] all URBs in transfer reaped --> complete!
[ 0.145428] [0002fea7] libusb: debug [usbi_handle_transfer_completion] transfer 0x7faab4001100 has callback 0x6eac30
[ 0.145624] [0002fea8] libusb: debug [libusb_free_transfer] transfer 0x7faab4001100
```

and if the OV519 controller is initialized with JPEG enabled will begin sending a JFIF file header set lacking image data (since the sensor is not yet initialized). There is little interaction with the camera to be had pre-controller-init.

# Development
There are a few things that need done to thinker against USB devices depending on the context. This project was only written for and intended to be used on a Linux system, particularly Debian based, and has not been tested by the maintainer elsewhere except in a WSL context.

### WSL>>Linux Dependencies
This has been tested to work on WSL, however you need to bind and then attach the USB device from windows.

```powershell
PS C:\WINDOWS\system32> usbipd list
Connected:
BUSID  VID:PID    DEVICE                                                        STATE
1-3    0489:e0c8  MediaTek Bluetooth Adapter                                    Not shared
1-4    0c45:6a10  Integrated Webcam                                             Not shared
2-2    054c:0154  Logitech EyeToy USB Camera                                    Not Shared
2-3    187c:0550  USB Input Device                                              Not shared
3-1    0d62:cabc  USB Input Device                                              Not shared

PS C:\WINDOWS\system32> usbipd bind --busid 2-2
PS C:\WINDOWS\system32> usbipd attach --wsl --busid 2-2
```

usbipd bind makes the device sharable, attach mounts it into the WSL instance.

Note: This requires an up to date usbipd.

### Linux
In order to more easily develop against USB devices, you can enable permissions on linux.

/etc/udev/rules.d/10-eyetoy.rules
```
SUBSYSTEM=="usb", ATTR{idVendor}=="054c", ATTR{idProduct}=="0154", MODE="0660", GROUP="plugdev"
```

then

```
sudo usermod -a -G plugdev $(whoami)
sudo udevadm control --reload-rules
sudo udevadm trigger
```

Then plug the device in (if it was plugged in already, plug-cycle it). If WSL disconnect, reconnect, then re-attach to WSL.

# References in no particular order
It's difficult to attribute every single piece of information and maintain any level of coherence in these documents, so here is a general listing of reference material used in assembling this information.

| Information Present  | Link |
|----------------------|------|
| Model manufacturers  | https://www.psdevwiki.com/ps2/EyeToy |
| Supported resolutions| https://en.wikipedia.org/wiki/EyeToy |
| Descriptor dump      | https://forums.pcsx2.net/Thread-Eyetoy-USB-Descriptors |

# Authors

| Author | Participation |
|--------|---------------|
| ASciifaceman | Owner / Maintainer |