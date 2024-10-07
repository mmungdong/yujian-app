package main

import (
	_ "yujian-app/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"yujian-app/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
