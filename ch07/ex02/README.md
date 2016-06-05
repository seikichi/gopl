# 練習問題 7.2

下記のシグニチャを持つ関数 `CountingWriter` を書きなさい．
`io.Writer` が与えられたなら，それを包む新たな `Writer` と `int64` 変数への
ポインタを返します．その変数は新たな `Writer` に書き込まれたバイト数を常に保持しています．

```go
func CountingWriter(w io.Writer) (io.Writer, *int64)
```
