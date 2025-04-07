package gotriesearch

import "testing"

func Test_trie_Case_Sensitive(t *testing.T) {
	t.Run("test_functionality", func(t *testing.T) {
		tr := NewTrie()
		tr.Insert("hello")
		tr.Insert("hell")
		tr.Insert("heaven")
		tr.Insert("heavy")
		tr.Insert("he")
		tr.Insert("hi")
		tr.Insert("hellowworld")
		tr.Insert("hellow")
		tr.Insert("Hellow")
		tr.Insert("football")
		tr.Insert("football")
		tr.Insert("foot")
		tr.Insert("feet")
		tr.Insert("hellowworld")
		tr.Insert("hellowild")
		tr.Insert("hellowild")

		if !tr.Exist("hi") {
			t.Errorf("trie.Exist() = %v, want %v", false, true)
		}

		if tr.Exist("hai") {
			t.Errorf("trie.Exist() = %v, want %v", true, false)
		}

		if tr.Size() != 13 {
			t.Errorf("trie.Size() = %v, want %v", tr.Size(), 13)
		}

		if tr.Delete("fuut") {
			t.Errorf("trie.Delete() = %v, want %v", true, false)
		}

		tempResult := tr.PrefixSearch("f")
		if len(tempResult) != 3 {
			t.Errorf("trie.PrefixSearch('f') = %v, want %v", len(tempResult), 3)
		}
		for _, word := range tempResult {
			if word != "football" && word != "foot" && word != "feet" {
				t.Errorf("trie.PrefixSearch('f') = %v, want %v", word, "football/foot/feet")
			}
		}

		if tr.Size() != 13 {
			t.Errorf("trie.Size() = %v, want %v", tr.Size(), 13)
		}

		if !tr.Delete("hellowild") {
			t.Errorf("trie.Delete() = %v, want %v", false, true)
		}

		if tr.Size() != 12 {
			t.Errorf("trie.Size() = %v, want %v", tr.Size(), 12)
		}

		tempResult2 := tr.PrefixSearch("hellowild")
		if len(tempResult2) != 0 {
			t.Errorf("trie.PrefixSearch('hellowild') = %v, want %v", len(tempResult2), 0)
		}

		tempResult3 := tr.PrefixSearch("hello")
		if len(tempResult3) != 3 {
			t.Errorf("trie.PrefixSearch('hello') = %v, want %v", len(tempResult3), 3)
		}
		for _, word := range tempResult3 {
			if word != "hellow" && word != "hello" && word != "hellowworld" {
				t.Errorf("trie.PrefixSearch('hello') = %v, want %v", word, "hellow/hello/hellowworld")
			}
		}

		tempResult4 := tr.PrefixSearch("Hel")
		if len(tempResult4) != 1 {
			t.Errorf("trie.PrefixSearch('Hel') = %v, want %v", len(tempResult4), 1)
		}
		for _, word := range tempResult4 {
			if word != "Hellow" {
				t.Errorf("trie.PrefixSearch('hello') = %v, want %v", word, "Hellow")
			}
		}

	})
}

func Test_trie_Case_Insensitive(t *testing.T) {
	t.Run("test_functionality", func(t *testing.T) {
		tr := NewCaseInsensitiveTrie()
		tr.Insert("hello")
		tr.Insert("hell")
		tr.Insert("heaven")
		tr.Insert("heavy")
		tr.Insert("he")
		tr.Insert("hi")
		tr.Insert("hellowworld")
		tr.Insert("hellow")
		tr.Insert("football")
		tr.Insert("football")
		tr.Insert("foot")
		tr.Insert("feet")
		tr.Insert("hellowworld")
		tr.Insert("hellowild")
		tr.Insert("hellowild")

		if !tr.Exist("hi") {
			t.Errorf("trie.Exist() = %v, want %v", false, true)
		}

		if tr.Exist("hai") {
			t.Errorf("trie.Exist() = %v, want %v", true, false)
		}

		if tr.Size() != 12 {
			t.Errorf("trie.Size() = %v, want %v", tr.Size(), 12)
		}

		if tr.Delete("fuut") {
			t.Errorf("trie.Delete() = %v, want %v", true, false)
		}

		tempResult := tr.PrefixSearch("f")
		if len(tempResult) != 3 {
			t.Errorf("trie.PrefixSearch('f') = %v, want %v", len(tempResult), 3)
		}
		for _, word := range tempResult {
			if word != "football" && word != "foot" && word != "feet" {
				t.Errorf("trie.PrefixSearch('f') = %v, want %v", word, "football/foot/feet")
			}
		}

		if tr.Size() != 12 {
			t.Errorf("trie.Size() = %v, want %v", tr.Size(), 12)
		}

		if !tr.Delete("hellowild") {
			t.Errorf("trie.Delete() = %v, want %v", false, true)
		}

		if tr.Size() != 11 {
			t.Errorf("trie.Size() = %v, want %v", tr.Size(), 11)
		}

		tempResult2 := tr.PrefixSearch("hellowild")
		if len(tempResult2) != 0 {
			t.Errorf("trie.PrefixSearch('hellowild') = %v, want %v", len(tempResult2), 0)
		}

		tempResult3 := tr.PrefixSearch("hello")
		if len(tempResult3) != 3 {
			t.Errorf("trie.PrefixSearch('hello') = %v, want %v", len(tempResult3), 3)
		}
		for _, word := range tempResult3 {
			if word != "hellow" && word != "hello" && word != "hellowworld" {
				t.Errorf("trie.PrefixSearch('hello') = %v, want %v", word, "hellow/hello/hellowworld")
			}
		}

	})
}
