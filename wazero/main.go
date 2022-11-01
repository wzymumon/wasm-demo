package main

import (
	"context"
	"fmt"
	"github.com/tetratelabs/wazero"
	"os"
	"time"
)

var addWasm []byte

// main implements a basic function in both Go and WebAssembly.
func main() {
	start := time.Now() // 获取当前时间

	// Choose the context to use for function calls.
	ctx := context.Background()

	// Read a WebAssembly binary containing an exported "sum" function.
	addWasm, err := os.ReadFile("demo/simple.wasm")
	check(err)

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	// Add a module to the runtime named "wasm/math" which exports one function "add", implemented in WebAssembly.
	module, err := r.InstantiateModuleFromBinary(ctx, addWasm)
	check(err)

	// exported "sum" function.
	sum := module.ExportedFunction("sum")

	// get result
	result, err := sum.Call(ctx, 5, 37)
	check(err)

	// wazero Call only return []int64
	fmt.Printf("sum(5, 37) = %d\n", result[0])

	elapsed := time.Since(start)
	fmt.Println("wazero-go 执行耗时：", elapsed)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
