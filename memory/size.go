package memory

const (
	_           = iota
	KILO uint64 = 1 << (10 * iota)
	MEGA
	GIGA
)

type storageSizes interface {
	[512 * GIGA]byte |
		[256 * GIGA]byte |
		[128 * GIGA]byte |
		[64 * GIGA]byte |
		[32 * GIGA]byte |
		[16 * GIGA]byte |
		[8 * GIGA]byte |
		[4 * GIGA]byte |
		[2 * GIGA]byte |
		[1 * GIGA]byte |
		[512 * MEGA]byte |
		[256 * MEGA]byte |
		[128 * MEGA]byte |
		[64 * MEGA]byte |
		[32 * MEGA]byte |
		[16 * MEGA]byte |
		[8 * MEGA]byte |
		[4 * MEGA]byte |
		[2 * MEGA]byte |
		[1 * MEGA]byte |
		[512 * KILO]byte |
		[256 * KILO]byte |
		[128 * KILO]byte |
		[64 * KILO]byte |
		[32 * KILO]byte |
		[16 * KILO]byte |
		[8 * KILO]byte |
		[4 * KILO]byte |
		[2 * KILO]byte |
		[1 * KILO]byte
}
