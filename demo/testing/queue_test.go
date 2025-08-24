package queue

import "testing"

func TestAddQueue(t *testing.T) {
	q := New(3)
	for i := 0; i < 3; i++ {
		if !q.AddQueue(i) {
			t.Errorf("Failed to append item %v to queue", i)
		}
		if len(q.items) != i+1 {
			t.Errorf("Incorrect queue element count: %v, want %v", len(q.items), i+1)
		}
		if q.items[i] != i {
			t.Errorf("Incorrect item in queue: %v, want %v", q.items[i], i)
		}
	}

	if q.AddQueue(4) {
		t.Errorf("Should not be able to aadd to a full queue")
	}
}

func TestNext(t *testing.T) {
	q := New(3)
	for i := 0; i<3; i++ {
		q.AddQueue(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			t.Errorf("Should be able to get item from queue")
		}
		if item != i {
			t.Errorf("got item in wrong order: %v, want %v", item, i)
		}
	}

	// Queue is empty at this point
	item, ok := q.Next()
	if ok {
		t.Errorf("Should not be any more items in queue, got: %v", item)
	}
}