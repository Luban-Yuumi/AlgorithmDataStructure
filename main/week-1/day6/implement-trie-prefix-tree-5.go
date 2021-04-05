package day6

const END_OF_WORLD = '#'

type Trie struct {
	storage map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	var root = Trie{
		make(map[byte]*Trie),
	}
	return root
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	charSet := []byte(word)
	for _, v := range charSet {
		if this.storage[v] == nil {
			newTrie := Constructor()
			this.storage[v] = &newTrie
		}
		this = this.storage[v]
	}
	this.storage[END_OF_WORLD] = &Trie{nil}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	charSet := []byte(word)
	for _, v := range charSet {
		if this.storage[v] == nil {
			return false
		} else {
			this = this.storage[v]
		}
	}
	return this.storage[END_OF_WORLD] != nil
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	charSet := []byte(prefix)
	for _, v := range charSet {
		if this.storage[v] == nil {
			return false
		} else {
			this = this.storage[v]
		}
	}
	return true
}
