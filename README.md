# Function Inject tool

## Install

`go get github.com/xxlv/go-function-inject@v1.0.0`

## Use

```go 
    // mock function cost 10MB mem
	f := &FunctionInject{CostMem: 10 * MB}
	f.Run()

```



