package quarantine

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

// RegisterLegacyAminoCodec registers all the necessary types and interfaces for the
// governance module.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgOptIn{}, "cosmos-sdk/MsgQuarantineOptIn")
	legacy.RegisterAminoMsg(cdc, &MsgOptOut{}, "cosmos-sdk/MsgQuarantineOptOut")
	legacy.RegisterAminoMsg(cdc, &MsgAccept{}, "cosmos-sdk/MsgQuarantineAccept")
	legacy.RegisterAminoMsg(cdc, &MsgDecline{}, "cosmos-sdk/MsgQuarantineDecline")
	legacy.RegisterAminoMsg(cdc, &MsgUpdateAutoResponses{}, "cosmos-sdk/MsgUpdateQuarantineAutoResp")
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgOptIn{},
		&MsgOptOut{},
		&MsgAccept{},
		&MsgDecline{},
		&MsgUpdateAutoResponses{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/quarantine module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/quarantine and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
