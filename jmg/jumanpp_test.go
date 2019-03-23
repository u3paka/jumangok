package jmg

import (
	"testing"

	"github.com/k0kubun/pp"
)

const txt = "美味い！美味すぎるっ！ 十万石饅頭！！風が...語りかけます。"

func TestJumanpp(t *testing.T) {
	cl, err := NewService("jumanpp")
	if err != nil {
		return
	}
	ws, _ := cl.GetWords(txt)
	ews := Extract(ws, func(w *Word) bool {
		// if w.HasDomain("料理・食事") || w.HasCategory("抽象物") {
		// 	return true
		// }
		// return false
		return true
	})
	pp.Println(ews)
}
func TestRawParse(t *testing.T) {
	cl, err := NewService("jumanpp")
	if err != nil {
		return
	}
	w := cl.RawParse(txt)
	pp.Println(w)
}
func TestNewClient(t *testing.T) {
	cl, err := NewService("jumanpp")
	if err != nil {
		return
	}
	cl.RawParse(txt)
}

//66706508 faster than mutex
func BenchmarkChannel(b *testing.B) {
	cl, err := NewService("jumanpp")
	if err != nil {
		return
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cl.RawParse(txt)
	}
}
