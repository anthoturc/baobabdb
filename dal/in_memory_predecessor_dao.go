package dal

import (
	"errors"
	"github.com/anthoturc/baobabdb/model"
)

type InMemoryPredecessorDao struct {
	Predecessor *model.Node
}

func (impd *InMemoryPredecessorDao) PutPredecessor(ip string) error {
	if impd.Predecessor != nil {
		return errors.New("cannot have more than one predecessor")
	}
	impd.Predecessor = model.NewNode(ip)
	return nil
}

func (impd *InMemoryPredecessorDao) GetPredecessor() (*model.Node, error) {
	return impd.Predecessor, nil
}

func (impd *InMemoryPredecessorDao) DeletePredecessor() error {
	impd.Predecessor = nil
	return nil
}
