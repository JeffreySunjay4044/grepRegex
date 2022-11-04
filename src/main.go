package main

import (
	"fmt"
	"io/ioutil"
	_ "io/ioutil"
	"regexp"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Error in resolving ", err)
	}
	findStringsMatchedOuter := regexp.MustCompile(`\[dev\]\ndefaultFactoryAddress = \".*\"\n`)

	matchedOuter := findStringsMatchedOuter.FindAllSubmatch(bytes, -1)
	for i := 0; i < len(matchedOuter); i++ {
		innerRegex := regexp.MustCompile(`defaultFactoryAddress = .*`)
		var innerSentence string
		for j := 0; j < len(matchedOuter[i]); j++ {
			innerSentence += string(matchedOuter[i][j])
		}
		matchedInner := innerRegex.FindAllSubmatch([]byte(innerSentence), -1)
		tempResult := matchedInner[0]
		var sentence string
		for j := 0; j < len(tempResult); j++ {
			sentence += string(tempResult[j])
		}
		splitBySpace := strings.Split(sentence, " ")
		fmt.Println("Result ", splitBySpace[len(splitBySpace)-1])
	}
	//stringsMatechedOuter = string(stringsMatchedOuterBytes)
	//findStringsMatchedInnerRegex = regexp.MustCompile(`defaultFactoryAddress = .*`)
	//regexMatchedString := findStringsMatchedOuter.String()
	//findStringsMatchedInnerByte = regexp.MustCompile(`defaultFactoryAddress = .*`, regexMatchedString)
	//findStringsMatchedInner = findStringsMatchedInnerByte.String()
	//matchedInnerBytes = findStringsMatchedInner
	//words := strings.Split(words_str, " ")
	//var countRegex int
	//for _, word := range words {
	//
	//	found, err := regexp.MatchString(`"\[$NETWORK\]\ndefaultFactoryAddress = \".*\"\n"`, word)
	//
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	//addr="$(grep -Pzo "\[$NETWORK\]\ndefaultFactoryAddress = \".*\"\n" ../scripts/generated/factoryState | grep -a "defaultFactoryAddress = .*" | awk '{print $NF}')"
	//
	//	if found {
	//		countRegex++
	//		fmt.Printf("%s matches\n", word)
	//	} else {
	//		fmt.Printf("%s does not match\n", word)
	//	}
	//}

}
