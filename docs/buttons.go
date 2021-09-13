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
				Children: func() vecty.List {
					return vecty.List{
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color: components.PrimaryColor,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Primary")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color: components.SecondaryColor,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Secondary")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color: components.SuccessColor,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Success")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color: components.DangerColor,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Danger")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color: components.WarningColor,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Warning")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color: components.InfoColor,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Info")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color: components.LightColor,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Light")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color: components.DarkColor,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Dark")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							IsLink: true,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Link")
							},
						},
					}
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
					Child: func() vecty.ComponentOrHTML {
						return vecty.Text("Link")
					},
				},
			},
		),
		elem.Section(
			elem.Heading2(
				vecty.Text("Outline Buttons"),
			),
			&components.Flex{
				Children: func() vecty.List {
					return vecty.List{
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color:     components.PrimaryColor,
							IsOutline: true,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Primary")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color:     components.SecondaryColor,
							IsOutline: true,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Secondary")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color:     components.SuccessColor,
							IsOutline: true,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Success")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color:     components.DangerColor,
							IsOutline: true,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Danger")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color:     components.WarningColor,
							IsOutline: true,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Warning")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color:     components.InfoColor,
							IsOutline: true,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Info")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color:     components.LightColor,
							IsOutline: true,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Light")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color:     components.DarkColor,
							IsOutline: true,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Dark")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							IsLink:    true,
							IsOutline: true,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Link")
							},
						},
					}
				},
			},
		),
		elem.Section(
			elem.Heading2(
				vecty.Text("Sizes"),
			),
			&components.Flex{
				Children: func() vecty.List {
					return vecty.List{
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color: components.PrimaryColor,
							Size:  components.LargeSize,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Large Button")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color: components.SecondaryColor,
							Size:  components.LargeSize,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Large Button")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color: components.PrimaryColor,
							Size:  components.SmallSize,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Small Button")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color: components.SecondaryColor,
							Size:  components.SmallSize,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Small Button")
							},
						},
					}
				},
			},
		),
		elem.Section(
			elem.Heading2(
				vecty.Text("Disabled State"),
			),
			&components.Flex{
				Children: func() vecty.List {
					return vecty.List{
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color:      components.PrimaryColor,
							Size:       components.LargeSize,
							IsDisabled: true,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Primary Button")
							},
						},
						&components.Button{
							MarkupData: internal.MarkupData{
								Styles: []vecty.Style{
									vecty.Margin(vecty.Px(8)),
								},
							},
							Color:      components.SecondaryColor,
							Size:       components.LargeSize,
							IsDisabled: true,
							Child: func() vecty.ComponentOrHTML {
								return vecty.Text("Button")
							},
						},
						&components.AnchorButton{
							Button: components.Button{
								MarkupData: internal.MarkupData{
									Styles: []vecty.Style{
										vecty.Margin(vecty.Px(8)),
									},
								},
								Color:      components.PrimaryColor,
								IsDisabled: true,
								Child: func() vecty.ComponentOrHTML {
									return vecty.Text("Primary Link")
								},
							},
						},
						&components.AnchorButton{
							Button: components.Button{
								MarkupData: internal.MarkupData{
									Styles: []vecty.Style{
										vecty.Margin(vecty.Px(8)),
									},
								},
								Color:      components.SecondaryColor,
								IsDisabled: true,
								Child: func() vecty.ComponentOrHTML {
									return vecty.Text("Link")
								},
							},
						},
					}
				},
			},
		),
	)
}
