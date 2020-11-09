package web

type IDData struct {
	ID   int
	Data interface{}
}

type IDDataContainer struct {
	Ul []*IDData
}

// 获得数据 ID 或初始化数据 ID
func (c *IDDataContainer) IDs(data []interface{}) []int {
	var ids []int
	for _, datum := range data {
		var id int
		func() {
			for _, li := range c.Ul {
				if li == datum {
					id = li.ID
					return
				}
			}
			id = len(c.Ul) + 1
			c.Ul = append(c.Ul, &IDData{
				ID:   id,
				Data: datum,
			})
		}()
		ids = append(ids, id)
	}
	return ids
}
