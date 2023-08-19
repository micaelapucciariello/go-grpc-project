package service

import (
	"github.com/jinzhu/copier"
	"github.com/micaelapucciariello/grpc-project/pb"
	"sync"
)

type PCStore interface {
	Save(pc *pb.PC) error
	Find(id string) (*pb.PC, error)
}

type InMemoryPCStore struct {
	mutex sync.Mutex
	data  map[string]*pb.PC
}

func NewInMemoryPCStore() *InMemoryPCStore {
	return &InMemoryPCStore{
		data: make(map[string]*pb.PC),
	}
}

func (store *InMemoryPCStore) Save(pc *pb.PC) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[pc.Id] != nil {
		return ErrAlreadyExists
	}

	// other
	other := &pb.PC{}
	err := copier.Copy(other, pc)
	if err != nil {
		return ErrCopyingItem
	}

	store.data[other.Id] = other
	return nil
}

func (store *InMemoryPCStore) Find(id string) (*pb.PC, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	pc := store.data[id]

	if pc == nil {
		return nil, ErrNotExists
	}

	other := &pb.PC{}

	err := copier.Copy(other, pc)
	if err != nil {
		return nil, ErrCopyingItem
	}

	return other, nil
}
