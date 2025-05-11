package element

import (
	"fmt"
	"log"

	"home.ru/internal/database"
)

const (
	PROPERTY_TYPE_S = "S"
	PROPERTY_TYPE_L = "L"
	PROPERTY_TYPE_E = "E"
)

type Element struct {
	NAME  string
	CODE  string
	VALUE string
}

func GetList() {
	propertiesElement := make(map[string]Element)
	var (
		name  string
		code  string
		value string
	)

	db := database.GetConnection()
	query := "SELECT " +
		"`prop`.NAME AS Name, " +
		"`prop`.CODE AS Code, " +
		"IF (`prop`.PROPERTY_TYPE = 'F', IFNULL (CONCAT(`file`.SUBDIR, '/', `file`.FILE_NAME), ''), IFNULL(`prop_enum`.VALUE, `el_prop`.VALUE)) AS Value  " +
		"FROM `b_iblock_element` el " +
		"LEFT JOIN `b_iblock_element_property` el_prop ON `el_prop`.IBLOCK_ELEMENT_ID = `el`.ID " +
		"LEFT JOIN `b_iblock_property` prop ON `prop`.ID = `el_prop`.IBLOCK_PROPERTY_ID " +
		"LEFT JOIN `b_iblock_property_enum` prop_enum ON `prop_enum`.ID = `el_prop`.VALUE " +
		"LEFT JOIN `b_file` file ON `file`.ID = `el_prop`.VALUE " +
		"WHERE `el`.ID = 322 AND `el`.IBLOCK_ID = 5"
	fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&name,
			&code,
			&value,
		)
		if err != nil {
			log.Fatal(err)
		}

		propertiesElement[name] = Element{
			NAME:  name,
			CODE:  code,
			VALUE: value,
		}
	}

	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(propertiesElement)
}
