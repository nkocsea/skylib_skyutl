package skyutl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func UnAccent(str string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, str)
	result = strings.Replace(result, "Đ", "D", -1)
	result = strings.Replace(result, "Ð", "D", -1)
	result = strings.Replace(result, "đ", "d", 1)
	return result
}

func LowerUnAccent(str string) string {
	return strings.ToLower(UnAccent(str))
}

func UpperUnAccent(str string) string {
	return strings.ToUpper(UnAccent(str))
}

func StringPadding(input string, padLength int, padString string, padType string) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		return input
	}

	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))

	switch padType {
	case "RIGHT":
		output = input + strings.Repeat(padString, int(repeat))
		output = output[:padLength]
	case "LEFT":
		output = strings.Repeat(padString, int(repeat)) + input
		output = output[len(output)-padLength:]
	case "BOTH":
		length := (float64(padLength - inputLength)) / float64(2)
		repeat = math.Ceil(length / float64(padStringLength))
		output = strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]
	}

	return output
}

func StringRightPaddingList(input []string, padLength []int) string {
	result := []string{}
	for i, str := range input {
		result = append(result, StringPadding(str, padLength[i], " ", "RIGHT"))
	}

	return strings.Join(result, "")
}

func JsonPrettyAny(v interface{}) string {
	in, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("JsonPrettyAny error: %v", err)
	}

	var out bytes.Buffer
	err = json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return string(in)
	}
	return out.String()
}

func WildCardLike(query string) string {
	return WildCardLikeSensitive(query, false)
}

func WildCardLikeSensitive(query string, sensitive bool) string {
	if query == "" {
		return "%"
	}
	if !sensitive {
		query = strings.ToUpper(query)
	}
	query = strings.Replace(query, "\\", "\\\\", -1)
	query = strings.Replace(query, "%", "\\%", -1)
	query = strings.Replace(query, "_", "\\_", -1)

	return "%" + query + "%"
}

func WildCardFull(query string) string {
	return WildCardFullSensitive(query, false)
}

func WildCardFullSensitive(query string, sensitive bool) string {
	if query == "" {
		return "%"
	}
	if !sensitive {
		query = strings.ToUpper(query)
	}
	query = strings.Replace(query, "\\", "\\\\", -1)
	query = strings.Replace(query, "%", "\\%", -1)
	query = strings.Replace(query, "_", "\\_", -1)

	query = strings.ReplaceAll(query, "\\s+", " ")
	query = strings.TrimSpace(query)

	arr := strings.Split(query, " ")
	query = "%"
	for _, q := range arr {
		query += q + "%"
	}

	return query
}
