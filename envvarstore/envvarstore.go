package envvarstore

import "os"

// EnvStore conforms to VarStore interface to get environment variables
type EnvStore struct {
}

// New return a instance of EnvStore
func New() *EnvStore {
	return &EnvStore{}
}

// Get returns the value of an environment variable
func (e *EnvStore) Get(key string) string {
	return os.Getenv(key)
}
