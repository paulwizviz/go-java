# Overview

This repository shows a side-by-side comparison of the steps involved in creating the same application in Go and Java. The focus of the comparison is the effort invovle in working with the source code **NOT** the executable (size, execution speed, etc).

We'll use the process of creating application to simulate a shopping cart as a basis for comparison.

## The shopping cart requirement

Items are presented one at a time, in a list, identified by name - for example "Apple" or "Banana" - to the cart.

Alternatively, multiple items are present multiple times in the list, so for example ["Apple", "Apple", "Banana"] is a basket with two apples and one banana.
 
The items are priced as follows:

 - Apples are 35p each
 - Bananas are 20p each
 - Melons are 50p each, but are available as ‘buy one get one free’
 - Limes are 15p each, but are available in a ‘three for the price of two’ offer

Given a list of shopping, calculate the total cost of those items.

## Implementation

For the purpose of comparisons, both Go and Java are implemented based on a monorepo architecture. In the case of Java, we have opted to use Gradle and to use only Java with no framework (e.g. Spring)

We will only implement a CLI base app where the app is given a list of fruits and the total is displayed as console output.