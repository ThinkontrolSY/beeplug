package beeplug

import "encoding/json"

type ValueIntermediate struct {
	Int32Value    *int32            `json:"Int32Value,omitempty"`
	Int64Value    *int64            `json:"Int64Value,omitempty"`
	Uint32Value   *uint32           `json:"Uint32Value,omitempty"`
	Uint64Value   *uint64           `json:"Uint64Value,omitempty"`
	FloatValue    *float32          `json:"FloatValue,omitempty"`
	DoubleValue   *float64          `json:"DoubleValue,omitempty"`
	BooleanValue  *bool             `json:"BooleanValue,omitempty"`
	StringValue   *string           `json:"StringValue,omitempty"`
	BytesValue    []byte            `json:"BytesValue,omitempty"`
	ResponseValue *Payload_Response `json:"ResponseValue,omitempty"`
}

type MetricIntermediate struct {
	Name      string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`                                                     // @gotags: yaml:"name"
	DataType  DataType           `protobuf:"varint,2,opt,name=data_type,json=dataType,proto3,enum=beeplug.DataType" json:"data_type,omitempty" yaml:"data_type"` // @gotags: yaml:"data_type"
	Timestamp uint64             `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty" yaml:"timestamp"`                                     // @gotags: yaml:"timestamp"
	IsNull    bool               `protobuf:"varint,4,opt,name=is_null,json=isNull,proto3" json:"is_null,omitempty" yaml:"is_null"`                               // @gotags: yaml:"is_null"
	Status    string             `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty" yaml:"status"`                                               // @gotags: yaml:"status"
	Value     *ValueIntermediate `json:"Value,omitempty"`
}

func (m *Payload_Metric) UnmarshalJSON(data []byte) error {
	var inter MetricIntermediate
	if err := json.Unmarshal(data, &inter); err != nil {
		return err
	}

	m.Name = inter.Name
	m.DataType = inter.DataType
	m.Timestamp = inter.Timestamp
	m.IsNull = inter.IsNull
	m.Status = inter.Status

	if inter.Value != nil {
		switch m.DataType {
		case DataType_Int8, DataType_Int16, DataType_Int32:
			if inter.Value.Int32Value != nil {
				m.Value = &Payload_Metric_Int32Value{Int32Value: *inter.Value.Int32Value}
			}
		case DataType_Int64:
			if inter.Value.Int64Value != nil {
				m.Value = &Payload_Metric_Int64Value{Int64Value: *inter.Value.Int64Value}
			}
		case DataType_UInt8, DataType_UInt16, DataType_UInt32:
			if inter.Value.Uint32Value != nil {
				m.Value = &Payload_Metric_Uint32Value{Uint32Value: *inter.Value.Uint32Value}
			}
		case DataType_UInt64:
			if inter.Value.Uint64Value != nil {
				m.Value = &Payload_Metric_Uint64Value{Uint64Value: *inter.Value.Uint64Value}
			}
		case DataType_Float:
			if inter.Value.FloatValue != nil {
				m.Value = &Payload_Metric_FloatValue{FloatValue: *inter.Value.FloatValue}
			}
		case DataType_Double:
			if inter.Value.DoubleValue != nil {
				m.Value = &Payload_Metric_DoubleValue{DoubleValue: *inter.Value.DoubleValue}
			}
		case DataType_Boolean:
			if inter.Value.BooleanValue != nil {
				m.Value = &Payload_Metric_BooleanValue{BooleanValue: *inter.Value.BooleanValue}
			}
		case DataType_String:
			if inter.Value.StringValue != nil {
				m.Value = &Payload_Metric_StringValue{StringValue: *inter.Value.StringValue}
			}
		case DataType_Bytes:
			m.Value = &Payload_Metric_BytesValue{BytesValue: inter.Value.BytesValue}
		case DataType_Response:
			if inter.Value.ResponseValue != nil {
				m.Value = &Payload_Metric_ResponseValue{ResponseValue: inter.Value.ResponseValue}
			}
		}
	}

	return nil
}
