# SCWVision
A tinker tool playing with the Sony Eyetoy

# Development
In order to more easily develop against USB devices, you can enable permissions on linux.

/etc/udev/rules.d/10-eyetoy.rules
```
SUBSYSTEM=="usb", ATTR{idVendor}=="0x054c", ATTR{idProduct}=="0x0154", MODE="0660", GROUP="plugdev"
```

then

```
sudo udevadm control --reload-rules
```

Then plug the device in (if it was plugged in already, plug-cycle it)