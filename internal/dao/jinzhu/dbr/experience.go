package dbr

import (
	"time"

	"gorm.io/gorm"
)

// ExperienceChangeLog 经验值变动日志模型
type ExperienceChangeLog struct {
	LogID     int       `gorm:"column:log_id;primaryKey" json:"log_id"` // 日志ID（主键）
	UserID    int       `gorm:"column:user_id" json:"user_id"`          // 用户ID
	Action    int       `gorm:"column:action" json:"action"`            // 触发行为（对应动作类型编码）
	Amount    int       `gorm:"column:amount" json:"amount"`            // 经验值变动量
	Remark    int       `gorm:"column:remark" json:"remark"`            // 备注（若为状态码可对应文本说明）
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`    // 变动时间
}

// TableName 指定表名映射
func (e *ExperienceChangeLog) TableName() string {
	return "p_experience_change_log"
}

// Create 创建经验值变动记录
func (e *ExperienceChangeLog) Create(db *gorm.DB) (*ExperienceChangeLog, error) {
	err := db.Create(&e).Error
	return e, err
}

// ListByUserID 根据用户ID查询日志（示例方法）
func (e *ExperienceChangeLog) ListByUserID(db *gorm.DB, userID int, page, size int) ([]*ExperienceChangeLog, error) {
	var logs []*ExperienceChangeLog
	err := db.Where("user_id = ?", userID).
		Offset((page - 1) * size).
		Limit(size).
		Order("updated_at DESC").
		Find(&logs).Error
	return logs, err
}
