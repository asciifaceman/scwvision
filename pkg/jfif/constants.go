package jfif

const (
	MARKER                    uint8 = 0xff
	MARKER_START_IMAGE        uint8 = 0xd8 //
	MARKER_APPLICATION_HEADER uint8 = 0xe0 // following 16 bits
	MARKER_QUANTABLE          uint8 = 0xdb //3 bits then data
	MARKER_START_FRAME        uint8 = 0xc0
	MARKET_HUFFMAN            uint8 = 0xc4
	MARKER_START_SCAN         uint8 = 0xda
	MARKER_EOI                uint8 = 0xd9
)
