package user

// User struct represents a user with its authorizations
type User struct {
	ID    int            `json:"id"`
	Authz Authorizations `json:"authorization"`
}

const (
	// Authorization is the key for authorizations in the token structure
	Authorization = "authorization"
	// Read is the key for read authorization in the token structure
	Read = "read"
	// Insert is the key for insert authorization in the token structure
	Insert = "insert"
	// Update is the key for update authorization in the token structure
	Update = "update"
	// Delete is the key for delete authorization in the token structure
	Delete = "delete"
)

// Authorizations is a struct that contains all authorization for a user
type Authorizations struct {
	Read   bool `json:"read"`
	Update bool `json:"update"`
	Insert bool `json:"insert"`
	Delete bool `json:"delete"`
}
