package model

type ListMapItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func ToMap(items []ListMapItem) map[string]string {
	dataMap := make(map[string]string)
	for _, item := range items {
		dataMap[item.Key] = item.Value
	}
	return dataMap
}

func ToList(data map[string]string) []ListMapItem {
	listData := make([]ListMapItem, 0)
	for key, value := range data {
		listData = append(listData, ListMapItem{
			Key:   key,
			Value: value,
		})
	}
	return listData
}

func ToListWithByte(data map[string][]uint8) []ListMapItem {
	listData := make([]ListMapItem, 0)
	for key, value := range data {
		listData = append(listData, ListMapItem{
			Key:   key,
			Value: string(value),
		})
	}
	return listData
}
