package jmg

import (
	"context"
	"fmt"
	"testing"
)

const in = "美味い美味すぎるっ！十万石饅頭！"

func TestClient(t *testing.T) {
	fmt.Println(NewClient("localhost:12000").Jumanpp(context.Background(), in))
}

func BenchmarkClientDocker(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewClient("localhost:12000").Jumanpp(context.Background(), in)
	}
}

func BenchmarkClient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewClient("localhost:12001").Jumanpp(context.Background(), in)
	}
}
