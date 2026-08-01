package internal

func StartsWithInvalidNameChar(name string) bool {
	if len(name) == 0 {
		return false
	}
	return name[0] == '-' || name[0] == '\''
}

func EndsWithInvalidNameChar(name string) bool {
	if len(name) == 0 {
		return false
	}
	return name[len(name)-1] == '-' || name[len(name)-1] == '\''
}

func IsNumericOnly(username string) bool {
	if len(username) == 0 {
		return false
	}
	for i := 0; i < len(username); i++ {
		if username[i] < '0' || username[i] > '9' {
			return false
		}
	}
	return true
}
