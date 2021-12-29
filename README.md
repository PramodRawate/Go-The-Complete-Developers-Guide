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

#### Variable Initialization and then assigning value
```
var deckSize int
deckSize = 52
```
#### We can initialize a variable outside of a function. We just can't assign a value to it.
```
var deckSize int

func main() {
	deckSize = 50
	fmt.Println(deckSize)
}
```
Below code using walrus operator(:=) will not work:
```
deckSize:= 52

func main() {
	fmt.Println(deckSize)
}
```


#### Basic types in Go

```
| Type    |    Values      |
| ------- | -------------- |
|   bool  | true, false    |
|  string | "hi", "hello"  |
|   int   | 0, -1000, 9999 |
| float64 | 10.1, 0.0009, -100.003|
```
## Functions in Go
```
func <function-name>() <return-type-of-func> {
	// """
	// function body
	// """
}
```
Example:-
```
func newCard() string {
	return "Five of Diamonds"
}
```

## Arrays and slices in Go
#### 1. Array - Creates an array of fixed size.
Array variable is *not a pointer* to the first array element, array variable *stands for the whole array*

#### 2. Slice - Creates an array which can grow and shrink.
Slice consists a pointer to the array.
###### Both Array and Slice must have elements of same types.

### Adding new elements to slice
```
append(sliceName, <new-element-to-add>)
```
Example:
```
cards := []string{"Ace of spades", newCard()}
cards = append(cards, "Six of Spades")
fmt.Println(cards)
```

### Looping through slices in Go
```
cards := []string{"Ace of spades", newCard(), "Six of Spades"}
for i, card :=  range cards {
    fmt.Println(i, card)
}
```
Output:
```
0 Ace of spades
1 Five of Diamonds
2 Six of Spades
```
## Custom type declaration in Go - Similar to Class in Java/OOP
```
type deck []string
```
Above creates a custom type or say like class which is of type []string(array of strings).
This can now be used instead of '[]string' type as below.

Primitive type:
```
cards := []string{"Ace of spades", newCard()}
```
Customer Type:
```
cards := deck{"Ace of spades", newCard()}
```
[]string type can be replaced with type 'deck'.

## Receiver Functions in Go
For above custom type we can define a receiver function. Unlike other programming language, receiver functions in Go have parameter list(called as *Receiver*) before the function name as below:
```
deck.go
-------
type deck []string

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
```
Any variable of type *deck*  now gets access to the *print* method.
```
main.go
-------
cards := deck{"Ace of spades", newCard()}
cards.print()
```

In terms of Java, consider type *deck* as a class and *print()* as its method.


## Using '_'(underscore) for variables not being used.
```
cards := deck{"Ace of spades", "Two of Diamonds", "Three of Hearts"}
for index, card := range d {
		fmt.Println(card)
	}
```
In above example Go compiler will give error since we are not using index. So we need to use it.
```
for index, card := range d {
    fmt.Println(index,card)
}
```
Hence, to avoid this we can use '_' instead of *index* and can ignore adding it in Println() as below:
```
for _, card := range d {
    fmt.Println(card)
}
```

## Returning multiple values in Go & Slice range syntax(slicename[start,end])
func <function-name>(<parameter1>, ...<parameterN>) (return1, return2) {
    return return1 return 2;
}
```
deck.go
--------
func deal(d deck, handSize int) (deck, deck) {
	return d[:3], d[3:]
}
```

```
main.go
--------
cards := newDec()
hand, remainingCards := deal(cards, 3)
```