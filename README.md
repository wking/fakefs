Status: DEPRECATED

I'll keep this repository around because it's pretty tiny, but you can accomplish its goals using GoDoc's [mapfs][] and [httpfs][].  There's an example in the [GoDoc][].

# FakeFS

An in-memory [`net/http` `FileSystem`][FileSystem] for mock-testing [Go][] programs.
See the [GoDoc][] for more details.

## Related work

* [`github.com/arschles/sys`](https://github.com/arschles/sys) is an in-memory filesystem, but does not implement [`net/http`'s `FileSystem` interface][FileSystem].
* [`github.com/blang/vfs`](https://github.com/blang/vfs) is an in-memory filesystem, but does not implement [`net/http`'s `FileSystem` interface][FileSystem].
* [`github.com/mjibson/esc`](https://github.com/mjibson/esc) is an in-memory filesystem which implements [`net/http`'s `FileSystem` interface][FileSystem], but the types are private (you are expected to generate them with the `esc` tool).

[FileSystem]: https://golang.org/pkg/net/http/#FileSystem
[Go]: https://golang.org/
[GoDoc]: https://godoc.org/github.com/wking/fakefs
[httpfs]: https://godoc.org/golang.org/x/tools/godoc/vfs/httpfs
[mapfs]: https://godoc.org/golang.org/x/tools/godoc/vfs/mapfs
