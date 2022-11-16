package memento

type Snapshot struct {
	document Document
}

func NewDocumentSnapshot(document *Document) *Snapshot {
	return &Snapshot{document: *document}
}

func (ds *Snapshot) GetDocument() *Document {
	return &ds.document
}
