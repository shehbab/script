package main
import (
	"fmt"
	"regexp"
)
func uniqueNonEmptyElementsOf(s []string) []string {
  unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}

	return us
}

func main() {
	str1 := `Proxy Port Last Check Proxy Speed Proxy Country Anonymity 118.99.81.204
	118.99.81.204 8080 34 sec  - Tangerang Transparent 2.184.31.2 8080 58 sec 
	Iran Transparent 93.126.11.189 8080 1 min Iran - Esfahan Transparent 202.118.236.130 
	7777 1 min  - Harbin Transparent 62.201.207.9 8080 1 min Iraq Transparent`

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	fmt.Println(re.MatchString(str1)) // true

	submatchall := re.FindAllString(str1, -1)
	for l, element := range uniqueNonEmptyElementsOf(submatchall) {
		fmt.Println(l,element)
	}
}
