package sqlserver

import (
    "fmt"
    "main/models"
    "gorm.io/driver/sqlserver"
    "gorm.io/gorm"
)

var Database * gorm.DB

func Init() {
    dsn: = "sqlserver://sa:sa123@XI065?database=TestDB"
    var err error
    Database,
    err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

    Database.AutoMigrate(&models.Cliente{})
    if err != nil {
        fmt.Println(err.Error())
    }
}