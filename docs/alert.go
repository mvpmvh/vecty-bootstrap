package docs

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
	components "github.com/mvpmvh/vecty-bootstrap/v5"
)

type Alert struct {
	vecty.Core
}

func (a *Alert) Render() vecty.ComponentOrHTML {
	return elem.Section(
		&components.Alert{
			MarkupData:           internal.MarkupData{},
			Type:                 components.PrimaryAlert,
			ShouldAnimateDismiss: true,
			Child:                vecty.Text("A simple primary alert—check it out!"),
		},
	)
}
