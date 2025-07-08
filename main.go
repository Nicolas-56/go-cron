package main

import (
	"fmt"

	"home.ru/internal/bitrix/iblock/element"
	"home.ru/internal/config"
)

func main() {

	config := config.GetConfig()
	fmt.Print(config)
	//filepath.Walk("./internal/task", walkFile)

	paramsElementsForCities := element.ParamsElementsGetList{
		IblockID:             config.Infoblock.CitiesIblockId,
		SelectPropertiesCode: []string{"ID", "LID_AVARDA"},
	}

	cities := element.GetList(paramsElementsForCities)

	for id, value := range cities {
		fmt.Printf("%s ------ %s", id, value)
	}
	fmt.Print(cities)

	paramsElementsGetList := element.ParamsElementsGetList{
		IblockID:             config.Infoblock.StoresIblockId,
		SelectPropertiesCode: []string{"ID_FILIAL", "LID_AVARDA", "DEFAULT"},
	}

	//	mainStore:= make([]int)
	storesFromBitrix := element.GetList(paramsElementsGetList)

	fmt.Print(storesFromBitrix)

	// mainShoppingCenterRegion := make(map[int]int)

	// for i, val := range storesFromBitrix {
	// 	mainShoppingCenterRegion[i] :=
	// 	fmt.Println(i)
	// 	fmt.Println(val)
	// }

	//iblockProperties.GetList(4)
	//branches.GetList()

	//iblockElement := iblock.GetVersionIblock(4)
	//fmt.Println(s)
}
