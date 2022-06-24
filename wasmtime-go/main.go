package main

import (
    "fmt"
    "github.com/bytecodealliance/wasmtime-go"
	"io/ioutil"
	"time"
)

func main() {
	start := time.Now() // 获取当前时间

    // Almost all operations in wasmtime require a contextual `store`
    // argument to share, so create that first
    store := wasmtime.NewStore(wasmtime.NewEngine())

	wasm, err := ioutil.ReadFile("../demo/simple.wasm")
	check(err)

    // Once we have our binary `wasm` we can compile that into a `*Module`
    // which represents compiled JIT code.
    module, err := wasmtime.NewModule(store.Engine, wasm)
    check(err)

    // Instantiates the module
    instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})
    check(err)

    // After we've instantiated we can lookup our `run` function and call
    // it.
	sum := instance.GetFunc(store, "sum")
	result, err := sum.Call(store, 5, 37)
	check(err)

	fmt.Printf("sum(5, 37) = %d\n", result.(int32))
    check(err)

	elapsed := time.Since(start)
	fmt.Println("wasmtime-go 执行耗时：", elapsed)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}