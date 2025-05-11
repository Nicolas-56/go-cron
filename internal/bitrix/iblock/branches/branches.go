package branches

import (
	"fmt"

	"home.ru/internal/bitrix/iblock/properties"
)

const IBLOCK_ID = 4

func GetList() {
	properties := getPropertiesIblock()
	fmt.Println(properties)
}

func getPropertiesIblock() map[string]properties.PropertyType {
	propertiesIblock := properties.GetList(IBLOCK_ID)
	return propertiesIblock
}
