package main

import (
        rysrv "github.com/ryrpc/server"
        "github.com/valyala/fasthttp"
)

func main() {

        repo := rysrv.NewRepository()

        repo.Register("sum", func(ctx *rysrv.RequestCtx) {
                params := ctx.Params()

                a := params.GetInt("a")
                b := params.GetInt("b")

                ctx.SetResult(ctx.Arena().NewNumberInt(a + b))
        })
        repo.Register("sum_struct", func(ctx *rysrv.RequestCtx) {
                type (
                        sumRequest struct {
                                A int `json:"a"`
                                B int `json:"b"`
                        }
                        sumResponse int
                )

                var req sumRequest
                if err := ctx.ParamsUnmarshal(&req); err != nil {
                        ctx.SetError(err)
                        return
                }

                ctx.SetResult(sumResponse(req.A + req.B))
        })

        _ = fasthttp.ListenAndServe(":8080", repo.RequestHandler())
}
