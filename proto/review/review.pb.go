// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v4.25.2
// source: proto/review/review.proto

package review

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Review struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderId       string                 `protobuf:"bytes,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Rating        int32                  `protobuf:"varint,4,opt,name=rating,proto3" json:"rating,omitempty"`
	Comment       string                 `protobuf:"bytes,5,opt,name=comment,proto3" json:"comment,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Review) Reset() {
	*x = Review{}
	mi := &file_proto_review_review_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_proto_review_review_proto_rawDescGZIP(), []int{0}
}

func (x *Review) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Review) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Review) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Review) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Review) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Review) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderId       string                 `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Rating        int32                  `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
	Comment       string                 `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateReviewRequest) Reset() {
	*x = CreateReviewRequest{}
	mi := &file_proto_review_review_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReviewRequest) ProtoMessage() {}

func (x *CreateReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReviewRequest.ProtoReflect.Descriptor instead.
func (*CreateReviewRequest) Descriptor() ([]byte, []int) {
	return file_proto_review_review_proto_rawDescGZIP(), []int{1}
}

func (x *CreateReviewRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateReviewRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *CreateReviewRequest) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *CreateReviewRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type GetReviewRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetReviewRequest) Reset() {
	*x = GetReviewRequest{}
	mi := &file_proto_review_review_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewRequest) ProtoMessage() {}

func (x *GetReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewRequest.ProtoReflect.Descriptor instead.
func (*GetReviewRequest) Descriptor() ([]byte, []int) {
	return file_proto_review_review_proto_rawDescGZIP(), []int{2}
}

func (x *GetReviewRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ReviewResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Review        *Review                `protobuf:"bytes,1,opt,name=review,proto3" json:"review,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReviewResponse) Reset() {
	*x = ReviewResponse{}
	mi := &file_proto_review_review_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReviewResponse) ProtoMessage() {}

func (x *ReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_review_review_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReviewResponse.ProtoReflect.Descriptor instead.
func (*ReviewResponse) Descriptor() ([]byte, []int) {
	return file_proto_review_review_proto_rawDescGZIP(), []int{3}
}

func (x *ReviewResponse) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

var File_proto_review_review_proto protoreflect.FileDescriptor

var file_proto_review_review_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x22, 0x9d, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x7b, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x0e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x06, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x32, 0x93,
	0x01, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x43, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x1b, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x41, 0x6c, 0x65, 0x6b, 0x73, 0x4b, 0x69, 0x73, 0x6c, 0x6f, 0x76, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_review_review_proto_rawDescOnce sync.Once
	file_proto_review_review_proto_rawDescData = file_proto_review_review_proto_rawDesc
)

func file_proto_review_review_proto_rawDescGZIP() []byte {
	file_proto_review_review_proto_rawDescOnce.Do(func() {
		file_proto_review_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_review_review_proto_rawDescData)
	})
	return file_proto_review_review_proto_rawDescData
}

var file_proto_review_review_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_review_review_proto_goTypes = []any{
	(*Review)(nil),              // 0: review.Review
	(*CreateReviewRequest)(nil), // 1: review.CreateReviewRequest
	(*GetReviewRequest)(nil),    // 2: review.GetReviewRequest
	(*ReviewResponse)(nil),      // 3: review.ReviewResponse
}
var file_proto_review_review_proto_depIdxs = []int32{
	0, // 0: review.ReviewResponse.review:type_name -> review.Review
	1, // 1: review.ReviewService.CreateReview:input_type -> review.CreateReviewRequest
	2, // 2: review.ReviewService.GetReview:input_type -> review.GetReviewRequest
	3, // 3: review.ReviewService.CreateReview:output_type -> review.ReviewResponse
	3, // 4: review.ReviewService.GetReview:output_type -> review.ReviewResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_review_review_proto_init() }
func file_proto_review_review_proto_init() {
	if File_proto_review_review_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_review_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_review_review_proto_goTypes,
		DependencyIndexes: file_proto_review_review_proto_depIdxs,
		MessageInfos:      file_proto_review_review_proto_msgTypes,
	}.Build()
	File_proto_review_review_proto = out.File
	file_proto_review_review_proto_rawDesc = nil
	file_proto_review_review_proto_goTypes = nil
	file_proto_review_review_proto_depIdxs = nil
}