

func isMatch(s string, p string) bool {
    pattern := regexp.MustCompile("^" + p + "$")
    if pattern.MatchString(s) {
        return true
    }
    return false
}

