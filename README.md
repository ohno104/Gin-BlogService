# Gin-BlogService

使用gin framework練習實作blog service

## Descript

使用gin framework練習實作blog service，根據RESTful API的規則設計業務模型的增刪改查功能，使用Swagger生成API文件並採用JWT作為安全驗證，
service呼叫後使用validator函數庫的功能進行參數驗證後，透過ORM的方式存取、修改MySQL DB的資料，添加逾時控制、系統異常信件發送的middleware，
並使用viper函數庫作為系統參數設定的管理

## Related
* gin framework
* RESTful API
* swaggo
* jwt-go
* gorm
* gomail.v2
* viper
* validator
