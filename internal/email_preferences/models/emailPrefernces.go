package models

import (
    "time"
    "gorm.io/gorm"
)

type EmailPreference struct {
    ID        string         `gorm:"primaryKey;type:uuid;default:gen_random_uuid();column:id"`
    UserID    string            `gorm:"not null;type:uuid;column:userId"`
    Category  string         `gorm:"not null;column:category"`
    Type      string         `gorm:"not null;column:type"`
    IsEnabled bool           `gorm:"default:true;column:isEnabled"`
    CreatedAt time.Time      `gorm:"autoCreateTime;column:createdAt"`
    UpdatedAt time.Time      `gorm:"autoUpdateTime;column:updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index;column:deletedAt"`
}
