package lm

import (
	"fmt"
	"testing"
)

func BenchmarkIdentity(b *testing.B) {
	M := &Mat4x4{}
	for n := 0; n < b.N; n++ {
		M.Identity()
	}
}

func BenchmarkMatIdentity(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = MatIdentity()
	}
}

func ExampleMatIdentity() {
	fmt.Println(MatIdentity())
	// Output:
	// [[1 0 0 0] [0 1 0 0] [0 0 1 0] [0 0 0 1]]
}

func ExampleIdentity() {
	M := Mat4x4{
		Vec4{1, 2, 3, 4},
		Vec4{5, 6, 7, 8},
		Vec4{9, 10, 11, 12},
		Vec4{13, 14, 15, 16},
	}
	M.Identity()
	fmt.Println(M)
	// Output:
	// [[1 0 0 0] [0 1 0 0] [0 0 1 0] [0 0 0 1]]
}

func ExampleDup() {
	M := Mat4x4{
		Vec4{1, 2, 3, 4},
		Vec4{5, 6, 7, 8},
		Vec4{9, 10, 11, 12},
		Vec4{13, 14, 15, 16},
	}
	I := MatIdentity()
	M.Dup(I)
	fmt.Println(M)
	// Output:
	// [[1 0 0 0] [0 1 0 0] [0 0 1 0] [0 0 0 1]]
}

func BenchmarkDup(b *testing.B) {
	M := Mat4x4{
		Vec4{1, 2, 3, 4},
		Vec4{5, 6, 7, 8},
		Vec4{9, 10, 11, 12},
		Vec4{13, 14, 15, 16},
	}
	I := MatIdentity()
	for n := 0; n < b.N; n++ {
		M.Dup(I)
	}
}

func ExampleRow() {
	M := Mat4x4{
		Vec4{1, 2, 3, 4},
		Vec4{5, 6, 7, 8},
		Vec4{9, 10, 11, 12},
		Vec4{13, 14, 15, 16},
	}
	fmt.Println(M.Row(3))
	fmt.Println(M.Row(2))
	fmt.Println(M.Row(1))
	fmt.Println(M.Row(0))
	// Output:
	// [4 8 12 16]
	// [3 7 11 15]
	// [2 6 10 14]
	// [1 5 9 13]
}

func ExampleCol() {
	M := Mat4x4{
		Vec4{1, 2, 3, 4},
		Vec4{5, 6, 7, 8},
		Vec4{9, 10, 11, 12},
		Vec4{13, 14, 15, 16},
	}
	fmt.Println(M.Row(3))
	fmt.Println(M.Row(2))
	fmt.Println(M.Row(1))
	fmt.Println(M.Row(0))
	// Output:
	// [4 8 12 16]
	// [3 7 11 15]
	// [2 6 10 14]
	// [1 5 9 13]
}
