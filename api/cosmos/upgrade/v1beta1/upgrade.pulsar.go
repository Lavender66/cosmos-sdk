package upgradev1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Plan                       protoreflect.MessageDescriptor
	fd_Plan_name                  protoreflect.FieldDescriptor
	fd_Plan_time                  protoreflect.FieldDescriptor
	fd_Plan_height                protoreflect.FieldDescriptor
	fd_Plan_info                  protoreflect.FieldDescriptor
	fd_Plan_upgraded_client_state protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_upgrade_v1beta1_upgrade_proto_init()
	md_Plan = File_cosmos_upgrade_v1beta1_upgrade_proto.Messages().ByName("Plan")
	fd_Plan_name = md_Plan.Fields().ByName("name")
	fd_Plan_time = md_Plan.Fields().ByName("time")
	fd_Plan_height = md_Plan.Fields().ByName("height")
	fd_Plan_info = md_Plan.Fields().ByName("info")
	fd_Plan_upgraded_client_state = md_Plan.Fields().ByName("upgraded_client_state")
}

var _ protoreflect.Message = (*fastReflection_Plan)(nil)

type fastReflection_Plan Plan

func (x *Plan) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Plan)(x)
}

func (x *Plan) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Plan_messageType fastReflection_Plan_messageType
var _ protoreflect.MessageType = fastReflection_Plan_messageType{}

type fastReflection_Plan_messageType struct{}

func (x fastReflection_Plan_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Plan)(nil)
}
func (x fastReflection_Plan_messageType) New() protoreflect.Message {
	return new(fastReflection_Plan)
}
func (x fastReflection_Plan_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Plan
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Plan) Descriptor() protoreflect.MessageDescriptor {
	return md_Plan
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Plan) Type() protoreflect.MessageType {
	return _fastReflection_Plan_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Plan) New() protoreflect.Message {
	return new(fastReflection_Plan)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Plan) Interface() protoreflect.ProtoMessage {
	return (*Plan)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Plan) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_Plan_name, value) {
			return
		}
	}
	if x.Time != nil {
		value := protoreflect.ValueOfMessage(x.Time.ProtoReflect())
		if !f(fd_Plan_time, value) {
			return
		}
	}
	if x.Height != int64(0) {
		value := protoreflect.ValueOfInt64(x.Height)
		if !f(fd_Plan_height, value) {
			return
		}
	}
	if x.Info != "" {
		value := protoreflect.ValueOfString(x.Info)
		if !f(fd_Plan_info, value) {
			return
		}
	}
	if x.UpgradedClientState != nil {
		value := protoreflect.ValueOfMessage(x.UpgradedClientState.ProtoReflect())
		if !f(fd_Plan_upgraded_client_state, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Plan) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.Plan.name":
		return x.Name != ""
	case "cosmos.upgrade.v1beta1.Plan.time":
		return x.Time != nil
	case "cosmos.upgrade.v1beta1.Plan.height":
		return x.Height != int64(0)
	case "cosmos.upgrade.v1beta1.Plan.info":
		return x.Info != ""
	case "cosmos.upgrade.v1beta1.Plan.upgraded_client_state":
		return x.UpgradedClientState != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.Plan"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.Plan does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Plan) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.Plan.name":
		x.Name = ""
	case "cosmos.upgrade.v1beta1.Plan.time":
		x.Time = nil
	case "cosmos.upgrade.v1beta1.Plan.height":
		x.Height = int64(0)
	case "cosmos.upgrade.v1beta1.Plan.info":
		x.Info = ""
	case "cosmos.upgrade.v1beta1.Plan.upgraded_client_state":
		x.UpgradedClientState = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.Plan"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.Plan does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Plan) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.upgrade.v1beta1.Plan.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "cosmos.upgrade.v1beta1.Plan.time":
		value := x.Time
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "cosmos.upgrade.v1beta1.Plan.height":
		value := x.Height
		return protoreflect.ValueOfInt64(value)
	case "cosmos.upgrade.v1beta1.Plan.info":
		value := x.Info
		return protoreflect.ValueOfString(value)
	case "cosmos.upgrade.v1beta1.Plan.upgraded_client_state":
		value := x.UpgradedClientState
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.Plan"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.Plan does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Plan) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.Plan.name":
		x.Name = value.Interface().(string)
	case "cosmos.upgrade.v1beta1.Plan.time":
		x.Time = value.Message().Interface().(*timestamppb.Timestamp)
	case "cosmos.upgrade.v1beta1.Plan.height":
		x.Height = value.Int()
	case "cosmos.upgrade.v1beta1.Plan.info":
		x.Info = value.Interface().(string)
	case "cosmos.upgrade.v1beta1.Plan.upgraded_client_state":
		x.UpgradedClientState = value.Message().Interface().(*anypb.Any)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.Plan"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.Plan does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Plan) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.Plan.time":
		if x.Time == nil {
			x.Time = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Time.ProtoReflect())
	case "cosmos.upgrade.v1beta1.Plan.upgraded_client_state":
		if x.UpgradedClientState == nil {
			x.UpgradedClientState = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.UpgradedClientState.ProtoReflect())
	case "cosmos.upgrade.v1beta1.Plan.name":
		panic(fmt.Errorf("field name of message cosmos.upgrade.v1beta1.Plan is not mutable"))
	case "cosmos.upgrade.v1beta1.Plan.height":
		panic(fmt.Errorf("field height of message cosmos.upgrade.v1beta1.Plan is not mutable"))
	case "cosmos.upgrade.v1beta1.Plan.info":
		panic(fmt.Errorf("field info of message cosmos.upgrade.v1beta1.Plan is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.Plan"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.Plan does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Plan) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.Plan.name":
		return protoreflect.ValueOfString("")
	case "cosmos.upgrade.v1beta1.Plan.time":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "cosmos.upgrade.v1beta1.Plan.height":
		return protoreflect.ValueOfInt64(int64(0))
	case "cosmos.upgrade.v1beta1.Plan.info":
		return protoreflect.ValueOfString("")
	case "cosmos.upgrade.v1beta1.Plan.upgraded_client_state":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.Plan"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.Plan does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Plan) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.upgrade.v1beta1.Plan", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Plan) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Plan) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Plan) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Plan) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Plan)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Time != nil {
			l = options.Size(x.Time)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Height != 0 {
			n += 1 + runtime.Sov(uint64(x.Height))
		}
		l = len(x.Info)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.UpgradedClientState != nil {
			l = options.Size(x.UpgradedClientState)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Plan)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.UpgradedClientState != nil {
			encoded, err := options.Marshal(x.UpgradedClientState)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Info) > 0 {
			i -= len(x.Info)
			copy(dAtA[i:], x.Info)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Info)))
			i--
			dAtA[i] = 0x22
		}
		if x.Height != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Height))
			i--
			dAtA[i] = 0x18
		}
		if x.Time != nil {
			encoded, err := options.Marshal(x.Time)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Plan)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Plan: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Time == nil {
					x.Time = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Time); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
				}
				x.Height = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Height |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Info = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field UpgradedClientState", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.UpgradedClientState == nil {
					x.UpgradedClientState = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.UpgradedClientState); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_SoftwareUpgradeProposal             protoreflect.MessageDescriptor
	fd_SoftwareUpgradeProposal_title       protoreflect.FieldDescriptor
	fd_SoftwareUpgradeProposal_description protoreflect.FieldDescriptor
	fd_SoftwareUpgradeProposal_plan        protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_upgrade_v1beta1_upgrade_proto_init()
	md_SoftwareUpgradeProposal = File_cosmos_upgrade_v1beta1_upgrade_proto.Messages().ByName("SoftwareUpgradeProposal")
	fd_SoftwareUpgradeProposal_title = md_SoftwareUpgradeProposal.Fields().ByName("title")
	fd_SoftwareUpgradeProposal_description = md_SoftwareUpgradeProposal.Fields().ByName("description")
	fd_SoftwareUpgradeProposal_plan = md_SoftwareUpgradeProposal.Fields().ByName("plan")
}

var _ protoreflect.Message = (*fastReflection_SoftwareUpgradeProposal)(nil)

type fastReflection_SoftwareUpgradeProposal SoftwareUpgradeProposal

func (x *SoftwareUpgradeProposal) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SoftwareUpgradeProposal)(x)
}

func (x *SoftwareUpgradeProposal) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SoftwareUpgradeProposal_messageType fastReflection_SoftwareUpgradeProposal_messageType
var _ protoreflect.MessageType = fastReflection_SoftwareUpgradeProposal_messageType{}

type fastReflection_SoftwareUpgradeProposal_messageType struct{}

func (x fastReflection_SoftwareUpgradeProposal_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SoftwareUpgradeProposal)(nil)
}
func (x fastReflection_SoftwareUpgradeProposal_messageType) New() protoreflect.Message {
	return new(fastReflection_SoftwareUpgradeProposal)
}
func (x fastReflection_SoftwareUpgradeProposal_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SoftwareUpgradeProposal
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SoftwareUpgradeProposal) Descriptor() protoreflect.MessageDescriptor {
	return md_SoftwareUpgradeProposal
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SoftwareUpgradeProposal) Type() protoreflect.MessageType {
	return _fastReflection_SoftwareUpgradeProposal_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SoftwareUpgradeProposal) New() protoreflect.Message {
	return new(fastReflection_SoftwareUpgradeProposal)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SoftwareUpgradeProposal) Interface() protoreflect.ProtoMessage {
	return (*SoftwareUpgradeProposal)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SoftwareUpgradeProposal) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Title != "" {
		value := protoreflect.ValueOfString(x.Title)
		if !f(fd_SoftwareUpgradeProposal_title, value) {
			return
		}
	}
	if x.Description != "" {
		value := protoreflect.ValueOfString(x.Description)
		if !f(fd_SoftwareUpgradeProposal_description, value) {
			return
		}
	}
	if x.Plan != nil {
		value := protoreflect.ValueOfMessage(x.Plan.ProtoReflect())
		if !f(fd_SoftwareUpgradeProposal_plan, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_SoftwareUpgradeProposal) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.title":
		return x.Title != ""
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.description":
		return x.Description != ""
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.plan":
		return x.Plan != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.SoftwareUpgradeProposal"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.SoftwareUpgradeProposal does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SoftwareUpgradeProposal) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.title":
		x.Title = ""
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.description":
		x.Description = ""
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.plan":
		x.Plan = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.SoftwareUpgradeProposal"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.SoftwareUpgradeProposal does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SoftwareUpgradeProposal) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.title":
		value := x.Title
		return protoreflect.ValueOfString(value)
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.description":
		value := x.Description
		return protoreflect.ValueOfString(value)
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.plan":
		value := x.Plan
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.SoftwareUpgradeProposal"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.SoftwareUpgradeProposal does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SoftwareUpgradeProposal) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.title":
		x.Title = value.Interface().(string)
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.description":
		x.Description = value.Interface().(string)
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.plan":
		x.Plan = value.Message().Interface().(*Plan)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.SoftwareUpgradeProposal"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.SoftwareUpgradeProposal does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SoftwareUpgradeProposal) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.plan":
		if x.Plan == nil {
			x.Plan = new(Plan)
		}
		return protoreflect.ValueOfMessage(x.Plan.ProtoReflect())
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.title":
		panic(fmt.Errorf("field title of message cosmos.upgrade.v1beta1.SoftwareUpgradeProposal is not mutable"))
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.description":
		panic(fmt.Errorf("field description of message cosmos.upgrade.v1beta1.SoftwareUpgradeProposal is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.SoftwareUpgradeProposal"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.SoftwareUpgradeProposal does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SoftwareUpgradeProposal) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.title":
		return protoreflect.ValueOfString("")
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.description":
		return protoreflect.ValueOfString("")
	case "cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.plan":
		m := new(Plan)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.SoftwareUpgradeProposal"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.SoftwareUpgradeProposal does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SoftwareUpgradeProposal) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.upgrade.v1beta1.SoftwareUpgradeProposal", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SoftwareUpgradeProposal) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SoftwareUpgradeProposal) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_SoftwareUpgradeProposal) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SoftwareUpgradeProposal) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SoftwareUpgradeProposal)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Title)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Description)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Plan != nil {
			l = options.Size(x.Plan)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*SoftwareUpgradeProposal)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Plan != nil {
			encoded, err := options.Marshal(x.Plan)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Description) > 0 {
			i -= len(x.Description)
			copy(dAtA[i:], x.Description)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Description)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Title) > 0 {
			i -= len(x.Title)
			copy(dAtA[i:], x.Title)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Title)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*SoftwareUpgradeProposal)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SoftwareUpgradeProposal: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SoftwareUpgradeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Title = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Description = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Plan == nil {
					x.Plan = &Plan{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Plan); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_CancelSoftwareUpgradeProposal             protoreflect.MessageDescriptor
	fd_CancelSoftwareUpgradeProposal_title       protoreflect.FieldDescriptor
	fd_CancelSoftwareUpgradeProposal_description protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_upgrade_v1beta1_upgrade_proto_init()
	md_CancelSoftwareUpgradeProposal = File_cosmos_upgrade_v1beta1_upgrade_proto.Messages().ByName("CancelSoftwareUpgradeProposal")
	fd_CancelSoftwareUpgradeProposal_title = md_CancelSoftwareUpgradeProposal.Fields().ByName("title")
	fd_CancelSoftwareUpgradeProposal_description = md_CancelSoftwareUpgradeProposal.Fields().ByName("description")
}

var _ protoreflect.Message = (*fastReflection_CancelSoftwareUpgradeProposal)(nil)

type fastReflection_CancelSoftwareUpgradeProposal CancelSoftwareUpgradeProposal

func (x *CancelSoftwareUpgradeProposal) ProtoReflect() protoreflect.Message {
	return (*fastReflection_CancelSoftwareUpgradeProposal)(x)
}

func (x *CancelSoftwareUpgradeProposal) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_CancelSoftwareUpgradeProposal_messageType fastReflection_CancelSoftwareUpgradeProposal_messageType
var _ protoreflect.MessageType = fastReflection_CancelSoftwareUpgradeProposal_messageType{}

type fastReflection_CancelSoftwareUpgradeProposal_messageType struct{}

func (x fastReflection_CancelSoftwareUpgradeProposal_messageType) Zero() protoreflect.Message {
	return (*fastReflection_CancelSoftwareUpgradeProposal)(nil)
}
func (x fastReflection_CancelSoftwareUpgradeProposal_messageType) New() protoreflect.Message {
	return new(fastReflection_CancelSoftwareUpgradeProposal)
}
func (x fastReflection_CancelSoftwareUpgradeProposal_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_CancelSoftwareUpgradeProposal
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_CancelSoftwareUpgradeProposal) Descriptor() protoreflect.MessageDescriptor {
	return md_CancelSoftwareUpgradeProposal
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_CancelSoftwareUpgradeProposal) Type() protoreflect.MessageType {
	return _fastReflection_CancelSoftwareUpgradeProposal_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_CancelSoftwareUpgradeProposal) New() protoreflect.Message {
	return new(fastReflection_CancelSoftwareUpgradeProposal)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_CancelSoftwareUpgradeProposal) Interface() protoreflect.ProtoMessage {
	return (*CancelSoftwareUpgradeProposal)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_CancelSoftwareUpgradeProposal) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Title != "" {
		value := protoreflect.ValueOfString(x.Title)
		if !f(fd_CancelSoftwareUpgradeProposal_title, value) {
			return
		}
	}
	if x.Description != "" {
		value := protoreflect.ValueOfString(x.Description)
		if !f(fd_CancelSoftwareUpgradeProposal_description, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_CancelSoftwareUpgradeProposal) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.title":
		return x.Title != ""
	case "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.description":
		return x.Description != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CancelSoftwareUpgradeProposal) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.title":
		x.Title = ""
	case "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.description":
		x.Description = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_CancelSoftwareUpgradeProposal) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.title":
		value := x.Title
		return protoreflect.ValueOfString(value)
	case "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.description":
		value := x.Description
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CancelSoftwareUpgradeProposal) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.title":
		x.Title = value.Interface().(string)
	case "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.description":
		x.Description = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CancelSoftwareUpgradeProposal) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.title":
		panic(fmt.Errorf("field title of message cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal is not mutable"))
	case "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.description":
		panic(fmt.Errorf("field description of message cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_CancelSoftwareUpgradeProposal) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.title":
		return protoreflect.ValueOfString("")
	case "cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal.description":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_CancelSoftwareUpgradeProposal) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_CancelSoftwareUpgradeProposal) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_CancelSoftwareUpgradeProposal) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_CancelSoftwareUpgradeProposal) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_CancelSoftwareUpgradeProposal) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*CancelSoftwareUpgradeProposal)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Title)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Description)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*CancelSoftwareUpgradeProposal)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Description) > 0 {
			i -= len(x.Description)
			copy(dAtA[i:], x.Description)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Description)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Title) > 0 {
			i -= len(x.Title)
			copy(dAtA[i:], x.Title)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Title)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*CancelSoftwareUpgradeProposal)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CancelSoftwareUpgradeProposal: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: CancelSoftwareUpgradeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Title = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Description = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_ModuleVersion         protoreflect.MessageDescriptor
	fd_ModuleVersion_name    protoreflect.FieldDescriptor
	fd_ModuleVersion_version protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_upgrade_v1beta1_upgrade_proto_init()
	md_ModuleVersion = File_cosmos_upgrade_v1beta1_upgrade_proto.Messages().ByName("ModuleVersion")
	fd_ModuleVersion_name = md_ModuleVersion.Fields().ByName("name")
	fd_ModuleVersion_version = md_ModuleVersion.Fields().ByName("version")
}

var _ protoreflect.Message = (*fastReflection_ModuleVersion)(nil)

type fastReflection_ModuleVersion ModuleVersion

func (x *ModuleVersion) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ModuleVersion)(x)
}

func (x *ModuleVersion) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ModuleVersion_messageType fastReflection_ModuleVersion_messageType
var _ protoreflect.MessageType = fastReflection_ModuleVersion_messageType{}

type fastReflection_ModuleVersion_messageType struct{}

func (x fastReflection_ModuleVersion_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ModuleVersion)(nil)
}
func (x fastReflection_ModuleVersion_messageType) New() protoreflect.Message {
	return new(fastReflection_ModuleVersion)
}
func (x fastReflection_ModuleVersion_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ModuleVersion
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ModuleVersion) Descriptor() protoreflect.MessageDescriptor {
	return md_ModuleVersion
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ModuleVersion) Type() protoreflect.MessageType {
	return _fastReflection_ModuleVersion_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ModuleVersion) New() protoreflect.Message {
	return new(fastReflection_ModuleVersion)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ModuleVersion) Interface() protoreflect.ProtoMessage {
	return (*ModuleVersion)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ModuleVersion) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Name != "" {
		value := protoreflect.ValueOfString(x.Name)
		if !f(fd_ModuleVersion_name, value) {
			return
		}
	}
	if x.Version != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Version)
		if !f(fd_ModuleVersion_version, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ModuleVersion) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.ModuleVersion.name":
		return x.Name != ""
	case "cosmos.upgrade.v1beta1.ModuleVersion.version":
		return x.Version != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.ModuleVersion"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.ModuleVersion does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ModuleVersion) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.ModuleVersion.name":
		x.Name = ""
	case "cosmos.upgrade.v1beta1.ModuleVersion.version":
		x.Version = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.ModuleVersion"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.ModuleVersion does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ModuleVersion) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.upgrade.v1beta1.ModuleVersion.name":
		value := x.Name
		return protoreflect.ValueOfString(value)
	case "cosmos.upgrade.v1beta1.ModuleVersion.version":
		value := x.Version
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.ModuleVersion"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.ModuleVersion does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ModuleVersion) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.ModuleVersion.name":
		x.Name = value.Interface().(string)
	case "cosmos.upgrade.v1beta1.ModuleVersion.version":
		x.Version = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.ModuleVersion"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.ModuleVersion does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ModuleVersion) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.ModuleVersion.name":
		panic(fmt.Errorf("field name of message cosmos.upgrade.v1beta1.ModuleVersion is not mutable"))
	case "cosmos.upgrade.v1beta1.ModuleVersion.version":
		panic(fmt.Errorf("field version of message cosmos.upgrade.v1beta1.ModuleVersion is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.ModuleVersion"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.ModuleVersion does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ModuleVersion) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.upgrade.v1beta1.ModuleVersion.name":
		return protoreflect.ValueOfString("")
	case "cosmos.upgrade.v1beta1.ModuleVersion.version":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.upgrade.v1beta1.ModuleVersion"))
		}
		panic(fmt.Errorf("message cosmos.upgrade.v1beta1.ModuleVersion does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ModuleVersion) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.upgrade.v1beta1.ModuleVersion", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ModuleVersion) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ModuleVersion) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ModuleVersion) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ModuleVersion) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ModuleVersion)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Name)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Version != 0 {
			n += 1 + runtime.Sov(uint64(x.Version))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ModuleVersion)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Version != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Version))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Name) > 0 {
			i -= len(x.Name)
			copy(dAtA[i:], x.Name)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Name)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ModuleVersion)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ModuleVersion: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ModuleVersion: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Name = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
				}
				x.Version = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Version |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/upgrade/v1beta1/upgrade.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Plan specifies information about a planned upgrade and when it should occur.
type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Sets the name for the upgrade. This name will be used by the upgraded
	// version of the software to apply any special "on-upgrade" commands during
	// the first BeginBlock method after the upgrade is applied. It is also used
	// to detect whether a software version can handle a given upgrade. If no
	// upgrade handler with this name has been set in the software, it will be
	// assumed that the software is out-of-date when the upgrade Time or Height is
	// reached and the software will exit.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Deprecated: Time based upgrades have been deprecated. Time based upgrade logic
	// has been removed from the SDK.
	// If this field is not empty, an error will be thrown.
	//
	// Deprecated: Do not use.
	Time *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// The height at which the upgrade must be performed.
	// Only used if Time is not set.
	Height int64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	// Any application specific upgrade info to be included on-chain
	// such as a git commit that validators could automatically upgrade to
	Info string `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	// Deprecated: UpgradedClientState field has been deprecated. IBC upgrade logic has been
	// moved to the IBC module in the sub module 02-client.
	// If this field is not empty, an error will be thrown.
	//
	// Deprecated: Do not use.
	UpgradedClientState *anypb.Any `protobuf:"bytes,5,opt,name=upgraded_client_state,json=upgradedClientState,proto3" json:"upgraded_client_state,omitempty"`
}

func (x *Plan) Reset() {
	*x = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_cosmos_upgrade_v1beta1_upgrade_proto_rawDescGZIP(), []int{0}
}

func (x *Plan) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Deprecated: Do not use.
func (x *Plan) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Plan) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Plan) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

// Deprecated: Do not use.
func (x *Plan) GetUpgradedClientState() *anypb.Any {
	if x != nil {
		return x.UpgradedClientState
	}
	return nil
}

// SoftwareUpgradeProposal is a gov Content type for initiating a software
// upgrade.
type SoftwareUpgradeProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Plan        *Plan  `protobuf:"bytes,3,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *SoftwareUpgradeProposal) Reset() {
	*x = SoftwareUpgradeProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoftwareUpgradeProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoftwareUpgradeProposal) ProtoMessage() {}

// Deprecated: Use SoftwareUpgradeProposal.ProtoReflect.Descriptor instead.
func (*SoftwareUpgradeProposal) Descriptor() ([]byte, []int) {
	return file_cosmos_upgrade_v1beta1_upgrade_proto_rawDescGZIP(), []int{1}
}

func (x *SoftwareUpgradeProposal) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SoftwareUpgradeProposal) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SoftwareUpgradeProposal) GetPlan() *Plan {
	if x != nil {
		return x.Plan
	}
	return nil
}

// CancelSoftwareUpgradeProposal is a gov Content type for cancelling a software
// upgrade.
type CancelSoftwareUpgradeProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CancelSoftwareUpgradeProposal) Reset() {
	*x = CancelSoftwareUpgradeProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelSoftwareUpgradeProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelSoftwareUpgradeProposal) ProtoMessage() {}

// Deprecated: Use CancelSoftwareUpgradeProposal.ProtoReflect.Descriptor instead.
func (*CancelSoftwareUpgradeProposal) Descriptor() ([]byte, []int) {
	return file_cosmos_upgrade_v1beta1_upgrade_proto_rawDescGZIP(), []int{2}
}

func (x *CancelSoftwareUpgradeProposal) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CancelSoftwareUpgradeProposal) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// ModuleVersion specifies a module and its consensus version.
//
// Since: cosmos-sdk 0.43
type ModuleVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the app module
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// consensus version of the app module
	Version uint64 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ModuleVersion) Reset() {
	*x = ModuleVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleVersion) ProtoMessage() {}

// Deprecated: Use ModuleVersion.ProtoReflect.Descriptor instead.
func (*ModuleVersion) Descriptor() ([]byte, []int) {
	return file_cosmos_upgrade_v1beta1_upgrade_proto_rawDescGZIP(), []int{3}
}

func (x *ModuleVersion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModuleVersion) GetVersion() uint64 {
	if x != nil {
		return x.Version
	}
	return 0
}

var File_cosmos_upgrade_v1beta1_upgrade_proto protoreflect.FileDescriptor

var file_cosmos_upgrade_v1beta1_upgrade_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x75,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xda, 0x01, 0x0a, 0x04, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0x18, 0x01, 0xc8, 0xde, 0x1f, 0x00, 0x90,
	0xdf, 0x1f, 0x01, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x4c, 0x0a, 0x15, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x64, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x02, 0x18, 0x01, 0x52, 0x13,
	0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x3a, 0x08, 0x98, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x01, 0x22, 0x93, 0x01,
	0x0a, 0x17, 0x53, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x36, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x42, 0x04, 0xc8,
	0xde, 0x1f, 0x00, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x3a, 0x08, 0x98, 0xa0, 0x1f, 0x00, 0xe8,
	0xa0, 0x1f, 0x01, 0x22, 0x61, 0x0a, 0x1d, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x6f, 0x66,
	0x74, 0x77, 0x61, 0x72, 0x65, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x08, 0x98, 0xa0,
	0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x01, 0x22, 0x47, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x08, 0x98, 0xa0, 0x1f, 0x01, 0xe8, 0xa0, 0x1f, 0x01, 0x42,
	0xf0, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x75,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0c,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x55, 0x58, 0xaa, 0x02, 0x16, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x16, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x55,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02,
	0x22, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5c,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x55, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xc8, 0xe1,
	0x1e, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_upgrade_v1beta1_upgrade_proto_rawDescOnce sync.Once
	file_cosmos_upgrade_v1beta1_upgrade_proto_rawDescData = file_cosmos_upgrade_v1beta1_upgrade_proto_rawDesc
)

func file_cosmos_upgrade_v1beta1_upgrade_proto_rawDescGZIP() []byte {
	file_cosmos_upgrade_v1beta1_upgrade_proto_rawDescOnce.Do(func() {
		file_cosmos_upgrade_v1beta1_upgrade_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_upgrade_v1beta1_upgrade_proto_rawDescData)
	})
	return file_cosmos_upgrade_v1beta1_upgrade_proto_rawDescData
}

var file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cosmos_upgrade_v1beta1_upgrade_proto_goTypes = []interface{}{
	(*Plan)(nil),                          // 0: cosmos.upgrade.v1beta1.Plan
	(*SoftwareUpgradeProposal)(nil),       // 1: cosmos.upgrade.v1beta1.SoftwareUpgradeProposal
	(*CancelSoftwareUpgradeProposal)(nil), // 2: cosmos.upgrade.v1beta1.CancelSoftwareUpgradeProposal
	(*ModuleVersion)(nil),                 // 3: cosmos.upgrade.v1beta1.ModuleVersion
	(*timestamppb.Timestamp)(nil),         // 4: google.protobuf.Timestamp
	(*anypb.Any)(nil),                     // 5: google.protobuf.Any
}
var file_cosmos_upgrade_v1beta1_upgrade_proto_depIdxs = []int32{
	4, // 0: cosmos.upgrade.v1beta1.Plan.time:type_name -> google.protobuf.Timestamp
	5, // 1: cosmos.upgrade.v1beta1.Plan.upgraded_client_state:type_name -> google.protobuf.Any
	0, // 2: cosmos.upgrade.v1beta1.SoftwareUpgradeProposal.plan:type_name -> cosmos.upgrade.v1beta1.Plan
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cosmos_upgrade_v1beta1_upgrade_proto_init() }
func file_cosmos_upgrade_v1beta1_upgrade_proto_init() {
	if File_cosmos_upgrade_v1beta1_upgrade_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plan); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoftwareUpgradeProposal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelSoftwareUpgradeProposal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleVersion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_upgrade_v1beta1_upgrade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_upgrade_v1beta1_upgrade_proto_goTypes,
		DependencyIndexes: file_cosmos_upgrade_v1beta1_upgrade_proto_depIdxs,
		MessageInfos:      file_cosmos_upgrade_v1beta1_upgrade_proto_msgTypes,
	}.Build()
	File_cosmos_upgrade_v1beta1_upgrade_proto = out.File
	file_cosmos_upgrade_v1beta1_upgrade_proto_rawDesc = nil
	file_cosmos_upgrade_v1beta1_upgrade_proto_goTypes = nil
	file_cosmos_upgrade_v1beta1_upgrade_proto_depIdxs = nil
}
