package inject

import "testing"

const MB = 1024 * 1024

func TestInject(t *testing.T) {

	f := &FunctionInject{CostMem: 10 * MB}
	f.Run()

}
