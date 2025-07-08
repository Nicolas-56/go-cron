package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Database struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

type Infoblock struct {
	StoresIblockId int
	CitiesIblockId int
}

type Config struct {
	Database  Database
	Infoblock Infoblock
}

func GetConfig() *Config {

	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env")
	}

	storesIblockId, errorConvert := strconv.Atoi(getEnv("STORES_IBLOCK_ID", "5"))
	if errorConvert != nil {
		fmt.Println("Ошибка при конвертации:", err)
	} else {
		fmt.Println("Конвертированное значение:", storesIblockId)
	}

	citiesIblockId, errorConvert := strconv.Atoi(getEnv("CITIES_IBLOCK_ID", "4"))
	if errorConvert != nil {
		fmt.Println("Ошибка при конвертации:", err)
	} else {
		fmt.Println("Конвертированное значение:", storesIblockId)
	}
	return &Config{
		Database: Database{
			DbHost: getEnv("DB_HOST", "localhost"),
			DbPort: getEnv("DB_PORT", "3306"),
			DbUser: getEnv("DB_USER", "root"),
			DbPass: getEnv("DB_PASS", "root"),
			DbName: getEnv("DB_NAME", "dbname"),
		},
		Infoblock: Infoblock{
			StoresIblockId: storesIblockId,
			CitiesIblockId: citiesIblockId,
		},
	}
}

// getEnv получает значение переменной окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
