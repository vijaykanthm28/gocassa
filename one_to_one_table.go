package cmagic 

import(
	"fmt"
	"errors"
)

type oneToOne struct {
	t Table
	idField string
}

func (o *oneToOne) Set(v interface{}) error {
	return o.t.Set(v)
}

func (o *oneToOne) Update(id string, m map[string]interface{}) error {
	return o.t.Where(Eq(o.idField, id)).Update(m)
}

func (o *oneToOne) Delete(id string) error {
	return o.t.Where(Eq(o.idField, id)).Delete()
}

func (o *oneToOne) Read(id string) (interface{}, error) {
	res, err := o.t.Where(Eq(o.idField, id)).Query().Read()
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New(fmt.Sprintf("Row with id %v not found", id))
	}
	return res[0], nil
}