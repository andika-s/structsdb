package structs

import "errors"

type Structs struct {
	table []*struct{}
}

func NewTable(value []*struct{}) Structs {
	table := new(Structs)

	table.table = value

	return *table
}

func (s *Structs) Insert(value *struct{}) error {

	if value == nil {
		return errors.New("cannot insert empty value")
	}

	s.table = append(s.table, value)

	return nil
}

func (s *Structs) Select() ([]*struct{}, error) {

	if s.table == nil {
		return nil, errors.New("value empty")
	}

	if s.table == nil {
		return nil, errors.New("temporary struct empty")
	}

	return s.table, nil
}

func (s *Structs) Update(value *struct{}) error {

	if value == nil {
		return errors.New("cannot update empty value")
	}

	if s.table == nil {
		return errors.New("temporary struct empty")
	}

	for _, v := range s.table {
		if v == value {
			v = value
		}
	}

	return nil
}

func (s *Structs) Delete(value *struct{}) error {

	if value == nil {
		return errors.New("cannot delete empty value")
	}

	if s.table == nil {
		return errors.New("temporary struct empty")
	}

	delete(s.table, value)

	return nil
}
