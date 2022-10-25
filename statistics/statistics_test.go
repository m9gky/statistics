package statistics

import (
	"github.com/klauspost/cpuid/v2"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAvg(t *testing.T) {
	if Avg([]float64{1, 2, 3}) != 2 {
		t.Errorf("Avg of 1, 2, 3 expected to be 2")
	}
}

func TestFact(t *testing.T) {
	if Fact(5) != 120 {
		t.Errorf("Factorial of 5 is expected to be 120")
	}
}

func TestDistance(t *testing.T) {
	require := require.New(t)

	require.Equal(5.0, Distance(Point{1, 1}, Point{4, 5}), "Пример 1")
	require.Equal(0.0, Distance(Point{1, -2}, Point{4, 5}), "Пример с отрицательными координатами")
}

func TestIsValidCPU(t *testing.T) {
	i := new(MockedInformer)
	i.On("IsVendor", cpuid.Intel).Return(false)
	i.On("IsVendor", cpuid.AMD).Return(true)
	i.On("LogicalCPU").Return(3)

	require := require.New(t)
	require.False(IsValidCPU(i))

	i.AssertCalled(t, "IsVendor", cpuid.AMD)
	i.AssertCalled(t, "LogicalCPU")
}
