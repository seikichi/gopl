alexa.com にあるトップ10サイトを試した

```bash
> ./fetchall.alexatop10.sh
1.35s   129655  http://www.alexa.com/siteinfo/facebook.com
1.40s   103706  http://www.alexa.com/siteinfo/qq.com
1.45s   119566  http://www.alexa.com/siteinfo/wikipedia.org
1.60s   114700  http://www.alexa.com/siteinfo/amazon.com
1.60s   123220  http://www.alexa.com/siteinfo/google.com
1.60s   130313  http://www.alexa.com/siteinfo/twitter.com
1.60s   100315  http://www.alexa.com/siteinfo/baidu.com
1.61s    89036  http://www.alexa.com/siteinfo/google.co.in
1.66s   120031  http://www.alexa.com/siteinfo/yahoo.com
1.82s   115466  http://www.alexa.com/siteinfo/youtube.com
1.82s elapsed
```

また，応答しないウェブサイトの例として `server/main.go` を作成した．

```
> go run server/main.go &
> go run main.go http://localhost:8000
```

実行した結果，`fetchall` プログラムは停止してしまった．
