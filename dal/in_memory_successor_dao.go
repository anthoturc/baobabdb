package dal

import (
	"errors"
	"fmt"
	"github.com/anthoturc/baobabdb/model"
)

const MaxSuccessors = 1

type InMemorySuccessorDao struct {
	Successors []*model.Node
}

func (imsd *InMemorySuccessorDao) PutSuccessor(ip string) error {
	node := model.NewNode(ip)
	if len(imsd.Successors) == MaxSuccessors {
		return errors.New(fmt.Sprintf("cannot exceed %d successor", MaxSuccessors))
	}
	imsd.Successors = append(imsd.Successors, node)
	return nil
}

func (imsd *InMemorySuccessorDao) GetSuccessors(i int) (*model.Node, error) {
	return imsd.Successors[i], nil
}

func (imsd *InMemorySuccessorDao) DeleteSuccessor(ip string) error {
	return nil
}
