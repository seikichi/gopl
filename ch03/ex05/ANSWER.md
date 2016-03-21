サンプルの (`gopl.io/ch3/mandelbrot/main.go`) の
`mandelbrot` 関数内を下記のように変更し，
発散と判定されるまでの計算回数の偶奇によって色分けした．

```go
		if cmplx.Abs(v) > 2 {
			if n%2 == 0 {
				return color.RGBA{0xff, 0x80, 0x00, 0xff}
			}
			return color.RGBA{0x00, 0x80, 0xff, 0xff}
		}
```

参考: http://azisava.sakura.ne.jp/mandelbrot/algorithm.html
