package util

/*
Integer stack
*/
type IntStack []int

// IsEmpty: check if stack is empty
func (s *IntStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *IntStack) Push(str int) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *IntStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *IntStack) Top() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		return (*s)[len(*s)-1], true
	}
}

/*
String stack
*/
type StringStack []string

// IsEmpty: check if stack is empty
func (s *StringStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *StringStack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *StringStack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *StringStack) Top() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		return (*s)[len(*s)-1], true
	}
}

/*
Any stack
*/
type AnyStack []interface{}

// IsEmpty: check if stack is empty
func (s *AnyStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *AnyStack) Push(str interface{}) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *AnyStack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *AnyStack) Top() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		return (*s)[len(*s)-1], true
	}
}
