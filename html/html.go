package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	var n int
	scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		err := checkBlock()
		if nil == err {
			printf("CORRECT\n")
		} else {
			printf("%s\n", err.Error())
		}
	}
}

type Tag struct {
	Value   string
	IsClose bool
	Id      int
}
type TagList []Tag

func (t Tag) same(tag Tag) bool {
	return t.Value == tag.Value
}

func (t Tag) close(tag Tag) bool {
	return t.Value == tag.Value && t.IsClose && !tag.IsClose
}

func incorrectError() error {
	return fmt.Errorf("INCORRECT")
}

func almostError(tag Tag) error {
	slash := ""
	if tag.IsClose {
		slash = "/"
	}
	return fmt.Errorf("ALMOST <%s%s>", slash, tag.Value)
}

func checkBlock() error {
	var n int
	scanf("%d\n", &n)

	list, open, closed := readTags(n)
	return validate(list, open > closed)
}

func readTags(n int) (TagList, int, int) {
	list := make(TagList, n)
	closed, open := 0, 0
	var s string
	for i := 0; i < n; i++ {
		scanf("%s\n", &s)

		isClose := false
		if s[1] == '/' {
			isClose = true
			closed++
			s = s[2 : len(s)-1]
		} else {
			s = s[1 : len(s)-1]
			open++
		}

		list[i] = Tag{
			Value:   strings.ToUpper(s),
			IsClose: isClose,
			Id:      i,
		}
	}
	return list, open, closed
}

func validate(list TagList, reverse bool) error {
	var err error
	var stack = NewStack()
	var tag Tag

	for i := 0; i < len(list); i++ {

		if reverse {
			tag = list[len(list)-i-1]
		} else {
			tag = list[i]
		}

		if reverse == tag.IsClose {
			stack.push(tag.Id)
			continue
		}

		if stack.length() > 0 {
			topId := stack.getTop()
			if list[topId].same(tag) {
				stack.pop()
				continue
			}
		}

		if err != nil {
			return incorrectError()
		}

		err = almostError(tag)
	}

	return err
}

type IntStack struct {
	ids []int
}

func NewStack() *IntStack {
	return &IntStack{
		ids: make([]int, 0),
	}
}

func (s *IntStack) push(id int) {
	s.ids = append(s.ids, id)
}

func (s *IntStack) getTop() int {
	l := len(s.ids) - 1
	return s.ids[l]
}

func (s *IntStack) pop() int {
	l := len(s.ids) - 1
	tag := s.ids[l]
	s.ids = s.ids[:l]
	return tag
}

func (s *IntStack) length() int {
	return len(s.ids)
}
