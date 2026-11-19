# Composite Design Pattern in Go

This example models window parts with the Composite pattern. `Item` is the
common interface implemented by the individual `Glass` and `Frame` components
as well as the composite `Window`. A window stores a slice of `Item`, so it
can contain individual parts or other windows. Its `Price` and `Display`
methods recursively process the contained items.

Run the demonstration:

```sh
go run .
```

Run the tests:

```sh
go test ./...
```
