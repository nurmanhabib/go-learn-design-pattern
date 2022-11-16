package memento_test

import (
	"memento/memento"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDocumentSnapshot(t *testing.T) {
	t.Run("If undo success", func(t *testing.T) {
		history := memento.NewHistory()
		document := &memento.Document{
			ID:           "1",
			Title:        "Learn Memento Concept",
			Participants: []*memento.Participant{},
		}

		// Create first snapshot
		history.Push(document.CreateSnapshot())

		// Add two participants
		document.AddParticipant(
			&memento.Participant{
				ID:    "participant-1",
				Name:  "Jay",
				Email: "jay@gmail.com",
			},
			&memento.Participant{
				ID:    "participant-2",
				Name:  "Jeje",
				Email: "jeje@gmail.com",
			},
		)

		// Create snapshot
		history.Push(document.CreateSnapshot())

		// Add one participant again
		document.AddParticipant(
			&memento.Participant{
				ID:    "participant-3",
				Name:  "Ghiffary",
				Email: "ghiffary@gmail.com",
			},
		)

		// Create snapshot
		history.Push(document.CreateSnapshot())
		assert.Equal(t, 3, len(document.Participants))

		// Undo one time
		document.Restore(history.Undo())
		assert.Equal(t, 2, len(document.Participants))

		// Redo one time (latest revision)
		document.Restore(history.Redo())
		assert.Equal(t, 3, len(document.Participants))
	})

	t.Run("If create snapshot after undoing", func(t *testing.T) {
		history := memento.NewHistory()
		document := &memento.Document{
			ID:           "1",
			Title:        "Hello World",
			Participants: []*memento.Participant{},
		}

		// Add two participants
		document.AddParticipant(
			&memento.Participant{
				ID:    "participant-1",
				Name:  "Jay",
				Email: "jay@gmail.com",
			},
			&memento.Participant{
				ID:    "participant-2",
				Name:  "Jeje",
				Email: "jeje@gmail.com",
			},
		)

		// Create snapshot
		history.Push(document.CreateSnapshot())

		// Add one participant again
		document.AddParticipant(
			&memento.Participant{
				ID:    "participant-3",
				Name:  "Ghiffary",
				Email: "ghiffary@gmail.com",
			},
		)

		// Create snapshot
		history.Push(document.CreateSnapshot())

		// Add one participant again
		document.AddParticipant(
			&memento.Participant{
				ID:    "participant-4",
				Name:  "Ramdhani",
				Email: "ramdhani@gmail.com",
			},
		)

		// Create snapshot
		history.Push(document.CreateSnapshot())

		// Add one participant again
		document.AddParticipant(
			&memento.Participant{
				ID:    "participant-5",
				Name:  "Fadhila",
				Email: "fadhila@gmail.com",
			},
		)

		// Create snapshot
		history.Push(document.CreateSnapshot())
		assert.Equal(t, 5, len(document.Participants))

		// Undo multiple time
		document.Restore(history.Undo())
		document.Restore(history.Undo())
		assert.Equal(t, 3, len(document.Participants))
		assert.Equal(t, 4, len(history.All()))

		// Add new participant after undoing
		document.AddParticipant(&memento.Participant{
			ID:    "participant-99",
			Name:  "Devina",
			Email: "devina@gmail.com",
		})

		// Create new snapshot after undoing
		history.Push(document.CreateSnapshot())

		assert.Equal(t, 4, len(document.Participants))
		assert.Equal(t, 3, len(history.All()))
	})
}
