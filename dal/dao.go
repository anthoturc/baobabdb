package dal

import "github.com/anthoturc/baobabdb/model"

type SuccessorDao interface {
	PutSuccessor(ip string) error
	GetSuccessor(i int) (*model.Node, error)
	DeleteSuccessor(ip string) error
}

type PredecessorDao interface {
	PutPredecessor(ip string) error
	GetPredecessor() (*model.Node, error)
	DeletePredecessor() error
}

type FingerDao interface {
	PutFinger(ip string) error
	GetFinger(i int) (*model.Node, error)
	DeleteFinger(i int) error
}
