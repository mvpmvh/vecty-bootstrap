package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	components "github.com/mvpmvh/vecty-bootstrap/v5"
)

func main() {
	vecty.SetTitle("Button Demo")
	vecty.AddStylesheet(vecty.AddStyleSheetParams{
		URL: "https://cdn.jsdelivr.net/npm/bootstrap-icons@1.5.0/font/bootstrap-icons.css",
	})
	vecty.AddStylesheet(vecty.AddStyleSheetParams{
		URL:         "https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/css/bootstrap.min.css",
		Integrity:   "sha384-KyZXEAg3QhqLMpG8r+8fhAXLRk2vvoC2f3B09zVXn8CA5QIVfZOJ3BCsw2P0p/We",
		CrossOrigin: "anonymous",
	})
	vecty.AddScript(vecty.AddScriptParams{
		URL:         "https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/js/bootstrap.bundle.min.js",
		Integrity:   "sha384-U1DAWAznBHeqEIlVSCgzq+c9gqGAJn5c/t99JyeKa9xxaYpSvHU5awsuZVVFIhvj",
		CrossOrigin: "anonymous",
		AddToHead:   false,
	})
	vecty.RenderBody(&demo{})
}

type demo struct {
	vecty.Core
}

func (d *demo) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Heading1(
			vecty.Text("Buttons"),
		),
		elem.Section(
			elem.Heading2(
				vecty.Text("Styles"),
			),
			&components.Flex{
				Children: vecty.List{
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color: components.PrimaryColor,
						Child: vecty.Text("Primary"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color: components.SecondaryColor,
						Child: vecty.Text("Secondary"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color: components.SuccessColor,
						Child: vecty.Text("Success"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color: components.DangerColor,
						Child: vecty.Text("Danger"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color: components.WarningColor,
						Child: vecty.Text("Warning"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color: components.InfoColor,
						Child: vecty.Text("Info"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color: components.LightColor,
						Child: vecty.Text("Light"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color: components.DarkColor,
						Child: vecty.Text("Dark"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						IsLink: true,
						Child:  vecty.Text("Link"),
					},
				},
			},
		),
		elem.Section(
			elem.Heading2(
				vecty.Text("Anchor Button"),
			),
			&components.AnchorButton{
				Button: components.Button{
					Color: components.PrimaryColor,
					Child: vecty.Text("Link"),
				},
			},
		),
		elem.Section(
			elem.Heading2(
				vecty.Text("Outline Buttons"),
			),
			&components.Flex{
				Children: vecty.List{
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color:     components.PrimaryColor,
						IsOutline: true,
						Child:     vecty.Text("Primary"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color:     components.SecondaryColor,
						IsOutline: true,
						Child:     vecty.Text("Secondary"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color:     components.SuccessColor,
						IsOutline: true,
						Child:     vecty.Text("Success"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color:     components.DangerColor,
						IsOutline: true,
						Child:     vecty.Text("Danger"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color:     components.WarningColor,
						IsOutline: true,
						Child:     vecty.Text("Warning"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color:     components.InfoColor,
						IsOutline: true,
						Child:     vecty.Text("Info"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color:     components.LightColor,
						IsOutline: true,
						Child:     vecty.Text("Light"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color:     components.DarkColor,
						IsOutline: true,
						Child:     vecty.Text("Dark"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						IsLink:    true,
						IsOutline: true,
						Child:     vecty.Text("Link"),
					},
				},
			},
		),
		elem.Section(
			elem.Heading2(
				vecty.Text("Sizes"),
			),
			&components.Flex{
				Children: vecty.List{
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color: components.PrimaryColor,
						Size:  components.LargeSize,
						Child: vecty.Text("Large Button"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color: components.SecondaryColor,
						Size:  components.LargeSize,
						Child: vecty.Text("Large Button"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color: components.PrimaryColor,
						Size:  components.SmallSize,
						Child: vecty.Text("Small Button"),
					},
					&components.Button{
						Styles: []vecty.Style{
							vecty.Margin(vecty.Px(8)),
						},
						Color: components.SecondaryColor,
						Size:  components.SmallSize,
						Child: vecty.Text("Small Button"),
					},
				},
			},
		),
	)
}
