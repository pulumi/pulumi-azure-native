package tle

type Span struct {
	startIndex int
	length     int
}

// Get the start index of this span.
func (s Span) getStartIndex() int {
	return s.startIndex
}

// Get the getLength() of this span.
func (s Span) getLength() int {
	return s.length
}

// Get the index directly after this span.
//
// If this span started at 3 and had a getLength() of 4 ([3,7)), then the after
// end index would be 7.
func (s Span) afterEndIndex() int {
	return s.startIndex + s.getLength()
}

// Create a new span that is a union of this Span and the provided Span.
// If the provided Span is undefined, then this Span will be returned.
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

// Create a new span that is a union of the given spans.
// If both are undefined, undefined will be returned
func union(lhs *Span, rhs *Span) *Span {
	if lhs != nil {
		return lhs.union(rhs)
	}
	if rhs != nil {
		return rhs
	}
	return nil
}
