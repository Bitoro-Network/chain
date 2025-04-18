// This file is @generated by prost-build.
/// Account State
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AccountState {
    #[prost(string, tag = "1")]
    pub address: ::prost::alloc::string::String,
    #[prost(message, optional, tag = "2")]
    pub timestamp_nonce_details: ::core::option::Option<TimestampNonceDetails>,
}
impl ::prost::Name for AccountState {
    const NAME: &'static str = "AccountState";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.AccountState".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.AccountState".into()
    }
}
/// Timestamp nonce details
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct TimestampNonceDetails {
    /// unsorted list of n most recent timestamp nonces
    #[prost(uint64, repeated, tag = "1")]
    pub timestamp_nonces: ::prost::alloc::vec::Vec<u64>,
    /// max timestamp nonce that was ejected from list above
    #[prost(uint64, tag = "2")]
    pub max_ejected_nonce: u64,
}
impl ::prost::Name for TimestampNonceDetails {
    const NAME: &'static str = "TimestampNonceDetails";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.TimestampNonceDetails".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.TimestampNonceDetails".into()
    }
}
/// AccountAuthenticator represents a foundational model for all authenticators.
/// It provides extensibility by allowing concrete types to interpret and
/// validate transactions based on the encapsulated data.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AccountAuthenticator {
    /// ID uniquely identifies the authenticator instance.
    #[prost(uint64, tag = "1")]
    pub id: u64,
    /// Type specifies the category of the AccountAuthenticator.
    /// This type information is essential for differentiating authenticators
    /// and ensuring precise data retrieval from the storage layer.
    #[prost(string, tag = "2")]
    pub r#type: ::prost::alloc::string::String,
    /// Config is a versatile field used in conjunction with the specific type of
    /// account authenticator to facilitate complex authentication processes.
    /// The interpretation of this field is overloaded, enabling multiple
    /// authenticators to utilize it for their respective purposes.
    #[prost(bytes = "vec", tag = "3")]
    pub config: ::prost::alloc::vec::Vec<u8>,
}
impl ::prost::Name for AccountAuthenticator {
    const NAME: &'static str = "AccountAuthenticator";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.AccountAuthenticator".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.AccountAuthenticator".into()
    }
}
/// Params defines the parameters for the module.
#[derive(Clone, Copy, PartialEq, ::prost::Message)]
pub struct Params {
    /// IsSmartAccountActive defines the state of the authenticator.
    /// If set to false, the authenticator module will not be used
    /// and the classic cosmos sdk authentication will be used instead.
    #[prost(bool, tag = "1")]
    pub is_smart_account_active: bool,
}
impl ::prost::Name for Params {
    const NAME: &'static str = "Params";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.Params".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.Params".into()
    }
}
/// AuthenticatorData represents a genesis exported account with Authenticators.
/// The address is used as the key, and the account authenticators are stored in
/// the authenticators field.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct AuthenticatorData {
    /// address is an account address, one address can have many authenticators
    #[prost(string, tag = "1")]
    pub address: ::prost::alloc::string::String,
    /// authenticators are the account's authenticators, these can be multiple
    /// types including SignatureVerification, AllOfs, CosmWasmAuthenticators, etc
    #[prost(message, repeated, tag = "2")]
    pub authenticators: ::prost::alloc::vec::Vec<AccountAuthenticator>,
}
impl ::prost::Name for AuthenticatorData {
    const NAME: &'static str = "AuthenticatorData";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.AuthenticatorData".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.AuthenticatorData".into()
    }
}
/// Module genesis state
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GenesisState {
    #[prost(message, repeated, tag = "1")]
    pub accounts: ::prost::alloc::vec::Vec<AccountState>,
    /// params define the parameters for the authenticator module.
    #[prost(message, optional, tag = "2")]
    pub params: ::core::option::Option<Params>,
    /// next_authenticator_id is the next available authenticator ID.
    #[prost(uint64, tag = "3")]
    pub next_authenticator_id: u64,
    /// authenticator_data contains the data for multiple accounts, each with their
    /// authenticators.
    #[prost(message, repeated, tag = "4")]
    pub authenticator_data: ::prost::alloc::vec::Vec<AuthenticatorData>,
}
impl ::prost::Name for GenesisState {
    const NAME: &'static str = "GenesisState";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.GenesisState".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.GenesisState".into()
    }
}
/// QueryParamsRequest is request type for the Query/Params RPC method.
#[derive(Clone, Copy, PartialEq, ::prost::Message)]
pub struct QueryParamsRequest {}
impl ::prost::Name for QueryParamsRequest {
    const NAME: &'static str = "QueryParamsRequest";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.QueryParamsRequest".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.QueryParamsRequest".into()
    }
}
/// QueryParamsResponse is response type for the Query/Params RPC method.
#[derive(Clone, Copy, PartialEq, ::prost::Message)]
pub struct QueryParamsResponse {
    /// params holds all the parameters of this module.
    #[prost(message, optional, tag = "1")]
    pub params: ::core::option::Option<Params>,
}
impl ::prost::Name for QueryParamsResponse {
    const NAME: &'static str = "QueryParamsResponse";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.QueryParamsResponse".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.QueryParamsResponse".into()
    }
}
/// MsgGetAuthenticatorsRequest defines the Msg/GetAuthenticators request type.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetAuthenticatorsRequest {
    #[prost(string, tag = "1")]
    pub account: ::prost::alloc::string::String,
}
impl ::prost::Name for GetAuthenticatorsRequest {
    const NAME: &'static str = "GetAuthenticatorsRequest";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.GetAuthenticatorsRequest".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.GetAuthenticatorsRequest".into()
    }
}
/// MsgGetAuthenticatorsResponse defines the Msg/GetAuthenticators response type.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetAuthenticatorsResponse {
    #[prost(message, repeated, tag = "1")]
    pub account_authenticators: ::prost::alloc::vec::Vec<AccountAuthenticator>,
}
impl ::prost::Name for GetAuthenticatorsResponse {
    const NAME: &'static str = "GetAuthenticatorsResponse";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.GetAuthenticatorsResponse".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.GetAuthenticatorsResponse".into()
    }
}
/// MsgGetAuthenticatorRequest defines the Msg/GetAuthenticator request type.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetAuthenticatorRequest {
    #[prost(string, tag = "1")]
    pub account: ::prost::alloc::string::String,
    #[prost(uint64, tag = "2")]
    pub authenticator_id: u64,
}
impl ::prost::Name for GetAuthenticatorRequest {
    const NAME: &'static str = "GetAuthenticatorRequest";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.GetAuthenticatorRequest".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.GetAuthenticatorRequest".into()
    }
}
/// MsgGetAuthenticatorResponse defines the Msg/GetAuthenticator response type.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct GetAuthenticatorResponse {
    #[prost(message, optional, tag = "1")]
    pub account_authenticator: ::core::option::Option<AccountAuthenticator>,
}
impl ::prost::Name for GetAuthenticatorResponse {
    const NAME: &'static str = "GetAuthenticatorResponse";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.GetAuthenticatorResponse".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.GetAuthenticatorResponse".into()
    }
}
/// Generated client implementations.
pub mod query_client {
    #![allow(
        unused_variables,
        dead_code,
        missing_docs,
        clippy::wildcard_imports,
        clippy::let_unit_value,
    )]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    /// Query defines the gRPC querier service.
    #[derive(Debug, Clone)]
    pub struct QueryClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    #[cfg(feature = "grpc-transport")]
    impl QueryClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> QueryClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + std::marker::Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + std::marker::Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> QueryClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + std::marker::Send + std::marker::Sync,
        {
            QueryClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        /// Parameters queries the parameters of the module.
        pub async fn params(
            &mut self,
            request: impl tonic::IntoRequest<super::QueryParamsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::QueryParamsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/bitoroprotocol.accountplus.Query/Params",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("bitoroprotocol.accountplus.Query", "Params"));
            self.inner.unary(req, path, codec).await
        }
        /// Queries a single authenticator by account and authenticator ID.
        pub async fn get_authenticator(
            &mut self,
            request: impl tonic::IntoRequest<super::GetAuthenticatorRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetAuthenticatorResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/bitoroprotocol.accountplus.Query/GetAuthenticator",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("bitoroprotocol.accountplus.Query", "GetAuthenticator"),
                );
            self.inner.unary(req, path, codec).await
        }
        /// Queries all authenticators for a given account.
        pub async fn get_authenticators(
            &mut self,
            request: impl tonic::IntoRequest<super::GetAuthenticatorsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetAuthenticatorsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/bitoroprotocol.accountplus.Query/GetAuthenticators",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new(
                        "bitoroprotocol.accountplus.Query",
                        "GetAuthenticators",
                    ),
                );
            self.inner.unary(req, path, codec).await
        }
    }
}
/// MsgAddAuthenticatorRequest defines the Msg/AddAuthenticator request type.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgAddAuthenticator {
    #[prost(string, tag = "1")]
    pub sender: ::prost::alloc::string::String,
    #[prost(string, tag = "2")]
    pub authenticator_type: ::prost::alloc::string::String,
    #[prost(bytes = "vec", tag = "3")]
    pub data: ::prost::alloc::vec::Vec<u8>,
}
impl ::prost::Name for MsgAddAuthenticator {
    const NAME: &'static str = "MsgAddAuthenticator";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.MsgAddAuthenticator".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.MsgAddAuthenticator".into()
    }
}
/// MsgAddAuthenticatorResponse defines the Msg/AddAuthenticator response type.
#[derive(Clone, Copy, PartialEq, ::prost::Message)]
pub struct MsgAddAuthenticatorResponse {
    #[prost(bool, tag = "1")]
    pub success: bool,
}
impl ::prost::Name for MsgAddAuthenticatorResponse {
    const NAME: &'static str = "MsgAddAuthenticatorResponse";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.MsgAddAuthenticatorResponse".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.MsgAddAuthenticatorResponse".into()
    }
}
/// MsgRemoveAuthenticatorRequest defines the Msg/RemoveAuthenticator request
/// type.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgRemoveAuthenticator {
    #[prost(string, tag = "1")]
    pub sender: ::prost::alloc::string::String,
    #[prost(uint64, tag = "2")]
    pub id: u64,
}
impl ::prost::Name for MsgRemoveAuthenticator {
    const NAME: &'static str = "MsgRemoveAuthenticator";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.MsgRemoveAuthenticator".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.MsgRemoveAuthenticator".into()
    }
}
/// MsgRemoveAuthenticatorResponse defines the Msg/RemoveAuthenticator response
/// type.
#[derive(Clone, Copy, PartialEq, ::prost::Message)]
pub struct MsgRemoveAuthenticatorResponse {
    #[prost(bool, tag = "1")]
    pub success: bool,
}
impl ::prost::Name for MsgRemoveAuthenticatorResponse {
    const NAME: &'static str = "MsgRemoveAuthenticatorResponse";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.MsgRemoveAuthenticatorResponse".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.MsgRemoveAuthenticatorResponse".into()
    }
}
/// MsgSetActiveState sets the active state of the module.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct MsgSetActiveState {
    /// Authority is the address that may send this message.
    #[prost(string, tag = "1")]
    pub authority: ::prost::alloc::string::String,
    #[prost(bool, tag = "2")]
    pub active: bool,
}
impl ::prost::Name for MsgSetActiveState {
    const NAME: &'static str = "MsgSetActiveState";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.MsgSetActiveState".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.MsgSetActiveState".into()
    }
}
/// MsgSetActiveStateResponse defines the Msg/SetActiveState response type.
#[derive(Clone, Copy, PartialEq, ::prost::Message)]
pub struct MsgSetActiveStateResponse {}
impl ::prost::Name for MsgSetActiveStateResponse {
    const NAME: &'static str = "MsgSetActiveStateResponse";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.MsgSetActiveStateResponse".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.MsgSetActiveStateResponse".into()
    }
}
/// TxExtension allows for additional authenticator-specific data in
/// transactions.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct TxExtension {
    /// selected_authenticators holds the authenticator_id for the chosen
    /// authenticator per message.
    #[prost(uint64, repeated, tag = "1")]
    pub selected_authenticators: ::prost::alloc::vec::Vec<u64>,
}
impl ::prost::Name for TxExtension {
    const NAME: &'static str = "TxExtension";
    const PACKAGE: &'static str = "bitoroprotocol.accountplus";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.accountplus.TxExtension".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.accountplus.TxExtension".into()
    }
}
/// Generated client implementations.
pub mod msg_client {
    #![allow(
        unused_variables,
        dead_code,
        missing_docs,
        clippy::wildcard_imports,
        clippy::let_unit_value,
    )]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    /// Msg defines the Msg service.
    #[derive(Debug, Clone)]
    pub struct MsgClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    #[cfg(feature = "grpc-transport")]
    impl MsgClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> MsgClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + std::marker::Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + std::marker::Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> MsgClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + std::marker::Send + std::marker::Sync,
        {
            MsgClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        /// AddAuthenticator adds an authenticator to an account.
        pub async fn add_authenticator(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgAddAuthenticator>,
        ) -> std::result::Result<
            tonic::Response<super::MsgAddAuthenticatorResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/bitoroprotocol.accountplus.Msg/AddAuthenticator",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("bitoroprotocol.accountplus.Msg", "AddAuthenticator"),
                );
            self.inner.unary(req, path, codec).await
        }
        /// RemoveAuthenticator removes an authenticator from an account.
        pub async fn remove_authenticator(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgRemoveAuthenticator>,
        ) -> std::result::Result<
            tonic::Response<super::MsgRemoveAuthenticatorResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/bitoroprotocol.accountplus.Msg/RemoveAuthenticator",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new(
                        "bitoroprotocol.accountplus.Msg",
                        "RemoveAuthenticator",
                    ),
                );
            self.inner.unary(req, path, codec).await
        }
        /// SetActiveState sets the active state of the authenticator.
        /// Primarily used for circuit breaking.
        pub async fn set_active_state(
            &mut self,
            request: impl tonic::IntoRequest<super::MsgSetActiveState>,
        ) -> std::result::Result<
            tonic::Response<super::MsgSetActiveStateResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::unknown(
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/bitoroprotocol.accountplus.Msg/SetActiveState",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("bitoroprotocol.accountplus.Msg", "SetActiveState"),
                );
            self.inner.unary(req, path, codec).await
        }
    }
}
