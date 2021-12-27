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
#### 1. Executable Package:
```
package main
```
Defines a package that can be compiled and then *executed*
###### Must have a func called *'main'*

#### 2. Reusable Package:

```
package <any userdefined name other than main> 
```
Defines a package that can be used as a dependency(helper code)

## Variables/Types in Go
#### Variable declaration
```
var <vairable-name> <variable-type> = <value>
var card string = "Ace of spades"
```
Type can be omitted in Go. GO compiler will infer type based on the value.
```
var card = "Ace of spades"
```

#### Alternate Variable declaration for above - walrus operator(:=)
```
card:="Five of Diamonds"
```
###### This can only be used during new variable initialization and not for reassigning



#### Basic types in Go

```
| Type    |    Values      |
| ------- | -------------- |
|   bool  | true, false    |
|  string | "hi", "hello"  |
|   int   | 0, -1000, 9999 |
| float64 | 10.1, 0.0009, -100.003|
```
