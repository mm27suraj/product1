package model

type Inventory struct {
	ProductId    uint   `gorm:"primary_key;auto_increment" json:"prodid"`
	ProductName  string `json:"prodname" gorm:"type:varchar(100)"`
	ProductQty   int    `json:"prodqty" gorm:"type:int"`
	ProductPrice int    `json:"prodprice" gorm:"type:int"`
}

func (c *Inventory) TableName() string {
	return "inventory"
}
