package stringutil

func ContainStr(s string, arrayS []string) bool {
	for _, text := range arrayS {
		if text == s {
			return true
		}
	}

	return false
}
