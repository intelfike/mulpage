package types

import "errors"

type Method func() *Redirect

type Package map[string]Method

func (ms Package) SetMethod(key string, m Method) {
	ms[key] = m
}

type Content map[string]Package

func (p Content) SetPackage(key string, ms Package) {
	p[key] = ms
}

type Contents map[string]Content

func (c Contents) SetContent(key string, p Content) {
	c[key] = p
}
func (c Contents) Exec(info *PageInfo) (*Redirect, error) {
	ps, ok := c[info.Contents]
	if !ok {
		return nil, errors.New(info.Contents + ":コンテンツは定義されていません")
	}
	p, ok := ps[info.Package]
	if !ok {
		return nil, errors.New(info.Package + ":ページは定義されていません")
	}
	m, ok := p[info.Method]
	if !ok {
		return nil, errors.New(info.Method + ":メソッドは定義されていません")
	}
	redirect := m()
	return redirect, nil
}
