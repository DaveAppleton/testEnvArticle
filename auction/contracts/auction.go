// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableBin is the compiled bytecode used for deploying new contracts.
var OwnableBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556101cc806100326000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638da5cb5b1461003b578063f2fde38b1461005f575b600080fd5b610043610087565b604080516001600160a01b039092168252519081900360200190f35b6100856004803603602081101561007557600080fd5b50356001600160a01b0316610096565b005b6000546001600160a01b031690565b6000546001600160a01b031633146100e1576040805162461bcd60e51b81526020600482015260096024820152682337b93134b23232b760b91b604482015290519081900360640190fd5b6001600160a01b03811661013c576040805162461bcd60e51b815260206004820152601a60248201527f4e6f6e2d7a65726f20616464726573732072657175697265642e000000000000604482015290519081900360640190fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a723158204fca6cde65352771513de5d8a442e51d976f7bdb52efc0e3b133364c571e315e64736f6c634300050b0032"

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionABI is the input ABI used to generate the binding from.
const AuctionABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"withdrawRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inReveal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberOfRevealedValues\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_randString\",\"type\":\"bytes\"}],\"name\":\"calculateHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"theAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"flagged\",\"type\":\"bool\"}],\"name\":\"setWinningAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startWithdraw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endWithdraw\",\"type\":\"uint256\"}],\"name\":\"startWithdrawal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_winningAmount\",\"type\":\"uint256\"}],\"name\":\"setWinningAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"check\",\"type\":\"address\"}],\"name\":\"isWinner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"winner\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"inPeriod\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inBidding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"biddingTime\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inWithdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"information\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hashValue\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"haveBid\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"revealed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"refunded\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"flagged\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"winners\",\"type\":\"address[]\"}],\"name\":\"earlyWithdrawal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bid\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"randString\",\"type\":\"bytes\"}],\"name\":\"reveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"check\",\"type\":\"address\"}],\"name\":\"withdrawalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferAnyERC20Token\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"revealedValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reveals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"numberOfRevealsForValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startBids\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endBids\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startReveal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endReveal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimumBid\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_wallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_minimumBid\",\"type\":\"uint256\"}],\"name\":\"MinimumBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startBids\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endBids\",\"type\":\"uint256\"}],\"name\":\"BiddingPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startReveal\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endReveal\",\"type\":\"uint256\"}],\"name\":\"RevealPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startWithdraw\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endWithdraw\",\"type\":\"uint256\"}],\"name\":\"WithdrawPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"funding\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"BidSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winningAmount\",\"type\":\"uint256\"}],\"name\":\"WinningAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tiedWinner\",\"type\":\"address\"}],\"name\":\"WinnerWithTie\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"NothingToRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"change\",\"type\":\"uint256\"}],\"name\":\"RefundChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BalanceWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"Wallet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// AuctionFuncSigs maps the 4-byte function signature to its string representation.
var AuctionFuncSigs = map[string]string{
	"a056e620": "biddingTime(bytes32)",
	"500a2b59": "calculateHash(uint256,bytes)",
	"b856ed86": "earlyWithdrawal(address[])",
	"9eff376d": "inBidding()",
	"17108562": "inReveal()",
	"a0ace77a": "inWithdraw()",
	"a30ab2d0": "information(address)",
	"9d9ca28d": "isWinner(address)",
	"d3a86386": "minimumBid()",
	"41361284": "numberOfRevealedValues()",
	"ed295415": "numberOfRevealsForValue(uint256)",
	"8da5cb5b": "owner()",
	"ce805642": "reveal(uint256,bytes)",
	"e44d7a69": "revealedValue(uint256)",
	"e69650a6": "reveals(uint256,uint256)",
	"51c65bdf": "setWinningAddresses(address[],bool)",
	"7eb4d7d3": "setWinningAmount(uint256)",
	"70431d29": "startWithdrawal(uint256,uint256)",
	"dc39d06d": "transferAnyERC20Token(address,uint256)",
	"f2fde38b": "transferOwnership(address)",
	"476343ee": "withdrawFees()",
	"110f8874": "withdrawRefund()",
	"daa91f2e": "withdrawalAmount(address)",
}

// AuctionBin is the compiled bytecode used for deploying new contracts.
var AuctionBin = "0x60806040523480156200001157600080fd5b5060405162001fc638038062001fc6833981810160405260c08110156200003757600080fd5b508051602082015160408301516060840151608085015160a090950151600080546001600160a01b03191633179055939492939192909185620000ca57604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526016602482015260008051602062001fa6833981519152604482015290519081900360640190fd5b846200012657604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526016602482015260008051602062001fa6833981519152604482015290519081900360640190fd5b836200018257604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526016602482015260008051602062001fa6833981519152604482015290519081900360640190fd5b82620001de57604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526016602482015260008051602062001fa6833981519152604482015290519081900360640190fd5b816200024b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f6d696e696d756d20626964206d757374206265206e6f6e207a65726f00000000604482015290519081900360640190fd5b6001600160a01b038116620002c157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e76616c69642077616c6c6574206164647265737300000000000000000000604482015290519081900360640190fd5b6005869055600685905560078490556008839055600c829055600d80546001600160a01b0319166001600160a01b038316179055604080518781526020810187905281517f892ffcfb665b4ba99a581f2bfa4b8614aaa46bae9247c832b4bd6c1d2badb2e6929181900390910190a17fd23f2f662aba8ca991baba62f1219f0780865580d293d0c7bc89079757639333600754600854604051808381526020018281526020019250505060405180910390a1600c5460408051918252517f855d14e5037e9de69ed821e8030c2c0ccaf36f6ebdb3fdc4689f51aa85a518f09181900360200190a1600d54604080516001600160a01b039092168252517f299e6b07aa8d613d084182e45e438778c5e03395c860c2f28273c66e6f620f0c9181900360200190a1505050505050611ba980620003fd6000396000f3fe6080604052600436106101405760003560e01c8063a056e620116100b6578063daa91f2e1161006f578063daa91f2e14610633578063dc39d06d1461067f578063e44d7a69146106b8578063e69650a6146106e2578063ed29541514610712578063f2fde38b1461073c57610140565b8063a056e62014610417578063a0ace77a14610434578063a30ab2d014610449578063b856ed86146104b8578063ce80564214610566578063d3a863861461061e57610140565b806351c65bdf1161010857806351c65bdf1461027957806370431d29146103295780637eb4d7d3146103595780638da5cb5b146103835780639d9ca28d146103b45780639eff376d1461040257610140565b8063110f887414610145578063171085621461015c5780634136128414610185578063476343ee146101ac578063500a2b59146101c1575b600080fd5b34801561015157600080fd5b5061015a61076f565b005b34801561016857600080fd5b50610171610a2f565b604080519115158252519081900360200190f35b34801561019157600080fd5b5061019a610a48565b60408051918252519081900360200190f35b3480156101b857600080fd5b5061015a610a4e565b3480156101cd57600080fd5b5061019a600480360360408110156101e457600080fd5b81359190810190604081016020820135600160201b81111561020557600080fd5b82018360208201111561021757600080fd5b803590602001918460018302840111600160201b8311171561023857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b6a945050505050565b34801561028557600080fd5b5061015a6004803603604081101561029c57600080fd5b810190602081018135600160201b8111156102b657600080fd5b8201836020820111156102c857600080fd5b803590602001918460208302840111600160201b831117156102e957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295505050503515159050610bea565b34801561033557600080fd5b5061015a6004803603604081101561034c57600080fd5b5080359060200135610d49565b34801561036557600080fd5b5061015a6004803603602081101561037c57600080fd5b5035610e6b565b34801561038f57600080fd5b50610398610f42565b604080516001600160a01b039092168252519081900360200190f35b3480156103c057600080fd5b506103e7600480360360208110156103d757600080fd5b50356001600160a01b0316610f51565b60408051921515835290151560208301528051918290030190f35b34801561040e57600080fd5b50610171610fef565b61015a6004803603602081101561042d57600080fd5b5035611006565b34801561044057600080fd5b506101716111c2565b34801561045557600080fd5b5061047c6004803603602081101561046c57600080fd5b50356001600160a01b03166111d9565b60408051978852602088019690965293151586860152911515606086015215156080850152151560a084015260c0830152519081900360e00190f35b3480156104c457600080fd5b5061015a600480360360208110156104db57600080fd5b810190602081018135600160201b8111156104f557600080fd5b82018360208201111561050757600080fd5b803590602001918460208302840111600160201b8311171561052857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611221945050505050565b34801561057257600080fd5b5061015a6004803603604081101561058957600080fd5b81359190810190604081016020820135600160201b8111156105aa57600080fd5b8201836020820111156105bc57600080fd5b803590602001918460018302840111600160201b831117156105dd57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611493945050505050565b34801561062a57600080fd5b5061019a611757565b34801561063f57600080fd5b506106666004803603602081101561065657600080fd5b50356001600160a01b031661175d565b6040805192835290151560208301528051918290030190f35b34801561068b57600080fd5b50610171600480360360408110156106a257600080fd5b506001600160a01b038135169060200135611852565b3480156106c457600080fd5b5061019a600480360360208110156106db57600080fd5b5035611938565b3480156106ee57600080fd5b506103986004803603604081101561070557600080fd5b50803590602001356119a9565b34801561071e57600080fd5b5061019a6004803603602081101561073557600080fd5b50356119de565b34801561074857600080fd5b5061015a6004803603602081101561075f57600080fd5b50356001600160a01b03166119f0565b60095442116107c5576040805162461bcd60e51b815260206004820152601c60248201527f576974686472617720686173206e6f7420737461727465642079657400000000604482015290519081900360640190fd5b600a544210610812576040805162461bcd60e51b81526020600482015260146024820152732bb4ba34323930bb903832b934b7b21037bb32b960611b604482015290519081900360640190fd5b336000908152600160205260409020600381015461086a576040805162461bcd60e51b815260206004820152601060248201526f139bc8189a59081cdd589b5a5d1d195960821b604482015290519081900360640190fd5b6002810154610100900460ff166108bf576040805162461bcd60e51b8152602060048201526014602482015273189a59081dd85cc81b9bdd081c995d99585b195960621b604482015290519081900360640190fd5b600281015462010000900460ff1615610912576040805162461bcd60e51b815260206004820152601060248201526f105b1c9958591e481c99599d5b99195960821b604482015290519081900360640190fd5b60028101805462ff00001916620100001790556000806109313361175d565b90925090508061097b57604080513381526020810184905281517fbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d929181900390910190a16109f5565b81156109c157604080513381526020810184905281517f7417b321c40d8a0c375656909a6cb0a232297cc6addda0f3a386b9f761669d90929181900390910190a16109f5565b6040805133815290517fbb386939e3dac174d852c26db4592a53fb01fde4794fa3124c31416a0c9b9ec09181900360200190a15b8115610a2a57604051339083156108fc029084906000818181858888f19350505050158015610a28573d6000803e3d6000fd5b505b505050565b600060075442118015610a43575060085442105b905090565b60045490565b600a544211610a9b576040805162461bcd60e51b81526020600482015260146024820152732bb4ba34323930bb903832b934b7b21037bb32b960611b604482015290519081900360640190fd5b6000546001600160a01b03163314610ae6576040805162461bcd60e51b81526020600482015260096024820152682337b93134b23232b760b91b604482015290519081900360640190fd5b600d54604080516001600160a01b0390921682523031602083015280517fddc398b321237a8d40ac914388309c2f52a08c134e4dc4ce61e32f57cb7d80f19281900390910190a1600d546040516001600160a01b0390911690303180156108fc02916000818181858888f19350505050158015610b67573d6000803e3d6000fd5b50565b600082826040516020018083815260200182805190602001908083835b60208310610ba65780518252601f199092019160209182019101610b87565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405280519060200120905092915050565b6000546001600160a01b03163314610c35576040805162461bcd60e51b81526020600482015260096024820152682337b93134b23232b760b91b604482015290519081900360640190fd5b6008544211610c86576040805162461bcd60e51b815260206004820152601860248201527714995d99585b081a185cc81b9bdd08195b991959081e595d60421b604482015290519081900360640190fd5b60005b8251811015610a2a578160016000858481518110610ca357fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060020160036101000a81548160ff0219169083151502179055507fdeb65b46e26a4ed6abe63bd73c4613e1262a81bafb7bb35b92f5160e8f1724f0838281518110610d1257fe5b602002602001015160405180826001600160a01b03166001600160a01b0316815260200191505060405180910390a1600101610c89565b6000546001600160a01b03163314610d94576040805162461bcd60e51b81526020600482015260096024820152682337b93134b23232b760b91b604482015290519081900360640190fd5b6000600b5411610de4576040805162461bcd60e51b815260206004820152601660248201527515da5b9b9a5b99c8105b5bdd5b9d08139bdd0814d95d60521b604482015290519081900360640190fd5b428211610e225760405162461bcd60e51b8152600401808060200182810382526023815260200180611b2e6023913960400191505060405180910390fd5b6009829055600a819055604080518381526020810183905281517f6efecf63a32b9b51e8ead9e19a996f69916a2753adab220b8bced7624fb50e34929181900390910190a15050565b6000546001600160a01b03163314610eb6576040805162461bcd60e51b81526020600482015260096024820152682337b93134b23232b760b91b604482015290519081900360640190fd5b6008544211610f07576040805162461bcd60e51b815260206004820152601860248201527714995d99585b081a185cc81b9bdd08195b991959081e595d60421b604482015290519081900360640190fd5b600b8190556040805182815290517f85679d62923fa647b865510f1c0f24c7babbee0a25eca9118b226faf71f42f1c9181900360200190a150565b6000546001600160a01b031690565b600080600954421015610f6957506000905080610fea565b50600a54600b546001600160a01b0384166000908152600160205260409020600301544292909210911180610fe75750600b546001600160a01b038416600090815260016020526040902060030154148015610fe757506001600160a01b0383166000908152600160205260409020600201546301000000900460ff165b91505b915091565b600060055442118015610a43575050600654421090565b60055442101561105d576040805162461bcd60e51b815260206004820152601760248201527f42696464696e67206e6f74207374617274656420796574000000000000000000604482015290519081900360640190fd5b60065442106110a7576040805162461bcd60e51b8152602060048201526011602482015270109a59191a5b99c81a185cc8195b991959607a1b604482015290519081900360640190fd5b3360009081526001602052604090206002015460ff161561110f576040805162461bcd60e51b815260206004820152601860248201527f6f6e6c79206f6e65206269642070657220616464726573730000000000000000604482015290519081900360640190fd5b600c543410156111505760405162461bcd60e51b8152600401808060200182810382526024815260200180611b516024913960400191505060405180910390fd5b336000818152600160208181526040928390208581553481840181905560028201805460ff1916909417909355835194855290840191909152828201849052905190917f43dfb6c14d1b44aaeaef2c0487e8423e2843a99f9e09108814b95ce236ed8707919081900360600190a15050565b600060095442118015610a43575050600a54421090565b60016020819052600091825260409091208054918101546002820154600390920154909160ff8082169261010083048216926201000081048316926301000000909104169087565b6000546001600160a01b0316331461126c576040805162461bcd60e51b81526020600482015260096024820152682337b93134b23232b760b91b604482015290519081900360640190fd5b60095442116112c2576040805162461bcd60e51b815260206004820152601c60248201527f576974686472617720686173206e6f7420737461727465642079657400000000604482015290519081900360640190fd5b600a54421061130f576040805162461bcd60e51b81526020600482015260146024820152732bb4ba34323930bb903832b934b7b21037bb32b960611b604482015290519081900360640190fd5b805160009081905b8082101561140b57600084838151811061132d57fe5b6020908102919091018101516001600160a01b0381166000908152600390925260409091205490915060ff166113dc5760008061136983610f51565b90925090508180156113785750805b6113b8576040805162461bcd60e51b815260206004820152600c60248201526b2737ba1030903bb4b73732b960a11b604482015290519081900360640190fd5b50506001600160a01b03811660009081526001602052604090206003015493909301925b6001600160a01b03166000908152600360205260409020805460ff191660019081179091559190910190611317565b600d54604080516001600160a01b0390921682526020820185905280517fddc398b321237a8d40ac914388309c2f52a08c134e4dc4ce61e32f57cb7d80f19281900390910190a1600d546040516001600160a01b039091169084156108fc029085906000818181858888f1935050505015801561148c573d6000803e3d6000fd5b5050505050565b6007544210156114e3576040805162461bcd60e51b815260206004820152601660248201527514995d99585b081b9bdd081cdd185c9d1959081e595d60521b604482015290519081900360640190fd5b600854421061152c576040805162461bcd60e51b815260206004820152601060248201526f14995d99585b081a185cc8195b99195960821b604482015290519081900360640190fd5b60006115388383610b6a565b3360009081526001602052604090208054919250908214611596576040805162461bcd60e51b81526020600482015260136024820152720d0c2e6d0cae640c8de40dcdee840dac2e8c6d606b1b604482015290519081900360640190fd5b80600101548411156115df576040805162461bcd60e51b815260206004820152600d60248201526c109a59081b9bdd081d985b1a59609a1b604482015290519081900360640190fd5b600c54841015611636576040805162461bcd60e51b815260206004820152601d60248201527f42696420776173206c657373207468616e206d696e696d756d20626964000000604482015290519081900360640190fd5b6002810154610100900460ff161561168c576040805162461bcd60e51b8152602060048201526014602482015273109a5908185b1c9958591e481c995d99585b195960621b604482015290519081900360640190fd5b600381018490556002808201805461ff0019166101001790556000858152602091825260408120805460018181018084559284529390922090910180546001600160a01b031916331790559081141561171557600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018590555b604080513381526020810187905281517f9891bd98143d5a2c5f29be9a7ef8e5d7b723746fb3b75220234af7aa6caeecfe929181900390910190a15050505050565b600c5481565b60008060008061176c85610f51565b909250905080611785575060009250829150610fea9050565b61178d611af1565b506001600160a01b038516600090815260016020818152604092839020835160e081018552815481529281015491830191909152600281015460ff808216151594840194909452610100810484161515606084018190526201000082048516151560808501526301000000909104909316151560a08301526003015460c082015290611823575060009350839250610fea915050565b8215611842578060c00151816020015103600194509450505050610fea565b6020015195600095509350505050565b600080546001600160a01b0316331461189e576040805162461bcd60e51b81526020600482015260096024820152682337b93134b23232b760b91b604482015290519081900360640190fd5b826001600160a01b031663a9059cbb6118b5610f42565b846040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561190557600080fd5b505af1158015611919573d6000803e3d6000fd5b505050506040513d602081101561192f57600080fd5b50519392505050565b60045460009082111561198a576040805162461bcd60e51b8152602060048201526015602482015274706f736974696f6e206e6f7420696e20617272617960581b604482015290519081900360640190fd5b6004828154811061199757fe5b90600052602060002001549050919050565b600260205281600052604060002081815481106119c257fe5b6000918252602090912001546001600160a01b03169150829050565b60009081526002602052604090205490565b6000546001600160a01b03163314611a3b576040805162461bcd60e51b81526020600482015260096024820152682337b93134b23232b760b91b604482015290519081900360640190fd5b6001600160a01b038116611a96576040805162461bcd60e51b815260206004820152601a60248201527f4e6f6e2d7a65726f20616464726573732072657175697265642e000000000000604482015290519081900360640190fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091529056fe43616e6e6f74207374617274207769746864726177616c20696e207468652070617374416d6f756e742073656e74206973206c657373207468616e206d696e696d756d20626964a265627a7a723158208630119af668c80c7da3f992a13c916a69a216a5bd76d7bf18c3d962b4e237c564736f6c634300050b00326461746573206d757374206265206e6f6e207a65726f00000000000000000000"

// DeployAuction deploys a new Ethereum contract, binding an instance of Auction to it.
func DeployAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _startBids *big.Int, _endBids *big.Int, _startReveal *big.Int, _endReveal *big.Int, _minimumBid *big.Int, _wallet common.Address) (common.Address, *types.Transaction, *Auction, error) {
	parsed, err := abi.JSON(strings.NewReader(AuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AuctionBin), backend, _startBids, _endBids, _startReveal, _endReveal, _minimumBid, _wallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// Auction is an auto generated Go binding around an Ethereum contract.
type Auction struct {
	AuctionCaller     // Read-only binding to the contract
	AuctionTransactor // Write-only binding to the contract
	AuctionFilterer   // Log filterer for contract events
}

// AuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionSession struct {
	Contract     *Auction          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionCallerSession struct {
	Contract *AuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionTransactorSession struct {
	Contract     *AuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionRaw struct {
	Contract *Auction // Generic contract binding to access the raw methods on
}

// AuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionCallerRaw struct {
	Contract *AuctionCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionTransactorRaw struct {
	Contract *AuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuction creates a new instance of Auction, bound to a specific deployed contract.
func NewAuction(address common.Address, backend bind.ContractBackend) (*Auction, error) {
	contract, err := bindAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// NewAuctionCaller creates a new read-only instance of Auction, bound to a specific deployed contract.
func NewAuctionCaller(address common.Address, caller bind.ContractCaller) (*AuctionCaller, error) {
	contract, err := bindAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionCaller{contract: contract}, nil
}

// NewAuctionTransactor creates a new write-only instance of Auction, bound to a specific deployed contract.
func NewAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionTransactor, error) {
	contract, err := bindAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionTransactor{contract: contract}, nil
}

// NewAuctionFilterer creates a new log filterer instance of Auction, bound to a specific deployed contract.
func NewAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionFilterer, error) {
	contract, err := bindAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionFilterer{contract: contract}, nil
}

// bindAuction binds a generic wrapper to an already deployed contract.
func bindAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.AuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transact(opts, method, params...)
}

// CalculateHash is a free data retrieval call binding the contract method 0x500a2b59.
//
// Solidity: function calculateHash(uint256 _bid, bytes _randString) constant returns(bytes32)
func (_Auction *AuctionCaller) CalculateHash(opts *bind.CallOpts, _bid *big.Int, _randString []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Auction.contract.Call(opts, out, "calculateHash", _bid, _randString)
	return *ret0, err
}

// CalculateHash is a free data retrieval call binding the contract method 0x500a2b59.
//
// Solidity: function calculateHash(uint256 _bid, bytes _randString) constant returns(bytes32)
func (_Auction *AuctionSession) CalculateHash(_bid *big.Int, _randString []byte) ([32]byte, error) {
	return _Auction.Contract.CalculateHash(&_Auction.CallOpts, _bid, _randString)
}

// CalculateHash is a free data retrieval call binding the contract method 0x500a2b59.
//
// Solidity: function calculateHash(uint256 _bid, bytes _randString) constant returns(bytes32)
func (_Auction *AuctionCallerSession) CalculateHash(_bid *big.Int, _randString []byte) ([32]byte, error) {
	return _Auction.Contract.CalculateHash(&_Auction.CallOpts, _bid, _randString)
}

// InBidding is a free data retrieval call binding the contract method 0x9eff376d.
//
// Solidity: function inBidding() constant returns(bool)
func (_Auction *AuctionCaller) InBidding(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Auction.contract.Call(opts, out, "inBidding")
	return *ret0, err
}

// InBidding is a free data retrieval call binding the contract method 0x9eff376d.
//
// Solidity: function inBidding() constant returns(bool)
func (_Auction *AuctionSession) InBidding() (bool, error) {
	return _Auction.Contract.InBidding(&_Auction.CallOpts)
}

// InBidding is a free data retrieval call binding the contract method 0x9eff376d.
//
// Solidity: function inBidding() constant returns(bool)
func (_Auction *AuctionCallerSession) InBidding() (bool, error) {
	return _Auction.Contract.InBidding(&_Auction.CallOpts)
}

// InReveal is a free data retrieval call binding the contract method 0x17108562.
//
// Solidity: function inReveal() constant returns(bool)
func (_Auction *AuctionCaller) InReveal(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Auction.contract.Call(opts, out, "inReveal")
	return *ret0, err
}

// InReveal is a free data retrieval call binding the contract method 0x17108562.
//
// Solidity: function inReveal() constant returns(bool)
func (_Auction *AuctionSession) InReveal() (bool, error) {
	return _Auction.Contract.InReveal(&_Auction.CallOpts)
}

// InReveal is a free data retrieval call binding the contract method 0x17108562.
//
// Solidity: function inReveal() constant returns(bool)
func (_Auction *AuctionCallerSession) InReveal() (bool, error) {
	return _Auction.Contract.InReveal(&_Auction.CallOpts)
}

// InWithdraw is a free data retrieval call binding the contract method 0xa0ace77a.
//
// Solidity: function inWithdraw() constant returns(bool)
func (_Auction *AuctionCaller) InWithdraw(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Auction.contract.Call(opts, out, "inWithdraw")
	return *ret0, err
}

// InWithdraw is a free data retrieval call binding the contract method 0xa0ace77a.
//
// Solidity: function inWithdraw() constant returns(bool)
func (_Auction *AuctionSession) InWithdraw() (bool, error) {
	return _Auction.Contract.InWithdraw(&_Auction.CallOpts)
}

// InWithdraw is a free data retrieval call binding the contract method 0xa0ace77a.
//
// Solidity: function inWithdraw() constant returns(bool)
func (_Auction *AuctionCallerSession) InWithdraw() (bool, error) {
	return _Auction.Contract.InWithdraw(&_Auction.CallOpts)
}

// Information is a free data retrieval call binding the contract method 0xa30ab2d0.
//
// Solidity: function information(address ) constant returns(bytes32 hashValue, uint256 value, bool haveBid, bool revealed, bool refunded, bool flagged, uint256 bid)
func (_Auction *AuctionCaller) Information(opts *bind.CallOpts, arg0 common.Address) (struct {
	HashValue [32]byte
	Value     *big.Int
	HaveBid   bool
	Revealed  bool
	Refunded  bool
	Flagged   bool
	Bid       *big.Int
}, error) {
	ret := new(struct {
		HashValue [32]byte
		Value     *big.Int
		HaveBid   bool
		Revealed  bool
		Refunded  bool
		Flagged   bool
		Bid       *big.Int
	})
	out := ret
	err := _Auction.contract.Call(opts, out, "information", arg0)
	return *ret, err
}

// Information is a free data retrieval call binding the contract method 0xa30ab2d0.
//
// Solidity: function information(address ) constant returns(bytes32 hashValue, uint256 value, bool haveBid, bool revealed, bool refunded, bool flagged, uint256 bid)
func (_Auction *AuctionSession) Information(arg0 common.Address) (struct {
	HashValue [32]byte
	Value     *big.Int
	HaveBid   bool
	Revealed  bool
	Refunded  bool
	Flagged   bool
	Bid       *big.Int
}, error) {
	return _Auction.Contract.Information(&_Auction.CallOpts, arg0)
}

// Information is a free data retrieval call binding the contract method 0xa30ab2d0.
//
// Solidity: function information(address ) constant returns(bytes32 hashValue, uint256 value, bool haveBid, bool revealed, bool refunded, bool flagged, uint256 bid)
func (_Auction *AuctionCallerSession) Information(arg0 common.Address) (struct {
	HashValue [32]byte
	Value     *big.Int
	HaveBid   bool
	Revealed  bool
	Refunded  bool
	Flagged   bool
	Bid       *big.Int
}, error) {
	return _Auction.Contract.Information(&_Auction.CallOpts, arg0)
}

// IsWinner is a free data retrieval call binding the contract method 0x9d9ca28d.
//
// Solidity: function isWinner(address check) constant returns(bool winner, bool inPeriod)
func (_Auction *AuctionCaller) IsWinner(opts *bind.CallOpts, check common.Address) (struct {
	Winner   bool
	InPeriod bool
}, error) {
	ret := new(struct {
		Winner   bool
		InPeriod bool
	})
	out := ret
	err := _Auction.contract.Call(opts, out, "isWinner", check)
	return *ret, err
}

// IsWinner is a free data retrieval call binding the contract method 0x9d9ca28d.
//
// Solidity: function isWinner(address check) constant returns(bool winner, bool inPeriod)
func (_Auction *AuctionSession) IsWinner(check common.Address) (struct {
	Winner   bool
	InPeriod bool
}, error) {
	return _Auction.Contract.IsWinner(&_Auction.CallOpts, check)
}

// IsWinner is a free data retrieval call binding the contract method 0x9d9ca28d.
//
// Solidity: function isWinner(address check) constant returns(bool winner, bool inPeriod)
func (_Auction *AuctionCallerSession) IsWinner(check common.Address) (struct {
	Winner   bool
	InPeriod bool
}, error) {
	return _Auction.Contract.IsWinner(&_Auction.CallOpts, check)
}

// MinimumBid is a free data retrieval call binding the contract method 0xd3a86386.
//
// Solidity: function minimumBid() constant returns(uint256)
func (_Auction *AuctionCaller) MinimumBid(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Auction.contract.Call(opts, out, "minimumBid")
	return *ret0, err
}

// MinimumBid is a free data retrieval call binding the contract method 0xd3a86386.
//
// Solidity: function minimumBid() constant returns(uint256)
func (_Auction *AuctionSession) MinimumBid() (*big.Int, error) {
	return _Auction.Contract.MinimumBid(&_Auction.CallOpts)
}

// MinimumBid is a free data retrieval call binding the contract method 0xd3a86386.
//
// Solidity: function minimumBid() constant returns(uint256)
func (_Auction *AuctionCallerSession) MinimumBid() (*big.Int, error) {
	return _Auction.Contract.MinimumBid(&_Auction.CallOpts)
}

// NumberOfRevealedValues is a free data retrieval call binding the contract method 0x41361284.
//
// Solidity: function numberOfRevealedValues() constant returns(uint256)
func (_Auction *AuctionCaller) NumberOfRevealedValues(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Auction.contract.Call(opts, out, "numberOfRevealedValues")
	return *ret0, err
}

// NumberOfRevealedValues is a free data retrieval call binding the contract method 0x41361284.
//
// Solidity: function numberOfRevealedValues() constant returns(uint256)
func (_Auction *AuctionSession) NumberOfRevealedValues() (*big.Int, error) {
	return _Auction.Contract.NumberOfRevealedValues(&_Auction.CallOpts)
}

// NumberOfRevealedValues is a free data retrieval call binding the contract method 0x41361284.
//
// Solidity: function numberOfRevealedValues() constant returns(uint256)
func (_Auction *AuctionCallerSession) NumberOfRevealedValues() (*big.Int, error) {
	return _Auction.Contract.NumberOfRevealedValues(&_Auction.CallOpts)
}

// NumberOfRevealsForValue is a free data retrieval call binding the contract method 0xed295415.
//
// Solidity: function numberOfRevealsForValue(uint256 value) constant returns(uint256)
func (_Auction *AuctionCaller) NumberOfRevealsForValue(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Auction.contract.Call(opts, out, "numberOfRevealsForValue", value)
	return *ret0, err
}

// NumberOfRevealsForValue is a free data retrieval call binding the contract method 0xed295415.
//
// Solidity: function numberOfRevealsForValue(uint256 value) constant returns(uint256)
func (_Auction *AuctionSession) NumberOfRevealsForValue(value *big.Int) (*big.Int, error) {
	return _Auction.Contract.NumberOfRevealsForValue(&_Auction.CallOpts, value)
}

// NumberOfRevealsForValue is a free data retrieval call binding the contract method 0xed295415.
//
// Solidity: function numberOfRevealsForValue(uint256 value) constant returns(uint256)
func (_Auction *AuctionCallerSession) NumberOfRevealsForValue(value *big.Int) (*big.Int, error) {
	return _Auction.Contract.NumberOfRevealsForValue(&_Auction.CallOpts, value)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Auction *AuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Auction.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Auction *AuctionSession) Owner() (common.Address, error) {
	return _Auction.Contract.Owner(&_Auction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Auction *AuctionCallerSession) Owner() (common.Address, error) {
	return _Auction.Contract.Owner(&_Auction.CallOpts)
}

// RevealedValue is a free data retrieval call binding the contract method 0xe44d7a69.
//
// Solidity: function revealedValue(uint256 position) constant returns(uint256)
func (_Auction *AuctionCaller) RevealedValue(opts *bind.CallOpts, position *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Auction.contract.Call(opts, out, "revealedValue", position)
	return *ret0, err
}

// RevealedValue is a free data retrieval call binding the contract method 0xe44d7a69.
//
// Solidity: function revealedValue(uint256 position) constant returns(uint256)
func (_Auction *AuctionSession) RevealedValue(position *big.Int) (*big.Int, error) {
	return _Auction.Contract.RevealedValue(&_Auction.CallOpts, position)
}

// RevealedValue is a free data retrieval call binding the contract method 0xe44d7a69.
//
// Solidity: function revealedValue(uint256 position) constant returns(uint256)
func (_Auction *AuctionCallerSession) RevealedValue(position *big.Int) (*big.Int, error) {
	return _Auction.Contract.RevealedValue(&_Auction.CallOpts, position)
}

// Reveals is a free data retrieval call binding the contract method 0xe69650a6.
//
// Solidity: function reveals(uint256 , uint256 ) constant returns(address)
func (_Auction *AuctionCaller) Reveals(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Auction.contract.Call(opts, out, "reveals", arg0, arg1)
	return *ret0, err
}

// Reveals is a free data retrieval call binding the contract method 0xe69650a6.
//
// Solidity: function reveals(uint256 , uint256 ) constant returns(address)
func (_Auction *AuctionSession) Reveals(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Auction.Contract.Reveals(&_Auction.CallOpts, arg0, arg1)
}

// Reveals is a free data retrieval call binding the contract method 0xe69650a6.
//
// Solidity: function reveals(uint256 , uint256 ) constant returns(address)
func (_Auction *AuctionCallerSession) Reveals(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _Auction.Contract.Reveals(&_Auction.CallOpts, arg0, arg1)
}

// WithdrawalAmount is a free data retrieval call binding the contract method 0xdaa91f2e.
//
// Solidity: function withdrawalAmount(address check) constant returns(uint256, bool)
func (_Auction *AuctionCaller) WithdrawalAmount(opts *bind.CallOpts, check common.Address) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Auction.contract.Call(opts, out, "withdrawalAmount", check)
	return *ret0, *ret1, err
}

// WithdrawalAmount is a free data retrieval call binding the contract method 0xdaa91f2e.
//
// Solidity: function withdrawalAmount(address check) constant returns(uint256, bool)
func (_Auction *AuctionSession) WithdrawalAmount(check common.Address) (*big.Int, bool, error) {
	return _Auction.Contract.WithdrawalAmount(&_Auction.CallOpts, check)
}

// WithdrawalAmount is a free data retrieval call binding the contract method 0xdaa91f2e.
//
// Solidity: function withdrawalAmount(address check) constant returns(uint256, bool)
func (_Auction *AuctionCallerSession) WithdrawalAmount(check common.Address) (*big.Int, bool, error) {
	return _Auction.Contract.WithdrawalAmount(&_Auction.CallOpts, check)
}

// BiddingTime is a paid mutator transaction binding the contract method 0xa056e620.
//
// Solidity: function biddingTime(bytes32 _hash) returns()
func (_Auction *AuctionTransactor) BiddingTime(opts *bind.TransactOpts, _hash [32]byte) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "biddingTime", _hash)
}

// BiddingTime is a paid mutator transaction binding the contract method 0xa056e620.
//
// Solidity: function biddingTime(bytes32 _hash) returns()
func (_Auction *AuctionSession) BiddingTime(_hash [32]byte) (*types.Transaction, error) {
	return _Auction.Contract.BiddingTime(&_Auction.TransactOpts, _hash)
}

// BiddingTime is a paid mutator transaction binding the contract method 0xa056e620.
//
// Solidity: function biddingTime(bytes32 _hash) returns()
func (_Auction *AuctionTransactorSession) BiddingTime(_hash [32]byte) (*types.Transaction, error) {
	return _Auction.Contract.BiddingTime(&_Auction.TransactOpts, _hash)
}

// EarlyWithdrawal is a paid mutator transaction binding the contract method 0xb856ed86.
//
// Solidity: function earlyWithdrawal(address[] winners) returns()
func (_Auction *AuctionTransactor) EarlyWithdrawal(opts *bind.TransactOpts, winners []common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "earlyWithdrawal", winners)
}

// EarlyWithdrawal is a paid mutator transaction binding the contract method 0xb856ed86.
//
// Solidity: function earlyWithdrawal(address[] winners) returns()
func (_Auction *AuctionSession) EarlyWithdrawal(winners []common.Address) (*types.Transaction, error) {
	return _Auction.Contract.EarlyWithdrawal(&_Auction.TransactOpts, winners)
}

// EarlyWithdrawal is a paid mutator transaction binding the contract method 0xb856ed86.
//
// Solidity: function earlyWithdrawal(address[] winners) returns()
func (_Auction *AuctionTransactorSession) EarlyWithdrawal(winners []common.Address) (*types.Transaction, error) {
	return _Auction.Contract.EarlyWithdrawal(&_Auction.TransactOpts, winners)
}

// Reveal is a paid mutator transaction binding the contract method 0xce805642.
//
// Solidity: function reveal(uint256 _bid, bytes randString) returns()
func (_Auction *AuctionTransactor) Reveal(opts *bind.TransactOpts, _bid *big.Int, randString []byte) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "reveal", _bid, randString)
}

// Reveal is a paid mutator transaction binding the contract method 0xce805642.
//
// Solidity: function reveal(uint256 _bid, bytes randString) returns()
func (_Auction *AuctionSession) Reveal(_bid *big.Int, randString []byte) (*types.Transaction, error) {
	return _Auction.Contract.Reveal(&_Auction.TransactOpts, _bid, randString)
}

// Reveal is a paid mutator transaction binding the contract method 0xce805642.
//
// Solidity: function reveal(uint256 _bid, bytes randString) returns()
func (_Auction *AuctionTransactorSession) Reveal(_bid *big.Int, randString []byte) (*types.Transaction, error) {
	return _Auction.Contract.Reveal(&_Auction.TransactOpts, _bid, randString)
}

// SetWinningAddresses is a paid mutator transaction binding the contract method 0x51c65bdf.
//
// Solidity: function setWinningAddresses(address[] theAddresses, bool flagged) returns()
func (_Auction *AuctionTransactor) SetWinningAddresses(opts *bind.TransactOpts, theAddresses []common.Address, flagged bool) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "setWinningAddresses", theAddresses, flagged)
}

// SetWinningAddresses is a paid mutator transaction binding the contract method 0x51c65bdf.
//
// Solidity: function setWinningAddresses(address[] theAddresses, bool flagged) returns()
func (_Auction *AuctionSession) SetWinningAddresses(theAddresses []common.Address, flagged bool) (*types.Transaction, error) {
	return _Auction.Contract.SetWinningAddresses(&_Auction.TransactOpts, theAddresses, flagged)
}

// SetWinningAddresses is a paid mutator transaction binding the contract method 0x51c65bdf.
//
// Solidity: function setWinningAddresses(address[] theAddresses, bool flagged) returns()
func (_Auction *AuctionTransactorSession) SetWinningAddresses(theAddresses []common.Address, flagged bool) (*types.Transaction, error) {
	return _Auction.Contract.SetWinningAddresses(&_Auction.TransactOpts, theAddresses, flagged)
}

// SetWinningAmount is a paid mutator transaction binding the contract method 0x7eb4d7d3.
//
// Solidity: function setWinningAmount(uint256 _winningAmount) returns()
func (_Auction *AuctionTransactor) SetWinningAmount(opts *bind.TransactOpts, _winningAmount *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "setWinningAmount", _winningAmount)
}

// SetWinningAmount is a paid mutator transaction binding the contract method 0x7eb4d7d3.
//
// Solidity: function setWinningAmount(uint256 _winningAmount) returns()
func (_Auction *AuctionSession) SetWinningAmount(_winningAmount *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.SetWinningAmount(&_Auction.TransactOpts, _winningAmount)
}

// SetWinningAmount is a paid mutator transaction binding the contract method 0x7eb4d7d3.
//
// Solidity: function setWinningAmount(uint256 _winningAmount) returns()
func (_Auction *AuctionTransactorSession) SetWinningAmount(_winningAmount *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.SetWinningAmount(&_Auction.TransactOpts, _winningAmount)
}

// StartWithdrawal is a paid mutator transaction binding the contract method 0x70431d29.
//
// Solidity: function startWithdrawal(uint256 _startWithdraw, uint256 _endWithdraw) returns()
func (_Auction *AuctionTransactor) StartWithdrawal(opts *bind.TransactOpts, _startWithdraw *big.Int, _endWithdraw *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "startWithdrawal", _startWithdraw, _endWithdraw)
}

// StartWithdrawal is a paid mutator transaction binding the contract method 0x70431d29.
//
// Solidity: function startWithdrawal(uint256 _startWithdraw, uint256 _endWithdraw) returns()
func (_Auction *AuctionSession) StartWithdrawal(_startWithdraw *big.Int, _endWithdraw *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.StartWithdrawal(&_Auction.TransactOpts, _startWithdraw, _endWithdraw)
}

// StartWithdrawal is a paid mutator transaction binding the contract method 0x70431d29.
//
// Solidity: function startWithdrawal(uint256 _startWithdraw, uint256 _endWithdraw) returns()
func (_Auction *AuctionTransactorSession) StartWithdrawal(_startWithdraw *big.Int, _endWithdraw *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.StartWithdrawal(&_Auction.TransactOpts, _startWithdraw, _endWithdraw)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, uint256 amount) returns(bool)
func (_Auction *AuctionTransactor) TransferAnyERC20Token(opts *bind.TransactOpts, tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "transferAnyERC20Token", tokenAddress, amount)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, uint256 amount) returns(bool)
func (_Auction *AuctionSession) TransferAnyERC20Token(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.TransferAnyERC20Token(&_Auction.TransactOpts, tokenAddress, amount)
}

// TransferAnyERC20Token is a paid mutator transaction binding the contract method 0xdc39d06d.
//
// Solidity: function transferAnyERC20Token(address tokenAddress, uint256 amount) returns(bool)
func (_Auction *AuctionTransactorSession) TransferAnyERC20Token(tokenAddress common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.TransferAnyERC20Token(&_Auction.TransactOpts, tokenAddress, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Auction *AuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Auction *AuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Auction.Contract.TransferOwnership(&_Auction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Auction *AuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Auction.Contract.TransferOwnership(&_Auction.TransactOpts, newOwner)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x476343ee.
//
// Solidity: function withdrawFees() returns()
func (_Auction *AuctionTransactor) WithdrawFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "withdrawFees")
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x476343ee.
//
// Solidity: function withdrawFees() returns()
func (_Auction *AuctionSession) WithdrawFees() (*types.Transaction, error) {
	return _Auction.Contract.WithdrawFees(&_Auction.TransactOpts)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x476343ee.
//
// Solidity: function withdrawFees() returns()
func (_Auction *AuctionTransactorSession) WithdrawFees() (*types.Transaction, error) {
	return _Auction.Contract.WithdrawFees(&_Auction.TransactOpts)
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0x110f8874.
//
// Solidity: function withdrawRefund() returns()
func (_Auction *AuctionTransactor) WithdrawRefund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "withdrawRefund")
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0x110f8874.
//
// Solidity: function withdrawRefund() returns()
func (_Auction *AuctionSession) WithdrawRefund() (*types.Transaction, error) {
	return _Auction.Contract.WithdrawRefund(&_Auction.TransactOpts)
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0x110f8874.
//
// Solidity: function withdrawRefund() returns()
func (_Auction *AuctionTransactorSession) WithdrawRefund() (*types.Transaction, error) {
	return _Auction.Contract.WithdrawRefund(&_Auction.TransactOpts)
}

// AuctionBalanceWithdrawnIterator is returned from FilterBalanceWithdrawn and is used to iterate over the raw logs and unpacked data for BalanceWithdrawn events raised by the Auction contract.
type AuctionBalanceWithdrawnIterator struct {
	Event *AuctionBalanceWithdrawn // Event containing the contract specifics and raw log

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
func (it *AuctionBalanceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBalanceWithdrawn)
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
		it.Event = new(AuctionBalanceWithdrawn)
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
func (it *AuctionBalanceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBalanceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBalanceWithdrawn represents a BalanceWithdrawn event raised by the Auction contract.
type AuctionBalanceWithdrawn struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBalanceWithdrawn is a free log retrieval operation binding the contract event 0xddc398b321237a8d40ac914388309c2f52a08c134e4dc4ce61e32f57cb7d80f1.
//
// Solidity: event BalanceWithdrawn(address recipient, uint256 amount)
func (_Auction *AuctionFilterer) FilterBalanceWithdrawn(opts *bind.FilterOpts) (*AuctionBalanceWithdrawnIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "BalanceWithdrawn")
	if err != nil {
		return nil, err
	}
	return &AuctionBalanceWithdrawnIterator{contract: _Auction.contract, event: "BalanceWithdrawn", logs: logs, sub: sub}, nil
}

// WatchBalanceWithdrawn is a free log subscription operation binding the contract event 0xddc398b321237a8d40ac914388309c2f52a08c134e4dc4ce61e32f57cb7d80f1.
//
// Solidity: event BalanceWithdrawn(address recipient, uint256 amount)
func (_Auction *AuctionFilterer) WatchBalanceWithdrawn(opts *bind.WatchOpts, sink chan<- *AuctionBalanceWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "BalanceWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBalanceWithdrawn)
				if err := _Auction.contract.UnpackLog(event, "BalanceWithdrawn", log); err != nil {
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

// ParseBalanceWithdrawn is a log parse operation binding the contract event 0xddc398b321237a8d40ac914388309c2f52a08c134e4dc4ce61e32f57cb7d80f1.
//
// Solidity: event BalanceWithdrawn(address recipient, uint256 amount)
func (_Auction *AuctionFilterer) ParseBalanceWithdrawn(log types.Log) (*AuctionBalanceWithdrawn, error) {
	event := new(AuctionBalanceWithdrawn)
	if err := _Auction.contract.UnpackLog(event, "BalanceWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionBidRevealedIterator is returned from FilterBidRevealed and is used to iterate over the raw logs and unpacked data for BidRevealed events raised by the Auction contract.
type AuctionBidRevealedIterator struct {
	Event *AuctionBidRevealed // Event containing the contract specifics and raw log

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
func (it *AuctionBidRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBidRevealed)
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
		it.Event = new(AuctionBidRevealed)
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
func (it *AuctionBidRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBidRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBidRevealed represents a BidRevealed event raised by the Auction contract.
type AuctionBidRevealed struct {
	Bidder common.Address
	Bid    *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBidRevealed is a free log retrieval operation binding the contract event 0x9891bd98143d5a2c5f29be9a7ef8e5d7b723746fb3b75220234af7aa6caeecfe.
//
// Solidity: event BidRevealed(address bidder, uint256 bid)
func (_Auction *AuctionFilterer) FilterBidRevealed(opts *bind.FilterOpts) (*AuctionBidRevealedIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "BidRevealed")
	if err != nil {
		return nil, err
	}
	return &AuctionBidRevealedIterator{contract: _Auction.contract, event: "BidRevealed", logs: logs, sub: sub}, nil
}

// WatchBidRevealed is a free log subscription operation binding the contract event 0x9891bd98143d5a2c5f29be9a7ef8e5d7b723746fb3b75220234af7aa6caeecfe.
//
// Solidity: event BidRevealed(address bidder, uint256 bid)
func (_Auction *AuctionFilterer) WatchBidRevealed(opts *bind.WatchOpts, sink chan<- *AuctionBidRevealed) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "BidRevealed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBidRevealed)
				if err := _Auction.contract.UnpackLog(event, "BidRevealed", log); err != nil {
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

// ParseBidRevealed is a log parse operation binding the contract event 0x9891bd98143d5a2c5f29be9a7ef8e5d7b723746fb3b75220234af7aa6caeecfe.
//
// Solidity: event BidRevealed(address bidder, uint256 bid)
func (_Auction *AuctionFilterer) ParseBidRevealed(log types.Log) (*AuctionBidRevealed, error) {
	event := new(AuctionBidRevealed)
	if err := _Auction.contract.UnpackLog(event, "BidRevealed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionBidSubmittedIterator is returned from FilterBidSubmitted and is used to iterate over the raw logs and unpacked data for BidSubmitted events raised by the Auction contract.
type AuctionBidSubmittedIterator struct {
	Event *AuctionBidSubmitted // Event containing the contract specifics and raw log

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
func (it *AuctionBidSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBidSubmitted)
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
		it.Event = new(AuctionBidSubmitted)
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
func (it *AuctionBidSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBidSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBidSubmitted represents a BidSubmitted event raised by the Auction contract.
type AuctionBidSubmitted struct {
	Bidder  common.Address
	Funding *big.Int
	Hash    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBidSubmitted is a free log retrieval operation binding the contract event 0x43dfb6c14d1b44aaeaef2c0487e8423e2843a99f9e09108814b95ce236ed8707.
//
// Solidity: event BidSubmitted(address bidder, uint256 funding, bytes32 hash)
func (_Auction *AuctionFilterer) FilterBidSubmitted(opts *bind.FilterOpts) (*AuctionBidSubmittedIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "BidSubmitted")
	if err != nil {
		return nil, err
	}
	return &AuctionBidSubmittedIterator{contract: _Auction.contract, event: "BidSubmitted", logs: logs, sub: sub}, nil
}

// WatchBidSubmitted is a free log subscription operation binding the contract event 0x43dfb6c14d1b44aaeaef2c0487e8423e2843a99f9e09108814b95ce236ed8707.
//
// Solidity: event BidSubmitted(address bidder, uint256 funding, bytes32 hash)
func (_Auction *AuctionFilterer) WatchBidSubmitted(opts *bind.WatchOpts, sink chan<- *AuctionBidSubmitted) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "BidSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBidSubmitted)
				if err := _Auction.contract.UnpackLog(event, "BidSubmitted", log); err != nil {
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

// ParseBidSubmitted is a log parse operation binding the contract event 0x43dfb6c14d1b44aaeaef2c0487e8423e2843a99f9e09108814b95ce236ed8707.
//
// Solidity: event BidSubmitted(address bidder, uint256 funding, bytes32 hash)
func (_Auction *AuctionFilterer) ParseBidSubmitted(log types.Log) (*AuctionBidSubmitted, error) {
	event := new(AuctionBidSubmitted)
	if err := _Auction.contract.UnpackLog(event, "BidSubmitted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionBiddingPeriodIterator is returned from FilterBiddingPeriod and is used to iterate over the raw logs and unpacked data for BiddingPeriod events raised by the Auction contract.
type AuctionBiddingPeriodIterator struct {
	Event *AuctionBiddingPeriod // Event containing the contract specifics and raw log

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
func (it *AuctionBiddingPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBiddingPeriod)
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
		it.Event = new(AuctionBiddingPeriod)
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
func (it *AuctionBiddingPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBiddingPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBiddingPeriod represents a BiddingPeriod event raised by the Auction contract.
type AuctionBiddingPeriod struct {
	StartBids *big.Int
	EndBids   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBiddingPeriod is a free log retrieval operation binding the contract event 0x892ffcfb665b4ba99a581f2bfa4b8614aaa46bae9247c832b4bd6c1d2badb2e6.
//
// Solidity: event BiddingPeriod(uint256 startBids, uint256 endBids)
func (_Auction *AuctionFilterer) FilterBiddingPeriod(opts *bind.FilterOpts) (*AuctionBiddingPeriodIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "BiddingPeriod")
	if err != nil {
		return nil, err
	}
	return &AuctionBiddingPeriodIterator{contract: _Auction.contract, event: "BiddingPeriod", logs: logs, sub: sub}, nil
}

// WatchBiddingPeriod is a free log subscription operation binding the contract event 0x892ffcfb665b4ba99a581f2bfa4b8614aaa46bae9247c832b4bd6c1d2badb2e6.
//
// Solidity: event BiddingPeriod(uint256 startBids, uint256 endBids)
func (_Auction *AuctionFilterer) WatchBiddingPeriod(opts *bind.WatchOpts, sink chan<- *AuctionBiddingPeriod) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "BiddingPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBiddingPeriod)
				if err := _Auction.contract.UnpackLog(event, "BiddingPeriod", log); err != nil {
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

// ParseBiddingPeriod is a log parse operation binding the contract event 0x892ffcfb665b4ba99a581f2bfa4b8614aaa46bae9247c832b4bd6c1d2badb2e6.
//
// Solidity: event BiddingPeriod(uint256 startBids, uint256 endBids)
func (_Auction *AuctionFilterer) ParseBiddingPeriod(log types.Log) (*AuctionBiddingPeriod, error) {
	event := new(AuctionBiddingPeriod)
	if err := _Auction.contract.UnpackLog(event, "BiddingPeriod", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionMinimumBidIterator is returned from FilterMinimumBid and is used to iterate over the raw logs and unpacked data for MinimumBid events raised by the Auction contract.
type AuctionMinimumBidIterator struct {
	Event *AuctionMinimumBid // Event containing the contract specifics and raw log

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
func (it *AuctionMinimumBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMinimumBid)
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
		it.Event = new(AuctionMinimumBid)
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
func (it *AuctionMinimumBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMinimumBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMinimumBid represents a MinimumBid event raised by the Auction contract.
type AuctionMinimumBid struct {
	MinimumBid *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMinimumBid is a free log retrieval operation binding the contract event 0x855d14e5037e9de69ed821e8030c2c0ccaf36f6ebdb3fdc4689f51aa85a518f0.
//
// Solidity: event MinimumBid(uint256 _minimumBid)
func (_Auction *AuctionFilterer) FilterMinimumBid(opts *bind.FilterOpts) (*AuctionMinimumBidIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "MinimumBid")
	if err != nil {
		return nil, err
	}
	return &AuctionMinimumBidIterator{contract: _Auction.contract, event: "MinimumBid", logs: logs, sub: sub}, nil
}

// WatchMinimumBid is a free log subscription operation binding the contract event 0x855d14e5037e9de69ed821e8030c2c0ccaf36f6ebdb3fdc4689f51aa85a518f0.
//
// Solidity: event MinimumBid(uint256 _minimumBid)
func (_Auction *AuctionFilterer) WatchMinimumBid(opts *bind.WatchOpts, sink chan<- *AuctionMinimumBid) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "MinimumBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMinimumBid)
				if err := _Auction.contract.UnpackLog(event, "MinimumBid", log); err != nil {
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

// ParseMinimumBid is a log parse operation binding the contract event 0x855d14e5037e9de69ed821e8030c2c0ccaf36f6ebdb3fdc4689f51aa85a518f0.
//
// Solidity: event MinimumBid(uint256 _minimumBid)
func (_Auction *AuctionFilterer) ParseMinimumBid(log types.Log) (*AuctionMinimumBid, error) {
	event := new(AuctionMinimumBid)
	if err := _Auction.contract.UnpackLog(event, "MinimumBid", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionNothingToRefundIterator is returned from FilterNothingToRefund and is used to iterate over the raw logs and unpacked data for NothingToRefund events raised by the Auction contract.
type AuctionNothingToRefundIterator struct {
	Event *AuctionNothingToRefund // Event containing the contract specifics and raw log

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
func (it *AuctionNothingToRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionNothingToRefund)
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
		it.Event = new(AuctionNothingToRefund)
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
func (it *AuctionNothingToRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionNothingToRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionNothingToRefund represents a NothingToRefund event raised by the Auction contract.
type AuctionNothingToRefund struct {
	Bidder common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNothingToRefund is a free log retrieval operation binding the contract event 0xbb386939e3dac174d852c26db4592a53fb01fde4794fa3124c31416a0c9b9ec0.
//
// Solidity: event NothingToRefund(address bidder)
func (_Auction *AuctionFilterer) FilterNothingToRefund(opts *bind.FilterOpts) (*AuctionNothingToRefundIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "NothingToRefund")
	if err != nil {
		return nil, err
	}
	return &AuctionNothingToRefundIterator{contract: _Auction.contract, event: "NothingToRefund", logs: logs, sub: sub}, nil
}

// WatchNothingToRefund is a free log subscription operation binding the contract event 0xbb386939e3dac174d852c26db4592a53fb01fde4794fa3124c31416a0c9b9ec0.
//
// Solidity: event NothingToRefund(address bidder)
func (_Auction *AuctionFilterer) WatchNothingToRefund(opts *bind.WatchOpts, sink chan<- *AuctionNothingToRefund) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "NothingToRefund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionNothingToRefund)
				if err := _Auction.contract.UnpackLog(event, "NothingToRefund", log); err != nil {
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

// ParseNothingToRefund is a log parse operation binding the contract event 0xbb386939e3dac174d852c26db4592a53fb01fde4794fa3124c31416a0c9b9ec0.
//
// Solidity: event NothingToRefund(address bidder)
func (_Auction *AuctionFilterer) ParseNothingToRefund(log types.Log) (*AuctionNothingToRefund, error) {
	event := new(AuctionNothingToRefund)
	if err := _Auction.contract.UnpackLog(event, "NothingToRefund", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Auction contract.
type AuctionOwnershipTransferredIterator struct {
	Event *AuctionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuctionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionOwnershipTransferred)
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
		it.Event = new(AuctionOwnershipTransferred)
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
func (it *AuctionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionOwnershipTransferred represents a OwnershipTransferred event raised by the Auction contract.
type AuctionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Auction *AuctionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuctionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionOwnershipTransferredIterator{contract: _Auction.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Auction *AuctionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuctionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionOwnershipTransferred)
				if err := _Auction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Auction *AuctionFilterer) ParseOwnershipTransferred(log types.Log) (*AuctionOwnershipTransferred, error) {
	event := new(AuctionOwnershipTransferred)
	if err := _Auction.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the Auction contract.
type AuctionRefundIterator struct {
	Event *AuctionRefund // Event containing the contract specifics and raw log

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
func (it *AuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionRefund)
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
		it.Event = new(AuctionRefund)
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
func (it *AuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionRefund represents a Refund event raised by the Auction contract.
type AuctionRefund struct {
	Bidder common.Address
	Refund *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0xbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d.
//
// Solidity: event Refund(address bidder, uint256 refund)
func (_Auction *AuctionFilterer) FilterRefund(opts *bind.FilterOpts) (*AuctionRefundIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return &AuctionRefundIterator{contract: _Auction.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0xbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d.
//
// Solidity: event Refund(address bidder, uint256 refund)
func (_Auction *AuctionFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *AuctionRefund) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionRefund)
				if err := _Auction.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0xbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d.
//
// Solidity: event Refund(address bidder, uint256 refund)
func (_Auction *AuctionFilterer) ParseRefund(log types.Log) (*AuctionRefund, error) {
	event := new(AuctionRefund)
	if err := _Auction.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionRefundChangeIterator is returned from FilterRefundChange and is used to iterate over the raw logs and unpacked data for RefundChange events raised by the Auction contract.
type AuctionRefundChangeIterator struct {
	Event *AuctionRefundChange // Event containing the contract specifics and raw log

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
func (it *AuctionRefundChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionRefundChange)
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
		it.Event = new(AuctionRefundChange)
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
func (it *AuctionRefundChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionRefundChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionRefundChange represents a RefundChange event raised by the Auction contract.
type AuctionRefundChange struct {
	Bidder common.Address
	Change *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefundChange is a free log retrieval operation binding the contract event 0x7417b321c40d8a0c375656909a6cb0a232297cc6addda0f3a386b9f761669d90.
//
// Solidity: event RefundChange(address bidder, uint256 change)
func (_Auction *AuctionFilterer) FilterRefundChange(opts *bind.FilterOpts) (*AuctionRefundChangeIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "RefundChange")
	if err != nil {
		return nil, err
	}
	return &AuctionRefundChangeIterator{contract: _Auction.contract, event: "RefundChange", logs: logs, sub: sub}, nil
}

// WatchRefundChange is a free log subscription operation binding the contract event 0x7417b321c40d8a0c375656909a6cb0a232297cc6addda0f3a386b9f761669d90.
//
// Solidity: event RefundChange(address bidder, uint256 change)
func (_Auction *AuctionFilterer) WatchRefundChange(opts *bind.WatchOpts, sink chan<- *AuctionRefundChange) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "RefundChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionRefundChange)
				if err := _Auction.contract.UnpackLog(event, "RefundChange", log); err != nil {
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

// ParseRefundChange is a log parse operation binding the contract event 0x7417b321c40d8a0c375656909a6cb0a232297cc6addda0f3a386b9f761669d90.
//
// Solidity: event RefundChange(address bidder, uint256 change)
func (_Auction *AuctionFilterer) ParseRefundChange(log types.Log) (*AuctionRefundChange, error) {
	event := new(AuctionRefundChange)
	if err := _Auction.contract.UnpackLog(event, "RefundChange", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionRevealPeriodIterator is returned from FilterRevealPeriod and is used to iterate over the raw logs and unpacked data for RevealPeriod events raised by the Auction contract.
type AuctionRevealPeriodIterator struct {
	Event *AuctionRevealPeriod // Event containing the contract specifics and raw log

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
func (it *AuctionRevealPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionRevealPeriod)
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
		it.Event = new(AuctionRevealPeriod)
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
func (it *AuctionRevealPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionRevealPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionRevealPeriod represents a RevealPeriod event raised by the Auction contract.
type AuctionRevealPeriod struct {
	StartReveal *big.Int
	EndReveal   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRevealPeriod is a free log retrieval operation binding the contract event 0xd23f2f662aba8ca991baba62f1219f0780865580d293d0c7bc89079757639333.
//
// Solidity: event RevealPeriod(uint256 startReveal, uint256 endReveal)
func (_Auction *AuctionFilterer) FilterRevealPeriod(opts *bind.FilterOpts) (*AuctionRevealPeriodIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "RevealPeriod")
	if err != nil {
		return nil, err
	}
	return &AuctionRevealPeriodIterator{contract: _Auction.contract, event: "RevealPeriod", logs: logs, sub: sub}, nil
}

// WatchRevealPeriod is a free log subscription operation binding the contract event 0xd23f2f662aba8ca991baba62f1219f0780865580d293d0c7bc89079757639333.
//
// Solidity: event RevealPeriod(uint256 startReveal, uint256 endReveal)
func (_Auction *AuctionFilterer) WatchRevealPeriod(opts *bind.WatchOpts, sink chan<- *AuctionRevealPeriod) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "RevealPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionRevealPeriod)
				if err := _Auction.contract.UnpackLog(event, "RevealPeriod", log); err != nil {
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

// ParseRevealPeriod is a log parse operation binding the contract event 0xd23f2f662aba8ca991baba62f1219f0780865580d293d0c7bc89079757639333.
//
// Solidity: event RevealPeriod(uint256 startReveal, uint256 endReveal)
func (_Auction *AuctionFilterer) ParseRevealPeriod(log types.Log) (*AuctionRevealPeriod, error) {
	event := new(AuctionRevealPeriod)
	if err := _Auction.contract.UnpackLog(event, "RevealPeriod", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionWalletIterator is returned from FilterWallet and is used to iterate over the raw logs and unpacked data for Wallet events raised by the Auction contract.
type AuctionWalletIterator struct {
	Event *AuctionWallet // Event containing the contract specifics and raw log

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
func (it *AuctionWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionWallet)
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
		it.Event = new(AuctionWallet)
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
func (it *AuctionWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionWallet represents a Wallet event raised by the Auction contract.
type AuctionWallet struct {
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWallet is a free log retrieval operation binding the contract event 0x299e6b07aa8d613d084182e45e438778c5e03395c860c2f28273c66e6f620f0c.
//
// Solidity: event Wallet(address _wallet)
func (_Auction *AuctionFilterer) FilterWallet(opts *bind.FilterOpts) (*AuctionWalletIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Wallet")
	if err != nil {
		return nil, err
	}
	return &AuctionWalletIterator{contract: _Auction.contract, event: "Wallet", logs: logs, sub: sub}, nil
}

// WatchWallet is a free log subscription operation binding the contract event 0x299e6b07aa8d613d084182e45e438778c5e03395c860c2f28273c66e6f620f0c.
//
// Solidity: event Wallet(address _wallet)
func (_Auction *AuctionFilterer) WatchWallet(opts *bind.WatchOpts, sink chan<- *AuctionWallet) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Wallet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionWallet)
				if err := _Auction.contract.UnpackLog(event, "Wallet", log); err != nil {
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

// ParseWallet is a log parse operation binding the contract event 0x299e6b07aa8d613d084182e45e438778c5e03395c860c2f28273c66e6f620f0c.
//
// Solidity: event Wallet(address _wallet)
func (_Auction *AuctionFilterer) ParseWallet(log types.Log) (*AuctionWallet, error) {
	event := new(AuctionWallet)
	if err := _Auction.contract.UnpackLog(event, "Wallet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionWinnerWithTieIterator is returned from FilterWinnerWithTie and is used to iterate over the raw logs and unpacked data for WinnerWithTie events raised by the Auction contract.
type AuctionWinnerWithTieIterator struct {
	Event *AuctionWinnerWithTie // Event containing the contract specifics and raw log

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
func (it *AuctionWinnerWithTieIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionWinnerWithTie)
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
		it.Event = new(AuctionWinnerWithTie)
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
func (it *AuctionWinnerWithTieIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionWinnerWithTieIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionWinnerWithTie represents a WinnerWithTie event raised by the Auction contract.
type AuctionWinnerWithTie struct {
	TiedWinner common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWinnerWithTie is a free log retrieval operation binding the contract event 0xdeb65b46e26a4ed6abe63bd73c4613e1262a81bafb7bb35b92f5160e8f1724f0.
//
// Solidity: event WinnerWithTie(address tiedWinner)
func (_Auction *AuctionFilterer) FilterWinnerWithTie(opts *bind.FilterOpts) (*AuctionWinnerWithTieIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "WinnerWithTie")
	if err != nil {
		return nil, err
	}
	return &AuctionWinnerWithTieIterator{contract: _Auction.contract, event: "WinnerWithTie", logs: logs, sub: sub}, nil
}

// WatchWinnerWithTie is a free log subscription operation binding the contract event 0xdeb65b46e26a4ed6abe63bd73c4613e1262a81bafb7bb35b92f5160e8f1724f0.
//
// Solidity: event WinnerWithTie(address tiedWinner)
func (_Auction *AuctionFilterer) WatchWinnerWithTie(opts *bind.WatchOpts, sink chan<- *AuctionWinnerWithTie) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "WinnerWithTie")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionWinnerWithTie)
				if err := _Auction.contract.UnpackLog(event, "WinnerWithTie", log); err != nil {
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

// ParseWinnerWithTie is a log parse operation binding the contract event 0xdeb65b46e26a4ed6abe63bd73c4613e1262a81bafb7bb35b92f5160e8f1724f0.
//
// Solidity: event WinnerWithTie(address tiedWinner)
func (_Auction *AuctionFilterer) ParseWinnerWithTie(log types.Log) (*AuctionWinnerWithTie, error) {
	event := new(AuctionWinnerWithTie)
	if err := _Auction.contract.UnpackLog(event, "WinnerWithTie", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionWinningAmountIterator is returned from FilterWinningAmount and is used to iterate over the raw logs and unpacked data for WinningAmount events raised by the Auction contract.
type AuctionWinningAmountIterator struct {
	Event *AuctionWinningAmount // Event containing the contract specifics and raw log

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
func (it *AuctionWinningAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionWinningAmount)
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
		it.Event = new(AuctionWinningAmount)
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
func (it *AuctionWinningAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionWinningAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionWinningAmount represents a WinningAmount event raised by the Auction contract.
type AuctionWinningAmount struct {
	WinningAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWinningAmount is a free log retrieval operation binding the contract event 0x85679d62923fa647b865510f1c0f24c7babbee0a25eca9118b226faf71f42f1c.
//
// Solidity: event WinningAmount(uint256 winningAmount)
func (_Auction *AuctionFilterer) FilterWinningAmount(opts *bind.FilterOpts) (*AuctionWinningAmountIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "WinningAmount")
	if err != nil {
		return nil, err
	}
	return &AuctionWinningAmountIterator{contract: _Auction.contract, event: "WinningAmount", logs: logs, sub: sub}, nil
}

// WatchWinningAmount is a free log subscription operation binding the contract event 0x85679d62923fa647b865510f1c0f24c7babbee0a25eca9118b226faf71f42f1c.
//
// Solidity: event WinningAmount(uint256 winningAmount)
func (_Auction *AuctionFilterer) WatchWinningAmount(opts *bind.WatchOpts, sink chan<- *AuctionWinningAmount) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "WinningAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionWinningAmount)
				if err := _Auction.contract.UnpackLog(event, "WinningAmount", log); err != nil {
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

// ParseWinningAmount is a log parse operation binding the contract event 0x85679d62923fa647b865510f1c0f24c7babbee0a25eca9118b226faf71f42f1c.
//
// Solidity: event WinningAmount(uint256 winningAmount)
func (_Auction *AuctionFilterer) ParseWinningAmount(log types.Log) (*AuctionWinningAmount, error) {
	event := new(AuctionWinningAmount)
	if err := _Auction.contract.UnpackLog(event, "WinningAmount", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AuctionWithdrawPeriodIterator is returned from FilterWithdrawPeriod and is used to iterate over the raw logs and unpacked data for WithdrawPeriod events raised by the Auction contract.
type AuctionWithdrawPeriodIterator struct {
	Event *AuctionWithdrawPeriod // Event containing the contract specifics and raw log

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
func (it *AuctionWithdrawPeriodIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionWithdrawPeriod)
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
		it.Event = new(AuctionWithdrawPeriod)
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
func (it *AuctionWithdrawPeriodIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionWithdrawPeriodIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionWithdrawPeriod represents a WithdrawPeriod event raised by the Auction contract.
type AuctionWithdrawPeriod struct {
	StartWithdraw *big.Int
	EndWithdraw   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawPeriod is a free log retrieval operation binding the contract event 0x6efecf63a32b9b51e8ead9e19a996f69916a2753adab220b8bced7624fb50e34.
//
// Solidity: event WithdrawPeriod(uint256 startWithdraw, uint256 endWithdraw)
func (_Auction *AuctionFilterer) FilterWithdrawPeriod(opts *bind.FilterOpts) (*AuctionWithdrawPeriodIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "WithdrawPeriod")
	if err != nil {
		return nil, err
	}
	return &AuctionWithdrawPeriodIterator{contract: _Auction.contract, event: "WithdrawPeriod", logs: logs, sub: sub}, nil
}

// WatchWithdrawPeriod is a free log subscription operation binding the contract event 0x6efecf63a32b9b51e8ead9e19a996f69916a2753adab220b8bced7624fb50e34.
//
// Solidity: event WithdrawPeriod(uint256 startWithdraw, uint256 endWithdraw)
func (_Auction *AuctionFilterer) WatchWithdrawPeriod(opts *bind.WatchOpts, sink chan<- *AuctionWithdrawPeriod) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "WithdrawPeriod")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionWithdrawPeriod)
				if err := _Auction.contract.UnpackLog(event, "WithdrawPeriod", log); err != nil {
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

// ParseWithdrawPeriod is a log parse operation binding the contract event 0x6efecf63a32b9b51e8ead9e19a996f69916a2753adab220b8bced7624fb50e34.
//
// Solidity: event WithdrawPeriod(uint256 startWithdraw, uint256 endWithdraw)
func (_Auction *AuctionFilterer) ParseWithdrawPeriod(log types.Log) (*AuctionWithdrawPeriod, error) {
	event := new(AuctionWithdrawPeriod)
	if err := _Auction.contract.UnpackLog(event, "WithdrawPeriod", log); err != nil {
		return nil, err
	}
	return event, nil
}
