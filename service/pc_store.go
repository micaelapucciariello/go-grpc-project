package service

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/micaelapucciariello/grpc-project/pb"
	"log"
	"strconv"
	"sync"
)

type PCStore interface {
	Save(pc *pb.PC) error
	Find(id string) (*pb.PC, error)
	Search(ctx context.Context, filter *pb.Filter, found func(pc *pb.PC) error) error
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

	other, err := deepCopy(pc)
	if err != nil {
		return err
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

	return deepCopy(pc)
}

func (store *InMemoryPCStore) Search(ctx context.Context, filter *pb.Filter, found func(pc *pb.PC) error) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	for _, pc := range store.data {
		if isQualified(pc, filter) {
			other, err := deepCopy(pc)
			if err != nil {
				return err
			}

			err = found(other)
			if err != nil {
				return err
			}
		}

		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			log.Print("context is cancelled")
			return errors.New("context is cancelled")
		}
	}

	return nil
}

func isQualified(pc *pb.PC, filter *pb.Filter) bool {
	if pc.UsdPrice > filter.MaxPriceUsd {
		return false
	}
	if pc.Cpu.MinGhz < filter.MinCpuGhz {
		return false
	}
	if pc.Cpu.Cores < filter.MinCpuCores {
		return false
	}
	if toBit(pc.Memory[0]) < toBit(filter.MinRam) {
		return false
	}
	return true
}

func toBit(memory *pb.Memory) uint64 {
	value, _ := strconv.ParseUint(memory.Value, 10, 64)
	switch memory.GetUnit() {
	case pb.Memory_BIT:
		return value
	case pb.Memory_BYTE:
		return value << 3 // 8 = 2 sq3
	case pb.Memory_KBYTE:
		return value << 13
	case pb.Memory_MBYTE:
		return value << 23
	case pb.Memory_GBYTE:
		return value << 33
	default:
		return 0
	}
}

func deepCopy(pc *pb.PC) (*pb.PC, error) {
	// other
	other := &pb.PC{}
	err := copier.Copy(other, pc)
	if err != nil {
		return nil, ErrCopyingItem
	}

	return other, nil
}
