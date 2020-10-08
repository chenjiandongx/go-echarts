package opts

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/go-echarts/go-echarts/types"
)

// Initialization contains options for the canvas.
type Initialization struct {
	// HTML title
	PageTitle string `default:"Awesome go-echarts"`

	// Canvas width
	Width string `default:"900px"`

	// Canvas Height
	Height string `default:"500px"`

	// Canvas Background Color
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// Chart unique ID
	ChartID string

	// Assets host
	AssetsHost string `default:"https://go-echarts.github.io/go-echarts-assets/assets/"`

	// Chart Theme
	Theme string `default:"white"`
}

// Validate validates the initialization configurations.
func (opt *Initialization) Validate() {
	setDefaultValue(opt)
	if opt.ChartID == "" {
		opt.ChartID = generateUniqueID()
	}
}

// set default values for the struct field.
// origin from: https://github.com/mcuadros/go-defaults
func setDefaultValue(ptr interface{}) {
	elem := reflect.ValueOf(ptr).Elem()
	t := elem.Type()

	for i := 0; i < t.NumField(); i++ {
		// handle `default` tag only
		if defaultVal := t.Field(i).Tag.Get("default"); defaultVal != "" {
			setField(elem.Field(i), defaultVal)
		}
	}
}

// setField handles String/Bool types only.
func setField(field reflect.Value, defaultVal string) {
	switch field.Kind() {
	case reflect.String:
		if field.String() == "" {
			field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
		}
	case reflect.Bool:
		if val, err := strconv.ParseBool(defaultVal); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}
	}
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	chartIDSize   = 12
)

// generate the unique ID for each chart.
func generateUniqueID() string {
	seed := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, chartIDSize)
	for i, cache, remain := chartIDSize-1, seed.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = seed.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// Title is the option set for a title component.
type Title struct {
	// The main title text, supporting for \n for newlines.
	Title string `json:"text,omitempty"`

	// TextStyle of the main title.
	TitleStyle *TextStyle `json:"textStyle,omitempty"`

	// The hyper link of main title text.
	Link string `json:"link,omitempty"`

	// Subtitle text, supporting for \n for newlines.
	Subtitle string `json:"subtext,omitempty"`

	// TextStyle of the sub title.
	SubtitleStyle *TextStyle `json:"subtextStyle,omitempty"`

	// The hyper link of sub title text.
	SubLink string `json:"sublink,omitempty"`

	// Open the hyper link of main title in specified tab.
	// options:
	// 'self' opening it in current tab
	// 'blank' opening it in a new tab
	Target string `json:"target,omitempty"`

	// Distance between title component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between title component and the bottom side of the container.
	// bottom value can be instant pixel value like 20;
	// it can also be a percentage value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`

	// Distance between title component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between title component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	//Adaptive by default.
	Right string `json:"right,omitempty"`
}

// Legend is the option set for a legend component.
type Legend struct {
	// Whether to show the Legend, default true.
	Show bool `json:"show"`
	// Distance between legend component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right', then the component
	// will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between legend component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom', then the component
	// will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between legend component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right string `json:"right,omitempty"`

	// Distance between legend component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`

	// Data array of legend. An array item is usually a name representing string.
	// set Data as []string{} if you wants to hide the legend.
	Data interface{} `json:"data,omitempty"`

	// The layout orientation of legend.
	// Options: 'horizontal', 'vertical'
	Orient string `json:"orient,omitempty"`

	// Legend color when not selected.
	InactiveColor string `json:"inactiveColor,omitempty"`

	// State table of selected legend.
	// example:
	// var selected = map[string]bool{}
	// selected["series1"] = true
	// selected["series2"] = false
	Selected map[string]bool `json:"selected,omitempty"`

	// Selected mode of legend, which controls whether series can be toggled displaying by clicking legends.
	// It is enabled by default, and you may set it to be false to disabled it.
	// Besides, it can be set to 'single' or 'multiple', for single selection and multiple selection.
	SelectedMode string `json:"selectedMode,omitempty"`

	// Legend space around content. The unit is px.
	// Default values for each position are 5.
	// And they can be set to different values with left, right, top, and bottom.
	// Examples:
	// 1. Set padding to be 5
	//    padding: 5
	// 2. Set the top and bottom paddings to be 5, and left and right paddings to be 10
	//    padding: [5, 10]
	// 3. Set each of the four paddings seperately
	//    padding: [
	//      5,  // up
	//      10, // right
	//      5,  // down
	//      10, // left
	//    ]
	Padding interface{} `json:"padding,omitempty"`

	// Image width of legend symbol.
	ItemWidth int `json:"itemWidth,omitempty"`

	// Image height of legend symbol.
	ItemHeight int `json:"itemHeight,omitempty"`

	// Legend X position, right/left/center
	X string `json:"x,omitempty"`

	// Legend Y position, right/left/center
	Y string `json:"y,omitempty"`

	// Width of legend component. Adaptive by default.
	Width string `json:"width,omitempty"`

	// Height of legend component. Adaptive by default.
	Height string `json:"height,omitempty"`

	// Legend marker and text aligning.
	// By default, it automatically calculates from component location and orientation.
	// When left value of this component is 'right' and orient is 'vertical', it would be aligned to 'right'.
	// Options: auto/left/right
	Align string `json:"align,omitempty"`

	// Legend text style.
	*TextStyle `json:"textStyle,omitempty"`
}

// Tooltip is the option set for a tooltip component.
type Tooltip struct {
	// Whether to show the tooltip component, including tooltip floating layer and axisPointer.
	Show bool `json:"show"`

	// Type of triggering.
	// Options:
	// * 'item': Triggered by data item, which is mainly used for charts that
	//    don't have a category axis like scatter charts or pie charts.
	// * 'axis': Triggered by axes, which is mainly used for charts that have category axes,
	//    like bar charts or line charts.
	// * 'none': Trigger nothing.
	Trigger string `json:"trigger,omitempty"`

	// Conditions to trigger tooltip. Options:
	// * 'mousemove': Trigger when mouse moves.
	// * 'click': Trigger when mouse clicks.
	// * 'mousemove|click': Trigger when mouse clicks and moves.
	// * 'none': Do not triggered by 'mousemove' and 'click'. Tooltip can be triggered and hidden
	//    manually by calling action.tooltip.showTip and action.tooltip.hideTip.
	//    It can also be triggered by axisPointer.handle in this case.
	TriggerOn string `json:"triggerOn,omitempty"`

	// The content formatter of tooltip's floating layer which supports string template and callback function.
	//
	// 1. String template
	// The template variables are {a}, {b}, {c}, {d} and {e}, which stands for series name,
	// data name and data value and ect. When trigger is set to be 'axis', there may be data from multiple series.
	// In this time, series index can be refereed as {a0}, {a1}, or {a2}.
	// {a}, {b}, {c}, {d} have different meanings for different series types:
	//
	// * Line (area) charts, bar (column) charts, K charts: {a} for series name,
	//   {b} for category name, {c} for data value, {d} for none;
	// * Scatter (bubble) charts: {a} for series name, {b} for data name, {c} for data value, {d} for none;
	// * Map: {a} for series name, {b} for area name, {c} for merging data, {d} for none;
	// * Pie charts, gauge charts, funnel charts: {a} for series name, {b} for data item name,
	//   {c} for data value, {d} for percentage.
	//
	// 2. Callback function
	// The format of callback function:
	// (params: Object|Array, ticket: string, callback: (ticket: string, html: string)) => string
	// The first parameter params is the data that the formatter needs. Its format is shown as follows:
	// {
	//    componentType: 'series',
	//    // Series type
	//    seriesType: string,
	//    // Series index in option.series
	//    seriesIndex: number,
	//    // Series name
	//    seriesName: string,
	//    // Data name, or category name
	//    name: string,
	//    // Data index in input data array
	//    dataIndex: number,
	//    // Original data as input
	//    data: Object,
	//    // Value of data. In most series it is the same as data.
	//    // But in some series it is some part of the data (e.g., in map, radar)
	//    value: number|Array|Object,
	//    // encoding info of coordinate system
	//    // Key: coord, like ('x' 'y' 'radius' 'angle')
	//    // value: Must be an array, not null/undefined. Contain dimension indices, like:
	//    // {
	//    //     x: [2] // values on dimension index 2 are mapped to x axis.
	//    //     y: [0] // values on dimension index 0 are mapped to y axis.
	//    // }
	//    encode: Object,
	//    // dimension names list
	//    dimensionNames: Array<String>,
	//    // data dimension index, for example 0 or 1 or 2 ...
	//    // Only work in `radar` series.
	//    dimensionIndex: number,
	//    // Color of data
	//    color: string,
	//
	//    // the percentage of pie chart
	//    percent: number,
	// }
	Formatter string `json:"formatter,omitempty"`
}

// Toolbox is the option set for a toolbox component.
type Toolbox struct {
	// Whether to show toolbox component.
	Show bool `json:"show"`

	// The layout orientation of toolbox's icon.
	// Options: 'horizontal','vertical'
	Orient string `json:"orient,omitempty"`

	// Distance between toolbox component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right', then the component
	// will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between toolbox component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom', then the component
	// will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between toolbox component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right string `json:"right,omitempty"`

	// Distance between toolbox component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`

	// The configuration item for each tool.
	// Besides the tools we provide, user-defined toolbox is also supported.
	Feature *ToolBoxFeature `json:"feature,omitempty"`
}

// ToolBoxFeature is a feature component under toolbox.
type ToolBoxFeature struct {
	// Save as image tool
	SaveAsImage *ToolBoxFeatureSaveAsImage `json:"saveAsImage,omitempty"`

	// Data area zooming, which only supports rectangular coordinate by now.
	DataZoom *ToolBoxFeatureDataZoom `json:"dataZoom,omitempty"`

	// Data view tool, which could display data in current chart and updates chart after being edited.
	DataView *ToolBoxFeatureDataView `json:"dataView,omitempty"`

	// Restore configuration item.
	Restore *ToolBoxFeatureRestore `json:"restore,omitempty"`
}

// ToolBoxFeatureSaveAsImage is the option for saving chart as image.
type ToolBoxFeatureSaveAsImage struct {
	// Whether to show the tool.
	Show bool `json:"show"`

	// toolbox.feature.saveAsImage. type = 'png'
	// File suffix of the image saved.
	// If the renderer is set to be 'canvas' when chart initialized (default), t
	// hen 'png' (default) and 'jpeg' are supported.
	// If the renderer is set to be 'svg' when when chart initialized, then only 'svg' is supported
	// for type ('svg' type is supported since v4.8.0).
	Type string `json:"png,omitempty"`

	// Name to save the image, whose default value is title.text.
	Name string `json:"name,omitempty"`

	// title for the tool.
	Title string `json:"title,omitempty"`
}

// ToolBoxFeatureDataZoom
type ToolBoxFeatureDataZoom struct {
	// Whether to show the tool.
	Show bool `json:"show"`

	// Restored and zoomed title text.
	// m["zoom"] = 'area zooming'
	// m["back"] = 'restore area zooming'
	Title map[string]string `json:"title"`
}

// ToolBoxFeatureDataView
type ToolBoxFeatureDataView struct {
	// Whether to show the tool.
	Show bool `json:"show"`

	// title for the tool.
	Title string `json:"title,omitempty"`
}

// ToolBoxFeatureRestore
type ToolBoxFeatureRestore struct {
	// Whether to show the tool.
	Show bool `json:"show"`

	// title for the tool.
	Title string `json:"title,omitempty"`
}

// XAxis is the option set for X axis.
type XAxis struct {
	// Name of axis.
	Name string `json:"name,omitempty"`

	// Type of axis.
	// Option:
	// * 'value': Numerical axis, suitable for continuous data.
	// * 'category': Category axis, suitable for discrete category data.
	//   Category data can be auto retrieved from series.data or dataset.source,
	//   or can be specified via xAxis.data.
	// * 'time' Time axis, suitable for continuous time series data. As compared to value axis,
	//   it has a better formatting for time and a different tick calculation method. For example,
	//   it decides to use month, week, day or hour for tick based on the range of span.
	// * 'log' Log axis, suitable for log data.
	Type string `json:"type,omitempty"`

	// Set this to false to prevent the axis from showing.
	Show bool `json:"show,omitempty"`

	// Category data, available in type: 'category' axis.
	Data interface{} `json:"data,omitempty"`

	// Number of segments that the axis is split into. Note that this number serves only as a
	// recommendation, and the true segments may be adjusted based on readability.
	// This is unavailable for category axis.
	SplitNumber int `json:"splitNumber,omitempty"`

	// It is available only in numerical axis, i.e., type: 'value'.
	// It specifies whether not to contain zero position of axis compulsively.
	// When it is set to be true, the axis may not contain zero position,
	// which is useful in the scatter chart for both value axes.
	// This configuration item is unavailable when the min and max are set.
	Scale bool `json:"scale,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Max interface{} `json:"max,omitempty"`

	// The index of grid which the x axis belongs to. Defaults to be in the first grid.
	// default 0
	GridIndex int `json:"gridIndex,omitempty"`

	// Split area of X axis in grid area.
	*SplitArea `json:"splitArea,omitempty"`

	// Split line of X axis in grid area.
	*SplitLine `json:"splitLine,,omitempty"`
}

// YAxis is the option set for Y axis.
type YAxis struct {
	// Name of axis.
	Name string `json:"name,omitempty"`

	// Type of axis.
	// Option:
	// * 'value': Numerical axis, suitable for continuous data.
	// * 'category': Category axis, suitable for discrete category data.
	//   Category data can be auto retrieved from series.data or dataset.source,
	//   or can be specified via xAxis.data.
	// * 'time' Time axis, suitable for continuous time series data. As compared to value axis,
	//   it has a better formatting for time and a different tick calculation method. For example,
	//   it decides to use month, week, day or hour for tick based on the range of span.
	// * 'log' Log axis, suitable for log data.
	Type string `json:"type,omitempty"`

	// Set this to false to prevent the axis from showing.
	Show bool `json:"show,omitempty"`

	// Formatter of axis label, which supports string template and callback function.
	// 1. Use string template; template variable is the default label of axis {value}
	// formatter: '{value} kg'
	// 2. Use callback function; function parameters are axis index
	// formatter: function (value, index) {
	//    // Formatted to be month/day; display year only in the first label
	//    var date = new Date(value);
	//    var texts = [(date.getMonth() + 1), date.getDate()];
	//    if (index === 0) {
	//        texts.unshift(date.getYear());
	//    }
	//    return texts.join('/');
	// }
	AxisLabel *Label `json:"axisLabel,omitempty"`

	// Category data, available in type: 'category' axis.
	Data interface{} `json:"data,omitempty"`

	// Number of segments that the axis is split into. Note that this number serves only as a
	// recommendation, and the true segments may be adjusted based on readability.
	// This is unavailable for category axis.
	SplitNumber int `json:"splitNumber,omitempty"`

	// It is available only in numerical axis, i.e., type: 'value'.
	// It specifies whether not to contain zero position of axis compulsively.
	// When it is set to be true, the axis may not contain zero position,
	// which is useful in the scatter chart for both value axes.
	// This configuration item is unavailable when the min and max are set.
	Scale bool `json:"scale,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Max interface{} `json:"max,omitempty"`

	// The index of grid which the Y axis belongs to. Defaults to be in the first grid.
	// default 0
	GridIndex int `json:"gridIndex,omitempty"`

	// Split area of Y axis in grid area.
	*SplitArea `json:"splitArea,omitempty"`

	// Split line of Y axis in grid area.
	*SplitLine `json:"splitLine,,omitempty"`
}

// TextStyle is the option set for a text style component.
type TextStyle struct {
	// Font color
	Color string `json:"color,omitempty"`

	// Font style
	// Options: 'normal', 'italic', 'oblique'
	FontStyle string `json:"fontStyle,omitempty"`

	// Font size
	FontSize int `json:"fontSize,omitempty"`

	// compatible for WordCloud
	Normal *TextStyle `json:"normal,omitempty"`
}

// SplitArea is the option set for a split area.
type SplitArea struct {
	// Set this to true to show the splitArea.
	Show bool `json:"show"`

	// Split area style.
	*AreaStyle `json:"areaStyle,omitempty"`
}

// SplitLine is the option set for a split line.
type SplitLine struct {
	// Set this to true to show the splitLine.
	Show bool `json:"show"`

	// Split line style.
	*LineStyle `json:"lineStyle,omitempty"`
}

// VisualMap is a type of component for visual encoding, which maps the data to visual channels.
type VisualMap struct {
	// Mapping type.
	// Options: "continuous", "piecewise"
	Type string `json:"type,omitempty" default:"continuous"`

	// Whether show handles, which can be dragged to adjust "selected range".
	Calculable bool `json:"calculable"`

	// Specify the min dataValue for the visualMap component.
	// [visualMap.min, visualMax.max] make up the domain of visual mapping.
	Min float32 `json:"min,omitempty"`

	// Specify the max dataValue for the visualMap component.
	// [visualMap.min, visualMax.max] make up the domain of visual mapping.
	Max float32 `json:"max,omitempty"`

	// Specify selected range, that is, the dataValue corresponding to the two handles.
	Range []float32 `json:"range,omitempty"`

	// The label text on both ends, such as ['High', 'Low'].
	Text []string `json:"text,omitempty"`

	// Define visual channels that will mapped from dataValues that are in selected range.
	InRange *VisualMapInRange `json:"inRange,omitempty"`
}

// VisualMapInRange is a visual map instance in a range.
type VisualMapInRange struct {
	// Color
	Color []string `json:"color,omitempty"`

	// Symbol type at the two ends of the mark line. It can be an array for two ends, or assigned separately.
	// Options: "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Symbol string `json:"symbol,omitempty"`

	// Symbol size.
	SymbolSize float32 `json:"symbolSize,omitempty"`
}

// DataZoom is the option set for a zoom component.
type DataZoom struct {
	// Data zoom component of inside type, Options: "inside", "slider"
	Type string `json:"type" default:"inside"`

	// The start percentage of the window out of the data extent, in the range of 0 ~ 100.
	// default 0
	Start float32 `json:"start,omitempty"`

	// The end percentage of the window out of the data extent, in the range of 0 ~ 100.
	// default 100
	End float32 `json:"end,omitempty"`

	// Specify the frame rate of views refreshing, with unit millisecond (ms).
	// If animation set as true and animationDurationUpdate set as bigger than 0,
	// you can keep throttle as the default value 100 (or set it as a value bigger than 0),
	// otherwise it might be not smooth when dragging.
	// If animation set as false or animationDurationUpdate set as 0, and data size is not very large,
	// and it seems to be not smooth when dragging, you can set throttle as 0 to improve that.
	Throttle float32 `json:"throttle,omitempty"`

	// Specify which xAxis is/are controlled by the dataZoom-inside when Cartesian coordinate system is used.
	// By default the first xAxis that parallel to dataZoom are controlled when dataZoom-inside.
	// Orient is set as 'horizontal'. But it is recommended to specify it explicitly but not use default value.
	// If it is set as a single number, one axis is controlled, while if it is set as an Array ,
	// multiple axes are controlled.
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`

	// Specify which yAxis is/are controlled by the dataZoom-inside when Cartesian coordinate system is used.
	// By default the first yAxis that parallel to dataZoom are controlled when dataZoom-inside.
	// Orient is set as 'vertical'. But it is recommended to specify it explicitly but not use default value.
	// If it is set as a single number, one axis is controlled, while if it is set as an Array ,
	// multiple axes are controlled.
	YAxisIndex interface{} `json:"yAxisIndex,omitempty"`
}

// SingleAxis is the option set for single axis.
type SingleAxis struct {
	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	Max interface{} `json:"max,omitempty"`

	// Type of axis.
	// Option:
	// * 'value': Numerical axis, suitable for continuous data.
	// * 'category': Category axis, suitable for discrete category data.
	//   Category data can be auto retrieved from series.data or dataset.source,
	//   or can be specified via xAxis.data.
	// * 'time' Time axis, suitable for continuous time series data. As compared to value axis,
	//   it has a better formatting for time and a different tick calculation method. For example,
	//   it decides to use month, week, day or hour for tick based on the range of span.
	// * 'log' Log axis, suitable for log data.
	Type string `json:"type,omitempty"`

	// Distance between grid component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between grid component and the right side of the container.
	//right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	Right string `json:"right,omitempty"`

	// Distance between grid component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between grid component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	Bottom string `json:"bottom,omitempty"`
}

// Indicator is the option set for a radar chart.
type Indicator struct {
	// Indicator name
	Name string `json:"name,omitempty"`

	// The maximum value of indicator. It is an optional configuration, but we recommend to set it manually.
	Max float32 `json:"max,omitempty"`

	// The minimum value of indicator. It it an optional configuration, with default value of 0.
	Min float32 `json:"min,omitempty"`

	// Specify a color the the indicator.
	Color string `json:"color,omitempty"`
}

// RadarComponent is the option set for a radar component.
type RadarComponent struct {
	// Indicator of radar chart, which is used to assign multiple variables(dimensions) in radar chart.
	Indicator []Indicator `json:"indicator,omitempty"`

	// Radar render type, in which 'polygon' and 'circle' are supported.
	Shape string `json:"shape,omitempty"`

	// Segments of indicator axis.
	// default 5
	SplitNumber int `json:"splitNumber,omitempty"`

	// Center position of , the first of which is the horizontal position, and the second is the vertical position.
	// Percentage is supported. When set in percentage, the item is relative to the container width and height.
	Center interface{} `json:"center,omitempty"`

	// Split area of axis in grid area.
	*SplitArea `json:"splitArea,omitempty"`

	// Split line of axis in grid area.
	*SplitLine `json:"splitLine,omitempty"`
}

// GeoComponent is the option set for geo component.
type GeoComponent struct {
	Map       string    `json:"map,omitempty"`
	ItemStyle ItemStyle `json:"itemStyle,omitempty"`
	// Set this to true, to prevent interaction with the axis.
	Silent bool `json:"silent,omitempty"`
}

// ParallelComponent is the option set for parallel component.
type ParallelComponent struct {
	// Distance between parallel component and the left side of the container.
	// Left value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%';
	// and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between parallel component and the top side of the container.
	// Top value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	// and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between parallel component and the right side of the container.
	// Right value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	Right string `json:"right,omitempty"`

	// Distance between parallel component and the bottom side of the container.
	// Bottom value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	Bottom string `json:"bottom,omitempty"`
}

// ParallelAxis is the option set for a parallel axis.
type ParallelAxis struct {
	// Dimension index of coordinate axis.
	Dim int `json:"dim,omitempty"`

	// Name of axis.
	Name string `json:"name,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	// In category axis, it can also be set as the ordinal number.
	Max interface{} `json:"max,omitempty"`

	// Compulsively set segmentation interval for axis.
	Inverse bool `json:"inverse,omitempty"`

	// Location of axis name. Options: "start", "middle", "center", "end"
	// default "end"
	NameLocation string `json:"nameLocation,omitempty"`

	// Type of axis.
	// Options：
	// "value"：Numerical axis, suitable for continuous data.
	// "category" Category axis, suitable for discrete category data. Category data can be auto retrieved from series.
	// "log"  Log axis, suitable for log data.
	Type string `json:"type,omitempty"`

	// Category data，works on (type: "category").
	Data interface{} `json:"data,omitempty"`
}

type JSFunctions struct {
	Fns []string
}

// AddJSFuncs adds a new JS function.
func (f *JSFunctions) AddJSFuncs(fn ...string) {
	for i := 0; i < len(fn); i++ {
		f.Fns = append(f.Fns, replaceJsFuncs(fn[i]))
	}
}

// FuncOpts is the option set for handling function type.
func FuncOpts(fn string) string {
	return replaceJsFuncs(fn)
}

const funcMarker = "__x__"

// replace and clear up js functions string
func replaceJsFuncs(fn string) string {
	pat := regexp.MustCompile(`\n|\t`)
	return fmt.Sprintf("%s%s%s", funcMarker, pat.ReplaceAllString(fn, ""), funcMarker)
}

type Colors []string

// AssetsOpts contains options for static assets.
type Assets struct {
	JSAssets  types.OrderedSet
	CSSAssets types.OrderedSet
}

// InitAssets static assets initial options
func (opt *Assets) InitAssets() {
	opt.JSAssets.Init("echarts.min.js")
	opt.CSSAssets.Init("bulma.min.css")
}

// Validate static assets checks options, append host
func (opt *Assets) Validate(host string) {
	for i := 0; i < len(opt.JSAssets.Values); i++ {
		opt.JSAssets.Values[i] = host + opt.JSAssets.Values[i]
	}
	for j := 0; j < len(opt.CSSAssets.Values); j++ {
		opt.CSSAssets.Values[j] = host + opt.CSSAssets.Values[j]
	}
}

// XAxis3D contains options for X axis in the 3D coordinate.
type XAxis3D struct {
	// Whether to display the x-axis.
	Show bool `json:"show,omitempty"`

	// The name of the axis.
	Name bool `json:"name,omitempty"`

	// The index of the grid3D component used by the axis. The default is to use the first grid3D component.
	Grid3DIndex int `json:"grid3DIndex,omitempty"`

	// The type of the axis.
	// Optional:
	// * 'value' The value axis. Suitable for continuous data.
	// * 'category' The category axis. Suitable for the discrete category data.
	//  For this type, the category data must be set through data.
	// * 'time' The timeline. Suitable for the continuous timing data. The time axis has a
	//  time format compared to the value axis, and the scale calculation is also different.
	//  For example, the scale of the month, week, day, and hour ranges can be determined according to the range of the span.
	// * 'log' Logarithmic axis. Suitable for the logarithmic data.
	Type string `json:"type,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example,
	// if a category axis has data: ['categoryA', 'categoryB', 'categoryC'],
	// and the ordinal 2 represents 'categoryC'. Moreover, it can be set as a negative number, like -3.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of the axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example, if a category axis
	// has data: ['categoryA', 'categoryB', 'categoryC'], and the ordinal 2 represents 'categoryC'.
	// Moreover, it can be set as a negative number, like -3.
	Max interface{} `json:"max,omitempty"`

	// Category data, available in type: 'category' axis.
	// If type is specified as 'category', but axis.data is not specified, axis.data will be auto
	// collected from series.data. It brings convenience, but we should notice that axis.data provides
	// then value range of the 'category' axis. If it is auto collected from series.data,
	// Only the values appearing in series.data can be collected. For example,
	// if series.data is empty, nothing will be collected.
	Data interface{} `json:"data,omitempty"`
}

// YAxis3D contains options for Y axis in the 3D coordinate.
type YAxis3D struct {
	// Whether to display the y-axis.
	Show bool `json:"show,omitempty"`

	// The name of the axis.
	Name bool `json:"name,omitempty"`

	// The index of the grid3D component used by the axis. The default is to use the first grid3D component.
	Grid3DIndex int `json:"grid3DIndex,omitempty"`

	// The type of the axis.
	// Optional:
	// * 'value' The value axis. Suitable for continuous data.
	// * 'category' The category axis. Suitable for the discrete category data.
	//  For this type, the category data must be set through data.
	// * 'time' The timeline. Suitable for the continuous timing data. The time axis has a
	//  time format compared to the value axis, and the scale calculation is also different.
	//  For example, the scale of the month, week, day, and hour ranges can be determined according to the range of the span.
	// * 'log' Logarithmic axis. Suitable for the logarithmic data.
	Type string `json:"type,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example,
	// if a category axis has data: ['categoryA', 'categoryB', 'categoryC'],
	// and the ordinal 2 represents 'categoryC'. Moreover, it can be set as a negative number, like -3.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of the axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example, if a category axis
	// has data: ['categoryA', 'categoryB', 'categoryC'], and the ordinal 2 represents 'categoryC'.
	// Moreover, it can be set as a negative number, like -3.
	Max interface{} `json:"max,omitempty"`

	// Category data, available in type: 'category' axis.
	// If type is specified as 'category', but axis.data is not specified, axis.data will be auto
	// collected from series.data. It brings convenience, but we should notice that axis.data provides
	// then value range of the 'category' axis. If it is auto collected from series.data,
	// Only the values appearing in series.data can be collected. For example,
	// if series.data is empty, nothing will be collected.
	Data interface{} `json:"data,omitempty"`
}

// ZAxis3D contains options for Z axis in the 3D coordinate.
type ZAxis3D struct {
	// Whether to display the z-axis.
	Show bool `json:"show,omitempty"`

	// The name of the axis.
	Name bool `json:"name,omitempty"`

	// The index of the grid3D component used by the axis. The default is to use the first grid3D component.
	Grid3DIndex int `json:"grid3DIndex,omitempty"`

	// The type of the axis.
	// Optional:
	// * 'value' The value axis. Suitable for continuous data.
	// * 'category' The category axis. Suitable for the discrete category data.
	//  For this type, the category data must be set through data.
	// * 'time' The timeline. Suitable for the continuous timing data. The time axis has a
	//  time format compared to the value axis, and the scale calculation is also different.
	//  For example, the scale of the month, week, day, and hour ranges can be determined according to the range of the span.
	// * 'log' Logarithmic axis. Suitable for the logarithmic data.
	Type string `json:"type,omitempty"`

	// The minimum value of axis.
	// It can be set to a special value 'dataMin' so that the minimum value on this axis is set to be the minimum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example,
	// if a category axis has data: ['categoryA', 'categoryB', 'categoryC'],
	// and the ordinal 2 represents 'categoryC'. Moreover, it can be set as a negative number, like -3.
	Min interface{} `json:"min,omitempty"`

	// The maximum value of the axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure the axis tick is equally distributed when not set.
	// In the category axis, it can also be set as the ordinal number. For example, if a category axis
	// has data: ['categoryA', 'categoryB', 'categoryC'], and the ordinal 2 represents 'categoryC'.
	// Moreover, it can be set as a negative number, like -3.
	Max interface{} `json:"max,omitempty"`

	// Category data, available in type: 'category' axis.
	// If type is specified as 'category', but axis.data is not specified, axis.data will be auto
	// collected from series.data. It brings convenience, but we should notice that axis.data provides
	// then value range of the 'category' axis. If it is auto collected from series.data,
	// Only the values appearing in series.data can be collected. For example,
	// if series.data is empty, nothing will be collected.
	Data interface{} `json:"data,omitempty"`
}

// Grid3D contains options for the 3D coordinate.
type Grid3D struct {
	// Whether to show the coordinate.
	Show bool `json:"show,omitempty"`

	// 3D Cartesian coordinates width
	// default 100
	BoxWidth float32 `json:"boxWidth,omitempty"`

	// 3D Cartesian coordinates height
	// default 100
	BoxHeight float32 `json:"boxHeight,omitempty"`

	// 3D Cartesian coordinates depth
	// default 100
	BoxDepth float32 `json:"boxDepth,omitempty"`

	// Rotate or scale fellows the mouse
	ViewControl ViewControl `json:"viewControl,omitempty"`
}

// ViewControl contains options for view controlling.
type ViewControl struct {
	// Whether to rotate automatically.
	AutoRotate bool `json:"autoRotate,omitempty"`

	// Rotate Speed, (angle/s).
	// default 10
	AutoRotateSpeed float32 `json:"autoRotateSpeed,omitempty"`
}
