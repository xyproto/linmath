# Linear Math

[![Build Status](https://travis-ci.org/xyproto/lm.svg?branch=master)](https://travis-ci.org/xyproto/lm) [![GoDoc](https://godoc.org/github.com/xyproto/lm?status.svg)](https://godoc.org/github.com/xyproto/lm) [![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/lm/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/lm)](https://goreportcard.com/report/github.com/xyproto/lm)

Port of [linmath.h](https://github.com/datenwolf/linmath.h) by Wolfgang 'datenwolf' Draxinger &lt;code@datenwolf.net&gt; (licensed under WTFPL), from C to Go.

## Features and limitations

Several functions and types are provided:

* `Vec2`, a 2 element vector of float64 (x,y)
* `Vec3`, a 3 element vector of float64 (x,y,z)
* `Vec4`, a 4 element vector of float64 (4th component used for homogenous computations)
* `Mat4x4`, a 4 by 4 elements matrix (computations are done in column major order)
* `Quat`, a 4 element vector of float64 (x,y,z,w)

These may be useful for computer graphics programming.

## TODO

- [ ] Full test coverage
- [ ] Full benchmark coverage
- [ ] Optimize functions

## General info

* License: MIT
* Version: 1.0.0
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
