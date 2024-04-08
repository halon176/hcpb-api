package db

import (
	"database/sql"
	"fmt"
	"hcpb-api/configs"
	"hcpb-api/models"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Main() {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Rome",
	configs.DB_HOST, configs.DB_USER, configs.DB_PASS, configs.DB_NAME, configs.DB_PORT)
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

}

func GetLastCallsDriver() (string, error) {
    var jsonString string
    err := db.QueryRow(`SELECT json_agg(c) FROM (SELECT services.name as service_name, types.name as type_name, calls.chat_id, calls.coin, calls.created_at FROM calls JOIN services ON calls.service_id = services.id JOIN types ON calls.type_id = types.id ORDER BY calls.created_at DESC LIMIT 10) c`).Scan(&jsonString)
    if err != nil {
        log.Println(err)
        return "", err
    }
    return jsonString, nil
}

func GetStatistics() (string, error) {
	var jsonString string
	err := db.QueryRow(`SELECT json_agg(c) from (SELECT * from get_statistics()) c`).Scan(&jsonString)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return jsonString, nil
}

func InsertCall(call models.Call) error {
	_, err := db.Exec(`INSERT INTO calls (service_id, type_id, chat_id, coin) VALUES ($1, $2, $3, $4)`, call.ServiceID, call.TypeID, call.ChatID, call.Coin)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}


