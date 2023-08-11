package sample

import "github.com/micaelapucciariello/grpc-project/pb"

func NewCPU() *pb.CPU {
	brand := RandomCPUBrand()
	cores := RandomInt(2, 8)
	threads := RandomInt(cores, 16)

	minGhz := RandomFloat64(2.5, 3)
	maxGhz := RandomFloat64(minGhz, 5)

	return &pb.CPU{
		Brand:   brand,
		Name:    RandomCPUName(brand),
		Cores:   uint32(cores),
		Threads: uint32(threads),
		MinGhz:  minGhz,
		MaxGhz:  maxGhz,
	}
}

func NewGPU() *pb.GPU {
	cores := RandomInt(2, 8)
	threads := RandomInt(cores, 16)

	minGhz := RandomFloat64(2.5, 3)
	maxGhz := RandomFloat64(minGhz, 5)

	memory := NewMemory()

	return &pb.GPU{
		Brand:   "",
		Name:    "",
		Cores:   uint32(cores),
		Threads: uint32(threads),
		MinGhz:  minGhz,
		MaxGhz:  maxGhz,
		Memory:  memory,
	}
}

func NewMemory() *pb.Memory {
	return &pb.Memory{
		Value: "",
		Unit:  pb.Memory_Unit(RandomInt(0, 2)),
	}
}
