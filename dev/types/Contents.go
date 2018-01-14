package types

import "errors"

type Method func() *Redirect

type Page map[string]Method

func (ms Page) SetMethod(key string, m Method) {
	ms[key] = m
}

type Pages map[string]Page

func (p Pages) SetPage(key string, ms Page) {
	p[key] = ms
}

type Contents map[string]Pages

func (c Contents) SetPages(key string, p Pages) {
	c[key] = p
}
func (c Contents) Exec(info *PageInfo) (*Redirect, error) {
	ps, ok := c[info.Contents]
	if !ok {
		return nil, errors.New(info.Contents + ":コンテンツは定義されていません")
	}
	p, ok := ps[info.Page]
	if !ok {
		return nil, errors.New(info.Page + ":ページは定義されていません")
	}
	m, ok := p[info.Method]
	if !ok {
		return nil, errors.New(info.Method + ":メソッドは定義されていません")
	}
	redirect := m()
	return redirect, nil
}
