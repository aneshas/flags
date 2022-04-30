# flags

Configuration package inspired by this talk/[article](https://peter.bourgon.org/go-for-industrial-programming/) by Peter Bourgon.

The reasoning behind the package is that flags are the best way to configure your program and thus provides a thin wrapper
around go standard flag package while providig an extra degree of configurability via different extendable/composable mechanisms such as
`env variables`, `config files` etc... on an opt-in basis.
