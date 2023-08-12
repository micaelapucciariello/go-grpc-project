package sample

import "github.com/micaelapucciariello/grpc-project/pb"

func NewCPU() *pb.CPU {
	brand := RandomCPUBrand()
	cores := RandomInt(2, 8)
	threads := RandomInt(cores, 16)

	minGhz := RandomFloat64(2, 3.5)
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
	brand := RandomGPUBrand()

	cores := RandomInt(2, 8)
	threads := RandomInt(cores, 16)

	minGhz := RandomFloat64(1, 1.5)
	maxGhz := RandomFloat64(minGhz, 2.0)

	memory := NewMemory()

	return &pb.GPU{
		Brand:   brand,
		Name:    brand,
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

func NewSSD() *pb.Storage {
	memory := NewMemory()
	return &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: memory,
	}
}

func NewHDD() *pb.Storage {
	memory := NewMemory()
	return &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: memory,
	}
}

func NewScreen() *pb.Screen {
	width := RandomInt(1024, 4320)
	heigth := width * (9 / 16)
	return &pb.Screen{
		SizeInch: float32(RandomInt(10, 24)),
		Resolution: &pb.Screen_Resolution{
			Width:  uint32(width),
			Height: uint32(heigth),
		},
		Panel:      0,
		Multitouch: RandomBool(),
	}
}
