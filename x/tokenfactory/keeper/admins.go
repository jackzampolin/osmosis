package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/osmosis-labs/osmosis/x/tokenfactory/types"
)

// GetAuthorityMetadata returns the authority metadata for a specific denom
func (k Keeper) GetAuthorityMetadata(ctx sdk.Context, denom string) (types.DenomAuthorityMetadata, error) {
	bz := k.GetDenomPrefixStore(ctx, denom).Get([]byte(types.DenomAuthorityMetadataKey))

	metadata := types.DenomAuthorityMetadata{}
	err := proto.Unmarshal(bz, &metadata)
	if err != nil {
		return types.DenomAuthorityMetadata{}, err
	}
	return metadata, nil
}

func (k Keeper) SetAuthorityMetadata(ctx sdk.Context, denom string, metadata types.DenomAuthorityMetadata) error {
	if metadata.Admin != "" {
		_, err := sdk.AccAddressFromBech32(metadata.Admin)
		if err != nil {
			return err
		}
	}

	store := k.GetDenomPrefixStore(ctx, denom)

	bz, err := proto.Marshal(&metadata)
	if err != nil {
		return err
	}

	store.Set([]byte(types.DenomAuthorityMetadataKey), bz)
	return nil
}

func (k Keeper) setAdmin(ctx sdk.Context, denom string, admin string) error {
	metadata, err := k.GetAuthorityMetadata(ctx, denom)
	if err != nil {
		return err
	}

	metadata.Admin = admin

	return k.SetAuthorityMetadata(ctx, denom, metadata)
}
