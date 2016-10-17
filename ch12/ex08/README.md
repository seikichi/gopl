# 練習問題 12.8

`sexpr.Unmarshal` 関数は `json.Unmarshal` のようにデコードを開始する前に
バイトスライスの形で完全な入力を必要とします．`json.Decoder` のように，
`io.Reader` からデコードされる値の列を許す `sexpr.Decoder` 型を定義しなさい．
その新たな型を使うように `sexpr.Unmarshal` を変更しなさい．
