package sample

import "math/rand"

func RandomCPUBrand() string {
	return randomStringFromSet(
		"Intel",
		"AMD",
	)
}

func RandomPCBrand() string {
	return randomStringFromSet(
		"Apple",
		"Dell",
		"Acer",
	)
}

func RandomGPUBrand() string {
	return randomStringFromSet(
		"NVIDIA",
		"AMD",
	)
}

func RandomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"intel-i3",
			"intel-i5",
			"intel-i7",
			"intel-i9",
		)
	}

	return randomStringFromSet(
		"rizen-3-PRO",
		"rizen-5-PRO",
		"rizen-7-PRO",
	)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomBool() bool {
	return rand.Intn(2) == 0
}

func randomStringFromSet(s ...string) string {
	n := len(s)

	return s[rand.Intn(n)]
}
