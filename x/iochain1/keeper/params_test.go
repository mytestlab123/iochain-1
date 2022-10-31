package keeper_test

import (
	"testing"

	testkeeper "github.com/mytestlab123/iochain-1/testutil/keeper"
	"github.com/mytestlab123/iochain-1/x/iochain1/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Iochain1Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
