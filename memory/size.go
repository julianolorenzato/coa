package memory

type Size uint64

const (
	S1KILO Size = 1024 << iota
	S2KILO
	S4KILO
	S8KILO
	S16KILO
	S32KILO
	S64KILO
	S128KILO
	S256KILO
	S512KILO
	S1MEGA
	S2MEGA
	S4MEGA
	S8MEGA
	S16MEGA
	S32MEGA
	S64MEGA
	S128MEGA
	S256MEGA
	S512MEGA
	S1GIGA
	S2GIGA
	S4GIGA
	S8GIGA
	S16GIGA
	S32GIGA
	S64GIGA
	S128GIGA
	S256GIGA
	S512GIGA
)

type Space interface {
	[S1KILO]byte |
		[S2KILO]byte |
		[S4KILO]byte |
		[S8KILO]byte |
		[S16KILO]byte |
		[S32KILO]byte |
		[S64KILO]byte |
		[S128KILO]byte |
		[S256KILO]byte |
		[S512KILO]byte |
		[S1MEGA]byte |
		[S2MEGA]byte |
		[S4MEGA]byte |
		[S8MEGA]byte |
		[S16MEGA]byte |
		[S32MEGA]byte |
		[S64MEGA]byte |
		[S128MEGA]byte |
		[S256MEGA]byte |
		[S512MEGA]byte |
		[S1GIGA]byte |
		[S2GIGA]byte |
		[S4GIGA]byte |
		[S8GIGA]byte |
		[S16GIGA]byte |
		[S32GIGA]byte |
		[S64GIGA]byte |
		[S128GIGA]byte |
		[S256GIGA]byte |
		[S512GIGA]byte
}
