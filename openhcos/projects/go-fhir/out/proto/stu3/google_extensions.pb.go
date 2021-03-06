// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/stu3/google_extensions.proto

package google_fhir_stu3_google

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Auto-generated from StructureDefinition for separator-stride.
// Base64Binary rendering parameters.
// See https://g.co/fhir/StructureDefinition/base64Binary-separatorStride
type Base64BinarySeparatorStride struct {
	// xml:id (or equivalent in JSON)
	Id *String `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []*Extension `protobuf:"bytes,2,rep,name=extension,proto3" json:"extension,omitempty"`
	// The separator, usually a sequence of spaces.
	Separator *String `protobuf:"bytes,4,opt,name=separator,proto3" json:"separator,omitempty"`
	// The stride.
	Stride               *PositiveInt `protobuf:"bytes,5,opt,name=stride,proto3" json:"stride,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Base64BinarySeparatorStride) Reset()         { *m = Base64BinarySeparatorStride{} }
func (m *Base64BinarySeparatorStride) String() string { return proto.CompactTextString(m) }
func (*Base64BinarySeparatorStride) ProtoMessage()    {}
func (*Base64BinarySeparatorStride) Descriptor() ([]byte, []int) {
	return fileDescriptor_576ce3ea5cc707ed, []int{0}
}

func (m *Base64BinarySeparatorStride) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Base64BinarySeparatorStride.Unmarshal(m, b)
}
func (m *Base64BinarySeparatorStride) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Base64BinarySeparatorStride.Marshal(b, m, deterministic)
}
func (m *Base64BinarySeparatorStride) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Base64BinarySeparatorStride.Merge(m, src)
}
func (m *Base64BinarySeparatorStride) XXX_Size() int {
	return xxx_messageInfo_Base64BinarySeparatorStride.Size(m)
}
func (m *Base64BinarySeparatorStride) XXX_DiscardUnknown() {
	xxx_messageInfo_Base64BinarySeparatorStride.DiscardUnknown(m)
}

var xxx_messageInfo_Base64BinarySeparatorStride proto.InternalMessageInfo

func (m *Base64BinarySeparatorStride) GetId() *String {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Base64BinarySeparatorStride) GetExtension() []*Extension {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *Base64BinarySeparatorStride) GetSeparator() *String {
	if m != nil {
		return m.Separator
	}
	return nil
}

func (m *Base64BinarySeparatorStride) GetStride() *PositiveInt {
	if m != nil {
		return m.Stride
	}
	return nil
}

// Auto-generated from StructureDefinition for eventLabel.
// EventLabels define labels used for TensorFlow model training and evaluation.
// See https://g.co/fhir/StructureDefinition/eventLabel
type EventLabel struct {
	// xml:id (or equivalent in JSON)
	Id *String `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []*Extension `protobuf:"bytes,2,rep,name=extension,proto3" json:"extension,omitempty"`
	// The patient associated with this label
	Patient *Reference `protobuf:"bytes,4,opt,name=patient,proto3" json:"patient,omitempty"`
	// The type of label, part of the prediction task definition
	Type *Coding `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	// Time associated with the label event, if available
	EventTime *DateTime `protobuf:"bytes,6,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// Resource that owns this label
	Source               *Reference          `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	Label                []*EventLabel_Label `protobuf:"bytes,8,rep,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *EventLabel) Reset()         { *m = EventLabel{} }
func (m *EventLabel) String() string { return proto.CompactTextString(m) }
func (*EventLabel) ProtoMessage()    {}
func (*EventLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_576ce3ea5cc707ed, []int{1}
}

func (m *EventLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventLabel.Unmarshal(m, b)
}
func (m *EventLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventLabel.Marshal(b, m, deterministic)
}
func (m *EventLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventLabel.Merge(m, src)
}
func (m *EventLabel) XXX_Size() int {
	return xxx_messageInfo_EventLabel.Size(m)
}
func (m *EventLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_EventLabel.DiscardUnknown(m)
}

var xxx_messageInfo_EventLabel proto.InternalMessageInfo

func (m *EventLabel) GetId() *String {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *EventLabel) GetExtension() []*Extension {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *EventLabel) GetPatient() *Reference {
	if m != nil {
		return m.Patient
	}
	return nil
}

func (m *EventLabel) GetType() *Coding {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *EventLabel) GetEventTime() *DateTime {
	if m != nil {
		return m.EventTime
	}
	return nil
}

func (m *EventLabel) GetSource() *Reference {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *EventLabel) GetLabel() []*EventLabel_Label {
	if m != nil {
		return m.Label
	}
	return nil
}

// List of labels associated with this event
type EventLabel_Label struct {
	// xml:id (or equivalent in JSON)
	Id *String `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Class name in multi-class prediction tasks, e.g. code "780.60" for icd9
	ClassName            *Coding                      `protobuf:"bytes,4,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`
	ClassValue           *EventLabel_Label_ClassValue `protobuf:"bytes,5,opt,name=class_value,json=classValue,proto3" json:"class_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *EventLabel_Label) Reset()         { *m = EventLabel_Label{} }
func (m *EventLabel_Label) String() string { return proto.CompactTextString(m) }
func (*EventLabel_Label) ProtoMessage()    {}
func (*EventLabel_Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_576ce3ea5cc707ed, []int{1, 0}
}

func (m *EventLabel_Label) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventLabel_Label.Unmarshal(m, b)
}
func (m *EventLabel_Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventLabel_Label.Marshal(b, m, deterministic)
}
func (m *EventLabel_Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventLabel_Label.Merge(m, src)
}
func (m *EventLabel_Label) XXX_Size() int {
	return xxx_messageInfo_EventLabel_Label.Size(m)
}
func (m *EventLabel_Label) XXX_DiscardUnknown() {
	xxx_messageInfo_EventLabel_Label.DiscardUnknown(m)
}

var xxx_messageInfo_EventLabel_Label proto.InternalMessageInfo

func (m *EventLabel_Label) GetId() *String {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *EventLabel_Label) GetClassName() *Coding {
	if m != nil {
		return m.ClassName
	}
	return nil
}

func (m *EventLabel_Label) GetClassValue() *EventLabel_Label_ClassValue {
	if m != nil {
		return m.ClassValue
	}
	return nil
}

// The value of the label
type EventLabel_Label_ClassValue struct {
	// Types that are valid to be assigned to ClassValue:
	//	*EventLabel_Label_ClassValue_Boolean
	//	*EventLabel_Label_ClassValue_Decimal
	//	*EventLabel_Label_ClassValue_Integer
	//	*EventLabel_Label_ClassValue_StringValue
	ClassValue           isEventLabel_Label_ClassValue_ClassValue `protobuf_oneof:"class_value"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *EventLabel_Label_ClassValue) Reset()         { *m = EventLabel_Label_ClassValue{} }
func (m *EventLabel_Label_ClassValue) String() string { return proto.CompactTextString(m) }
func (*EventLabel_Label_ClassValue) ProtoMessage()    {}
func (*EventLabel_Label_ClassValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_576ce3ea5cc707ed, []int{1, 0, 0}
}

func (m *EventLabel_Label_ClassValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventLabel_Label_ClassValue.Unmarshal(m, b)
}
func (m *EventLabel_Label_ClassValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventLabel_Label_ClassValue.Marshal(b, m, deterministic)
}
func (m *EventLabel_Label_ClassValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventLabel_Label_ClassValue.Merge(m, src)
}
func (m *EventLabel_Label_ClassValue) XXX_Size() int {
	return xxx_messageInfo_EventLabel_Label_ClassValue.Size(m)
}
func (m *EventLabel_Label_ClassValue) XXX_DiscardUnknown() {
	xxx_messageInfo_EventLabel_Label_ClassValue.DiscardUnknown(m)
}

var xxx_messageInfo_EventLabel_Label_ClassValue proto.InternalMessageInfo

type isEventLabel_Label_ClassValue_ClassValue interface {
	isEventLabel_Label_ClassValue_ClassValue()
}

type EventLabel_Label_ClassValue_Boolean struct {
	Boolean *Boolean `protobuf:"bytes,1,opt,name=boolean,proto3,oneof"`
}

type EventLabel_Label_ClassValue_Decimal struct {
	Decimal *Decimal `protobuf:"bytes,2,opt,name=decimal,proto3,oneof"`
}

type EventLabel_Label_ClassValue_Integer struct {
	Integer *Integer `protobuf:"bytes,3,opt,name=integer,proto3,oneof"`
}

type EventLabel_Label_ClassValue_StringValue struct {
	StringValue *String `protobuf:"bytes,4,opt,name=string_value,json=string,proto3,oneof"`
}

func (*EventLabel_Label_ClassValue_Boolean) isEventLabel_Label_ClassValue_ClassValue() {}

func (*EventLabel_Label_ClassValue_Decimal) isEventLabel_Label_ClassValue_ClassValue() {}

func (*EventLabel_Label_ClassValue_Integer) isEventLabel_Label_ClassValue_ClassValue() {}

func (*EventLabel_Label_ClassValue_StringValue) isEventLabel_Label_ClassValue_ClassValue() {}

func (m *EventLabel_Label_ClassValue) GetClassValue() isEventLabel_Label_ClassValue_ClassValue {
	if m != nil {
		return m.ClassValue
	}
	return nil
}

func (m *EventLabel_Label_ClassValue) GetBoolean() *Boolean {
	if x, ok := m.GetClassValue().(*EventLabel_Label_ClassValue_Boolean); ok {
		return x.Boolean
	}
	return nil
}

func (m *EventLabel_Label_ClassValue) GetDecimal() *Decimal {
	if x, ok := m.GetClassValue().(*EventLabel_Label_ClassValue_Decimal); ok {
		return x.Decimal
	}
	return nil
}

func (m *EventLabel_Label_ClassValue) GetInteger() *Integer {
	if x, ok := m.GetClassValue().(*EventLabel_Label_ClassValue_Integer); ok {
		return x.Integer
	}
	return nil
}

func (m *EventLabel_Label_ClassValue) GetStringValue() *String {
	if x, ok := m.GetClassValue().(*EventLabel_Label_ClassValue_StringValue); ok {
		return x.StringValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*EventLabel_Label_ClassValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _EventLabel_Label_ClassValue_OneofMarshaler, _EventLabel_Label_ClassValue_OneofUnmarshaler, _EventLabel_Label_ClassValue_OneofSizer, []interface{}{
		(*EventLabel_Label_ClassValue_Boolean)(nil),
		(*EventLabel_Label_ClassValue_Decimal)(nil),
		(*EventLabel_Label_ClassValue_Integer)(nil),
		(*EventLabel_Label_ClassValue_StringValue)(nil),
	}
}

func _EventLabel_Label_ClassValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*EventLabel_Label_ClassValue)
	// class_value
	switch x := m.ClassValue.(type) {
	case *EventLabel_Label_ClassValue_Boolean:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Boolean); err != nil {
			return err
		}
	case *EventLabel_Label_ClassValue_Decimal:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Decimal); err != nil {
			return err
		}
	case *EventLabel_Label_ClassValue_Integer:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Integer); err != nil {
			return err
		}
	case *EventLabel_Label_ClassValue_StringValue:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.StringValue); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("EventLabel_Label_ClassValue.ClassValue has unexpected type %T", x)
	}
	return nil
}

func _EventLabel_Label_ClassValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*EventLabel_Label_ClassValue)
	switch tag {
	case 1: // class_value.boolean
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Boolean)
		err := b.DecodeMessage(msg)
		m.ClassValue = &EventLabel_Label_ClassValue_Boolean{msg}
		return true, err
	case 2: // class_value.decimal
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Decimal)
		err := b.DecodeMessage(msg)
		m.ClassValue = &EventLabel_Label_ClassValue_Decimal{msg}
		return true, err
	case 3: // class_value.integer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Integer)
		err := b.DecodeMessage(msg)
		m.ClassValue = &EventLabel_Label_ClassValue_Integer{msg}
		return true, err
	case 4: // class_value.string_value
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(String)
		err := b.DecodeMessage(msg)
		m.ClassValue = &EventLabel_Label_ClassValue_StringValue{msg}
		return true, err
	default:
		return false, nil
	}
}

func _EventLabel_Label_ClassValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*EventLabel_Label_ClassValue)
	// class_value
	switch x := m.ClassValue.(type) {
	case *EventLabel_Label_ClassValue_Boolean:
		s := proto.Size(x.Boolean)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventLabel_Label_ClassValue_Decimal:
		s := proto.Size(x.Decimal)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventLabel_Label_ClassValue_Integer:
		s := proto.Size(x.Integer)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *EventLabel_Label_ClassValue_StringValue:
		s := proto.Size(x.StringValue)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Auto-generated from StructureDefinition for eventTrigger.
// EventTriggers specify cutoff times for generated TensorFlow examples.
// See https://g.co/fhir/StructureDefinition/eventTrigger
type EventTrigger struct {
	// xml:id (or equivalent in JSON)
	Id *String `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Additional content defined by implementations
	Extension []*Extension `protobuf:"bytes,2,rep,name=extension,proto3" json:"extension,omitempty"`
	// The type of trigger, part of the prediction task definition.
	Type *Coding `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	// Prediction event time, more recent data will be elided.
	EventTime *DateTime `protobuf:"bytes,5,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// Resource that owns this trigger.
	Source               *Reference `protobuf:"bytes,6,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *EventTrigger) Reset()         { *m = EventTrigger{} }
func (m *EventTrigger) String() string { return proto.CompactTextString(m) }
func (*EventTrigger) ProtoMessage()    {}
func (*EventTrigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_576ce3ea5cc707ed, []int{2}
}

func (m *EventTrigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventTrigger.Unmarshal(m, b)
}
func (m *EventTrigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventTrigger.Marshal(b, m, deterministic)
}
func (m *EventTrigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventTrigger.Merge(m, src)
}
func (m *EventTrigger) XXX_Size() int {
	return xxx_messageInfo_EventTrigger.Size(m)
}
func (m *EventTrigger) XXX_DiscardUnknown() {
	xxx_messageInfo_EventTrigger.DiscardUnknown(m)
}

var xxx_messageInfo_EventTrigger proto.InternalMessageInfo

func (m *EventTrigger) GetId() *String {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *EventTrigger) GetExtension() []*Extension {
	if m != nil {
		return m.Extension
	}
	return nil
}

func (m *EventTrigger) GetType() *Coding {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *EventTrigger) GetEventTime() *DateTime {
	if m != nil {
		return m.EventTime
	}
	return nil
}

func (m *EventTrigger) GetSource() *Reference {
	if m != nil {
		return m.Source
	}
	return nil
}

// Auto-generated from StructureDefinition for primitive-has-no-value.
// Is the primitive missing a value?.
// See https://g.co/fhir/StructureDefinition/primitiveHasNoValue
type PrimitiveHasNoValue struct {
	// xml:id (or equivalent in JSON)
	Id *String `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Value of extension
	ValueBoolean         *Boolean `protobuf:"bytes,3,opt,name=value_boolean,json=valueBoolean,proto3" json:"value_boolean,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimitiveHasNoValue) Reset()         { *m = PrimitiveHasNoValue{} }
func (m *PrimitiveHasNoValue) String() string { return proto.CompactTextString(m) }
func (*PrimitiveHasNoValue) ProtoMessage()    {}
func (*PrimitiveHasNoValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_576ce3ea5cc707ed, []int{3}
}

func (m *PrimitiveHasNoValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimitiveHasNoValue.Unmarshal(m, b)
}
func (m *PrimitiveHasNoValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimitiveHasNoValue.Marshal(b, m, deterministic)
}
func (m *PrimitiveHasNoValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimitiveHasNoValue.Merge(m, src)
}
func (m *PrimitiveHasNoValue) XXX_Size() int {
	return xxx_messageInfo_PrimitiveHasNoValue.Size(m)
}
func (m *PrimitiveHasNoValue) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimitiveHasNoValue.DiscardUnknown(m)
}

var xxx_messageInfo_PrimitiveHasNoValue proto.InternalMessageInfo

func (m *PrimitiveHasNoValue) GetId() *String {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PrimitiveHasNoValue) GetValueBoolean() *Boolean {
	if m != nil {
		return m.ValueBoolean
	}
	return nil
}

func init() {
	proto.RegisterType((*Base64BinarySeparatorStride)(nil), "google.fhir.stu3.google.Base64BinarySeparatorStride")
	proto.RegisterType((*EventLabel)(nil), "google.fhir.stu3.google.EventLabel")
	proto.RegisterType((*EventLabel_Label)(nil), "google.fhir.stu3.google.EventLabel.Label")
	proto.RegisterType((*EventLabel_Label_ClassValue)(nil), "google.fhir.stu3.google.EventLabel.Label.ClassValue")
	proto.RegisterType((*EventTrigger)(nil), "google.fhir.stu3.google.EventTrigger")
	proto.RegisterType((*PrimitiveHasNoValue)(nil), "google.fhir.stu3.google.PrimitiveHasNoValue")
}

func init() { proto.RegisterFile("proto/stu3/google_extensions.proto", fileDescriptor_576ce3ea5cc707ed) }

var fileDescriptor_576ce3ea5cc707ed = []byte{
	// 718 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xbf, 0x6f, 0xd3, 0x4c,
	0x18, 0xc7, 0x9b, 0x5f, 0x6e, 0x73, 0x4d, 0x17, 0xbf, 0xc3, 0x6b, 0xa5, 0xaf, 0xde, 0x96, 0xb0,
	0x94, 0x01, 0x1b, 0xda, 0x0a, 0xd4, 0x20, 0x51, 0xc9, 0x49, 0x45, 0x2b, 0xa1, 0xaa, 0x72, 0x2b,
	0xd6, 0xe8, 0xe2, 0x3c, 0x75, 0x4f, 0xb2, 0xef, 0xa2, 0xbb, 0x4b, 0x45, 0x25, 0x56, 0x84, 0x58,
	0x18, 0x60, 0x40, 0x62, 0x41, 0xcc, 0x6c, 0x48, 0x74, 0x66, 0x65, 0xe3, 0x5f, 0x40, 0x65, 0x63,
	0x61, 0xec, 0x84, 0xd0, 0x9d, 0xed, 0x24, 0xa5, 0xb8, 0xb8, 0xad, 0xd4, 0x25, 0x4a, 0x72, 0xcf,
	0xe7, 0xf1, 0xf3, 0x3c, 0xdf, 0xfb, 0x3e, 0x46, 0x8d, 0x3e, 0x67, 0x92, 0x39, 0x42, 0x0e, 0x96,
	0x9c, 0x80, 0xb1, 0x20, 0x84, 0x0e, 0x3c, 0x96, 0x40, 0x05, 0x61, 0x54, 0xd8, 0xfa, 0xd0, 0xfc,
	0x37, 0x3e, 0xb0, 0x77, 0xf7, 0x08, 0xb7, 0x55, 0xa4, 0x1d, 0xff, 0x51, 0xff, 0x6f, 0x0c, 0xc6,
	0x94, 0x32, 0x89, 0xe5, 0x08, 0xab, 0xd7, 0xc7, 0x4e, 0x7b, 0x58, 0x62, 0x79, 0xd0, 0x87, 0xe4,
	0xac, 0xf1, 0xae, 0x84, 0x66, 0x5d, 0x2c, 0xe0, 0xce, 0xb2, 0x4b, 0x28, 0xe6, 0x07, 0xdb, 0xd0,
	0xc7, 0x1c, 0x4b, 0xc6, 0xb7, 0x25, 0x27, 0x3d, 0x30, 0x6d, 0x54, 0x24, 0x3d, 0xab, 0x30, 0x5f,
	0x58, 0x98, 0x5e, 0xfc, 0xdf, 0x3e, 0xf5, 0x7c, 0x9d, 0xc4, 0x56, 0xb1, 0x34, 0xf0, 0x8a, 0xa4,
	0x67, 0xae, 0xa2, 0xea, 0xb0, 0x6c, 0xab, 0x38, 0x5f, 0x5a, 0x98, 0x5e, 0xbc, 0x96, 0x85, 0xad,
	0xa5, 0x81, 0xde, 0x88, 0x31, 0xdb, 0xa8, 0x2a, 0xd2, 0x1a, 0xac, 0x72, 0x9e, 0xe7, 0xba, 0xc6,
	0x8f, 0x2f, 0xcf, 0xbe, 0x97, 0x0b, 0xde, 0x08, 0x34, 0x5b, 0xc8, 0x10, 0xba, 0x01, 0xab, 0xa2,
	0x53, 0x5c, 0xcf, 0x4a, 0xb1, 0xc5, 0x04, 0x91, 0x64, 0x1f, 0x36, 0xa8, 0x1c, 0xe6, 0x49, 0xd0,
	0xe6, 0xd3, 0xc2, 0xa7, 0xb7, 0x47, 0x87, 0x95, 0xe2, 0x9b, 0x8f, 0x2f, 0x5e, 0x19, 0xb7, 0xf7,
	0xa4, 0xec, 0x37, 0x1d, 0x67, 0x2f, 0xbc, 0x6b, 0x33, 0x1e, 0x38, 0x2a, 0x91, 0xb3, 0x2d, 0xf9,
	0xc0, 0x97, 0x03, 0x0e, 0x6d, 0xd8, 0x25, 0x94, 0xa8, 0xa1, 0x3b, 0xc3, 0xae, 0x3e, 0xfc, 0xfc,
	0xf6, 0xda, 0x70, 0x15, 0x26, 0x9a, 0x8e, 0x13, 0xd8, 0x3e, 0xcb, 0x86, 0xba, 0x63, 0x12, 0xdc,
	0x14, 0x27, 0x35, 0x68, 0xbc, 0x9f, 0x42, 0x68, 0x6d, 0x1f, 0xa8, 0x7c, 0x88, 0xbb, 0x10, 0x5e,
	0xbd, 0x24, 0x2d, 0x34, 0xd9, 0xc7, 0x92, 0x00, 0x95, 0x89, 0x20, 0x99, 0xb8, 0x07, 0xbb, 0xc0,
	0x81, 0xfa, 0x30, 0x9c, 0x65, 0x4a, 0x9a, 0x4d, 0x54, 0x56, 0xf7, 0x2e, 0xd1, 0x23, 0xb3, 0xee,
	0x16, 0xeb, 0x8d, 0x4b, 0xaa, 0x19, 0x73, 0x15, 0x21, 0x50, 0xfd, 0x77, 0x24, 0x89, 0xc0, 0x32,
	0x74, 0x86, 0xf9, 0xac, 0x0c, 0x6d, 0x2c, 0x61, 0x87, 0x44, 0xe0, 0x55, 0x35, 0xa3, 0xbe, 0x9a,
	0x2b, 0xc8, 0x10, 0x6c, 0xc0, 0x7d, 0xb0, 0x26, 0x73, 0x36, 0xe0, 0x25, 0x80, 0xb9, 0x8a, 0x2a,
	0xa1, 0x1a, 0xbb, 0x35, 0xa5, 0x27, 0x77, 0xc3, 0xce, 0xf0, 0xa0, 0x3d, 0x52, 0xc8, 0xd6, 0x9f,
	0x5e, 0xcc, 0xd5, 0x8f, 0x4b, 0xa8, 0x72, 0x31, 0xe1, 0xd6, 0x10, 0xf2, 0x43, 0x2c, 0x44, 0x87,
	0xe2, 0x08, 0xfe, 0xe6, 0x85, 0xdf, 0x06, 0x57, 0xd5, 0xe4, 0x26, 0x8e, 0xc0, 0xf4, 0xd1, 0x74,
	0x9c, 0x66, 0x1f, 0x87, 0x83, 0x54, 0x80, 0xe5, 0xdc, 0x7d, 0xd8, 0x2d, 0x05, 0x3f, 0x52, 0xac,
	0x5b, 0x8b, 0xb3, 0x1f, 0x7f, 0x3d, 0xfa, 0x5c, 0x29, 0x78, 0x71, 0x75, 0xfa, 0xa4, 0xfe, 0xb2,
	0x88, 0xd0, 0x28, 0xd0, 0xbc, 0x87, 0x26, 0xbb, 0x8c, 0x85, 0x80, 0x69, 0xd2, 0xef, 0x5c, 0x56,
	0xdd, 0x6e, 0x1c, 0xb6, 0x3e, 0xe1, 0xa5, 0x84, 0x82, 0x7b, 0xe0, 0x93, 0x08, 0x87, 0x56, 0xf1,
	0x6c, 0xb8, 0x1d, 0x87, 0x29, 0x38, 0x21, 0x14, 0x4c, 0xa8, 0x84, 0x00, 0xb8, 0x55, 0x3a, 0x1b,
	0xde, 0x88, 0xc3, 0x14, 0x9c, 0x10, 0xe6, 0x7d, 0x54, 0x13, 0x7a, 0xfe, 0xc9, 0xac, 0x72, 0xed,
	0x9f, 0xf5, 0x89, 0x78, 0x63, 0xd0, 0xc0, 0x9d, 0x39, 0x31, 0xea, 0xa6, 0xb8, 0xfc, 0xfe, 0xb8,
	0x95, 0x6f, 0x7f, 0xc0, 0x50, 0xb4, 0xc6, 0x61, 0x09, 0xd5, 0xb4, 0x86, 0x3b, 0x9c, 0x04, 0xaa,
	0xa9, 0x2b, 0xdf, 0x17, 0xa9, 0xd5, 0xcb, 0x17, 0xb0, 0xfa, 0x83, 0x13, 0x56, 0xaf, 0xe4, 0xb3,
	0xfa, 0xe8, 0xd6, 0xff, 0xc9, 0xf2, 0xc6, 0x39, 0x2d, 0xdf, 0x1c, 0x5c, 0x5e, 0xb6, 0xc5, 0x73,
	0xc8, 0x96, 0xe8, 0xd4, 0x78, 0x5e, 0x44, 0xff, 0x6c, 0x71, 0x12, 0xe9, 0xf7, 0xd1, 0x3a, 0x16,
	0x9b, 0x2c, 0xf6, 0xd2, 0x79, 0xf5, 0x6b, 0xa3, 0x19, 0x7d, 0xfd, 0x3a, 0xa9, 0x03, 0x4b, 0xb9,
	0x1c, 0xe8, 0xd5, 0x34, 0x95, 0xfc, 0x6a, 0x3e, 0xb9, 0xfc, 0x10, 0x56, 0xf2, 0x0d, 0xa1, 0x7f,
	0xba, 0x67, 0x77, 0x0e, 0xcd, 0xfa, 0x2c, 0xca, 0xda, 0x51, 0x5b, 0x85, 0xae, 0xa1, 0x6b, 0x5f,
	0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0x30, 0x20, 0x66, 0x30, 0x37, 0x09, 0x00, 0x00,
}
