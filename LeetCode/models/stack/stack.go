package stack

type StringStack []string

func (s *StringStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *StringStack) Push(str string) {
	*s = append(*s, str)
}

func (s *StringStack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := s.GetLen() - 1
	element := (*s)[index]
	(*s) = (*s)[:index]
	return element, true
}

func (s *StringStack) GetLen() int {
	return len(*s)
}

func (s *StringStack) Peek() interface{} {
	element := ""
	if s.IsEmpty() {
		return element
	}
	index := s.GetLen() - 1
	element = (*s)[index]
	return element
}

func (s *StringStack) Clear() {
	*s = nil
}

// Infix 中綴
func Infix() string {
	return ""
}

// Suffix 後綴
func Suffix() string {
	return ""
}
