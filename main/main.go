package main

import (
	"../wordstrie"
	"fmt"
)


func main() {
	wordsTrie := wordstrie.GetInstance()
	companiesList := []string{"something is here", "something else is here"}
	wordsTrie.Insert(companiesList)

	//fmt.Println(wordsTrie.Search("something"))
	var res wordstrie.SearchResult = wordsTrie.Search("something")
	for _, v := range res.GetSearchResult() {
		fmt.Println(v)
	}

	//wordsTrie.Search("something wow")
	//fmt.Println("successful")
}
