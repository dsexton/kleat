// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: security/ssl.proto

package security

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Setting for client authentication.
type ClientAuth int32

const (
	// No client authentication.
	ClientAuth_NONE ClientAuth = 0
	// Client authentication is optional.
	ClientAuth_WANT ClientAuth = 1
	// Client authentication is required.
	ClientAuth_NEED ClientAuth = 2
)

// Enum value maps for ClientAuth.
var (
	ClientAuth_name = map[int32]string{
		0: "NONE",
		1: "WANT",
		2: "NEED",
	}
	ClientAuth_value = map[string]int32{
		"NONE": 0,
		"WANT": 1,
		"NEED": 2,
	}
)

func (x ClientAuth) Enum() *ClientAuth {
	p := new(ClientAuth)
	*p = x
	return p
}

func (x ClientAuth) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientAuth) Descriptor() protoreflect.EnumDescriptor {
	return file_security_ssl_proto_enumTypes[0].Descriptor()
}

func (ClientAuth) Type() protoreflect.EnumType {
	return &file_security_ssl_proto_enumTypes[0]
}

func (x ClientAuth) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientAuth.Descriptor instead.
func (ClientAuth) EnumDescriptor() ([]byte, []int) {
	return file_security_ssl_proto_rawDescGZIP(), []int{0}
}

// Configuration for the API server's addressable URL and CORS policies.
type ApiSecurity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If you have authentication enabled, are accessing Spinnaker remotely, and
	// are logging in from sources other than the UI, provide a regex matching all
	// URLs authentication redirects may come from.
	CorsAccessPattern string `protobuf:"bytes,1,opt,name=corsAccessPattern,proto3" json:"corsAccessPattern,omitempty"`
	// If you want the API server to do SSL termination, it must be enabled and
	// configured here. If you are doing your own SSL termination, leave this disabled.
	Ssl *ApiSsl `protobuf:"bytes,2,opt,name=ssl,proto3" json:"ssl,omitempty"`
	// If you are accessing the API server remotely, provide the full base URL of
	// whatever proxy or load balancer is fronting the API requests
	OverrideBaseUrl string `protobuf:"bytes,3,opt,name=overrideBaseUrl,proto3" json:"overrideBaseUrl,omitempty"`
}

func (x *ApiSecurity) Reset() {
	*x = ApiSecurity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_ssl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiSecurity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiSecurity) ProtoMessage() {}

func (x *ApiSecurity) ProtoReflect() protoreflect.Message {
	mi := &file_security_ssl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiSecurity.ProtoReflect.Descriptor instead.
func (*ApiSecurity) Descriptor() ([]byte, []int) {
	return file_security_ssl_proto_rawDescGZIP(), []int{0}
}

func (x *ApiSecurity) GetCorsAccessPattern() string {
	if x != nil {
		return x.CorsAccessPattern
	}
	return ""
}

func (x *ApiSecurity) GetSsl() *ApiSsl {
	if x != nil {
		return x.Ssl
	}
	return nil
}

func (x *ApiSecurity) GetOverrideBaseUrl() string {
	if x != nil {
		return x.OverrideBaseUrl
	}
	return ""
}

// Configuration for the UI server's addressable URL.
type UiSecurity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for SSL termination by the UI gateway.
	Ssl *UiSsl `protobuf:"bytes,1,opt,name=ssl,proto3" json:"ssl,omitempty"`
	// If you are accessing the UI server remotely, provide the full base URL of
	// whatever proxy or load balancer is fronting the UI requests.
	OverrideBaseUrl string `protobuf:"bytes,2,opt,name=overrideBaseUrl,proto3" json:"overrideBaseUrl,omitempty"`
}

func (x *UiSecurity) Reset() {
	*x = UiSecurity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_ssl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UiSecurity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UiSecurity) ProtoMessage() {}

func (x *UiSecurity) ProtoReflect() protoreflect.Message {
	mi := &file_security_ssl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UiSecurity.ProtoReflect.Descriptor instead.
func (*UiSecurity) Descriptor() ([]byte, []int) {
	return file_security_ssl_proto_rawDescGZIP(), []int{1}
}

func (x *UiSecurity) GetSsl() *UiSsl {
	if x != nil {
		return x.Ssl
	}
	return nil
}

func (x *UiSecurity) GetOverrideBaseUrl() string {
	if x != nil {
		return x.OverrideBaseUrl
	}
	return ""
}

// Configuration for SSL termination by the API server.
type ApiSsl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether SSL is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Name of your keystore entry as generated with your keytool.
	KeyAlias string `protobuf:"bytes,2,opt,name=keyAlias,proto3" json:"keyAlias,omitempty"`
	// Path to the keystore holding your security certificates.
	KeyStore string `protobuf:"bytes,3,opt,name=keyStore,proto3" json:"keyStore,omitempty"`
	// The type of your keystore. Examples include JKS, and PKCS12.
	KeyStoreType string `protobuf:"bytes,4,opt,name=keyStoreType,proto3" json:"keyStoreType,omitempty"`
	// The password to unlock your keystore. Due to a limitation in Tomcat, this
	// must match your key's password in the keystore.
	KeyStorePassword string `protobuf:"bytes,5,opt,name=keyStorePassword,proto3" json:"keyStorePassword,omitempty"`
	// Path to the truststore holding your trusted certificates.
	TrustStore string `protobuf:"bytes,6,opt,name=trustStore,proto3" json:"trustStore,omitempty"`
	// The type of your truststore. Examples include JKS, and PKCS12.
	TrustStoreType string `protobuf:"bytes,7,opt,name=trustStoreType,proto3" json:"trustStoreType,omitempty"`
	// The password to unlock your truststore.
	TrustStorePassword string `protobuf:"bytes,8,opt,name=trustStorePassword,proto3" json:"trustStorePassword,omitempty"`
	// Whether to require or allow client authentication.
	ClientAuth ClientAuth `protobuf:"varint,9,opt,name=clientAuth,proto3,enum=proto.security.ClientAuth" json:"clientAuth,omitempty"`
}

func (x *ApiSsl) Reset() {
	*x = ApiSsl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_ssl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiSsl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiSsl) ProtoMessage() {}

func (x *ApiSsl) ProtoReflect() protoreflect.Message {
	mi := &file_security_ssl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiSsl.ProtoReflect.Descriptor instead.
func (*ApiSsl) Descriptor() ([]byte, []int) {
	return file_security_ssl_proto_rawDescGZIP(), []int{2}
}

func (x *ApiSsl) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *ApiSsl) GetKeyAlias() string {
	if x != nil {
		return x.KeyAlias
	}
	return ""
}

func (x *ApiSsl) GetKeyStore() string {
	if x != nil {
		return x.KeyStore
	}
	return ""
}

func (x *ApiSsl) GetKeyStoreType() string {
	if x != nil {
		return x.KeyStoreType
	}
	return ""
}

func (x *ApiSsl) GetKeyStorePassword() string {
	if x != nil {
		return x.KeyStorePassword
	}
	return ""
}

func (x *ApiSsl) GetTrustStore() string {
	if x != nil {
		return x.TrustStore
	}
	return ""
}

func (x *ApiSsl) GetTrustStoreType() string {
	if x != nil {
		return x.TrustStoreType
	}
	return ""
}

func (x *ApiSsl) GetTrustStorePassword() string {
	if x != nil {
		return x.TrustStorePassword
	}
	return ""
}

func (x *ApiSsl) GetClientAuth() ClientAuth {
	if x != nil {
		return x.ClientAuth
	}
	return ClientAuth_NONE
}

// Configuration for SSL termination by the UI gateway.
type UiSsl struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether SSL is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Path to your .crt file.
	SslCertificateFile string `protobuf:"bytes,2,opt,name=sslCertificateFile,proto3" json:"sslCertificateFile,omitempty"`
	// Path to your .key file.
	SslCertificateKeyFile string `protobuf:"bytes,3,opt,name=sslCertificateKeyFile,proto3" json:"sslCertificateKeyFile,omitempty"`
	// Path to the .crt file for the CA that issued your SSL certificate. This is
	// only needed for local git deployments that serve the UI using webpack dev server.
	SslCACertificateFile string `protobuf:"bytes,4,opt,name=sslCACertificateFile,proto3" json:"sslCACertificateFile,omitempty"`
	// The passphrase needed to unlock your SSL certificate. This will be provided
	// to Apache on startup.
	SslCertificatePassphrase string `protobuf:"bytes,5,opt,name=sslCertificatePassphrase,proto3" json:"sslCertificatePassphrase,omitempty"`
}

func (x *UiSsl) Reset() {
	*x = UiSsl{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_ssl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UiSsl) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UiSsl) ProtoMessage() {}

func (x *UiSsl) ProtoReflect() protoreflect.Message {
	mi := &file_security_ssl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UiSsl.ProtoReflect.Descriptor instead.
func (*UiSsl) Descriptor() ([]byte, []int) {
	return file_security_ssl_proto_rawDescGZIP(), []int{3}
}

func (x *UiSsl) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *UiSsl) GetSslCertificateFile() string {
	if x != nil {
		return x.SslCertificateFile
	}
	return ""
}

func (x *UiSsl) GetSslCertificateKeyFile() string {
	if x != nil {
		return x.SslCertificateKeyFile
	}
	return ""
}

func (x *UiSsl) GetSslCACertificateFile() string {
	if x != nil {
		return x.SslCACertificateFile
	}
	return ""
}

func (x *UiSsl) GetSslCertificatePassphrase() string {
	if x != nil {
		return x.SslCertificatePassphrase
	}
	return ""
}

var File_security_ssl_proto protoreflect.FileDescriptor

var file_security_ssl_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x73, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x22, 0x8f, 0x01, 0x0a, 0x0b, 0x41, 0x70, 0x69, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6f, 0x72, 0x73, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x63, 0x6f, 0x72, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x12, 0x28, 0x0a, 0x03, 0x73, 0x73, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x2e, 0x41, 0x70, 0x69, 0x53, 0x73, 0x6c, 0x52, 0x03, 0x73, 0x73, 0x6c, 0x12, 0x28, 0x0a, 0x0f,
	0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x42, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x42,
	0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x5f, 0x0a, 0x0a, 0x55, 0x69, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x03, 0x73, 0x73, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x55, 0x69, 0x53, 0x73, 0x6c, 0x52, 0x03, 0x73, 0x73, 0x6c, 0x12, 0x28, 0x0a,
	0x0f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x42, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x42, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xde, 0x02, 0x0a, 0x06, 0x41, 0x70, 0x69, 0x53,
	0x73, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6b, 0x65, 0x79, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6b, 0x65, 0x79, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6b, 0x65, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x65, 0x79, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6b, 0x65, 0x79, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x6b, 0x65, 0x79, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x72, 0x75, 0x73, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x72, 0x75, 0x73, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x72,
	0x75, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x12,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x74, 0x72, 0x75, 0x73, 0x74, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x3a, 0x0a, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x52, 0x0a, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x22, 0xf7, 0x01, 0x0a, 0x05, 0x55, 0x69, 0x53,
	0x73, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x12,
	0x73, 0x73, 0x6c, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x73, 0x6c, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x15,
	0x73, 0x73, 0x6c, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x73, 0x6c,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x73, 0x73, 0x6c, 0x43, 0x41, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x73, 0x73, 0x6c, 0x43, 0x41, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x3a, 0x0a, 0x18, 0x73, 0x73, 0x6c, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61,
	0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x73, 0x73, 0x6c, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61,
	0x73, 0x65, 0x2a, 0x2a, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41,
	0x4e, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x45, 0x45, 0x44, 0x10, 0x02, 0x42, 0x30,
	0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69,
	0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_security_ssl_proto_rawDescOnce sync.Once
	file_security_ssl_proto_rawDescData = file_security_ssl_proto_rawDesc
)

func file_security_ssl_proto_rawDescGZIP() []byte {
	file_security_ssl_proto_rawDescOnce.Do(func() {
		file_security_ssl_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_ssl_proto_rawDescData)
	})
	return file_security_ssl_proto_rawDescData
}

var file_security_ssl_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_security_ssl_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_security_ssl_proto_goTypes = []interface{}{
	(ClientAuth)(0),     // 0: proto.security.ClientAuth
	(*ApiSecurity)(nil), // 1: proto.security.ApiSecurity
	(*UiSecurity)(nil),  // 2: proto.security.UiSecurity
	(*ApiSsl)(nil),      // 3: proto.security.ApiSsl
	(*UiSsl)(nil),       // 4: proto.security.UiSsl
}
var file_security_ssl_proto_depIdxs = []int32{
	3, // 0: proto.security.ApiSecurity.ssl:type_name -> proto.security.ApiSsl
	4, // 1: proto.security.UiSecurity.ssl:type_name -> proto.security.UiSsl
	0, // 2: proto.security.ApiSsl.clientAuth:type_name -> proto.security.ClientAuth
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_security_ssl_proto_init() }
func file_security_ssl_proto_init() {
	if File_security_ssl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_security_ssl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiSecurity); i {
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
		file_security_ssl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UiSecurity); i {
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
		file_security_ssl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiSsl); i {
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
		file_security_ssl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UiSsl); i {
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
			RawDescriptor: file_security_ssl_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_security_ssl_proto_goTypes,
		DependencyIndexes: file_security_ssl_proto_depIdxs,
		EnumInfos:         file_security_ssl_proto_enumTypes,
		MessageInfos:      file_security_ssl_proto_msgTypes,
	}.Build()
	File_security_ssl_proto = out.File
	file_security_ssl_proto_rawDesc = nil
	file_security_ssl_proto_goTypes = nil
	file_security_ssl_proto_depIdxs = nil
}
