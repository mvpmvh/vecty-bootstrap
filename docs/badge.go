package docs

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
	components "github.com/mvpmvh/vecty-bootstrap/v5"
	"github.com/mvpmvh/vecty-bootstrap/v5/utilities"
)

type Badge struct {
	vecty.Core
}

func (b *Badge) Render() vecty.ComponentOrHTML {
	return elem.Section(
		&components.Badge{
			MarkupData: internal.MarkupData{
				Classes: []string{utilities.BackgroundPrimary},
			},
			Child: vecty.Text("badge"),
		},
		&components.Badge{
			Child: vecty.Text("pill"),
			MarkupData: internal.MarkupData{
				Classes: []string{utilities.BackgroundSecondary, utilities.BorderRadiusRoundedPill},
			},
		},
	)
}
