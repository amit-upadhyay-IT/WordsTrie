package wordstrie

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

type SearchResult struct {
	matchState     MatchState
	searchResult   []string
	matchingFactor int
}

func getSearchResultInstance() SearchResult {
	return SearchResult{
		matchState:     UNKNOWN,
		searchResult:   make([]string, 1),
		matchingFactor: 0,
	}
}

func (sr *SearchResult) GetSearchResult() []string {
	return sr.searchResult
}

func (sr *SearchResult) GetMatchState() MatchState {
	return sr.matchState
}

func (sr *SearchResult) GetMatchingFactor() int {
	return sr.matchingFactor
}
