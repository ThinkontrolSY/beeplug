package beeplug

import "errors"

func ConvertToType(input interface{}, targetType DataType) (interface{}, error) {
	switch targetType {
	case DataType_Boolean:
		if b, ok := input.(bool); ok {
			return b, nil
		} else {
			return false, errors.New("cannot convert to bool")
		}
	case DataType_Bytes:
		if b, ok := input.([]byte); ok {
			return b, nil
		} else {
			return nil, errors.New("cannot convert to []byte")
		}
	case DataType_String, DataType_WString:
		if s, ok := input.(string); ok {
			return s, nil
		} else {
			return "", errors.New("cannot convert to string")
		}
	case DataType_Int8:
		return ToInt8(input)
	case DataType_Int16:
		return ToInt16(input)
	case DataType_Int32:
		return ToInt32(input)
	case DataType_Int64:
		return ToInt64(input)
	case DataType_UInt8:
		return ToUint8(input)
	case DataType_UInt16:
		return ToUint16(input)
	case DataType_UInt32:
		return ToUint32(input)
	case DataType_UInt64:
		return ToUint64(input)
	case DataType_Float:
		return ToFloat32(input)
	case DataType_Double:
		return ToFloat64(input)
	default:
		return nil, errors.New("unsupported target type")
	}
}

func ToInt(input interface{}) (int, error) {
	switch v := input.(type) {
	case float64:
		return int(v), nil
	case float32:
		return int(v), nil
	case int:
		return v, nil
	case int64:
		return int(v), nil
	case int32:
		return int(v), nil
	case int16:
		return int(v), nil
	case int8:
		return int(v), nil
	case uint:
		return int(v), nil
	case uint64:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint8:
		return int(v), nil
	default:
		return 0, errors.New("cannot convert to int")
	}
}

func ToInt8(input interface{}) (int8, error) {
	switch v := input.(type) {
	case float64:
		return int8(v), nil
	case float32:
		return int8(v), nil
	case int:
		return int8(v), nil
	case int64:
		return int8(v), nil
	case int32:
		return int8(v), nil
	case int16:
		return int8(v), nil
	case int8:
		return v, nil
	case uint:
		return int8(v), nil
	case uint64:
		return int8(v), nil
	case uint32:
		return int8(v), nil
	case uint16:
		return int8(v), nil
	case uint8:
		return int8(v), nil
	default:
		return 0, errors.New("cannot convert to int8")
	}
}

// 此处省略 toInt16、toInt32、toInt64 等其他函数，使用类似的方法实现

func ToFloat32(input interface{}) (float32, error) {
	switch v := input.(type) {
	case float64:
		return float32(v), nil
	case float32:
		return v, nil
	case int:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case uint:
		return float32(v), nil
	case uint64:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	default:
		return 0, errors.New("cannot convert to float32")
	}
}

func ToFloat64(input interface{}) (float64, error) {
	switch v := input.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	default:
		return 0, errors.New("cannot convert to float64")
	}
}

// toInt16 将 input 转换为 int16
func ToInt16(input interface{}) (int16, error) {
	switch v := input.(type) {
	case float64:
		return int16(v), nil
	case float32:
		return int16(v), nil
	case int:
		return int16(v), nil
	case int64:
		return int16(v), nil
	case int32:
		return int16(v), nil
	case int16:
		return v, nil
	case int8:
		return int16(v), nil
	case uint:
		return int16(v), nil
	case uint64:
		return int16(v), nil
	case uint32:
		return int16(v), nil
	case uint16:
		return int16(v), nil
	case uint8:
		return int16(v), nil
	default:
		return 0, errors.New("cannot convert to int16")
	}
}

// toInt32 将 input 转换为 int32
func ToInt32(input interface{}) (int32, error) {
	switch v := input.(type) {
	case float64:
		return int32(v), nil
	case float32:
		return int32(v), nil
	case int:
		return int32(v), nil
	case int64:
		return int32(v), nil
	case int32:
		return v, nil
	case int16:
		return int32(v), nil
	case int8:
		return int32(v), nil
	case uint:
		return int32(v), nil
	case uint64:
		return int32(v), nil
	case uint32:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint8:
		return int32(v), nil
	default:
		return 0, errors.New("cannot convert to int32")
	}
}

// toInt64 将 input 转换为 int64
func ToInt64(input interface{}) (int64, error) {
	switch v := input.(type) {
	case float64:
		return int64(v), nil
	case float32:
		return int64(v), nil
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case int32:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	default:
		return 0, errors.New("cannot convert to int64")
	}
}

// toUint 将 input 转换为 uint
func ToUint(input interface{}) (uint, error) {
	switch v := input.(type) {
	case float64:
		return uint(v), nil
	case float32:
		return uint(v), nil
	case int:
		return uint(v), nil
	case int64:
		return uint(v), nil
	case int32:
		return uint(v), nil
	case int16:
		return uint(v), nil
	case int8:
		return uint(v), nil
	case uint:
		return v, nil
	case uint64:
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case uint8:
		return uint(v), nil
	default:
		return 0, errors.New("cannot convert to uint")
	}
}

// toUint8 将 input 转换为 uint8
func ToUint8(input interface{}) (uint8, error) {
	switch v := input.(type) {
	case float64:
		return uint8(v), nil
	case float32:
		return uint8(v), nil
	case int:
		return uint8(v), nil
	case int64:
		return uint8(v), nil
	case int32:
		return uint8(v), nil
	case int16:
		return uint8(v), nil
	case int8:
		return uint8(v), nil
	case uint:
		return uint8(v), nil
	case uint64:
		return uint8(v), nil
	case uint32:
		return uint8(v), nil
	case uint16:
		return uint8(v), nil
	case uint8:
		return v, nil
	default:
		return 0, errors.New("cannot convert to uint8")
	}
}

// toUint16 将 input 转换为 uint16
func ToUint16(input interface{}) (uint16, error) {
	switch v := input.(type) {
	case float64:
		return uint16(v), nil
	case float32:
		return uint16(v), nil
	case int:
		return uint16(v), nil
	case int64:
		return uint16(v), nil
	case int32:
		return uint16(v), nil
	case int16:
		return uint16(v), nil
	case int8:
		return uint16(v), nil
	case uint:
		return uint16(v), nil
	case uint64:
		return uint16(v), nil
	case uint32:
		return uint16(v), nil
	case uint16:
		return v, nil
	case uint8:
		return uint16(v), nil
	default:
		return 0, errors.New("cannot convert to uint16")
	}
}

// toUint32 将 input 转换为 uint32
func ToUint32(input interface{}) (uint32, error) {
	switch v := input.(type) {
	case float64:
		return uint32(v), nil
	case float32:
		return uint32(v), nil
	case int:
		return uint32(v), nil
	case int64:
		return uint32(v), nil
	case int32:
		return uint32(v), nil
	case int16:
		return uint32(v), nil
	case int8:
		return uint32(v), nil
	case uint:
		return uint32(v), nil
	case uint64:
		return uint32(v), nil
	case uint32:
		return v, nil
	case uint16:
		return uint32(v), nil
	case uint8:
		return uint32(v), nil
	default:
		return 0, errors.New("cannot convert to uint32")
	}
}

// toUint64 将 input 转换为 uint64
func ToUint64(input interface{}) (uint64, error) {
	switch v := input.(type) {
	case float64:
		return uint64(v), nil
	case float32:
		return uint64(v), nil
	case int:
		return uint64(v), nil
	case int64:
		return uint64(v), nil
	case int32:
		return uint64(v), nil
	case int16:
		return uint64(v), nil
	case int8:
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint64:
		return v, nil
	case uint32:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	default:
		return 0, errors.New("cannot convert to uint64")
	}
}
