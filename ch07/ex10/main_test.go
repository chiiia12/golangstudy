package main

import "testing"

func TestIsPalindromeTrue(t *testing.T) {
	s := sortPalindrome("hogeegoh")
	if !IsPalindrome(&s) {
		t.Errorf("hogeegoh is not palindrome")
	}
}

func TestIsPalindromeFalse(t *testing.T) {
	s := sortPalindrome("hogeego")
	if IsPalindrome(&s) {
		t.Errorf("hogeego is palindrome")
	}

}
