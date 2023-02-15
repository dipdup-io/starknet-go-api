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

func TestGetBlockWithTxHashesCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get block with tx hashes by hash call", func(t *testing.T) {
		if _, err := api.GetBlockWithTxHashes(
			context.Background(),
			BlockFilter{
				Hash: "0x144282ad24c23c724b537f5afe759de928e30048862caf105a36e121aaedbe",
			},
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetBlockWithTxHashes(hash) error = %v", err)
		}

		if _, err := api.GetBlockWithTxHashes(
			context.Background(),
			BlockFilter{
				Number: 1,
			},
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetBlockWithTxHashes(number) error = %v", err)
		}

		if _, err := api.GetBlockWithTxHashes(
			context.Background(),
			BlockFilter{
				String: Latest,
			},
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetBlockWithTxHashes(latest) error = %v", err)
		}
	})
}

func TestGetBlockWithTxsCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get block with txs by hash call", func(t *testing.T) {
		if _, err := api.GetBlockWithTxs(
			context.Background(),
			BlockFilter{
				Hash: "0x144282ad24c23c724b537f5afe759de928e30048862caf105a36e121aaedbe",
			},
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetBlockWithTxs(hash) error = %v", err)
			return
		}

		if _, err := api.GetBlockWithTxs(
			context.Background(),
			BlockFilter{
				Number: 1,
			},
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetBlockWithTxs(number) error = %v", err)
			return
		}

		if _, err := api.GetBlockWithTxs(
			context.Background(),
			BlockFilter{
				String: Latest,
			},
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetBlockWithTxs(latest) error = %v", err)
			return
		}
	})
}

func TestGetClassAtCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get get class at call", func(t *testing.T) {
		if _, err := api.GetClassAt(
			context.Background(),
			BlockFilter{
				Hash: "0x75e00250d4343326f322e370df4c9c73c7be105ad9f532eeb97891a34d9e4a5",
			},
			"0x2fb7ff5b1b474e8e691f5bebad9aa7aa3009f6ef22ccc2816f96cdfe217604d",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassAt(hash) error = %v", err)
		}

		if _, err := api.GetClassAt(
			context.Background(),
			BlockFilter{
				Number: 20,
			},
			"0x2fb7ff5b1b474e8e691f5bebad9aa7aa3009f6ef22ccc2816f96cdfe217604d",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassAt(number) error = %v", err)
		}

		if _, err := api.GetClassAt(
			context.Background(),
			BlockFilter{
				String: Latest,
			},
			"0x2fb7ff5b1b474e8e691f5bebad9aa7aa3009f6ef22ccc2816f96cdfe217604d",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassAt(latest) error = %v", err)
		}
	})
}

func TestGetClassHashAtCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get get class hash at call", func(t *testing.T) {
		if _, err := api.GetClassHashAt(
			context.Background(),
			BlockFilter{
				Hash: "0x75e00250d4343326f322e370df4c9c73c7be105ad9f532eeb97891a34d9e4a5",
			},
			"0x2fb7ff5b1b474e8e691f5bebad9aa7aa3009f6ef22ccc2816f96cdfe217604d",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassHashAt(hash) error = %v", err)
		}

		if _, err := api.GetClassHashAt(
			context.Background(),
			BlockFilter{
				Number: 20,
			},
			"0x2fb7ff5b1b474e8e691f5bebad9aa7aa3009f6ef22ccc2816f96cdfe217604d",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassHashAt(number) error = %v", err)
		}

		if _, err := api.GetClassHashAt(
			context.Background(),
			BlockFilter{
				String: Latest,
			},
			"0x2fb7ff5b1b474e8e691f5bebad9aa7aa3009f6ef22ccc2816f96cdfe217604d",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClassHashAt(latest) error = %v", err)
		}
	})
}

func TestGetClassByBlockNumberCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get class call", func(t *testing.T) {
		if _, err := api.GetClass(
			context.Background(),
			BlockFilter{
				Hash: "0x11172ea58125f54df2c07df73accd9236558944ec0ee650d80968f863267764",
			},
			"0x10455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClass(hash) error = %v", err)
		}

		if _, err := api.GetClass(
			context.Background(),
			BlockFilter{
				Number: 2,
			},
			"0x10455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClass(number) error = %v", err)
		}

		if _, err := api.GetClass(
			context.Background(),
			BlockFilter{
				String: Latest,
			},
			"0x10455c752b86932ce552f2b0fe81a880746649b9aee7e0d842bf3f52378f9f8",
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetClass(number) error = %v", err)
		}
	})
}

func TestGetBlockTransactionCountCall(t *testing.T) {
	api := NewAPI(testUrl)

	t.Run("test get transaction count call", func(t *testing.T) {
		if _, err := api.GetBlockTransactionCount(
			context.Background(),
			BlockFilter{
				Number: 10,
			},
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetBlockTransactionCount(number) error = %v", err)
		}

		if _, err := api.GetBlockTransactionCount(
			context.Background(),
			BlockFilter{
				Hash: "0x144282ad24c23c724b537f5afe759de928e30048862caf105a36e121aaedbe",
			},
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetBlockTransactionCount(hash) error = %v", err)
		}

		if _, err := api.GetBlockTransactionCount(
			context.Background(),
			BlockFilter{
				String: Latest,
			},
			WithTimeout(2),
		); err != nil {
			t.Errorf("GetBlockTransactionCount(latest) error = %v", err)
		}
	})
}
