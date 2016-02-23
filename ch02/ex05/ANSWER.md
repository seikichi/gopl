ハッカーのたのしみ版が最速

```bash
> go test -bench .                                                                                                                                                              [~/go/src/github.com/seikichi/gopl/ch02/ex05]
testing: warning: no tests to run
PASS
BenchmarkPopCount-4             200000000                7.05 ns/op
BenchmarkPopCountByLoop-4       100000000               15.0 ns/op
BenchmarkPopCountByShift-4      30000000                50.7 ns/op
BenchmarkPopCountByClear-4      30000000                55.3 ns/op
BenchmarkPopCountByHD-4         1000000000               2.81 ns/op
ok      github.com/seikichi/gopl/ch02/ex05      10.045s
```
