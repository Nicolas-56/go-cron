package properties

import (
	"database/sql"
	"log"

	"home.ru/internal/database"
)

type PropertyType struct {
	ID             int
	NAME           string
	ACTIVE         bool
	CODE           string
	PROPERTY_TYPE  string
	LINK_IBLOCK_ID int
}

var properties = make(map[string]PropertyType)

func GetList(iblockId int) map[string]PropertyType {
	var (
		id           int
		name         string
		active       string
		code         string
		propertyType sql.NullString
		linkIblockId sql.NullInt64
	)

	db := database.GetConnection()
	query := `SELECT id, name, active, code, property_type, link_iblock_id  FROM b_iblock_property WHERE iblock_id=?`
	rows, err := db.Query(query, iblockId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&id,
			&name,
			&active,
			&code,
			&propertyType,
			&linkIblockId,
		)
		if err != nil {
			log.Fatal(err)
		}

		propertyTypeValue := database.GetValuesFromSqlString(propertyType)
		properties[code] = PropertyType{
			ID:             id,
			NAME:           name,
			ACTIVE:         active == "Y",
			CODE:           code,
			PROPERTY_TYPE:  propertyTypeValue,
			LINK_IBLOCK_ID: int(linkIblockId.Int64),
		}
	}

	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	return properties
}
