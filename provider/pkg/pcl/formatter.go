// Copyright 2021, Pulumi Corporation.  All rights reserved.

package pcl

import (
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/model"
	"github.com/pulumi/pulumi/pkg/v2/codegen/hcl2/syntax"
	"github.com/zclconf/go-cty/cty"
)

// a formatter formats PCL into a canonical style.
type formatter struct {
	indentLevel int
}

// space returns a TriviaList composed of a single space.
func (f *formatter) space() syntax.TriviaList {
	return syntax.TriviaList{syntax.NewWhitespace(' ')}
}

// newline returns a TriviaList composed of a single newline.
func (f *formatter) newline() syntax.TriviaList {
	return syntax.TriviaList{syntax.NewWhitespace('\n')}
}

// indent returns a TriviaList composed of whitespace that corresponds to the current indent level.
func (f *formatter) indent() syntax.TriviaList {
	return syntax.TriviaList{syntax.NewWhitespace([]byte(strings.Repeat("\t", f.indentLevel))...)}
}

// indented increments the indent level, runs the provided function, then decrements the indent level.
func (f *formatter) indented(fn func()) {
	f.indentLevel++
	defer func() { f.indentLevel-- }()
	fn()
}

// formatObjectCons formats an object construction expression.
// - An object with no elements is formatted as `{}`
// - An object with elements is formatted as
//
//       {
//           key0 = value0,
//           ...
//           keyN = valueN
//       }
//
func (f *formatter) formatObjectCons(x *model.ObjectConsExpression) {
	x.Tokens = syntax.NewObjectConsTokens(len(x.Items))

	if len(x.Items) == 0 {
		x.Tokens.CloseBrace.LeadingTrivia = nil
	} else {
		x.Tokens.OpenBrace.TrailingTrivia = f.newline()
		f.indented(func() {
			for i, item := range x.Items {
				f.formatExpression(item.Key)
				item.Key.SetLeadingTrivia(f.indent())

				x.Tokens.Items[i].Equals.LeadingTrivia = f.space()
				x.Tokens.Items[i].Equals.TrailingTrivia = f.space()

				f.formatExpression(item.Value)
				if x.Tokens.Items[i].Comma == nil {
					item.Value.SetTrailingTrivia(f.newline())
				} else {
					x.Tokens.Items[i].Comma.TrailingTrivia = f.newline()
				}
			}
		})
		x.Tokens.CloseBrace.LeadingTrivia = f.indent()
	}
}

// formatTupleCons formats a tuple construction expression.
// - A tuple with no elements is formatted as `[]`
// - A tuple with one element is formatted as `[value]`
// - A tuple with more than one element is formatted as
//
//       [
//           value0,
//           ...
//           valueN
//       ]
//
func (f *formatter) formatTupleCons(x *model.TupleConsExpression) {
	x.Tokens = syntax.NewTupleConsTokens(len(x.Expressions))

	switch len(x.Expressions) {
	case 0:
		x.Tokens.CloseBracket.LeadingTrivia = nil
	case 1:
		f.formatExpression(x.Expressions[0])

		x.Expressions[0].SetLeadingTrivia(nil)
		x.Expressions[0].SetTrailingTrivia(nil)
	default:
		x.Tokens.OpenBracket.TrailingTrivia = f.newline()
		f.indented(func() {
			for i, item := range x.Expressions {
				f.formatExpression(item)

				item.SetLeadingTrivia(f.indent())

				if i < len(x.Tokens.Commas) {
					x.Tokens.Commas[i].TrailingTrivia = f.newline()
				} else {
					item.SetTrailingTrivia(f.newline())
				}
			}
		})
		x.Tokens.CloseBracket.LeadingTrivia = f.indent()
	}
}

// formatExpression formats an expression. Extraneous leading and trailing whitespace is eliminated, and single spaces
// are inserted around specific tokens.
func (f *formatter) formatExpression(x model.Expression) model.Expression {
	x.SetLeadingTrivia(nil)
	x.SetTrailingTrivia(nil)

	switch x := x.(type) {
	case *model.AnonymousFunctionExpression:
		f.formatExpression(x.Body)
	case *model.BinaryOpExpression:
		x.Tokens = syntax.NewBinaryOpTokens(x.Operation)

		f.formatExpression(x.LeftOperand)
		f.formatExpression(x.RightOperand).SetLeadingTrivia(f.space())
	case *model.ConditionalExpression:
		x.Tokens = syntax.NewConditionalTokens()

		f.formatExpression(x.Condition)
		f.formatExpression(x.TrueResult).SetLeadingTrivia(f.space())
		f.formatExpression(x.FalseResult).SetLeadingTrivia(f.space())
	case *model.ForExpression:
		keyName := ""
		if x.KeyVariable != nil {
			keyName = x.KeyVariable.Name
		}
		x.Tokens = syntax.NewForTokens(keyName, x.ValueVariable.Name, x.Key != nil, x.Group, x.Condition != nil)

		f.formatExpression(x.Collection).SetLeadingTrivia(f.space())
		if x.Key != nil {
			f.formatExpression(x.Key).SetLeadingTrivia(f.space())
		}
		f.formatExpression(x.Value).SetLeadingTrivia(f.space())
		if x.Condition != nil {
			f.formatExpression(x.Condition).SetLeadingTrivia(f.space())
		}
	case *model.FunctionCallExpression:
		x.Tokens = syntax.NewFunctionCallTokens(x.Name, len(x.Args))

		for i, x := range x.Args {
			f.formatExpression(x)
			if i > 0 {
				x.SetLeadingTrivia(f.space())
			}
		}
	case *model.IndexExpression:
		x.Tokens = syntax.NewIndexTokens()

		f.formatExpression(x.Collection)
		f.formatExpression(x.Key)
	case *model.RelativeTraversalExpression:
		x.Tokens = syntax.NewRelativeTraversalTokens(x.Traversal)

		f.formatExpression(x.Source)
	case *model.ScopeTraversalExpression:
		x.Tokens = syntax.NewScopeTraversalTokens(x.Traversal)
	case *model.SplatExpression:
		x.Tokens = syntax.NewSplatTokens(false)

		f.formatExpression(x.Source)
		f.formatExpression(x.Each)
	case *model.TemplateExpression:
		x.Tokens = syntax.NewTemplateTokens()

		for i, part := range x.Parts {
			f.formatExpression(part)

			if lit, ok := part.(*model.LiteralValueExpression); ok && lit.Value.Type() == cty.String {
				if i > 0 {
					part.SetLeadingTrivia(syntax.TriviaList{syntax.NewTemplateDelimiter(hclsyntax.TokenTemplateSeqEnd)})
				}
				if i < len(x.Parts)-1 {
					part.SetTrailingTrivia(syntax.TriviaList{syntax.NewTemplateDelimiter(hclsyntax.TokenTemplateInterp)})
				}
			}
		}
	case *model.TemplateJoinExpression:
		f.formatExpression(x.Tuple)
	case *model.UnaryOpExpression:
		x.Tokens = syntax.NewUnaryOpTokens(x.Operation)

		f.formatExpression(x.Operand)
	case *model.ObjectConsExpression:
		f.formatObjectCons(x)
	case *model.TupleConsExpression:
		f.formatTupleCons(x)
	}
	return x
}

// formatBlock formats a PCL block as
//
//     type labels... {
//         item0
//         ...
//         itemN
//     }
//
// Unless the block is the first in its containing body, a newline is inserted before the block.
func (f *formatter) formatBlock(block *model.Block, first bool) {
	block.Tokens = syntax.NewBlockTokens(block.Type, block.Labels...)

	leadingTrivia := f.indent()
	if !first {
		leadingTrivia = append(f.newline(), leadingTrivia...)
	}
	block.Tokens.Type.LeadingTrivia = leadingTrivia

	f.indented(func() {
		f.formatBody(block.Body)
	})

	block.Tokens.CloseBrace.LeadingTrivia = f.indent()
	block.Tokens.CloseBrace.TrailingTrivia = f.newline()
}

// formatAttribute formats a PCL attribute as
//
//     name = value
//
// If the attribute follows a block, a newline is inserted before the attribute.
func (f *formatter) formatAttribute(attr *model.Attribute, afterBlock bool) {
	attr.Tokens = syntax.NewAttributeTokens(attr.Name)

	leadingTrivia := f.indent()
	if afterBlock {
		leadingTrivia = append(f.newline(), leadingTrivia...)
	}
	attr.Tokens.Name.LeadingTrivia = leadingTrivia

	f.formatExpression(attr.Value)

	attr.Value.SetLeadingTrivia(f.space())
	attr.Value.SetTrailingTrivia(f.newline())
}

// formatBody formats a PCL body.
func (f *formatter) formatBody(body *model.Body) {
	first, afterBlock := true, false
	for _, item := range body.Items {
		switch item := item.(type) {
		case *model.Block:
			f.formatBlock(item, first)
			afterBlock = true
		case *model.Attribute:
			f.formatAttribute(item, afterBlock)
			afterBlock = false
		}
		first = false
	}
}

// FormatBody formats a PCL body.
func FormatBody(body *model.Body) {
	var f formatter
	f.formatBody(body)
}
