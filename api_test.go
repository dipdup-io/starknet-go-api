package starknetgoapi

import (
	"context"
	"testing"
)

const testUrl = "https://starknet-testnet.public.blastapi.io"

func TestBlockAndHashCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test block and hash call", func(t *testing.T) {
		if _, err := api.BlockHashAndNumber(context.Background(), WithTimeout(2)); err != nil {
			t.Errorf("BlockHashAndNumber() error = %v", err)
		}
	})
}

func TestBlockNumberCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test block number call", func(t *testing.T) {
		if _, err := api.BlockNumber(context.Background(), WithTimeout(2)); err != nil {
			t.Errorf("BlockNumber() error = %v", err)
		}
	})
}

func TestChainIDCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test chain id call", func(t *testing.T) {
		if _, err := api.ChainID(context.Background(), WithTimeout(2)); err != nil {
			t.Errorf("ChainID() error = %v", err)
		}
	})
}

func TestGetBlockWithTxHashesByHashCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get block with tx hashes by hash call", func(t *testing.T) {
		if _, err := api.GetBlockWithTxHashesByHash(
			context.Background(),
			"0x144282ad24c23c724b537f5afe759de928e30048862caf105a36e121aaedbe",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetBlockWithTxHashesByHash() error = %v", err)
		}
	})
}

func TestGetBlockWithTxHashesByNumberCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get block with tx hashes by number call", func(t *testing.T) {
		if _, err := api.GetBlockWithTxHashesByNumber(
			context.Background(),
			1,
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetBlockWithTxHashesByNumber() error = %v", err)
		}
	})
}

func TestGetBlockWithTxsByHashCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get block with txs by hash call", func(t *testing.T) {
		if _, err := api.GetBlockWithTxsByHash(
			context.Background(),
			"0x144282ad24c23c724b537f5afe759de928e30048862caf105a36e121aaedbe",
			WithTimeout(2),
		); err != nil {
			t.Errorf("TestGetBlockWithTxsByHashCall() error = %v", err)
		}
	})
}

func TestGetBlockWithTxsByNumberCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get block with txs by number call", func(t *testing.T) {
		if _, err := api.GetBlockWithTxsByNumber(
			context.Background(),
			1,
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetBlockWithTxsByNumber() error = %v", err)
		}
	})
}

func TestGetClassAtBlockHashCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get get class at block hash call", func(t *testing.T) {
		if _, err := api.GetClassAtBlockHash(
			context.Background(),
			"0x75e00250d4343326f322e370df4c9c73c7be105ad9f532eeb97891a34d9e4a5",
			"0x2fb7ff5b1b474e8e691f5bebad9aa7aa3009f6ef22ccc2816f96cdfe217604d",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassAtBlockHash() error = %v", err)
		}
	})
}

func TestGetClassAtBlockNumberCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get class at block number call", func(t *testing.T) {
		if _, err := api.GetClassAtBlockNumber(
			context.Background(),
			1,
			"0x2fb7ff5b1b474e8e691f5bebad9aa7aa3009f6ef22ccc2816f96cdfe217604d",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassAtBlockNumber() error = %v", err)
		}
	})
}

func TestGetClassAtLatestBlockCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get class at latest call", func(t *testing.T) {
		if _, err := api.GetClassAtLatestBlock(
			context.Background(),
			"0x2fb7ff5b1b474e8e691f5bebad9aa7aa3009f6ef22ccc2816f96cdfe217604d",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassAtLatestBlock() error = %v", err)
		}
	})
}

func TestGetClassAtPendingBlockCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get class at pending call", func(t *testing.T) {
		if _, err := api.GetClassAtPendingBlock(
			context.Background(),
			"0x2fb7ff5b1b474e8e691f5bebad9aa7aa3009f6ef22ccc2816f96cdfe217604d",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassAtPendingBlock() error = %v", err)
		}
	})
}

func TestGetClassByBlockHashCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get class by block hash call", func(t *testing.T) {
		if _, err := api.GetClassByBlockHash(
			context.Background(),
			"0x11172ea58125f54df2c07df73accd9236558944ec0ee650d80968f863267764",
			"0x10455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassByBlockHash() error = %v", err)
		}
	})
}

func TestGetClassByBlockNumberCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get class by block number call", func(t *testing.T) {
		if _, err := api.GetClassByBlockNumber(
			context.Background(),
			2,
			"0x10455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassByBlockNumber() error = %v", err)
		}
	})
}

func TestGetClassByLatestCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get class by latest block call", func(t *testing.T) {
		if _, err := api.GetClassByLatestBlock(
			context.Background(),
			"0x10455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassByLatestBlock() error = %v", err)
		}
	})
}

func TestGetClassByPendingCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get class by pending block call", func(t *testing.T) {
		if _, err := api.GetClassByPendingBlock(
			context.Background(),
			"0x10455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassByPendingBlock() error = %v", err)
		}
	})
}

func TestGetTransactionCountByBlockNumberCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get transaction count by block number call", func(t *testing.T) {
		if _, err := api.GetTransactionCountByBlockNumber(
			context.Background(),
			10,
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetTransactionCountByBlockNumber() error = %v", err)
		}
	})
}
