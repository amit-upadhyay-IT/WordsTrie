package wordstrie

type TrieNode struct {
	childNode map[string]*TrieNode
	level int  // will help to find matching factor
	terminatingCount int32
}

func GetTrieNode() *TrieNode {
	return &TrieNode{
		childNode:        make(map[string]*TrieNode),
		level:            0,
		terminatingCount: 0,
	}
}
