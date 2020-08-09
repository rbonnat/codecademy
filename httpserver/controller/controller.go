package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"

	"github.com/rbonnat/codecademy/picture"
)

const (
	maxMBMemory = 32
)

// HandleGetPic returns a handler for the endpoint to get a picture
// returns a picture by its ID
func HandleGetPic(fs FileStore, dbs DBStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling Get picture request")

		// Extract ID
		picID := chi.URLParam(r, "ID")
		log.Printf("Picture ID: %s", picID)

		// Fetch metadata in DB
		pic, err := dbs.Get(picID)
		if err != nil {
			log.Printf("Error while getting picture metadata in DB: '%v'", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Not found
		if pic == nil {
			log.Printf("Picture with ID '%s' not found\n", picID)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// Fetch picture in file storage
		content, err := fs.Get(picID)
		if err != nil {
			log.Printf("Error while getting picture in file storage: '%v'", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Prepare response
		gr := GetResponse{
			MetaData: *pic,
			Content:  content,
		}
		resp, err := json.Marshal(gr)
		if err != nil {
			log.Printf("error when marshalling response '%+v' : '%v'", gr, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Success: Write response for return
		w.Write(resp)
	}
}

// HandleDeletePic returns a handler for endpoint to delete a picture
// delete a pictures of a specific ID
func HandleDeletePic(fs FileStore, dbs DBStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling delete picture request")

		// Extract ID
		picID := chi.URLParam(r, "ID")
		log.Printf("Picture ID: %s", picID)

		// Delete pic metadata in DB
		n, errDB := dbs.Delete(picID)
		if errDB != nil {
			log.Printf("Error deleting picture in database: '%v'", errDB)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if n == 0 {
			log.Printf("Picture with ID `%s` not found", picID)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// Delete pic in file storage
		errFs := fs.Delete(picID)
		if errFs != nil {
			log.Printf("Error deleting picture in file storage: '%v'", errFs)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
}

// HandleInsertPic returns a handler for endpoint to insert a picture
// insert an uploaded a picture
func HandleInsertPic(fs FileStore, dbs DBStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling Insert picture request")

		// Get post data
		errFormData := r.ParseMultipartForm(maxMBMemory << 20) // maxMemory 32MB
		if errFormData != nil {
			log.Printf("Error getting data from multipart form: '%v'", errFormData)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Get Name
		name := r.FormValue("name")

		// Get file and headers
		f, headers, fileErr := r.FormFile("picture")
		if fileErr != nil {
			log.Printf("Error getting picture from multipart form: '%v'", fileErr)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		defer f.Close()

		// Read file to buffer
		fileBuffer := make([]byte, headers.Size)
		f.Read(fileBuffer)

		// Define picture
		pic := &picture.Picture{
			ID:          uuid.NewV4().String(), // Generate a new UUID
			Name:        name,
			FileName:    headers.Filename,
			ContentType: headers.Header.Get("Content-Type"),
			Size:        headers.Size,
		}

		// Insert into DB
		errDB := dbs.Insert(pic)
		if errDB != nil {
			log.Printf("Error when inserting data in database: '%v'", errDB)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Insert into filesystem
		_, errFs := fs.Insert(fileBuffer, pic.ID)
		if errFs != nil {
			log.Printf("Error when writing file to file storage: '%v'", errFs)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Prepare response
		ir := InsertResponse{ID: pic.ID}
		resp, errMarshall := json.Marshal(ir)
		if errMarshall != nil {
			log.Printf("error when marshalling response: '%v'", errMarshall)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.Write(resp)
	}
}

// HandleUpdatePic returns a handler for endpoint to update a pic
// Update a pic
func HandleUpdatePic(fs FileStore, dbs DBStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling Update pic request")

		// Extract ID
		picID := chi.URLParam(r, "ID")
		log.Printf("Picture ID: %s", picID)

		// Get post data
		errFormData := r.ParseMultipartForm(maxMBMemory << 20) // maxMemory 32MB
		if errFormData != nil {
			log.Printf("Error getting data from multipart form: '%v'", errFormData)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		// Get Name
		name := r.FormValue("name")

		// Get file and headers
		f, headers, fileErr := r.FormFile("picture")
		if fileErr != nil {
			log.Printf("Error getting picture from multipart form: '%v'", fileErr)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		defer f.Close()

		fileBuffer := make([]byte, headers.Size)
		f.Read(fileBuffer)

		// Update into DB
		pic := &picture.Picture{
			ID:          picID,
			Name:        name,
			FileName:    headers.Filename,
			ContentType: headers.Header.Get("Content-Type"),
			Size:        headers.Size,
		}

		// Update into DB
		n, errDB := dbs.Update(pic)
		if errDB != nil {
			log.Printf("Error when updating data in database: '%v'", errDB)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if n == 0 {
			log.Printf("Picture with ID `%s` not found", picID)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// Update into filesystem
		errFs := fs.Update(fileBuffer, picID)
		if errFs != nil {
			log.Printf("Error when updating picture in file storage: '%v'", errFs)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Success
	}
}

// HandleGetPics returns a handler for endpoint to get a list of all pics
// Return a list of all pics
func HandleGetPics(fs FileStore, dbs DBStore) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling list pics request")

		// Fetch metadata in DB
		pics, errDB := dbs.GetAll()
		if errDB != nil {
			log.Printf("Error while getting all pictures: '%v'", errDB)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Prepare response
		gr := GetAllResponse{
			Pictures: pics,
		}
		resp, err := json.Marshal(gr)
		if err != nil {
			log.Printf("Error when marshalling response '%+v' : '%v'", gr, err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// Success
		w.Write(resp)
	}
}
