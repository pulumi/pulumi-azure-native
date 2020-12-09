package arm2pulumi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type node struct {
	TemplateElement
	name string
}

func newNode(name string) *dependencyTracking {
	return newDependencyTracking(&node{name: name})
}
func (n *node) Name() string {
	return n.name
}

func TestTopologicalSort(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		// A-> B -> C
		// \       /
		//  D ->  E
		//         \
		// F -> G -> H
		// I
		a := newNode("A")
		b := newNode("B")
		c := newNode("C")
		d := newNode("D")
		e := newNode("E")
		f := newNode("F")
		g := newNode("G")
		h := newNode("H")
		i := newNode("I")

		b.RefersTo(a)
		c.RefersTo(b)
		d.RefersTo(a)
		e.RefersTo(d)
		c.RefersTo(e)
		g.RefersTo(f)
		h.RefersTo(g)
		h.RefersTo(e)

		els := TemplateElements{
			elements: map[string]*dependencyTracking{
				"a": a,
				"b": b,
				"c": c,
				"d": d,
				"e": e,
				"f": f,
				"g": g,
				"h": h,
				"i": i,
			},
		}

		ordered, err := els.topologicalOrder()
		require.NoError(t, err)
		var actual []string
		for _, o := range ordered {
			t.Logf(o.Name())
			actual = append(actual, o.Name())
		}

		assert.Equal(t, []string{"A", "B", "D", "E", "C", "F", "G", "H", "I"}, actual)
	})

	t.Run("Cycle", func(t *testing.T) {
		// A -> B -> C
		//     /       \
		//    G <- E <- D
		// H

		a := newNode("A")
		b := newNode("B")
		c := newNode("C")
		d := newNode("D")
		e := newNode("E")
		f := newNode("F")
		g := newNode("G")
		h := newNode("H")

		a.RefersTo(b)
		b.RefersTo(c)
		c.RefersTo(d)
		d.RefersTo(e)
		e.RefersTo(g)
		g.RefersTo(b)

		els := TemplateElements{
			elements: map[string]*dependencyTracking{
				"a": a,
				"b": b,
				"c": c,
				"d": d,
				"e": e,
				"f": f,
				"g": g,
				"h": h,
			},
		}

		_, err := els.topologicalOrder()
		require.EqualError(t, err, "detected cycle at B")
	})
}
