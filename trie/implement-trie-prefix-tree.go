package trie

type dict map[string]*Trie

type Trie struct {
	node  *dict
	isEnd bool
}

func Constructor() Trie {
	node := Trie{
		node: &dict{},
	}
	return node
}

func (t *Trie) Insert(word string) {
	cur := t.node
	wordLen := len(word)
	for i, v := range word {

		if (*cur)[string(v)] == nil {
			t := Constructor()
			(*cur)[string(v)] = &t
		}

		if wordLen-1 == i {
			(*cur)[string(v)].isEnd = true
		}

		cur = (*cur)[string(v)].node
	}

}

func (t *Trie) Search(word string) bool {
	root := t.node
	wordLen := len(word)

	for i, v := range word {
		curNode := (*root)[string(v)]
		if curNode != nil && wordLen-1 == i && curNode.isEnd {
			return true
		}

		if curNode == nil {
			return false
		}
		root = curNode.node

	}
	return false
}

func (t *Trie) StartsWith(prefix string) bool {
	root := t.node
	wordLen := len(prefix)

	for i, v := range prefix {
		curNode := (*root)[string(v)]
		if wordLen-1 == i && curNode != nil {
			return true
		}

		if curNode == nil {
			return false
		}
		root = curNode.node

	}
	return false
}
