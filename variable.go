package beeplug

import (
	"fmt"
	reflect "reflect"
	sync "sync"
	"time"
)

type BeeVariable struct {
	Name         string
	DataType     DataType
	Status       string
	mux          sync.Mutex
	snapshot     interface{}
	snapshotTime uint64
	previous     interface{}
	previousTime uint64
	Threshold    *float64
	Interval     *uint64 // ms
}

func (f *BeeVariable) Get(previous bool) (interface{}, uint64) {
	f.mux.Lock()
	defer f.mux.Unlock()
	if previous {
		return f.previous, f.previousTime
	} else {
		return f.snapshot, f.snapshotTime
	}
}

func (f *BeeVariable) GetPrevious() interface{} {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.previous
}

func (f *BeeVariable) GetPreviousTime() uint64 {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.previousTime
}

func (f *BeeVariable) GetSnapshot() interface{} {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.snapshot
}

func (f *BeeVariable) GetSnapshotTime() uint64 {
	f.mux.Lock()
	defer f.mux.Unlock()
	return f.snapshotTime
}

func (f *BeeVariable) WriteValue(v interface{}, status string) bool {
	f.mux.Lock()
	defer f.mux.Unlock()
	defer func() {
		f.Status = status
	}()
	timeNow := time.Now().UnixNano()
	exception := f.checkException(v)
	if exception {
		f.snapshot = v
		f.snapshotTime = uint64(timeNow)
	} else {
		f.previous = v
		f.previousTime = uint64(timeNow)
	}
	return exception || f.Status != status
}

func (f *BeeVariable) checkException(value interface{}) bool {
	if value == nil {
		return false
	}
	if f.snapshot == nil {
		return true
	}

	if f.Threshold != nil {
		newValue := getFloat64(value)
		oldValue := getFloat64(f.snapshot)
		if newValue != nil && oldValue != nil {
			return *newValue < *oldValue-*f.Threshold || *newValue > *oldValue+*f.Threshold
		}
	} else if f.DataType != DataType_Float && f.DataType != DataType_Double {
		return !(reflect.TypeOf(f.snapshot) == reflect.TypeOf(value) && reflect.DeepEqual(f.snapshot, value))
	}

	if f.Interval == nil {
		// default interval is 1 hour
		defaultInterval := uint64(3600000)
		f.Interval = &defaultInterval
	}

	return time.Now().UnixNano()-int64(f.snapshotTime) > int64(*f.Interval)*int64(1000000)
}

func getFloat64(value interface{}) *float64 {
	var fv float64
	switch v := value.(type) {
	case int:
		fv = float64(v)
	case int8:
		fv = float64(v)
	case int16:
		fv = float64(v)
	case int32:
		fv = float64(v)
	case int64:
		fv = float64(v)
	case uint:
		fv = float64(v)
	case uint8:
		fv = float64(v)
	case uint16:
		fv = float64(v)
	case uint32:
		fv = float64(v)
	case uint64:
		fv = float64(v)
	case float32:
		fv = float64(v)
	case float64:
		fv = v
	default:
		return nil
	}
	return &fv
}

func (f *BeeVariable) GetMetricPayloads() []*Payload_Metric {
	if f.previous == nil {
		return []*Payload_Metric{f.getMetricPayload(false)}
	}
	return []*Payload_Metric{f.getMetricPayload(true), f.getMetricPayload(false)}
}

func (f *BeeVariable) getMetricPayload(previous bool) *Payload_Metric {
	metric := &Payload_Metric{
		Name:     f.Name,
		DataType: f.DataType,
		Status:   f.Status,
	}
	v, t := f.Get(previous)
	metric.SetValue(v)
	metric.Timestamp = t
	return metric
}

func (metric *Payload_Metric) SetValue(value interface{}) error {
	switch metric.DataType {
	case DataType_Boolean:
		b, ok := value.(bool)
		if !ok {
			return fmt.Errorf("value: %T is not a boolean", value)
		}
		metric.Value = &Payload_Metric_BooleanValue{
			BooleanValue: b,
		}
	case DataType_Int8, DataType_Int16, DataType_Int32:
		switch v := value.(type) {
		case int:
			metric.Value = &Payload_Metric_Int32Value{
				Int32Value: int32(v),
			}
		case int8:
			metric.Value = &Payload_Metric_Int32Value{
				Int32Value: int32(v),
			}
		case int16:
			metric.Value = &Payload_Metric_Int32Value{
				Int32Value: int32(v),
			}
		case int32:
			metric.Value = &Payload_Metric_Int32Value{
				Int32Value: v,
			}
		case int64:
			metric.Value = &Payload_Metric_Int32Value{
				Int32Value: int32(v),
			}
		case uint:
			metric.Value = &Payload_Metric_Int32Value{
				Int32Value: int32(v),
			}
		case uint8:
			metric.Value = &Payload_Metric_Int32Value{
				Int32Value: int32(v),
			}
		case uint16:
			metric.Value = &Payload_Metric_Int32Value{
				Int32Value: int32(v),
			}
		case uint32:
			metric.Value = &Payload_Metric_Int32Value{
				Int32Value: int32(v),
			}
		case uint64:
			metric.Value = &Payload_Metric_Int32Value{
				Int32Value: int32(v),
			}
		default:
			return fmt.Errorf("value: %T cannot be converted to int32", value)
		}
	case DataType_UInt8, DataType_UInt16, DataType_UInt32:
		switch v := value.(type) {
		case int:
			metric.Value = &Payload_Metric_Uint32Value{
				Uint32Value: uint32(v),
			}
		case int8:
			metric.Value = &Payload_Metric_Uint32Value{
				Uint32Value: uint32(v),
			}
		case int16:
			metric.Value = &Payload_Metric_Uint32Value{
				Uint32Value: uint32(v),
			}
		case int32:
			metric.Value = &Payload_Metric_Uint32Value{
				Uint32Value: uint32(v),
			}
		case int64:
			metric.Value = &Payload_Metric_Uint32Value{
				Uint32Value: uint32(v),
			}
		case uint:
			metric.Value = &Payload_Metric_Uint32Value{
				Uint32Value: uint32(v),
			}
		case uint8:
			metric.Value = &Payload_Metric_Uint32Value{
				Uint32Value: uint32(v),
			}
		case uint16:
			metric.Value = &Payload_Metric_Uint32Value{
				Uint32Value: uint32(v),
			}
		case uint32:
			metric.Value = &Payload_Metric_Uint32Value{
				Uint32Value: uint32(v),
			}
		case uint64:
			metric.Value = &Payload_Metric_Uint32Value{
				Uint32Value: uint32(v),
			}
		default:
			return fmt.Errorf("value: %T cannot be converted to uint32", value)
		}
	case DataType_Int64:
		switch v := value.(type) {
		case int:
			metric.Value = &Payload_Metric_Int64Value{
				Int64Value: int64(v),
			}
		case int8:
			metric.Value = &Payload_Metric_Int64Value{
				Int64Value: int64(v),
			}
		case int16:
			metric.Value = &Payload_Metric_Int64Value{
				Int64Value: int64(v),
			}
		case int32:
			metric.Value = &Payload_Metric_Int64Value{
				Int64Value: int64(v),
			}
		case int64:
			metric.Value = &Payload_Metric_Int64Value{
				Int64Value: int64(v),
			}
		case uint:
			metric.Value = &Payload_Metric_Int64Value{
				Int64Value: int64(v),
			}
		case uint8:
			metric.Value = &Payload_Metric_Int64Value{
				Int64Value: int64(v),
			}
		case uint16:
			metric.Value = &Payload_Metric_Int64Value{
				Int64Value: int64(v),
			}
		case uint32:
			metric.Value = &Payload_Metric_Int64Value{
				Int64Value: int64(v),
			}
		case uint64:
			metric.Value = &Payload_Metric_Int64Value{
				Int64Value: int64(v),
			}
		default:
			return fmt.Errorf("value: %T cannot be converted to int64", value)
		}
	case DataType_UInt64:
		switch v := value.(type) {
		case int:
			metric.Value = &Payload_Metric_Uint64Value{
				Uint64Value: uint64(v),
			}
		case int8:
			metric.Value = &Payload_Metric_Uint64Value{
				Uint64Value: uint64(v),
			}
		case int16:
			metric.Value = &Payload_Metric_Uint64Value{
				Uint64Value: uint64(v),
			}
		case int32:
			metric.Value = &Payload_Metric_Uint64Value{
				Uint64Value: uint64(v),
			}
		case int64:
			metric.Value = &Payload_Metric_Uint64Value{
				Uint64Value: uint64(v),
			}
		case uint:
			metric.Value = &Payload_Metric_Uint64Value{
				Uint64Value: uint64(v),
			}
		case uint8:
			metric.Value = &Payload_Metric_Uint64Value{
				Uint64Value: uint64(v),
			}
		case uint16:
			metric.Value = &Payload_Metric_Uint64Value{
				Uint64Value: uint64(v),
			}
		case uint32:
			metric.Value = &Payload_Metric_Uint64Value{
				Uint64Value: uint64(v),
			}
		case uint64:
			metric.Value = &Payload_Metric_Uint64Value{
				Uint64Value: uint64(v),
			}
		default:
			return fmt.Errorf("value: %T cannot be converted to uint64", value)
		}
	case DataType_Float:
		switch v := value.(type) {
		case float32:
			metric.Value = &Payload_Metric_FloatValue{
				FloatValue: v,
			}
		case float64:
			metric.Value = &Payload_Metric_FloatValue{
				FloatValue: float32(v),
			}
		default:
			return fmt.Errorf("value: %T cannot be converted to float32", value)
		}
	case DataType_Double:
		switch v := value.(type) {
		case float32:
			metric.Value = &Payload_Metric_DoubleValue{
				DoubleValue: float64(v),
			}
		case float64:
			metric.Value = &Payload_Metric_DoubleValue{
				DoubleValue: v,
			}
		default:
			return fmt.Errorf("value: %T cannot be converted to float64", value)
		}
	case DataType_String, DataType_WString:
		s, ok := value.(string)
		if !ok {
			return fmt.Errorf("value: %T is not a string", value)
		}
		metric.Value = &Payload_Metric_StringValue{
			StringValue: s,
		}
	case DataType_Bytes:
		b, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value: %T is not a byte array", value)
		}
		metric.Value = &Payload_Metric_BytesValue{
			BytesValue: b,
		}
	default:
		return fmt.Errorf("unsupported data type: %v", metric.DataType)
	}
	return nil
}
