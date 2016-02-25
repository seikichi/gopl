Wikipedia の [長いページ](https://ja.wikipedia.org/wiki/%E7%89%B9%E5%88%A5:%E9%95%B7%E3%81%84%E3%83%9A%E3%83%BC%E3%82%B8) を参考に，
[過去に存在したダイエーの店舗](https://ja.wikipedia.org/wiki/%E9%81%8E%E5%8E%BB%E3%81%AB%E5%AD%98%E5%9C%A8%E3%81%97%E3%81%9F%E3%83%80%E3%82%A4%E3%82%A8%E3%83%BC%E3%81%AE%E5%BA%97%E8%88%97) を調査対象とした．

```bash
> go run main.go https://ja.wikipedia.org/wiki/%E9%81%8E%E5%8E%BB%E3%81%AB%E5%AD%98%E5%9C%A8%E3%81%97%E3%81%9F%E3%83%80%E3%82%A4%E3%82%A8%E3%83%BC%E3%81%AE%E5%BA%97%E8%88%97
4.59s elapsed
> go run main.go https://ja.wikipedia.org/wiki/%E9%81%8E%E5%8E%BB%E3%81%AB%E5%AD%98%E5%9C%A8%E3%81%97%E3%81%9F%E3%83%80%E3%82%A4%E3%82%A8%E3%83%BC%E3%81%AE%E5%BA%97%E8%88%97
1.98s elapsed
```

とキャッシュされてそう (?) である．

続けて取得したページを標準出力にリダイレクトし，
その内容を保存したが，二回とも内容は同じであった (see [main.go](./main.go))．
