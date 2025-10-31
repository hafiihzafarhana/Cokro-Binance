package arraydata

import "strings"

func SplitAndTrim(s, sep string) []string {
    parts := strings.Split(s, sep)
    for i := range parts {
        parts[i] = strings.TrimSpace(parts[i])
    }
    return parts
}