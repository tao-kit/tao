package timer

// Len is used to implement the interface of sort.Interface.
func (h *priorityQueueHeap) Len() int {
	return len(h.array)
}

// Less is used to implement the interface of sort.Interface.
func (h *priorityQueueHeap) Less(i, j int) bool {
	return h.array[i].priority < h.array[j].priority
}

// Swap is used to implement the interface of sort.Interface.
func (h *priorityQueueHeap) Swap(i, j int) {
	if len(h.array) == 0 {
		return
	}
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

// Push pushes an item to the heap.
func (h *priorityQueueHeap) Push(x any) {
	h.array = append(h.array, x.(priorityQueueItem))
}

// Pop retrieves, removes and returns the most high priority item from the heap.
func (h *priorityQueueHeap) Pop() any {
	length := len(h.array)
	if length == 0 {
		return nil
	}
	item := h.array[length-1]
	h.array = h.array[0 : length-1]
	return item
}
