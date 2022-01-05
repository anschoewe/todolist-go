# Todo List: Gin + Gorm + Bootstrap example
This is a basic web application that shows how to use Gin/Gonic and GORM.  It attempts to show some common use cases and how to test for them.  Right now this includes a web interface and an API.

It uses an in-memory SqlLite database that is seeded with 3 records.  All data is lost when the applciation stops

# Test
Execute all unit tests
```bash
go test ./...
```

Execute all unit tests with code coverage
```bash
go test -v ./... -coverprofile cover.out
go tool cover -html=cover.out -o cover.html
open cover.html
```

# Run
Start the web application on local port 8080
```bash
go run main.go
```

# Future enhancements
- Docker and Skaffold build pipeline (and Github actions)
- Dependabot for auto patching
- Support for POST, PATCH, DELETE methods
- Authentication of web app and APIs
- More unit testing
- 
