# 練習問題 5.17

HTML ノードツリーとゼロ個以上の名前が与えられたら，
それらの名前の一つと一致する要素をすべてを返す可変個引数関数
`ElementsByTagName` を書きなさい．二つの呼び出し例を次に示します.

```go
func ElementsByTagName(doc *html.Node, name ...string) []*html.Node

images := ElementsByTagName(doc, "img")
headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
```
