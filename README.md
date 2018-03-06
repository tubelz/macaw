# Macaw
Macaw is a 2D Game Engine using SDL2.
Macaw is written in Go with the [ECS architecture pattern](https://en.wikipedia.org/wiki/Entity%E2%80%93component%E2%80%93system).

Still under development and many improvements to be done.


![Demo](https://github.com/tubelz/pong-macaw/blob/master/pong.gif)

## Installation and requirements

* Go: https://golang.org/dl/
* SDL2:
	You will need to install SDL2 in your machine and the binding for Go.
	You can find more information here: [https://github.com/veandco/go-sdl2](https://github.com/veandco/go-sdl2)
* Macaw framework: `go get github.com/tubelz/macaw`

## Usage

You can find a working example in the repository [https://github.com/tubelz/pong-macaw/](https://github.com/tubelz/pong-macaw/)
That example covers many functionalities such as:

* Initialization
* Usage of gameloop
* Usage of components
* Usage of systems
* Usage of observers
* Creating a new system
* Usage of fonts

## Discussion (issues/suggestions)
If you have questions or suggestions you can go to **##macaw** at [freenode.net](https://freenode.net).
If there is a bug you can open an issue here.

## Development
There are many features to be developed and improve current functionalities. In the queue:

* Tests
* Sound system
* Scene manager
* GUI

## License
The code here is under the zlib license. You can read more [here](https://github.com/tubelz/macaw/LICENSE.txt)
