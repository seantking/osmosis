package client

import (
	"github.com/osmosis-labs/osmosis/v19/x/superfluid/client/cli"
	"github.com/osmosis-labs/osmosis/v19/x/superfluid/client/rest"

	govclient "github.com/osmosis-labs/osmosis/v19/x/gov/client"
)

var (
	SetSuperfluidAssetsProposalHandler    = govclient.NewProposalHandler(cli.NewCmdSubmitSetSuperfluidAssetsProposal, rest.ProposalSetSuperfluidAssetsRESTHandler)
	RemoveSuperfluidAssetsProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitRemoveSuperfluidAssetsProposal, rest.ProposalRemoveSuperfluidAssetsRESTHandler)
	UpdateUnpoolWhitelistProposalHandler  = govclient.NewProposalHandler(cli.NewCmdUpdateUnpoolWhitelistProposal, rest.ProposalUpdateUnpoolWhitelistProposal)
)
