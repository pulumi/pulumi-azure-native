package tle

// Visitor visits Value types, allowing resolution to their component types
type Visitor interface {
	// VisitArrayAccessValue visits an ArrayAccessValue type
	VisitArrayAccessValue(*ArrayAccessValue) error
	VisitFunctionCallValue(*FunctionCallValue) error
	VisitNumberValue(*NumberValue) error
	VisitPropertyAccessValue(*PropertyAccessValue) error
	VisitStringValue(*StringValue) error
}

type wrappedVisitor struct {
	visitArrayAccessValue    func(*ArrayAccessValue) error
	visitFunctionCallValue   func(*FunctionCallValue) error
	visitNumberValue         func(*NumberValue) error
	visitPropertyAccessValue func(*PropertyAccessValue) error
	visitStringValue         func(*StringValue) error
}

func (w wrappedVisitor) VisitArrayAccessValue(v *ArrayAccessValue) error {
	if w.visitArrayAccessValue != nil {
		return w.visitArrayAccessValue(v)
	}
	panic(v)
}

func (w wrappedVisitor) VisitFunctionCallValue(v *FunctionCallValue) error {
	if w.visitFunctionCallValue != nil {
		return w.visitFunctionCallValue(v)
	}
	panic(v)
}

func (w wrappedVisitor) VisitNumberValue(v *NumberValue) error {
	if w.visitNumberValue != nil {
		return w.visitNumberValue(v)
	}
	panic(v)
}

func (w wrappedVisitor) VisitPropertyAccessValue(v *PropertyAccessValue) error {
	if w.visitPropertyAccessValue != nil {
		return w.visitPropertyAccessValue(v)
	}
	panic(v)
}

func (w wrappedVisitor) VisitStringValue(v *StringValue) error {
	if w.visitStringValue != nil {
		return w.visitStringValue(v)
	}
	panic(v)
}

// StringValueVisitor creates a new Visitor specifically designed to use with
// string values. Calling any other visit function will cause a panic
func StringValueVisitor(f func(*StringValue) error) Visitor {
	return &wrappedVisitor{visitStringValue: f}
}

// NumberValueVisitor creates a new Visitor specifically designed to use with
// number values. Calling any other visit function will cause a panic
func NumberValueVisitor(f func(*NumberValue) error) Visitor {
	return &wrappedVisitor{visitNumberValue: f}
}

func FunctionCallValueVisitor(f func(*FunctionCallValue) error) Visitor {
	return &wrappedVisitor{visitFunctionCallValue: f}
}

func PropertyAccessValueVisitor(p func(*PropertyAccessValue) error) Visitor {
	return &wrappedVisitor{visitPropertyAccessValue: p}
}

func ArrayAccessValueVisitor(p func(*ArrayAccessValue) error) Visitor {
	return &wrappedVisitor{visitArrayAccessValue: p}
}
