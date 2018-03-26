package ex03

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q)=false", p)
		}
	}
}

func randomNotPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) + 2
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; {
		r := rune(rng.Intn(0x1000))
		if unicode.IsLetter(r) {
			runes[i] = r
			for {
				r2 := rune(rng.Intn(0x1000))
				if unicode.IsLetter(r2) && r2 != r {
					runes[n-1-i] = r2
					break
				}
			}
			i++
		}
	}
	return string(runes)
}

//これだと一回でも対照なペアがあったらはじくかなり厳しい条件になってる。
//一個でも非対称なペアをみつけたら回文じゃないとして返しちゃえばよかった
//for {
//	if r[i]==r[n-1-i]{
//		return randomNotPalindrome()
//	}
//}
//return string(rune)

func TestRandomNotPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	//seedを定数にしちゃえばいつも同じ数字が得られる→debugに使える
	//randomのテストは再現できなくなるのでseedの出力は必須。
	t.Logf("Random seed:%d", seed)
	rng := rand.New(rand.NewSource(seed))

	count := 0;
	for i := 0; i < 1000; i++ {
		p := randomNotPalindrome(rng)

		if IsPalindrome(p) {
			count++
			t.Errorf("IsPalindrome(%v)=true,%v", p, []rune(p))
		}
	}
	t.Log("error count is", count)
}
