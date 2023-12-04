package aoclib

import "regexp"

func PatternToMap(re *regexp.Regexp, b []byte) map[string][]byte {
	matches := re.FindAllSubmatch(b, -1)
	result := make(map[string][]byte)
	if matches == nil {
		return nil
	}
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matches[0][i]
		}
	}
	return result
}
