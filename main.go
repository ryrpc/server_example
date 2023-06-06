package main

import (
	"fmt"

	rysrv "github.com/ryrpc/server"
	"github.com/valyala/fasthttp"
)

type ip_t struct {
	Country string `json:"country"`
	Region  string `json:"region"`
	City    string `json:"city"`
}

func main() {

	repo := rysrv.NewRepository()

	repo.Register("/ip/location/one", func(ctx *fasthttp.RequestCtx) {

		if !ctx.PostArgs().Has("params") {
			err := fmt.Errorf("IpToAddress StringBytes (params) = $s\r\n", string(ctx.PostArgs().Peek("params")))
			rysrv.SetError(ctx, err)
			return
		}

		ipStr := string(ctx.PostArgs().Peek("params"))
		fmt.Println("ipStr = ", ipStr)

		resp := ip_t{
			Country: "中国",
			Region:  "beijing",
			City:    "超一米五",
		}

		rysrv.SetResult(ctx, resp)
	})

	_ = fasthttp.ListenAndServe(":8080", repo.RequestHandler())
}
