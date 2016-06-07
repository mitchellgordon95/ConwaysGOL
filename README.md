# ConwaysGOL
An implementation of [Conway's Game Of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) on a very large board (2^64 x 2^64 cells). The main implementation of the game of life algorithm comes from [Bill Gosper's HashLife](http://www.conwaylife.com/wiki/HashLife). For more information on the details of HashLife, see [this paper](http://www.drdobbs.com/jvm/an-algorithm-for-compressing-space-and-t/184406478).

## Setup

Clone this repository into
```
$GOPATH/src/github.com/mitchellgordon95
```
Or run
```
go get github.com/mitchellgordon95/ConwaysGOL
```
To download and install dependencies, make sure you have [Glide](https://github.com/Masterminds/glide), and run
```
glide install
```
in the root directory of the project.

To build, run the normal
```
go build
```
To run all tests (minus the vendored packages), run
```
go test $(glide novendor)
```

## Usage
Running
```
./ConwaysGOL
```
will put you in a text-mode interface.

You can use the -h flag for more startup options.

## Feature Wishlist

- Parallelize the HashLife generation routine
- Finish unit tests
- Implement the "true" HashLife algorithm so we can do trillions of generations quickly
- A GUI interface using OpenGL bindings for Go
- Board deserialization to files
- Better text animation for short delays
- Ability to stop animations halfway
