# {{.Name}}

{{.Description}}

## Installation

```bash
go get -v {{.Package}}
```

## Usage

Navigate into the project directory and run the following command

```bash
go run main.go
```

### Environment Variables

| Enivironment Variable           | Description                                                                            |
| --------------------------------| -------------------------------------------------------------------------------------- |
| **PORT**                        | Port on to which application server listens to. Default value is 8080                  |
| **RESPONSE_TIMEOUT**            | Timeout for the server to write response. Default value is 100ms                       |
| **REQUEST_BODY_READ_TIMEOUT**   | Timeout for reading the request body send to the server. Default value is 20ms         |
| **RESPONSE_BODY_WRITE_TIMEOUT** | Timeout for writing the response body. Default value is 20ms                           |
| **PRODUCTION**                  | Flag to denote whether the server is running in production. Default value is false     |

## Author

{{.Author.Name}}<{{.Author.Email}}>
