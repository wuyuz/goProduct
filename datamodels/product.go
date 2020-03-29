package datamodels

type Product struct {
	ID            int64  `json:"id" sql:"id" imooc:"id"` // 可以用来映射表字段
	ProductName   string `json:"ProductName" sql:"productName" imooc:"ProductName"`
	ProductNum    int64  `json:"ProductNum" sql:"productNum" imooc:"ProductNum"`
	ProductImage string `json:"ProductImage" sql:"productImage" imooc:"ProductImage"`
	ProductUrl    string `json:"ProductUrl" sql:"productUrl" imooc:"ProductImage"`
}
