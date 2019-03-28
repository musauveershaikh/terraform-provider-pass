// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/model_evaluation.proto

package automl // import "google.golang.org/genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Evaluation results of a model.
type ModelEvaluation struct {
	// Output only. Problem type specific evaluation metrics.
	//
	// Types that are valid to be assigned to Metrics:
	//	*ModelEvaluation_ClassificationEvaluationMetrics
	//	*ModelEvaluation_TranslationEvaluationMetrics
	Metrics isModelEvaluation_Metrics `protobuf_oneof:"metrics"`
	// Output only.
	// Resource name of the model evaluation.
	// Format:
	//
	// `projects/{project_id}/locations/{location_id}/models/{model_id}/modelEvaluations/{model_evaluation_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only.
	// The ID of the annotation spec that the model evaluation applies to. The
	// ID is empty for overall model evaluation.
	// NOTE: Currently there is no way to obtain the display_name of the
	// annotation spec from its ID. To see the display_names, review the model
	// evaluations in the UI.
	AnnotationSpecId string `protobuf:"bytes,2,opt,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
	// Output only.
	// Timestamp when this model evaluation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The number of examples used for model evaluation.
	EvaluatedExampleCount int32    `protobuf:"varint,6,opt,name=evaluated_example_count,json=evaluatedExampleCount,proto3" json:"evaluated_example_count,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ModelEvaluation) Reset()         { *m = ModelEvaluation{} }
func (m *ModelEvaluation) String() string { return proto.CompactTextString(m) }
func (*ModelEvaluation) ProtoMessage()    {}
func (*ModelEvaluation) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_evaluation_1263308439036312, []int{0}
}
func (m *ModelEvaluation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModelEvaluation.Unmarshal(m, b)
}
func (m *ModelEvaluation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModelEvaluation.Marshal(b, m, deterministic)
}
func (dst *ModelEvaluation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModelEvaluation.Merge(dst, src)
}
func (m *ModelEvaluation) XXX_Size() int {
	return xxx_messageInfo_ModelEvaluation.Size(m)
}
func (m *ModelEvaluation) XXX_DiscardUnknown() {
	xxx_messageInfo_ModelEvaluation.DiscardUnknown(m)
}

var xxx_messageInfo_ModelEvaluation proto.InternalMessageInfo

type isModelEvaluation_Metrics interface {
	isModelEvaluation_Metrics()
}

type ModelEvaluation_ClassificationEvaluationMetrics struct {
	ClassificationEvaluationMetrics *ClassificationEvaluationMetrics `protobuf:"bytes,8,opt,name=classification_evaluation_metrics,json=classificationEvaluationMetrics,proto3,oneof"`
}

type ModelEvaluation_TranslationEvaluationMetrics struct {
	TranslationEvaluationMetrics *TranslationEvaluationMetrics `protobuf:"bytes,9,opt,name=translation_evaluation_metrics,json=translationEvaluationMetrics,proto3,oneof"`
}

func (*ModelEvaluation_ClassificationEvaluationMetrics) isModelEvaluation_Metrics() {}

func (*ModelEvaluation_TranslationEvaluationMetrics) isModelEvaluation_Metrics() {}

func (m *ModelEvaluation) GetMetrics() isModelEvaluation_Metrics {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *ModelEvaluation) GetClassificationEvaluationMetrics() *ClassificationEvaluationMetrics {
	if x, ok := m.GetMetrics().(*ModelEvaluation_ClassificationEvaluationMetrics); ok {
		return x.ClassificationEvaluationMetrics
	}
	return nil
}

func (m *ModelEvaluation) GetTranslationEvaluationMetrics() *TranslationEvaluationMetrics {
	if x, ok := m.GetMetrics().(*ModelEvaluation_TranslationEvaluationMetrics); ok {
		return x.TranslationEvaluationMetrics
	}
	return nil
}

func (m *ModelEvaluation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ModelEvaluation) GetAnnotationSpecId() string {
	if m != nil {
		return m.AnnotationSpecId
	}
	return ""
}

func (m *ModelEvaluation) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *ModelEvaluation) GetEvaluatedExampleCount() int32 {
	if m != nil {
		return m.EvaluatedExampleCount
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ModelEvaluation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ModelEvaluation_OneofMarshaler, _ModelEvaluation_OneofUnmarshaler, _ModelEvaluation_OneofSizer, []interface{}{
		(*ModelEvaluation_ClassificationEvaluationMetrics)(nil),
		(*ModelEvaluation_TranslationEvaluationMetrics)(nil),
	}
}

func _ModelEvaluation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ModelEvaluation)
	// metrics
	switch x := m.Metrics.(type) {
	case *ModelEvaluation_ClassificationEvaluationMetrics:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClassificationEvaluationMetrics); err != nil {
			return err
		}
	case *ModelEvaluation_TranslationEvaluationMetrics:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TranslationEvaluationMetrics); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ModelEvaluation.Metrics has unexpected type %T", x)
	}
	return nil
}

func _ModelEvaluation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ModelEvaluation)
	switch tag {
	case 8: // metrics.classification_evaluation_metrics
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClassificationEvaluationMetrics)
		err := b.DecodeMessage(msg)
		m.Metrics = &ModelEvaluation_ClassificationEvaluationMetrics{msg}
		return true, err
	case 9: // metrics.translation_evaluation_metrics
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TranslationEvaluationMetrics)
		err := b.DecodeMessage(msg)
		m.Metrics = &ModelEvaluation_TranslationEvaluationMetrics{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ModelEvaluation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ModelEvaluation)
	// metrics
	switch x := m.Metrics.(type) {
	case *ModelEvaluation_ClassificationEvaluationMetrics:
		s := proto.Size(x.ClassificationEvaluationMetrics)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ModelEvaluation_TranslationEvaluationMetrics:
		s := proto.Size(x.TranslationEvaluationMetrics)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ModelEvaluation)(nil), "google.cloud.automl.v1beta1.ModelEvaluation")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/model_evaluation.proto", fileDescriptor_model_evaluation_1263308439036312)
}

var fileDescriptor_model_evaluation_1263308439036312 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x8b, 0x13, 0x31,
	0x18, 0x36, 0xab, 0xbb, 0xda, 0xec, 0x41, 0x09, 0x88, 0xc3, 0xec, 0x62, 0xab, 0xa7, 0x1e, 0x34,
	0x71, 0x2b, 0x08, 0x52, 0x2f, 0x6d, 0x29, 0xea, 0xa1, 0x20, 0x63, 0xf1, 0x20, 0x85, 0x21, 0xcd,
	0xa4, 0x43, 0x20, 0x1f, 0xc3, 0x24, 0x53, 0xbc, 0x0a, 0x9e, 0xfc, 0x69, 0xfe, 0x0e, 0x7f, 0xc8,
	0x32, 0xc9, 0xb4, 0x33, 0x85, 0x32, 0xb7, 0xc9, 0x3c, 0x1f, 0xef, 0xf3, 0xe4, 0x0d, 0x9c, 0xe4,
	0xc6, 0xe4, 0x92, 0x13, 0x26, 0x4d, 0x95, 0x11, 0x5a, 0x39, 0xa3, 0x24, 0xd9, 0xdf, 0x6d, 0xb9,
	0xa3, 0x77, 0x44, 0x99, 0x8c, 0xcb, 0x94, 0xef, 0xa9, 0xac, 0xa8, 0x13, 0x46, 0xe3, 0xa2, 0x34,
	0xce, 0xa0, 0x9b, 0xa0, 0xc1, 0x5e, 0x83, 0x83, 0x06, 0x37, 0x9a, 0xf8, 0xb6, 0x31, 0xa4, 0x85,
	0x20, 0x54, 0x6b, 0xe3, 0xbc, 0xd2, 0x06, 0x69, 0xfc, 0xae, 0x6f, 0x1c, 0x93, 0xd4, 0x5a, 0xb1,
	0x13, 0xac, 0x33, 0x2c, 0x7e, 0xdb, 0xa7, 0x70, 0x25, 0xd5, 0x56, 0x76, 0xe9, 0xc3, 0x86, 0xee,
	0x4f, 0xdb, 0x6a, 0x47, 0x9c, 0x50, 0xdc, 0x3a, 0xaa, 0x8a, 0x40, 0x78, 0xfd, 0xff, 0x21, 0x7c,
	0xba, 0xaa, 0x7b, 0x2d, 0x8f, 0xb5, 0xd0, 0x5f, 0x00, 0x5f, 0x9d, 0x0e, 0xef, 0x94, 0x4e, 0x15,
	0x77, 0xa5, 0x60, 0x36, 0x7a, 0x32, 0x02, 0xe3, 0xeb, 0xc9, 0x27, 0xdc, 0xd3, 0x1e, 0x2f, 0x4e,
	0x5c, 0xda, 0x11, 0xab, 0xe0, 0xf1, 0xe5, 0x41, 0x32, 0x64, 0xfd, 0x14, 0xf4, 0x1b, 0xc0, 0x97,
	0x9d, 0x5e, 0xe7, 0x92, 0x0c, 0x7c, 0x92, 0x8f, 0xbd, 0x49, 0xd6, 0xad, 0xc5, 0xb9, 0x18, 0xb7,
	0xae, 0x07, 0x47, 0x08, 0x3e, 0xd2, 0x54, 0xf1, 0x08, 0x8c, 0xc0, 0x78, 0x90, 0xf8, 0x6f, 0xf4,
	0x06, 0xa2, 0x76, 0x9f, 0xa9, 0x2d, 0x38, 0x4b, 0x45, 0x16, 0x5d, 0x78, 0xc6, 0xb3, 0x16, 0xf9,
	0x5e, 0x70, 0xf6, 0x35, 0x43, 0x53, 0x78, 0xcd, 0x4a, 0x4e, 0x1d, 0x4f, 0xeb, 0x05, 0x44, 0x97,
	0x3e, 0x71, 0x7c, 0x48, 0x7c, 0xd8, 0x0e, 0x5e, 0x1f, 0xb6, 0x93, 0xc0, 0x40, 0xaf, 0x7f, 0xa0,
	0x0f, 0xf0, 0x45, 0xd3, 0x9a, 0x67, 0x29, 0xff, 0x45, 0x55, 0x21, 0x79, 0xca, 0x4c, 0xa5, 0x5d,
	0x74, 0x35, 0x02, 0xe3, 0xcb, 0xe4, 0xf9, 0x11, 0x5e, 0x06, 0x74, 0x51, 0x83, 0xf3, 0x01, 0x7c,
	0xdc, 0x5c, 0xd1, 0xfc, 0x0f, 0x80, 0x43, 0x66, 0x54, 0xdf, 0x15, 0x7d, 0x03, 0x3f, 0x67, 0x0d,
	0x9c, 0x1b, 0x49, 0x75, 0x8e, 0x4d, 0x99, 0x93, 0x9c, 0x6b, 0x9f, 0x8e, 0x04, 0x88, 0x16, 0xc2,
	0x9e, 0x7d, 0x7b, 0xd3, 0x70, 0xfc, 0x77, 0x71, 0xf3, 0xd9, 0x13, 0x37, 0x8b, 0x9a, 0xb4, 0x99,
	0x55, 0xce, 0xac, 0xe4, 0xe6, 0x47, 0x20, 0x6d, 0xaf, 0xbc, 0xd7, 0xfb, 0xfb, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5e, 0x71, 0x31, 0xe7, 0x67, 0x03, 0x00, 0x00,
}
