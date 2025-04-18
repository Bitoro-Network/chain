// This file is @generated by prost-build.
/// OrderPlace messages contain the order placed/replaced.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OrderPlaceV1 {
    #[prost(message, optional, tag = "1")]
    pub order: ::core::option::Option<super::protocol::v1::IndexerOrder>,
    #[prost(enumeration = "order_place_v1::OrderPlacementStatus", tag = "2")]
    pub placement_status: i32,
    /// The timestamp of the order placement.
    #[prost(message, optional, tag = "3")]
    pub time_stamp: ::core::option::Option<::prost_types::Timestamp>,
}
/// Nested message and enum types in `OrderPlaceV1`.
pub mod order_place_v1 {
    /// OrderPlacementStatus is an enum for the resulting status after an order is
    /// placed.
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
    pub enum OrderPlacementStatus {
        /// Default value, this is invalid and unused.
        Unspecified = 0,
        /// A best effort opened order is one that has only been confirmed to be
        /// placed on the Bitoro node sending the off-chain update message.
        /// The cases where this happens includes:
        /// - The Bitoro node places an order in it's in-memory orderbook during the
        ///    CheckTx flow.
        /// A best effort placed order may not have been placed on other Bitoro
        /// nodes including other Bitoro validator nodes and may still be excluded in
        /// future order matches.
        BestEffortOpened = 1,
        /// An opened order is one that is confirmed to be placed on all Bitoro nodes
        /// (discounting dishonest Bitoro nodes) and will be included in any future
        /// order matches.
        /// This status is used internally by the indexer and will not be sent
        /// out by protocol.
        Opened = 2,
    }
    impl OrderPlacementStatus {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub fn as_str_name(&self) -> &'static str {
            match self {
                Self::Unspecified => "ORDER_PLACEMENT_STATUS_UNSPECIFIED",
                Self::BestEffortOpened => "ORDER_PLACEMENT_STATUS_BEST_EFFORT_OPENED",
                Self::Opened => "ORDER_PLACEMENT_STATUS_OPENED",
            }
        }
        /// Creates an enum from field names used in the ProtoBuf definition.
        pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
            match value {
                "ORDER_PLACEMENT_STATUS_UNSPECIFIED" => Some(Self::Unspecified),
                "ORDER_PLACEMENT_STATUS_BEST_EFFORT_OPENED" => {
                    Some(Self::BestEffortOpened)
                }
                "ORDER_PLACEMENT_STATUS_OPENED" => Some(Self::Opened),
                _ => None,
            }
        }
    }
}
impl ::prost::Name for OrderPlaceV1 {
    const NAME: &'static str = "OrderPlaceV1";
    const PACKAGE: &'static str = "bitoroprotocol.indexer.off_chain_updates";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.indexer.off_chain_updates.OrderPlaceV1".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.indexer.off_chain_updates.OrderPlaceV1".into()
    }
}
/// OrderRemove messages contain the id of the order removed, the reason for the
/// removal and the resulting status from the removal.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OrderRemoveV1 {
    #[prost(message, optional, tag = "1")]
    pub removed_order_id: ::core::option::Option<super::protocol::v1::IndexerOrderId>,
    #[prost(enumeration = "super::shared::OrderRemovalReason", tag = "2")]
    pub reason: i32,
    #[prost(enumeration = "order_remove_v1::OrderRemovalStatus", tag = "3")]
    pub removal_status: i32,
    /// The timestamp of the order removal.
    #[prost(message, optional, tag = "4")]
    pub time_stamp: ::core::option::Option<::prost_types::Timestamp>,
}
/// Nested message and enum types in `OrderRemoveV1`.
pub mod order_remove_v1 {
    /// OrderRemovalStatus is an enum for the resulting status after an order is
    /// removed.
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
    pub enum OrderRemovalStatus {
        /// Default value, this is invalid and unused.
        Unspecified = 0,
        /// A best effort canceled order is one that has only been confirmed to be
        /// removed on the Bitoro node sending the off-chain update message.
        /// The cases where this happens includes:
        /// - the order was removed due to the Bitoro node receiving a CancelOrder
        ///    transaction for the order.
        /// - the order was removed due to being undercollateralized during
        ///    optimistic matching.
        /// A best effort canceled order may not have been removed on other Bitoro
        /// nodes including other Bitoro validator nodes and may still be included in
        /// future order matches.
        BestEffortCanceled = 1,
        /// A canceled order is one that is confirmed to be removed on all Bitoro nodes
        /// (discounting dishonest Bitoro nodes) and will not be included in any future
        /// order matches.
        /// The cases where this happens includes:
        /// - the order is expired.
        Canceled = 2,
        /// An order was fully-filled. Only sent by the Indexer for stateful orders.
        Filled = 3,
    }
    impl OrderRemovalStatus {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub fn as_str_name(&self) -> &'static str {
            match self {
                Self::Unspecified => "ORDER_REMOVAL_STATUS_UNSPECIFIED",
                Self::BestEffortCanceled => "ORDER_REMOVAL_STATUS_BEST_EFFORT_CANCELED",
                Self::Canceled => "ORDER_REMOVAL_STATUS_CANCELED",
                Self::Filled => "ORDER_REMOVAL_STATUS_FILLED",
            }
        }
        /// Creates an enum from field names used in the ProtoBuf definition.
        pub fn from_str_name(value: &str) -> ::core::option::Option<Self> {
            match value {
                "ORDER_REMOVAL_STATUS_UNSPECIFIED" => Some(Self::Unspecified),
                "ORDER_REMOVAL_STATUS_BEST_EFFORT_CANCELED" => {
                    Some(Self::BestEffortCanceled)
                }
                "ORDER_REMOVAL_STATUS_CANCELED" => Some(Self::Canceled),
                "ORDER_REMOVAL_STATUS_FILLED" => Some(Self::Filled),
                _ => None,
            }
        }
    }
}
impl ::prost::Name for OrderRemoveV1 {
    const NAME: &'static str = "OrderRemoveV1";
    const PACKAGE: &'static str = "bitoroprotocol.indexer.off_chain_updates";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.indexer.off_chain_updates.OrderRemoveV1".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.indexer.off_chain_updates.OrderRemoveV1".into()
    }
}
/// OrderUpdate messages contain the id of the order being updated, and the
/// updated total filled quantums of the order.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OrderUpdateV1 {
    #[prost(message, optional, tag = "1")]
    pub order_id: ::core::option::Option<super::protocol::v1::IndexerOrderId>,
    #[prost(uint64, tag = "2")]
    pub total_filled_quantums: u64,
}
impl ::prost::Name for OrderUpdateV1 {
    const NAME: &'static str = "OrderUpdateV1";
    const PACKAGE: &'static str = "bitoroprotocol.indexer.off_chain_updates";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.indexer.off_chain_updates.OrderUpdateV1".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.indexer.off_chain_updates.OrderUpdateV1".into()
    }
}
/// OrderReplace messages contain the old order ID and the replacement order.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OrderReplaceV1 {
    /// vault replaces orders with a different order ID
    #[prost(message, optional, tag = "1")]
    pub old_order_id: ::core::option::Option<super::protocol::v1::IndexerOrderId>,
    #[prost(message, optional, tag = "2")]
    pub order: ::core::option::Option<super::protocol::v1::IndexerOrder>,
    #[prost(enumeration = "order_place_v1::OrderPlacementStatus", tag = "3")]
    pub placement_status: i32,
    #[prost(message, optional, tag = "4")]
    pub time_stamp: ::core::option::Option<::prost_types::Timestamp>,
}
impl ::prost::Name for OrderReplaceV1 {
    const NAME: &'static str = "OrderReplaceV1";
    const PACKAGE: &'static str = "bitoroprotocol.indexer.off_chain_updates";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.indexer.off_chain_updates.OrderReplaceV1".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.indexer.off_chain_updates.OrderReplaceV1".into()
    }
}
/// An OffChainUpdate message is the message type which will be sent on Kafka to
/// the Indexer.
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OffChainUpdateV1 {
    /// Contains one of an OrderPlaceV1, OrderRemoveV1, OrderUpdateV1, and
    /// OrderReplaceV1 message.
    #[prost(oneof = "off_chain_update_v1::UpdateMessage", tags = "1, 2, 3, 4")]
    pub update_message: ::core::option::Option<off_chain_update_v1::UpdateMessage>,
}
/// Nested message and enum types in `OffChainUpdateV1`.
pub mod off_chain_update_v1 {
    /// Contains one of an OrderPlaceV1, OrderRemoveV1, OrderUpdateV1, and
    /// OrderReplaceV1 message.
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum UpdateMessage {
        #[prost(message, tag = "1")]
        OrderPlace(super::OrderPlaceV1),
        #[prost(message, tag = "2")]
        OrderRemove(super::OrderRemoveV1),
        #[prost(message, tag = "3")]
        OrderUpdate(super::OrderUpdateV1),
        #[prost(message, tag = "4")]
        OrderReplace(super::OrderReplaceV1),
    }
}
impl ::prost::Name for OffChainUpdateV1 {
    const NAME: &'static str = "OffChainUpdateV1";
    const PACKAGE: &'static str = "bitoroprotocol.indexer.off_chain_updates";
    fn full_name() -> ::prost::alloc::string::String {
        "bitoroprotocol.indexer.off_chain_updates.OffChainUpdateV1".into()
    }
    fn type_url() -> ::prost::alloc::string::String {
        "/bitoroprotocol.indexer.off_chain_updates.OffChainUpdateV1".into()
    }
}
