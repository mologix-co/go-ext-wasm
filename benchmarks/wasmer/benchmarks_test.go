package wasmer_test

import (
	"github.com/mologix-co/wasmer-go/wasmer"
	"path"
	"runtime"
	"testing"
)

const N = 100000

func GetBytes(wasmFile string) []byte {
	_, filename, _, _ := runtime.Caller(0)
	modulePath := path.Join(path.Dir(filename), "testdata", "benchmarks", wasmFile)

	bytes, _ := wasmer.ReadBytes(modulePath)

	return bytes
}

func benchmarkWasmer(b *testing.B, wasmFile string, exportName string, exportValues ...interface{}) {
	wasmBytes := GetBytes(wasmFile)
	instance, _ := wasmer.NewInstance(wasmBytes)
	export := instance.Exports[exportName]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, _ := export(exportValues...)

		_ = result
	}
}

func Benchmark_Wasmer_Nbody(b *testing.B) {
	benchmarkWasmer(b, "nbody.wasm", "main", N)
}

func Benchmark_Wasmer_Fibonacci_Recursive(b *testing.B) {
	benchmarkWasmer(b, "fib_recursive.wasm", "app_main")
}

func Benchmark_Wasmer_Pollard_Rho_128(b *testing.B) {
	benchmarkWasmer(b, "pollard_rho_128.wasm", "app_main")
}

func Benchmark_Wasmer_Snappy_Compress(b *testing.B) {
	benchmarkWasmer(b, "snappy_compress.wasm", "app_main")
}
