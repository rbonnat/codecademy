# Codecademy assessment in Go : Cat pictures API

## Initialization
Before running `docker-compose` command, two steps are required:
* Create a common network for containers
    * `docker create network dev`
* Create image
    * `make dev-image`

## Running containers
Run `docker-compose up -d` to launch the three containers:
* codecademy - The API service
* localstack - AWS services (s3, secrets manager)
* mysql - The RDBMS database for metadata

## API description
The application can be reach using the following endpoints :
* Upload a new picture: `POST` http://localhost:8080/cat/picture
* Get an uploaded picture: `GET` http://localhost:8080/cat/picture/{id}
* Update a picture : `PUT` http://localhost:8080/cat/picture/{id}
* Delete a picture : `DELETE` http://localhost:8080/cat/picture/{id}
* List all pictures : `GET` http://localhost:8080/cat/pictures

The complete documentation can be found [here](https://github.com/rbonnat/codecademy/blob/master/apidoc.md)

The API uses a JWT Bearer Token to authenticate and authorize access

Example :
`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIzNDU2Nzg5MCwiYXV0aG9yaXphdGlvbiI6eyJyZWFkIjp0cnVlLCJ1cGRhdGUiOnRydWUsImluc2VydCI6dHJ1ZSwiZGVsZXRlIjp0cnVlfX0.CiDOe4g7toUvAR72H8gQRU70SdfE0xCGq7t-_41nl4s`

The token is signed using HS256 symmetric encryption algorithm. The secret key `secret` and is stored in secrets manager.

Example of payload :
```json
{
  "id": 1234567890,
  "authorization": {
    "read": true,
    "update": true,
    "insert": true,
    "delete": true
  }
}
```

## Make commands
* Check data
    * `make s3` - List all files in the bucket
    * `make db` - List of pictures in DB
* Development
    * `make tests` - launch unit tests (require `go test`)
    * `make test-coverage` - build HTML for test coverage (require `go tool cover`)
    * `make lint` - launch code linter (require `golangci-lint`)
