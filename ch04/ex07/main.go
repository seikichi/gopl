package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(string(reverse([]byte("こんにちは!!"))))
}

type subslice struct {
	s    []byte
	i, j int
}

func (res *subslice) popFront() rune {
	r, d := utf8.DecodeRune(res.s[res.i:res.j])
	res.i += d
	return r
}

func (res *subslice) popBack() rune {
	r, d := utf8.DecodeLastRune(res.s[res.i:res.j])
	res.j -= d
	return r
}

func (res *subslice) pushFront(r rune) {
	copy(res.s[res.i:], []byte(string(r)))
	res.i += utf8.RuneLen(r)
}

func (res *subslice) pushBack(r rune) {
	d := utf8.RuneLen(r)
	copy(res.s[res.j-d:], []byte(string(r)))
	res.j -= d
}

func (res *subslice) empty() bool {
	return res.i >= res.j
}

type queue []rune

func (q queue) front() rune  { return q[0] }
func (q *queue) push(r rune) { *q = append(*q, r) }
func (q *queue) pop()        { *q = (*q)[1:] }

func frontInsertSpace(rest, result subslice) int {
	return rest.i - result.i
}

func backInsertSpace(rest, result subslice) int {
	return result.j - rest.j
}

func reverse(s []byte) []byte {
	rest := subslice{s, 0, len(s)}
	result := subslice{s, 0, len(s)}
	// note: que は高々サイズ4の []rune
	que := queue{}
	for !rest.empty() {
		// 後ろから要素を取り出して...
		r := rest.popBack()
		// 先頭に十分な空きができるまで先頭の要素をキューに追加
		for !rest.empty() && frontInsertSpace(rest, result) < utf8.RuneLen(r) {
			que.push(rest.popFront())
		}
		// 後ろから取り出した要素を先頭に移動
		result.pushFront(r)
		// 先頭から取り出した要素を詰めれるだけ後ろに逆順で詰める
		for len(que) > 0 {
			if backInsertSpace(rest, result) < utf8.RuneLen(que.front()) {
				break
			}
			result.pushBack(que.front())
			que.pop()
		}
	}
	// 取り出せる要素が無くなったら，キューの要素を余った隙間に後ろに逆順で追加
	for len(que) > 0 {
		result.pushBack(que.front())
		que.pop()
	}
	return s
}
