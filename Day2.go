package main

import "strings"

func UniqueWordCount(textContent string) (map[string]uint64, []string) {
	textContent = strings.Replace(textContent,","," ",-1)
	textContent = strings.Replace(textContent, ".", " ", -1)
	textContent = strings.Replace(textContent, "\""," ", -1)
	textContent = strings.ToLower(textContent)
	words := strings.Fields(textContent)
	wordMap := make(map[string]uint64)
	for _,word := range words {
		count, present := wordMap[word]
		if present {
			count++
		} else {
			count = 1
		}
		wordMap[word] = count
	}
	var wordsList = make([]string,0)
	for key := range wordMap {
		wordsList = append(wordsList,key)
	}
	return wordMap, wordsList
}