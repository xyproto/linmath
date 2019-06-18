package linmath

import "math"

type Quat [4]float64

func Quatidentity() (r Quat) {
	r[0] = 0
	r[1] = 0
	r[2] = 0
	r[3] = 1.0
	return r
}

func QuatAdd(a, b Quat) (r Quat) {
	for i := 0; i < 4; i++ {
		r[i] = a[i] + b[i]
	}
	return r
}

func QuatSub(a, b Quat) (r Quat) {
	for i := 0; i < 4; i++ {
		r[i] = a[i] - b[i]
	}
	return r
}

func Vec3FromQuat(q Quat) (r Vec3) {
	r[0] = q[0]
	r[1] = q[1]
	r[2] = q[2]
	return r
}

func QuatMul(p, q Quat) (r Quat) {
	v := Vec3MulCross(Vec3FromQuat(p), Vec3FromQuat(q))
	w := Vec3Scale(Vec3FromQuat(p), q[3])
	v = Vec3Add(v, w)
	w = Vec3Scale(Vec3FromQuat(q), p[3])
	v = Vec3Add(v, w)
	r[0] = v[0]
	r[1] = v[1]
	r[2] = v[2]
	r[3] = p[3]*q[3] - Vec3MulInner(Vec3FromQuat(p), Vec3FromQuat(q))
	return r
}

func QuatScale(v Quat, s float64) (r Quat) {
	for i := 0; i < 4; i++ {
		r[i] = v[i] * s
	}
	return r
}

func QuatInnerProduct(a, b Quat) (p float64) {
	for i := 0; i < 4; i++ {
		p += b[i] * a[i]
	}
	return p
}

func QuatConj(q Quat) (r Quat) {
	for i := 0; i < 3; i++ {
		r[i] = -q[i]
	}
	r[3] = q[3]
	return r
}

func QuatRotate(angle float64, axis Vec3) (r Quat) {
	v := Vec3Scale(axis, math.Sin(angle/2))
	for i := 0; i < 3; i++ {
		r[i] = v[i]
	}
	r[3] = math.Cos(angle / 2)
	return r
}

func QuatFromVec4(v Vec4) (q Quat) {
	q[0] = v[0]
	q[1] = v[1]
	q[2] = v[2]
	q[3] = v[3]
	return q
}

func Vec4FromQuat(q Quat) (v Vec4) {
	v[0] = q[0]
	v[1] = q[1]
	v[2] = q[2]
	v[3] = q[3]
	return v
}

func QuatNorm(q Quat) Quat {
	return QuatFromVec4(Vec4Norm(Vec4FromQuat(q)))
}

func QuatMulVec3(q Quat, v Vec3) (r Vec3) {
	/*
	    * Method by Fabian 'ryg' Giessen (of Farbrausch)
	   t = 2 * cross(q.xyz, v)
	   v' = v + q.w * t + cross(q.xyz, t)
	*/
	qXyz := Vec3{q[0], q[1], q[2]}
	u := Vec3{q[0], q[1], q[2]}

	t := Vec3MulCross(qXyz, v)
	t = Vec3Scale(t, 2)

	u = Vec3MulCross(qXyz, t)
	t = Vec3Scale(t, q[3])

	r = Vec3Add(v, t)
	r = Vec3Add(r, u)
	return r
}

func Mat4x4FromQuat(M *Mat4x4, q Quat) {
	a := q[3]
	b := q[0]
	c := q[1]
	d := q[2]

	a2 := a * a
	b2 := b * b
	c2 := c * c
	d2 := d * d

	(*M)[0][0] = a2 + b2 - c2 - d2
	(*M)[0][1] = 2.0 * (b*c + a*d)
	(*M)[0][2] = 2.0 * (b*d - a*c)
	(*M)[0][3] = 0

	(*M)[1][0] = 2 * (b*c - a*d)
	(*M)[1][1] = a2 - b2 + c2 - d2
	(*M)[1][2] = 2.0 * (c*d + a*b)
	(*M)[1][3] = 0

	(*M)[2][0] = 2.0 * (b*d + a*c)
	(*M)[2][1] = 2.0 * (c*d - a*b)
	(*M)[2][2] = a2 - b2 - c2 + d2
	(*M)[2][3] = 0

	(*M)[3][0] = 0
	(*M)[3][1] = 0
	(*M)[3][2] = 0
	(*M)[3][3] = 1.0
}

func Mat4x4oMulQuat(R *Mat4x4, M Mat4x4, q Quat) {
	/*  XXX: The way this is written only works for othogonal matrices. */
	/* TODO: Take care of non-orthogonal case. */
	(*R)[0] = Vec43(QuatMulVec3(q, Vec34(M[0])))
	(*R)[1] = Vec43(QuatMulVec3(q, Vec34(M[1])))
	(*R)[2] = Vec43(QuatMulVec3(q, Vec34(M[2])))
	(*R)[3][0] = 0
	(*R)[3][1] = 0
	(*R)[3][2] = 0
	(*R)[3][3] = 1.0
}

func (M Mat4x4) Quat() (q Quat) {
	r := 0.0

	for i := 0; i < 3; i++ {
		m := M[i][i]
		if m < r {
			continue
		}
		m = r
	}

	p0 := 2
	p1 := 0
	p2 := 1

	r = math.Sqrt(1.0 + M[p0][p0] - M[p1][p1] - M[p2][p2])

	if r < 1e-6 {
		q[0] = 1.0
		q[1] = 0
		q[2] = 0
		q[3] = 0
		return q
	}

	q[0] = r / 2.0
	q[1] = (M[p0][p1] - M[p1][p0]) / (2.0 * r)
	q[2] = (M[p2][p0] - M[p0][p2]) / (2.0 * r)
	q[3] = (M[p2][p1] - M[p1][p2]) / (2.0 * r)
	return q
}
