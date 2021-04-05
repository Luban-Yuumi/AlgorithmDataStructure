package main

type Trie struct {
	isEnd   bool
	nextMap map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		nextMap: map[rune]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if _, ok := this.nextMap[v]; !ok {
			this.nextMap[v] = &Trie{
				nextMap: map[rune]*Trie{},
			}
		}
		this = this.nextMap[v]
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if _, ok := this.nextMap[v]; !ok {
			return false
		}
		this = this.nextMap[v]
	}
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if _, ok := this.nextMap[v]; !ok {
			return false
		}
		this = this.nextMap[v]
	}
	return true
}
