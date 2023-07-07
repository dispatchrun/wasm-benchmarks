.PHONY: all clean native wasmedge wasmtime wazero

test ?= go
bench ?= .
count ?= 1
ifeq ($(test), c)
args = test $(count)
else
args = -test.bench '$(bench)' -test.count '$(count)'
endif
src.c = c/main.c
src.go = go/main.go

all: go.wasm tinygo.wasm

clean:
	rm -f *.test *.wasm

c.test: $(src.c)
	zig cc -O2 -o $@ $<

c.wasm: $(src.c)
	zig cc -O2 --target=wasm32-wasi -o $@ $<

go.test: $(src.go)
	gotip test -c -o $@ ./go

go.wasm: $(src.go)
	GOOS=wasip1 GOARCH=wasm gotip test -c -o $@ ./go

tinygo.test:
	tinygo test -opt 2 -c -o $@ ./go

tinygo.wasm: $(src.go)
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
