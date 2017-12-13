package plugin

import "strings"

func Prefix(s string) (string, bool) {
	if s != "" {
		if n := strings.Index(s, "."); n > -1 {
			return s[0:n], true
		}
	}

	return "", false
}

func Suffix(s string) (string, bool) {
	if s != "" {
		if n := strings.LastIndex(s, "."); n > -1 {
            return s[n+1:], true
		}
	}

	return "", false
}

func Name(name string) (string, bool) {
	if name != "" {
		return Flag + "." + name, true
	}
	return "", false
}

func Domain(key, name string) (string, bool) {
	if key != "" && name != "" {
		return Flag + "." + key + "." + name, true
	}
	return "", false
}

func Check(name, s string) (string, bool) {
	if s != "" {
		if prefix, ok := Prefix(s); ok {
			if name == prefix {
				return s, true
			}
		}

		return "", false
	}

	return "", false
}
