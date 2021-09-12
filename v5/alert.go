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
	CloseButton          *CloseButton
	ShouldAnimateDismiss bool
	Child                vecty.ComponentOrHTML
}

func (a *Alert) Render() vecty.ComponentOrHTML {
	if a.CloseButton != nil {
		data := vecty.Data{"bsDismiss": "alert"}
		for k, v := range a.CloseButton.Data {
			data[k] = v
		}
		a.CloseButton.Data = data
	}

	return elem.Div(
		vecty.Markup(a.markup()),
		a.Child,
		vecty.If(a.CloseButton != nil, a.CloseButton),
	)
}

func (a *Alert) markup() vecty.MarkupList {
	return vecty.Markup(
		a.MarkupData.Markup(),
		vecty.Class("alert", fmt.Sprintf("alert-%s", a.Type)),
		vecty.ClassMap{
			"alert-dismissible": a.CloseButton != nil,
			"fade":              a.ShouldAnimateDismiss,
			"show":              a.ShouldAnimateDismiss,
		},
		vecty.Attribute("role", "alert"),
	)
}
