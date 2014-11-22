automatic or algorithmic differentiation in go
----------------------------------------------

a simple mathematical formula language, which is auto-differentiated
and compiled to http://golang.org for high performance.

see https://autodiff.info for live demo.

to get started: make sure you have latest golang.org installed
(e.g. https://golang.org/dl/), or build it yourself via
https://github.com/xoba/goinit

then:

    git clone --recursive https://github.com/xoba/ad.git
    cd ad
    source goinit.sh
    ./install.sh
    run compile -formula "f := sqrt(abs(a+b*b))"
    go run compute.go

for help, you can try:

    run
    run compile -help
    run nn -help

to develop with emacs:

    ./ide.sh

to auto-generate various code:

    lib/gogenerate.sh

to run a simple neural network example:

    run nn

