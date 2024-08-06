package keeper

import (
	"bytes"
	"cosmossdk.io/store/prefix"

	"encoding/binary"
	"github.com/Skwizi_4/TrackerChain/x/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func Uint64ToBytes(n uint64) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, n)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (k Keeper) SetResult(ctx sdk.Context, lastBlockID uint64, result types.UserResults) (uint64, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BlockKeyPrefix))
	b := k.cdc.MustMarshal(&result)
	id, err := Uint64ToBytes(lastBlockID + 1)
	if err != nil {
		return 0, err
	}
	store.Set(id, b)
	return lastBlockID + 1, nil
}

func (k Keeper) GetLastIndex(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	var lastIndex uint64
	lastIndex = 0
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()
		if bytes.HasPrefix(key, types.KeyPrefix(types.BlockKeyPrefix)) {
			index := binary.BigEndian.Uint64(key[len(types.KeyPrefix(types.BlockKeyPrefix)):])
			if index > lastIndex {
				lastIndex = index
			}
		}
	}

	return lastIndex
}
