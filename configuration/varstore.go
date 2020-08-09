package configuration

// VarStore is an interface to a key value store
type VarStore interface {
	Get(key string) string
}
