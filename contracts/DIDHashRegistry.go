// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"HashUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"didToHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"setHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506109888061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80635b6beeb914610043578063d61983e714610073578063e15fe023146100a3575b5f5ffd5b61005d60048036038101906100589190610319565b6100bf565b60405161006a91906103d4565b60405180910390f35b61008d6004803603810190610088919061051c565b61016f565b60405161009a91906103d4565b60405180910390f35b6100bd60048036038101906100b89190610563565b610221565b005b60605f83836040516100d292919061060f565b908152602001604051809103902080546100eb90610654565b80601f016020809104026020016040519081016040528092919081815260200182805461011790610654565b80156101625780601f1061013957610100808354040283529160200191610162565b820191905f5260205f20905b81548152906001019060200180831161014557829003601f168201915b5050505050905092915050565b5f818051602081018201805184825260208301602085012081835280955050505050505f9150905080546101a290610654565b80601f01602080910402602001604051908101604052809291908181526020018280546101ce90610654565b80156102195780601f106101f057610100808354040283529160200191610219565b820191905f5260205f20905b8154815290600101906020018083116101fc57829003601f168201915b505050505081565b81815f868660405161023492919061060f565b9081526020016040518091039020918261024f929190610837565b50838360405161026092919061060f565b60405180910390207fda8d62b518e9adca07dd6211c684efee3f3e9eead56c1c628228c081766ada558383604051610299929190610930565b60405180910390a250505050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126102d9576102d86102b8565b5b8235905067ffffffffffffffff8111156102f6576102f56102bc565b5b602083019150836001820283011115610312576103116102c0565b5b9250929050565b5f5f6020838503121561032f5761032e6102b0565b5b5f83013567ffffffffffffffff81111561034c5761034b6102b4565b5b610358858286016102c4565b92509250509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6103a682610364565b6103b0818561036e565b93506103c081856020860161037e565b6103c98161038c565b840191505092915050565b5f6020820190508181035f8301526103ec818461039c565b905092915050565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61042e8261038c565b810181811067ffffffffffffffff8211171561044d5761044c6103f8565b5b80604052505050565b5f61045f6102a7565b905061046b8282610425565b919050565b5f67ffffffffffffffff82111561048a576104896103f8565b5b6104938261038c565b9050602081019050919050565b828183375f83830152505050565b5f6104c06104bb84610470565b610456565b9050828152602081018484840111156104dc576104db6103f4565b5b6104e78482856104a0565b509392505050565b5f82601f830112610503576105026102b8565b5b81356105138482602086016104ae565b91505092915050565b5f60208284031215610531576105306102b0565b5b5f82013567ffffffffffffffff81111561054e5761054d6102b4565b5b61055a848285016104ef565b91505092915050565b5f5f5f5f6040858703121561057b5761057a6102b0565b5b5f85013567ffffffffffffffff811115610598576105976102b4565b5b6105a4878288016102c4565b9450945050602085013567ffffffffffffffff8111156105c7576105c66102b4565b5b6105d3878288016102c4565b925092505092959194509250565b5f81905092915050565b5f6105f683856105e1565b93506106038385846104a0565b82840190509392505050565b5f61061b8284866105eb565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061066b57607f821691505b60208210810361067e5761067d610627565b5b50919050565b5f82905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026106ea7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826106af565b6106f486836106af565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61073861073361072e8461070c565b610715565b61070c565b9050919050565b5f819050919050565b6107518361071e565b61076561075d8261073f565b8484546106bb565b825550505050565b5f5f905090565b61077c61076d565b610787818484610748565b505050565b5b818110156107aa5761079f5f82610774565b60018101905061078d565b5050565b601f8211156107ef576107c08161068e565b6107c9846106a0565b810160208510156107d8578190505b6107ec6107e4856106a0565b83018261078c565b50505b505050565b5f82821c905092915050565b5f61080f5f19846008026107f4565b1980831691505092915050565b5f6108278383610800565b9150826002028217905092915050565b6108418383610684565b67ffffffffffffffff81111561085a576108596103f8565b5b6108648254610654565b61086f8282856107ae565b5f601f83116001811461089c575f841561088a578287013590505b610894858261081c565b8655506108fb565b601f1984166108aa8661068e565b5f5b828110156108d1578489013582556001820191506020850194506020810190506108ac565b868310156108ee57848901356108ea601f891682610800565b8355505b6001600288020188555050505b50505050505050565b5f61090f838561036e565b935061091c8385846104a0565b6109258361038c565b840190509392505050565b5f6020820190508181035f830152610949818486610904565b9050939250505056fea26469706673582212208ed6ee2dd4f4a1589d045f22e4f9409b86c921b06a2cbf1f97a1882da958b88664736f6c634300081d0033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// DidToHash is a free data retrieval call binding the contract method 0xd61983e7.
//
// Solidity: function didToHash(string ) view returns(string)
func (_Contracts *ContractsCaller) DidToHash(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "didToHash", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DidToHash is a free data retrieval call binding the contract method 0xd61983e7.
//
// Solidity: function didToHash(string ) view returns(string)
func (_Contracts *ContractsSession) DidToHash(arg0 string) (string, error) {
	return _Contracts.Contract.DidToHash(&_Contracts.CallOpts, arg0)
}

// DidToHash is a free data retrieval call binding the contract method 0xd61983e7.
//
// Solidity: function didToHash(string ) view returns(string)
func (_Contracts *ContractsCallerSession) DidToHash(arg0 string) (string, error) {
	return _Contracts.Contract.DidToHash(&_Contracts.CallOpts, arg0)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string did) view returns(string)
func (_Contracts *ContractsCaller) GetHash(opts *bind.CallOpts, did string) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getHash", did)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string did) view returns(string)
func (_Contracts *ContractsSession) GetHash(did string) (string, error) {
	return _Contracts.Contract.GetHash(&_Contracts.CallOpts, did)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string did) view returns(string)
func (_Contracts *ContractsCallerSession) GetHash(did string) (string, error) {
	return _Contracts.Contract.GetHash(&_Contracts.CallOpts, did)
}

// SetHash is a paid mutator transaction binding the contract method 0xe15fe023.
//
// Solidity: function setHash(string did, string hash) returns()
func (_Contracts *ContractsTransactor) SetHash(opts *bind.TransactOpts, did string, hash string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setHash", did, hash)
}

// SetHash is a paid mutator transaction binding the contract method 0xe15fe023.
//
// Solidity: function setHash(string did, string hash) returns()
func (_Contracts *ContractsSession) SetHash(did string, hash string) (*types.Transaction, error) {
	return _Contracts.Contract.SetHash(&_Contracts.TransactOpts, did, hash)
}

// SetHash is a paid mutator transaction binding the contract method 0xe15fe023.
//
// Solidity: function setHash(string did, string hash) returns()
func (_Contracts *ContractsTransactorSession) SetHash(did string, hash string) (*types.Transaction, error) {
	return _Contracts.Contract.SetHash(&_Contracts.TransactOpts, did, hash)
}

// ContractsHashUpdatedIterator is returned from FilterHashUpdated and is used to iterate over the raw logs and unpacked data for HashUpdated events raised by the Contracts contract.
type ContractsHashUpdatedIterator struct {
	Event *ContractsHashUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractsHashUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsHashUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractsHashUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractsHashUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsHashUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsHashUpdated represents a HashUpdated event raised by the Contracts contract.
type ContractsHashUpdated struct {
	Did  common.Hash
	Hash string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHashUpdated is a free log retrieval operation binding the contract event 0xda8d62b518e9adca07dd6211c684efee3f3e9eead56c1c628228c081766ada55.
//
// Solidity: event HashUpdated(string indexed did, string hash)
func (_Contracts *ContractsFilterer) FilterHashUpdated(opts *bind.FilterOpts, did []string) (*ContractsHashUpdatedIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "HashUpdated", didRule)
	if err != nil {
		return nil, err
	}
	return &ContractsHashUpdatedIterator{contract: _Contracts.contract, event: "HashUpdated", logs: logs, sub: sub}, nil
}

// WatchHashUpdated is a free log subscription operation binding the contract event 0xda8d62b518e9adca07dd6211c684efee3f3e9eead56c1c628228c081766ada55.
//
// Solidity: event HashUpdated(string indexed did, string hash)
func (_Contracts *ContractsFilterer) WatchHashUpdated(opts *bind.WatchOpts, sink chan<- *ContractsHashUpdated, did []string) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "HashUpdated", didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsHashUpdated)
				if err := _Contracts.contract.UnpackLog(event, "HashUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHashUpdated is a log parse operation binding the contract event 0xda8d62b518e9adca07dd6211c684efee3f3e9eead56c1c628228c081766ada55.
//
// Solidity: event HashUpdated(string indexed did, string hash)
func (_Contracts *ContractsFilterer) ParseHashUpdated(log types.Log) (*ContractsHashUpdated, error) {
	event := new(ContractsHashUpdated)
	if err := _Contracts.contract.UnpackLog(event, "HashUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
