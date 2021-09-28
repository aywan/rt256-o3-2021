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

var err error

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

func (t Tag) same(tag Tag) bool {
	return t.Value == tag.Value
}

func (t Tag) close(tag Tag) bool {
	return t.Value == tag.Value && t.IsClose && ! tag.IsClose
}

type TagList []Tag

func checkBlock() error {
	err = nil
	var n int
	scanf("%d\n", &n)

	list := readTags(n)

	res := clearBlocks(list, 0, n)
	if nil != err && nil == res {
		return err
	}
	return res
}

func clearBlocks(list TagList, s int, e int) error {
	if s >= e {
		return nil
	}

	top := list[s]
	isTagClosed := false
	var closeTag, cur Tag

	if top.IsClose {
		err = almostError(top)
		s++
	} else {
		cnt := 1
		for j := s + 1; j < e; j++ {
			cur = list[j]
			if ! cur.same(top) {
				continue
			}

			if cur.close(top) {
				cnt --
				if cnt == 0 {
					if ! isTagClosed {
						closeTag = cur
					}
					isTagClosed = true
				} else if isTagClosed && cnt < 0 {
					closeTag = cur
					break
				}
			} else {
				cnt ++
			}
		}

		if isTagClosed {
			err2 := clearBlocks(list, s+1, closeTag.Id)
			if err != nil && err2 != nil {
				return incorrectError()
			}
			if err2 != nil {
				err = err2
			}
			s = closeTag.Id + 1
		} else {
			if err != nil {
				return incorrectError()
			}
			err = almostError(top)
			s++
		}

	}

	err2 := clearBlocks(list, s, e)
	if err != nil && err2 != nil {
		return incorrectError()
	}
	if err2 != nil {
		err = err2
	}

	return nil
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
