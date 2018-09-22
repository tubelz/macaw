# Macaw [![Build Status](https://travis-ci.org/tubelz/macaw.svg?branch=master)](https://travis-ci.org/tubelz/macaw.svg?branch=master) [![Coverage Status](https://codecov.io/gh/tubelz/macaw/branch/master/graph/badge.svg)](https://codecov.io/gh/tubelz/macaw) [![GoDoc](https://godoc.org/github.com/tubelz/macaw?status.svg)](https://godoc.org/github.com/tubelz/macaw) [![Go Report Card](https://goreportcard.com/badge/github.com/tubelz/macaw)](https://goreportcard.com/report/github.com/tubelz/macaw)

Macaw is a 2D Game Engine using SDL2.
Macaw is written in Go with the [ECS architecture pattern](https://en.wikipedia.org/wiki/Entity%E2%80%93component%E2%80%93system).

![Demo](https://github.com/tubelz/pong-macaw/blob/master/pong.gif)

## Installation and requirements

* Go: https://golang.org/dl/
* SDL2:
	You will need to install SDL2 in your machine and the binding for Go.
	You can find more information on how to install on your OS here: [https://github.com/veandco/go-sdl2](https://github.com/veandco/go-sdl2)
	Also, make sure if you are compiling from source code to enable CGO (`export CGO_ENABLED=1`)
* Macaw framework: `go get github.com/tubelz/macaw`

## Usage

You can find a working example in the repository [https://github.com/tubelz/pong-macaw/](https://github.com/tubelz/pong-macaw/)
That example covers many functionalities such as:

* Initialization
* Game loop
* Usage of entities, components and systems (**ECS**)
* Scene
* Camera
* Observers
* Creating a new system
* Fonts
* Input handler

## Discussion (issues/suggestions)

If you have questions, suggestions, or just want to chat about our Game Engine you can go to **##macaw** at [freenode.net](https://webchat.freenode.net?channels=%23%23macaw).
If there is a bug you can open an issue here.
Your input is fundamental for the project's success. :)

## Contributing

There's always something to be worked on! Don't be afraid to open an issue or submit a PR.
Please check the [contributing guide](https://github.com/tubelz/macaw/blob/master/CONTRIBUTING.md) for more information!

## License

The code here is under the zlib license. You can read more [here](https://github.com/tubelz/macaw/LICENSE.txt)
