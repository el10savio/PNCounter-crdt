package pncounter

import (
	"github.com/el10savio/gcounter-crdt/gcounter"
)

// PNCounter is the PNCounter CRDT data type
type PNCounter struct {
	// Add is a gcounter to store the values added
	Add gcounter.GCounter `json:"add"`
	// Delete is a gcounter to store the values removed
	Delete gcounter.GCounter `json:"delete"`
}

const (
	// singleNodeName is the node name assigned to
	// the node when it is the only node
	// present in the cluster
	singleNodeName = "node"
)

// Initialize returns a new empty PNCounter
func Initialize(node string) PNCounter {
	// Set the node name as singleNodeName
	// in case of single node
	if node == "" {
		node = singleNodeName
	}

	return PNCounter{
		Add:    gcounter.Initialize(node),
		Delete: gcounter.Initialize(node),
	}
}

// Increment increases the Add count for
// a given node in the PNCounter by 1
func (pncounter PNCounter) Increment(node string) PNCounter {
	// Set the node name as singleNodeName
	// in case of single node
	if node == "" {
		node = singleNodeName
	}

	// Increment the PNCounter Count
	pncounter.Add.Increment(node)

	// Return the updated PNCounter
	// Count value
	return pncounter
}

// Decrement increases the Delete count for
// a given node in the PNCounter by 1
func (pncounter PNCounter) Decrement(node string) PNCounter {
	// Set the node name as singleNodeName
	// in case of single node
	if node == "" {
		node = singleNodeName
	}

	// Increment the PNCounter Count
	pncounter.Delete.Increment(node)

	// Return the updated PNCounter
	// Count value
	return pncounter
}

// GetTotal returns the sum of the counts
// of all the nodes in GCounter.Count
func (pncounter PNCounter) GetTotal() int {
	// Return the total count
	return pncounter.Add.GetTotal() - pncounter.Delete.GetTotal()
}

// Merge combines multiple PNCounters into a single PNCounter
// for the same node in multiple PNCounters the merged
// node's count is the max count obtained
func Merge(PNCounters ...PNCounter) PNCounter {
	// Initialize the merged PNCounter
	pncounterMerged := PNCounters[0]

	// pncounterMerged = Max(pncounterMerged, pncounterToMergeWith)
	for _, pncounter := range PNCounters {
		for node, value := range pncounter.Add.Count {
			pncounterMerged.Add.Count[node] = Max(pncounterMerged.Add.Count[node], value)
		}
		for node, value := range pncounter.Delete.Count {
			pncounterMerged.Delete.Count[node] = Max(pncounterMerged.Delete.Count[node], value)
		}
	}

	// Return the merged PNCounter
	return pncounterMerged
}

// SetCount is a utility function  used in tests
// that assigns the PNCounter's Count
// value to a specified value
func (pncounter PNCounter) SetCount(node string, addValue int, deleteValue int) PNCounter {
	// Set the node name as singleNodeName
	// in case of single node
	if node == "" {
		node = singleNodeName
	}

	// Set the count of the node
	pncounter.Add.Count[node] = addValue
	pncounter.Delete.Count[node] = deleteValue

	// Return the updated PNCounter Count
	return pncounter
}

// Clear is a utility function used for tests that clears the PNCounter
func (pncounter PNCounter) Clear(node string) PNCounter {
	pncounter.Add = gcounter.Initialize(node)
	pncounter.Delete = gcounter.Initialize(node)

	// Return the Cleared PNCounter
	return pncounter
}

// Max computes the maximum of
// 2 int values passed to it
func Max(value1, value2 int) int {
	if value1 > value2 {
		return value1
	}
	return value2
}
