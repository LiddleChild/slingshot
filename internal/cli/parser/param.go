package parser

type Param struct {
	index  int
	params []string
}

func NewParam(params []string) *Param {
	return &Param{
		0,
		params,
	}
}

func (p *Param) HasNext() bool {
	return p.index < len(p.params)
}

func (p *Param) Next() (string, bool) {
	if !p.HasNext() {
		return "", false
	}

	param := p.params[p.index]
	p.index++

	return param, true
}
