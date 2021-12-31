// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: nhk.proto

package nhk

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=endDate,proto3" json:"endDate,omitempty"`
}

func (x *NewsRequest) Reset() {
	*x = NewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nhk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsRequest) ProtoMessage() {}

func (x *NewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nhk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsRequest.ProtoReflect.Descriptor instead.
func (*NewsRequest) Descriptor() ([]byte, []int) {
	return file_nhk_proto_rawDescGZIP(), []int{0}
}

func (x *NewsRequest) GetStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.StartDate
	}
	return nil
}

func (x *NewsRequest) GetEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EndDate
	}
	return nil
}

type NewsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error  `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	News  []*News `protobuf:"bytes,2,rep,name=news,proto3" json:"news,omitempty"`
}

func (x *NewsReply) Reset() {
	*x = NewsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nhk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsReply) ProtoMessage() {}

func (x *NewsReply) ProtoReflect() protoreflect.Message {
	mi := &file_nhk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsReply.ProtoReflect.Descriptor instead.
func (*NewsReply) Descriptor() ([]byte, []int) {
	return file_nhk_proto_rawDescGZIP(), []int{1}
}

func (x *NewsReply) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *NewsReply) GetNews() []*News {
	if x != nil {
		return x.News
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nhk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_nhk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_nhk_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type News struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsId          string                 `protobuf:"bytes,1,opt,name=newsId,proto3" json:"newsId,omitempty"`
	Title           string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	TitleWithRuby   string                 `protobuf:"bytes,3,opt,name=titleWithRuby,proto3" json:"titleWithRuby,omitempty"`
	OutlineWithRuby string                 `protobuf:"bytes,4,opt,name=outlineWithRuby,proto3" json:"outlineWithRuby,omitempty"`
	Body            string                 `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Url             string                 `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	M3U8Url         string                 `protobuf:"bytes,7,opt,name=m3u8Url,proto3" json:"m3u8Url,omitempty"`
	ImageUrl        string                 `protobuf:"bytes,8,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
	PublishedAtUtc  *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=publishedAtUtc,proto3" json:"publishedAtUtc,omitempty"`
}

func (x *News) Reset() {
	*x = News{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nhk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *News) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*News) ProtoMessage() {}

func (x *News) ProtoReflect() protoreflect.Message {
	mi := &file_nhk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use News.ProtoReflect.Descriptor instead.
func (*News) Descriptor() ([]byte, []int) {
	return file_nhk_proto_rawDescGZIP(), []int{3}
}

func (x *News) GetNewsId() string {
	if x != nil {
		return x.NewsId
	}
	return ""
}

func (x *News) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *News) GetTitleWithRuby() string {
	if x != nil {
		return x.TitleWithRuby
	}
	return ""
}

func (x *News) GetOutlineWithRuby() string {
	if x != nil {
		return x.OutlineWithRuby
	}
	return ""
}

func (x *News) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *News) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *News) GetM3U8Url() string {
	if x != nil {
		return x.M3U8Url
	}
	return ""
}

func (x *News) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *News) GetPublishedAtUtc() *timestamppb.Timestamp {
	if x != nil {
		return x.PublishedAtUtc
	}
	return nil
}

var File_nhk_proto protoreflect.FileDescriptor

var file_nhk_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6e, 0x68, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0b,
	0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x44, 0x0a, 0x09, 0x4e,
	0x65, 0x77, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x04, 0x6e, 0x65, 0x77, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x04, 0x6e, 0x65, 0x77,
	0x73, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa4, 0x02, 0x0a, 0x04, 0x4e, 0x65, 0x77,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x65, 0x77, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x73, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x75, 0x62, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x52, 0x75, 0x62, 0x79, 0x12, 0x28, 0x0a, 0x0f, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x6e, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x52, 0x75, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x6f, 0x75, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x57, 0x69, 0x74, 0x68, 0x52, 0x75, 0x62, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x33, 0x75, 0x38, 0x55, 0x72, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x33, 0x75, 0x38, 0x55, 0x72, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x42, 0x0a, 0x0e, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x55, 0x74, 0x63, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x55, 0x74, 0x63, 0x32,
	0x33, 0x0a, 0x0a, 0x4e, 0x68, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x77, 0x73, 0x12, 0x0c, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x64, 0x65, 0x6b, 0x69, 0x72, 0x75, 0x2e, 0x61,
	0x70, 0x70, 0x2f, 0x6e, 0x68, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nhk_proto_rawDescOnce sync.Once
	file_nhk_proto_rawDescData = file_nhk_proto_rawDesc
)

func file_nhk_proto_rawDescGZIP() []byte {
	file_nhk_proto_rawDescOnce.Do(func() {
		file_nhk_proto_rawDescData = protoimpl.X.CompressGZIP(file_nhk_proto_rawDescData)
	})
	return file_nhk_proto_rawDescData
}

var file_nhk_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nhk_proto_goTypes = []interface{}{
	(*NewsRequest)(nil),           // 0: NewsRequest
	(*NewsReply)(nil),             // 1: NewsReply
	(*Error)(nil),                 // 2: Error
	(*News)(nil),                  // 3: News
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_nhk_proto_depIdxs = []int32{
	4, // 0: NewsRequest.startDate:type_name -> google.protobuf.Timestamp
	4, // 1: NewsRequest.endDate:type_name -> google.protobuf.Timestamp
	2, // 2: NewsReply.error:type_name -> Error
	3, // 3: NewsReply.news:type_name -> News
	4, // 4: News.publishedAtUtc:type_name -> google.protobuf.Timestamp
	0, // 5: NhkService.GetNews:input_type -> NewsRequest
	1, // 6: NhkService.GetNews:output_type -> NewsReply
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_nhk_proto_init() }
func file_nhk_proto_init() {
	if File_nhk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nhk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsRequest); i {
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
		file_nhk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewsReply); i {
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
		file_nhk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_nhk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*News); i {
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
			RawDescriptor: file_nhk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nhk_proto_goTypes,
		DependencyIndexes: file_nhk_proto_depIdxs,
		MessageInfos:      file_nhk_proto_msgTypes,
	}.Build()
	File_nhk_proto = out.File
	file_nhk_proto_rawDesc = nil
	file_nhk_proto_goTypes = nil
	file_nhk_proto_depIdxs = nil
}
