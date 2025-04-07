package gotriesearch

import "strings"

type (
	// TrieNode represents a node in the trie.
	trieNode struct {
		// children is a map of child nodes, where the key is the character.
		children map[byte]*trieNode
		// endOfWord indicates if the node represents the end of a word.
		endOfWord bool
	}

	// Trie represents the trie data structure.
	trie struct {
		// root is the root node of the trie.
		root *trieNode
		// size is the number of words in the trie.
		size int
		// caseInsensitive indicates if the trie is case insensitive.
		caseInsensitive bool
	}

	TrieSearch interface {
		// Insert adds a word to the trie.
		Insert(word string)
		// PrefixSearch checks if there is any word in the trie that starts with the given prefix.
		PrefixSearch(word string) []string
		// Exist return if the word exists in the trie.
		Exist(word string) (isExist bool)
		// Delete removes a word from the trie.
		Delete(word string) (deleted bool)
		// Size returns the number of words in the trie.
		Size() int
	}
)

// NewTrie initializes a new Trie
func NewTrie() TrieSearch {
	return &trie{root: newTrieNode()}
}

// NewCaseInsensitiveTrie initializes a new Trie
func NewCaseInsensitiveTrie() TrieSearch {
	return &trie{root: newTrieNode(), caseInsensitive: true}
}

// Size returns the number of words in the trie
func (t trie) Size() int {
	return t.size
}

func (t *trie) Exist(word string) bool {
	node := t.root
	for _, char := range word {
		if _, ok := node.getChildren(char); !ok {
			return false
		}
		node, _ = node.getChildren(char)
	}
	return node.isEndOfWord()
}

// Insert adds a word to the trie
func (t *trie) Insert(word string) {
	// If the word is empty, we can't insert it
	if word == "" {
		return
	}

	// If the trie is case insensitive, we need to convert the word to lower case
	if t.caseInsensitive {
		word = strings.ToLower(word)
	}

	// If the word already exists, we don't need to insert it again
	// and we can return early.
	if t.Exist(word) {
		return
	}

	// If the word doesn't exist, we need to insert it
	node := t.root
	for _, char := range word {
		child, exist := node.getChildren(char)
		if !exist {
			child = newTrieNode()
			node.setChildren(char, child)
		}
		node = child
	}
	node.setAsEndOfWord()
	t.size++
}

// PrefixSearch returns all words in the trie that start with the given prefix
// If no words are found, it returns an empty slice.
func (t *trie) PrefixSearch(word string) []string {
	if word == "" {
		return nil
	}

	// If the trie is case insensitive, we need to convert the word to lower case
	if t.caseInsensitive {
		word = strings.ToLower(word)
	}

	// Traverse the trie to find the node corresponding to the last character of the prefix
	// If the prefix doesn't exist, return an empty slice
	node := t.root
	for _, char := range word {
		if _, ok := node.getChildren(char); !ok {
			return nil
		}
		node, _ = node.getChildren(char)
	}

	var result []string
	t.collectWords(node, word, &result)

	return result
}

func (t *trie) collectWords(node *trieNode, prefix string, result *[]string) {
	if node.isEndOfWord() {
		*result = append(*result, prefix)
	}
	for ch, child := range node.getChildrens() {
		t.collectWords(child, prefix+string(ch), result)
	}
}

// Delete removes a word from the Trie
func (t *trie) Delete(word string) bool {
	// If the word doesn't exist, we can't delete it
	if word == "" {
		return false
	}

	// If the trie is case insensitive, we need to convert the word to lower case
	if t.caseInsensitive {
		word = strings.ToLower(word)
	}

	// the delete worker function already handles this case but we need this to reduce the size
	if !t.Exist(word) {
		return false
	}
	// If the word exists, we need to delete it
	// and we can return true.
	t.deleteWorker(t.root, word, 0)
	t.size--
	return true
}

// deleteWorker is a helper function that deletes a word from the Trie
func (t *trie) deleteWorker(node *trieNode, word string, depth int) bool {
	if node == nil {
		return false
	}

	// End of word reached
	if depth == len(word) {
		if node.isEndOfWord() {
			node.setAsNotEndOfWord()
			// If node has no children, it can be pruned
			return !node.hasChildren()
		}
		return false
	}

	char := rune(word[depth])
	child, exists := node.getChildren(char)
	if !exists {
		return false
	}

	// Recursively delete child
	needDeleteChild := t.deleteWorker(child, word, depth+1)
	if needDeleteChild {
		node.deleteChildren(char)
		return !node.hasChildren() && !node.isEndOfWord()
	}

	return false
}

// newTrieNode creates a new trie node
func newTrieNode() *trieNode {
	return &trieNode{children: make(map[byte]*trieNode)}
}

// isEndOfWord checks if the node is the end of a word
func (tn trieNode) isEndOfWord() bool {
	return tn.endOfWord
}

// setAsEndOfWord sets the endOfWord flag to true
func (tn *trieNode) setAsEndOfWord() {
	tn.endOfWord = true
}

// setAsNotEndOfWord sets the endOfWord flag to false
func (tn *trieNode) setAsNotEndOfWord() {
	tn.endOfWord = false
}

// hasChildren checks if the node has any children
func (tn trieNode) hasChildren() bool {
	return len(tn.children) > 0
}

// getChildren returns the child node for a given character
func (tn trieNode) getChildren(char rune) (*trieNode, bool) {
	if tn.children == nil {
		return nil, false
	}
	child, exist := tn.children[byte(char)]
	return child, exist
}

// getChildrens return all childrens of the node
func (tn trieNode) getChildrens() map[byte]*trieNode {
	return tn.children
}

// setChildren sets the children of the node
func (tn *trieNode) setChildren(char rune, child *trieNode) {
	tn.children[byte(char)] = child
}

// deleteChildren deletes a child node from the node's children
func (tn *trieNode) deleteChildren(char rune) {
	delete(tn.children, byte(char))
}
