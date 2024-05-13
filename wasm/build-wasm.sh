CGO_ENABLED=0 GOOS=js GOARCH=wasm go build -tags=no_cgo -o wasm/flow-emulator.wasm wasm/main.go
