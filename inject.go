package inject

import (
	"bufio"
	"os"
	"time"
)

type FunctionInject struct {
	CpuWindowSecond int     // cpu时间窗口 单位s
	CostCpu         float64 // 花费的cps 百分比
	CostMem         int     // 花费的内存 byte数
	CostDelayMill   uint    // 执行时间延迟
	EmptySpinCount  uint    // 自旋次数
}

func (f *FunctionInject) Run() {

	// 消耗cpu
	if f.CostCpu > 0 {
		compute(float64(f.CpuWindowSecond), f.CostCpu, 0)
	}
	// cpu空跑
	if f.EmptySpinCount > 0 {
		i := 0
		for {
			i += 1
			if i > int(f.EmptySpinCount) {
				break
			}
		}
	}
	// 消耗内存
	var data []byte
	if f.CostMem > 0 {
		data = make([]byte, f.CostMem)
		for i := range data {
			data[i] = 1
		}
	}
	println(len(data))
	// delay
	if f.CostDelayMill > 0 {
		time.Sleep(time.Duration(f.CostDelayMill) * time.Millisecond)
	}

}

func compute(t, percent float64, id int) {
	var r int64 = 1000 * 1000
	totalNanoTime := t * (float64)(r)               // 纳秒
	runtime := totalNanoTime * percent              // 纳秒
	sleeptime := totalNanoTime - runtime            // 纳秒
	starttime := time.Now().UnixNano()              // 当前的纳秒数
	d := time.Duration(sleeptime) * time.Nanosecond // 休眠时间
	for float64(time.Now().UnixNano())-float64(starttime) < runtime {
	}
	time.Sleep(d)
}

func main() {

	f := &FunctionInject{}
	f.Run()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
	}

}
