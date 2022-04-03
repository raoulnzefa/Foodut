package dto

type Product struct {
	ProductId          int      `form:"productId" json:"productId"`
	StoreName          string   `form:"storeName" json:"storeName"`
	ProductName        string   `form:"productName" json:"productName"`
	ProductCategory    string   `form:"productCategory" json:"productCategory"`
	ProductStock       int      `form:"productStock" json:"productStock"`
	ProductPrice       float64  `form:"productPrice" json:"productPrice"`
	ProductRate        float32  `form:"productRate" json:"productRate"`
	ProductDescription string   `form:"productDescription" json:"productDescription"`
	ProductPicture     []string `form:"productPicture" json:"productPicture"`
}

type PostProduct struct {
	SellerId           int      `form:"sellerId" json:"sellerId"`
	ProductName        string   `form:"productName" json:"productName"`
	ProductCategory    int      `form:"productCategory" json:"productCategory"`
	ProductStock       int      `form:"productStock" json:"productStock"`
	ProductPrice       float64  `form:"productPrice" json:"productPrice"`
	ProductDescription string   `form:"productDescription" json:"productDescription"`
	ProductPicture     []string `form:"productPicture" json:"productPicture"`
}

type EditProduct struct {
	ProductName        string  `form:"productName" json:"productName"`
	ProductPrice       float64 `form:"productPrice" json:"productPrice"`
	ProductStock       int     `form:"productStock" json:"productStock"`
	ProductCategory    int     `form:"productCategory" json:"productCategory"`
	ProductDescription string  `form:"productDescription" json:"productDescription"`
}
