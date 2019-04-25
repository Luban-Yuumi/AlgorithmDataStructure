package main

import "fmt"

type Trie struct {
	storage map[byte]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		make(map[byte]*Trie),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	charSet := []byte(word)
	for _, v := range charSet {
		if this.storage[v] == nil {
			this.storage[v] = &Trie{make(map[byte]*Trie)}
		}
		this = this.storage[v]
	}
	this.storage['#'] = &Trie{nil}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	charSet := []byte(word)
	for _, v := range charSet {
		if this.storage[v] == nil {
			return false
		}
		this = this.storage[v]
	}
	return this.storage['#'] != nil
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	charSet := []byte(prefix)
	for _, v := range charSet {
		if this.storage[v] == nil {
			return false
		}
		this = this.storage[v]
	}
	return true
}

func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 || len(board) == 0 {
		return nil
	}
	var trie = Constructor()
	var res []string
	for _, v := range words {
		trie.Insert(v)
	}
	m, n := len(board), len(board[0])
	visit := make(map[byte]int)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			_dfs(board, i, j, &trie, "", &res, visit)
		}
	}
	return res
}

func _dfs(board [][]byte, m int, n int, currentTrie *Trie, word string, res *[]string, visit map[byte]int) {
	if m >= len(board) || n >= len(board[0]) || n < 0 || m < 0 {
		return
	}
	if _, ok := visit[board[m][n]]; ok {
		return
	}
	word += string(board[m][n])
	fmt.Println(word)
	if !currentTrie.StartsWith(word) {
		return
	}
	if currentTrie.Search(word) {
		*res = append(*res, word)
	}
	visit[board[m][n]] = 1
	_dfs(board, m+1, n, currentTrie, word, res, visit)
	_dfs(board, m-1, n, currentTrie, word, res, visit)
	_dfs(board, m, n+1, currentTrie, word, res, visit)
	_dfs(board, m, n-1, currentTrie, word, res, visit)
	delete(visit, board[m][n])
}
