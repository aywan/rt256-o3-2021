package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

var stack *IntStack = NewStack()

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

func checkBlock() error {
	var n int
	scanf("%d\n", &n)

	list := readTags(n)

	state, ids := isValid(list, -1)
	if state {
		return nil
	}

	if len(ids) > 1 {
		return incorrectError()
	} else {
		return almostError(list[ids[0]])
	}

	for _, i := range ids {
		if s, _ := isValid(list, i); s {
			return almostError(list[i])
		}
	}

	return incorrectError()
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

func isValid(list TagList, skipId int) (bool, []int) {

	errorIds := make([]int, 0, len(list))
	state := true
	for _, tag := range list {
		if tag.Id == skipId {
			continue
		}
		if !tag.IsClose {
			stack.push(tag.Id)
			continue
		}

		if stack.length() > 0 {
			topId := stack.getTop()

			if list[topId].Value == tag.Value {
				stack.pop()
				continue
			}
		}

		errorIds = append(errorIds, tag.Id)
		state = false
	}

	for stack.length() > 0 {
		tagId := stack.pop()
		errorIds = append(errorIds, tagId)
		state = false
	}

	return state, errorIds
}

func readTags(n int) TagList {
	list := make(TagList, n)
	var s string
	for i := 0; i < n; i++ {
		scanf("%s\n", &s)

		isClose := false
		if s[1] == '/' {
			isClose = true
			s = s[2 : len(s)-1]
		} else {
			s = s[1 : len(s)-1]
		}

		list[i] = Tag{
			Value:   strings.ToUpper(s),
			IsClose: isClose,
			Id:      i,
		}
	}
	return list
}

type IntStack struct {
	ids []int
}

type TagList []Tag

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
