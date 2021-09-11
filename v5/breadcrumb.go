package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
)

// Breadcrumbs creates a collection of breadcrumb items (see: https://getbootstrap.com/docs/5.1/components/breadcrumb/)
type Breadcrumbs struct {
	vecty.Core
	internal.MarkupData
	Items   []*BreadcrumbItem
	Divider *string
}

func (b *Breadcrumbs) Render() vecty.ComponentOrHTML {
	children := []vecty.MarkupOrChild{
		vecty.Markup(vecty.Class("breadcrumb")),
	}

	for _, item := range b.Items {
		children = append(children, item)
	}

	var divider string
	if b.Divider != nil {
		divider = *b.Divider
	}

	return elem.Navigation(
		vecty.Markup(
			vecty.Property{"aria-label": "breadcrumb"},
			vecty.MarkupIf(b.Divider != nil, vecty.Style{"--bs-breadcrumb-divider": divider}),
		),
		elem.OrderedList(children...),
	)
}

// BreadcrumbItem is an individual breadcrumb in a containing list of breadcrumbs
type BreadcrumbItem struct {
	vecty.Core
	internal.MarkupData
	IsActive bool
	Child    vecty.ComponentOrHTML
}

func (b *BreadcrumbItem) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		vecty.Markup(
			vecty.Class("breadcrumb-item"),
			vecty.ClassMap{"active": b.IsActive},
			vecty.MarkupIf(b.IsActive, vecty.Property{"aria-current": "page"}),
		),
		b.Child,
	)
}
