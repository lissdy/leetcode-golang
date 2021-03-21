package main

func main() {

}

type Trie struct {
	letter   string
	children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		"/", [26]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		index := v - 'a'
		if this.children[index] == nil {
			this.children[index] = &Trie{
				string(v), [26]*Trie{},
			}
		}
		this = this.children[index]
	}
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
	return true
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
