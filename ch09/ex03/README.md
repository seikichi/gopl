# 練習問題 9.3

`Func` 型と `(*Memo).Get` メソッドを拡張して，呼び出しもとがオプションの
`done` チャネルを渡して，そのチャネルを介して操作をキャンセルできるようにしなさい (8.9 節)．
キャンセルされた `Func` 呼び出しの結果はキャッシュされるべきではありません．
