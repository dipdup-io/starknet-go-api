package abi

import (
	"log"
	"testing"

	"github.com/goccy/go-json"
	"github.com/tfkhsr/jsonschema"
)

func TestAbi_JsonSchema(t *testing.T) {
	tests := []struct {
		name string
		abi  []byte
	}{
		{
			name: "test 1",
			abi: []byte(`[
				{
				  "name": "Uint256",
				  "size": 2,
				  "type": "struct",
				  "members": [
					{
					  "name": "low",
					  "type": "felt",
					  "offset": 0
					},
					{
					  "name": "high",
					  "type": "felt",
					  "offset": 1
					}
				  ]
				},
				{
				  "data": [
					{
					  "name": "from_",
					  "type": "felt"
					},
					{
					  "name": "to",
					  "type": "felt"
					},
					{
					  "name": "value",
					  "type": "Uint256"
					}
				  ],
				  "keys": [],
				  "name": "Transfer",
				  "type": "event"
				},
				{
				  "data": [
					{
					  "name": "owner",
					  "type": "felt"
					},
					{
					  "name": "spender",
					  "type": "felt"
					},
					{
					  "name": "value",
					  "type": "Uint256"
					}
				  ],
				  "keys": [],
				  "name": "Approval",
				  "type": "event"
				},
				{
				  "name": "constructor",
				  "type": "constructor",
				  "inputs": [
					{
					  "name": "recipient",
					  "type": "felt"
					}
				  ],
				  "outputs": []
				},
				{
				  "name": "name",
				  "type": "function",
				  "inputs": [],
				  "outputs": [
					{
					  "name": "name",
					  "type": "felt"
					}
				  ],
				  "stateMutability": "view"
				},
				{
				  "name": "symbol",
				  "type": "function",
				  "inputs": [],
				  "outputs": [
					{
					  "name": "symbol",
					  "type": "felt"
					}
				  ],
				  "stateMutability": "view"
				},
				{
				  "name": "totalSupply",
				  "type": "function",
				  "inputs": [],
				  "outputs": [
					{
					  "name": "totalSupply",
					  "type": "Uint256"
					}
				  ],
				  "stateMutability": "view"
				},
				{
				  "name": "decimals",
				  "type": "function",
				  "inputs": [],
				  "outputs": [
					{
					  "name": "decimals",
					  "type": "felt"
					}
				  ],
				  "stateMutability": "view"
				},
				{
				  "name": "balanceOf",
				  "type": "function",
				  "inputs": [
					{
					  "name": "account",
					  "type": "felt"
					}
				  ],
				  "outputs": [
					{
					  "name": "balance",
					  "type": "Uint256"
					}
				  ],
				  "stateMutability": "view"
				},
				{
				  "name": "allowance",
				  "type": "function",
				  "inputs": [
					{
					  "name": "owner",
					  "type": "felt"
					},
					{
					  "name": "spender",
					  "type": "felt"
					}
				  ],
				  "outputs": [
					{
					  "name": "remaining",
					  "type": "Uint256"
					}
				  ],
				  "stateMutability": "view"
				},
				{
				  "name": "transfer",
				  "type": "function",
				  "inputs": [
					{
					  "name": "recipient",
					  "type": "felt"
					},
					{
					  "name": "amount",
					  "type": "Uint256"
					}
				  ],
				  "outputs": [
					{
					  "name": "success",
					  "type": "felt"
					}
				  ]
				},
				{
				  "name": "transferFrom",
				  "type": "function",
				  "inputs": [
					{
					  "name": "sender",
					  "type": "felt"
					},
					{
					  "name": "recipient",
					  "type": "felt"
					},
					{
					  "name": "amount",
					  "type": "Uint256"
					}
				  ],
				  "outputs": [
					{
					  "name": "success",
					  "type": "felt"
					}
				  ]
				},
				{
				  "name": "approve",
				  "type": "function",
				  "inputs": [
					{
					  "name": "spender",
					  "type": "felt"
					},
					{
					  "name": "amount",
					  "type": "Uint256"
					}
				  ],
				  "outputs": [
					{
					  "name": "success",
					  "type": "felt"
					}
				  ]
				},
				{
				  "name": "increaseAllowance",
				  "type": "function",
				  "inputs": [
					{
					  "name": "spender",
					  "type": "felt"
					},
					{
					  "name": "added_value",
					  "type": "Uint256"
					}
				  ],
				  "outputs": [
					{
					  "name": "success",
					  "type": "felt"
					}
				  ]
				},
				{
				  "name": "decreaseAllowance",
				  "type": "function",
				  "inputs": [
					{
					  "name": "spender",
					  "type": "felt"
					},
					{
					  "name": "subtracted_value",
					  "type": "Uint256"
					}
				  ],
				  "outputs": [
					{
					  "name": "success",
					  "type": "felt"
					}
				  ]
				}
			  ]`),
		},
		{
			name: "test 2",
			abi:  []byte(`[{"type": "event", "name": "Hello", "keys": [{"name": "from", "type": "core::starknet::contract_address::ContractAddress"}, {"name": "value", "type": "core::integer::u128"}]}]`),
		},
		{
			name: "test 3",
			abi:  []byte(`[{"type": "function", "name": "constructor", "inputs": [{"name": "owner", "type": "core::felt252"}, {"name": "guardian", "type": "core::felt252"}], "outputs": [], "state_mutability": "external"}, {"type": "struct", "name": "lib::calls::Call", "members": [{"name": "to", "type": "core::starknet::contract_address::ContractAddress"}, {"name": "selector", "type": "core::felt252"}, {"name": "calldata", "type": "core::array::Array::<core::felt252>"}]}, {"type": "function", "name": "__validate__", "inputs": [{"name": "calls", "type": "core::array::Array::<lib::calls::Call>"}], "outputs": [{"type": "core::felt252"}], "state_mutability": "external"}, {"type": "function", "name": "__validate_declare__", "inputs": [{"name": "class_hash", "type": "core::felt252"}], "outputs": [{"type": "core::felt252"}], "state_mutability": "external"}, {"type": "function", "name": "__validate_deploy__", "inputs": [{"name": "class_hash", "type": "core::felt252"}, {"name": "contract_address_salt", "type": "core::felt252"}, {"name": "owner", "type": "core::felt252"}, {"name": "guardian", "type": "core::felt252"}], "outputs": [{"type": "core::felt252"}], "state_mutability": "external"}, {"type": "function", "name": "__execute__", "inputs": [{"name": "calls", "type": "core::array::Array::<lib::calls::Call>"}], "outputs": [{"type": "core::array::Span::<core::array::Span::<core::felt252>>"}], "state_mutability": "external"}, {"type": "struct", "name": "lib::outside_execution::OutsideExecution", "members": [{"name": "caller", "type": "core::starknet::contract_address::ContractAddress"}, {"name": "nonce", "type": "core::felt252"}, {"name": "execute_after", "type": "core::integer::u64"}, {"name": "execute_before", "type": "core::integer::u64"}, {"name": "calls", "type": "core::array::Array::<lib::calls::Call>"}]}, {"type": "function", "name": "execute_from_outside", "inputs": [{"name": "outside_execution", "type": "lib::outside_execution::OutsideExecution"}, {"name": "signature", "type": "core::array::Array::<core::felt252>"}], "outputs": [{"type": "core::array::Span::<core::array::Span::<core::felt252>>"}], "state_mutability": "external"}, {"type": "function", "name": "change_owner", "inputs": [{"name": "new_owner", "type": "core::felt252"}, {"name": "signature_r", "type": "core::felt252"}, {"name": "signature_s", "type": "core::felt252"}], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "change_guardian", "inputs": [{"name": "new_guardian", "type": "core::felt252"}], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "change_guardian_backup", "inputs": [{"name": "new_guardian_backup", "type": "core::felt252"}], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "trigger_escape_owner", "inputs": [{"name": "new_owner", "type": "core::felt252"}], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "trigger_escape_guardian", "inputs": [{"name": "new_guardian", "type": "core::felt252"}], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "escape_owner", "inputs": [], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "escape_guardian", "inputs": [], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "cancel_escape", "inputs": [], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "upgrade", "inputs": [{"name": "implementation", "type": "core::starknet::class_hash::ClassHash"}, {"name": "calldata", "type": "core::array::Array::<core::felt252>"}], "outputs": [{"type": "core::array::Array::<core::felt252>"}], "state_mutability": "external"}, {"type": "function", "name": "execute_after_upgrade", "inputs": [{"name": "data", "type": "core::array::Array::<core::felt252>"}], "outputs": [{"type": "core::array::Array::<core::felt252>"}], "state_mutability": "external"}, {"type": "function", "name": "get_owner", "inputs": [], "outputs": [{"type": "core::felt252"}], "state_mutability": "view"}, {"type": "function", "name": "get_guardian", "inputs": [], "outputs": [{"type": "core::felt252"}], "state_mutability": "view"}, {"type": "function", "name": "get_guardian_backup", "inputs": [], "outputs": [{"type": "core::felt252"}], "state_mutability": "view"}, {"type": "struct", "name": "account::escape::Escape", "members": [{"name": "ready_at", "type": "core::integer::u64"}, {"name": "escape_type", "type": "core::felt252"}, {"name": "new_signer", "type": "core::felt252"}]}, {"type": "function", "name": "get_escape", "inputs": [], "outputs": [{"type": "account::escape::Escape"}], "state_mutability": "view"}, {"type": "struct", "name": "lib::version::Version", "members": [{"name": "major", "type": "core::integer::u8"}, {"name": "minor", "type": "core::integer::u8"}, {"name": "patch", "type": "core::integer::u8"}]}, {"type": "function", "name": "get_version", "inputs": [], "outputs": [{"type": "lib::version::Version"}], "state_mutability": "view"}, {"type": "function", "name": "getVersion", "inputs": [], "outputs": [{"type": "core::felt252"}], "state_mutability": "view"}, {"type": "function", "name": "get_name", "inputs": [], "outputs": [{"type": "core::felt252"}], "state_mutability": "view"}, {"type": "function", "name": "getName", "inputs": [], "outputs": [{"type": "core::felt252"}], "state_mutability": "view"}, {"type": "function", "name": "get_guardian_escape_attempts", "inputs": [], "outputs": [{"type": "core::integer::u32"}], "state_mutability": "view"}, {"type": "function", "name": "get_owner_escape_attempts", "inputs": [], "outputs": [{"type": "core::integer::u32"}], "state_mutability": "view"}, {"type": "enum", "name": "account::escape::EscapeStatus", "variants": [{"name": "None", "type": "()"}, {"name": "NotReady", "type": "()"}, {"name": "Ready", "type": "()"}, {"name": "Expired", "type": "()"}]}, {"type": "function", "name": "get_escape_and_status", "inputs": [], "outputs": [{"type": "(account::escape::Escape, account::escape::EscapeStatus)"}], "state_mutability": "view"}, {"type": "function", "name": "supports_interface", "inputs": [{"name": "interface_id", "type": "core::felt252"}], "outputs": [{"type": "core::bool"}], "state_mutability": "view"}, {"type": "function", "name": "supportsInterface", "inputs": [{"name": "interface_id", "type": "core::felt252"}], "outputs": [{"type": "core::felt252"}], "state_mutability": "view"}, {"type": "function", "name": "is_valid_signature", "inputs": [{"name": "hash", "type": "core::felt252"}, {"name": "signatures", "type": "core::array::Array::<core::felt252>"}], "outputs": [{"type": "core::felt252"}], "state_mutability": "view"}, {"type": "function", "name": "isValidSignature", "inputs": [{"name": "hash", "type": "core::felt252"}, {"name": "signatures", "type": "core::array::Array::<core::felt252>"}], "outputs": [{"type": "core::felt252"}], "state_mutability": "view"}, {"type": "function", "name": "get_outside_execution_message_hash", "inputs": [{"name": "outside_execution", "type": "lib::outside_execution::OutsideExecution"}], "outputs": [{"type": "core::felt252"}], "state_mutability": "view"}, {"type": "event", "name": "AccountCreated", "inputs": [{"name": "account", "type": "core::starknet::contract_address::ContractAddress"}, {"name": "key", "type": "core::felt252"}, {"name": "guardian", "type": "core::felt252"}]}, {"type": "event", "name": "TransactionExecuted", "inputs": [{"name": "hash", "type": "core::felt252"}, {"name": "response", "type": "core::array::Span::<core::array::Span::<core::felt252>>"}]}, {"type": "event", "name": "EscapeOwnerTriggered", "inputs": [{"name": "ready_at", "type": "core::integer::u64"}, {"name": "new_owner", "type": "core::felt252"}]}, {"type": "event", "name": "EscapeGuardianTriggered", "inputs": [{"name": "ready_at", "type": "core::integer::u64"}, {"name": "new_guardian", "type": "core::felt252"}]}, {"type": "event", "name": "OwnerEscaped", "inputs": [{"name": "new_owner", "type": "core::felt252"}]}, {"type": "event", "name": "GuardianEscaped", "inputs": [{"name": "new_guardian", "type": "core::felt252"}]}, {"type": "event", "name": "EscapeCanceled", "inputs": []}, {"type": "event", "name": "OwnerChanged", "inputs": [{"name": "new_owner", "type": "core::felt252"}]}, {"type": "event", "name": "GuardianChanged", "inputs": [{"name": "new_guardian", "type": "core::felt252"}]}, {"type": "event", "name": "GuardianBackupChanged", "inputs": [{"name": "new_guardian", "type": "core::felt252"}]}, {"type": "event", "name": "AccountUpgraded", "inputs": [{"name": "new_implementation", "type": "core::starknet::class_hash::ClassHash"}]}]`),
		},
		{
			name: "test 4",
			abi:  []byte(`[{"name":"register_proof","type":"function","inputs":[{"name":"A","type":"core::ec::EcPoint"},{"name":"B","type":"core::ec::EcPoint"},{"name":"c","type":"core::felt252"},{"name":"proof_data","type":"core::array::Array::<core::felt252>"}],"outputs":[],"state_mutability":"external"}]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var abi Abi
			if err := json.Unmarshal(tt.abi, &abi); err != nil {
				t.Errorf("can't unmarshal abi: %s", string(tt.abi))
				return
			}
			got := abi.JsonSchema()
			b, err := json.MarshalIndent(got, "", " ")
			if err != nil {
				t.Errorf("can't marshal jsonschema: %s", string(tt.abi))
				return
			}

			log.Print(string(b))
			if _, err := jsonschema.Parse(b); err != nil {
				t.Errorf("can't parse resulting json schema: %s", string(tt.abi))
			}

		})
	}
}
