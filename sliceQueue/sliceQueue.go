package sliceQueue

//implement queue by slice

type SliceQueue struct {
	array []int
	front int //队头索引
	rear  int //队尾索引
}

//check queue is empty?
func (s *SliceQueue) IsEmpty() bool {
	return s.front == s.rear
}

//queue's size
func (s *SliceQueue) Size() int {
	return s.rear - s.front
}

//queue's first element
func (s *SliceQueue) GetFront() int {
	if s.IsEmpty() {
		panic("slice queue is empty")
	}
	return s.array[s.front]
}

//queue's last element
func (s *SliceQueue) GetRear() int {
	if s.IsEmpty() {
		panic("slice queue is empty")
	}
	return s.array[s.rear-1]
}

//delete first element of queue
func (s *SliceQueue) DeQueue() {
	if s.rear > s.front {
		s.rear--
		s.array = s.array[1:]
	} else {
		panic("slice queue is empty")
	}
}

//add element to queue
func (s *SliceQueue) EnQueue(val int) {
	s.rear++
	s.array = append(s.array, val)
}
