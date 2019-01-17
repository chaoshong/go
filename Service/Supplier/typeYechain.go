package service

type GetProduct struct {
	sku       []string
	page_size int
	page      string
}

type ProductResult struct {
	Ask         int           `json:"ask"`
	Message     string        `json:"message"`
	Total       int           `json:"total"`
	ProductList []ProductList `json:"data"`
	Page        int           `json:"page"`
	PageSize    int           `json:"page_size"`
}

type ProductList struct {
	SpID              string      `json:"sp_id"`
	ProductSku        string      `json:"product_sku"`
	ProductTitleEn    string      `json:"product_title_en"`
	ProductTitle      string      `json:"product_title"`
	SharedStatus      string      `json:"shared_status"`
	ProductWeight     string      `json:"product_weight"`
	ProductAllWeight  string      `json:"product_all_weight"`
	ProductLength     string      `json:"product_length"`
	ProductWidth      string      `json:"product_width"`
	ProductHeight     string      `json:"product_height"`
	ProductPickHeight string      `json:"product_pick_height"`
	ProductPickLength string      `json:"product_pick_length"`
	ProductPickWidth  string      `json:"product_pick_width"`
	IsShowInventory   string      `json:"is_show_inventory"`
	ProductImages     string      `json:"product_images"`
	ImagesArr         []string    `json:"images_arr"`
	IsTop             string      `json:"is_top"`
	IsNew             string      `json:"is_new"`
	IsSale            string      `json:"is_sale"`
	IsCombination     string      `json:"is_combination"`
	IsPacking         string      `json:"is_packing"`
	ProductAddedTime  string      `json:"product_added_time"`
	ProductKeyword    string      `json:"product_keyword"`
	IsFocus           string      `json:"is_focus"`
	IsProductLibrary  string      `json:"is_product_library"`
	SpaSeries         string      `json:"spa_series"`
	Warehouse         []Warehouse `json:"warehouse"`
}

type Warehouse struct {
	CurrencyCode              string `json:"currency_code"`
	AssignLogisticsCode       string `json:"assign_logistics_code"`
	LogisticsIsCoerce         string `json:"logistics_is_coerce"`
	ShippingType              string `json:"shipping_type"`
	ShippingValue             string `json:"shipping_value"`
	WarehouseCode             string `json:"warehouse_code"`
	WarehouseName             string `json:"warehouse_name"`
	ProductInventory          int    `json:"product_inventory"`
	ProductExclusiveInventory int    `json:"product_exclusive_inventory"`
	UnitPrice                 string `json:"unit_price"`
	SuggestSalePrice          string `json:"suggest_sale_price"`
}
