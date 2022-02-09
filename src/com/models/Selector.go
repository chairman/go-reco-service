package models

import (
	"encoding/json"
	"strconv"
)

type Selector interface {
	Judge(context Context) bool
	getName() string
}

type Eq struct {
	key   string
	value interface{}
}

func (o *Eq) Judge(context Context) bool {
	current := context.Params[o.key]
	if current == "" {
		return false
	}
	value := Strval(o.value)
	return current == value
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

func (o *Eq) getName() string {
	return "$eq"
}

type True struct {
}

func (o *True) Judge(context Context) bool {
	return true
}

func (o *True) getName() string {
	return "$true"
}

type And struct {
	selectors []Selector
}

func (o *And) Judge(context Context) bool {
	if o.selectors != nil {
		for _, selector := range o.selectors {
			if !selector.Judge(context) {
				return false
			}
		}
	}
	return true
}

func (o *And) getName() string {
	return "$and"
}

type Or struct {
	selectors []Selector
}

func (o *Or) Judge(context Context) bool {
	if len(o.selectors) == 0 {
		return true
	}
	if o.selectors != nil {
		for _, selector := range o.selectors {
			if selector.Judge(context) {
				return true
			}
		}
	}
	return false
}

func (o *Or) getName() string {
	return "$or"
}

type Gt struct {
}

func (o *Gt) Judge(context Context) bool {
	return true
}

func (o *Gt) getName() string {
	return "$gt"
}

type Gte struct {
}

func (o *Gte) Judge(context Context) bool {
	return true
}

func (o *Gte) getName() string {
	return "$gte"
}

type Lt struct {
}

func (o *Lt) Judge(context Context) bool {
	return true
}

func (o *Lt) getName() string {
	return "$lt"
}

type Lte struct {
}

func (o *Lte) Judge(context Context) bool {
	return true
}

func (o *Lte) getName() string {
	return "$lte"
}
