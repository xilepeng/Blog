package main

import (
	"github.com/xilepeng/Blog/model"
	"github.com/xilepeng/Blog/routers"
)

func main() {
	model.InitDb()
	routers.InitRouter()

}
