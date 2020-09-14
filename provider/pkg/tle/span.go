package tle

import "github.com/pulumi/pulumi/sdk/go/common/util/contract"

func fromStartAndAfterEnd(startIndex int, afterEndIndex int) Span {
	return Span{startIndex: startIndex, length: afterEndIndex - startIndex}
}

type Span struct {
	startIndex int
	length     int
}

type ContainsBehavior string

const (
	strict   ContainsBehavior = "strict"
	extended ContainsBehavior = "extended"
	enclosed ContainsBehavior = "enclosed"
)

/**
 * Get the start index of this span.
 */

func (s Span) getStartIndex() int {
	return s.startIndex
}

/**
 * Get the getLength() of this span.
 */
func (s Span) getLength() int {
	return s.length
}

/**
 * Get the last index of this span.
 */

func (s Span) endIndex() int {
	l := 0
	if s.getLength() > 0 {
		return s.getLength() - 1
	}
	return s.startIndex + l
}

/**
 * Get the index directly after this span.
 *
 * If this span started at 3 and had a getLength() of 4 ([3,7)), then the after
 * end index would be 7.
 */

func (s Span) afterEndIndex() int {
	return s.startIndex + s.getLength()
}

/**
 * Determine if the provided index is contained by this span.
 *
 * If this span started at 3 and had a getLength() of 4, i.e. [3, 7), then all
 * indexes between 3 and 6 (inclusive) would be contained. 2 and 7 would
 * not be contained.
 *
 * If includeAfterEndIndex=true, then 3-7 inclusive would be contained.
 */
func (s Span) contains(index int, containsBehavior ContainsBehavior) bool {
	switch containsBehavior {
	case strict:
		return s.startIndex <= index && index <= s.endIndex()

	case extended:
		return s.startIndex <= index && index <= s.afterEndIndex()

	case enclosed:
		return s.startIndex+1 <= index && index <= s.endIndex()

	default:
		contract.Failf("Unexpected containsBehavior %v", containsBehavior)
	}
	return false
}

/**
 * Create a new span that is a union of this Span and the provided Span.
 * If the provided Span is undefined, then this Span will be returned.
 */

func (s Span) union(rhs *Span) *Span {
	var result Span
	if rhs != nil {
		minStart := s.startIndex
		if minStart > rhs.startIndex {
			minStart = rhs.startIndex
		}
		maxAfterEndIndex := s.afterEndIndex()
		if maxAfterEndIndex < rhs.afterEndIndex() {
			maxAfterEndIndex = rhs.afterEndIndex()
		}
		result = Span{startIndex: minStart, length: maxAfterEndIndex - minStart}
	} else {
		result = s
	}
	return &result
}

/**
 * Create a new span that is a union of the given spans.
 * If both are undefined, undefined will be returned
 */

func union(lhs *Span, rhs *Span) *Span {
	if lhs != nil {
		return lhs.union(rhs)
	}
	if rhs != nil {
		return rhs
	}
	return nil
}

/**
 * Create a new span that is the intersection of this and a given span.
 * If the provided span is undefined, or they do no intersect, undefined will be returned
 */

func (s Span) intersect(rhs *Span) *Span {
	if rhs != nil {
		// tslint:disable-next-line:no-this-assignment
		lhss := s
		lhs := &lhss
		if rhs.startIndex < s.startIndex {
			lhs, rhs = rhs, lhs
		}

		start := rhs.startIndex
		var afterEnd int
		if lhs.afterEndIndex() < rhs.afterEndIndex() {
			afterEnd = lhs.afterEndIndex()
		} else {
			afterEnd = rhs.afterEndIndex()
		}
		if afterEnd >= start {
			return &Span{startIndex: start, length: afterEnd - start}
		} else {
			return nil
		}
	}

	return nil
}

/**
 * Create a new span that is the intersection of the given spans.
 * If either is undefined, or they do no intersect, undefined will be returned
 */

func intersect(lhs *Span, rhs *Span) *Span {
	if lhs != nil {
		return lhs.intersect(rhs)
	}
	return nil
}

func (s Span) translate(movement int) Span {
	if movement == 0 {
		return s
	}
	return Span{startIndex: s.startIndex + movement, length: s.getLength()}
}

func (s Span) extendLeft(extendLeft int) Span {
	return Span{startIndex: s.startIndex - extendLeft, length: s.getLength() + extendLeft}
}

func (s Span) extendRight(extendRight int) Span {
	return Span{startIndex: s.startIndex, length: s.getLength() + extendRight}
}

func (s Span) getText(text string, offsetIndex ...int) string {
	offset := 0
	if len(offsetIndex) > 0 {
		offset = offsetIndex[0]
	}
	start := s.startIndex + offset
	return text[start : start+s.getLength()]
}
