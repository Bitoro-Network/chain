// This file is @generated by prost-build.
/// RedisOrder is a proto for orders stored in Redis. This proto holds some
/// human-readable values such as price, size and ticker as well as the original
/// `Order` proto from the Bitoro application.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct RedisOrder {
    /// uuid of the Order generated by the Indexer based on the `OrderId`.
    #[prost(string, tag = "1")]
    pub id: ::prost::alloc::string::String,
    /// Order proto from the protocol.
    #[prost(message, optional, tag = "2")]
    pub order: ::core::option::Option<super::protocol::v1::IndexerOrder>,
    /// Ticker for the exchange pair for the order.
    #[prost(string, tag = "3")]
    pub ticker: ::prost::alloc::string::String,
    /// Type of the ticker, PERPETUAL or SPOT.
    #[prost(enumeration = "redis_order::TickerType", tag = "4")]
    pub ticker_type: i32,
    /// Human-readable price of the order.
    #[prost(string, tag = "5")]
    pub price: ::prost::alloc::string::String,
    /// Human-readable size of the order.
    #[prost(string, tag = "6")]
    pub size: ::prost::alloc::string::String,
}
/// Nested message and enum types in `RedisOrder`.
pub mod redis_order {
    /// Enum for the ticker type, PERPETUAL or SPOT.
    #[derive(
        Clone,
        Copy,
        Debug,
        PartialEq,
        Eq,
        Hash,
        PartialOrd,
        Ord,
        ::prost::Enumeration
    )]
    #[repr(i32)]
    pub enum TickerType {
        /// Default value for the enum. Should never be used in an initialized
        /// `RedisOrder`.
        Unspecified = 0,
        /// Ticker is for a perpetual pair.
        Perpetual = 1,
        /// Ticker is for a spot pair.
        Spot = 2,
    }
    impl TickerType {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub fn as_str_name(&self) -> &'static str {
            match self {
                Self::Unspecified => "TICKER_TYPE_UNSPECIFIED",
                Self::Perpetual => "TICKER_TYPE_PERPETUAL",
                Self::Spot => "TICKER_TYPE_SPOT",
            }
        }
        /// Creates an enum from field names used in the ProtoBuf definition.
        pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
            match value {
                "TICKER_TYPE_UNSPECIFIED" => Some(Self::Unspecified),
                "TICKER_TYPE_PERPETUAL" => Some(Self::Perpetual),
                "TICKER_TYPE_SPOT" => Some(Self::Spot),
                _ => None,
            }
        }
    }
}
impl ::prost::Name for RedisOrder {
    const NAME: &'static str = "RedisOrder";
    const PACKAGE: &'static str = "bitoroprotocol.indexer.redis";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.indexer.redis.RedisOrder".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.indexer.redis.RedisOrder".into()
    }
}
