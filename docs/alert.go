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
	return elem.Div(
		elem.Paragraph(
			elem.Heading1(
				vecty.Text("Alerts"),
			),
			elem.Heading4(
				vecty.Text("Provide contextual feedback messages for typical user actions with the handful of available and flexible alert messages."),
			),
		),
		a.example(),
		a.linkColorExample(),
		a.additionalContent(),
	)
}

func (a *Alert) example() vecty.ComponentOrHTML {
	safeHTML := `
<pre>
	<code class="d-block" style="font-family: monospace; background: black; color: white;">
  &<span style="color: yellow">components</span>.<span style="color: teal">Alert</span> {
    Type: <span style="color: yellow">components</span>.<span style="color: purple">PrimaryAlert</span><span style="color: orange">,</span>
    Child: <span style="color: orange">func()</span> <span style="color: yellow">vecty</span>.<span style="color: teal">ComponentOrHTML</span> {
      <span style="color: orange">return</span> <span style="color: yellow">vecty</span>.<span style="color: brown">Text</span>(<span style="color: green">"A simple primary alert—check it out!")</span>
    }<span style="color: orange">,</span>
  }
  &<span style="color: yellow">components</span>.<span style="color: teal">Alert</span> {
    Type: <span style="color: yellow">components</span>.<span style="color: purple">SecondaryAlert</span><span style="color: orange">,</span>
    Child: <span style="color: orange">func()</span> <span style="color: yellow">vecty</span>.<span style="color: teal">ComponentOrHTML</span> {
      <span style="color: orange">return</span> <span style="color: yellow">vecty</span>.<span style="color: brown">Text</span>(<span style="color: green">"A simple secondary alert—check it out!")</span>
    }<span style="color: orange">,</span>
  }
  &<span style="color: yellow">components</span>.<span style="color: teal">Alert</span> {
    Type: <span style="color: yellow">components</span>.<span style="color: purple">SuccessAlert</span><span style="color: orange">,</span>
    Child: <span style="color: orange">func()</span> <span style="color: yellow">vecty</span>.<span style="color: teal">ComponentOrHTML</span> {
      <span style="color: orange">return</span> <span style="color: yellow">vecty</span>.<span style="color: brown">Text</span>(<span style="color: green">"A simple success alert—check it out!")</span>
    }<span style="color: orange">,</span>
  }
  &<span style="color: yellow">components</span>.<span style="color: teal">Alert</span> {
    Type: <span style="color: yellow">components</span>.<span style="color: purple">DangerAlert</span><span style="color: orange">,</span>
    Child: <span style="color: orange">func()</span> <span style="color: yellow">vecty</span>.<span style="color: teal">ComponentOrHTML</span> {
      <span style="color: orange">return</span> <span style="color: yellow">vecty</span>.<span style="color: brown">Text</span>(<span style="color: green">"A simple danger alert—check it out!")</span>
    }<span style="color: orange">,</span>
  }
  &<span style="color: yellow">components</span>.<span style="color: teal">Alert</span> {
    Type: <span style="color: yellow">components</span>.<span style="color: purple">WarningAlert</span><span style="color: orange">,</span>
    Child: <span style="color: orange">func()</span> <span style="color: yellow">vecty</span>.<span style="color: teal">ComponentOrHTML</span> {
      <span style="color: orange">return</span> <span style="color: yellow">vecty</span>.<span style="color: brown">Text</span>(<span style="color: green">"A simple warning alert—check it out!")</span>
    }<span style="color: orange">,</span>
  }
  &<span style="color: yellow">components</span>.<span style="color: teal">Alert</span> {
    Type: <span style="color: yellow">components</span>.<span style="color: purple">InfoAlert</span><span style="color: orange">,</span>
    Child: <span style="color: orange">func()</span> <span style="color: yellow">vecty</span>.<span style="color: teal">ComponentOrHTML</span> {
      <span style="color: orange">return</span> <span style="color: yellow">vecty</span>.<span style="color: brown">Text</span>(<span style="color: green">"A simple info alert—check it out!")</span>
    }<span style="color: orange">,</span>
  }
  &<span style="color: yellow">components</span>.<span style="color: teal">Alert</span> {
    Type: <span style="color: yellow">components</span>.<span style="color: purple">LightAlert</span><span style="color: orange">,</span>
    Child: <span style="color: orange">func()</span> <span style="color: yellow">vecty</span>.<span style="color: teal">ComponentOrHTML</span> {
      <span style="color: orange">return</span> <span style="color: yellow">vecty</span>.<span style="color: brown">Text</span>(<span style="color: green">"A simple light alert—check it out!")</span>
    }<span style="color: orange">,</span>
  }
  &<span style="color: yellow">components</span>.<span style="color: teal">Alert</span> {
    Type: <span style="color: yellow">components</span>.<span style="color: purple">DarkAlert</span><span style="color: orange">,</span>
    Child: <span style="color: orange">func()</span> <span style="color: yellow">vecty</span>.<span style="color: teal">ComponentOrHTML</span> {
      <span style="color: orange">return</span> <span style="color: yellow">vecty</span>.<span style="color: brown">Text</span>(<span style="color: green">"A simple dark alert—check it out!")</span>
    }<span style="color: orange">,</span>
  }
	</code>
</pre>
`
	return elem.Section(
		elem.Heading5(
			vecty.Text("Examples"),
		),
		elem.Paragraph(
			vecty.Text("Alerts are available for any length of text, as well as an optional close button."),
		),
		elem.Section(
			vecty.Markup(
				vecty.Class("border", "border-1", "p-3"),
			),
			&components.Alert{
				Type: components.PrimaryAlert,
				Child: func() vecty.ComponentOrHTML {
					return vecty.Text("A simple primary alert—check it out!")
				},
			},
			&components.Alert{
				Type: components.SecondaryAlert,
				Child: func() vecty.ComponentOrHTML {
					return vecty.Text("A simple secondary alert—check it out!")
				},
			},
			&components.Alert{
				Type: components.SuccessAlert,
				Child: func() vecty.ComponentOrHTML {
					return vecty.Text("A simple success alert—check it out!")
				},
			},
			&components.Alert{
				Type: components.DangerAlert,
				Child: func() vecty.ComponentOrHTML {
					return vecty.Text("A simple danger alert—check it out!")
				},
			},
			&components.Alert{
				Type: components.WarningAlert,
				Child: func() vecty.ComponentOrHTML {
					return vecty.Text("A simple warning alert—check it out!")
				},
			},
			&components.Alert{
				Type: components.InfoAlert,
				Child: func() vecty.ComponentOrHTML {
					return vecty.Text("A simple info alert—check it out!")
				},
			},
			&components.Alert{
				Type: components.LightAlert,
				Child: func() vecty.ComponentOrHTML {
					return vecty.Text("A simple light alert—check it out!")
				},
			},
			&components.Alert{
				MarkupData: internal.MarkupData{
					Classes: []string{"mb-0"},
				},
				Type: components.DarkAlert,
				Child: func() vecty.ComponentOrHTML {
					return vecty.Text("A simple dark alert—check it out!")
				},
			},
		),
		elem.Section(
			vecty.Markup(
				vecty.UnsafeHTML(safeHTML),
			),
		),
	)
}

func (a *Alert) linkColorExample() vecty.ComponentOrHTML {
	safeHTML := `
<pre>
	<code class="d-block" style="font-family: monospace; background: black; color: white;">
  &<span style="color: yellow">components</span>.<span style="color: teal">Alert</span> {
    Type: <span style="color: yellow">components</span>.<span style="color: purple">PrimaryAlert</span><span style="color: orange">,</span>
    Child: <span style="color: orange">func()</span> <span style="color: yellow">vecty</span>.<span style="color: teal">ComponentOrHTML</span> {
      <span style="color: orange">return</span> <span style="color: yellow">elem</span>.<span style="color: brown">Div</span>(
        <span style="color: yellow">vecty</span>.<span style="color: brown">Text</span>(<span style="color: green">"A simple primary alert with ")</span><span style="color: orange">,</span>
        <span style="color: yellow">elem</span>.<span style="color: brown">Anchor</span>(
          <span style="color: yellow">vecty</span>.<span style="color: brown">Markup</span>(
            <span style="color: yellow">vecty</span>.<span style="color: brown">Class</span>(<span style="color: green">"alert-link"</span>)<span style="color: orange">,</span>
            <span style="color: yellow">vecty</span>.<span style="color:brown">Href</span>(<span style="color: green">"#"</span>)<span style="color: orange">,</span>
          )<span style="color: orange">,</span>
          <span style="color: yellow">vecty</span>.<span style="color: brown">Text</span>(<span style="color: green">"an example link."</span>)<span style="color: orange">,</span>
        )<span style="color: orange">,</span>
      )
    }<span style="color: orange">,</span>
  }
  </code>
</pre>`

	return elem.Section(
		elem.Heading5(
			vecty.Text("Link Color"),
		),
		elem.Paragraph(
			vecty.Text("Use the "),
			elem.Code(vecty.Text(".alert-link ")),
			vecty.Text("utility class to quickly provide matching colored links within any alert."),
		),
		elem.Section(
			vecty.Markup(
				vecty.Class("border", "border-1", "p-3"),
			),
			&components.Alert{
				MarkupData: internal.MarkupData{
					Classes: []string{"mb-0"},
				},
				Type: components.PrimaryAlert,
				Child: func() vecty.ComponentOrHTML {
					return elem.Div(
						vecty.Text("A simple primary alert with "),
						elem.Anchor(
							vecty.Markup(
								vecty.Class("alert-link"),
								vecty.Href("#"),
							),
							vecty.Text("an example link."),
						),
					)
				},
			},
		),
		elem.Section(
			vecty.Markup(
				vecty.UnsafeHTML(safeHTML),
			),
		),
	)
}

func (a *Alert) additionalContent() vecty.ComponentOrHTML {
	safeHTML := `
<pre>
	<code class="d-block" style="font-family: monospace; background: black; color: white;">
  &<span style="color: yellow">components</span>.<span style="color: teal">Alert</span> {
    Type: <span style="color: yellow">components</span>.<span style="color: purple">SuccessAlert</span><span style="color: orange">,</span>
    Child: <span style="color: orange">func()</span> <span style="color: yellow">vecty</span>.<span style="color: teal">ComponentOrHTML</span> {
      <span style="color: orange">return</span> <span style="color: yellow">elem</span>.<span style="color: brown">Div</span>(
        <span style="color: yellow">elem</span>.<span style="color: brown">Heading4</span>(
          <span style="color: yellow">vecty</span>.<span style="color: brown">Markup</span>(
            <span style="color: yellow">vecty</span>.<span style="color: brown">Class</span>(<span style="color: green">"alert-heading"</span>)<span style="color: orange">,</span>
          )<span style="color: orange">,</span>
          <span style="color: yellow">vecty</span>.<span style="color: brown">Text</span>(<span style="color: green">"Well done!"</span>)<span style="color: orange">,</span>
        )<span style="color: orange">,</span>
        <span style="color: yellow">elem</span>.<span style="color: brown">Paragraph</span>(
          <span style="color: yellow">vecty</span>.<span style="color: brown">Text</span>(<span style="color: green">"Aww yeah, you successfully read this important alert message. This example text is going to run a bit longer so that you can see how spacing within an alert works with this kind of content."</span>)<span style="color: orange">,</span>
        )<span style="color: orange">,</span>
        <span style="color: yellow">elem</span>.<span style="color: brown">HorizontalRule</span>()<span style="color: orange">,</span>
        <span style="color: yellow">elem</span>.<span style="color: brown">Paragraph</span>(
          <span style="color: yellow">vecty</span>.<span style="color: brown">Markup</span>(
            <span style="color: yellow">vecty</span>.<span style="color: brown">Class</span>(<span style="color: green">"mb-0"</span>)<span style="color: orange">,</span>
          )<span style="color: orange">,</span>
          <span style="color: yellow">vecty</span>.<span style="color: brown">Text</span>(<span style="color: green">"Whenever you need to, be sure to use margin utilities to keep things nice and tidy."</span>)<span style="color: orange">,</span>
        )<span style="color: orange">,</span>
      )
    }<span style="color: orange">,</span>
  }
  </code>
</pre>`
	return elem.Section(
		elem.Heading5(
			vecty.Text("Additional Content"),
		),
		elem.Paragraph(
			vecty.Text("Alerts can also contain additional HTML elements like headings, paragraphs and dividers."),
		),
		elem.Section(
			vecty.Markup(
				vecty.Class("border", "border-1", "p-3"),
			),
			&components.Alert{
				MarkupData: internal.MarkupData{
					Classes: []string{"mb-0"},
				},
				Type: components.SuccessAlert,
				Child: func() vecty.ComponentOrHTML {
					return elem.Div(
						elem.Heading4(
							vecty.Markup(
								vecty.Class("alert-heading"),
							),
							vecty.Text("Well done!"),
						),
						elem.Paragraph(
							vecty.Text("Aww yeah, you successfully read this important alert message. This example text is going to run a bit longer so that you can see how spacing within an alert works with this kind of content."),
						),
						elem.HorizontalRule(),
						elem.Paragraph(
							vecty.Markup(
								vecty.Class("mb-0"),
							),
							vecty.Text("Whenever you need to, be sure to use margin utilities to keep things nice and tidy."),
						),
					)
				},
			},
		),
		elem.Section(
			vecty.Markup(
				vecty.UnsafeHTML(safeHTML),
			),
		),
	)
}
