package Question_by_day

type Trie struct {
	isWord bool
	Word   string
	next   map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{false, "", make(map[byte]*Trie)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(this.Word) > 0 {
		this.isWord = false
		curWord := this.Word
		this.Word = ""
		this.Insert(curWord)
	}
	if len(word) == 0 {
		this.isWord = true
		return
	}
	word_ := []byte(word)
	first := word_[0]
	if node, ok := this.next[first]; ok {
		node.Insert(string(word_[1:]))
	} else {
		this.next[first] = &Trie{true, string(word_[1:]), make(map[byte]*Trie)}
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.isWord && len(this.Word) == 0
	}
	if this.isWord && len(this.Word) > 0 {
		return this.Word == word
	}
	word_ := []byte(word)
	first := word_[0]
	if node, ok := this.next[first]; ok {
		return node.Search(string(word_[1:]))
	} else {
		return false
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	// fmt.Println(this.isWord,this.Word,prefix)
	prefix_ := []byte(prefix)
	first := prefix_[0]
	if node, ok := this.next[first]; ok {
		return node.StartsWith(string(prefix_[1:]))
	} else {
		if !this.isWord || len(this.Word) < len(prefix_) {
			return false
		}
		flag := true
		for i := 0; i < len(prefix_); i++ {
			if prefix_[i] != this.Word[i] {
				flag = false
				break
			}
		}
		return flag
	}
}
