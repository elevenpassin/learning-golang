module github.com/buoyantair/learning-golang

go 1.18

require (
	github.com/buoyantair/learning-golang/greetings v0.0.0
	github.com/buoyantair/learning-golang/hello v0.0.0
)

replace github.com/buoyantair/learning-golang/greetings => ./greetings

replace github.com/buoyantair/learning-golang/hello => ./hello
