# Device Registers
The Sony EyeToy uses an OV519 Chip.

These are the device registers

| Address | Name   | R/W | Bit | Purpose           | Default |
|---------|--------|-----|-----|-------------------|---------|
| 0x50    | RESET0 |  RW |     |                   | 00      |
|         |        |     | 0   | UDC Reset         | 0       |
|         |        |     | 1   | UDCIF Reset       | 0       |
|         |        |     | 2   | uController Reset | 0       |
|         |        |     | 3   | Audio Reset       | 0       |
|         |        |     | 4   | SCCB Reset        | 0       |
|         |        |     | 5   | Register Reset    | 0       |
|         |        |     | 6   | Snapshot Reset    | 0       |
|         |        |     | 7   | GPIO Reset        | 0       |
| 0x51    | RESET1 | RW  |     |                   | 00      |
|         |        |     | 0   | CIF Reset         | 0       |
|         |        |     | 1   | SFIFO Reset       | 0       |
|         |        |     | 2   | JPEG Reset        | 0       |
|         |        |     | 3   | Video FIFO Reset  | 0       |
| 0x53    |EN_CLK0 | RW  |     |                   | 87      |
|         |        |     | 0   | UDC Enable        | 1       |
|         |        |     | 1   | UDCIF Enable      | 0       |
|         |        |     | 2   | Microcontroller Enable | 0  |
|         |        |     | 3   | Audio Enable      | 0       |
|         |        |     | 4   | SCCB              | 0       |
|         |        |     | 5   | RESERVED          | 1       |
|         |        |     | 6   | ISP Clock Enable  | 1       |
|         |        |     | 7   | Transceiver Enable| 1       |
| 0x54    | EN_CLK1| RW  |     |                   | 00      |
|         |        |     | 1   | SFIFO Enable      | 0       |
|         |        |     | 2   | JPEG Enable       | 0       |
|         |        |     | 3   | Video FIFO Enable | 0       |
| 0x55    | AUDIO_CLK| RW|     |                   | 01      |
|         |        |     | 1:0 | Clock Select      | 0       | 
|         |        |     | 1:0:0 | 2.048 MHz       |         |
|         |        |     | 1:0:1 | 2.048 MHz       |         |
|         |        |     | 1:0:2 | 4.096 MHz       |         |
|         |        |     | 1:0:3 | 6.144 MHz       |         |
|         |        |     | 2   | SD_CLK Divide by 2 Enable |0|
|         |        |     | 3   | FIR_BLK Divide by 2 Enable|0|
|         |        |     | 4   | Fixed phase clock Enable | 0|
|         |        |     | 5   | 24 MHz Clock Enable      | 0|
| 0x57    | SNAPSHOT| RW |     |                   | 01      |
|         |        |     | 0   | Snapshot Stat Pin |         |
|         |        |     | 1   | Snapshot Stat Debounce |    |
|         |        |     | 2   | Snapshot Wakeup Enable |    |
|         |        |     | 3   | Host Snapshot     |         |
|         |        |     | 4   | Snapshot Clear    |         |
|         |        |     | 5   | Snapshot Enable   |         |
| 0x58    | PONOFF | RW  |     | Power On/Off      | 00      |
|         |        |     | 0   | PONOFF Stat Pin   |         |
|         |        |     | 1   | PONOFF Stat Debnc |         |
|         |        |     | 2   | PONOFF Wake Enable|         |
|         |        |     | 3   | PONOFF Status Clr |         |
|         |        |     | 4   | PONOFF Enable     |         |
| 0x59    | CAMERA_CLOCK|RW|   | Camera Clock Divider| 02    |
|         |        |     | 4:0 | CCLK=48MHz / Camera Clock[4:0]| |
|         |        |     |     | 00000 = 48MHz | |
| 0x5A    | YS_CTRL| RW  |     |                    | 6C     |
|         |           |     | 0 | System 1 1011 00 Enable | 0      |
|         |           |     | 1 | System Reset Mask Enable| 1      | 
|         |           |     | 2 | Oscillator Power-Down SUSPEND Enable - microcontroller | 1 |
|         |           |     | 3 | Oscillator Power-Down SUSPEND Enable - USB | 0 |
|         |           |     | 4 | Wakeup Enable | 1 |
|         |           |     | 5 | USB Reset Enable | 1 |
|         |           |     | 6 | Power Down SUSPEND Mode Enable | 0 |
|         |           |     | 7 | UNUSED | 0 |