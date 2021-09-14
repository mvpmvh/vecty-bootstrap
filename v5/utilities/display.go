package utilities

import "fmt"

// DisplayType represents different css display options
type DisplayType string

const (
	DisplayNone        DisplayType = "none"
	DisplayInline      DisplayType = "inline"
	DisplayInlineBlock DisplayType = "inline-block"
	DisplayBlock       DisplayType = "block"
	DisplayGrid        DisplayType = "grid"
	DisplayTable       DisplayType = "table"
	DisplayTableCell   DisplayType = "table-cell"
	DisplayTableRow    DisplayType = "display-row"
	DisplayFlex        DisplayType = "flex"
	DisplayInlineFlex  DisplayType = "inline-flex"
)

// DisplayParams configures a new display utility class value
type DisplayParams struct {
	Breakpoint
	DisplayType
}

// Display returns a bootstrap display utility class
func Display(params DisplayParams) string {
	if params.Breakpoint != "" {
		return fmt.Sprintf("d-%s-%s", params.Breakpoint, params.DisplayType)
	}
	return fmt.Sprintf("d-%s", params.DisplayType)
}

// Print returns a bootstrap display utility for print
func Print(d DisplayType) string {
	return fmt.Sprintf("d-print-%s", d)
}
