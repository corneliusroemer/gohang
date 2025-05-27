# MRE showing Go uiprogress hang

Go terminal progress bar package [`github.com/gosuri/uiprogress`](https://github.com/gosuri/uiprogress) hangs on exit
when the indirect dependency [`github.com/gosuri/uilive`](https://github.com/gosuri/uilive) version 0.0.4 is used
together with Go version 1.21 or higher on ARM macs running macOS >= 13.

This was discovered in the NCBI datasets CLI that uses uiprogress: https://github.com/ncbi/datasets/issues/490

Repro can be executed as simply as:

```sh
go run github.com/corneliusroemer/gohang@main
```
