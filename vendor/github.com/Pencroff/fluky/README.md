# fluky

![fluky](./assets/fluky-logo.min.png)


Happening by or depending on chance in Golang

All rng provides uint64 values (1.845e19) and float64 values

## Dieharder summary

| Name       | Test Data | Time  | PASS | WEAK | FAIL | Total | Draw test                       | References                                                                                                                                      |
|------------|:---------:|:-----:|:----:|:----:|:----:|:------|---------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------|
| Built In   |  229GiB   | 35:44 | 112  |  2   |  0   | 114   | [img](./out/built-in_out.png)   | [Ref](https://pkg.go.dev/math/rand)                                                                                                             |
| Pcg64      |  229GiB   | 34:34 | 113  |  1   |  0   | 114   | [img](./out/pcg64_out.png)      | [Ref](https://www.pcg-random.org/)                                                                                                              |
| Small Prng |  229GiB   | 37:49 | 113  |  1   |  0   | 114   | [img](./out/small_prng_out.png) | [Ref1](https://burtleburtle.net/bob/rand/smallprng.html), [Ref2](https://www.pcg-random.org/posts/bob-jenkins-small-prng-passes-practrand.html) | 
| ---------- | --------- | ----- | ---- | ---- | ---- | ----- | ---------                       |                                                                                                                                                 |
| Squares    |  229GiB   | 36:02 | 103  |  5   |  6   | 114   | [img](./out/squares_out.png)    | [Ref](https://arxiv.org/abs/2004.06278)                                                                                                         |
| MMIX       |  229GiB   | 36:19 |  71  |  6   |  37  | 114   | [img](./out/mmix_out.png)       | [Ref](https://en.wikipedia.org/wiki/Linear_congruential_generator)                                                                              |
| Musl       |  229GiB   | 35:52 |  70  |  7   |  37  | 114   | [img](./out/musl_out.png)       | [Ref](https://en.wikipedia.org/wiki/Linear_congruential_generator)                                                                              |

Detailed results please check in the [`dieharder-result`](./dieharder-result) directory.

Few more RNG and test results can be found in [rngset repo](https://github.com/TyeolRik/rngset).

## Benchmark

goos: windows
goarch: amd64
pkg: github.com/Pencroff/fluky/rng
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
BenchmarkRndGen 

| Name             |    Ops    | Execution time |
|------------------|:---------:|:--------------:|
| BuiltIn          | 243794713 |  4.721 ns/op   |
| PcgCRng with CGO | 18878521  |  63.00 ns/op   |
| PcgRng pure GO   | 205031472 |  5.755 ns/op   |
| Small Prng       | 518096238 |  2.327 ns/op   |
| Squares          | 146871903 |  8.507 ns/op   |


## Nist tests

TestFrequencyMonobits stats (1342177280 numbers) ~ 10 Gb of random bytes each

| Name       |   P-value    |               Values                | Status |
|------------|:------------:|:-----------------------------------:|-------:|
| mmix       | 5.940685e-01 | [ 1: 42949594860 ; 0: 42949751060 ] |   PASS |
| Musl       | 9.624453e-01 | [ 1: 42949679860 ; 0: 42949666060 ] |   PASS |
| built_in   | 1.700484e-01 | [ 1: 42949874024 ; 0: 42949471896 ] |   PASS |
| small_prng | 4.221443e-01 | [ 1: 42949790591 ; 0: 42949555329 ] |   PASS |
| pcg64      | 1.271434e-01 | [ 1: 42949449414 ; 0: 42949896506 ] |   PASS |

Max value: Musl (9.624453e-01)

--- PASS: TestFrequencyMonobits (50.06s)

## Execute dieharder tests

Check the [**README.md**](./container/README.md).

## References

* [List of random number generators](https://en.wikipedia.org/wiki/List_of_random_number_generators)
* [Dieharder test](https://webhome.phy.duke.edu/~rgb/General/dieharder.php)
* [NIST Statistical Test Suite](https://csrc.nist.gov/Projects/Random-Bit-Generation/Documentation-and-Software)
* [NIST Statistical Test Suite implementation (Python)](https://github.com/GINARTeam/NIST-statistical-test)
* Similar projects - [ref1](https://github.com/skeeto/rng-go), [ref2](https://github.com/TyeolRik/rngset)