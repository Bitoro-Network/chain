package basic_manager

import (
	"cosmossdk.io/x/evidence"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	"cosmossdk.io/x/upgrade"
	delaymsgmodule "github.com/bitoro-network/chain/protocol/x/delaymsg"
	listingmodule "github.com/bitoro-network/chain/protocol/x/listing"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/ibc-go/modules/capability"

	custommodule "github.com/bitoro-network/chain/protocol/app/module"
	accountplusmodule "github.com/bitoro-network/chain/protocol/x/accountplus"
	affiliatesmodule "github.com/bitoro-network/chain/protocol/x/affiliates"
	assetsmodule "github.com/bitoro-network/chain/protocol/x/assets"
	blocktimemodule "github.com/bitoro-network/chain/protocol/x/blocktime"
	bridgemodule "github.com/bitoro-network/chain/protocol/x/bridge"
	clobmodule "github.com/bitoro-network/chain/protocol/x/clob"
	epochsmodule "github.com/bitoro-network/chain/protocol/x/epochs"
	feetiersmodule "github.com/bitoro-network/chain/protocol/x/feetiers"
	govplusmodule "github.com/bitoro-network/chain/protocol/x/govplus"
	perpetualsmodule "github.com/bitoro-network/chain/protocol/x/perpetuals"
	pricesmodule "github.com/bitoro-network/chain/protocol/x/prices"
	ratelimitmodule "github.com/bitoro-network/chain/protocol/x/ratelimit"
	revsharemodule "github.com/bitoro-network/chain/protocol/x/revshare"
	rewardsmodule "github.com/bitoro-network/chain/protocol/x/rewards"
	sendingmodule "github.com/bitoro-network/chain/protocol/x/sending"
	statsmodule "github.com/bitoro-network/chain/protocol/x/stats"
	subaccountsmodule "github.com/bitoro-network/chain/protocol/x/subaccounts"
	vaultmodule "github.com/bitoro-network/chain/protocol/x/vault"
	vestmodule "github.com/bitoro-network/chain/protocol/x/vest"
	marketmapmodule "github.com/skip-mev/slinky/x/marketmap"

	ica "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts"
	"github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	// Upgrades
)

var (
	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	// TODO(CORE-538): Remove ModuleBasics as it doesn't create the AppModuleBasic correctly since the fields
	// of the types aren't set causing panic during DefaultGenesis.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			[]govclient.ProposalHandler{
				paramsclient.ProposalHandler,
			},
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		custommodule.SlashingModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		ica.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		consensus.AppModuleBasic{},
		authzmodule.AppModuleBasic{},

		// Custom modules
		pricesmodule.AppModuleBasic{},
		assetsmodule.AppModuleBasic{},
		blocktimemodule.AppModuleBasic{},
		bridgemodule.AppModuleBasic{},
		feetiersmodule.AppModuleBasic{},
		perpetualsmodule.AppModuleBasic{},
		statsmodule.AppModuleBasic{},
		subaccountsmodule.AppModuleBasic{},
		clobmodule.AppModuleBasic{},
		vestmodule.AppModuleBasic{},
		rewardsmodule.AppModuleBasic{},
		delaymsgmodule.AppModuleBasic{},
		sendingmodule.AppModuleBasic{},
		epochsmodule.AppModuleBasic{},
		ratelimitmodule.AppModuleBasic{},
		govplusmodule.AppModuleBasic{},
		vaultmodule.AppModuleBasic{},
		revsharemodule.AppModuleBasic{},
		listingmodule.AppModuleBasic{},
		marketmapmodule.AppModuleBasic{},
		accountplusmodule.AppModuleBasic{},
		affiliatesmodule.AppModuleBasic{},
	)
)
