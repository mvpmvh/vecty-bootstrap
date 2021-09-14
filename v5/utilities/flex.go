package utilities

import "fmt"

// FlexDirectionType represents different css flex direction options
type FlexDirectionType string

const (
	FlexDirectionRow           FlexDirectionType = "row"
	FlexDirectionRowReverse    FlexDirectionType = "row-reverse"
	FlexDirectionColumn        FlexDirectionType = "column"
	FlexDirectionColumnReverse FlexDirectionType = "column-reverse"
)

// FlexDirectionParams configures a new flex direction utility class value
type FlexDirectionParams struct {
	Breakpoint
	FlexDirectionType
}

// FlexDirection returns a bootstrap flex direction utility class
func FlexDirection(params FlexDirectionParams) string {
	if params.Breakpoint != "" {
		return fmt.Sprintf("flex-%s-%s", params.Breakpoint, params.FlexDirectionType)
	}
	return fmt.Sprintf("flex-%s", params.FlexDirectionType)
}

// JustifyContentType represents different css flex justify-content options
type JustifyContentType string

const (
	JustifyContentStart   JustifyContentType = "start"
	JustifyContentEnd     JustifyContentType = "end"
	JustifyContentCenter  JustifyContentType = "center"
	JustifyContentBetween JustifyContentType = "between"
	JustifyContentAround  JustifyContentType = "around"
	JustifyContentEvenly  JustifyContentType = "evenly"
)

// JustifyContentParams configures a new flex justify-content utility class value
type JustifyContentParams struct {
	Breakpoint
	JustifyContentType
}

// JustifyContent returns a bootstrap flex direction utility class
func JustifyContent(params JustifyContentParams) string {
	if params.Breakpoint != "" {
		return fmt.Sprintf("justify-content-%s-%s", params.Breakpoint, params.JustifyContentType)
	}
	return fmt.Sprintf("justify-content-%s", params.JustifyContentType)
}

// AlignItemsType represents different css flex align-items options
type AlignItemsType string

const (
	AlignItemsStart    AlignItemsType = "start"
	AlignItemsEnd      AlignItemsType = "end"
	AlignItemsCenter   AlignItemsType = "center"
	AlignItemsBaseline AlignItemsType = "baseline"
	AlignItemsStretch  AlignItemsType = "stretch"
)

// AlignItemsParams configures a new flex align-items utility class value
type AlignItemsParams struct {
	Breakpoint
	AlignItemsType
}

// AlignItems returns a bootstrap flex align-items utility class
func AlignItems(params AlignItemsParams) string {
	if params.Breakpoint != "" {
		return fmt.Sprintf("align-items-%s-%s", params.Breakpoint, params.AlignItemsType)
	}
	return fmt.Sprintf("align-items-%s", params.AlignItemsType)
}

// AlignSelfType represents different css flex align-self options
type AlignSelfType string

const (
	AlignSelfStart    AlignSelfType = "start"
	AlignSelfEnd      AlignSelfType = "end"
	AlignSelfCenter   AlignSelfType = "center"
	AlignSelfBaseline AlignSelfType = "baseline"
	AlignSelfStretch  AlignSelfType = "stretch"
)

// AlignSelfParams configures a new flex align-self utility class value
type AlignSelfParams struct {
	Breakpoint
	AlignSelfType
}

// AlignSelf returns a bootstrap flex align-self utility class
func AlignSelf(params AlignSelfParams) string {
	if params.Breakpoint != "" {
		return fmt.Sprintf("align-self-%s-%s", params.Breakpoint, params.AlignSelfType)
	}
	return fmt.Sprintf("align-self-%s", params.AlignSelfType)
}

// FlexFill returns a bootstrap flex-fill utility class. If the variadic breakpoint param is set, only the first value
// is used.
func FlexFill(breakpoint ...Breakpoint) string {
	if breakpoint == nil {
		return "flex-fill"
	}
	return fmt.Sprintf("flex-%s-fill", breakpoint[0])
}

// FlexGrowVal represents different css flex-grow or flex-shrink amounts
type FlexGrowVal int

const (
	FlexGrowZero FlexGrowVal = iota
	FlexGrowOne
)

// FlexGrowOrShrinkParams configures a new flex-grow or flex-shrink utility class value
type FlexGrowOrShrinkParams struct {
	Breakpoint
	FlexGrowVal
}

// FlexGrow returns a bootstrap flex-grow utility class
func FlexGrow(params FlexGrowOrShrinkParams) string {
	if params.Breakpoint != "" {
		return fmt.Sprintf("flex-%s-%s-%d", params.Breakpoint, "grow", params.FlexGrowVal)
	}
	return fmt.Sprintf("flex-grow-%d", params.FlexGrowVal)
}

// FlexShrink returns a bootstrap flex-shrink utility class
func FlexShrink(params FlexGrowOrShrinkParams) string {
	if params.Breakpoint != "" {
		return fmt.Sprintf("flex-%s-%s-%d", params.Breakpoint, "shrink", params.FlexGrowVal)
	}
	return fmt.Sprintf("flex-shrink-%d", params.FlexGrowVal)
}

// FlexWrapType represents different css flex wrap options
type FlexWrapType string

const (
	FlexWrapTypeNoWrap      FlexWrapType = "nowrap"
	FlexWrapTypeWrap        FlexWrapType = "wrap"
	FlexWrapTypeWrapReverse FlexWrapType = "wrap-reverse"
)

// FlexWrapParams configures a new flex wrap utility class value
type FlexWrapParams struct {
	Breakpoint
	FlexWrapType
}

// FlexWrap returns a bootstrap flex-wrap utility class. If the variadic breakpoint param is set, only the first value
// is used.
func FlexWrap(params FlexWrapParams) string {
	if params.Breakpoint != "" {
		return fmt.Sprintf("flex-%s-%s", params.Breakpoint, params.FlexWrapType)
	}
	return fmt.Sprintf("flex-%s", params.FlexWrapType)
}

// FlexOrderVal represents different css flex order options
type FlexOrderVal int

const (
	FlexOrderFirst FlexOrderVal = -1
	FlexOrderZero  FlexOrderVal = iota
	FlexOrderOne
	FlexOrderTwo
	FlexOrderThree
	FlexOrderFour
	FlexOrderFive
	FlexOrderLast
)

// FlexOrderParams configures a new flex order utility class value
type FlexOrderParams struct {
	Breakpoint
	FlexOrderVal
}

// FlexOrder returns a bootstrap flex order utility class
func FlexOrder(params FlexOrderParams) string {
	if params.Breakpoint != "" {
		return fmt.Sprintf("order-%s-%d", params.Breakpoint, params.FlexOrderVal)
	}
	return fmt.Sprintf("order-%d", params.FlexOrderVal)
}

// AlignContentType represents different css flex align-content options
type AlignContentType string

const (
	AlignContentStart   AlignContentType = "start"
	AlignContentEnd     AlignContentType = "end"
	AlignContentCenter  AlignContentType = "center"
	AlignContentBetween AlignContentType = "between"
	AlignContentAround  AlignContentType = "around"
	AlignContentStretch AlignContentType = "stretch"
)

// AlignContentParams configures a new flex align-content utility class value
type AlignContentParams struct {
	Breakpoint
	AlignContentType
}

// AlignContent returns a bootstrap flex align-content utility class
func AlignContent(params AlignContentParams) string {
	if params.Breakpoint != "" {
		return fmt.Sprintf("align-content-%s-%s", params.Breakpoint, params.AlignContentType)
	}
	return fmt.Sprintf("align-content-%s", params.AlignContentType)
}
