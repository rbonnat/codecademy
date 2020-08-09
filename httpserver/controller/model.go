package controller

import (
	"github.com/rbonnat/codecademy/picture"
)

// DBStore is an interface for storing files data
type DBStore interface {
	Get(string) (*picture.Picture, error)
	Delete(string) (int, error)
	Update(*picture.Picture) (int, error)
	Insert(*picture.Picture) error
	GetAll() ([]picture.Picture, error)
}

// FileStore is an interface for storing files
type FileStore interface {
	Get(string) ([]byte, error)
	Delete(string) error
	Update([]byte, string) error
	Insert([]byte, string) (int, error)
}

// InsertResponse is a struct for the response of insert pic endpoint
type InsertResponse struct {
	ID string `json:"id"`
}

// GetResponse is a struct for the response of get pic endpoint
type GetResponse struct {
	MetaData picture.Picture `json:"meta_data"`
	Content  []byte          `json:"content"`
}

// GetAllResponse is a struct for the response of get list of pics endpoint
type GetAllResponse struct {
	Pictures []picture.Picture `json:"pictures"`
}
