# 練習問題 7.11

データベースのエントリをクライアントが作成，読み出し，更新，削除できるように
ハンドラを追加しなさい．たとえば，フォーム `/update?item=socks&price=6` 形式のリクエストは，
商品全体の中の一つの商品の価格を更新し，その商品がないもしくは価格が不正であればエラーを報告します．
(警告: この変更は，変数の並行な更新を発生させます)
