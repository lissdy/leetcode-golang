package main

var res []string

func main() {
}
func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}
	row := len(board)
	col := len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			backtrace(i, j, board, "", trie)
		}
	}
	return removeDuplicates(res)
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func backtrace(i int, j int, board [][]byte, sub string, trie Trie) {
	m := len(board)
	n := len(board[0])
	if i < 0 || i > m-1 || j < 0 || j > n-1 {
		return
	}
	tmp := board[i][j]
	sub += string(tmp)
	if !trie.StartsWith(sub) {
		return
	}
	if trie.Search(sub) {
		res = append(res, sub)
		return
	}
	board[i][j] = '.'
	backtrace(i-1, j, board, sub, trie)
	backtrace(i+1, j, board, sub, trie)
	backtrace(i, j-1, board, sub, trie)
	backtrace(i, j+1, board, sub, trie)
	board[i][j] = tmp
}

type Trie struct {
	letter   string
	children [26]*Trie
	isEnd    bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		"/", [26]*Trie{}, false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		index := v - 'a'
		if this.children[index] == nil {
			this.children[index] = &Trie{
				string(v), [26]*Trie{}, false,
			}
		}
		this = this.children[index]
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		index := v - 'a'
		if index < 0 { // '.'
			return false
		}
		if this.children[index] == nil {
			return false
		}
		if this.children[index].letter != string(v) {
			return false
		}
		this = this.children[index]
	}
	if this.isEnd {
		return true
	} else {
		return false
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		index := v - 'a'
		if index < 0 { // '.'
			return false
		}
		if this.children[index] == nil {
			return false
		}
		if this.children[index].letter != string(v) {
			return false
		}
		this = this.children[index]
	}
	return true
}
