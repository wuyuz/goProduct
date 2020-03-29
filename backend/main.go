package main

import (
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"test-produce/backend/web/controllers"
	"test-produce/common"
	"test-produce/repositories"
	"test-produce/services"
)

func main() {
	// 1. 创建iris 实例
	app := iris.New()
	// 2. 设置错误模式,在mvc模式下提示错误
	app.Logger().SetLevel("debug")

	// 注册模板
	tmplate := iris.HTML("./backend/web/view", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(tmplate)

	// 设置模板目标, 屏蔽静态资源
	app.StaticWeb("/assets", "./backend/web/assets")

	// 异常页面,出现异常就跳转到此
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错"))
		ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	db, err := common.NewMysqlConn()
	if err != nil {
		log.Fatal(err)
	}

	// 上下文
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 注册控制器
	productRepository := repositories.NewProductManager("product",db)
	productService := services.NewProductService(productRepository)
	productParty := app.Party("/product")
	product := mvc.New(productParty)
	product.Register(ctx, productService)
	product.Handle(new(controllers.ProductController))


	// 启动服务
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed), // 忽略框架的错误
		iris.WithOptimizations,
	)

}
