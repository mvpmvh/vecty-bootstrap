package components

// Breakpoint represents a responsive breakpoint (see: https://getbootstrap.com/docs/5.1/layout/breakpoints/)
type Breakpoint string

// The supported responsive breakpoints
const (
	BreakpointExtraSmall      Breakpoint = "xs"
	BreakpointSmall           Breakpoint = "sm"
	BreakpointMedium          Breakpoint = "md"
	BreakpointLarge           Breakpoint = "lg"
	BreakpointXtraLarge       Breakpoint = "xl"
	BreakpointExtraExtraLarge Breakpoint = "xxl"
)
