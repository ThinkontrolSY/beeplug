package beeplug

import (
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
	exception := f.checkException(v)
	timeNow := time.Now().UnixNano()
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
	if f.Threshold == nil {
		return !(reflect.TypeOf(f.snapshot) == reflect.TypeOf(value) && reflect.DeepEqual(f.snapshot, value))
	}

	newValue := getFloat64(value)
	oldValue := getFloat64(f.snapshot)
	if newValue != nil && oldValue != nil {
		return *newValue < *oldValue-*f.Threshold || *newValue > *oldValue+*f.Threshold
	}

	return false
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
		Name:     "Variables." + f.Name,
		DataType: f.DataType,
		Status:   f.Status,
	}
	switch f.DataType {
	case DataType_Boolean:
		v, t := f.Get(previous)
		metric.Value = &Payload_Metric_BooleanValue{
			BooleanValue: v.(bool),
		}
		metric.Timestamp = t
	case DataType_Int8:
		v, t := f.Get(previous)
		metric.Value = &Payload_Metric_Int32Value{
			Int32Value: int32(v.(int8)),
		}
		metric.Timestamp = t
	case DataType_UInt8:
		v, t := f.Get(previous)
		metric.Value = &Payload_Metric_Uint32Value{
			Uint32Value: uint32(v.(uint8)),
		}
		metric.Timestamp = t
	case DataType_Int16:
		v, t := f.Get(previous)
		metric.Value = &Payload_Metric_Int32Value{
			Int32Value: int32(v.(int16)),
		}
		metric.Timestamp = t
	case DataType_UInt16:
		v, t := f.Get(previous)
		metric.Value = &Payload_Metric_Uint32Value{
			Uint32Value: uint32(v.(uint16)),
		}
		metric.Timestamp = t
	case DataType_Int32:
		v, t := f.Get(previous)
		metric.Value = &Payload_Metric_Int32Value{
			Int32Value: v.(int32),
		}
		metric.Timestamp = t
	case DataType_UInt32:
		v, t := f.Get(previous)
		metric.Value = &Payload_Metric_Uint32Value{
			Uint32Value: v.(uint32),
		}
		metric.Timestamp = t
	case DataType_Int64:
		v, t := f.Get(previous)
		metric.Value = &Payload_Metric_Int64Value{
			Int64Value: v.(int64),
		}
		metric.Timestamp = t
	case DataType_UInt64:
		v, t := f.Get(previous)
		metric.Value = &Payload_Metric_Uint64Value{
			Uint64Value: v.(uint64),
		}
		metric.Timestamp = t
	case DataType_Float:
		v, t := f.Get(previous)
		metric.Value = &Payload_Metric_FloatValue{
			FloatValue: v.(float32),
		}
		metric.Timestamp = t
	case DataType_Double:
		v, t := f.Get(previous)
		metric.Value = &Payload_Metric_DoubleValue{
			DoubleValue: v.(float64),
		}
		metric.Timestamp = t
	case DataType_String, DataType_WString:
		v, t := f.Get(previous)
		metric.Value = &Payload_Metric_StringValue{
			StringValue: v.(string),
		}
		metric.Timestamp = t
	case DataType_Bytes:
		v, t := f.Get(previous)
		metric.Value = &Payload_Metric_BytesValue{
			BytesValue: v.([]byte),
		}
		metric.Timestamp = t
	}
	return metric
}
