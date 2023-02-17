package types

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// IBCPortID is the default port id that profiles module binds to.
	IBCPortID = "feeabs"
)

type SpotPrice struct {
	SpotPrice string `json:"spot_price"`
}

type Result struct {
	Success bool   `json:"success"`
	Data    []byte `json:"data"`
}

type IcqRespones struct {
	Respones []Result `json:"responses"`
}

var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

// IBCPortKey defines the key to store the port ID in store.
var (
	IBCPortKey        = []byte{0x01}
	FeePoolAddressKey = []byte{0x02}
)

// NewOsmosisQueryRequestPacketData create new packet for ibc.
func NewOsmosisQueryRequestPacketData(poolId uint64, baseDenom string, quoteDenom string) OsmosisQuerySpotPriceRequestPacketData {
	return OsmosisQuerySpotPriceRequestPacketData{
		PoolId:          poolId,
		BaseAssetDenom:  baseDenom,
		QuoteAssetDenom: quoteDenom,
	}
}

// GetBytes is a helper for serializing.
func (p OsmosisQuerySpotPriceRequestPacketData) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&p))
}

// NewQueryArithmeticTwapToNowRequest create new packet for ibc.
func NewQueryArithmeticTwapToNowRequest(
	poolID uint64,
	baseDenom string,
	quoteDenom string,
	startTime time.Time,
) QueryArithmeticTwapToNowRequest {
	return QueryArithmeticTwapToNowRequest{
		PoolId:     poolID,
		BaseAsset:  baseDenom,
		QuoteAsset: quoteDenom,
		StartTime:  startTime,
	}
}

func (p QueryArithmeticTwapToNowRequest) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&p))
}

func NewInterchainQueryRequest(path string, data []byte) InterchainQueryRequest {
	return InterchainQueryRequest{
		Data: data,
		Path: path,
	}
}

// GetBytes is a helper for serializing.
func (p InterchainQueryRequest) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&p))
}

func NewInterchainQueryRequestPacket(req []InterchainQueryRequest) InterchainQueryRequestPacket {
	return InterchainQueryRequestPacket{
		Requests: req,
	}
}

// GetBytes is a helper for serializing.
func (p InterchainQueryRequestPacket) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&p))
}
