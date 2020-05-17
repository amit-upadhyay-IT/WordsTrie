# WordsTrie

import it this way:
```go
go get https://github.com/amit-upadhyay-IT/WordsTrie.git
```

## Application:
- Search the similar stock name from a list of stock names

## How is works?
- Using Trie DS

Data structure name `WordsTrie`

```
type wordsTrie struct {
	root* TrieNode
}
```

Where TrieNode is
```
type TrieNode struct {
	childNode map[string]*TrieNode
	level int  // will help to find matching factor
	terminatingCount int32
}
```

- methods it contains:
    - `insert(list<string> sentence)` # ["Larsen & Turbo Infotech Pvt. Lmt", "HDFC BANK", "ICICI", "LARSEN & Turbo pvt. ltd"]
    - returns `void`
    
    - `search(string keyword)`
    - returns
    ```
    type SearchResult struct {
        matchState     MatchState
        searchResult   []string
        matchingFactor int
    }
    ```
  
  ### MatchState is enum 
  ```
  type MatchState int
  const (
  	PERFECT_MATCH MatchState = iota  // all words in keyword string match completely to one or  more items in WordTrie
  
  	PARTIAL_MATCH					 // just few words are matched in keyword
  
  	DOUBTFUL_MATCH					 // few words of keyword are matched and rest of the words couldn't be searched as
  									 // the termination node reached while searching, this is the case when length of
  									 // keyword is greater than the partially matching items in WordTrie, basically
  									 // this can also be considered as NO_MATCH case coz all words couldn't be found
  
  	NO_MATCH						 // No match found for any of the word in the keyword
  
  	UNKNOWN
  )
  ```

### searchResult:
- list of matching keywords

### matchingFactor:
- an integer ranging from 0 to len(searched_keyword)
