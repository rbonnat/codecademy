package mysqlstore

import (
	"database/sql"
	"log"

	// Imported for side effect
	_ "github.com/go-sql-driver/mysql"
	"github.com/rbonnat/codecademy/picture"
)

// MySQLStore is a structure containing MySQL client
type MySQLStore struct {
	client *sql.DB
}

// New returns a pointer to a MySQLStore instance
func New(DSN string) (*MySQLStore, error) {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Printf("Error while opening MySQL connection with DSN '%s'", DSN)
		return nil, err
	}

	return &MySQLStore{client: db}, nil
}

// Get data of a pic
func (ms *MySQLStore) Get(ID string) (*picture.Picture, error) {
	log.Printf("Getting picture with ID '%s'", ID)

	picture := picture.Picture{}

	query := "SELECT id, name, filename, content_type, size FROM pictures WHERE id = ?"

	err := ms.client.QueryRow(query, ID).Scan(
		&picture.ID,
		&picture.Name,
		&picture.FileName,
		&picture.ContentType,
		&picture.Size,
	)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}

	return &picture, nil
}

// Delete data of a pic
func (ms *MySQLStore) Delete(ID string) (int, error) {
	log.Printf("Deleting picture with ID '%s'", ID)

	var err error
	stm, err := ms.client.Prepare("DELETE FROM pictures WHERE id = ?")
	if err != nil {
		return 0, err
	}

	r, err := stm.Exec(ID)
	if err != nil {
		return 0, err
	}

	n, err := r.RowsAffected()

	return int(n), err
}

// Update data of a pic
func (ms *MySQLStore) Update(p *picture.Picture) (int, error) {
	log.Printf("Updating picture : '%+v'", *p)

	var err error
	stm, err := ms.client.Prepare("UPDATE pictures SET name = ?, filename = ?, content_type = ?, size = ? WHERE id = ?")
	if err != nil {
		return 0, err
	}

	r, err := stm.Exec(p.Name, p.FileName, p.ContentType, p.Size, p.ID)
	if err != nil {
		return 0, err
	}

	n, err := r.RowsAffected()

	return int(n), err
}

// Insert data of a pic
func (ms *MySQLStore) Insert(p *picture.Picture) error {
	log.Printf("Inserting picture with name '%s'", p.Name)

	var err error

	stm, err := ms.client.Prepare("INSERT pictures (id, name, filename, content_type, size) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stm.Exec(p.ID, p.Name, p.FileName, p.ContentType, p.Size)
	if err != nil {
		return err
	}

	return nil
}

// GetAll return list of pics
func (ms *MySQLStore) GetAll() ([]picture.Picture, error) {
	log.Println("Getting all pictures")

	pictures := []picture.Picture{}

	query := "SELECT id, name, filename, content_type, size FROM pictures LIMIT 10"

	rows, errQuery := ms.client.Query(query)
	if errQuery != nil {
		return nil, errQuery
	}
	defer rows.Close()

	for rows.Next() {
		pic := &picture.Picture{}
		err := rows.Scan(
			&pic.ID,
			&pic.Name,
			&pic.FileName,
			&pic.ContentType,
			&pic.Size,
		)
		if err != nil {
			log.Printf(" \n")
		}
		pictures = append(pictures, *pic)
	}

	errNext := rows.Err()
	if errNext != nil {
		return nil, errNext
	}

	return pictures, nil
}
