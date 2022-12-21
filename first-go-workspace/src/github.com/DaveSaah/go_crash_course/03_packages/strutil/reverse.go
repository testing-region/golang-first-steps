package strutil


// Edit: I understand how this function works
func Reverse(s string) string {
    bytes := []byte(s)
    
    for i, j := 0, len(bytes) - 1; i < j; i, j = i+1, j-1 {
	bytes[i], bytes[j] = bytes[j], bytes[i]
    }

    return string(bytes)
}
