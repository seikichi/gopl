# 練習問題 7.10

`sort.Interface` 型は，他の利用にも応用できます．
列 `s` が回文 (palindrome) であるか，つまり列を逆順にしても変わらないかを報告する関数
`IsPalindrome(s sort.Interface) bool` を書きなさい．
インデックス `i` と `j` の要素は，`!s.Less(i, j) && !s.Less(j, i)` であれば等しいとみなしなさい．
