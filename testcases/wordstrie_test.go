package testcases

import (
	"../wordstrie"
	"encoding/json"
	"github.com/amit-upadhyay-it/goutils/io"
	"path/filepath"
	"strings"
	"testing"
)

// covers all possibilities for the input file
func TestPerfectMatch_01(t *testing.T) {
	keywordsSet := getKeysToInsert()
	wordsTrie := wordstrie.GetInstance()
	wordsTrie.Insert(keywordsSet)

	for _, item := range keywordsSet {
		if item == "" {  // empty string can't be searched, so ignoring such occurring
			continue
		}
		res := wordsTrie.Search(item)
		failureAssertionForPerfectMatch(res, item, t)
	}
}

func TestPerfectMatch_02(t *testing.T) {
	keywordsSet := getKeysToInsert()
	wordsTrie := wordstrie.GetInstance()
	wordsTrie.Insert(keywordsSet)

	res := wordsTrie.Search("Avenue Supermarts Limited")
	failureAssertionForPerfectMatch(res, "Avenue Supermarts Limited", t)
}

func failureAssertionForPerfectMatch(res wordstrie.SearchResult, item string, t *testing.T) {
	if res.GetMatchState() != 0 {
		t.Error(
			"For", item,
			"expected MatchState", 0,
			"got", res.GetMatchState(),
		)
	}
	if len(res.GetSearchResult()) != 1 && res.GetSearchResult()[0] != item {
		t.Error(
			"For", item,
			"expected SearchResult", []string{item},
			"got", res.GetSearchResult(),
		)
	}
	if res.GetMatchingFactor() != len(strings.Fields(item)) {
		t.Error(
			"For", item,
			"expected MatchingFactor", len(strings.Fields(item)),
			"got", res.GetMatchingFactor(),
		)
	}
}


// method for reading file content, converting to map and return the keys of the map for inserting into wordsTrie
// so it accepts nothing and always returns the sample input for tests
func getKeysToInsert() []string {
	var companyMap map[string]string
	abs, _ := filepath.Abs("../nse_data/name_to_symbol.json")
	content, _ := io.ReadFileBytes(abs)
	_ = json.Unmarshal(content, &companyMap)
	return getKeysFromMap(companyMap)
}

func getKeysFromMap(dic map[string]string) []string {
	result := make([]string, 0)
	for k := range dic {
		result = append(result, k)
	}
	return result
}
