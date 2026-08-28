package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before running all test in this package")
	m.Run()
	fmt.Println("After running all test in this package")
}

func TestAdd(t *testing.T) {
	t.Run("PositiveInt", func(t *testing.T) {
		result := Add(1, 2)
		if result != 3 {
			t.Error("Result of add operation should 3")
		}
	})
	t.Run("PositiveIntAssert", func(t *testing.T) {
		result := Add(1, 2)
		assert.Equal(t, 3, result, "Result must be 3")
		fmt.Println("Test executed") // assert using t.Fail, so it wont break the codes below
	})
	t.Run("PositiveIntRequire", func(t *testing.T) {
		result := Add(1, 2)
		require.Equal(t, 3, result, "Result must be 3")
		fmt.Println("Test executed") // require using t.FailNow, so it will break the codes below
	})
	t.Run("NegativeInt", func(t *testing.T) {
		result := Add(-1, 2)
		if result != 1 {
			t.Fatal("Result of add operation should 1")
		}
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Test can not run on Mac")
	}

	result := Add(1, 2)
	assert.Equal(t, 3, result, "Result must be 3")
}

func BenchmarkMinus(b *testing.B) {
	var result int
	for i := 0; b.Loop(); i++ {
		result = Minus(i+1, i)
	}
	_ = result
}
func TestMinus(t *testing.T) {
	t.Run("PositiveIntAssert", func(t *testing.T) {
		result := Minus(2, 1)
		assert.Equal(t, 1, result, "Result must be 1")
	})

	t.Run("NegativeIntAssert", func(t *testing.T) {
		result := Minus(-2, 1)
		assert.Equal(t, -3, result, "Result must be -3")
	})

	t.Run("PositiveFloatAssert", func(t *testing.T) {
		result := Minus(2.5, 1)
		assert.Equal(t, 1.5, result, "Result must be 1.5")
	})

	t.Run("NegativeFloatAssert", func(t *testing.T) {
		result := Minus(-2.5, 1)
		assert.Equal(t, -3.5, result, "Result must be -3.5")
	})
}

func TestMultiply(t *testing.T) {
	t.Run("PositiveIntAssert", func(t *testing.T) {
		result := Multiply(2, 1)
		assert.Equal(t, 2, result, "Result must be 2")
	})

	t.Run("NegativeIntAssert", func(t *testing.T) {
		result := Multiply(-2, 1)
		assert.Equal(t, -2, result, "Result must be -2")
	})

	t.Run("PositiveFloatAssert", func(t *testing.T) {
		result := Multiply(2.5, 1)
		assert.Equal(t, 2.5, result, "Result must be 2.5")
	})

	t.Run("NegativeFloatAssert", func(t *testing.T) {
		result := Multiply(-2.5, 1)
		assert.Equal(t, -2.5, result, "Result must be -2.5")
	})
}

type NumericConstraint interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type TestCasesMultiplyType[T NumericConstraint] struct {
	name     string
	a        T
	b        T
	expected T
	error    string
}

func BenchmarkMultiply(b *testing.B) {
	testCases := []TestCasesMultiplyType[float64]{
		{
			name: "PositiveIntAssert",
			a:    2.0,
			b:    1.0,
		},
		{
			name: "NegativeIntAssert",
			a:    -2.0,
			b:    1.0,
		},
		{
			name: "PositiveFloatAssert",
			a:    2.5,
			b:    1.0,
		},
		{
			name: "NegativeFloatAssert",
			a:    -2.5,
			b:    1.0,
		},
	}
	for _, test := range testCases {
		b.Run(test.name, func(b *testing.B) {
			var result int
			for i := 0; b.Loop(); i++ {
				result = Multiply(int(test.a), int(test.b))
			}
			_ = result
		})
	}
}
func TestTableMultiply(t *testing.T) {
	testCases := []TestCasesMultiplyType[float64]{
		{
			name:     "PositiveIntAssert",
			a:        2.0,
			b:        1.0,
			expected: 2.0,
			error:    "Result must be 2",
		},
		{
			name:     "NegativeIntAssert",
			a:        -2.0,
			b:        1.0,
			expected: -2.0,
			error:    "Result must be -2",
		},
		{
			name:     "PositiveFloatAssert",
			a:        2.5,
			b:        1.0,
			expected: 2.5,
			error:    "Result must be 2.5",
		},
		{
			name:     "NegativeFloatAssert",
			a:        -2.5,
			b:        1.0,
			expected: -2.5,
			error:    "Result must be -2.5",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := Multiply(test.a, test.b)
			assert.Equal(t, test.expected, result, test.error)
		})
	}
}
