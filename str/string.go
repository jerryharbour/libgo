package str

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type StrLineRange struct {
	StartLine int
	EndLine   int
}

func StringEqInAarry(str string, arrStr []string) bool {
	for _, as := range arrStr {
		if strings.EqualFold(str, as) {
			return true
		}
	}

	return false
}

func StringLeftEq(src string, comp string) bool {
	if len(comp) > len(src) {
		return false
	}

	str := src[:len(comp)]
	return strings.EqualFold(str, comp)
}

func StringRightEq(src string, comp string) bool {
	srcLen := len(src)
	compLen := len(comp)
	if compLen > srcLen {
		return false
	}

	str := src[srcLen-compLen : srcLen]
	return strings.EqualFold(str, comp)
}

func StringContains(src string, comp string) bool {
	return strings.Contains(src, comp)
}

func StringContainsInAarry(str string, arrStr []string) bool {
	for _, as := range arrStr {
		if strings.Contains(str, as) {
			return true
		}
	}
	return false
}

func EndsWith(src, comp string) bool {
	srcLen := len(src)
	compLen := len(comp)
	if compLen > srcLen {
		return false
	}

	str := src[srcLen-compLen : srcLen]
	return str == comp
}

func EndsWithInArray(src string, arrComp []string) bool {
	for _, comp := range arrComp {
		if EndsWith(src, comp) {
			return true
		}
	}

	return false
}

func CompressStr(str string) string {
	if str == "" {
		return ""
	}

	reg := regexp.MustCompile(`[\s\p{Zs}]{1,}`)
	return reg.ReplaceAllString(str, "")
}

func SubStr(str string, start int, end int) string {
	strLen := len(str)
	if start < 0 || end > strLen || start > end {
		return ""
	}

	if start == 0 && end == strLen {
		return str
	}

	return str[start:end]
}

func StrReplace(str, substr, rpsubstr string) string {
	return strings.ReplaceAll(str, substr, rpsubstr)
}

func ExtractStringInRange(str string, startLine, endLine int, isPreLineNum bool) (string, error) {
	if startLine < 0 || endLine < 0 || endLine < startLine {
		return "", fmt.Errorf("range parameter is wrong, startLine[%d], endLine[%d]", startLine, endLine)
	}

	data := []byte(str)
	scanner := bufio.NewScanner(bytes.NewReader(data))

	var buf bytes.Buffer
	currentLine := 0

	for scanner.Scan() {
		currentLine++
		if currentLine >= startLine && currentLine <= endLine {
			lineNum := ""
			if isPreLineNum {
				lineNum = IntToString(currentLine) + ":"
			}
			_, err := buf.WriteString(lineNum + scanner.Text() + "\n")
			if err != nil {
				return "", err
			}
		}

		if currentLine > endLine {
			break
		}

		if scanner.Err() != nil {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		if err == io.EOF {
			return buf.String(), nil
		}

		return "", err
	}

	return buf.String(), nil
}

func StringInRange(str string, lineRanges []*StrLineRange, isPreLineNum bool) (string, error) {
	if len(lineRanges) <= 0 {
		return "", fmt.Errorf("line ranges is empty")
	}

	data := []byte(str)
	scanner := bufio.NewScanner(bytes.NewReader(data))

	var buf bytes.Buffer
	currentLine := 0
	lrIndex := 0
	size := len(lineRanges)

	for scanner.Scan() {
		currentLine++

		if lrIndex > 0 && currentLine == lineRanges[lrIndex].StartLine {
			_, err := buf.WriteString("...\n")
			if err != nil {
				return "", err
			}
		}

		if lrIndex >= size {
			break
		}

		if currentLine >= lineRanges[lrIndex].StartLine && currentLine <= lineRanges[lrIndex].EndLine {
			lineNum := ""
			if isPreLineNum {
				lineNum = IntToString(currentLine) + ":"
			}
			_, err := buf.WriteString(lineNum + scanner.Text() + "\n")
			if err != nil {
				return "", err
			}
		}

		if currentLine > lineRanges[lrIndex].EndLine {
			lrIndex++
		}

		if currentLine >= lineRanges[size-1].EndLine {
			break
		}

		if scanner.Err() != nil {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		if err == io.EOF {
			return buf.String(), nil
		}

		return "", err
	}

	return buf.String(), nil
}

func TrimLeft(str string, cutset string) string {
	return SubStr(str, len(cutset), len(str))
}

func TrimRight(str string, cutstr string) string {
	return strings.TrimRight(str, cutstr)
}

func TrimBfLastChar(str string, ch string) string {
	idx := strings.LastIndex(str, ch)
	if idx == -1 {
		return ""
	}

	return str[idx+1:]
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

func ProcessStringByLine(content string, lineHandle func(string)) error {
	bufscanner := bufio.NewScanner(strings.NewReader(content))

	for bufscanner.Scan() {
		line := bufscanner.Text()

		lineHandle(line)
	}

	if err := bufscanner.Err(); err != nil {
		return err
	}

	return nil
}

func ToMd5(content string) string {
	cxtSum := md5.Sum([]byte(content))
	return hex.EncodeToString(cxtSum[:])
}
