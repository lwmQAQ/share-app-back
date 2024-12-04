package enum

type SearchType int

const (
	SearchTypeMache = 1 << iota // 1
	SearchTypeImage             // 2
	SearchTypeVideo             // 4
)

// 定义搜索类型的映射
var searchTypeHints = map[SearchType]string{
	SearchTypeMache: "搜索文本信息",
	SearchTypeImage: "搜索图片资源",
	SearchTypeVideo: "搜索视频内容",
}

func GetTips(Type SearchType) string {
	if str, ok := searchTypeHints[Type]; ok {
		return str
	}
	return "Unkown"
}
