import * as _5 from "./accountplus/accountplus";
import * as _6 from "./accountplus/genesis";
import * as _7 from "./accountplus/models";
import * as _8 from "./accountplus/params";
import * as _9 from "./accountplus/query";
import * as _10 from "./accountplus/tx";
import * as _11 from "./affiliates/affiliates";
import * as _12 from "./affiliates/genesis";
import * as _13 from "./affiliates/query";
import * as _14 from "./affiliates/tx";
import * as _15 from "./assets/asset";
import * as _16 from "./assets/genesis";
import * as _17 from "./assets/query";
import * as _18 from "./assets/tx";
import * as _19 from "./blocktime/blocktime";
import * as _20 from "./blocktime/genesis";
import * as _21 from "./blocktime/params";
import * as _22 from "./blocktime/query";
import * as _23 from "./blocktime/tx";
import * as _24 from "./bridge/bridge_event_info";
import * as _25 from "./bridge/bridge_event";
import * as _26 from "./bridge/genesis";
import * as _27 from "./bridge/params";
import * as _28 from "./bridge/query";
import * as _29 from "./bridge/tx";
import * as _30 from "./clob/block_rate_limit_config";
import * as _31 from "./clob/clob_pair";
import * as _32 from "./clob/equity_tier_limit_config";
import * as _33 from "./clob/finalize_block";
import * as _34 from "./clob/genesis";
import * as _35 from "./clob/liquidations_config";
import * as _36 from "./clob/liquidations";
import * as _37 from "./clob/matches";
import * as _38 from "./clob/mev";
import * as _39 from "./clob/operation";
import * as _40 from "./clob/order_removals";
import * as _41 from "./clob/order";
import * as _42 from "./clob/process_proposer_matches_events";
import * as _43 from "./clob/query";
import * as _44 from "./clob/streaming";
import * as _45 from "./clob/tx";
import * as _46 from "./daemons/bridge/bridge";
import * as _47 from "./daemons/liquidation/liquidation";
import * as _48 from "./daemons/pricefeed/price_feed";
import * as _49 from "./delaymsg/block_message_ids";
import * as _50 from "./delaymsg/delayed_message";
import * as _51 from "./delaymsg/genesis";
import * as _52 from "./delaymsg/query";
import * as _53 from "./delaymsg/tx";
import * as _54 from "./epochs/epoch_info";
import * as _55 from "./epochs/genesis";
import * as _56 from "./epochs/query";
import * as _57 from "./feetiers/genesis";
import * as _58 from "./feetiers/params";
import * as _59 from "./feetiers/query";
import * as _60 from "./feetiers/tx";
import * as _61 from "./govplus/genesis";
import * as _62 from "./govplus/query";
import * as _63 from "./govplus/tx";
import * as _64 from "./indexer/events/events";
import * as _65 from "./indexer/indexer_manager/event";
import * as _66 from "./indexer/off_chain_updates/off_chain_updates";
import * as _67 from "./indexer/protocol/v1/clob";
import * as _68 from "./indexer/protocol/v1/perpetual";
import * as _69 from "./indexer/protocol/v1/subaccount";
import * as _70 from "./indexer/protocol/v1/vault";
import * as _71 from "./indexer/redis/redis_order";
import * as _72 from "./indexer/shared/removal_reason";
import * as _73 from "./indexer/socks/messages";
import * as _74 from "./listing/genesis";
import * as _75 from "./listing/params";
import * as _76 from "./listing/query";
import * as _77 from "./listing/tx";
import * as _78 from "./perpetuals/genesis";
import * as _79 from "./perpetuals/params";
import * as _80 from "./perpetuals/perpetual";
import * as _81 from "./perpetuals/query";
import * as _82 from "./perpetuals/tx";
import * as _83 from "./prices/genesis";
import * as _84 from "./prices/market_param";
import * as _85 from "./prices/market_price";
import * as _86 from "./prices/query";
import * as _87 from "./prices/streaming";
import * as _88 from "./prices/tx";
import * as _89 from "./ratelimit/capacity";
import * as _90 from "./ratelimit/genesis";
import * as _91 from "./ratelimit/limit_params";
import * as _92 from "./ratelimit/pending_send_packet";
import * as _93 from "./ratelimit/query";
import * as _94 from "./ratelimit/tx";
import * as _95 from "./revshare/genesis";
import * as _96 from "./revshare/params";
import * as _97 from "./revshare/query";
import * as _98 from "./revshare/revshare";
import * as _99 from "./revshare/tx";
import * as _100 from "./rewards/genesis";
import * as _101 from "./rewards/params";
import * as _102 from "./rewards/query";
import * as _103 from "./rewards/reward_share";
import * as _104 from "./rewards/tx";
import * as _105 from "./sending/genesis";
import * as _106 from "./sending/query";
import * as _107 from "./sending/transfer";
import * as _108 from "./sending/tx";
import * as _109 from "./stats/genesis";
import * as _110 from "./stats/params";
import * as _111 from "./stats/query";
import * as _112 from "./stats/stats";
import * as _113 from "./stats/tx";
import * as _114 from "./subaccounts/asset_position";
import * as _115 from "./subaccounts/genesis";
import * as _116 from "./subaccounts/perpetual_position";
import * as _117 from "./subaccounts/query";
import * as _118 from "./subaccounts/streaming";
import * as _119 from "./subaccounts/subaccount";
import * as _120 from "./vault/genesis";
import * as _121 from "./vault/params";
import * as _122 from "./vault/query";
import * as _123 from "./vault/share";
import * as _124 from "./vault/tx";
import * as _125 from "./vault/vault";
import * as _126 from "./vest/genesis";
import * as _127 from "./vest/query";
import * as _128 from "./vest/tx";
import * as _129 from "./vest/vest_entry";
import * as _137 from "./accountplus/query.lcd";
import * as _138 from "./affiliates/query.lcd";
import * as _139 from "./assets/query.lcd";
import * as _140 from "./blocktime/query.lcd";
import * as _141 from "./bridge/query.lcd";
import * as _142 from "./clob/query.lcd";
import * as _143 from "./delaymsg/query.lcd";
import * as _144 from "./epochs/query.lcd";
import * as _145 from "./feetiers/query.lcd";
import * as _146 from "./listing/query.lcd";
import * as _147 from "./perpetuals/query.lcd";
import * as _148 from "./prices/query.lcd";
import * as _149 from "./ratelimit/query.lcd";
import * as _150 from "./revshare/query.lcd";
import * as _151 from "./rewards/query.lcd";
import * as _152 from "./stats/query.lcd";
import * as _153 from "./subaccounts/query.lcd";
import * as _154 from "./vault/query.lcd";
import * as _155 from "./vest/query.lcd";
import * as _156 from "./accountplus/query.rpc.Query";
import * as _157 from "./affiliates/query.rpc.Query";
import * as _158 from "./assets/query.rpc.Query";
import * as _159 from "./blocktime/query.rpc.Query";
import * as _160 from "./bridge/query.rpc.Query";
import * as _161 from "./clob/query.rpc.Query";
import * as _162 from "./delaymsg/query.rpc.Query";
import * as _163 from "./epochs/query.rpc.Query";
import * as _164 from "./feetiers/query.rpc.Query";
import * as _165 from "./govplus/query.rpc.Query";
import * as _166 from "./listing/query.rpc.Query";
import * as _167 from "./perpetuals/query.rpc.Query";
import * as _168 from "./prices/query.rpc.Query";
import * as _169 from "./ratelimit/query.rpc.Query";
import * as _170 from "./revshare/query.rpc.Query";
import * as _171 from "./rewards/query.rpc.Query";
import * as _172 from "./sending/query.rpc.Query";
import * as _173 from "./stats/query.rpc.Query";
import * as _174 from "./subaccounts/query.rpc.Query";
import * as _175 from "./vault/query.rpc.Query";
import * as _176 from "./vest/query.rpc.Query";
import * as _177 from "./accountplus/tx.rpc.msg";
import * as _178 from "./affiliates/tx.rpc.msg";
import * as _179 from "./blocktime/tx.rpc.msg";
import * as _180 from "./bridge/tx.rpc.msg";
import * as _181 from "./clob/tx.rpc.msg";
import * as _182 from "./delaymsg/tx.rpc.msg";
import * as _183 from "./feetiers/tx.rpc.msg";
import * as _184 from "./govplus/tx.rpc.msg";
import * as _185 from "./listing/tx.rpc.msg";
import * as _186 from "./perpetuals/tx.rpc.msg";
import * as _187 from "./prices/tx.rpc.msg";
import * as _188 from "./ratelimit/tx.rpc.msg";
import * as _189 from "./revshare/tx.rpc.msg";
import * as _190 from "./rewards/tx.rpc.msg";
import * as _191 from "./sending/tx.rpc.msg";
import * as _192 from "./stats/tx.rpc.msg";
import * as _193 from "./vault/tx.rpc.msg";
import * as _194 from "./vest/tx.rpc.msg";
import * as _195 from "./lcd";
import * as _196 from "./rpc.query";
import * as _197 from "./rpc.tx";
export namespace bitoroprotocol {
  export const accountplus = { ..._5,
    ..._6,
    ..._7,
    ..._8,
    ..._9,
    ..._10,
    ..._137,
    ..._156,
    ..._177
  };
  export const affiliates = { ..._11,
    ..._12,
    ..._13,
    ..._14,
    ..._138,
    ..._157,
    ..._178
  };
  export const assets = { ..._15,
    ..._16,
    ..._17,
    ..._18,
    ..._139,
    ..._158
  };
  export const blocktime = { ..._19,
    ..._20,
    ..._21,
    ..._22,
    ..._23,
    ..._140,
    ..._159,
    ..._179
  };
  export const bridge = { ..._24,
    ..._25,
    ..._26,
    ..._27,
    ..._28,
    ..._29,
    ..._141,
    ..._160,
    ..._180
  };
  export const clob = { ..._30,
    ..._31,
    ..._32,
    ..._33,
    ..._34,
    ..._35,
    ..._36,
    ..._37,
    ..._38,
    ..._39,
    ..._40,
    ..._41,
    ..._42,
    ..._43,
    ..._44,
    ..._45,
    ..._142,
    ..._161,
    ..._181
  };
  export namespace daemons {
    export const bridge = { ..._46
    };
    export const liquidation = { ..._47
    };
    export const pricefeed = { ..._48
    };
  }
  export const delaymsg = { ..._49,
    ..._50,
    ..._51,
    ..._52,
    ..._53,
    ..._143,
    ..._162,
    ..._182
  };
  export const epochs = { ..._54,
    ..._55,
    ..._56,
    ..._144,
    ..._163
  };
  export const feetiers = { ..._57,
    ..._58,
    ..._59,
    ..._60,
    ..._145,
    ..._164,
    ..._183
  };
  export const govplus = { ..._61,
    ..._62,
    ..._63,
    ..._165,
    ..._184
  };
  export namespace indexer {
    export const events = { ..._64
    };
    export const indexer_manager = { ..._65
    };
    export const off_chain_updates = { ..._66
    };
    export namespace protocol {
      export const v1 = { ..._67,
        ..._68,
        ..._69,
        ..._70
      };
    }
    export const redis = { ..._71
    };
    export const shared = { ..._72
    };
    export const socks = { ..._73
    };
  }
  export const listing = { ..._74,
    ..._75,
    ..._76,
    ..._77,
    ..._146,
    ..._166,
    ..._185
  };
  export const perpetuals = { ..._78,
    ..._79,
    ..._80,
    ..._81,
    ..._82,
    ..._147,
    ..._167,
    ..._186
  };
  export const prices = { ..._83,
    ..._84,
    ..._85,
    ..._86,
    ..._87,
    ..._88,
    ..._148,
    ..._168,
    ..._187
  };
  export const ratelimit = { ..._89,
    ..._90,
    ..._91,
    ..._92,
    ..._93,
    ..._94,
    ..._149,
    ..._169,
    ..._188
  };
  export const revshare = { ..._95,
    ..._96,
    ..._97,
    ..._98,
    ..._99,
    ..._150,
    ..._170,
    ..._189
  };
  export const rewards = { ..._100,
    ..._101,
    ..._102,
    ..._103,
    ..._104,
    ..._151,
    ..._171,
    ..._190
  };
  export const sending = { ..._105,
    ..._106,
    ..._107,
    ..._108,
    ..._172,
    ..._191
  };
  export const stats = { ..._109,
    ..._110,
    ..._111,
    ..._112,
    ..._113,
    ..._152,
    ..._173,
    ..._192
  };
  export const subaccounts = { ..._114,
    ..._115,
    ..._116,
    ..._117,
    ..._118,
    ..._119,
    ..._153,
    ..._174
  };
  export const vault = { ..._120,
    ..._121,
    ..._122,
    ..._123,
    ..._124,
    ..._125,
    ..._154,
    ..._175,
    ..._193
  };
  export const vest = { ..._126,
    ..._127,
    ..._128,
    ..._129,
    ..._155,
    ..._176,
    ..._194
  };
  export const ClientFactory = { ..._195,
    ..._196,
    ..._197
  };
}