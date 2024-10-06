package parser

import (
	"regexp"
)

func (s *Settings) ReadObject(htmlData string) (output []string) {

	var (
		regExp *regexp.Regexp
	)

	if s.Attribute == "" && s.TagName == "" {
		mapa := make([]string, 0)
		return append(mapa, htmlData)
	}

	if s.Attribute == "" {
		regExp = regexp.MustCompile(`<` + s.TagName + `[^>]*?>[\s\S]*?</` + s.TagName + `>`)
	} else {
		regExp = regexp.MustCompile(`<` + s.TagName + `.*?` + s.Attribute + `\s*=\s*['"]([^\s'"]+)[\s'"]`)
	}

	result := regExp.FindAllString(htmlData, -1)

	for _, s := range result {
		output = append(output, s)
	}

	return output
}
