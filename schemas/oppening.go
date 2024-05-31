package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Oppening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

// cria uma objeto json e iformar para o go que vai receber um json com o nome que vai chegar na requisicao
type OppeningResponse struct {
	ID        uint      `json: "id"`
	CreatedAt time.Time `json: "createdate"`
	UpdatedAt time.Time `json: "updatedat"`
	DeletedAt time.Time `json: "deletedat, omitempty"`
	Role      string    `json: role"`
	Company   string    `json: "company"`
	Location  string    `json: "location"`
	Remote    bool      `json: "remote"`
	Link      string    `json: "link"`
	Salary    int64     `json: "salary"`
}
