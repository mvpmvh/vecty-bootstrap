package components

// Color is an "enum" for Bootstrap colors
type Color int

// Color themes
const (
	UnknownColor Color = iota
	PrimaryColor
	SecondaryColor
	SuccessColor
	DangerColor
	WarningColor
	InfoColor
	LightColor
	DarkColor
)
