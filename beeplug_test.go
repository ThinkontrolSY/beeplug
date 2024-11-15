package beeplug

import (
	"encoding/json"
	"testing"
)

func TestMetricJSONSerializationDeserialization(t *testing.T) {
	// 创建不同类型的Metric实例用于测试
	testCases := []*Payload_Metric{
		{
			Name:      "example_metric_int32",
			DataType:  DataType_Int32,
			Timestamp: 1234567890,
			Value: &Payload_Metric_Int32Value{
				Int32Value: 42,
			},
		},
		{
			Name:      "example_metric_int64",
			DataType:  DataType_Int64,
			Timestamp: 1234567890,
			Value: &Payload_Metric_Int64Value{
				Int64Value: 4112,
			},
		},
		{
			Name:      "example_metric_string",
			DataType:  DataType_String,
			Timestamp: 1234567890,
			Value:     &Payload_Metric_StringValue{StringValue: "hello"},
		},
		{
			Name:      "example_metric_bool",
			DataType:  DataType_Boolean,
			Timestamp: 1234567890,
			Value:     &Payload_Metric_BooleanValue{BooleanValue: true},
		},
		{
			Name:      "example_metric_float",
			DataType:  DataType_Float,
			Timestamp: 1234567890,
			Value:     &Payload_Metric_FloatValue{FloatValue: 0.3},
		},
		{
			Name:      "example_metric_double",
			DataType:  DataType_Double,
			Timestamp: 1234567890,
			Value:     &Payload_Metric_DoubleValue{DoubleValue: 0.44},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			// 测试JSON序列化
			jsonData, err := json.Marshal(tc)
			if err != nil {
				t.Fatalf("Error serializing to JSON: %v", err)
			}

			t.Log(string(jsonData))

			// 测试JSON反序列化
			var deserialized Payload_Metric
			if err := json.Unmarshal(jsonData, &deserialized); err != nil {
				t.Fatalf("Error deserializing from JSON: %v", err)
			}

			t.Log(&deserialized)
		})
	}
}
