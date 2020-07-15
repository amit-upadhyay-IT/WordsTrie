package wordstrie

import (
	"strings"
	"sync"
)

type wordsTrie struct {
	root* TrieNode
}

var singletonInstance *wordsTrie = nil
var once sync.Once

func GetSingletonInstance() *wordsTrie {
	once.Do(func() {
		singletonInstance = &wordsTrie{
			&TrieNode{
				childNode:make(map[string]*TrieNode),
				level: 0,
				terminatingCount:0,
			},
		}
	})
	return singletonInstance
}

func GetInstance() *wordsTrie {
	return &wordsTrie {
		&TrieNode {
			childNode:make(map[string]*TrieNode),
			level: 0,
			terminatingCount:0,
		},
	}
}

// [larsen and turbo , larsen and turbo infotech, ICICI]
func (wt *wordsTrie) Insert(sentencesList []string)  {
	for _, sentence := range sentencesList {

		words := strings.Fields(sentence)

		localRoot := wt.root

		for _, word := range words {

			// TODO: this can be optimized by not getting a new trie node always, as this isn't required in case
			// trie has already a node with the current key
			nextNode := GetTrieNode()

			if val, ok := localRoot.childNode[word]; ok {
				nextNode = val
			} else {
				nextNode.level = localRoot.level + 1
				localRoot.childNode[word] = nextNode
			}
			localRoot = nextNode
		}

		localRoot.terminatingCount = localRoot.terminatingCount + 1
	}
}


func (wt *wordsTrie) Search(keyword string) SearchResult {
	var result = getSearchResultInstance()

	words := strings.Fields(keyword)

	localRoot := wt.root
	for _, word := range words {

		if val, ok := localRoot.childNode[word]; ok {
			localRoot = val
		} else {
			break
		}
	}

	result.matchState = getMatchState(words, localRoot)
	result.matchingFactor = getMatchingFactor(words, result.matchState, localRoot)
	result.searchResult = getSearchResult(words, result.matchState, localRoot)

	return result
}


func getMatchState(words []string, localRoot *TrieNode) MatchState {
	var result = UNKNOWN
	if localRoot.level == 0 {
		result = NO_MATCH
	} else {
		if localRoot.terminatingCount > 0 {
			result = PERFECT_MATCH
		} else {
			if len(words) == localRoot.level {
				result = PARTIAL_MATCH
			} else if len(words) > localRoot.level {
				result = DOUBTFUL_MATCH
			}
		}
	}

	// UNKNOWN state shouldn't be ideally possible, so panic if unknown found
	if result == UNKNOWN {
		panic("NOT possible to have len(words) < localRoot.level, se panicking explicitly")
	}

	return result
}

func getMatchingFactor(words []string, state MatchState, endNode *TrieNode) int {
	result := 0
	if state == PERFECT_MATCH || state == PARTIAL_MATCH {
		result = len(words)
	} else if state == DOUBTFUL_MATCH {
		result = endNode.level
	}
	return result
}

/*
 NOTE that in case of PERFECT_MATCH, we don't wanna traverse down the Trie to get more possible matches, coz its
 already a perfect match and giving more possibilities of matched items would be a semantics error
 */
func getSearchResult(words []string, state MatchState, endNode *TrieNode) []string {
	result := make([]string, 0)

	if state == PERFECT_MATCH {
		 result = append(result, strings.Join(words, " "))
	} else if state == PARTIAL_MATCH {
		result = getPossibleSearchResult(words, endNode)
	} else if state == DOUBTFUL_MATCH {
		result = append(result, strings.Join(words, " "))
	} else {
		result = nil
	}

	return result
}

func getPossibleSearchResult(words []string, endNode *TrieNode) []string {
	result := make([]string, 0)

	firstHalf := strings.Join(words, " ")

	downwardPossibilities := traverseDown(endNode)
	for _, item := range downwardPossibilities {
		secondHalf := strings.Join(item, " ")
		result = append(result, firstHalf + " " + secondHalf)
	}

	return result
}

/*
	traverses down from the current node and returns the list terminating nodes
	eg:
	word trie has: ["larsen and turbo pvt ltd", "larsen and turbo infotech pvt ltd", "Ruchi Soya"]
	and search keyword is "larsen"
	then, its a case of partial match, and so we would
	return: []string{"larsen and turbo pvt ltd", "larsen and turbo infotech pvt ltd"}
 */


type ResultStrings struct {
	result [][]string
}

func traverseDown(currentNode *TrieNode) [][]string {

	result := &ResultStrings{result: make([][]string, 0)}
	wordsList := make([]string, 0)
	traverse(currentNode, wordsList, result)

	return result.result
}

func traverse(node *TrieNode, words []string, res *ResultStrings) {

	if node.terminatingCount > 0 {
		//finalResult = append(finalResult, deepCopyList(words))
		res.result = append(res.result, deepCopyList(words))
		if len(node.childNode) > 0 {
			for key, val := range node.childNode {
				words = append(words, key)
				traverse(val, words, res)
				words = words[:len(words)-1]
			}
		}
	} else {
		for key, val := range node.childNode {
			words = append(words, key)
			traverse(val, words, res)
			words = words[:len(words)-1]
		}
	}
}

func deepCopyList(list []string)  []string {
	res := make([]string, len(list))
	copy(res, list)
	return res
}
