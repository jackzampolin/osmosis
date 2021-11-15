package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/osmosis-labs/osmosis/x/tokenfactory/types"
)

// ConvertToBaseToken converts a fee amount in a whitelisted fee token to the base fee token amount
func (k Keeper) CreateDenom(ctx sdk.Context, creatorAddr string, denomnonce string) error {
	denom, err := types.GetTokenDenom(creatorAddr, denomnonce)
	if err != nil {
		return err
	}

	_, found := k.bankKeeper.GetDenomMetaData(ctx, denom)
	if found {
		return types.ErrDenomExists
	}

	baseDenomUnit := banktypes.DenomUnit{
		Denom:    denom,
		Exponent: 0,
	}

	denomMetaData := banktypes.Metadata{
		DenomUnits: []*banktypes.DenomUnit{&baseDenomUnit},
		Base:       denom,
	}

	k.bankKeeper.SetDenomMetaData(ctx, denomMetaData)

	authorityMetadata := types.DenomAuthorityMetadata{
		Admin: creatorAddr,
	}
	err = k.SetAuthorityMetadata(ctx, denom, authorityMetadata)
	if err != nil {
		return err
	}

	k.addDenomFromCreator(ctx, creatorAddr, denom)
	return nil
}
