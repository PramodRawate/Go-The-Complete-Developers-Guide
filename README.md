# Go-The-Complete-Developers-Guide

## Go Command line tools
#### 1. go build - compiles a bunch of go source code files
```
go build <filename.go>
go build main.go
``` 
#### 2. go run - Compiles and executes one or two files
```
go run <filename.go>
go run main.go
```
#### 3. go fmt - Formats all the code in each file in the current directory
```
go fmt <filename.go> || go fmt <modulename>
go fmt main.go
```
#### 4. go installs - Compiles and "installs" a package
```
go install <filename.go> || go fmt <modulename>
go install main.go
```

#### 5. go get - Downloads the source code of someone else's package
```
go get <pacakge/module path>
```
#### 6. go test - Run any tests associated with the current project
```
go test <file/package name>
```


## Packages in Go