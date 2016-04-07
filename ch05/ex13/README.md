# Exercise 5.13
Modify `crawl` to make local copies of the pages it finds, creating directories as necessary. Don't make copies of pages that come from a different domain. For example, if the original pages comes from `golang.org`, save all files from there, but exclude ones from `vimeo.com`.

---
# 練習問題 5.13
`crawl`を修正して、必要に応じてディレクトリを作成しながら見つけたページの複製をローカルに作成するようにしなさい。異なるドメインのページの複製はしないようにしなさい。たとえば、もとのページが`golang.org`からであれば、そこにあるすべてのファイルは保存しますが、`vimeo.com`からのファイルは保存しないということです。


# Result

````
$  go run crawl.go links.go http://golang.org
$  tree -L 3 local
local
└── golang.org
    ├── CONTRIBUTORS
    ├── LICENSE
    ├── blog
    │   ├── c-go-cgo
    │   ├── defer-panic-and-recover
    │   ├── error-handling-and-go
    │   ├── gif-decoder-exercise-in-go-interfaces
    │   ├── go-concurrency-patterns-timing-out-and
    │   ├── go-image-package
    │   ├── go-imagedraw-package
    │   ├── go-slices-usage-and-internals
    │   ├── gobs-of-data
    │   ├── godoc-documenting-go-code
    │   ├── gos-declaration-syntax
    │   ├── json-and-go
    │   ├── json-rpc-tale-of-interfaces
    │   ├── laws-of-reflection
    │   ├── organizing-go-code
    │   ├── profiling-go-programs
    │   └── race-detector
    ├── change
    ├── cmd
    │   ├── cgo
    │   ├── go
    │   └── godoc
    ├── conduct
    ├── doc
    │   ├── articles
    │   ├── asm
    │   ├── cmd
    │   ├── code.html
    │   ├── codewalk
    │   ├── contribute.html
    │   ├── devel
    │   ├── effective_go.html
    │   ├── faq
    │   ├── gdb
    │   ├── go1
    │   ├── go1.1
    │   ├── go1.2
    │   ├── go1.3
    │   ├── go1.4
    │   ├── go1.5
    │   ├── go1.6
    │   ├── go1compat
    │   ├── go1compat.html
    │   ├── go_faq.html
    │   ├── install
    │   ├── install.html
    │   └── tos.html
    ├── issue
    ├── lib
    │   └── godoc
    ├── pkg
    │   ├── go
    │   ├── regexp
    │   ├── runtime
    │   └── sync
    ├── play
    ├── ref
    │   ├── mem
    │   └── spec
    ├── s
    │   ├── comments
    │   ├── go13compiler
    │   ├── go14customimport
    │   ├── go14gc
    │   ├── go14internal
    │   ├── go15bootstrap
    │   ├── go15vendor
    │   └── releasesched
    ├── security
    ├── src
    │   └── fmt
    ├── talks
    ├── test
    │   └── bench
    ├── wiki
    └── x
        └── text

18 directories, 63 files
````
