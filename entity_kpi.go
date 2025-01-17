package beeplug

type EntityKpi struct {
	Bucket string            `json:"bucket,omitempty" yaml:"bucket"`
	Name   string            `json:"name,omitempty" yaml:"name"`
	Fields map[string]any    `json:"fields,omitempty" yaml:"fields"`
	Tags   map[string]string `json:"tags,omitempty" yaml:"tags"`
}
