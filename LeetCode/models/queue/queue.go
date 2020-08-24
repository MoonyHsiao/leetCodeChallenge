package queue

type StringQueue []string

func (s *StringQueue) IsEmpty() bool {
	return len(*s) == 0
}

func (s *StringQueue) Push(str string) {
	*s = append(*s, str)
}

func (s *StringQueue) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return "", false
	}
	element := (*s)[0]
	(*s) = (*s)[1:]
	return element, true
}

func (s *StringQueue) GetLen() int {
	return len(*s)
}

func (s *StringQueue) PeekFront() interface{} {
	element := ""
	if s.IsEmpty() {
		return element
	}
	element = (*s)[0]
	return element
}

func (s *StringQueue) PeekBack() interface{} {
	element := ""
	if s.IsEmpty() {
		return element
	}
	index := s.GetLen() - 1
	element = (*s)[index]
	return element
}

func (s *StringQueue) Clear() {
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
