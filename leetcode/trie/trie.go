package trie

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func NewTrie() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	prefix := t.SearchPrefix(word)
	return prefix != nil && prefix.isEnd
}

func (t *Trie) StartWith(prefix string) bool {
	searchPrefix := t.SearchPrefix(prefix)
	return searchPrefix != nil
}
