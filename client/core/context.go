package core

import (
	rpcclient "github.com/tendermint/tendermint/rpc/client"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CoreContext struct {
	ChainID         string
	Height          int64
	TrustNode       bool
	NodeURI         string
	FromAddressName string
	Sequence        int64
	Client          rpcclient.Client
	Decoder         sdk.AccountDecoder
	AccountStore    string
}

// WithChainID - return a copy of the context with an updated chainID
func (c CoreContext) WithChainID(chainID string) CoreContext {
	c.ChainID = chainID
	return c
}

// WithHeight - return a copy of the context with an updated height
func (c CoreContext) WithHeight(height int64) CoreContext {
	c.Height = height
	return c
}

// WithTrustNode - return a copy of the context with an updated TrustNode flag
func (c CoreContext) WithTrustNode(trustNode bool) CoreContext {
	c.TrustNode = trustNode
	return c
}

// WithNodeURI - return a copy of the context with an updated node URI
func (c CoreContext) WithNodeURI(nodeURI string) CoreContext {
	c.NodeURI = nodeURI
	c.Client = rpcclient.NewHTTP(nodeURI, "/websocket")
	return c
}

// WithFromAddressName - return a copy of the context with an updated from address
func (c CoreContext) WithFromAddressName(fromAddressName string) CoreContext {
	c.FromAddressName = fromAddressName
	return c
}

// WithSequence - return a copy of the context with an updated sequence number
func (c CoreContext) WithSequence(sequence int64) CoreContext {
	c.Sequence = sequence
	return c
}

// WithClient - return a copy of the context with an updated RPC client instance
func (c CoreContext) WithClient(client rpcclient.Client) CoreContext {
	c.Client = client
	return c
}

// WithDecoder - return a copy of the context with an updated Decoder
func (c CoreContext) WithDecoder(decoder sdk.AccountDecoder) CoreContext {
	c.Decoder = decoder
	return c
}

// WithAccountStore - return a copy of the context with an updated AccountStore
func (c CoreContext) WithAccountStore(accountStore string) CoreContext {
	c.AccountStore = accountStore
	return c
}
