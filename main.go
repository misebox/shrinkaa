package main

import "bufio"
import "fmt"
import "math"
import "os"
import "strconv"

func ReadFile(fileName string) []string {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	var lines []string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func MakeItRunes(lines []string) [][]rune {
	maxLen := 0
	for _, line := range lines {
		length := len(line)
		if maxLen < length {
			maxLen = length
		}
	}
	nLines := len(lines)
	runes := make([][]rune, nLines)
	for y, line := range lines {
		runes[y] = make([]rune, maxLen)
		for x := 0; x < maxLen; x++ {
			if x < len(line) {
				runes[y][x] = rune(line[x])
			} else {
				runes[y][x] = ' '
			}
		}
	}
	return runes
}

func Shrink(runes [][]rune, ratio float64) [][]rune {
	w := len(runes[0])
	h := len(runes)
	nw, nh := int(math.Ceil(float64(w)*ratio)), int(math.Ceil(float64(h)*ratio))
	nRunes := make([][]rune, nh, nh)
	for y := 0; y < nh; y++ {
		nRunes[y] = make([]rune, nw, nw)
	}
	ConvPos := func(dst, src [][]rune, x, y int, ratio float64) {
		nx, ny := int(float64(x)*ratio), int(float64(y)*ratio)
		dst[ny][nx] = src[y][x]
	}
	hw := w / 2
	hh := h / 2
	for i := 0; i < hh; i++ {
		for j := 0; j < hw; j++ {
			ConvPos(nRunes, runes, hw+j, hh+i, ratio)
			ConvPos(nRunes, runes, hw+j, hh-1-i, ratio)
			ConvPos(nRunes, runes, hw-1-j, hh-1-i, ratio)
			ConvPos(nRunes, runes, hw-1-j, hh+i, ratio)
		}
	}
	for y, line := range runes {
		for x, r := range line {
			nx, ny := int(float64(x)*ratio), int(float64(y)*ratio)
			nRunes[ny][nx] = r
		}
	}
	return nRunes
}

func Show(runes [][]rune) {
	for _, line := range runes {
		fmt.Println(string(line))
	}
}

func main() {
	if len(os.Args) < 3 {
		panic("Target filename is required")
	}
	fn := os.Args[1]
	per, err := strconv.ParseInt(os.Args[2], 10, 32)
	if err != nil {
		panic("Parse error")
	}
	lines := ReadFile(fn)
	runes := MakeItRunes(lines)

	ratio := 1.0 - float64(100-per)/100.0
	nRunes := Shrink(runes, ratio)
	Show(nRunes)
}
