package linmath

import "math"

type Quat [4]float64

func QuatIdentity() (r Quat) {
	r[0] = 0
	r[1] = 0
	r[2] = 0
	r[3] = 1.0
	return r
}

func (q Quat) Add(b Quat) (r Quat) {
	for i := 0; i < 4; i++ {
		r[i] = q[i] + b[i]
	}
	return r
}

func (q Quat) Sub(b Quat) (r Quat) {
	for i := 0; i < 4; i++ {
		r[i] = q[i] - b[i]
	}
	return r
}

func (q Quat) Vec3() (r Vec3) {
	r[0] = q[0]
	r[1] = q[1]
	r[2] = q[2]
	return r
}

func (q Quat) Mul(b Quat) (r Quat) {
	v := q.Vec3().MulCross(b.Vec3()).Add(q.Vec3().Scale(b[3])).Add(b.Vec3().Scale(q[3]))
	r[0] = v[0]
	r[1] = v[1]
	r[2] = v[2]
	r[3] = q[3]*b[3] - q.Vec3().MulInner(b.Vec3())
	return r
}

func (q Quat) Scale(s float64) (r Quat) {
	for i := 0; i < 4; i++ {
		r[i] = q[i] * s
	}
	return r
}

func (q Quat) InnerProduct(b Quat) (p float64) {
	for i := 0; i < 4; i++ {
		p += q[i] * b[i]
	}
	return p
}

func (q Quat) Conj() (r Quat) {
	for i := 0; i < 3; i++ {
		r[i] = -q[i]
	}
	r[3] = q[3]
	return r
}

func QuatRotate(angle float64, axis Vec3) (r Quat) {
	v := axis.Scale(math.Sin(angle / 2))
	for i := 0; i < 3; i++ {
		r[i] = v[i]
	}
	r[3] = math.Cos(angle / 2)
	return r
}

func (v Vec4) Quat() (q Quat) {
	return Quat(v)
}

func (q Quat) Vec4() (v Vec4) {
	return Vec4(q)
}

func (q Quat) Norm() Quat {
	return q.Vec4().Norm().Quat()
}

func (q Quat) MulVec3(v Vec3) (r Vec3) {
	/*
	   Method by Fabian 'ryg' Giessen (of Farbrausch)
	   t = 2 * cross(q.xyz, v)
	   v' = v + q.w * t + cross(q.xyz, t)
	*/
	qXyz := Vec3{q[0], q[1], q[2]}

	u := Vec3{q[0], q[1], q[2]}
	t := qXyz.MulCross(v).Scale(2)
	u = qXyz.MulCross(t)
	t = t.Scale(q[3])

	return v.Add(t).Add(u)
}

func (q Quat) Mat4x4() (M Mat4x4) {
	a := q[3]
	b := q[0]
	c := q[1]
	d := q[2]

	a2 := a * a
	b2 := b * b
	c2 := c * c
	d2 := d * d

	M[0][0] = a2 + b2 - c2 - d2
	M[0][1] = 2.0 * (b*c + a*d)
	M[0][2] = 2.0 * (b*d - a*c)
	M[0][3] = 0

	M[1][0] = 2 * (b*c - a*d)
	M[1][1] = a2 - b2 + c2 - d2
	M[1][2] = 2.0 * (c*d + a*b)
	M[1][3] = 0

	M[2][0] = 2.0 * (b*d + a*c)
	M[2][1] = 2.0 * (c*d - a*b)
	M[2][2] = a2 - b2 - c2 + d2
	M[2][3] = 0

	M[3][0] = 0
	M[3][1] = 0
	M[3][2] = 0
	M[3][3] = 1.0

	return M
}

func (M *Mat4x4) MulQuat(RM Mat4x4, q Quat) {
	/*  XXX: The way this is written only works for othogonal matrices. */
	/* TODO: Take care of non-orthogonal case. */
	(*M)[0] = q.MulVec3(RM[0].Vec3()).Vec4()
	(*M)[1] = q.MulVec3(RM[1].Vec3()).Vec4()
	(*M)[2] = q.MulVec3(RM[2].Vec3()).Vec4()
	(*M)[3][0] = 0
	(*M)[3][1] = 0
	(*M)[3][2] = 0
	(*M)[3][3] = 1.0
}

func (M *Mat4x4) Quat() (q Quat) {
	r := 0.0

	for i := 0; i < 3; i++ {
		m := (*M)[i][i]
		if m < r {
			continue
		}
		m = r
	}

	p0 := 2
	p1 := 0
	p2 := 1

	r = math.Sqrt(1.0 + (*M)[p0][p0] - (*M)[p1][p1] - (*M)[p2][p2])

	if r < 1e-6 {
		q[0] = 1.0
		q[1] = 0
		q[2] = 0
		q[3] = 0
		return q
	}

	q[0] = r / 2.0
	q[1] = ((*M)[p0][p1] - (*M)[p1][p0]) / (2.0 * r)
	q[2] = ((*M)[p2][p0] - (*M)[p0][p2]) / (2.0 * r)
	q[3] = ((*M)[p2][p1] - (*M)[p1][p2]) / (2.0 * r)
	return q
}
