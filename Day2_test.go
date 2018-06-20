package main

import (
	"testing"
	"io/ioutil"
)

func TestWordCount(t *testing.T) {
	filecontent, err := ioutil.ReadFile("./Day2_SampleCorpus.txt")

	if err != nil {
		t.Fatal("Unable to read from the sample data file ./Day2_SampleCorpus.txt")
	}

	var fileContentAsString = string(filecontent)
	wordCountMap,wordList := UniqueWordCount(fileContentAsString)
	if wordCountMap == nil {
		t.Logf("The word count map (1st return value) is nil")
		t.Fail()
	}
	if wordList == nil {
		t.Logf("The word count list (2nd return value) is nil")
		t.Fail()
	}
	if len(wordCountMap) != len(wordList) {
		t.Logf("Unique Word count in the count map <%d> and word list <%d> do not match.",len(wordCountMap),len(wordList))
		t.Fail()
	}

	if len(wordList) != 207 {
		t.Logf("Total number of unique words doesn't match. Expected: 207, actual: <%d>",len(wordList))
		t.Fail()
	}
	totalWordCount := uint64(0)
	for _,count := range wordCountMap {
		totalWordCount += count
	}

	if totalWordCount != 495 {
		t.Logf("Total words count doesn't match. Expected: 495, actual: <%d>",len(wordList))
		t.Fail()
	}
}
