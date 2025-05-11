package iblock

import (
	"database/sql"
	"log"

	"home.ru/internal/database"
)

const IBLOCK_VERSION_2 = 2
const IBLOCK_VERSION_1 = 1

type IblockElement struct {
	id      int
	code    string
	version int
}

func GetListById(idIblock int) IblockElement {
	var (
		id      int
		code    sql.NullString
		version sql.NullInt32
	)

	db := database.GetConnection()
	query := `SELECT id, code, version FROM b_iblock WHERE id=?`
	rows, err := db.Query(query, idIblock)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	IblockElement := IblockElement{}

	for rows.Next() {
		err := rows.Scan(
			&id,
			&code,
			&version,
		)
		if err != nil {
			log.Fatal(err)
		}

		codeValue := database.ConvertSqlStringToString(code)
		versionValue := database.ConvertSqlInt32ToInt(version)
		IblockElement.id = id
		IblockElement.code = codeValue
		IblockElement.version = versionValue
	}

	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	return IblockElement
}

func GetVersionIblock(idIblock int) int {
	iblock := GetListById(idIblock)
	if iblock.id == IBLOCK_VERSION_2 {
		return IBLOCK_VERSION_2
	} else {
		return IBLOCK_VERSION_1
	}
}
