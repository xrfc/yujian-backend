package model

// EsModel 定义了一个ES模型
type EsModel interface {
	// SetScore 设置对象的评分
	SetScore(score float64)

	// GetScore 获取对象的评分
	GetScore() float64

	// GetID 获取对象的唯一标识符
	GetID() string

	// 获取对象的索引名称
	GetIndexName() string

	// 获取内容
	GetContent() string

	// 获取标题
	GetTitle() string
}
