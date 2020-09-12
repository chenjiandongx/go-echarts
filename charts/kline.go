package charts

import (
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// Kline represents a kline chart.
type Kline struct {
	RectChart
}

func (Kline) Type() string { return types.ChartKline }

// NewKLine creates a new kline chart.
func NewKLine() *Kline {
	c := &Kline{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.HasXYAxis = true
	return c
}

// SetXAxis adds the X axis.
func (c *Kline) SetXAxis(xAxis interface{}) *Kline {
	c.xAxisData = xAxis
	return c
}

// AddSeries adds the new series.
func (c *Kline) AddSeries(name string, data []opts.KlineData, opts ...SeriesOpts) *Kline {
	series := SingleSeries{Name: name, Type: types.ChartKline, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// TODO: add more Kline validate cases
func (c *Kline) Validate() {
	c.XAxisList[0].Data = c.xAxisData
	c.Assets.Validate(c.AssetsHost)
}