package arm2pulumi

import (
	"fmt"
	"sort"
)

func newDependencyTracking(e TemplateElement) *dependencyTracking {
	return &dependencyTracking{
		TemplateElement: e,
		dependencies:    map[*dependencyTracking]bool{},
		incoming:        map[*dependencyTracking]bool{},
	}
}

// dependencyTracking allows TemplateElement to be represented in a DAG.
type dependencyTracking struct {
	TemplateElement
	dependencies map[*dependencyTracking]bool // things this item references (i.e. must precede in topo sort order)
	incoming     map[*dependencyTracking]bool // things that refer to this item
}

func (d *dependencyTracking) RefersTo(ref *dependencyTracking) {
	d.dependencies[ref] = true
	ref.incoming[d] = true
}

func (d *dependencyTracking) Dependencies() []*dependencyTracking {
	var result []*dependencyTracking
	for k := range d.dependencies {
		result = append(result, k)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name() < result[j].Name()
	})

	return result
}

// topoSort performs a topological sort of elements accessible from e.
// visited is a map to track already seen elements and detect cycles.
// Absence from visited means an item has not been seen before. Presence
// but the map value set to false means that all its dependencies have not
// been processed. Finally if set to true, the the item has had all
// its neighbors visited. Returns the sorted order or an error if a
// cycle is detected.
func topoSort(
	e *dependencyTracking,
	visited map[*dependencyTracking]bool,
	order []*dependencyTracking,
) ([]*dependencyTracking, error) {

	committed, seen := visited[e]
	if committed {
		return order, nil
	}
	if seen {
		return nil, fmt.Errorf("detected cycle at %s", e.Name())
	}
	visited[e] = false // Mark as visited (but not committed)
	var sorted []*dependencyTracking
	for _, el := range e.Dependencies() {
		s, err := topoSort(el, visited, order)
		if err != nil {
			return nil, err
		}
		sorted = append(sorted, s...)
	}
	sorted = append(sorted, e)
	visited[e] = true // Mark as committed
	return sorted, nil
}
