package strutil


// this functions honestly makes no sense to me
// maybe later, but not now. Copied from stackoverflow
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes) - 1; i < j; i, j = i+1, j-1 {
	runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)
}
