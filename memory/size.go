package memory

type Size uint32

const (
	S1KIBI Size = 1024 << iota
	S2KIBI
	S4KIBI
	S8KIBI
	S16KIBI
	S32KIBI
	S64KIBI
	S128KIBI
	S256KIBI
	S512KIBI
	S1MIBI
	S2MIBI
	S4MIBI
	S8MIBI
	S16MIBI
	S32MIBI
	S64MIBI
	S128MIBI
	S256MIBI
	S512MIBI
	S1GIBI
	S2GIBI
	//S4GIBI
)

var availableSizes = map[Size]bool{
	S1KIBI:   true,
	S2KIBI:   true,
	S4KIBI:   true,
	S8KIBI:   true,
	S16KIBI:  true,
	S32KIBI:  true,
	S64KIBI:  true,
	S128KIBI: true,
	S256KIBI: true,
	S512KIBI: true,
	S1MIBI:   true,
	S2MIBI:   true,
	S4MIBI:   true,
	S8MIBI:   true,
	S16MIBI:  true,
	S32MIBI:  true,
	S64MIBI:  true,
	S128MIBI: true,
	S256MIBI: true,
	S512MIBI: true,
	S1GIBI:   true,
	S2GIBI:   true,
	//S4GIBI:   true,
}
