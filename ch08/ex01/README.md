# 練習問題 8.1

`clock2` を修正しポート番号を受け付けるようにしなさい．
そして，一度に複数の時計サーバのクライアントとして振る舞い，
それぞれのサーバから時刻を読み出し，ビジネスオフィスで見かける
壁にかかった複数の時計に似せた表で結果を表示するプログラム
`clockwall` を書きなさい．みなさんが地理的に分散したコンピュータへ
アクセスできるなら，サーバをリモートで実行しなさい．
そうでなければ，次のように擬似的なタイムゾーンを用いて異なる
ポートでローカルにサーバを実行しなさい．

```go

$ TZ=US/Eastern    ./clock2 -port 8010 &
$ TZ=Asia/Tokyo    ./clock2 -port 8020 &
$ TZ=Europe/London ./clock2 -port 8030 &
$ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030
```
