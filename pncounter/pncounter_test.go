package pncounter

import (
	"testing"

	"github.com/el10savio/gcounter-crdt/gcounter"
	"github.com/stretchr/testify/assert"
)

const (
	// testNode is the node name used
	// for single node PNCounter tests
	testNode = "test-node"
)

var (
	pncounter PNCounter
)

func init() {
	// Initialize the test PNCounter
	pncounter = Initialize(testNode)
}

// TestGetCount checks the basic functionality of PNCounter GetTotal()
// GetTotal() should return all count map entries added to the PNCounter
func TestGetCount(t *testing.T) {
	pncounter = pncounter.Increment(testNode)

	expectedCount := 1
	actualCount := pncounter.GetTotal()

	assert.Equal(t, expectedCount, actualCount)

	pncounter = pncounter.Clear(testNode)
}

// TestGetCount_UpdatedValue checks the functionality of PNCounter GetTotal()
// when values are incremented in the PNCounter it should return
// all the count map entries added to the PNCounter
func TestGetCount_UpdatedValue(t *testing.T) {
	pncounter = pncounter.Increment(testNode)
	pncounter = pncounter.Increment(testNode)
	pncounter = pncounter.Increment(testNode)

	expectedCount := 3
	actualCount := pncounter.GetTotal()

	assert.Equal(t, expectedCount, actualCount)

	pncounter = pncounter.Clear(testNode)
}

// TestGetCount_NoValue checks the functionality of PNCounter GetTotal()
// when no values are added to PNCounter, it should return an
// initialized empty count map
func TestGetCount_NoValue(t *testing.T) {
	expectedCount := 0
	actualCount := pncounter.GetTotal()

	assert.Equal(t, expectedCount, actualCount)

	pncounter = pncounter.Clear(testNode)
}

// TestIncrement checks the basic functionality of PNCounter Increment
// it should return the PNCounter node count incremented by 1
func TestIncrement(t *testing.T) {
	pncounter = pncounter.Increment(testNode)

	expectedCount := 1
	actualCount := pncounter.GetTotal()

	assert.Equal(t, expectedCount, actualCount)

	pncounter = pncounter.Clear(testNode)
}

// TestIncrement checks the basic functionality of PNCounter Decrement
// it should return the PNCounter node count decremented by 1
func TestDecrement(t *testing.T) {
	pncounter = pncounter.Increment(testNode)
	pncounter = pncounter.Increment(testNode)
	pncounter = pncounter.Decrement(testNode)

	expectedCount := 1
	actualCount := pncounter.GetTotal()

	assert.Equal(t, expectedCount, actualCount)

	pncounter = pncounter.Clear(testNode)
}

// TestIncrement checks the basic functionality of PNCounter Decrement
// as the total count becomes negative
func TestDecrement_TotalBecomesNegative(t *testing.T) {
	pncounter = pncounter.Decrement(testNode)
	pncounter = pncounter.Decrement(testNode)
	pncounter = pncounter.Decrement(testNode)
	pncounter = pncounter.Decrement(testNode)

	expectedCount := -4
	actualCount := pncounter.GetTotal()

	assert.Equal(t, expectedCount, actualCount)

	pncounter = pncounter.Clear(testNode)
}

// TestClear checks the basic functionality of PNCounter Clear
// this utility function it clears all the values in a PNCounter
func TestClear(t *testing.T) {
	pncounter = pncounter.Increment(testNode)
	pncounter = pncounter.Increment(testNode)
	pncounter = pncounter.Decrement(testNode)
	pncounter = pncounter.Clear(testNode)

	expectedCount := 0
	actualCount := pncounter.GetTotal()

	assert.Equal(t, expectedCount, actualCount)

	pncounter = pncounter.Clear(testNode)
}

// TestClear_EmptyStore checks the functionality of PNCounter Clear
// utility function when no values are in it
func TestClear_EmptyStore(t *testing.T) {
	pncounter = pncounter.Clear(testNode)

	expectedCount := 0
	actualCount := pncounter.GetTotal()

	assert.Equal(t, expectedCount, actualCount)

	pncounter = pncounter.Clear(testNode)
}

// TestGetTotal checks the basic functionality of PNCounter GetTotal()
// this function should return the total of all nodes in count
func TestGetTotal(t *testing.T) {
	pncounter = pncounter.SetCount("testNode1", 1, 1)
	pncounter = pncounter.SetCount("testNode2", 3, 1)
	pncounter = pncounter.SetCount("testNode3", 5, 3)

	expectedCount := 4
	actualCount := pncounter.GetTotal()

	assert.Equal(t, expectedCount, actualCount)

	pncounter = pncounter.Clear(testNode)
}

// TestGetTotal checks GetTotal function when
// no node counts are present, it should then return 0
func TestGetTotal_EmptyCount(t *testing.T) {
	expectedCount := 0
	actualCount := pncounter.GetTotal()

	assert.Equal(t, expectedCount, actualCount)

	pncounter = pncounter.Clear(testNode)
}

// TestMerge checks the basic functionality of the Merge() function on multiple PNCounters
// it returns all the PNCounters merged together with unique elements as a single PNCounter
func TestMerge(t *testing.T) {
	pncounter1 := PNCounter{
		Add:    gcounter.GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}},
		Delete: gcounter.GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}},
	}
	pncounter2 := PNCounter{
		Add:    gcounter.GCounter{map[string]int{"node1": 4, "node2": 6, "node3": 8}},
		Delete: gcounter.GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}},
	}
	pncounter3 := PNCounter{
		Add:    gcounter.GCounter{map[string]int{"node1": 2, "node2": 4, "node3": 9}},
		Delete: gcounter.GCounter{map[string]int{"node1": 2, "node2": 4, "node3": 9}},
	}

	pncounterExpected := PNCounter{
		Add:    gcounter.GCounter{map[string]int{"node1": 4, "node2": 6, "node3": 9}},
		Delete: gcounter.GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 9}},
	}

	pncounterActual := Merge(pncounter1, pncounter2, pncounter3)

	countExpected := 2
	countActual := pncounterActual.GetTotal()

	assert.Equal(t, pncounterExpected, pncounterActual)
	assert.Equal(t, countExpected, countActual)

	pncounter = pncounter.Clear(testNode)
}

// TestMerge_Empty checks the functionality of the Merge() function on multiple PNCounters
// when one PNCounters are empty, it returns an empty PNCounter followed by an error
func TestMerge_Empty(t *testing.T) {
	pncounter1 := PNCounter{}
	pncounter2 := PNCounter{}
	pncounter3 := PNCounter{}

	pncounterExpected := PNCounter{}
	pncounterActual := Merge(pncounter1, pncounter2, pncounter3)

	countExpected := 0
	countActual := pncounterActual.GetTotal()

	assert.Equal(t, pncounterExpected, pncounterActual)
	assert.Equal(t, countExpected, countActual)

	pncounter = pncounter.Clear(testNode)
}

// TestMerge_Duplicate checks the functionality of the Merge() function on multiple PNCounters
// when duplicate values are passed with the PNCounter it returns all the PNCounters
// merged together with maximum counts as one single PNCounter
func TestMerge_Duplicate(t *testing.T) {
	pncounter1 := PNCounter{
		Add:    gcounter.GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}},
		Delete: gcounter.GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}},
	}
	pncounter2 := PNCounter{
		Add:    gcounter.GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}},
		Delete: gcounter.GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}},
	}
	pncounter3 := PNCounter{
		Add:    gcounter.GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}},
		Delete: gcounter.GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}},
	}

	pncounterExpected := PNCounter{
		Add:    gcounter.GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}},
		Delete: gcounter.GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}},
	}

	pncounterActual := Merge(pncounter1, pncounter2, pncounter3)

	countExpected := 0
	countActual := pncounterActual.GetTotal()

	assert.Equal(t, pncounterExpected, pncounterActual)
	assert.Equal(t, countExpected, countActual)

	pncounter = pncounter.Clear(testNode)
}
