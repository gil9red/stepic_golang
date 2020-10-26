package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
Данная задача ориентирована на реальную работу с данными в формате JSON. Для решения мы будем
использовать справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности),
который был размещен на web-портале data.gov.ru.

Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json,
а сами данные, которые вам потребуется декодировать, содержатся в файле data-20190514T0100.json.
Файлы размещены в https://github.com/semyon-dev/stepik-go/tree/master/work_with_json.

Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо в
качестве ответа записать сумму полей global_id всех элементов, закодированных в файле.
*/

type Item struct {
	SystemObjectId string `json:"system_object_id"`
	Kod            string `json:"Kod"`
	IsDeleted      int    `json:"is_deleted"`
	SignatureDate  string `json:"signature_date"`
	Nomdescr       string `json:"Nomdescr"`
	GlobalId       int    `json:"global_id"`
	Idx            string `json:"Idx"`
	Razdel         string `json:"Razdel"`
	Name           string `json:"Name"`
}

func main() {
	const URL = "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"

	rs, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	dec := json.NewDecoder(rs.Body)

	var items []Item
	err = dec.Decode(&items)
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, x := range items {
		sum += x.GlobalId
	}
	fmt.Println(sum)
	// 763804981288
}
