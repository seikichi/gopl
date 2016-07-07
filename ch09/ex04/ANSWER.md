```bash
% uname -a
Darwin leticia.local 15.3.0 Darwin Kernel Version 15.3.0: Thu Dec 10 18:40:58 PST 2015; root:xnu-3248.30.4~1/RELEASE_X86_64 x86_64 i386 MacBookPro11,1 Darwin
% sysctl -n machdep.cpu.brand_string
Intel(R) Core(TM) i5-4258U CPU @ 2.40GHz
% go test -bench .
testing: warning: no tests to run
PASS
BenchmarkRun10-4          500000              2904 ns/op
BenchmarkRun100-4          50000             28490 ns/op
BenchmarkRun1000-4          5000            321181 ns/op
BenchmarkRun10000-4          300           4162918 ns/op
ok      github.com/seikichi/gopl/ch09/ex04      6.720s
```
