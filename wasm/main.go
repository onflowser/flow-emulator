package main

import (
	"fmt"
	"syscall/js"

	"github.com/onflow/flow-emulator/emulator"
	flowgo "github.com/onflow/flow-go/model/flow"
)

func main() {
	fmt.Println("Starting emulator")

	emulator, error := emulator.New()

	if error != nil {
		panic(error)
	}

	// Mount the function on the JavaScript global object.
	js.Global().Set("GetAccount", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		address := flowgo.HexToAddress(args[0].String())
		account, error := emulator.GetAccount(address)

		if error != nil {
			panic(error)
		}

		return map[string]interface{}{
			"address": account.Address.String(),
			"balance": account.Balance,
			// "contracts": account.Contracts,
		}
	}))

	// Prevent the function from returning, which is required in a wasm module
	select {}
}
