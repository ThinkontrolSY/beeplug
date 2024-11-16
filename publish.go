package beeplug

type Publish struct {
	Topic   string
	Payload *Payload
}

func (metric *Payload_Metric) GetMetricValue() interface{} {
	switch DataType(metric.GetDataType()) {
	case DataType_Int8:
		return int8(metric.GetInt32Value())
	case DataType_Int16:
		return int16(metric.GetInt32Value())
	case DataType_Int32:
		return metric.GetInt32Value()
	case DataType_Int64:
		return metric.GetInt64Value()
	case DataType_UInt8:
		return uint8(metric.GetUint32Value())
	case DataType_UInt16:
		return uint16(metric.GetUint32Value())
	case DataType_UInt32:
		return metric.GetUint32Value()
	case DataType_UInt64:
		return metric.GetUint64Value()
	case DataType_Float:
		return metric.GetFloatValue()
	case DataType_Double:
		return metric.GetDoubleValue()
	case DataType_Boolean:
		return metric.GetBooleanValue()
	case DataType_String, DataType_WString:
		return metric.GetStringValue()
	case DataType_Bytes:
		return metric.GetBytesValue()
	case DataType_Response:
		return metric.GetResponseValue()
	default:
		return nil
	}
}
