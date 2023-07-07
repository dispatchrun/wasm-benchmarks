.PHONY: all native wasmedge wasmtime wazero

test ?= go
bench ?= .
count ?= 1
args = -test.bench '$(bench)' -test.count '$(count)'
src = $(wildcard go/*.go)

all: go.wasm tinygo.wasm

go.test: $(src)
	gotip test -c -o $@ ./go

go.wasm: $(src)
	GOOS=wasip1 GOARCH=wasm gotip test -c -o $@ ./go

tinygo.test:
	tinygo test -opt 2 -c -o $@ ./go

tinygo.wasm: $(src)
	tinygo test -opt 2 -c -o $@ -target wasi ./go

%.so: %.wasm
	wasmedgec $< $@

native: $(test).test
	./$(test).test $(args) | tee bench.$(test).native

wasmedge: $(test).so
	wasmedge $(test).so $(args) | tee bench.$(test).wasmedge

wasmtime: $(test).wasm
	wasmtime run -- $(test).wasm $(args)  | tee bench.$(test).wasmtime

wazero: $(test).wasm
	wazero run $(test).wasm $(args)  | tee bench.$(test).wazero
