# 練習問題 8.4

活動している `echo` ゴルーチンを数えるために接続ごとに
`sync.WaitGroup` を使うように `reverb2` を修正しなさい．
ゼロになったら，練習問題 8.3 で説明れているように TCP 接続の書き込み側を閉じなさい．
その練習問題で修正した `netcat3` クライアントは，標準入力が閉じられた後でも，
複数の並行な叫びの最後のエコーを持つことを検証しなさい．