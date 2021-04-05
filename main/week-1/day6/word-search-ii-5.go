package day6



func findWords(board [][]byte, words []string) []string {
	var (
		trie = Constructor()
		res  []string
		n    = len(board)
		m    = len(board[0])
	)
	for i := range words {
		trie.Insert(words[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			findWordsDfs(board, &trie, n, m, i, j, "", &res) //每个位置都要作为开始点
		}
	}
	return res
}

func findWordsDfs(board [][]byte, trie *Trie, n, m, i, j int, word string, res *[]string) {
	if i < 0 || j < 0 || i >= n || j >= m || board[i][j] == '@' {
		return
	}
	tmp := board[i][j]
	word = word + string(tmp)

	if !trie.StartsWith(string(tmp)) {
		return
	}
	if trie.Search(string(tmp)) {
		*res = append(*res, word)
		trie.storage[tmp].storage['#'] = nil // 防止出现重复单词
	}
	board[i][j] = '@'
	findWordsDfs(board, trie.storage[tmp], n, m, i+1, j, word, res)
	findWordsDfs(board, trie.storage[tmp], n, m, i-1, j, word, res)
	findWordsDfs(board, trie.storage[tmp], n, m, i, j+1, word, res)
	findWordsDfs(board, trie.storage[tmp], n, m, i, j-1, word, res)
	board[i][j] = tmp
	return
}
