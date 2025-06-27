package main

import (
	"slices"
	"testing"
)

func wordBreak(s string, wordDict []string) bool {
	memo := make(map[string]bool)
	return wordBreakMemo(s, wordDict, memo)
}

func wordBreakMemo(s string, wordDict []string, memo map[string]bool) bool {
	if val, found := memo[s]; found {
		return val
	}

	if s == "" {
		return true
	}

	for i := 1; i <= len(s); i++ {
		prefix := s[0:i]
		if slices.Contains(wordDict, prefix) {
			if wordBreakMemo(s[i:], wordDict, memo) {
				memo[s] = true
				return true
			}
		}
	}

	memo[s] = false
	return false
}

func TestWordBreak(t *testing.T) {
	s := "bb"
	wordDict := []string{"a", "b", "bbb", "bbbb"}
	res := wordBreak(s, wordDict)
	if res != true {
		t.Errorf("wordBreak returned %v, expected %v", res, true)
	}
}
