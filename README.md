# Simple in memory key value cache

### Notes
* Caches key value pairs in memory
* Will overwrite existing values
* Entries will expire after a set duration, default `30 minutes`
* Expired items will be cleaned after a set interval, default `1 hour`
* Logging to file `logs.txt`

### Build and run

#### Development

`go env -w GO111MODULE=on`

`go run main.go`

#### Docker

`docker build --tag key-value-cache .`

`docker run --publish 8080:8080 key-value-cache`

#### Tests

`go test ./cache/...`

### Usage

The application will accept these http requests to port `8080`

#### GET /{key}

Response status:`200 OK`, `404 Not found`

Example: `curl -v "localhost:8080/1"`


#### POST /{key}

Response status: `200 OK`, `400 Bad request`

Example: `curl -v "localhost:8080/1" -X POST -d 'data'` 




