package elementproperties

import (
	"fmt"
	"log"
	"strings"

	"home.ru/internal/database"
)

type ElementProperty struct {
	ID                 int
	IBLOCK_PROPERTY_ID int
	IBLOCK_ELEMENT_ID  int
	VALUE              string
}

func GetList(ids []int) {
	propertiesElement := make(map[int]ElementProperty)
	var (
		id               int
		iblockPropertyId int
		iblockElementId  int
		value            string
	)

	db := database.GetConnection()
	idsForSql := database.ConvertToInterfaceSlice(ids)
	query := `SELECT id, iblock_property_id, iblock_element_id, value FROM b_iblock_element_property WHERE iblock_element_id IN(` + strings.Repeat("?", len(ids)) + `)`
	fmt.Println(query)
	rows, err := db.Query(query, idsForSql...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&id,
			&iblockPropertyId,
			&iblockElementId,
			&value,
		)
		if err != nil {
			log.Fatal(err)
		}

		propertiesElement[id] = ElementProperty{
			ID:                 id,
			IBLOCK_PROPERTY_ID: iblockPropertyId,
			IBLOCK_ELEMENT_ID:  iblockElementId,
			VALUE:              value,
		}
	}

	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(propertiesElement)
}
