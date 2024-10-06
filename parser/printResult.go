package parser

import (
	"fmt"
)

func (s *Settings) PrintResult(html, separator string) {
	for _, res := range s.ReadObject(html) {
		fmt.Print(res)
		fmt.Print(separator)
	}
}
