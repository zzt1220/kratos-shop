// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: api/goods/v1/goods.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GoodsClient is the client API for Goods service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodsClient interface {
	// 商品分类
	GetAllCategoryList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CategoryListResponse, error)
	GetSubCategory(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*SubCategoryListResponse, error)
	CreateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*CategoryInfoResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 商品品牌
	BrandList(ctx context.Context, in *BrandListRequest, opts ...grpc.CallOption) (*BrandListResponse, error)
	CreateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*BrandInfoResponse, error)
	DeleteBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 商品类型 goods_property_names
	// 商品类型不同于商品分类，指的是依据某一类商品的相同属性归纳成的属性集合 // 手机类型都有屏幕尺寸、网络制式等共同的属性
	CreateGoodsType(ctx context.Context, in *GoodsTypeRequest, opts ...grpc.CallOption) (*GoodsTypeResponse, error)
	// 创建商品规格 也就是 手机的颜色、内存版本、购买方式之类的
	// 商品规格的值，比如手机颜色对应的有 红、白、黑，内存，128g、256g, 也一起创建了
	CreateGoodsSpecification(ctx context.Context, in *SpecificationRequest, opts ...grpc.CallOption) (*SpecificationResponse, error)
	// 商品参数属性 ,手机:主体,屏幕, 操作系统,网络支持之类的
	CreateAttrGroup(ctx context.Context, in *AttrGroupRequest, opts ...grpc.CallOption) (*AttrGroupResponse, error)
	// 商品参数属性组下的一些信息 ,主体:上市年份 产品名称 ,网络支持 5G网络,双卡双待类型,
	CreateAttrValue(ctx context.Context, in *AttrRequest, opts ...grpc.CallOption) (*AttrResponse, error)
	// 商品接口
	CreateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...grpc.CallOption) (*CreateGoodsResponse, error)
	UpdateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GoodsList(ctx context.Context, in *GoodsFilterRequest, opts ...grpc.CallOption) (*GoodsListResponse, error)
}

type goodsClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodsClient(cc grpc.ClientConnInterface) GoodsClient {
	return &goodsClient{cc}
}

func (c *goodsClient) GetAllCategoryList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CategoryListResponse, error) {
	out := new(CategoryListResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/GetAllCategoryList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GetSubCategory(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*SubCategoryListResponse, error) {
	out := new(SubCategoryListResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/GetSubCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*CategoryInfoResponse, error) {
	out := new(CategoryInfoResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/DeleteCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateCategory(ctx context.Context, in *CategoryInfoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/UpdateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) BrandList(ctx context.Context, in *BrandListRequest, opts ...grpc.CallOption) (*BrandListResponse, error) {
	out := new(BrandListResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/BrandList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*BrandInfoResponse, error) {
	out := new(BrandInfoResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/CreateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) DeleteBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/DeleteBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateBrand(ctx context.Context, in *BrandRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/UpdateBrand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateGoodsType(ctx context.Context, in *GoodsTypeRequest, opts ...grpc.CallOption) (*GoodsTypeResponse, error) {
	out := new(GoodsTypeResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/CreateGoodsType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateGoodsSpecification(ctx context.Context, in *SpecificationRequest, opts ...grpc.CallOption) (*SpecificationResponse, error) {
	out := new(SpecificationResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/CreateGoodsSpecification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateAttrGroup(ctx context.Context, in *AttrGroupRequest, opts ...grpc.CallOption) (*AttrGroupResponse, error) {
	out := new(AttrGroupResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/CreateAttrGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateAttrValue(ctx context.Context, in *AttrRequest, opts ...grpc.CallOption) (*AttrResponse, error) {
	out := new(AttrResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/CreateAttrValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) CreateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...grpc.CallOption) (*CreateGoodsResponse, error) {
	out := new(CreateGoodsResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/CreateGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) UpdateGoods(ctx context.Context, in *CreateGoodsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/UpdateGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodsClient) GoodsList(ctx context.Context, in *GoodsFilterRequest, opts ...grpc.CallOption) (*GoodsListResponse, error) {
	out := new(GoodsListResponse)
	err := c.cc.Invoke(ctx, "/goods.v1.Goods/GoodsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodsServer is the s API for Goods service.
// All implementations must embed UnimplementedGoodsServer
// for forward compatibility
type GoodsServer interface {
	// 商品分类
	GetAllCategoryList(context.Context, *emptypb.Empty) (*CategoryListResponse, error)
	GetSubCategory(context.Context, *CategoryListRequest) (*SubCategoryListResponse, error)
	CreateCategory(context.Context, *CategoryInfoRequest) (*CategoryInfoResponse, error)
	DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error)
	UpdateCategory(context.Context, *CategoryInfoRequest) (*emptypb.Empty, error)
	// 商品品牌
	BrandList(context.Context, *BrandListRequest) (*BrandListResponse, error)
	CreateBrand(context.Context, *BrandRequest) (*BrandInfoResponse, error)
	DeleteBrand(context.Context, *BrandRequest) (*emptypb.Empty, error)
	UpdateBrand(context.Context, *BrandRequest) (*emptypb.Empty, error)
	// 商品类型 goods_property_names
	// 商品类型不同于商品分类，指的是依据某一类商品的相同属性归纳成的属性集合 // 手机类型都有屏幕尺寸、网络制式等共同的属性
	CreateGoodsType(context.Context, *GoodsTypeRequest) (*GoodsTypeResponse, error)
	// 创建商品规格 也就是 手机的颜色、内存版本、购买方式之类的
	// 商品规格的值，比如手机颜色对应的有 红、白、黑，内存，128g、256g, 也一起创建了
	CreateGoodsSpecification(context.Context, *SpecificationRequest) (*SpecificationResponse, error)
	// 商品参数属性 ,手机:主体,屏幕, 操作系统,网络支持之类的
	CreateAttrGroup(context.Context, *AttrGroupRequest) (*AttrGroupResponse, error)
	// 商品参数属性组下的一些信息 ,主体:上市年份 产品名称 ,网络支持 5G网络,双卡双待类型,
	CreateAttrValue(context.Context, *AttrRequest) (*AttrResponse, error)
	// 商品接口
	CreateGoods(context.Context, *CreateGoodsRequest) (*CreateGoodsResponse, error)
	UpdateGoods(context.Context, *CreateGoodsRequest) (*emptypb.Empty, error)
	GoodsList(context.Context, *GoodsFilterRequest) (*GoodsListResponse, error)
	mustEmbedUnimplementedGoodsServer()
}

// UnimplementedGoodsServer must be embedded to have forward compatible implementations.
type UnimplementedGoodsServer struct {
}

func (UnimplementedGoodsServer) GetAllCategoryList(context.Context, *emptypb.Empty) (*CategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategoryList not implemented")
}
func (UnimplementedGoodsServer) GetSubCategory(context.Context, *CategoryListRequest) (*SubCategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubCategory not implemented")
}
func (UnimplementedGoodsServer) CreateCategory(context.Context, *CategoryInfoRequest) (*CategoryInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedGoodsServer) DeleteCategory(context.Context, *DeleteCategoryRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedGoodsServer) UpdateCategory(context.Context, *CategoryInfoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategory not implemented")
}
func (UnimplementedGoodsServer) BrandList(context.Context, *BrandListRequest) (*BrandListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BrandList not implemented")
}
func (UnimplementedGoodsServer) CreateBrand(context.Context, *BrandRequest) (*BrandInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBrand not implemented")
}
func (UnimplementedGoodsServer) DeleteBrand(context.Context, *BrandRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBrand not implemented")
}
func (UnimplementedGoodsServer) UpdateBrand(context.Context, *BrandRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBrand not implemented")
}
func (UnimplementedGoodsServer) CreateGoodsType(context.Context, *GoodsTypeRequest) (*GoodsTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoodsType not implemented")
}
func (UnimplementedGoodsServer) CreateGoodsSpecification(context.Context, *SpecificationRequest) (*SpecificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoodsSpecification not implemented")
}
func (UnimplementedGoodsServer) CreateAttrGroup(context.Context, *AttrGroupRequest) (*AttrGroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttrGroup not implemented")
}
func (UnimplementedGoodsServer) CreateAttrValue(context.Context, *AttrRequest) (*AttrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAttrValue not implemented")
}
func (UnimplementedGoodsServer) CreateGoods(context.Context, *CreateGoodsRequest) (*CreateGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGoods not implemented")
}
func (UnimplementedGoodsServer) UpdateGoods(context.Context, *CreateGoodsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoods not implemented")
}
func (UnimplementedGoodsServer) GoodsList(context.Context, *GoodsFilterRequest) (*GoodsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoodsList not implemented")
}
func (UnimplementedGoodsServer) mustEmbedUnimplementedGoodsServer() {}

// UnsafeGoodsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodsServer will
// result in compilation errors.
type UnsafeGoodsServer interface {
	mustEmbedUnimplementedGoodsServer()
}

func RegisterGoodsServer(s grpc.ServiceRegistrar, srv GoodsServer) {
	s.RegisterService(&Goods_ServiceDesc, srv)
}

func _Goods_GetAllCategoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetAllCategoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/GetAllCategoryList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetAllCategoryList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GetSubCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GetSubCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/GetSubCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GetSubCategory(ctx, req.(*CategoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateCategory(ctx, req.(*CategoryInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/DeleteCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteCategory(ctx, req.(*DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/UpdateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateCategory(ctx, req.(*CategoryInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_BrandList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).BrandList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/BrandList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).BrandList(ctx, req.(*BrandListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/CreateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateBrand(ctx, req.(*BrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_DeleteBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).DeleteBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/DeleteBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).DeleteBrand(ctx, req.(*BrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/UpdateBrand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateBrand(ctx, req.(*BrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateGoodsType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateGoodsType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/CreateGoodsType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateGoodsType(ctx, req.(*GoodsTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateGoodsSpecification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SpecificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateGoodsSpecification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/CreateGoodsSpecification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateGoodsSpecification(ctx, req.(*SpecificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateAttrGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttrGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateAttrGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/CreateAttrGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateAttrGroup(ctx, req.(*AttrGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateAttrValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateAttrValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/CreateAttrValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateAttrValue(ctx, req.(*AttrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_CreateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).CreateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/CreateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).CreateGoods(ctx, req.(*CreateGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_UpdateGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).UpdateGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/UpdateGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).UpdateGoods(ctx, req.(*CreateGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Goods_GoodsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GoodsFilterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodsServer).GoodsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goods.v1.Goods/GoodsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodsServer).GoodsList(ctx, req.(*GoodsFilterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Goods_ServiceDesc is the grpc.ServiceDesc for Goods service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Goods_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goods.v1.Goods",
	HandlerType: (*GoodsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllCategoryList",
			Handler:    _Goods_GetAllCategoryList_Handler,
		},
		{
			MethodName: "GetSubCategory",
			Handler:    _Goods_GetSubCategory_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _Goods_CreateCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _Goods_DeleteCategory_Handler,
		},
		{
			MethodName: "UpdateCategory",
			Handler:    _Goods_UpdateCategory_Handler,
		},
		{
			MethodName: "BrandList",
			Handler:    _Goods_BrandList_Handler,
		},
		{
			MethodName: "CreateBrand",
			Handler:    _Goods_CreateBrand_Handler,
		},
		{
			MethodName: "DeleteBrand",
			Handler:    _Goods_DeleteBrand_Handler,
		},
		{
			MethodName: "UpdateBrand",
			Handler:    _Goods_UpdateBrand_Handler,
		},
		{
			MethodName: "CreateGoodsType",
			Handler:    _Goods_CreateGoodsType_Handler,
		},
		{
			MethodName: "CreateGoodsSpecification",
			Handler:    _Goods_CreateGoodsSpecification_Handler,
		},
		{
			MethodName: "CreateAttrGroup",
			Handler:    _Goods_CreateAttrGroup_Handler,
		},
		{
			MethodName: "CreateAttrValue",
			Handler:    _Goods_CreateAttrValue_Handler,
		},
		{
			MethodName: "CreateGoods",
			Handler:    _Goods_CreateGoods_Handler,
		},
		{
			MethodName: "UpdateGoods",
			Handler:    _Goods_UpdateGoods_Handler,
		},
		{
			MethodName: "GoodsList",
			Handler:    _Goods_GoodsList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/goods/v1/goods.proto",
}
