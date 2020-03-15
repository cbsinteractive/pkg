##  Common Go packages

This repository contains stand-alone common go libraries useful for import into our projects written in Go.

### Adding a new package

When adding a new package, create a new directory with the name of your package at the root of this repository.
Add a `go.mod` file in this new directory and import any libraries your package needs.

__NOTE: Reusable packages have few dependencies. Be critical of each external dependency you add. Sometimes a little duplication should be preferred over incurring unneeded dependency bloat.__

Good packages:

- Minimize external decencies and are as lean as they can practically be
- Look to solve a singular need, if you find your package does many things, perhaps they should be different packages
- Are well tested and include comments that would lead to a clear, usable GoDoc

Reference:  [The Zen of Go](https://dave.cheney.net/2020/02/23/the-zen-of-go)
