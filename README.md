# Use Case

Using Java, write a simple program that calculates the price of a basket of shopping.

The solution should be accomplished in roughly two hours.

Items are presented one at a time, in a list, identified by name - for example "Apple" or "Banana".

Multiple items are present multiple times in the list, so for example ["Apple", "Apple", "Banana"] is a basket with two apples and one banana.
 
Items are priced as follows:

 - Apples are 35p each
 - Bananas are 20p each
 - Melons are 50p each, but are available as ‘buy one get one free’
 - Limes are 15p each, but are available in a ‘three for the price of two’ offer

Given a list of shopping, calculate the total cost of those items.

## Solutions

There are two solutions: one in Go and the other in Pure Java.

For the Go version (install Go):

   1) Run the unit test, `cd go` and run the command `go test -v ./...`
   2) Run the application, run the command `go run main.go`

The Java version is replicated from Go. There is a precompiled Java class. To run the app from project root:

`java -cp jdk/src/main/java com.pjs.shoppingcart.ShoppingCart`

NOTE:

1. Both version of has a hard coded basket consisting of:

`["Apple", "Apple", "Banana", "Melon", "Melon", "Lime", "Lime", "Lime"]`

Which should add up to a total cost of `170p`

2. I am not familiar with Java Unit test hence relies on Go unit tests