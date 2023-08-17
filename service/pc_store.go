package service

import (
	"github.com/jinzhu/copier"
	"github.com/micaelapucciariello/grpc-project/pb"
	"sync"
)

type PCStore interface {
	Save(pc *pb.PC) error
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

	if store.data[pc.Id] == nil {
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
