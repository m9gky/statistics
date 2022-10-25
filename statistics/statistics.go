package statistics

import (
	"github.com/klauspost/cpuid/v2"
	"github.com/stretchr/testify/mock"
	"math"
)

func Avg(nums []float64) float64 {
	var sum float64
	for _, n := range nums {
		sum += n
	}
	return sum / float64(len(nums))
}

func Fact(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	return f
}

type Point struct {
	X, Y int
}

// Distance вычисляет расстояние между двумя точками.
func Distance(a, b Point) float64 {

	if a.X < 0 || b.X < 0 || a.Y < 0 || b.Y < 0 {
		return 0
	}
	return math.Sqrt(float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)))
}

type CPUInformer interface {
	IsVendor(v cpuid.Vendor) bool
	LogicalCPU() int
}

func IsValidCPU(i CPUInformer) bool {
	return (i.IsVendor(cpuid.AMD) || i.IsVendor(cpuid.Intel)) && i.LogicalCPU() >= 6
}

type MockedInformer struct {
	mock.Mock
}

func (m *MockedInformer) IsVendor(v cpuid.Vendor) bool {
	args := m.Called(v)
	return args.Bool(0)
}

func (m *MockedInformer) LogicalCPU() int {
	args := m.Called()
	return args.Int(0)
}
