ベンチマークを取得したところ以下の結果が得られた

```bash
leticia:seikichi% go test -bench .
testing: warning: no tests to run
PASS
BenchmarkRun10-4                 3000000               434 ns/op
BenchmarkRun100-4                1000000              1821 ns/op
BenchmarkRun1000-4                100000             16823 ns/op
BenchmarkRunInefficiently10-4    2000000               923 ns/op
BenchmarkRunInefficiently100-4    100000             12043 ns/op
BenchmarkRunInefficiently1000-4     3000            398830 ns/op
ok      github.com/seikichi/gopl/ch01/ex03      10.799s
```
