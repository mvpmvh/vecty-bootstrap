package docs

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	components "github.com/mvpmvh/vecty-bootstrap/v5"
)

type Breadcrumb struct {
	vecty.Core
}

func (b *Breadcrumb) Render() vecty.ComponentOrHTML {
	return elem.Section(
		&components.Breadcrumbs{
			Items: []*components.BreadcrumbItem{
				{
					Child: elem.Anchor(
						vecty.Markup(vecty.Href("#")),
						vecty.Text("Home"),
					),
				},
				{
					Child: elem.Anchor(
						vecty.Markup(vecty.Href("#")),
						vecty.Text("Library"),
					),
				},
				{
					Child:    vecty.Text("Data"),
					IsActive: true,
				},
			},
		},
	)
}
