package eyetoy

// Instruction is a register:value pair
type Instruction struct {
	Desc string
	Reg  uint16
	Val  uint8
}

/*
0v519_controller_init defines registers and values to set in a sequence
that will initialize the sony eyetoy

these are largely lifted from the OV519 linux kernel driver which was never
really completed for this chipset and served multiple chips so not every
instruction may be important, accurate, or valid
*/
var ov519_controller_init []*Instruction = []*Instruction{
	{"Enable System", OV519_YS_CTRL, 0x6d},
}
