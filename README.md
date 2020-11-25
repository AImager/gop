# GOP

🇬🇧 English | 🇨🇳 [中文](README_zh.md)

## Install

`go get github.com/AImager/gop`

## Usage

register point in method

```go
package main

import (
	"fmt"
	"reflect"

	"github.com/AImager/gop"
)

type Aspect struct{}

func (a *Aspect) Before(point *gop.JoinPoint) bool {
	fmt.Println("before")
	return true
}

func (a *Aspect) After(point *gop.JoinPoint) {
	fmt.Println("after")
}

func (a *Aspect) Finally(point *gop.JoinPoint) {
	fmt.Println("finally")
}

func (a *Aspect) GetAspectExpress() string {
	return ".*\\.HelloAop"
}

type HelloAop struct {
}

func (h *HelloAop) HelloAop() {
	fmt.Println("helloAop")
}

func main() {
	gop.RegisterAspect([]gop.AspectInterface{&Aspect{}})
	gop.RegisterMethodPoint(reflect.TypeOf((*HelloAop)(nil)))

	h := &HelloAop{}
	h.HelloAop()
}
```

register point in function

```go
package main

import (
	"fmt"

	"github.com/AImager/gop"
)

type Aspect struct{}

func (a *Aspect) Before(point *gop.JoinPoint) bool {
	fmt.Println("before")
	return true
}

func (a *Aspect) After(point *gop.JoinPoint) {
	fmt.Println("after")
}

func (a *Aspect) Finally(point *gop.JoinPoint) {
	fmt.Println("finally")
}

func (a *Aspect) GetAspectExpress() string {
	return ".*"
}

func test() {
	fmt.Println("nihao test")
}

func main() {
	gop.RegisterAspect([]gop.AspectInterface{&Aspect{}})
	gop.RegisterFuncPoint(test)

	test()
}

```

## Contributing

PRs accepted.

## License

MIT © AImager