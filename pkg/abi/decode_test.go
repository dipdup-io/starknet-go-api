package abi

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/require"
)

func TestDecodeFunctionCallData(t *testing.T) {
	type args struct {
		calldata []string
		abi      string
		endpoint string
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "test 1",
			args: args{
				calldata: []string{
					"0x327d34747122d7a40f4670265b098757270a449ec80c4871450fffdab7c2fa8",
					"0x0",
				},
				endpoint: "test_re_entrance",
				abi: `[{
																"type": "function",
																"name": "test_re_entrance",
																"inputs": [
																  {
																	"name": "other_contract_address",
																	"type": "felt"
																  },
																  {
																	"name": "depth",
																	"type": "felt"
																  }
																],
																"outputs": []
															}]`,
			},
			want: map[string]any{
				"other_contract_address": "0x327d34747122d7a40f4670265b098757270a449ec80c4871450fffdab7c2fa8",
				"depth":                  "0x0",
			},
		}, {
			name: "test 2",
			args: args{
				calldata: []string{
					"0x327d34747122d7a40f4670265b098757270a449ec80c4871450fffdab7c2fa8",
					"0x317eb442b72a9fae758d4fb26830ed0d9f31c8e7da4dbff4e8c59ea6a158e7f",
					"0x4",
					"0x5bd24b507fcc2fd77dc7847babb8df01363d58e9b0bbcd2d06d982e1f3e0c86",
					"0x2",
					"0x26b5943d4a0c420607cee8030a8cdd859bf2814a06633d165820960a42c6aed",
					"0x1518eec76afd5397cefd14eda48d01ad59981f9ce9e70c233ca67acd8754008",
				},
				endpoint: "test_call_contract",
				abi: `[{
															"type": "function",
															"name": "test_call_contract",
															"inputs": [
															  {
																"name": "contract_address",
																"type": "felt"
															  },
															  {
																"name": "function_selector",
																"type": "felt"
															  },
															  {
																"name": "calldata_len",
																"type": "felt"
															  },
															  {
																"name": "calldata",
																"type": "felt*"
															  }
															],
															"outputs": []
														}]`,
			},
			want: map[string]any{
				"contract_address":  "0x327d34747122d7a40f4670265b098757270a449ec80c4871450fffdab7c2fa8",
				"function_selector": "0x317eb442b72a9fae758d4fb26830ed0d9f31c8e7da4dbff4e8c59ea6a158e7f",
				"calldata_len":      "0x4",
				"calldata": []any{
					"0x5bd24b507fcc2fd77dc7847babb8df01363d58e9b0bbcd2d06d982e1f3e0c86",
					"0x2",
					"0x26b5943d4a0c420607cee8030a8cdd859bf2814a06633d165820960a42c6aed",
					"0x1518eec76afd5397cefd14eda48d01ad59981f9ce9e70c233ca67acd8754008",
				},
			},
		}, {
			name: "test 3",
			args: args{
				calldata: []string{
					"0xf899bfebffcf95f54cf61ac07a85587988678ed4",
					"0xde0b6b3a7640000",
					"0x0",
				},
				endpoint: "initiate_withdraw",
				abi: `[{
															"type": "function",
															"name": "initiate_withdraw",
															"inputs": [
															  {
																"name": "l1_recipient",
																"type": "felt"
															  },
															  {
																"name": "amount",
																"type": "Uint256"
															  }
															],
															"outputs": []
														},{
															"name": "Uint256",
															"size": 2,
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
															],
															"type": "struct"
														}]`,
			},
			want: map[string]any{
				"l1_recipient": "0xf899bfebffcf95f54cf61ac07a85587988678ed4",
				"amount": map[string]any{
					"low":  "0xde0b6b3a7640000",
					"high": "0x0",
				},
			},
		}, {
			name: "test 4",
			args: args{
				calldata: []string{
					"0x1",
					"0x41a78e741e5af2fec34b695679bc6891742439f7afb8484ecd7766661ad02bf",
					"0x1987cbd17808b9a23693d4de7e246a443cfe37e6e7fbaeabd7d7e6532b07c3d",
					"0xa",
					"0x48e09e58a43b0794821ac868ed3c2d940a2ff13a2a7edc57dea999cf881e439",
					"0x52554c45532056322e30",
					"0x0",
					"0x6",
					"0x2",
					"0x68747470733a2f2f6170692e72756c65732e6172742f6d657461646174612f",
					"0x7b69647d2e6a736f6e",
					"0x7ca95733f27cbd125214ada1c131a100daca25354cb9bfcac1ef5ea479646b5",
					"0x4b5cab01f7e6e59df39c14a3f3436c721554f719e42618fd4e69e4a605d39ef",
					"0x63a4b3b0122cdaa6ba244739add94aed1d31e3330458cda833a8d119f28cbe8",
				},
				endpoint: "__execute__",
				abi: `[{
														"name": "rules_account::account::interface::Call",
														"type": "struct",
														"members": [
														  {
															"name": "to",
															"type": "core::starknet::contract_address::ContractAddress"
														  },
														  {
															"name": "selector",
															"type": "core::felt252"
														  },
														  {
															"name": "calldata",
															"type": "core::array::Array::<core::felt252>"
														  }
														]
													},{
														"name": "__execute__",
														"type": "function",
														"inputs": [
														  {
															"name": "calls",
															"type": "core::array::Array::<rules_account::account::interface::Call>"
														  }
														],
														"outputs": [
														  {
															"type": "core::array::Array::<core::array::Span::<core::felt252>>"
														  }
														],
														"state_mutability": "external"
													}]`,
			},
			want: map[string]any{
				"calls": []any{
					map[string]any{
						"selector": "0x1987cbd17808b9a23693d4de7e246a443cfe37e6e7fbaeabd7d7e6532b07c3d",
						"to":       "0x41a78e741e5af2fec34b695679bc6891742439f7afb8484ecd7766661ad02bf",
						"calldata": []any{
							"0x48e09e58a43b0794821ac868ed3c2d940a2ff13a2a7edc57dea999cf881e439",
							"0x52554c45532056322e30",
							"0x0",
							"0x6",
							"0x2",
							"0x68747470733a2f2f6170692e72756c65732e6172742f6d657461646174612f",
							"0x7b69647d2e6a736f6e",
							"0x7ca95733f27cbd125214ada1c131a100daca25354cb9bfcac1ef5ea479646b5",
							"0x4b5cab01f7e6e59df39c14a3f3436c721554f719e42618fd4e69e4a605d39ef",
							"0x63a4b3b0122cdaa6ba244739add94aed1d31e3330458cda833a8d119f28cbe8",
						},
					},
				},
			},
		}, {
			name: "test 5",
			args: args{
				calldata: []string{
					"0x342be14f26ee1d8388f6bb876227311aec71cf9a38d962e78f57f18764c980a",
					"0x7",
					"0x0",
					"0x1",
				},
				endpoint: "test_deploy",
				abi: `[{
														"name": "test_deploy",
														"type": "function",
														"inputs": [
														  {
															"name": "class_hash",
															"type": "core::starknet::class_hash::ClassHash"
														  },
														  {
															"name": "contract_address_salt",
															"type": "core::felt252"
														  },
														  {
															"name": "calldata",
															"type": "core::array::Array::<core::felt252>"
														  },
														  {
															"name": "deploy_from_zero",
															"type": "core::bool"
														  }
														],
														"outputs": [],
														"state_mutability": "external"
													  }]`,
			},
			want: map[string]any{
				"class_hash":            "0x342be14f26ee1d8388f6bb876227311aec71cf9a38d962e78f57f18764c980a",
				"contract_address_salt": "0x7",
				"deploy_from_zero":      true,
				"calldata":              []any{},
			},
		}, {
			name: "test 6",
			args: args{
				calldata: []string{
					"0x2c133a0c1b78b0b16cf787bfeeb6a8978ef26e9d97a993cc7ffce9d192a313",
					"0x744f02e8b25f48d70fc865bfe2d1282a796f68607ed19c20a99b151b86b5ebc",
					"0xa",
					"0xe1f9f64cc1ae4ed64862b752db544ec8",
					"0x11a4",
					"0xc64ba59b83c589b886270cc611ed8183",
					"0x10cc",
					"0x5b3014d566ebc40534220cab925d68a9",
					"0x1189",
					"0xbd25422deabd085a3e0f9de4f36a736",
					"0x11c3",
					"0x71539945d31799e04899aed68e7510db",
					"0xd97",
					"0x9e455df7442c79cfb26ef95d7b4f2f94",
					"0x1260",
					"0x112bbd2afaa1219009e5b94cb18aa814",
					"0x1291",
					"0xc64ba59b83c589b886270cc611ed8183",
					"0x126e",
					"0x4ff3a57a3ae31a6e4b6e30ac86d6854",
					"0x12f1",
					"0x503a7ecedcc23347c0ddf70ff8771625",
					"0x1243",
					"0xa",
					"0x1",
					"0x0",
					"0x1",
					"0x0",
					"0x1",
					"0x0",
					"0x1",
					"0x0",
					"0x1",
					"0x0",
					"0x1",
					"0x0",
					"0x1",
					"0x0",
					"0x1",
					"0x0",
					"0x1",
					"0x0",
					"0x1",
					"0x0",
					"0x0",
				},
				endpoint: "safe_batch_transfer_from",
				abi: `[ {
													"name": "safe_batch_transfer_from",
													"type": "function",
													"inputs": [
													  {
														"name": "from",
														"type": "core::starknet::contract_address::ContractAddress"
													  },
													  {
														"name": "to",
														"type": "core::starknet::contract_address::ContractAddress"
													  },
													  {
														"name": "ids",
														"type": "core::array::Span::<core::integer::u256>"
													  },
													  {
														"name": "amounts",
														"type": "core::array::Span::<core::integer::u256>"
													  },
													  {
														"name": "data",
														"type": "core::array::Span::<core::felt252>"
													  }
													],
													"outputs": [],
													"state_mutability": "external"
												  }]`,
			},
			want: map[string]any{
				"from": "0x2c133a0c1b78b0b16cf787bfeeb6a8978ef26e9d97a993cc7ffce9d192a313",
				"to":   "0x744f02e8b25f48d70fc865bfe2d1282a796f68607ed19c20a99b151b86b5ebc",
				"ids": []any{
					"1537015543191481526425763139164378848579272",
					"1463477757684383927484695633877883616592259",
					"1527648754508514042686069948630541910501545",
					"1547279635986286197371806523182383403280182",
					"1183992991350848610714626900544229899309275",
					"1600898632193766335772009556738516408283028",
					"1617384913956694341372619145443166416185364",
					"1605715787057336205212386219784362728980867",
					"1650035839330635939173146333879076828047444",
					"1590926707320240837995093984183187355932197",
				},
				"amounts": []any{
					"1", "1", "1", "1", "1", "1", "1", "1", "1", "1",
				},
				"data": []any{},
			},
		}, {
			name: "test 6",
			args: args{
				calldata: []string{
					"0x3c0c5cb60d57873e6d74d4240300c5d0601303595b92bea80ca2f9e59aa3692",
					"0x2a2b93d24a5278aaff4a3ec913a3d616f21d8c30f9a3b2ab123a4f24aad75de",
					"0xd545533296acb37934b889a95b392f8021ef61633a010da565a68975be8f55",
					"0x5498d2c9437143a390c42532ee70c750573a8a196d0070a8e478afd6f25e03c",
					"0x7a83193b9f136a80445649a764051bbb004bd7bce7089edef9e1ae5226ff36",
					"0xe",
					"0x4a4d170ac5d8749bf7b7230d44b64d449f83802c0a4d8a55df5b27da5730ad",
					"0xce0bf2d1c917a039b2663409bbfe749d2bb7a0ae4019ff32e7e69e88c8c9fb",
					"0x2ca95f97a154230ed28be357dae4c73b635517dac78c78a3765e855b5d93402",
					"0x3f910281f29d7078852cadc7c261bcb86e24d85e717adf9f60b5e96b4ecc652",
					"0x1f8fbad42408418f0871372ea5272d16109b77e3a2ff748cd502f86fc06cc31",
					"0x55495050b2b69ba162b13790cf42fcac39627e58f649fcd6fc655b10fe86b4c",
					"0x4c08652376ec0779410ccaa9f28705dedbc273bcb4fac0c2fdaf68d74742c08",
					"0x5f63751a7bfbbb4bc92d30a1d27abd41f41fca6961c052cbfc8b6c708f0ee1b",
					"0xf89a68bb8715d2ca78c37f66bed1b22f13bcefec9b4fa3fa2d1cd5d02b539b",
					"0x180a21515f534b464bb93700c9d558e84a2221b7e2d903edf2867c0dc38027e",
					"0x3f9d8504820a9419714e74c93d8b0f9013058433000953612ad7fae6f4ccd9c",
					"0x63b67c4ad9915a0b47292253ff6afc90c7343ca2e24dfdbcae18893441e49d0",
					"0x59f0fc402eee19c7ed9efcdb9074a9c631b82c8c7d3089a7b9ffb9d1a690eed",
					"0xda3f81b91d292a5c2efd960c57d4a23fa4c003facba8539a76cdddea9cb2b1",
				},
				endpoint: "register_proof",
				abi: `[
													{
													  "name": "register_proof",
													  "type": "function",
													  "inputs": [
														{
														  "name": "A",
														  "type": "core::ec::EcPoint"
														},
														{
														  "name": "B",
														  "type": "core::ec::EcPoint"
														},
														{
														  "name": "c",
														  "type": "core::felt252"
														},
														{
														  "name": "proof_data",
														  "type": "core::array::Array::<core::felt252>"
														}
													  ],
													  "outputs": [],
													  "state_mutability": "external"
													}
												  ]`,
			},
			want: map[string]any{
				"A": map[string]string{
					"x": "0x3c0c5cb60d57873e6d74d4240300c5d0601303595b92bea80ca2f9e59aa3692",
					"y": "0x2a2b93d24a5278aaff4a3ec913a3d616f21d8c30f9a3b2ab123a4f24aad75de",
				},
				"B": map[string]string{
					"x": "0xd545533296acb37934b889a95b392f8021ef61633a010da565a68975be8f55",
					"y": "0x5498d2c9437143a390c42532ee70c750573a8a196d0070a8e478afd6f25e03c",
				},
				"c": "0x7a83193b9f136a80445649a764051bbb004bd7bce7089edef9e1ae5226ff36",
				"proof_data": []any{
					"0x4a4d170ac5d8749bf7b7230d44b64d449f83802c0a4d8a55df5b27da5730ad",
					"0xce0bf2d1c917a039b2663409bbfe749d2bb7a0ae4019ff32e7e69e88c8c9fb",
					"0x2ca95f97a154230ed28be357dae4c73b635517dac78c78a3765e855b5d93402",
					"0x3f910281f29d7078852cadc7c261bcb86e24d85e717adf9f60b5e96b4ecc652",
					"0x1f8fbad42408418f0871372ea5272d16109b77e3a2ff748cd502f86fc06cc31",
					"0x55495050b2b69ba162b13790cf42fcac39627e58f649fcd6fc655b10fe86b4c",
					"0x4c08652376ec0779410ccaa9f28705dedbc273bcb4fac0c2fdaf68d74742c08",
					"0x5f63751a7bfbbb4bc92d30a1d27abd41f41fca6961c052cbfc8b6c708f0ee1b",
					"0xf89a68bb8715d2ca78c37f66bed1b22f13bcefec9b4fa3fa2d1cd5d02b539b",
					"0x180a21515f534b464bb93700c9d558e84a2221b7e2d903edf2867c0dc38027e",
					"0x3f9d8504820a9419714e74c93d8b0f9013058433000953612ad7fae6f4ccd9c",
					"0x63b67c4ad9915a0b47292253ff6afc90c7343ca2e24dfdbcae18893441e49d0",
					"0x59f0fc402eee19c7ed9efcdb9074a9c631b82c8c7d3089a7b9ffb9d1a690eed",
					"0xda3f81b91d292a5c2efd960c57d4a23fa4c003facba8539a76cdddea9cb2b1",
				},
			},
		}, {
			name: "test 7",
			args: args{
				calldata: []string{
					"0x54328a1075b8820eb43caf0caa233923148c983742402dcfc38541dd843d01a",
					"0x1",
				},
				endpoint: "set_user1",
				abi: `[{
											"name": "token_sender::hello::UserData",
											"type": "struct",
											"members": [
											  {
												"name": "address",
												"type": "core::starknet::contract_address::ContractAddress"
											  },
											  {
												"name": "is_claimed",
												"type": "core::bool"
											  }
											]
										  },{
											"name": "token_sender::hello::IHelloStarknet",
											"type": "interface",
											"items": [
											  {
												"name": "set_user1",
												"type": "function",
												"inputs": [
												  {
													"name": "user",
													"type": "token_sender::hello::UserData"
												  }
												],
												"outputs": [],
												"state_mutability": "external"
											  }, {
												"name": "array2ddd_felt",
												"type": "function",
												"inputs": [
												  {
													"name": "testdd",
													"type": "core::array::Array::<core::array::Array::<core::felt252>>"
												  }
												],
												"outputs": [
												  {
													"type": "core::felt252"
												  }
												],
												"state_mutability": "view"
											  }
											]
										  }]`,
			},
			want: map[string]any{
				"user": map[string]any{
					"address":    "0x54328a1075b8820eb43caf0caa233923148c983742402dcfc38541dd843d01a",
					"is_claimed": true,
				},
			},
		}, {
			name: "test 8",
			args: args{
				calldata: []string{
					"0x52b3783eb2ac90130f536c7202eaa8fddb99746d31676cbde3d4a028e6eb967",
					"0xc6001be915d5bae7bb5940bd3787e876",
					"0x62",
					"0x1",
					"0x0",
					"0x414721218ead2f830421dfba040037ffa3d16e34c8b2d78f18f7d8c164938af",
					"0x2",
					"0x21cbf489c776a18afa5ff5b0dfd6e68d8e154eb078e15274464c239a16b6912",
					"0x5424903a62dede5fa3163c6fd76b0c5059091fe4729c8b764239086af72f0a7",
				},
				endpoint: "redeem_voucher",
				abi: `[{
										"name": "redeem_voucher",
										"type": "function",
										"inputs": [
										  {
											"name": "voucher",
											"type": "rules_tokens::core::voucher::Voucher"
										  },
										  {
											"name": "signature",
											"type": "core::array::Span::<core::felt252>"
										  }
										],
										"outputs": [],
										"state_mutability": "external"
									  },{
										"name": "rules_tokens::core::voucher::Voucher",
										"type": "struct",
										"members": [
										  {
											"name": "receiver",
											"type": "core::starknet::contract_address::ContractAddress"
										  },
										  {
											"name": "token_id",
											"type": "core::integer::u256"
										  },
										  {
											"name": "amount",
											"type": "core::integer::u256"
										  },
										  {
											"name": "salt",
											"type": "core::felt252"
										  }
										]
									  }]`,
			},
			want: map[string]any{
				"voucher": map[string]any{
					"receiver": "0x52b3783eb2ac90130f536c7202eaa8fddb99746d31676cbde3d4a028e6eb967",
					"token_id": "33610859667509361495705047181898111510646",
					"amount":   "1",
					"salt":     "0x414721218ead2f830421dfba040037ffa3d16e34c8b2d78f18f7d8c164938af",
				},
				"signature": []any{
					"0x21cbf489c776a18afa5ff5b0dfd6e68d8e154eb078e15274464c239a16b6912",
					"0x5424903a62dede5fa3163c6fd76b0c5059091fe4729c8b764239086af72f0a7",
				},
			},
		}, {
			name: "test 9",
			args: args{
				calldata: []string{
					"0x2",
					"0x5b456031650c1de9eec123d3e7d06a684d321a346d4a7cac9fd86c2b77cf70f",
					"0x34ca17ba0da4a2d6eb37fc61d82b68793d4b3bb23aeef721bc6045152fcfc00",
					"0x9",
					"0x52b3783eb2ac90130f536c7202eaa8fddb99746d31676cbde3d4a028e6eb967",
					"0xc6001be915d5bae7bb5940bd3787e876",
					"0x62",
					"0x1",
					"0x0",
					"0x414721218ead2f830421dfba040037ffa3d16e34c8b2d78f18f7d8c164938af",
					"0x2",
					"0x21cbf489c776a18afa5ff5b0dfd6e68d8e154eb078e15274464c239a16b6912",
					"0x5424903a62dede5fa3163c6fd76b0c5059091fe4729c8b764239086af72f0a7",
					"0x5b456031650c1de9eec123d3e7d06a684d321a346d4a7cac9fd86c2b77cf70f",
					"0x3556ee435402e506fc85acb898a9acb9daf2855fdec20673ec29a8cb1196cb7",
					"0x9",
					"0x52b3783eb2ac90130f536c7202eaa8fddb99746d31676cbde3d4a028e6eb967",
					"0x347e0e06f7dd8594d51834407f3c7e19a4b6552e822a0040add8b0edcdec280",
					"0x1",
					"0xc6001be915d5bae7bb5940bd3787e876",
					"0x62",
					"0x1",
					"0x1",
					"0x0",
					"0x0",
				},
				endpoint: "__execute__",
				abi:      `[{"type": "function", "name": "constructor", "inputs": [{"name": "signer_public_key_", "type": "core::felt252"}, {"name": "guardian_public_key_", "type": "core::felt252"}], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "upgrade", "inputs": [{"name": "new_implementation", "type": "core::starknet::class_hash::ClassHash"}], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "get_version", "inputs": [], "outputs": [{"type": "core::felt252"}], "state_mutability": "view"}, {"type": "function", "name": "get_signer_public_key", "inputs": [], "outputs": [{"type": "core::felt252"}], "state_mutability": "view"}, {"type": "function", "name": "get_guardian_public_key", "inputs": [], "outputs": [{"type": "core::felt252"}], "state_mutability": "view"}, {"type": "function", "name": "get_signer_escape_activation_date", "inputs": [], "outputs": [{"type": "core::integer::u64"}], "state_mutability": "view"}, {"type": "function", "name": "set_signer_public_key", "inputs": [{"name": "new_public_key", "type": "core::felt252"}], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "set_guardian_public_key", "inputs": [{"name": "new_public_key", "type": "core::felt252"}], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "supports_interface", "inputs": [{"name": "interface_id", "type": "core::integer::u32"}], "outputs": [{"type": "core::bool"}], "state_mutability": "view"}, {"type": "function", "name": "is_valid_signature", "inputs": [{"name": "message", "type": "core::felt252"}, {"name": "signature", "type": "core::array::Span::<core::felt252>"}], "outputs": [{"type": "core::integer::u32"}], "state_mutability": "view"}, {"type": "struct", "name": "rules_account::account::interface::Call", "members": [{"name": "to", "type": "core::starknet::contract_address::ContractAddress"}, {"name": "selector", "type": "core::felt252"}, {"name": "calldata", "type": "core::array::Array::<core::felt252>"}]}, {"type": "function", "name": "__execute__", "inputs": [{"name": "calls", "type": "core::array::Array::<rules_account::account::interface::Call>"}], "outputs": [{"type": "core::array::Array::<core::array::Span::<core::felt252>>"}], "state_mutability": "external"}, {"type": "function", "name": "__validate__", "inputs": [{"name": "calls", "type": "core::array::Array::<rules_account::account::interface::Call>"}], "outputs": [{"type": "core::felt252"}], "state_mutability": "external"}, {"type": "function", "name": "__validate_declare__", "inputs": [{"name": "class_hash", "type": "core::felt252"}], "outputs": [{"type": "core::felt252"}], "state_mutability": "external"}, {"type": "function", "name": "__validate_deploy__", "inputs": [{"name": "class_hash", "type": "core::felt252"}, {"name": "contract_address_salt", "type": "core::felt252"}, {"name": "signer_public_key_", "type": "core::felt252"}, {"name": "guardian_public_key_", "type": "core::felt252"}], "outputs": [{"type": "core::felt252"}], "state_mutability": "external"}, {"type": "function", "name": "trigger_signer_escape", "inputs": [], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "cancel_escape", "inputs": [], "outputs": [], "state_mutability": "external"}, {"type": "function", "name": "escape_signer", "inputs": [{"name": "new_public_key", "type": "core::felt252"}], "outputs": [], "state_mutability": "external"}, {"type": "event", "name": "AccountUpgraded", "inputs": [{"name": "new_implementation", "type": "core::starknet::class_hash::ClassHash"}]}, {"type": "event", "name": "AccountInitialized", "inputs": [{"name": "signer_public_key", "type": "core::felt252"}, {"name": "guardian_public_key", "type": "core::felt252"}]}, {"type": "event", "name": "SignerPublicKeyChanged", "inputs": [{"name": "new_public_key", "type": "core::felt252"}]}, {"type": "event", "name": "GuardianPublicKeyChanged", "inputs": [{"name": "new_public_key", "type": "core::felt252"}]}, {"type": "event", "name": "SignerEscapeTriggered", "inputs": [{"name": "active_at", "type": "core::integer::u64"}]}, {"type": "event", "name": "SignerEscaped", "inputs": [{"name": "new_public_key", "type": "core::felt252"}]}, {"type": "event", "name": "EscapeCanceled", "inputs": []}]`,
			},
			want: map[string]any{
				"calls": []any{
					map[string]any{
						"to":       "0x5b456031650c1de9eec123d3e7d06a684d321a346d4a7cac9fd86c2b77cf70f",
						"selector": "0x34ca17ba0da4a2d6eb37fc61d82b68793d4b3bb23aeef721bc6045152fcfc00",
						"calldata": []any{
							"0x52b3783eb2ac90130f536c7202eaa8fddb99746d31676cbde3d4a028e6eb967",
							"0xc6001be915d5bae7bb5940bd3787e876",
							"0x62",
							"0x1",
							"0x0",
							"0x414721218ead2f830421dfba040037ffa3d16e34c8b2d78f18f7d8c164938af",
							"0x2",
							"0x21cbf489c776a18afa5ff5b0dfd6e68d8e154eb078e15274464c239a16b6912",
							"0x5424903a62dede5fa3163c6fd76b0c5059091fe4729c8b764239086af72f0a7",
						},
					}, map[string]any{
						"to":       "0x5b456031650c1de9eec123d3e7d06a684d321a346d4a7cac9fd86c2b77cf70f",
						"selector": "0x3556ee435402e506fc85acb898a9acb9daf2855fdec20673ec29a8cb1196cb7",
						"calldata": []any{
							"0x52b3783eb2ac90130f536c7202eaa8fddb99746d31676cbde3d4a028e6eb967",
							"0x347e0e06f7dd8594d51834407f3c7e19a4b6552e822a0040add8b0edcdec280",
							"0x1",
							"0xc6001be915d5bae7bb5940bd3787e876",
							"0x62",
							"0x1",
							"0x1",
							"0x0",
							"0x0",
						},
					},
				},
			},
		}, {
			name: "test 10",
			args: args{
				calldata: []string{
					"0x1",
					"0x454f0bd015e730e5adbb4f080b075fdbf55654ff41ee336203aa2e1ac4d4309",
					"0x32a99297e1d12a9b91d4f90d5dd4b160d93c84a9e3b4daa916fec14ec852e05",
					"0x2",
					"0x333338373738616361363538666133353533666137303662313132323661303",
					"0x366431356532333239643637636137613166323061343166643763316161653",
				},
				endpoint: "__execute__",
				abi: `[
							{
								"type": "struct",
								"name": "core::starknet::account::Call",
								"members": [
									{
										"name": "to",
										"type": "core::starknet::contract_address::ContractAddress"
									},
									{
										"name": "selector",
										"type": "core::felt252"
									},
									{
										"name": "calldata",
										"type": "core::array::Array::<core::felt252>"
									}
								]
							},
							{
								"type": "function",
								"name": "__execute__",
								"inputs": [
									{
										"name": "calls",
										"type": "core::array::Array::<core::starknet::account::Call>"
									}
								],
								"outputs": [
									{
										"type": "core::array::Array::<core::array::Span::<core::felt252>>"
									}
								],
								"state_mutability": "external"
							}
						]`,
			},
			want: map[string]any{
				"calls": []any{
					map[string]any{
						"to":       "0x454f0bd015e730e5adbb4f080b075fdbf55654ff41ee336203aa2e1ac4d4309",
						"selector": "0x32a99297e1d12a9b91d4f90d5dd4b160d93c84a9e3b4daa916fec14ec852e05",
						"calldata": []any{
							"0x333338373738616361363538666133353533666137303662313132323661303",
							"0x366431356532333239643637636137613166323061343166643763316161653",
						},
					},
				},
			},
		}, {
			name: "test 11",
			args: args{
				calldata: []string{
					"0x1",
					"0x53c7c49260f079bf5b79dfc05227e1da316429a0",
					"0x4",
					"0x1",
					"0x2",
					"0x3",
					"0x4",
					"0x5678",
					"0x0",
					"0x4",
					"0xffffffffffffffffffffffffffffffffffffffffff",
					"0x1234",
					"0x5678",
					"0x9abc",
				},
				endpoint: "propose",
				abi: `[
						{
							"name": "sx::types::strategy::Strategy",
							"type": "struct",
							"members": [
							  {
								"name": "address",
								"type": "core::starknet::contract_address::ContractAddress"
							  },
							  {
								"name": "params",
								"type": "core::array::Array::<core::felt252>"
							  }
							]
						},
						{
							"name": "core::starknet::eth_address::EthAddress",
							"type": "struct",
							"members": [
							  {
								"name": "address",
								"type": "core::felt252"
							  }
							]
						},
						{
							"name": "sx::types::user_address::UserAddress",
							"type": "enum",
							"variants": [
							  {
								"name": "Starknet",
								"type": "core::starknet::contract_address::ContractAddress"
							  },
							  {
								"name": "Ethereum",
								"type": "core::starknet::eth_address::EthAddress"
							  },
							  {
								"name": "Custom",
								"type": "core::integer::u256"
							  }
							]
						},
						{
							"name": "propose",
							"type": "function",
							"inputs": [
							  {
								"name": "author",
								"type": "sx::types::user_address::UserAddress"
							  },
							  {
								"name": "metadata_uri",
								"type": "core::array::Array::<core::felt252>"
							  },
							  {
								"name": "execution_strategy",
								"type": "sx::types::strategy::Strategy"
							  },
							  {
								"name": "user_proposal_validation_params",
								"type": "core::array::Array::<core::felt252>"
							  }
							],
							"outputs": [],
							"state_mutability": "external"
						}
					]`,
			},
			want: map[string]any{
				"author": map[string]any{
					"Ethereum": map[string]any{
						"address": "0x53c7c49260f079bf5b79dfc05227e1da316429a0",
					},
				},
				"execution_strategy": map[string]any{
					"address": "0x5678",
					"params":  []any{},
				},
				"metadata_uri": []any{"0x1", "0x2", "0x3", "0x4"},
				"user_proposal_validation_params": []any{
					"0xffffffffffffffffffffffffffffffffffffffffff", "0x1234", "0x5678", "0x9abc",
				},
			},
		}, {
			name: "test 12",
			args: args{
				calldata: []string{
					"0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7",
					"0x53c91253bc9682c04929ca02ed00b3e423f6710d2ee7e0d5ebb06f3ecf368a8",
					"0x20c49ba5e353f80000000000000000",
					"0x3e8",
					"0x0",
					"0x132bb70",
					"0x1",
					"0x12fae30",
					"0x1",
					"0x566875f644c",
				},
				endpoint: "mint_and_deposit",
				abi: `[
					{
						"name": "core::bool",
						"type": "enum",
						"variants": [
						  {
							"name": "False",
							"type": "()"
						  },
						  {
							"name": "True",
							"type": "()"
						  }
						]
					},
					{
						"name": "ekubo::types::i129::i129",
						"type": "struct",
						"members": [
						  {
							"name": "mag",
							"type": "core::integer::u128"
						  },
						  {
							"name": "sign",
							"type": "core::bool"
						  }
						]
					},
					{
						"name": "ekubo::types::bounds::Bounds",
						"type": "struct",
						"members": [
						  {
							"name": "lower",
							"type": "ekubo::types::i129::i129"
						  },
						  {
							"name": "upper",
							"type": "ekubo::types::i129::i129"
						  }
						]
					},
					{
					  "name": "ekubo::types::keys::PoolKey",
					  "type": "struct",
					  "members": [
						{
						  "name": "token0",
						  "type": "core::starknet::contract_address::ContractAddress"
						},
						{
						  "name": "token1",
						  "type": "core::starknet::contract_address::ContractAddress"
						},
						{
						  "name": "fee",
						  "type": "core::integer::u128"
						},
						{
						  "name": "tick_spacing",
						  "type": "core::integer::u128"
						},
						{
						  "name": "extension",
						  "type": "core::starknet::contract_address::ContractAddress"
						}
					  ]
					},
					{
						"name": "mint_and_deposit",
						"type": "function",
						"inputs": [
						  {
							"name": "pool_key",
							"type": "ekubo::types::keys::PoolKey"
						  },
						  {
							"name": "bounds",
							"type": "ekubo::types::bounds::Bounds"
						  },
						  {
							"name": "min_liquidity",
							"type": "core::integer::u128"
						  }
						],
						"outputs": [
						  {
							"type": "(core::integer::u64, core::integer::u128)"
						  }
						],
						"state_mutability": "external"
					}
				]`,
			},
			want: map[string]any{
				"pool_key": map[string]any{
					"token0":       "0x49d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7",
					"token1":       "0x53c91253bc9682c04929ca02ed00b3e423f6710d2ee7e0d5ebb06f3ecf368a8",
					"fee":          "170141183460469235273462165868118016",
					"tick_spacing": "1000",
					"extension":    "0x0",
				},
				"bounds": map[string]any{
					"lower": map[string]any{
						"mag":  "20102000",
						"sign": "True",
					},
					"upper": map[string]any{
						"mag":  "19902000",
						"sign": "True",
					},
				},
				"min_liquidity": "5937915978828",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var a Abi
			err := json.Unmarshal([]byte(tt.args.abi), &a)
			require.NoError(t, err)

			got, err := DecodeFunctionCallData(tt.args.calldata, *a.Functions[tt.args.endpoint], a.Structs, a.Enums)
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
