package linmath

import "math"

type Mat4x4 [4]Vec4

func equalIsOne(a, b int) float64 {
	if a == b {
		return 1.
	}
	return 0
}

func Mat4x4Identity(M *Mat4x4) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			(*M)[i][j] = equalIsOne(i, j)
		}
	}
}

func Mat4x4Dup(M *Mat4x4, N Mat4x4) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			(*M)[i][j] = N[i][j]
		}
	}
}

func Mat4x4Row(M Mat4x4, i int) (r Vec4) {
	for k := 0; k < 4; k++ {
		r[k] = M[k][i]
	}
	return r
}

func Mat4x4Col(M Mat4x4, i int) (r Vec4) {
	for k := 0; k < 4; k++ {
		r[k] = M[i][k]
	}
	return r
}

func Mat4x4Transpose(M *Mat4x4, N Mat4x4) {
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			(*M)[i][j] = N[j][i]
		}
	}
}
func Mat4x4Add(M *Mat4x4, a, b Mat4x4) {
	for i := 0; i < 4; i++ {
		(*M)[i] = Vec4Add(a[i], b[i])
	}
}

func Mat4x4Sub(M *Mat4x4, a, b Mat4x4) {
	for i := 0; i < 4; i++ {
		(*M)[i] = Vec4Sub(a[i], b[i])
	}
}
func Mat4x4Scale(M *Mat4x4, a Mat4x4, k float64) {
	for i := 0; i < 4; i++ {
		(*M)[i] = Vec4Scale(a[i], k)
	}
}

func Mat4x4ScaleAniso(M *Mat4x4, a Mat4x4, x, y, z float64) {
	(*M)[0] = Vec4Scale(a[0], x)
	(*M)[1] = Vec4Scale(a[1], y)
	(*M)[2] = Vec4Scale(a[2], z)
	for i := 0; i < 4; i++ {
		(*M)[3][i] = a[3][i]
	}
}

func Mat4x4Mul(M *Mat4x4, a, b Mat4x4) {
	var temp Mat4x4
	for c := 0; c < 4; c++ {
		for r := 0; r < 4; r++ {
			temp[c][r] = 0
			for k := 0; k < 4; k++ {
				temp[c][r] += a[k][r] * b[c][k]
			}
		}
	}
	Mat4x4Dup(M, temp)
}

func Mat4x4MulVec4(M Mat4x4, v Vec4) (r Vec4) {
	for j := 0; j < 4; j++ {
		r[j] = 0
		for i := 0; i < 4; i++ {
			r[j] += M[i][j] * v[i]
		}
	}
	return r
}

func Mat4x4Translate(T *Mat4x4, x, y, z float64) {
	Mat4x4Identity(T)
	(*T)[3][0] = x
	(*T)[3][1] = y
	(*T)[3][2] = z
}

func Mat4x4TranslateInPlace(M *Mat4x4, x, y, z float64) {
	t := Vec4{x, y, z, 0}
	var r Vec4
	for i := 0; i < 4; i++ {
		r = Mat4x4Row(*M, i)
		(*M)[3][i] += Vec4MulInner(r, t)
	}
}

func Mat4x4FromVec3MulOuter(M *Mat4x4, a, b Vec3) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i < 3 && j < 3 {
				(*M)[i][j] = a[i] * b[j]
			} else {
				(*M)[i][j] = 0
			}
		}
	}
}

func Mat4x4Rotate(R *Mat4x4, M Mat4x4, x, y, z, angle float64) {
	s := math.Sin(angle)
	c := math.Cos(angle)
	u := Vec3{x, y, z}

	if Vec3Len(u) > 1e-4 {
		u = Vec3Norm(u)
		var T Mat4x4
		Mat4x4FromVec3MulOuter(&T, u, u)

		S := Mat4x4{
			Vec4{0, u[2], -u[1], 0},
			Vec4{-u[2], 0, u[0], 0},
			Vec4{u[1], -u[0], 0, 0},
			Vec4{0, 0, 0, 0},
		}
		Mat4x4Scale(&S, S, s)

		var C Mat4x4
		Mat4x4Identity(&C)
		Mat4x4Sub(&C, C, T)

		Mat4x4Scale(&C, C, c)

		Mat4x4Add(&T, T, C)
		Mat4x4Add(&T, T, S)

		T[3][3] = 1.0
		Mat4x4Mul(R, M, T)
	} else {
		Mat4x4Dup(R, M)
	}
}

func Mat4x4RotateX(Q *Mat4x4, M Mat4x4, angle float64) {
	s := math.Sin(angle)
	c := math.Cos(angle)
	R := Mat4x4{
		Vec4{1.0, 0, 0, 0},
		Vec4{0, c, s, 0},
		Vec4{0, -s, c, 0},
		Vec4{0, 0, 0, 1.0},
	}
	Mat4x4Mul(Q, M, R)
}

func Mat4x4RotateY(Q *Mat4x4, M Mat4x4, angle float64) {
	s := math.Sin(angle)
	c := math.Cos(angle)
	R := Mat4x4{
		Vec4{c, 0, s, 0},
		Vec4{0, 1.0, 0, 0},
		Vec4{-s, 0, c, 0},
		Vec4{0, 0, 0, 1.0},
	}
	Mat4x4Mul(Q, M, R)
}

func Mat4x4RotateZ(Q *Mat4x4, M Mat4x4, angle float64) {
	s := math.Sin(angle)
	c := math.Cos(angle)
	R := Mat4x4{
		Vec4{c, s, 0, 0},
		Vec4{-s, c, 0, 0},
		Vec4{0, 0, 1.0, 0},
		Vec4{0, 0, 0, 1.0},
	}
	Mat4x4Mul(Q, M, R)
}

func Mat4x4Invert(T *Mat4x4, M Mat4x4) {
	var (
		s [6]float64
		c [6]float64
	)

	s[0] = M[0][0]*M[1][1] - M[1][0]*M[0][1]
	s[1] = M[0][0]*M[1][2] - M[1][0]*M[0][2]
	s[2] = M[0][0]*M[1][3] - M[1][0]*M[0][3]
	s[3] = M[0][1]*M[1][2] - M[1][1]*M[0][2]
	s[4] = M[0][1]*M[1][3] - M[1][1]*M[0][3]
	s[5] = M[0][2]*M[1][3] - M[1][2]*M[0][3]

	c[0] = M[2][0]*M[3][1] - M[3][0]*M[2][1]
	c[1] = M[2][0]*M[3][2] - M[3][0]*M[2][2]
	c[2] = M[2][0]*M[3][3] - M[3][0]*M[2][3]
	c[3] = M[2][1]*M[3][2] - M[3][1]*M[2][2]
	c[4] = M[2][1]*M[3][3] - M[3][1]*M[2][3]
	c[5] = M[2][2]*M[3][3] - M[3][2]*M[2][3]

	// Assumes it is invertible
	idet := 1.0 / (s[0]*c[5] - s[1]*c[4] + s[2]*c[3] + s[3]*c[2] - s[4]*c[1] + s[5]*c[0])

	(*T)[0][0] = (M[1][1]*c[5] - M[1][2]*c[4] + M[1][3]*c[3]) * idet
	(*T)[0][1] = (-M[0][1]*c[5] + M[0][2]*c[4] - M[0][3]*c[3]) * idet
	(*T)[0][2] = (M[3][1]*s[5] - M[3][2]*s[4] + M[3][3]*s[3]) * idet
	(*T)[0][3] = (-M[2][1]*s[5] + M[2][2]*s[4] - M[2][3]*s[3]) * idet

	(*T)[1][0] = (-M[1][0]*c[5] + M[1][2]*c[2] - M[1][3]*c[1]) * idet
	(*T)[1][1] = (M[0][0]*c[5] - M[0][2]*c[2] + M[0][3]*c[1]) * idet
	(*T)[1][2] = (-M[3][0]*s[5] + M[3][2]*s[2] - M[3][3]*s[1]) * idet
	(*T)[1][3] = (M[2][0]*s[5] - M[2][2]*s[2] + M[2][3]*s[1]) * idet

	(*T)[2][0] = (M[1][0]*c[4] - M[1][1]*c[2] + M[1][3]*c[0]) * idet
	(*T)[2][1] = (-M[0][0]*c[4] + M[0][1]*c[2] - M[0][3]*c[0]) * idet
	(*T)[2][2] = (M[3][0]*s[4] - M[3][1]*s[2] + M[3][3]*s[0]) * idet
	(*T)[2][3] = (-M[2][0]*s[4] + M[2][1]*s[2] - M[2][3]*s[0]) * idet

	(*T)[3][0] = (-M[1][0]*c[3] + M[1][1]*c[1] - M[1][2]*c[0]) * idet
	(*T)[3][1] = (M[0][0]*c[3] - M[0][1]*c[1] + M[0][2]*c[0]) * idet
	(*T)[3][2] = (-M[3][0]*s[3] + M[3][1]*s[1] - M[3][2]*s[0]) * idet
	(*T)[3][3] = (M[2][0]*s[3] - M[2][1]*s[1] + M[2][2]*s[0]) * idet
}

func Mat4x4Orthonormalize(R *Mat4x4, M Mat4x4) {
	Mat4x4Dup(R, M)

	vn := Vec3Norm(Vec3FromVec4((*R)[2]))

	(*R)[2][0] = vn[0]
	(*R)[2][1] = vn[1]
	(*R)[2][2] = vn[2]

	s := Vec3MulInner(Vec3FromVec4((*R)[1]), Vec3FromVec4((*R)[2]))
	h := Vec3Scale(Vec3FromVec4((*R)[2]), s)
	vs := Vec3Sub(Vec3FromVec4((*R)[1]), h)

	(*R)[1][0] = vs[0]
	(*R)[1][1] = vs[1]
	(*R)[1][2] = vs[2]

	vn = Vec3Norm(Vec3FromVec4((*R)[1]))

	(*R)[1][0] = vn[0]
	(*R)[1][1] = vn[1]
	(*R)[1][2] = vn[2]

	s = Vec3MulInner(Vec3FromVec4((*R)[0]), Vec3FromVec4((*R)[2]))
	h = Vec3Scale(Vec3FromVec4((*R)[2]), s)
	vs = Vec3Sub(Vec3FromVec4((*R)[0]), h)

	(*R)[0][0] = vs[0]
	(*R)[0][1] = vs[1]
	(*R)[0][2] = vs[2]

	s = Vec3MulInner(Vec3FromVec4((*R)[0]), Vec3FromVec4((*R)[1]))
	h = Vec3Scale(Vec3FromVec4((*R)[1]), s)
	vs = Vec3Sub(Vec3FromVec4((*R)[0]), h)
	vn = Vec3Norm(vs)

	(*R)[0][0] = vn[0]
	(*R)[0][1] = vn[1]
	(*R)[0][2] = vn[2]
}

func Mat4x4Frustum(M *Mat4x4, l, r, b, t, n, f float64) {
	(*M)[0][0] = 2.0 * n / (r - l)
	(*M)[0][1] = 0
	(*M)[0][2] = 0
	(*M)[0][3] = 0

	(*M)[1][1] = 2. * n / (t - b)
	(*M)[1][0] = 0
	(*M)[1][2] = 0
	(*M)[1][3] = 0

	(*M)[2][0] = (r + l) / (r - l)
	(*M)[2][1] = (t + b) / (t - b)
	(*M)[2][2] = -(f + n) / (f - n)
	(*M)[2][3] = -1.0

	(*M)[3][2] = -2.0 * (f * n) / (f - n)
	(*M)[3][0] = 0
	(*M)[3][1] = 0
	(*M)[3][3] = 0
}

func Mat4x4Ortho(M *Mat4x4, l, r, b, t, n, f float64) {
	(*M)[0][0] = 2.0 / (r - l)
	(*M)[0][1] = 0
	(*M)[0][2] = 0
	(*M)[0][3] = 0

	(*M)[1][1] = 2.0 / (t - b)
	(*M)[1][0] = 0
	(*M)[1][2] = 0
	(*M)[1][3] = 0

	(*M)[2][2] = -2.0 / (f - n)
	(*M)[2][0] = 0
	(*M)[2][1] = 0
	(*M)[2][3] = 0

	(*M)[3][0] = -(r + l) / (r - l)
	(*M)[3][1] = -(t + b) / (t - b)
	(*M)[3][2] = -(f + n) / (f - n)
	(*M)[3][3] = 1.0
}

func Mat4x4Perspective(M *Mat4x4, yFov, aspect, n, f float64) {
	/* NOTE: Degrees are an unhandy unit to work with.
	 * linmath.h uses radians for everything! */
	a := 1.0 / math.Tan(yFov/2.0)

	(*M)[0][0] = a / aspect
	(*M)[0][1] = 0
	(*M)[0][2] = 0
	(*M)[0][3] = 0

	(*M)[1][0] = 0
	(*M)[1][1] = a
	(*M)[1][2] = 0
	(*M)[1][3] = 0

	(*M)[2][0] = 0
	(*M)[2][1] = 0
	(*M)[2][2] = -((f + n) / (f - n))
	(*M)[2][3] = -1.0

	(*M)[3][0] = 0
	(*M)[3][1] = 0
	(*M)[3][2] = -((2.0 * f * n) / (f - n))
	(*M)[3][3] = 0
}

func Mat4x4LookAt(M *Mat4x4, eye, center, up Vec3) {
	/* Adapted from Android's OpenGL Matrix.java.                        */
	/* See the OpenGL GLUT documentation for gluLookAt for a description */
	/* of the algorithm. We implement it in a straightforward way:       */

	/* TODO: The negation of of can be spared by swapping the order of
	 *       operands in the following cross products in the right way. */

	f := Vec3Norm(Vec3Sub(center, eye))
	s := Vec3Norm(Vec3MulCross(f, up))
	t := Vec3MulCross(s, f)

	(*M)[0][0] = s[0]
	(*M)[0][1] = t[0]
	(*M)[0][2] = -f[0]
	(*M)[0][3] = 0

	(*M)[1][0] = s[1]
	(*M)[1][1] = t[1]
	(*M)[1][2] = -f[1]
	(*M)[1][3] = 0

	(*M)[2][0] = s[2]
	(*M)[2][1] = t[2]
	(*M)[2][2] = -f[2]
	(*M)[2][3] = 0

	(*M)[3][0] = 0
	(*M)[3][1] = 0
	(*M)[3][2] = 0
	(*M)[3][3] = 1.0

	Mat4x4TranslateInPlace(M, -eye[0], -eye[1], -eye[2])
}