### FirstAPI
This is a small web api for testing out go's gin gonin framework for the rest apis.

### Dependencies
```bash
go >= 1.16
gin
gorm
mysql drivers
```

### How to run?
In order to run this application you just need to run following commands:

- To install all the missing deps
```bash
go mod tidy
```

- To run the webserver
```bash
go run .
```

- Alternatively, you can also run the webserver by first building the application and then running it
```bash
go build

./firstapi
``` 
