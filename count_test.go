package wordscount_test

import (
	"testing"

	"github.com/polaris1119/wordscount"
)

func TestStat(t *testing.T) {
	tests := []struct {
		name  string
		input string
		total int
		words int
	}{
		{"en1", "hello,playground", 3, 2},
		{"en2", "hello, playground", 3, 2},
		{"cn1", "你好世界", 4, 4},
		{"encn1", "Hello你好世界", 5, 5},
		{"encn2", "Hello 你好世界", 5, 5},
		{"encn3", "Hello，你好世界。", 7, 5},
		{"link1", "Hello，你好世界。https://studygolang.com Go中文网", 11, 9},
		{"en3", "I'm Lily.", 3, 2},
		{"en4-hyphen", "non-governmental organization", 2, 2},
		//{"en5", "", 3, 2},
		{"num1", "Five is 5.", 4, 3},
		{"num2", "1st and 2nd", 3, 3},
		{"num3", "100 and -100", 3, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := &wordscount.Counter{}
			if counter.Stat(tt.input); counter.Total != tt.total || counter.Words != tt.words {
				t.Errorf("Total = %v, want %v; Words=%v, want %v",
					counter.Total, tt.total, counter.Words, tt.words)
			}
		})
	}
}
