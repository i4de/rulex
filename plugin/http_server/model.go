package httpserver

import (
	"database/sql/driver"
	"time"

	"gopkg.in/square/go-jose.v2/json"
)

type RulexModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
}
type stringList []string

func (f stringList) Value() (driver.Value, error) {
	b, err := json.Marshal(f)
	return string(b), err
}

func (f *stringList) Scan(data interface{}) error {
	return json.Unmarshal([]byte(data.(string)), f)
}

type MRule struct {
	RulexModel
	UUID        string `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string
	FromSource  stringList `gorm:"not null type:string[]"`
	FromDevice  stringList `gorm:"not null type:string[]"`
	Actions     string     `gorm:"not null"`
	Success     string     `gorm:"not null"`
	Failed      string     `gorm:"not null"`
}

type MInEnd struct {
	RulexModel
	// UUID for origin source ID
	UUID        string `gorm:"not null"`
	Type        string `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string
	Config      string
	XDataModels string
}

type MOutEnd struct {
	RulexModel
	// UUID for origin source ID
	UUID        string `gorm:"not null"`
	Type        string `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string
	Config      string
}

type MUser struct {
	RulexModel
	Role        string `gorm:"not null"`
	Username    string `gorm:"not null"`
	Password    string `gorm:"not null"`
	Description string
}

// 设备元数据
type MDevice struct {
	RulexModel
	UUID         string `gorm:"not null"`
	Name         string `gorm:"not null"`
	Type         string `gorm:"not null"`
	ActionScript string
	Config       string
	Description  string
}

//
// 外挂
//

type MGoods struct {
	RulexModel
	UUID        string     `gorm:"not null"`
	Addr        string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	Args        stringList `gorm:"not null"`
}
