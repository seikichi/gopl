# 練習問題 7.5

`io` パッケージの `LimitReader` 関数は 
`io.Reader` である `r` とバイト数 `n` を受け取り，
`r` から呼び出す別の `Reader` を返しますが，`n` バイト呼び出した後に
ファイルの終わりの状態を報告します．その関数を実装しなさい．

```go
func LimitReader(r io.Reader, n int64) io.Reader
```
