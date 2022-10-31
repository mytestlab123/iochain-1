package iochain1_test

import (
	"testing"

	keepertest "github.com/mytestlab123/iochain-1/testutil/keeper"
	"github.com/mytestlab123/iochain-1/testutil/nullify"
	"github.com/mytestlab123/iochain-1/x/iochain1"
	"github.com/mytestlab123/iochain-1/x/iochain1/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Iochain1Keeper(t)
	iochain1.InitGenesis(ctx, *k, genesisState)
	got := iochain1.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
