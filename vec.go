package lm

import "math"

type (
	Vec2 [2]float64
	Vec3 [3]float64
	Vec4 [4]float64
)

func (v Vec2) Add(b Vec2) (r Vec2) {
	for i := 0; i < 2; i++ {
		r[i] = v[i] + b[i]
	}
	return r
}

func (v Vec3) Add(b Vec3) (r Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = v[i] + b[i]
	}
	return r
}

func (v Vec4) Add(b Vec4) (r Vec4) {
	for i := 0; i < 4; i++ {
		r[i] = v[i] + b[i]
	}
	return r
}

func (v Vec2) Sub(b Vec2) (r Vec2) {
	for i := 0; i < 2; i++ {
		r[i] = v[i] - b[i]
	}
	return r
}

func (v Vec3) Sub(b Vec3) (r Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = v[i] - b[i]
	}
	return r
}

func (v Vec4) Sub(b Vec4) (r Vec4) {
	for i := 0; i < 4; i++ {
		r[i] = v[i] - b[i]
	}
	return r
}

func (v Vec2) Scale(s float64) (r Vec2) {
	for i := 0; i < 2; i++ {
		r[i] = v[i] * s
	}
	return r
}

func (v Vec3) Scale(s float64) (r Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = v[i] * s
	}
	return r
}

func (v Vec4) Scale(s float64) (r Vec4) {
	for i := 0; i < 4; i++ {
		r[i] = v[i] * s
	}
	return r
}

func (v Vec2) MulInner(b Vec2) (r float64) {
	for i := 0; i < 2; i++ {
		r += v[i] * b[i]
	}
	return r
}

func (v Vec3) MulInner(b Vec3) (r float64) {
	for i := 0; i < 3; i++ {
		r += v[i] * b[i]
	}
	return r
}

func (v Vec4) MulInner(b Vec4) (r float64) {
	for i := 0; i < 4; i++ {
		r += v[i] * b[i]
	}
	return r
}

func (v Vec2) Len() float64 {
	return math.Sqrt(v.MulInner(v))
}

func (v Vec3) Len() float64 {
	return math.Sqrt(v.MulInner(v))
}

func (v Vec4) Len() float64 {
	return math.Sqrt(v.MulInner(v))
}

func (v Vec2) Norm() Vec2 {
	return v.Scale(1.0 / v.Len())
}

func (v Vec3) Norm() Vec3 {
	return v.Scale(1.0 / v.Len())
}

func (v Vec4) Norm() Vec4 {
	return v.Scale(1.0 / v.Len())
}

func (v Vec2) Min(b Vec2) (r Vec2) {
	for i := 0; i < 2; i++ {
		r[i] = math.Min(v[i], b[i])
	}
	return r
}

func (v Vec3) Min(b Vec3) (r Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = math.Min(v[i], b[i])
	}
	return r
}

func (v Vec4) Min(b Vec4) (r Vec4) {
	for i := 0; i < 4; i++ {
		r[i] = math.Min(v[i], b[i])
	}
	return r
}

func (v Vec2) Max(b Vec2) (r Vec2) {
	for i := 0; i < 2; i++ {
		r[i] = math.Max(v[i], b[i])
	}
	return r
}

func (v Vec3) Max(b Vec3) (r Vec3) {
	for i := 0; i < 3; i++ {
		r[i] = math.Max(v[i], b[i])
	}
	return r
}

func (v Vec4) Max(b Vec4) (r Vec4) {
	for i := 0; i < 4; i++ {
		r[i] = math.Max(v[i], b[i])
	}
	return r
}

func (v Vec3) MulCross(b Vec3) (r Vec3) {
	r[0] = v[1]*b[2] - v[2]*b[1]
	r[1] = v[2]*b[0] - v[0]*b[2]
	r[2] = v[0]*b[1] - v[1]*b[0]
	return r
}

func (v Vec3) Reflect(n Vec3) (r Vec3) {
	p := 2.0 * v.MulInner(n)
	for i := 0; i < 3; i++ {
		r[i] = v[i] - p*n[i]
	}
	return r
}

func (v Vec4) MulCross(b Vec4) (r Vec4) {
	r[0] = v[1]*b[2] - v[2]*b[1]
	r[1] = v[2]*b[0] - v[0]*b[2]
	r[2] = v[0]*b[1] - v[1]*b[0]
	r[3] = 1.0
	return r
}

func (v Vec4) Reflect(n Vec4) (r Vec4) {
	p := 2.0 * v.MulInner(n)
	for i := 0; i < 4; i++ {
		r[i] = v[i] - p*n[i]
	}
	return r
}

// Vec3 creates a Vec3 from a given Vec4
func (v Vec4) Vec3() (r Vec3) {
	r[0] = v[0]
	r[1] = v[1]
	r[2] = v[2]
	// Drop v[3]
	return r
}

// Vec4 creates a Vec4 from a given Vec3
func (v Vec3) Vec4() (r Vec4) {
	r[0] = v[0]
	r[1] = v[1]
	r[2] = v[2]
	// There is no v[3], use 0
	r[3] = 0
	return r
}
