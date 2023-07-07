# wasm-benchmarks
Benchmarks to compare the performance of wasm compilers

## Running Benchmarks

```sh
make native wasmedge wasmtime wazero count=10 bench='BenchmarkJoinPath/#00' test=go
make native wasmedge wasmtime wazero count=10 bench='BenchmarkJoinPath/#00' test=tinygo
```

## Comparing Benchmark Results

```
$ benchstat -ignore=goarch,goos bench.*
pkg:
             │ bench.tinygo.native │        bench.tinygo.wasmedge        │        bench.tinygo.wasmtime        │         bench.tinygo.wazero          │
             │       sec/op        │   sec/op     vs base                │   sec/op     vs base                │    sec/op     vs base                │
JoinPath/#00          126.85n ± 1%   46.40n ± 1%  -63.42% (p=0.000 n=10)   74.80n ± 2%  -41.03% (p=0.000 n=10)   143.05n ± 1%  +12.77% (p=0.000 n=10)

pkg: github.com/stealthrocket/benchmarks/go
             │ bench.go.native │          bench.go.wasmedge           │           bench.go.wasmtime            │            bench.go.wazero             │
             │     sec/op      │   sec/op     vs base                 │    sec/op     vs base                  │    sec/op     vs base                  │
JoinPath/#00       16.09n ± 0%   77.85n ± 1%  +383.84% (p=0.000 n=10)   435.70n ± 0%  +2607.89% (p=0.000 n=10)   238.85n ± 0%  +1384.46% (p=0.000 n=10)
```
