package components

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
)

type Layout string

const (
	PageNoneLayout   Layout = "none"
	PageCenterLayout Layout = "center"
	PageFlexLayout   Layout = "flex"
)

// Charter represents a chart value which provides its type, assets and can be validated.
type Charter interface {
	Type() string
	GetAssets() opts.Assets
	Validate()
}

// Page represents a page chart.
type Page struct {
	render.Renderer
	opts.Initialization
	opts.Assets

	Charts []interface{}
	Layout Layout
}

type PageOpts func(p *Page)

func (p *Page) SetPageOptions(opts ...PageOpts) *Page {
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func WithInitializationOpts(opt opts.Initialization) PageOpts {
	return func(p *Page) {
		p.Initialization = opt
		p.Initialization.Validate()
	}
}

// NewPage creates a new page.
func NewPage() *Page {
	page := &Page{}
	page.Assets.InitAssets()
	page.Renderer = render.NewPageRender(page, page.Validate)
	page.Layout = PageCenterLayout
	return page
}

// SetLayout sets the layout of the Page.
func (page *Page) SetLayout(layout Layout) *Page {
	page.Layout = layout
	return page
}

// AddCharts adds new charts to the page.
func (page *Page) AddCharts(charts ...Charter) *Page {
	for i := 0; i < len(charts); i++ {
		assets := charts[i].GetAssets()
		for _, v := range assets.JSAssets.Values {
			page.JSAssets.Add(v)
		}

		for _, v := range assets.CSSAssets.Values {
			page.CSSAssets.Add(v)
		}
		charts[i].Validate()
		page.Charts = append(page.Charts, charts[i])
	}
	return page
}

// Validate validates the given configuration.
func (page *Page) Validate() {
	page.Initialization.Validate()
	page.Assets.Validate(page.AssetsHost)
}
