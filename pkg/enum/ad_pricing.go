package enum

type AdPricing string

const (
	PRICING_CPC  AdPricing = "PRICING_CPC"  // CPC（点击付费），出价范围（单位元）:0.2-100
	PRICING_CPM  AdPricing = "PRICING_CPM"  // CPM（展示付费），出价范围（单位元）: 4-100
	PRICING_OCPC AdPricing = "PRICING_OCPC" // OCPC（已下线，仅投放范围为穿山甲可用）
	PRICING_OCPM AdPricing = "PRICING_OCPM" // OCPM（转化量付费），出价范围（单位元）:0.1-10000
	PRICING_CPV  AdPricing = "PRICING_CPV"  // CPV ，出价范围（单位元）:0.07-100
	PRICING_CPA  AdPricing = "PRICING_CPA"  // CPA（已下线，不支持进行出价方式为CPA的投放行为
)
