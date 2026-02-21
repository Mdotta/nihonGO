package stack

type StackElement[T any] struct {
	next  *StackElement[T]
	Value T
}

type Stack[T any] struct {
	current *StackElement[T]
}

func (s *Stack[T]) Push(value T) {
	s.current = &StackElement[T]{
		next:  s.current,
		Value: value,
	}
}

func (s *Stack[T]) Pop() *StackElement[T] {
	this := s.current
	if this != nil {
		s.current = this.next
	}
	return this
}

func (s *Stack[T]) Peek() *StackElement[T] {
	return s.current
}
