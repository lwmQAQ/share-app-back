package models

// Task 表对应的 GORM 结构体
type Task struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`                     // 对应 tasks 表的主键 taskid
	UserID    uint   `gorm:"column:user_id"`                               // 对应 userid 字段，外键索引
	FileName  string `gorm:"column:file_name"`                             //文件名
	SourceURL string `gorm:"type:varchar(255);not null;column:source_url"` // 对应 sourceurl 字段
	MonoURL   string `gorm:"type:varchar(255);column:mono_url"`            // 对应 monourl 字段
	DualURL   string `gorm:"type:varchar(255);column:dual_url"`            // 对应 dualurl 字段
	Status    int    `gorm:"type:int;column:status"`                       // 任务状态
}
