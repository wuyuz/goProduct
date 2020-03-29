package controllers

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"strconv"
	"test-produce/common"
	"test-produce/datamodels"
	"test-produce/services"
)

type ProductController struct {
	Ctx iris.Context
	ProductService services.IProductService
}


// 获取所有的数据
func (p *ProductController) GetAll() mvc.View {
	fmt.Println("触发all视图")
	productArray, _ := p.ProductService.GetAllProduct()
	return mvc.View {
		Name:"product/view.html",
		Data: iris.Map{
			"productArray": productArray,
		},
	}
}

// 修改商品
func (p *ProductController) PostUpdate() {
	product := &datamodels.Product{}
	p.Ctx.Request().ParseForm()
	// 使用form.go以及结构体的tag来映射到具体的结构体中  如imooc下的所有字段
	des :=common.NewDecoder(&common.DecoderOptions{TagName:"imooc"})

	if err := des.Decode(p.Ctx.Request().Form,product); err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	err := p.ProductService.UpdateProduct(product)
	if err != nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

// 添加页面
func (p *ProductController) GetAdd() mvc.View {
	return mvc.View{
		Name:"product/add.html",
	}
}

// 新增函数
func (p *ProductController) PostAdd() {
	product :=&datamodels.Product{}
	p.Ctx.Request().ParseForm()
	dec := common.NewDecoder(&common.DecoderOptions{TagName:"imooc"})
	if err:= dec.Decode(p.Ctx.Request().Form,product);err!=nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	_,err:=p.ProductService.InsertProduct(product)
	if err !=nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	p.Ctx.Redirect("/product/all")
}

// 单条数据搜索
func (p *ProductController) GetManager() mvc.View {
	idString := p.Ctx.URLParam("id")
	id,err :=strconv.ParseInt(idString,10,16)
	if err !=nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	product,err:=p.ProductService.GetProductByID(id)
	if err !=nil {
		p.Ctx.Application().Logger().Debug(err)
	}

	return mvc.View{
		Name:"product/manager.html",
		Data:iris.Map{
			"product":product,
		},
	}
}

// 删除函数
func (p *ProductController) GetDelete() {
	idString:=p.Ctx.URLParam("id")
	id ,err := strconv.ParseInt(idString,10,64)
	if err !=nil {
		p.Ctx.Application().Logger().Debug(err)
	}
	isOk:=p.ProductService.DeleteProductByID(id)
	if isOk{
		p.Ctx.Application().Logger().Debug("删除商品成功，ID为："+idString)
	} else {
		p.Ctx.Application().Logger().Debug("删除商品失败，ID为："+idString)
	}
	p.Ctx.Redirect("/product/all")
}

