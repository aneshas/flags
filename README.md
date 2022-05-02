# Go flags

[![Go](https://github.com/aneshas/flags/actions/workflows/go.yml/badge.svg)](https://github.com/aneshas/flags/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/aneshas/flags/badge.svg?branch=trunk)](https://coveralls.io/github/aneshas/flags?branch=trunk)
[![Go Report Card](https://goreportcard.com/badge/github.com/aneshas/flags)](https://goreportcard.com/report/github.com/aneshas/flags)

Configuration package inspired by this talk/[article](https://peter.bourgon.org/go-for-industrial-programming/) by Peter Bourgon.

The guiding idea behind this package is that `flags are the best way to configure your program` and thus it provides a thin wrapper
around go standard flag package while providig an extra degree of configurability via different extendable/composable mechanisms such as
`env variables`, `config files` etc... on an opt-in basis.
