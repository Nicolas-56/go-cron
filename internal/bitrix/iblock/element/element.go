package element

import (
	"log"
	"strings"

	"home.ru/internal/database"
)

type ParamsElementsGetList struct {
	Offset               int
	Limit                int
	IblockID             int
	SelectPropertiesCode []string
}

type Element struct {
	NAME  string
	CODE  string
	VALUE string
}

func GetList(paramsGetList ParamsElementsGetList) map[string][]Element {
	propertiesElement := make(map[string][]Element)
	var (
		id    string
		name  string
		code  string
		value string
	)

	db, _ := database.GetConnection()
	query := "SELECT " +
		"`el`.ID as ID," +
		"`prop`.NAME AS Name, " +
		"`prop`.CODE AS Code, " +
		"IF (`prop`.PROPERTY_TYPE = 'F', IFNULL (CONCAT(`file`.SUBDIR, '/', `file`.FILE_NAME), ''), IFNULL(`prop_enum`.VALUE, `el_prop`.VALUE)) AS Value  " +
		"FROM `b_iblock_element` el " +
		"LEFT JOIN `b_iblock_element_property` el_prop ON `el_prop`.IBLOCK_ELEMENT_ID = `el`.ID " +
		"LEFT JOIN `b_iblock_property` prop ON `prop`.ID = `el_prop`.IBLOCK_PROPERTY_ID " +
		"LEFT JOIN `b_iblock_property_enum` prop_enum ON `prop_enum`.ID = `el_prop`.VALUE " +
		"LEFT JOIN `b_file` file ON `file`.ID = `el_prop`.VALUE " +
		"WHERE `el`.ACTIVE = 'Y'"

	params := []any{}
	whereParts := []string{}
	if paramsGetList.IblockID != 0 {
		whereParts = append(whereParts, "el.IBLOCK_ID = ?")
		params = append(params, paramsGetList.IblockID)
	}

	if len(paramsGetList.SelectPropertiesCode) != 0 {
		placeholders := make([]string, len(paramsGetList.SelectPropertiesCode))
		for i, val := range paramsGetList.SelectPropertiesCode {
			placeholders[i] = "?" // используя плейсхолдеры
			params = append(params, val)
		}
		whereParts = append(whereParts, "prop.CODE IN ("+strings.Join(placeholders, ", ")+")")
	}

	where := ""
	// Объединяем все части условия where
	if len(whereParts) > 0 {
		where = strings.Join(whereParts, " AND ")
	}

	if where != "" {
		query += " AND " + where
	}

	//fmt.Println(query)
	rows, err := db.DB.Query(query, params...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	i := 0
	for rows.Next() {
		err := rows.Scan(
			&id,
			&name,
			&code,
			&value,
		)
		if err != nil {
			log.Fatal(err)
		}

		propertiesElement[id] = append(propertiesElement[id], Element{
			NAME:  name,
			CODE:  code,
			VALUE: value,
		})
		i++
	}

	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	return propertiesElement
}
