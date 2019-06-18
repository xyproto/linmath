package linmath

import "math"

type (
	Vec2 [2]float64
	Vec3 [3]float64
	Vec4 [4]float64
)

func Vec2Add(a, b Vec2) (r Vec2) {
	for i := 0; i < 2; i++ {
		r[i] = a[i] + b[i]
	}
	return r
}

func Vec3Add(a, b Vec3) (r Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = a[i] + b[i]
	}
	return r
}

func Vec4Add(a, b Vec4) (r Vec4) {
	for i := 0; i < 4; i++ {
		r[i] = a[i] + b[i]
	}
	return r
}

func Vec2Sub(a, b Vec2) (r Vec2) {
	for i := 0; i < 2; i++ {
		r[i] = a[i] - b[i]
	}
	return r
}

func Vec3Sub(a, b Vec3) (r Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = a[i] - b[i]
	}
	return r
}

func Vec4Sub(a, b Vec4) (r Vec4) {
	for i := 0; i < 4; i++ {
		r[i] = a[i] - b[i]
	}
	return r
}

func Vec2Scale(v Vec2, s float64) (r Vec2) {
	for i := 0; i < 2; i++ {
		r[i] = v[i] * s
	}
	return r
}

func Vec3Scale(v Vec3, s float64) (r Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = v[i] * s
	}
	return r
}

func Vec4Scale(v Vec4, s float64) (r Vec4) {
	for i := 0; i < 4; i++ {
		r[i] = v[i] * s
	}
	return r
}

func Vec2MulInner(a, b Vec2) (r float64) {
	for i := 0; i < 2; i++ {
		r += a[i] * b[i]
	}
	return r
}

func Vec3MulInner(a, b Vec3) (r float64) {
	for i := 0; i < 3; i++ {
		r += a[i] * b[i]
	}
	return r
}

func Vec4MulInner(a, b Vec4) (r float64) {
	for i := 0; i < 4; i++ {
		r += a[i] * b[i]
	}
	return r
}

func Vec2Len(v Vec2) float64 {
	return math.Sqrt(Vec2MulInner(v, v))
}

func Vec3Len(v Vec3) float64 {
	return math.Sqrt(Vec3MulInner(v, v))
}

func Vec4Len(v Vec4) float64 {
	return math.Sqrt(Vec4MulInner(v, v))
}

func Vec2Norm(v Vec2) Vec2 {
	return Vec2Scale(v, 1.0/Vec2Len(v))
}

func Vec3Norm(v Vec3) Vec3 {
	return Vec3Scale(v, 1.0/Vec3Len(v))
}

func Vec4Norm(v Vec4) Vec4 {
	return Vec4Scale(v, 1.0/Vec4Len(v))
}

func Vec2Min(a, b Vec2) (r Vec2) {
	for i := 0; i < 2; i++ {
		r[i] = math.Min(a[i], b[i])
	}
	return r
}

func Vec3Min(a, b Vec3) (r Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = math.Min(a[i], b[i])
	}
	return r
}

func Vec4Min(a, b Vec4) (r Vec4) {
	for i := 0; i < 4; i++ {
		r[i] = math.Min(a[i], b[i])
	}
	return r
}

func Vec2Max(a, b Vec2) (r Vec2) {
	for i := 0; i < 2; i++ {
		r[i] = math.Max(a[i], b[i])
	}
	return r
}

func Vec3Max(a, b Vec3) (r Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = math.Max(a[i], b[i])
	}
	return r
}

func Vec4Max(a, b Vec4) (r Vec4) {
	for i := 0; i < 4; i++ {
		r[i] = math.Max(a[i], b[i])
	}
	return r
}

func Vec3MulCross(a, b Vec3) (r Vec3) {
	r[0] = a[1]*b[2] - a[2]*b[1]
	r[1] = a[2]*b[0] - a[0]*b[2]
	r[2] = a[0]*b[1] - a[1]*b[0]
	return r
}

func Vec3Reflect(v, n Vec3) (r Vec3) {
	p := 2.0 * Vec3MulInner(v, n)
	for i := 0; i < 3; i++ {
		r[i] = v[i] - p*n[i]
	}
	return r
}

func Vec4MulCross(a, b Vec4) (r Vec4) {
	r[0] = a[1]*b[2] - a[2]*b[1]
	r[1] = a[2]*b[0] - a[0]*b[2]
	r[2] = a[0]*b[1] - a[1]*b[0]
	r[3] = 1.0
	return r
}

func Vec4Reflect(v, n Vec4) (r Vec4) {
	p := 2.0 * Vec4MulInner(v, n)
	for i := 0; i < 4; i++ {
		r[i] = v[i] - p*n[i]
	}
	return r
}

// Vec34 creates a Vec3 from a given Vec4
func Vec34(v Vec4) (r Vec3) {
	r[0] = v[0]
	r[1] = v[1]
	r[2] = v[2]
	// Drop v[3]
	return r
}

// Vec43 creates a Vec4 from a given Vec3
func Vec43(v Vec3) (r Vec4) {
	r[0] = v[0]
	r[1] = v[1]
	r[2] = v[2]
	// There is no v[3], use 0
	r[3] = 0
	return r
}
