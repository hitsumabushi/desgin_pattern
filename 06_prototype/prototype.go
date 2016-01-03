package prototype

import "errors"

// clone()はManagerにcloneの役割を一任したい場合には、unexportedにしておく
type product interface {
	Use(s string) string
	clone() product
}

type Manager struct {
	showcase map[string]product
}

func (self *Manager) Register(name string, product product) {
	self.showcase[name] = product
}

func (self *Manager) Create(name string) (product, error) {
	var err error
	result, ok := self.showcase[name]
	if !ok {
		err = errors.New("Not registerd")
		return nil, err
	}
	return result.clone(), nil
}

type LinePen struct {
	lchar string
}

func (self *LinePen) Use(s string) string {
	return self.lchar + s + self.lchar
}

func (self *LinePen) clone() product {
	return &LinePen{self.lchar}
}
