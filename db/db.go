package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"hcpb-api/configs"
)

// DB is the database connection
var DB *gorm.DB

// Init is the function to initialize the database connection
func Init() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Rome",
		configs.DB_HOST, configs.DB_USER, configs.DB_PASS, configs.DB_NAME, configs.DB_PORT)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}

func QueryAllCalls() ([]Call, error) {
	var calls []Call
	result := DB.Find(&calls)
	if result.Error != nil {
		log.Println(result.Error)
		return calls, result.Error
	}
	return calls, nil
}

func InsertCall(call Call) error {
	result := DB.Create(&call)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

// scrivo una query che conta tutte le chiamate effettuate da un chat_id nel mese corrente
func CountCallsByChatID(chatID string) (int64, error) {
	var count int64
	result := DB.Model(&Call{}).Where("chat_id = ? AND date_trunc('month', created_at) = date_trunc('month', now())", chatID).Count(&count)
	if result.Error != nil {
		log.Println(result.Error)
		return count, result.Error
	}
	return count, nil
}

func QueryLastCalls() ([]CallInfo, error) {
    var calls []CallInfo
    result := DB.Table("calls").
        Select("services.name AS service_name, types.name AS type_name, calls.chat_id, calls.coin, calls.created_at").
        Joins("JOIN services ON calls.service_id = services.id").
        Joins("JOIN types ON calls.type_id = types.id").
        Order("calls.created_at desc").
        Limit(200).
        Scan(&calls)
    if result.Error != nil {
        log.Println(result.Error)
        return calls, result.Error
    }
    return calls, nil
}

