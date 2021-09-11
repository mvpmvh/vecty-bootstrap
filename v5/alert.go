package components

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
)

// AlertType controls the appearance of the alert (see: https://getbootstrap.com/docs/5.1/components/buttons/#sizes)
type AlertType string

const (
	PrimaryAlert   AlertType = "primary"
	SecondaryAlert AlertType = "secondary"
	SuccessAlert   AlertType = "success"
	DangerAlert    AlertType = "danger"
	WarningAlert   AlertType = "warning"
	InfoAlert      AlertType = "info"
	LightAlert     AlertType = "light"
	DarkAlert      AlertType = "dark"
)

// Alert creates an alert (see: https://getbootstrap.com/docs/5.1/components/alerts/)
type Alert struct {
	vecty.Core
	internal.MarkupData
	Type                 AlertType
	DismissButton        *Button
	ShouldAnimateDismiss bool
	Child                vecty.ComponentOrHTML
}

func (a *Alert) Render() vecty.ComponentOrHTML {
	if a.DismissButton != nil {
		a.DismissButton.Classes = append(a.DismissButton.Classes, "btn-close")
		a.DismissButton.Data = append(a.DismissButton.Data, vecty.Data{"bsDismiss": "alert"})
		a.DismissButton.Properties = append(a.DismissButton.Properties, vecty.Property{"aria-label": "close"})
	}

	return elem.Div(
		vecty.Markup(a.markup()),
		a.Child,
		vecty.If(a.DismissButton != nil, a.DismissButton),
	)
}

func (a *Alert) markup() vecty.MarkupList {
	return vecty.Markup(
		a.MarkupData.Markup(),
		vecty.Class("alert", fmt.Sprintf("alert-%s", a.Type)),
		vecty.ClassMap{
			"alert-dismissible": a.DismissButton != nil,
			"fade":              a.ShouldAnimateDismiss,
			"show":              a.ShouldAnimateDismiss,
		},
		vecty.Attribute("role", "alert"),
	)
}
