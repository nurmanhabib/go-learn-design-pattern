package memento

type Document struct {
	ID           string
	Title        string
	Participants []*Participant
}

func (d *Document) AddParticipant(participant ...*Participant) {
	d.Participants = append(d.Participants, participant...)
}

func (d *Document) CreateSnapshot() *Snapshot {
	return NewDocumentSnapshot(d)
}

func (d *Document) Restore(snapshot *Snapshot) {
	if snapshot != nil {
		d.ID = snapshot.document.ID
		d.Title = snapshot.document.Title
		d.Participants = snapshot.document.Participants
	} else {
		d.ID = ""
		d.Title = ""
		d.Participants = []*Participant{}
	}
}
