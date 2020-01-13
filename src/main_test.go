package main

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func generator(n int) []int {
	s := make([]int, n)

	src := rand.NewSource(time.Now().UnixNano())
	randN := rand.New(src)
	maxValue := 2*n/2
	for i:=0; i < n; i++  {
		s[i] = randN.Intn(maxValue)
	}
	return s
}

const MAXN = 1000000

func TestMergesort(t *testing.T)  {
	s := []int{5,5,6,1,2,3,9,7,5,0}
	mergesort(s)

	assert.Equal(t, []int{0,1,2,3,5,5,5,6,7,9}, s)
}

func BenchmarkMergesort(b *testing.B) {
	for i := 0 ; i < b.N; i++ {
		s := generator(MAXN)
		b.StartTimer()
		mergesort(s)
		b.StopTimer()
	}
}
