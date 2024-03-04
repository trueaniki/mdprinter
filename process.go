package mdprinter

type Processor struct {
	jsonData  interface{}
	style     string
	align     string
	customCss string
}

func New() *Processor {
	return &Processor{
		jsonData:  nil,
		style:     "air",
		align:     "left",
		customCss: "none",
	}
}

func (p *Processor) Process(md []byte) ([]byte, error) {
	return Print(p.generateHtml(md))
}

func (p *Processor) Html(md []byte) []byte {
	return p.generateHtml(md)
}

func (p *Processor) generateHtml(md []byte) []byte {
	html := Parse(md)
	if p.jsonData != nil {
		html = Interpolate(html, p.jsonData)
	}
	css := FormupCss(p.style, p.align, p.customCss)
	html = append(html, []byte("<style>"+css+"</style>")...)
	return html
}

func (p *Processor) WithInterpolation(jsonData interface{}) *Processor {
	p.jsonData = jsonData
	return p
}

func (p *Processor) WithStyle(style string) *Processor {
	p.style = style
	return p
}

func (p *Processor) WithAlign(align string) *Processor {
	p.align = align
	return p
}

func (p *Processor) WithCustomCss(css string) *Processor {
	p.customCss = css
	return p
}
