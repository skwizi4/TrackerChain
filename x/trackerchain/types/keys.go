package types

const (
	// ModuleName defines the module name
	ModuleName = "trackerchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_trackerchain"
)

var (
	ParamsKey = []byte("p_trackerchain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
