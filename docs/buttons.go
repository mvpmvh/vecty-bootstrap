package docs

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
	components "github.com/mvpmvh/vecty-bootstrap/v5"
)

type Button struct {
	vecty.Core
}

func (b *Button) Render() vecty.ComponentOrHTML {
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
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color: components.PrimaryColor,
						Child: vecty.Text("Primary"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color: components.SecondaryColor,
						Child: vecty.Text("Secondary"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color: components.SuccessColor,
						Child: vecty.Text("Success"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color: components.DangerColor,
						Child: vecty.Text("Danger"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color: components.WarningColor,
						Child: vecty.Text("Warning"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color: components.InfoColor,
						Child: vecty.Text("Info"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color: components.LightColor,
						Child: vecty.Text("Light"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color: components.DarkColor,
						Child: vecty.Text("Dark"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
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
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color:     components.PrimaryColor,
						IsOutline: true,
						Child:     vecty.Text("Primary"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color:     components.SecondaryColor,
						IsOutline: true,
						Child:     vecty.Text("Secondary"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color:     components.SuccessColor,
						IsOutline: true,
						Child:     vecty.Text("Success"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color:     components.DangerColor,
						IsOutline: true,
						Child:     vecty.Text("Danger"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color:     components.WarningColor,
						IsOutline: true,
						Child:     vecty.Text("Warning"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color:     components.InfoColor,
						IsOutline: true,
						Child:     vecty.Text("Info"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color:     components.LightColor,
						IsOutline: true,
						Child:     vecty.Text("Light"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color:     components.DarkColor,
						IsOutline: true,
						Child:     vecty.Text("Dark"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
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
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color: components.PrimaryColor,
						Size:  components.LargeSize,
						Child: vecty.Text("Large Button"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color: components.SecondaryColor,
						Size:  components.LargeSize,
						Child: vecty.Text("Large Button"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color: components.PrimaryColor,
						Size:  components.SmallSize,
						Child: vecty.Text("Small Button"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color: components.SecondaryColor,
						Size:  components.SmallSize,
						Child: vecty.Text("Small Button"),
					},
				},
			},
		),
		elem.Section(
			elem.Heading2(
				vecty.Text("Disabled State"),
			),
			&components.Flex{
				Children: vecty.List{
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color:      components.PrimaryColor,
						Size:       components.LargeSize,
						IsDisabled: true,
						Child:      vecty.Text("Primary Button"),
					},
					&components.Button{
						StyleData: internal.StyleData{
							Styles: []vecty.Style{
								vecty.Margin(vecty.Px(8)),
							},
						},
						Color:      components.SecondaryColor,
						Size:       components.LargeSize,
						IsDisabled: true,
						Child:      vecty.Text("Button"),
					},
					&components.AnchorButton{
						Button: components.Button{
							StyleData: internal.StyleData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color:      components.PrimaryColor,
							IsDisabled: true,
							Child:      vecty.Text("Primary Link"),
						},
					},
					&components.AnchorButton{
						Button: components.Button{
							StyleData: internal.StyleData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color:      components.SecondaryColor,
							IsDisabled: true,
							Child:      vecty.Text("Link"),
						},
					},
				},
			},
		),
	)
}
