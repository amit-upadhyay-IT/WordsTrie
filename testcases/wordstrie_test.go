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


func TestPartialMatch_01(t *testing.T) {
	wordsTrie := wordstrie.GetInstance()
	insertionData := getKeysToInsert()
	wordsTrie.Insert(insertionData)
	testCaseDic := getExpectedPartialTestValues()
	for key, val := range testCaseDic {
		searchResult := wordsTrie.Search(key)
		matchingFactor := searchResult.GetMatchingFactor()
		failureAssertionForPartialMatch(key, val, insertionData, matchingFactor, searchResult, t)
	}
}

func failureAssertionForPartialMatch(key string, val []string, insertionData []string,
		matchingFactor int, res wordstrie.SearchResult, t *testing.T) {

	if len(strings.Fields(key)) != matchingFactor {
		t.Error(
			"For", key,
			"expected MatchingFactor", len(strings.Fields(key)),
			"got", res.GetMatchingFactor(),
		)
	}
	if contains(insertionData, key) {
		if res.GetMatchState() != 0 {
			t.Error(
				"For", key,
				"expected MatchState", 0,
				"got", res.GetMatchState(),
			)
		}
	} else {
		if res.GetMatchState() != 1 {
			t.Error(
				"For", key,
				"expected MatchState", 1,
				"got", res.GetMatchState(),
			)
		}
	}
	if !isListsEqual(res.GetSearchResult(), val) {
		t.Error(
			"For", key,
			"expected SearchResult", val,
			"got", res.GetSearchResult(),
		)
	}
}

func isListsEqual(lis1, lis2 []string) bool {

	if len(lis1) != len(lis2) {
		return false
	}

	for _, item := range lis1 {
		if !contains(lis2, item) {
			return false
		}
	}

	//for _, item := range lis2 {
	//	if !contains(lis1, item) {
	//		return false
	//	}
	//}
	return true;
}

func contains(lis []string, keyword string) bool {
	for _, v := range lis {
		if keyword == v {
			return true
		}
	}
	return false
}

func getExpectedPartialTestValues() map[string][]string {
	companyNames := getKeysToInsert()
	dic := make(map[string][]string)
	for _, company := range companyNames {
		words := strings.Fields(company)
		toMatch := make([]string, 0)
		for _, word := range words {
			toMatch = append(toMatch, word)
			matchingComp := getMatchIfAvailable(companyNames, strings.Join(toMatch, " "))
			dic[strings.Join(toMatch, " ")] = matchingComp
		}
	}
	return dic
}

func getMatchIfAvailable(companyNames []string, word string) []string {

	matchingCompanies  := make([]string, 0)
	for _, company := range companyNames {
		if strings.HasPrefix(company, word) && isMatchFinal(company, word) {
			matchingCompanies = append(matchingCompanies, company)
		}
	}
	return matchingCompanies
}

func isMatchFinal(company, word string) bool {
	begIndex := strings.Index(company, word)
	if begIndex != -1 {
		finalIndex := begIndex + len(word)
		if finalIndex < len(company) {  // Nandani And Arman, Nandan, finalIndex = 6, company[finalIndex] should be a space
			if rune(company[finalIndex]) == ' ' {
				return true
			}
		} else if company == word {
			return true
		}
	}
	return false
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
	return removeMulSpaces(getKeysFromMap(companyMap))
}

func getKeysFromMap(dic map[string]string) []string {
	result := make([]string, 0)
	for k := range dic {
		result = append(result, k)
	}
	return result
}

func removeMulSpaces(lis []string) []string {
	res := make([]string, 0)
	for _, item := range lis {
		res = append(res, strings.Join(strings.Fields(item), " "))
	}
	return res
}

