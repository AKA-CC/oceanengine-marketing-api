package report

import "github.com/AKA-CC/oceanengine-marketing-api/pkg/enum"

type BaseResponseData struct {
	StatDatetime         string                 `json:"stat_datetime,omitempty"`          // 数据起始时间，分组条件包含 STAT_GROUP_BY_FIELD_STAT_TIME 时返回，格式为：yyyy-MM-dd HH:mm:ss
	Inventory            enum.StatInventoryType `json:"inventory,omitempty"`              // 投放广告位，分组条件包含STAT_GROUP_BY_INVENTORY时返回
	CreativeMaterialMode string                 `json:"creative_material_mode,omitempty"` // 创意类型，分组条件包含STAT_GROUP_BY_CREATIVE_MATERIAL_MODE时返回，允许值：STATIC_ASSEMBLE（程序化创意）、INTERVENE（自定义创意）
	LandingType          enum.LandingType       `json:"landing_type,omitempty"`           // 推广目的类型，分组条件包含STAT_GROUP_BY_LANDING_TYPE时返回
	Pricing              enum.AdPricing         `json:"pricing,omitempty"`                // 出价类型，分组条件包含STAT_GROUP_BY_PRICING时返回
	ImageMode            enum.ImageMode         `json:"image_mode,omitempty"`             // 素材类型，分组条件STAT_GROUP_BY_IMAGE_MODE返回
	ProvinceName         string                 `json:"province_name,omitempty"`          // 省份。如果分组条件中包括 STAT_GROUP_BY_PROVINCE_NAME 时返回
	CityName             string                 `json:"city_name,omitempty"`              // 城市。如果分组条件中包括 STAT_GROUP_BY_CITY_NAME 时返回。
	Gender               string                 `json:"gender,omitempty"`                 // 性别。如果分组条件中包括 STAT_GROUP_BY_GENDER 时返回。
	Age                  string                 `json:"age,omitempty"`                    // 年龄。如果分组条件中包括 STAT_GROUP_BY_AGE
	Platform             string                 `json:"platform,omitempty"`               // 平台。如果分组条件中包括 STAT_GROUP_BY_PLATFORM 时返回。
	Ac                   string                 `json:"ac,omitempty"`                     // 网络类型。如果分组条件中包括 STAT_GROUP_BY_AC 时返回。
	TargetField
}
