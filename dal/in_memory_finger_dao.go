package dal

import (
	"errors"
	"fmt"
	"github.com/anthoturc/baobabdb/model"
)

const MaxFingers = 0

type InMemoryFigerDao struct {
	Fingers []*model.Node
}

func (imfd *InMemoryFigerDao) PutFinger(ip string) error {
	if len(imfd.Fingers) == MaxFingers {
		return errors.New(fmt.Sprintf("cannot exceed %d fingers", MaxFingers))
	}
	imfd.Fingers = append(imfd.Fingers, model.NewNode(ip))
	return nil
}

func (imfd *InMemoryFigerDao) GetFinger(i int) (*model.Node, error) {
	if i < 0 || i > len(imfd.Fingers) {
		return nil, errors.New("invalid finger entry")
	}
	return imfd.Fingers[i], nil
}

func (imfd *InMemoryFigerDao) DeleteFinger(i int) error {
	if i < 0 || i > len(imfd.Fingers) {
		return errors.New("inavlid finger entry")
	}
	imfd.Fingers = append(imfd.Fingers[:i], imfd.Fingers[i+1:]...)
	return nil
}
