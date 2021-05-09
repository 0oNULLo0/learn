package main

import (
    "net/http"

    xContext "golang.org/x/net/context"
    "golang.org/x/sync/errgroup"
)

func main() { // {{{
    ctx, cancel := xContext.WithCancel(xContext.Background())
    group, errCtx := errgroup.WithContext(ctx)

    group.Go(func() error {
        http.ListenAndServe("0.0.0.0:8080", nil)
        return nil 
    })  

    group.Go(func() error {
        http.ListenAndServe("0.0.0.0:8081", nil)
        return nil 
    })  

    // 对不起老师， 我下面的不会写了。。。应该是用无缓冲区chan来完成

} // }}}
