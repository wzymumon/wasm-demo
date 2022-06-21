package main

import (
	"fmt"
	"github.com/wasmerio/wasmer-go/wasmer"
	"io/ioutil"
	"time"
)

func main() {
	start := time.Now() // 获取当前时间
	wasmBytes, _ := ioutil.ReadFile("../demo/simple.wasm")

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, _ := wasmer.NewModule(store, wasmBytes)

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	instance, _ := wasmer.NewInstance(module, importObject)

	// Gets the `sum` exported function from the WebAssembly instance.
	sum, _ := instance.Exports.GetFunction("sum")

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, _ := sum(5, 37)
	fmt.Println(result) // 42!
	elapsed := time.Since(start)
	fmt.Println("wasmer-go 执行耗时：", elapsed)
}
