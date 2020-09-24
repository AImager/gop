module github.com/aimager/gop/example/basic

go 1.13

require (
	bou.ke/monkey v1.0.2
	github.com/aimager/gop v1.2.3
)

replace (
	bou.ke/monkey => ../../../../bouk/monkey
	github.com/aimager/gop => ../../
)
