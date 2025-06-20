// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: productService.proto

package productpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetProductRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	mi := &file_productService_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_productService_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_productService_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetProductResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Product       *Product               `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetProductResponse) Reset() {
	*x = GetProductResponse{}
	mi := &file_productService_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductResponse) ProtoMessage() {}

func (x *GetProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_productService_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductResponse.ProtoReflect.Descriptor instead.
func (*GetProductResponse) Descriptor() ([]byte, []int) {
	return file_productService_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductResponse) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

var File_productService_proto protoreflect.FileDescriptor

const file_productService_proto_rawDesc = "" +
	"\n" +
	"\x14productService.proto\x12\x0eproductservice\x1a\rproduct.proto\"#\n" +
	"\x11GetProductRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\"@\n" +
	"\x12GetProductResponse\x12*\n" +
	"\aproduct\x18\x01 \x01(\v2\x10.product.ProductR\aproduct2^\n" +
	"\aProduct\x12S\n" +
	"\n" +
	"GetProduct\x12!.productservice.GetProductRequest\x1a\".productservice.GetProductResponseB\x1aZ\x18productService/productpbb\x06proto3"

var (
	file_productService_proto_rawDescOnce sync.Once
	file_productService_proto_rawDescData []byte
)

func file_productService_proto_rawDescGZIP() []byte {
	file_productService_proto_rawDescOnce.Do(func() {
		file_productService_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_productService_proto_rawDesc), len(file_productService_proto_rawDesc)))
	})
	return file_productService_proto_rawDescData
}

var file_productService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_productService_proto_goTypes = []any{
	(*GetProductRequest)(nil),  // 0: productservice.GetProductRequest
	(*GetProductResponse)(nil), // 1: productservice.GetProductResponse
	(*Product)(nil),            // 2: product.Product
}
var file_productService_proto_depIdxs = []int32{
	2, // 0: productservice.GetProductResponse.product:type_name -> product.Product
	0, // 1: productservice.Product.GetProduct:input_type -> productservice.GetProductRequest
	1, // 2: productservice.Product.GetProduct:output_type -> productservice.GetProductResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_productService_proto_init() }
func file_productService_proto_init() {
	if File_productService_proto != nil {
		return
	}
	file_product_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_productService_proto_rawDesc), len(file_productService_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_productService_proto_goTypes,
		DependencyIndexes: file_productService_proto_depIdxs,
		MessageInfos:      file_productService_proto_msgTypes,
	}.Build()
	File_productService_proto = out.File
	file_productService_proto_goTypes = nil
	file_productService_proto_depIdxs = nil
}
