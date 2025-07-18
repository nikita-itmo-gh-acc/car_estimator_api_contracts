// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: profile.proto

package profile_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type UUID struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UUID) Reset() {
	*x = UUID{}
	mi := &file_profile_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{0}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SourceData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ip            string                 `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	UserAgent     string                 `protobuf:"bytes,2,opt,name=userAgent,proto3" json:"userAgent,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SourceData) Reset() {
	*x = SourceData{}
	mi := &file_profile_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SourceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceData) ProtoMessage() {}

func (x *SourceData) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceData.ProtoReflect.Descriptor instead.
func (*SourceData) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{1}
}

func (x *SourceData) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *SourceData) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

type RegiserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        *UUID                  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegiserRequest) Reset() {
	*x = RegiserRequest{}
	mi := &file_profile_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegiserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegiserRequest) ProtoMessage() {}

func (x *RegiserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegiserRequest.ProtoReflect.Descriptor instead.
func (*RegiserRequest) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{2}
}

func (x *RegiserRequest) GetUserId() *UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

type TokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TokenResponse) Reset() {
	*x = TokenResponse{}
	mi := &file_profile_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenResponse) ProtoMessage() {}

func (x *TokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenResponse.ProtoReflect.Descriptor instead.
func (*TokenResponse) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{3}
}

func (x *TokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *TokenResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Source        *SourceData            `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_profile_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginRequest) GetSource() *SourceData {
	if x != nil {
		return x.Source
	}
	return nil
}

type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        *UUID                  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Tokens        *TokenResponse         `protobuf:"bytes,2,opt,name=tokens,proto3" json:"tokens,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_profile_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResponse) GetUserId() *UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *LoginResponse) GetTokens() *TokenResponse {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type RegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        *UUID                  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_profile_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterResponse) GetUserId() *UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

type RegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Fullname      string                 `protobuf:"bytes,1,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Password      string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Birthdate     int64                  `protobuf:"varint,5,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_profile_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterRequest) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *RegisterRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetBirthdate() int64 {
	if x != nil {
		return x.Birthdate
	}
	return 0
}

type UserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        *UUID                  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_profile_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{8}
}

func (x *UserRequest) GetUserId() *UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

type UserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        *UUID                  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Fullname      string                 `protobuf:"bytes,2,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone         string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Birthdate     int64                  `protobuf:"varint,5,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	Registerdate  int64                  `protobuf:"varint,6,opt,name=registerdate,proto3" json:"registerdate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	mi := &file_profile_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_profile_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_profile_proto_rawDescGZIP(), []int{9}
}

func (x *UserResponse) GetUserId() *UUID {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *UserResponse) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *UserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserResponse) GetBirthdate() int64 {
	if x != nil {
		return x.Birthdate
	}
	return 0
}

func (x *UserResponse) GetRegisterdate() int64 {
	if x != nil {
		return x.Registerdate
	}
	return 0
}

var File_profile_proto protoreflect.FileDescriptor

const file_profile_proto_rawDesc = "" +
	"\n" +
	"\rprofile.proto\x12\aprofile\x1a\x1bgoogle/protobuf/empty.proto\"\x1c\n" +
	"\x04UUID\x12\x14\n" +
	"\x05value\x18\x01 \x01(\tR\x05value\":\n" +
	"\n" +
	"SourceData\x12\x0e\n" +
	"\x02ip\x18\x01 \x01(\tR\x02ip\x12\x1c\n" +
	"\tuserAgent\x18\x02 \x01(\tR\tuserAgent\"7\n" +
	"\x0eRegiserRequest\x12%\n" +
	"\x06userId\x18\x01 \x01(\v2\r.profile.UUIDR\x06userId\"U\n" +
	"\rTokenResponse\x12 \n" +
	"\vaccessToken\x18\x01 \x01(\tR\vaccessToken\x12\"\n" +
	"\frefreshToken\x18\x02 \x01(\tR\frefreshToken\"m\n" +
	"\fLoginRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\x12+\n" +
	"\x06source\x18\x03 \x01(\v2\x13.profile.SourceDataR\x06source\"f\n" +
	"\rLoginResponse\x12%\n" +
	"\x06userId\x18\x01 \x01(\v2\r.profile.UUIDR\x06userId\x12.\n" +
	"\x06tokens\x18\x02 \x01(\v2\x16.profile.TokenResponseR\x06tokens\"9\n" +
	"\x10RegisterResponse\x12%\n" +
	"\x06userId\x18\x01 \x01(\v2\r.profile.UUIDR\x06userId\"\x93\x01\n" +
	"\x0fRegisterRequest\x12\x1a\n" +
	"\bfullname\x18\x01 \x01(\tR\bfullname\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x03 \x01(\tR\x05phone\x12\x1a\n" +
	"\bpassword\x18\x04 \x01(\tR\bpassword\x12\x1c\n" +
	"\tbirthdate\x18\x05 \x01(\x03R\tbirthdate\"4\n" +
	"\vUserRequest\x12%\n" +
	"\x06userId\x18\x01 \x01(\v2\r.profile.UUIDR\x06userId\"\xbf\x01\n" +
	"\fUserResponse\x12%\n" +
	"\x06userId\x18\x01 \x01(\v2\r.profile.UUIDR\x06userId\x12\x1a\n" +
	"\bfullname\x18\x02 \x01(\tR\bfullname\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x04 \x01(\tR\x05phone\x12\x1c\n" +
	"\tbirthdate\x18\x05 \x01(\x03R\tbirthdate\x12\"\n" +
	"\fregisterdate\x18\x06 \x01(\x03R\fregisterdate2\xfd\x02\n" +
	"\x0eProfileService\x12A\n" +
	"\bRegister\x12\x18.profile.RegisterRequest\x1a\x19.profile.RegisterResponse\"\x00\x12>\n" +
	"\n" +
	"Unregister\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\"\x00\x128\n" +
	"\aGetUser\x12\x14.profile.UserRequest\x1a\x15.profile.UserResponse\"\x00\x128\n" +
	"\x05Login\x12\x15.profile.LoginRequest\x1a\x16.profile.LoginResponse\"\x00\x12:\n" +
	"\x06Logout\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\"\x00\x128\n" +
	"\aRefresh\x12\x13.profile.SourceData\x1a\x16.profile.TokenResponse\"\x00B\fZ\n" +
	"profile.v1b\x06proto3"

var (
	file_profile_proto_rawDescOnce sync.Once
	file_profile_proto_rawDescData []byte
)

func file_profile_proto_rawDescGZIP() []byte {
	file_profile_proto_rawDescOnce.Do(func() {
		file_profile_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_profile_proto_rawDesc), len(file_profile_proto_rawDesc)))
	})
	return file_profile_proto_rawDescData
}

var file_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_profile_proto_goTypes = []any{
	(*UUID)(nil),             // 0: profile.UUID
	(*SourceData)(nil),       // 1: profile.SourceData
	(*RegiserRequest)(nil),   // 2: profile.RegiserRequest
	(*TokenResponse)(nil),    // 3: profile.TokenResponse
	(*LoginRequest)(nil),     // 4: profile.LoginRequest
	(*LoginResponse)(nil),    // 5: profile.LoginResponse
	(*RegisterResponse)(nil), // 6: profile.RegisterResponse
	(*RegisterRequest)(nil),  // 7: profile.RegisterRequest
	(*UserRequest)(nil),      // 8: profile.UserRequest
	(*UserResponse)(nil),     // 9: profile.UserResponse
	(*emptypb.Empty)(nil),    // 10: google.protobuf.Empty
}
var file_profile_proto_depIdxs = []int32{
	0,  // 0: profile.RegiserRequest.userId:type_name -> profile.UUID
	1,  // 1: profile.LoginRequest.source:type_name -> profile.SourceData
	0,  // 2: profile.LoginResponse.userId:type_name -> profile.UUID
	3,  // 3: profile.LoginResponse.tokens:type_name -> profile.TokenResponse
	0,  // 4: profile.RegisterResponse.userId:type_name -> profile.UUID
	0,  // 5: profile.UserRequest.userId:type_name -> profile.UUID
	0,  // 6: profile.UserResponse.userId:type_name -> profile.UUID
	7,  // 7: profile.ProfileService.Register:input_type -> profile.RegisterRequest
	10, // 8: profile.ProfileService.Unregister:input_type -> google.protobuf.Empty
	8,  // 9: profile.ProfileService.GetUser:input_type -> profile.UserRequest
	4,  // 10: profile.ProfileService.Login:input_type -> profile.LoginRequest
	10, // 11: profile.ProfileService.Logout:input_type -> google.protobuf.Empty
	1,  // 12: profile.ProfileService.Refresh:input_type -> profile.SourceData
	6,  // 13: profile.ProfileService.Register:output_type -> profile.RegisterResponse
	10, // 14: profile.ProfileService.Unregister:output_type -> google.protobuf.Empty
	9,  // 15: profile.ProfileService.GetUser:output_type -> profile.UserResponse
	5,  // 16: profile.ProfileService.Login:output_type -> profile.LoginResponse
	10, // 17: profile.ProfileService.Logout:output_type -> google.protobuf.Empty
	3,  // 18: profile.ProfileService.Refresh:output_type -> profile.TokenResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_profile_proto_init() }
func file_profile_proto_init() {
	if File_profile_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_profile_proto_rawDesc), len(file_profile_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_profile_proto_goTypes,
		DependencyIndexes: file_profile_proto_depIdxs,
		MessageInfos:      file_profile_proto_msgTypes,
	}.Build()
	File_profile_proto = out.File
	file_profile_proto_goTypes = nil
	file_profile_proto_depIdxs = nil
}
