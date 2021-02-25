package ArrayList

import (
	"container/list"
)

type ArrayList struct {
	list.List
	index []*list.Element // 维护一个list的索引，方便通过下标方式检索
}

func NewArrayList(row string) *ArrayList {
	c := new(ArrayList)
	c.Init()
	t := getColumns(row)
	for _, i := range t {
		e := c.PushBack(i)
		c.index = append(c.index, e)
	}
	return c
}

func (p *ArrayList) Append(s string) {
	p.index = append(p.index, p.PushBack(s))
}

func (p *ArrayList) IndexOf(index int) *list.Element {
	e := p.Front()
	for i := 0; i < index; i++ {
		e = e.Next()
	}
	return e
}

func (p *ArrayList) Delete(i int) *ArrayList {
	e := p.index[i]
	p.Remove(e)
	p.index = append(p.index[:i], p.index[i+1:]...)
	return p
}

func (p *ArrayList) ToString() string {
	var ret string
	for _, i := range p.index {
		ret += i.Value.(string)
	}
	return ret
}

