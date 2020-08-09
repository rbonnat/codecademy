# API Documentation

## Upload a picture
Upload a picture

```
POST http://localhost:8080/cat/picture
```

### Parameters
`ID (string) - ID of the picture`

additionalMetadata

`name (string) - Name of the picture`

`picture (file) - file to upload`

### Response
`400 Bad request`

`200  Successful`

Example :
```json
{
  "id": "98f79fe3-75c4-4dbe-b0e6-b985c987bf85"
}
```

## Get a picture
Get a picture by its ID

```
GET http://localhost:8080/cat/picture/{id}
```

### Parameters
`ID (string) - ID of the picture`

### Response
`200  Successful`

```json
{
  "meta_data": {
    "ID": "98f79fe3-75c4-4dbe-b0e6-b985c987bf85",
    "Name": "My cat",
    "FileName": "cat.png",
    "ContentType": "image/png",
    "Size": 3945646
  },
  "content": "iVBORw0KGgoAAAANSUhEUgAACjwAAAXECAYAAABzqjUlAAABYWlDQ1BrQ0dDb2xvclNwYWNlRGlzcGxheVAzAAAokWN..."
}
```

### Example
```
curl --location --request GET 'http://localhost:8080/cat/picture/98f79fe3-75c4-4dbe-b0e6-b985c987bf85' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIzNDU2Nzg5MCwiYXV0aG9yaXphdGlvbiI6eyJyZWFkIjp0cnVlLCJ1cGRhdGUiOnRydWUsImluc2VydCI6dHJ1ZSwiZGVsZXRlIjp0cnVlfX0.CiDOe4g7toUvAR72H8gQRU70SdfE0xCGq7t-_41nl4s'
```

## Update a picture
Update a picture by its ID

```
PUT http://localhost:8080/cat/picture/{id}
```

### Parameters
`ID (string) - ID of the picture to delete`

additionalMetadata

`name (string) - Name of the picture`

`picture (file) - file to upload`

### Response
`404 Picture not found`

`200  Successful`

### Example
```
curl --location --request PUT 'http://localhost:8080/cat/picture/98f79fe3-75c4-4dbe-b0e6-b985c987bf85' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIzNDU2Nzg5MCwiYXV0aG9yaXphdGlvbiI6eyJyZWFkIjp0cnVlLCJ1cGRhdGUiOnRydWUsImluc2VydCI6dHJ1ZSwiZGVsZXRlIjp0cnVlfX0.CiDOe4g7toUvAR72H8gQRU70SdfE0xCGq7t-_41nl4s' \
--form 'name=My cat Picture updated' \
--form 'picture=@documents/new-cat-pic.png'
```

## Delete a picture
Delete a picture of by its ID
```
DELETE http://localhost:8080/cat/picture/{id}
```

### Parameters
`ID (string) - ID of the picture to delete`

### Response
`404 Picture not found`

`200  Successful`

### Example
```
curl --location --request DELETE 'http://localhost:8080/cat/picture/98f79fe3-75c4-4dbe-b0e6-b985c987bf85' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIzNDU2Nzg5MCwiYXV0aG9yaXphdGlvbiI6eyJyZWFkIjp0cnVlLCJ1cGRhdGUiOnRydWUsImluc2VydCI6dHJ1ZSwiZGVsZXRlIjp0cnVlfX0.CiDOe4g7toUvAR72H8gQRU70SdfE0xCGq7t-_41nl4s'
```

## List all pictures
Fetch the list of all uploaded pictures

```
GET http://localhost:8080/cat/pictures
```

### Parameters


### Response

`200 - Successful`

Example:
```json
{
  "pictures": [
    {
      "ID": "98f79fe3-75c4-4dbe-b0e6-b985c987bf85",
      "Name": "My Cat",
      "FileName": "cat.png",
      "ContentType": "image/png",
      "Size": 3945646
    }
  ]
}
```

### Example
```
curl --location --request GET 'http://localhost:8080/cat/pictures' \
 --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIzNDU2Nzg5MCwiYXV0aG9yaXphdGlvbiI6eyJyZWFkIjp0cnVlLCJ1cGRhdGUiOnRydWUsImluc2VydCI6dHJ1ZSwiZGVsZXRlIjp0cnVlfX0.CiDOe4g7toUvAR72H8gQRU70SdfE0xCGq7t-_41nl4s
```

