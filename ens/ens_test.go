package ens

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/umbracle/go-web3"
	"github.com/umbracle/go-web3/testutil"
)

func TestENS_Resolve(t *testing.T) {
	ens, err := NewENS(WithAddress(testutil.TestInfuraEndpoint(t)))
	assert.NoError(t, err)

	addr, err := ens.Resolve("arachnid.eth")
	assert.NoError(t, err)
	assert.Equal(t, web3.HexToAddress("0xb8c2C29ee19D8307cb7255e1Cd9CbDE883A267d5"), addr)
}
