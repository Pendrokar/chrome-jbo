package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"sort"

	"regexp"

	"github.com/BenLubar/jbo/jbovlaste"
)

func main() {
	var dict jbovlaste.Dictionary
	err := xml.NewDecoder(os.Stdin).Decode(&dict)
	if err != nil {
		panic(err)
	}

	var selmaho []string
	selmahoValsi := make(map[string][]string)

	re := regexp.MustCompile(`\A[A-Zh]+`)

	for i := range dict.Direction[0].Valsi {
		valsi := &dict.Direction[0].Valsi[i]

		if valsi.Type == "cmavo" || valsi.Type == "experimental cmavo" {
			s := re.FindString(valsi.Selmaho)
			if len(selmahoValsi[s]) == 0 {
				selmaho = append(selmaho, s)
			}
			selmahoValsi[s] = append(selmahoValsi[s], valsi.Word)
		}
	}

	sort.Strings(selmaho)

	for _, s := range selmaho {
		fmt.Print(s, " = &cmavo expr:(")
		for i, v := range selmahoValsi[s] {
			if i != 0 {
				fmt.Print(" /")
			}
			for j := range v {
				if v[j] == '\'' {
					fmt.Print(" h")
				} else {
					fmt.Print(" ", string(v[j]))
				}
			}
		}
		fmt.Print(" ) &post_word {return [\"", s, "\", _join(expr)];}\n\n")
	}
}