# GOP

## 安装

`go get github.com/AImager/gop`

## 使用

在方法上注册切点

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

在函数上注册切点

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

## 贡献者

PRs accepted.

## 协议

MIT © AImager