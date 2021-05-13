package utils

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"

	uuid "github.com/satori/go.uuid"
)

// String2Int 字符串转int
func String2Int(val string) int {

	goodsID, err := strconv.Atoi(val)
	if err != nil {
		return -1
	}
	return goodsID

}

// Int2String 数字转字符串
func Int2String(val int) string {
	return strconv.Itoa(val)
}

// Int642String Int642String
func Int642String(val int64) string {
	return strconv.FormatInt(val, 10)
}

// Float642String Float642String
func Float642String(val float64) string {
	return strconv.FormatFloat(val, 'E', -1, 64)
}

// GetUUID GetUUID
func GetUUID() string {
	UUID := uuid.NewV4()
	return UUID.String()
}

// GetTimestamp the result likes 1423361979
func GetTimestamp() int64 {
	return time.Now().Unix()
}

// FormatTimestamp the result likes 2015-02-08 10:19:39 AM
func FormatTimestamp(timestamp int64, format string) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format(format)
}

// ExactMapValues2Int64Array map转int64串数组
func ExactMapValues2Int64Array(maparray []orm.Params, key string) []int64 {

	var values []int64
	for _, value := range maparray {
		values = append(values, value[key].(int64))
	}
	return values
}

// ExactMapValues2StringArray map转字符串数组
func ExactMapValues2StringArray(mapArray []orm.Params, key string) []string {

	var values []string
	for _, value := range mapArray {
		values = append(values, value[key].(string))
	}
	return values
}

// PageData 页数结构体
type PageData struct {
	NumsPerPage int         `json:"pageSize"`
	CurrentPage int         `json:"currentPage"`
	Count       int         `json:"count"`
	TotalPages  int         `json:"totalPages"`
	Data        interface{} `json:"data"`
}

// GetPageData 获取页面数据
func GetPageData(rawData []orm.Params, page int, size int) PageData {

	count := len(rawData)
	totalPages := (count + size - 1) / size
	var pageData []orm.Params

	for idx := (page - 1) * size; idx < page*size && idx < count; idx++ {
		pageData = append(pageData, rawData[idx])
	}

	return PageData{NumsPerPage: size, CurrentPage: page, Count: count, TotalPages: totalPages, Data: pageData}
}

// ContainsInt 判断数组内是否包含某个值
func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// DateEqual 日期判断
func DateEqual(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
