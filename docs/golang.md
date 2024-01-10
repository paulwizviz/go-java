# Go Implementation

Here are the steps of implementing the shopping cart.

## Setup

* Install Go

Refer to this [documentation](https://go.dev/doc/install) for instruction on how to install Go toolchain

## Project layout

Go does not mandate any layout.

However, there are some consensus conventions. Here are some

* [Organizing a Go module](https://go.dev/doc/modules/layout) - from the official blog
*[Design the architecture, name the components, document the details](https://paulwizviz.github.io/go/2022/12/23/go-proverb-architecture.html)
* [Go - Project Structure and Guidelines](https://dev.to/jinxankit/go-project-structure-and-guidelines-4ccm)
* [How to structure a Go project?](https://vsupalov.com/go-folder-structure/)

The consensus layout typically involves a structure that includes these folders at the project root:

* `cmd` - this is typically where you organise your main packages your application entry point
* `internal` - this contain shared packages common for the entire project.

As part of a Go project, you will find a `go.mod` file which is used to describe a project (also known as `module`) including it's dependencies. You can find the official description of go.mod file [here](https://go.dev/doc/modules/gomod-ref)

