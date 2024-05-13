package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"syscall/js"

	"github.com/onflow/flow-emulator/emulator"
	flowgo "github.com/onflow/flow-go/model/flow"
)

type Config struct {
	Verbose   bool
	LogFormat string
}

func main() {
	fmt.Println("Starting emulator")

	config := Config{
		Verbose:   true,
		LogFormat: "text",
	}

	logger := initLogger(config)

	blockchain, err := emulator.New(
		emulator.WithLogger(*logger),
	)

	if err != nil {
		panic(err)
	}

	// Mount the function on the JavaScript global object.
	js.Global().Set("GetAccount", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		address := flowgo.HexToAddress(args[0].String())
		account, err := blockchain.GetAccount(address)

		if err != nil {
			panic(err)
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

func initLogger(conf Config) *zerolog.Logger {

	level := zerolog.InfoLevel
	if conf.Verbose {
		level = zerolog.DebugLevel
	}
	zerolog.MessageFieldName = "msg"

	switch strings.ToLower(conf.LogFormat) {
	case "json":
		logger := zerolog.New(os.Stdout).With().Timestamp().Logger().Level(level)
		return &logger
	default:
		writer := zerolog.ConsoleWriter{Out: os.Stdout}
		writer.FormatMessage = func(i interface{}) string {
			if i == nil {
				return ""
			}
			return fmt.Sprintf("%-44s", i)
		}
		logger := zerolog.New(writer).With().Timestamp().Logger().Level(level)
		return &logger
	}

}
