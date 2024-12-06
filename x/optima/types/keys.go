package types

const (
	// ModuleName defines the module name
	ModuleName = "optima"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_optima"
)

var (
	ParamsKey = []byte("p_optima")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
