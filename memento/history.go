package memento

type History struct {
	revisions    []*Snapshot
	currentIndex int
}

func NewHistory() *History {
	return &History{
		revisions:    []*Snapshot{},
		currentIndex: -1,
	}
}

func (h *History) Push(snapshot *Snapshot) {
	h.cutOff()
	h.revisions = append(h.revisions, snapshot)
	h.currentIndex++
}

func (h *History) All() []*Snapshot {
	return h.revisions
}

func (h *History) cutOff() {
	if h.currentIndex > 0 {
		h.revisions = h.revisions[:h.currentIndex+1]
	}
}

func (h *History) Undo() *Snapshot {
	if h.currentIndex > h.minIndex() {
		h.currentIndex--
	}

	return h.revisions[h.currentIndex]
}

func (h *History) Redo() *Snapshot {
	if h.currentIndex < h.maxIndex() {
		h.currentIndex++
	}

	return h.revisions[h.currentIndex]
}

func (h *History) maxIndex() int {
	return len(h.revisions) - 1
}

func (h *History) minIndex() int {
	return 0
}
