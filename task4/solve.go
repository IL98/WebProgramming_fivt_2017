package main

import (
	"fmt"
	"regexp"
	"unicode"
)	


func RemoveEven(slice []int) []int{
	var result []int
	for  _,iter := range slice {
		if iter % 2 == 1 {
			result = append(result, iter)
		}
		
	}
	return result  
} 

func PowerGenerator(n int) func()int {
	value := n
	degree := n
	gen := func() int {
		result := value
		value  = value * degree
		return result
	}
	
	return gen 
}


func IsEqualWords(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	if word1 == word2 {
		return true
	}
	runes1 := []rune(word1)
	runes2 := []rune(word2)
	for i, _ := range word1 {
		if unicode.ToLower(runes1[i]) != unicode.ToLower(runes2[i]) {
			return false
		}  
	}
	return true	
}


func DifferentWordsCount(text string) int{
	RegEx, _ := regexp.Compile("[a-zA-Z]+")
	array := RegEx.FindAllString(text, -1)	
	var counter int = 0
	
	for i, str := range array {
		for j := 0; j < i; j++ {
			if IsEqualWords(array[j], str) {
				counter--
				break
			}
		}
		counter++
	}
	return counter
}

/*
func main() {
	gen := PowerGenerator(5)
	arr := make([]int, 10)
	for i := range arr {
		fmt.Println(gen())
		i++
	}
	
	fmt.Println(DifferentWordsCount("Hello, world!HELLO  wOrlD...12"))
}
*/
