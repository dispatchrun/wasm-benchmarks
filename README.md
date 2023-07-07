# wasm-benchmarks
Benchmarks to compare the performance of wasm compilers

## Running Benchmarks

```sh
make native wasmedge wasmtime wazero count=10 bench='BenchmarkJoinPath/#00' test=go
make native wasmedge wasmtime wazero count=10 bench='BenchmarkJoinPath/#00' test=tinygo
```

## Comparing Benchmark Results

```
$ benchstat -ignore=goarch,goos bench.tinygo.*
             │ bench.tinygo.native │          bench.tinygo.node          │        bench.tinygo.wasmedge        │        bench.tinygo.wasmtime        │         bench.tinygo.wazero          │
             │       sec/op        │   sec/op     vs base                │   sec/op     vs base                │   sec/op     vs base                │    sec/op     vs base                │
JoinPath/#00          126.85n ± 1%   47.46n ± 1%  -62.58% (p=0.000 n=10)   46.40n ± 1%  -63.42% (p=0.000 n=10)   74.80n ± 2%  -41.03% (p=0.000 n=10)   143.05n ± 1%  +12.77% (p=0.000 n=10)
$ benchstat -ignore=goarch,goos bench.go.*
pkg: github.com/stealthrocket/wasm-benchmarks/go
             │ bench.go.native │          bench.go.node          │        bench.go.wasmedge        │        bench.go.wasmtime         │         bench.go.wazero          │
             │     sec/op      │   sec/op     vs base            │    sec/op     vs base           │    sec/op      vs base           │    sec/op      vs base           │
JoinPath/#00      15.99n ± ∞ ¹   96.75n ± 1%  ~ (p=0.182 n=1+10)   80.98n ± ∞ ¹  ~ (p=1.000 n=1) ²   445.00n ± ∞ ¹  ~ (p=1.000 n=1) ²   254.50n ± ∞ ¹  ~ (p=1.000 n=1) ²
$ benchstat -ignore=goarch,goos bench.c.*
             │ bench.c.native │            bench.c.node            │          bench.c.wasmedge          │          bench.c.wasmtime           │            bench.c.wazero            │
             │     sec/op     │   sec/op     vs base               │   sec/op     vs base               │   sec/op     vs base                │   sec/op     vs base                 │
JoinPath/#00      24.53n ± 0%   24.29n ± 0%  -0.96% (p=0.001 n=10)   26.15n ± 2%  +6.63% (p=0.000 n=10)   48.84n ± 1%  +99.14% (p=0.000 n=10)   73.01n ± 1%  +197.72% (p=0.000 n=10)
```
