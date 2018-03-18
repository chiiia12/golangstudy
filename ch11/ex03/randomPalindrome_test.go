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
	n := rng.Intn(600) + 2
	runes := make([]rune, n)
	//TODO:可能性的に回文が出てしまうのをなんとかする
	for i := 0; i < (n+1)/2; {
		r := rune(rng.Intn(0x1000))
		if unicode.IsLetter(r) {
			runes[i] = r
			i++
			for {
				r2 := rune(rng.Intn(0x1000))
				if unicode.IsLetter(r2) && r2 != r {
					runes[n-1-i] = r2
					break
				}
			}
		}
	}
	return string(runes)
}

func TestRandomNotPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Randome seed:%d", seed)
	rng := rand.New(rand.NewSource(seed))

	count := 0;
	for i := 0; i < 1000; i++ {
		p := randomNotPalindrome(rng)

		if IsPalindrome(p) {
			count++
			t.Errorf("IsPalindrome(%v)=true", p)
		}
	}
	t.Log("error count is", count)
}
