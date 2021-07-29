package enum

type ImageMode string

const (
	CREATIVE_IMAGE_MODE_SMALL              ImageMode = "CREATIVE_IMAGE_MODE_SMALL"              // 小图，宽高比1.52，大小1.5M以下，下限：456 & 300，上限：1368 & 900
	CREATIVE_IMAGE_MODE_LARGE              ImageMode = "CREATIVE_IMAGE_MODE_LARGE"              // 大图，横版大图宽高比1.78，大小1.5M以下，下限：1280 & 720，上限：2560 & 1440
	CREATIVE_IMAGE_MODE_GROUP              ImageMode = "CREATIVE_IMAGE_MODE_GROUP"              // 组图，宽高比1.52，大小1.5M以下，下限：456 & 300，上限：1368 & 900
	CREATIVE_IMAGE_MODE_VIDEO              ImageMode = "CREATIVE_IMAGE_MODE_VIDEO"              // 横版视频，封面图宽高比1.78（下限：1280 & 720，上限：2560 & 1440））,视频资源支持mp4、mpeg、3gp、avi文件格式，宽高比16:9，大小不超过1000M
	CREATIVE_IMAGE_MODE_GIF                ImageMode = "CREATIVE_IMAGE_MODE_GIF"                // GIF图,宽高比(690, 388),大小不超过1.5M
	CREATIVE_IMAGE_MODE_LARGE_VERTICAL     ImageMode = "CREATIVE_IMAGE_MODE_LARGE_VERTICAL"     // 大图竖图，宽高比0.56，大小1.5M以下，下限：720 & 1280，上限：1440 & 2560
	CREATIVE_IMAGE_MODE_VIDEO_VERTICAL     ImageMode = "CREATIVE_IMAGE_MODE_VIDEO_VERTICAL"     // 竖版视频，封面图宽高比0.56（9:16），下限：720 & 1280，上限：1440 & 2560，视频资源支持mp4、mpeg、3gp、avi文件格式，大小不超过100M
	TOUTIAO_SEARCH_AD_IMAGE                ImageMode = "TOUTIAO_SEARCH_AD_IMAGE"                // 搜索大图; 仅限搜索广告使用，下限：345 & 138，上限：690 & 276
	SEARCH_AD_SMALL_IMAGE                  ImageMode = "SEARCH_AD_SMALL_IMAGE"                  // 搜索小图; 仅限搜索广告使用，下限：108 & 72，上限：432 & 288
	CREATIVE_IMAGE_MODE_UNION_SPLASH       ImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH"       // 穿山甲开屏图片; 仅限穿山甲开屏使用，下限：1080 & 1920，上限：2160 & 3840，比例0.56
	CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO ImageMode = "CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO" // 穿山甲开屏视频; 仅限穿山甲开屏使用，宽高比（0.56，9:16）,视频码率≥516kbps,大小≤100M,分辨率≥720*1280,1s<时长<6s
	CREATIVE_IMAGE_MODE_DISPLAY_WINDOW     ImageMode = "CREATIVE_IMAGE_MODE_DISPLAY_WINDOW"     // 搜索橱窗，宽高比1，下限：109 & 109，上限：109 & 109
	MATERIAL_IMAGE_MODE_TITLE              ImageMode = "MATERIAL_IMAGE_MODE_TITLE"              // 标题类型，非创意的素材类型，仅报表接口会区分此素材类型
	CREATIVE_IMAGE_MODE_AWEME_LIVE         ImageMode = "CREATIVE_IMAGE_MODE_AWEME_LIVE"         // 直播画面类型
)