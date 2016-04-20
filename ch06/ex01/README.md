# 練習問題 6.1

これらの追加メソッドを実装しなさい．

```go
func (*IntSet) Len() int
func (*IntSet) Remove(x int)
func (*IntSet) Clear()
func (*IntSet) Copy() *IntSet
```

# 練習問題 6.2

`s.AddAll(1, 2, 3)` などのように
追加すべき値のリストが可能である可変個引数
`(*IntSet).AddAll(...int)` を実装しなさい．

# 練習問題 6.3

`(*IntSet).UnionWith` はワード単位のビット和算子である `|` を使用して
二つのセットの和集合を計算しています．セット操作に対応するメソッド
`IntersectWith`，`DifferenceWith`，`SymmetricDifference` を実装しなさい．

二つの集合の対照差 (symmetric difference) は，
どちらかの集合にはあるが，両方にはない要素を含む集合です．

二つの集合の差 (difference) は，
一方の集合にはあるが，もう一方にはない要素を含む集合です．

# 練習問題 6.4

`range` ループでの繰り返しに適した，
セットの要素を含むスライスを返す `Elems` メソッドを追加しなさい．
