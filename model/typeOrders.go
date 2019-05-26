package model

import (
	"time"
)

type OrderBalance struct {
	Id                   int
	OrderId              string    `gorm:"column:OrderId" order:"orderid"`
	SupplierId           string    `gorm:"column:SupplierId"`
	FullName             string    `gorm:"column:FullName" order:"buyername"`
	Country              string    `gorm:"column:Country" order:"country"`
	Address1             string    `gorm:"column:Address1" order:"address1"`
	Address2             string    `gorm:"column:Address2" order:"address2"`
	City                 string    `gorm:"column:City" order:"city"`
	State                string    `gorm:"column:State" order:"state"`
	Postcode             string    `gorm:"column:Postcode" order:"postcode"`
	Email                string    `gorm:"column:Email" order:"email"`
	Company              string    `gorm:"column:Company" order:"company"`
	BuyerPhoneNumber     string    `gorm:"column:BuyerPhoneNumber" order:"phonenumber"`
	Shipping             string    `gorm:"column:Shipping" order:"shippingservice"`
	OrderDate            time.Time `gorm:"column:OrderDate"`
	ReplyDate            time.Time `gorm:"column:ReplyDate"`
	SalesSku             string    `gorm:"column:Sku" order:"ordersku"`
	TrackingNumber       string    `gorm:"column:TrackingNumber"`
	Qty                  int       `gorm:"column:Qty" order:"orderqty"`
	BuyerUserName        string    `gorm:"column:BuyerUserName" order:"buyerusername"`
	SalesPrice           float64   `gorm:"column:SalesPrice" order:"productprice"`
	ShippingPay          float64   `gorm:"column:ShippingPay" order:"shippingcost"`
	CostPrice            float64   `gorm:"column:CostPrice"`
	EbayFees             float64   `gorm:"column:EbayFees" order:"ebayfees"`
	PaypalFees           float64   `gorm:"column:PaypalFees"  order:"paypalfees"`
	ShippingCost         float64   `gorm:"column:ShippingCost"`
	TotalCost            float64   `gorm:"column:TotalCost"`
	Profit               float64
	EbayFeeRate          float64   `gorm:"column:EbayFeeRate"`
	PaypalFeeRate        float64   `gorm:"column:PaypalFeeRate"`
	ProfitRate           float64   `gorm:"column:ProfitRate"`
	CostRMB              float64   `gorm:"column:CostRMB"`
	CostExchangeRate     float64   `gorm:"column:CostExchangeRate"`
	PaymentRMB           float64   `gorm:"column:PaymentRMB"`
	PaymentExchangeRate  float64   `gorm:"column:PaymentExchangeRate"`
	ProfitRMB            float64   `gorm:"column:ProfitRMB"`
	GrossProfitRate      float64   `gorm:"column:GrossProfitRate"`
	RateOfExchangeLoss   float64   `gorm:"column:RateOfExchangeLoss"`
	TotalCostRate        float64   `gorm:"column:TotalCostRate"`
	ShippingCostBySystem float64   `gorm:"column:ShippingCostBySystem"`
	PaymentDate          time.Time `gorm:"column:PaymentDate" order:"paymentdate"`
	Supplier             string    `gorm:"column:Supplier"`
	ProductSKU           string    `gorm:"column:ProductSKU" order:"ProductSKU"`
	ProductQty           int       `gorm:"column:ProductQty" order:"ProductQty"`
	ProductStockNum      int       `gorm:"column:ProductStockNum"`
	ProductStockDate     time.Time `gorm:"column:ProductStokDate"`
	OrderHandled         string    `gorm:"column:OrderHandled"`
}

func (OrderBalance) TableName() string {
	return "OrderBalance"
}

type OrderLittleBoss struct {
	ID                     int
	OrderIDLb              string    `gorm:"column:小老板订单号"`
	OrderId                string    `gorm:"column:平台订单号" order:"orderid"`
	StatusLb               string    `gorm:"column:小老板订单状态"`
	StatusLg               string    `gorm:"column:物流状态"`
	StatusPlatform         string    `gorm:"column:平台状态"`
	NotePayment            string    `gorm:"column:付款备注"`
	Note                   string    `gorm:"column:买家留言"`
	DIY                    string    `gorm:"column:自定义标签"`
	NoteOrder              string    `gorm:"column:订单备注"`
	ProductAtt             string    `gorm:"column:订单商品属性"`
	Platform               string    `gorm:"column:平台"`
	Station                string    `gorm:"column:站点"`
	SellerAccount          string    `gorm:"column:卖家账号"`
	BuyerAccount           string    `gorm:"column:买家账号"  order:"buyerusername"`
	EbayOrderNum           string    `gorm:"column:ebay交易号"`
	Currency               string    `gorm:"column:货币"`
	Price                  float64   `gorm:"column:单价"`
	OrderAmount            float64   `gorm:"column:订单金额"`
	ProductTotalPrice      float64   `gorm:"column:产品总价格" order:"productprice"`
	ShippingCost           float64   `gorm:"column:运费"  order:"shippingcost"`
	Discount               float64   `gorm:"column:折扣"`
	ProductPrice           float64   `gorm:"column:订单商品成本"`
	PlatformFees           float64   `gorm:"column:佣金"  order:"ebayfees"`
	PaypalFees             float64   `gorm:"column:paypal手续费(ebay)" order:"paypalfees"`
	OrderDate              time.Time `gorm:"column:下单日期"`
	PrintDate              time.Time `gorm:"column:打单时间"`
	PaymentDate            time.Time `gorm:"column:订单付款时间" order:"paymentdate"`
	ShippingDate           time.Time `gorm:"column:订单发货时间"`
	ConsigneeName          string    `gorm:"column:收货人姓名" order:"buyername"`
	ConsigneeCountry       string    `gorm:"column:收件人国家" order:"country"`
	ConsigneePostCode      string    `gorm:"column:收货人邮编" order:"postcode"`
	ConsigneePhone         string    `gorm:"column:收货人电话" order:"phonenumber"`
	ConsigneeMobile        string    `gorm:"column:收货人手机"`
	ConsigneeEmail         string    `gorm:"column:收货人Email" order:"email"`
	ConsigneeCompany       string    `gorm:"column:收货人公司" order:"company"`
	ConsigneeCountryCode   string    `gorm:"column:收货人国家代码"`
	ConsigneeCountryCn     string    `gorm:"column:收货人国家中文"`
	ConsigneeCity          string    `gorm:"column:收货人城市" order:"city"`
	ConsigneeState         string    `gorm:"column:收货人省" order:"state"`
	ConsigneeDistrit       string    `gorm:"column:收货人区"`
	ConsigneeTown          string    `gorm:"column:收货人镇"`
	ConsigneeAddress1      string    `gorm:"column:收货人地址1" order:"address1"`
	ConsigneeAddress2      string    `gorm:"column:收货人地址2" order:"address2"`
	ConsigneeAddress3      string    `gorm:"column:收货人地址3"`
	ConsigneeDetailAddress string    `gorm:"column:详细地址"`
	Warehouse              string    `gorm:"column:仓库"`
	SelectedLogistics      string    `gorm:"column:客选物流" order:"shippingservice"`
	DefaultlLogisticalCode string    `gorm:"column:默认物流商代码"`
	PostService            string    `gorm:"column:运输服务"`
	TrackingNumber         string    `gorm:"column:物流跟踪号"`
	ListingSKu             string    `gorm:"column:店铺SKU" order:"ordersku"`
	LocalSku               string    `gorm:"column:本地商品sku"`
	Quantity               int       `gorm:"column:数量" order:"orderqty"`
	CnPickName             string    `gorm:"column:中文配货名称"`
	EnPickName             string    `gorm:"column:英文配货名称"`
	ProductName            string    `gorm:"column:商品名称"`
	Title                  string    `gorm:"column:商品标题"`
	GrossWeight            float64   `gorm:"column:称重重量"`
	ProductWeight          float64   `gorm:"column:商品重量"`
	ProductCustomsCn       string    `gorm:"column:商品报关中文名"`
	ProductCustomsEn       string    `gorm:"column:商品报关英文名"`
	ProductValue           float64   `gorm:"column:商品报关价值"`
	ProductCurrency        string    `gorm:"column:商品报关货币"`
	ProductPhoto           string    `gorm:"column:商品主图"`
	ProductPhotoURL        string    `gorm:"column:商品主图URL"`
	PlatformItemId         string    `gorm:"column:平台商品itemId"`
	MultiName              string    `gorm:"column:多品名"`
	OrderHandled           string    `gorm:"column:OrderHandled"`
}

func (OrderLittleBoss) TableName() string {
	return "OrderLittleBoss"
}

type OrderIsale struct {
	Id                int
	OrderReference    string    `gorm:"column:OrderReference"`
	PaymentDate       time.Time `gorm:"column:OrderDate"`
	ConsigneeName     string    `gorm:"column:FullName" order:"buyername"`
	ConsigneeCountry  string    `gorm:"column:Country" order:"country"`
	ConsigneeAddress1 string    `gorm:"column:Address1" order:"address1"`
	ConsigneeAddress2 string    `gorm:"column:Address2" order:"address2"`
	ConsigneeCity     string    `gorm:"column:City" order:"city"`
	ConsigneeState    string    `gorm:"column:State" order:"state"`
	ConsigneePostCode string    `gorm:"column:Postcode" order:"postcode"`
	ConsigneeEmail    string    `gorm:"column:Email" order:"email"`
	ConsigneeCompany  string    `gorm:"column:Company" order:"company"`
	ConsigneePhone    string    `gorm:"column:BuyerPhoneNumber" order:"phonenumber"`
	SelectedLogistics string    `gorm:"column:PostageServiceTag" order:"shippingservice"`
	PackagingGroupTag string    `gorm:"column:PackagingGroupTag"`
	Source            string    `gorm:"column:Source"`
	ListingSKu        string    `gorm:"column:OrderItemNumber" order:"ProductSKU"`
	Title             string    `gorm:"column:OrderItemTitle"`
	Quantity          int       `gorm:"column:OrderItemQty" order:"ProductQty"`
	Unitcost          float64   `gorm:"column:Unitcost"`
	OrderId           string    `gorm:"column:OrderId" order:"orderid"`
}

func (OrderIsale) TableName() string {
	return "OrderIsale"
}
