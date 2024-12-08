package req

type ProductReq struct {
	Id         int32    `json:"id"`
	Name       string   `json:"name" binding:"required, min=2, max=32"`
	SN         string   `json:"sn" binding:"required,min=2,max=16"`
	Stocks     int32    `json:"stocks" binding:"required, min=1"`
	CategoryId int32    `json:"category_id" binding:"required"`
	Price      float32  `json:"price" binding:"required, min=0"`
	RealPrice  float32  `json:"real_price" binding:"required, min=0"`
	ShortDesc  string   `json:"short_desc" binding:"required, min=3"`
	Desc       string   `json:"desc" binding:"required, min=3"`
	Images     []string `json:"images" binding:"required, min=1"`
	DescImages []string `json:"desc_images" binding:"required, min=1"`
	CoverImage string   `json:"cover_image" binding:"required, url"`
	BrandId    int32    `json:"brand_id" binding:"required"`
	IsNew      bool     `json:"is_new" binding:"required"`
	IsPop      bool     `json:"is_pop" binding:"required"`
	Selling    bool     `json:"selling" binding:"required"`
	FavNum     int32    `json:"fav_num" binding:"required"`
	SoldNum    int32    `json:"sold_num" binding:"required"`
}
