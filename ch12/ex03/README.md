# 練習問題 12.3

`encode` 関数の不足している場合を実装しなさい．
ブーリアンを `t` と `nil` として，浮動小数点数に Go の表記を使い，
`1 + 2i` などの複素数は `#C(1.0 2.0)` としてエンコードしなさい．
インタフェースは，たとえば `("[]int" (1 2 3))` といったように
型の名前と値の組としてエンコードできますが，この表記は曖昧であることに注意しなさい．
`reflect.Type.String` メソッドは異なる型に対して同じ文字列を返すかもしれません．