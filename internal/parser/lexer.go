package parser

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	LineEnd       = '\n'
	CommentBegin  = '#'
	InstructBegin = '@'
	ArgBegin      = '('
	ArgEnd        = ')'
	BlockBegin    = '{'
	BlockEnd      = '}'
)

type Lexer struct {
	input  []byte // 原始数据
	len    int    // 总长度
	line   int    // 当前正在处理的行
	end    int    // 已处理字的偏移量
	offset int    // 即将处理的开始位置
	chars  rune   // 最近解析的字符
	width  int    // 最近解析的字符宽度
	tree   *Tree
	err    error
}

func NewLexer(input []byte) *Lexer {
	return &Lexer{input: input, len: len(input), line: 1}
}

// Parse 解析
// @Instruct
// @Instruct(Arg1, Arg2, ..., ArgN)
// @Instruct(Arg1, Arg2, ..., ArgN) Expression(Arg1, Arg2, ..., ArgN)
// @Instruct { ... SubBlock ... }
// @Instruct(Arg1, Arg2, ..., ArgN) { ... SubBlock ... }
// @Instruct(Arg1, Arg2, ..., ArgN) Expression(Arg1, Arg2, ..., ArgN) { ... SubBlock ... }
// @Instruct(Arg1, Arg2, ..., ArgN) Expression(Arg1, Arg2, ..., ArgN) { ... SubBlock ... }
// @Instruct(Arg1, Arg2, ..., ArgN) Expression(Arg1, Arg2, ..., ArgN) { ... SubBlock ... }
func (x *Lexer) Parse() (*Tree, error) {
	x.tree = new(Tree)
	x.parse(true, x.instruct)

	return x.tree, x.err
}

func (x *Lexer) parse(skip bool, fn func() bool) bool {
	for {
		if skip {
			if x.skip(nil) {
				x.mark()
			} else {
				return false
			}
		}
		if x.next() && fn() {
			continue
		}
		return false
	}
}

func (x *Lexer) instruct() bool {
	if x.chars != InstructBegin {
		x.err = fmt.Errorf("unexpected character at position %d (line %d)", x.offset, x.line)
		return false
	}

	if !x.parse(false, x.letter) {
		return false
	}
	switch instruct := x.get(strings.ToLower); instruct {
	case "@types":
	case "@server":

	default:
		if ok, arg := x.arg(); ok {
			if instruct == "@include" {

			} else {
				_ = &AttributeNode{Name: instruct, Value: arg}
			}
		} else {
			return false
		}
	}
	return true
}

func (x *Lexer) arg() (ok bool, value string) {
	var begin bool
	ok = x.skip(func() bool {
		switch x.chars {
		case ArgBegin:
			x.mark()
		case ArgEnd:
			value = x.get(strings.TrimSpace)
		default:
			if !begin {
				x.err = fmt.Errorf("unexpected character at position %d (line %d)", x.offset, x.line)
				return false
			}
		}
		return true
	})
	return
}

func (x *Lexer) letter() bool {
	return unicode.IsLetter(x.chars)
}

func (x *Lexer) next() bool {
	if x.err != nil || x.offset >= len(x.input) {
		return false
	}
	if x.input[x.offset] == LineEnd {
		x.line++
	}
	r, w := utf8.DecodeRune(x.input[x.offset:])
	if r == utf8.RuneError {
		x.err = fmt.Errorf("invalid utf-8 byte at position %d (line %d)", x.offset, x.line)
		return false
	}
	x.chars = r
	x.width = w
	x.offset += w
	return true
}

// Ignore all comments and space text.
func (x *Lexer) skip(fn func() bool) bool {
	var comment bool
	for x.next() {
		if comment {
			// If the line has ended, the comment is ended.
			// We only support line comment.
			if x.chars == LineEnd {
				comment = false
			}
			// Ignore all comment text.
			continue
		}
		switch x.chars {
		case CommentBegin:
			// Ignore all comment start characters.
			comment = true
		default:
			// Ignore all space text.
			if !unicode.IsSpace(x.chars) {
				if fn != nil {
					return fn()
				}
				return true
			}
		}
	}
	return false
}

func (x *Lexer) mark() {
	x.end = x.offset - x.width
}

func (x *Lexer) get(fn ...func(string) string) string {
	s := string(x.input[x.end : x.offset-x.width])
	for i, j := 0, len(fn); i < j; i++ {
		s = fn[i](s)
	}
	return s
}
