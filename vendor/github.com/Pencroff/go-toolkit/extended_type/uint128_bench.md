goos: windows

goarch: amd64

pkg: github.com/Pencroff/go-toolkit/extended_type

cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz

| Name |     Ops | Execution time    |     
|------|--------|-------------------|
|BenchmarkArithmetic|||
|BenchmarkArithmetic/Add_native |             1000000000       | 0.2635 ns/op      |
|BenchmarkArithmetic/Sub_native |             1000000000       | 0.2673 ns/op      |
|BenchmarkArithmetic/Mul_native |             1000000000       | 0.2868 ns/op      |
|BenchmarkArithmetic/Lsh_native |             1000000000       | 0.2413 ns/op      |
|BenchmarkArithmetic/Rsh_native |             1000000000       | 0.2908 ns/op      |
|BenchmarkArithmetic/Add        |             1000000000       | 0.2436 ns/op      |
|BenchmarkArithmetic/Sub        |             1000000000       | 0.2831 ns/op      |
|BenchmarkArithmetic/Mul        |             1000000000       | 0.2706 ns/op      |
|BenchmarkArithmetic/Lsh        |             1000000000       | 0.2887 ns/op      |
|BenchmarkArithmetic/Rsh        |             1000000000       | 0.2888 ns/op      |
|BenchmarkArithmetic/Cmp        |             1000000000       | 0.2918 ns/op      |
|BenchmarkArithmetic/Cmp64      |             1000000000       | 0.2656 ns/op      |
|BenchmarkDivision|||
|BenchmarkDivision/native_64/64 |             1000000000       | 0.2856 ns/op      |
|BenchmarkDivision/Div64_64/64  |             1000000000       | 0.5297 ns/op      |
|BenchmarkDivision/Div64_128/64 |             1000000000       | 0.5858 ns/op      |
|BenchmarkDivision/Div_64/64    |             100000000        | 10.50 ns/op       |
|BenchmarkDivision/Div_128/64-Lo|             50000415         | 25.54 ns/op       |
|BenchmarkDivision/Div_128/64-Hi|             35371622         | 33.25 ns/op       |
|BenchmarkDivision/Div_128/128  |             42097581         | 30.48 ns/op       |
|BenchmarkDivision/big.Int_128/64|            17251863         | 67.19 ns/op       |
|BenchmarkDivision/big.Int_128/128|           25723030         | 47.55 ns/op       |
|BenchmarkString|||
|BenchmarkString/Uint128          |           9631150          |       124.0 ns/op |
|BenchmarkString/big.Int          |           4403035          |       280.0 ns/op |