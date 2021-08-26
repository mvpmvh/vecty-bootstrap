package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

type Input struct {
	vecty.Core
	Name        string
	IsFloating  bool
	Placeholder string
	InputType   prop.InputType
	Id          string
	Value       string
	IsInvalid   bool
	IsReadOnly  bool
	// Custom classes to apply to the top-most container div
	ContainerClasses []string
}

// Key implements the vecty.Keyer interface.
// func (i *Input) Key() interface{} {
// 	return i.Id // TODO: should I generate a value for an unexported "key" field instead?
// }

func (i *Input) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class(i.ContainerClasses...),
			vecty.ClassMap{
				"form-floating": i.IsFloating,
			},
		),
		elem.Input(
			vecty.Markup(
				vecty.Class("form-control"),
				vecty.ClassMap{
					"is-invalid": i.IsInvalid,
				},
				prop.Type(i.InputType),
				vecty.MarkupIf(i.Id != "", prop.ID(i.Id)),
				vecty.MarkupIf(i.Name != "", prop.Name(i.Name)),
				vecty.MarkupIf(i.IsReadOnly, vecty.Property("readOnly", true)),
				vecty.MarkupIf(i.Placeholder != "", prop.Placeholder(i.Placeholder)),
				prop.Value(i.Value),
			),
		),
		vecty.If(i.Placeholder != "", elem.Label(
			vecty.Markup(
				vecty.Class("form-label"),
				vecty.MarkupIf(i.Id != "", prop.For(i.Id)),
			),
			vecty.Text(i.Placeholder),
		)),
	)
}
