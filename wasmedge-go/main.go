package main

import (
	"fmt"
	"time"
	"github.com/second-state/WasmEdge-go/wasmedge"
)

func ListInsts(name *string, mod *wasmedge.Module) {
	if name == nil {
		fmt.Println(" --- Exported instances of the anonymous module")
	} else {
		fmt.Println(" --- Exported instances of the module", *name)
	}
	nf := mod.ListFunction()
	fmt.Println("     --- Functions (", len(nf), ") : ", nf)
	nt := mod.ListTable()
	fmt.Println("     --- Tables    (", len(nt), ") : ", nt)
	nm := mod.ListMemory()
	fmt.Println("     --- Memories  (", len(nm), ") : ", nm)
	ng := mod.ListGlobal()
	fmt.Println("     --- Globals   (", len(ng), ") : ", ng)
}

func main() {
	start := time.Now() // 获取当前时间
	/// Set not to print debug info
	wasmedge.SetLogErrorLevel()

	/// Create configure
	var conf = wasmedge.NewConfigure(wasmedge.WASI)

	/// Create store
	var store = wasmedge.NewStore()

	/// Create VM by configure and external store
	var vm = wasmedge.NewVMWithConfigAndStore(conf, store)

	/// Register fibonacci wasm as module name "wasm"
	vm.RegisterWasmFile("wasm", "demo/simple.wasm")

	/// Run fib[22] directly
	result, err := vm.ExecuteRegistered("wasm", "sum", uint32(5), uint32(37))
	if err != nil {g
		fmt.Println(" !!! Error: ", err.Error())
	}
	// } else if result != nil {
	// 	for _, val := range result {
	// 		fmt.Println(" Return value: ", val)
	// 	}
	// }
	fmt.Printf("sum(5, 37) = %d\n", result[0])
	vm.Release()
	conf.Release()
	store.Release()
	elapsed := time.Since(start)
	fmt.Println("wasmedge-go 执行耗时：", elapsed)
}
