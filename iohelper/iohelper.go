package iohelper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var bio = bufio.NewReader(os.Stdin)

//GetLine gets a line from stdin
func GetLine() string {
	b, _, _ := bio.ReadLine()
	sb := string(b)
	sb = strings.TrimSpace(sb)
	if sb == "" {
		return GetLine()
	}
	return sb
}

//GetWords gets a slice of words from stdin (finishes on line)
func GetWords() []string {
	return strings.Fields(GetLine())
}

//GetInt gets an int from stdin
func GetInt() int {
	var temp int
	fmt.Fscan(bio, &temp)
	return temp
}

//GetIntsN gets multiple ints from stdin
func GetIntsN(n int) []int {
	list := make([]int, n)
	for i := range list {
		fmt.Fscan(bio, &list[i])
	}
	return list
}

//GetInts gets all the ints in the line from stdin
func GetInts() []int {
	line := GetLine()
	fields := strings.Fields(line)
	list := make([]int, len(fields))
	for i, a := range fields {
		list[i], _ = strconv.Atoi(a)
	}
	return list
}

//GetInt64 gets int64 from stdin
func GetInt64() int64 {
	var temp int64
	fmt.Fscan(bio, &temp)
	return temp
}
