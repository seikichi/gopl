# 練習問題 3.4

1.7節のリサジューの例の方法に従って，
面を計算して SVG データをクライアントに対して書き出す
ウェブサーバを作成しなさい．サーバは次の様に `Content-Type` を指定しなければなりません．

```go
w.Header().Set("Content-Type", "image/svg+xml")
```
