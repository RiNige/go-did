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
	Bin: "0x608060405234801561001057600080fd5b50610801806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80635b6beeb914610046578063d61983e714610076578063e15fe023146100a6575b600080fd5b610060600480360381019061005b919061041b565b6100c2565b60405161006d91906105fa565b60405180910390f35b610090600480360381019061008b91906104e9565b610175565b60405161009d91906105fa565b60405180910390f35b6100c060048036038101906100bb9190610468565b61022b565b005b6060600083836040516100d69291906105bd565b908152602001604051809103902080546100ef906106db565b80601f016020809104026020016040519081016040528092919081815260200182805461011b906106db565b80156101685780601f1061013d57610100808354040283529160200191610168565b820191906000526020600020905b81548152906001019060200180831161014b57829003601f168201915b5050505050905092915050565b60008180516020810182018051848252602083016020850120818352809550505050505060009150905080546101aa906106db565b80601f01602080910402602001604051908101604052809291908181526020018280546101d6906106db565b80156102235780601f106101f857610100808354040283529160200191610223565b820191906000526020600020905b81548152906001019060200180831161020657829003601f168201915b505050505081565b81816000868660405161023f9291906105bd565b9081526020016040518091039020919061025a9291906102b2565b50838360405161026b9291906105bd565b60405180910390207fda8d62b518e9adca07dd6211c684efee3f3e9eead56c1c628228c081766ada5583836040516102a49291906105d6565b60405180910390a250505050565b8280546102be906106db565b90600052602060002090601f0160209004810192826102e05760008555610327565b82601f106102f957803560ff1916838001178555610327565b82800160010185558215610327579182015b8281111561032657823582559160200191906001019061030b565b5b5090506103349190610338565b5090565b5b80821115610351576000816000905550600101610339565b5090565b600061036861036384610641565b61061c565b905082815260208101848484011115610384576103836107ab565b5b61038f848285610699565b509392505050565b60008083601f8401126103ad576103ac6107a1565b5b8235905067ffffffffffffffff8111156103ca576103c961079c565b5b6020830191508360018202830111156103e6576103e56107a6565b5b9250929050565b600082601f830112610402576104016107a1565b5b8135610412848260208601610355565b91505092915050565b60008060208385031215610432576104316107b5565b5b600083013567ffffffffffffffff8111156104505761044f6107b0565b5b61045c85828601610397565b92509250509250929050565b60008060008060408587031215610482576104816107b5565b5b600085013567ffffffffffffffff8111156104a05761049f6107b0565b5b6104ac87828801610397565b9450945050602085013567ffffffffffffffff8111156104cf576104ce6107b0565b5b6104db87828801610397565b925092505092959194509250565b6000602082840312156104ff576104fe6107b5565b5b600082013567ffffffffffffffff81111561051d5761051c6107b0565b5b610529848285016103ed565b91505092915050565b600061053e838561067d565b935061054b838584610699565b610554836107ba565b840190509392505050565b600061056b838561068e565b9350610578838584610699565b82840190509392505050565b600061058f82610672565b610599818561067d565b93506105a98185602086016106a8565b6105b2816107ba565b840191505092915050565b60006105ca82848661055f565b91508190509392505050565b600060208201905081810360008301526105f1818486610532565b90509392505050565b600060208201905081810360008301526106148184610584565b905092915050565b6000610626610637565b9050610632828261070d565b919050565b6000604051905090565b600067ffffffffffffffff82111561065c5761065b61076d565b5b610665826107ba565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b82818337600083830152505050565b60005b838110156106c65780820151818401526020810190506106ab565b838111156106d5576000848401525b50505050565b600060028204905060018216806106f357607f821691505b602082108114156107075761070661073e565b5b50919050565b610716826107ba565b810181811067ffffffffffffffff821117156107355761073461076d565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f830116905091905056fea2646970667358221220bbfd61375088df958083bdf12720b5330eb85f647727e1c0869690935e66b34364736f6c63430008060033",
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
