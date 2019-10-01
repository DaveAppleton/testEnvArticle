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
const IERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

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
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableBin is the compiled bytecode used for deploying new contracts.
var OwnableBin = "0x608060405234801561001057600080fd5b5060008054600160a060020a03191633179055610237806100326000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b14610081575b600080fd5b34801561005c57600080fd5b506100656100a4565b60408051600160a060020a039092168252519081900360200190f35b34801561008d57600080fd5b506100a2600160a060020a03600435166100b3565b005b600054600160a060020a031690565b600054600160a060020a0316331461012c57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600960248201527f466f7262696464656e0000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a03811615156101a357604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f4e6f6e2d7a65726f20616464726573732072657175697265642e000000000000604482015290519081900360640190fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058202d823bc12e2dbf6ca936a07071e029e152c27862a6f816cbea229fe683b8e7690029"

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
const AuctionABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"withdrawRefund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inReveal\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberOfRevealedValues\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_bid\",\"type\":\"uint256\"},{\"name\":\"_randString\",\"type\":\"bytes\"}],\"name\":\"calculateHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"theAddresses\",\"type\":\"address[]\"},{\"name\":\"flagged\",\"type\":\"bool\"}],\"name\":\"setWinningAddresses\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_startWithdraw\",\"type\":\"uint256\"},{\"name\":\"_endWithdraw\",\"type\":\"uint256\"}],\"name\":\"startWithdrawal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_winningAmount\",\"type\":\"uint256\"}],\"name\":\"setWinningAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"check\",\"type\":\"address\"}],\"name\":\"isWinner\",\"outputs\":[{\"name\":\"winner\",\"type\":\"bool\"},{\"name\":\"inPeriod\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inBidding\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"biddingTime\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"inWithdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"information\",\"outputs\":[{\"name\":\"hashValue\",\"type\":\"bytes32\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"haveBid\",\"type\":\"bool\"},{\"name\":\"revealed\",\"type\":\"bool\"},{\"name\":\"refunded\",\"type\":\"bool\"},{\"name\":\"flagged\",\"type\":\"bool\"},{\"name\":\"bid\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"winners\",\"type\":\"address[]\"}],\"name\":\"earlyWithdrawal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bid\",\"type\":\"uint256\"},{\"name\":\"randString\",\"type\":\"bytes\"}],\"name\":\"reveal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumBid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"check\",\"type\":\"address\"}],\"name\":\"withdrawalAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferAnyERC20Token\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"position\",\"type\":\"uint256\"}],\"name\":\"revealedValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reveals\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"numberOfRevealsForValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_startBids\",\"type\":\"uint256\"},{\"name\":\"_endBids\",\"type\":\"uint256\"},{\"name\":\"_startReveal\",\"type\":\"uint256\"},{\"name\":\"_endReveal\",\"type\":\"uint256\"},{\"name\":\"_minimumBid\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_minimumBid\",\"type\":\"uint256\"}],\"name\":\"MinimumBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"startBids\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endBids\",\"type\":\"uint256\"}],\"name\":\"BiddingPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"startReveal\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endReveal\",\"type\":\"uint256\"}],\"name\":\"RevealPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"startWithdraw\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endWithdraw\",\"type\":\"uint256\"}],\"name\":\"WithdrawPeriod\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"funding\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"BidSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bid\",\"type\":\"uint256\"}],\"name\":\"BidRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"winningAmount\",\"type\":\"uint256\"}],\"name\":\"WinningAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tiedWinner\",\"type\":\"address\"}],\"name\":\"WinnerWithTie\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"NothingToRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"change\",\"type\":\"uint256\"}],\"name\":\"RefundChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BalanceWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"Wallet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

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
var AuctionBin = "0x60806040523480156200001157600080fd5b5060405160c08062001f8f83398101604090815281516020830151918301516060840151608085015160a09095015160008054600160a060020a0319163317905592949192909190851515620000b757604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526016602482015260008051602062001f6f833981519152604482015290519081900360640190fd5b8415156200011557604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526016602482015260008051602062001f6f833981519152604482015290519081900360640190fd5b8315156200017357604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526016602482015260008051602062001f6f833981519152604482015290519081900360640190fd5b821515620001d157604080517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526016602482015260008051602062001f6f833981519152604482015290519081900360640190fd5b8115156200024057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f6d696e696d756d20626964206d757374206265206e6f6e207a65726f00000000604482015290519081900360640190fd5b600160a060020a0381161515620002b857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e76616c69642077616c6c6574206164647265737300000000000000000000604482015290519081900360640190fd5b6005869055600685905560078490556008839055600c829055600d8054600160a060020a031916600160a060020a038316179055604080518781526020810187905281517f892ffcfb665b4ba99a581f2bfa4b8614aaa46bae9247c832b4bd6c1d2badb2e6929181900390910190a17fd23f2f662aba8ca991baba62f1219f0780865580d293d0c7bc89079757639333600754600854604051808381526020018281526020019250505060405180910390a1600c5460408051918252517f855d14e5037e9de69ed821e8030c2c0ccaf36f6ebdb3fdc4689f51aa85a518f09181900360200190a1600d5460408051600160a060020a039092168252517f299e6b07aa8d613d084182e45e438778c5e03395c860c2f28273c66e6f620f0c9181900360200190a1505050505050611b7b80620003f46000396000f3006080604052600436106101325763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663110f88748114610137578063171085621461014e5780634136128414610177578063476343ee1461019e578063500a2b59146101b357806351c65bdf1461021157806370431d291461026a5780637eb4d7d3146102855780638da5cb5b1461029d5780639d9ca28d146102ce5780639eff376d1461030a578063a056e6201461031f578063a0ace77a1461032a578063a30ab2d01461033f578063b856ed861461039c578063ce805642146103f1578063d3a863861461044f578063daa91f2e14610464578063dc39d06d1461049e578063e44d7a69146104c2578063e69650a6146104da578063ed295415146104f5578063f2fde38b1461050d575b600080fd5b34801561014357600080fd5b5061014c61052e565b005b34801561015a57600080fd5b5061016361083b565b604080519115158252519081900360200190f35b34801561018357600080fd5b5061018c610854565b60408051918252519081900360200190f35b3480156101aa57600080fd5b5061014c61085a565b3480156101bf57600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261018c9583359536956044949193909101919081908401838280828437509497506109879650505050505050565b34801561021d57600080fd5b506040805160206004803580820135838102808601850190965280855261014c95369593946024949385019291829185019084908082843750949750505050913515159250610a5a915050565b34801561027657600080fd5b5061014c600435602435610bc5565b34801561029157600080fd5b5061014c600435610d35565b3480156102a957600080fd5b506102b2610e19565b60408051600160a060020a039092168252519081900360200190f35b3480156102da57600080fd5b506102ef600160a060020a0360043516610e28565b60408051921515835290151560208301528051918290030190f35b34801561031657600080fd5b50610163610ec6565b61014c600435610edd565b34801561033657600080fd5b506101636110ef565b34801561034b57600080fd5b50610360600160a060020a0360043516611106565b60408051978852602088019690965293151586860152911515606086015215156080850152151560a084015260c0830152519081900360e00190f35b3480156103a857600080fd5b506040805160206004803580820135838102808601850190965280855261014c9536959394602494938501929182918501908490808284375094975061114e9650505050505050565b3480156103fd57600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261014c9583359536956044949193909101919081908401838280828437509497506113f69650505050505050565b34801561045b57600080fd5b5061018c611717565b34801561047057600080fd5b50610485600160a060020a036004351661171d565b6040805192835290151560208301528051918290030190f35b3480156104aa57600080fd5b50610163600160a060020a036004351660243561180e565b3480156104ce57600080fd5b5061018c600435611915565b3480156104e657600080fd5b506102b2600435602435611992565b34801561050157600080fd5b5061018c6004356119c9565b34801561051957600080fd5b5061014c600160a060020a03600435166119db565b60008060006009544211151561058e576040805160e560020a62461bcd02815260206004820152601c60248201527f576974686472617720686173206e6f7420737461727465642079657400000000604482015290519081900360640190fd5b600a5442106105e7576040805160e560020a62461bcd02815260206004820152601460248201527f576974686472617720706572696f64206f766572000000000000000000000000604482015290519081900360640190fd5b336000908152600160205260408120600381015490945011610653576040805160e560020a62461bcd02815260206004820152601060248201527f4e6f20626964207375626d697474656400000000000000000000000000000000604482015290519081900360640190fd5b6002830154610100900460ff1615156106b6576040805160e560020a62461bcd02815260206004820152601460248201527f62696420776173206e6f742072657665616c6564000000000000000000000000604482015290519081900360640190fd5b600283015462010000900460ff1615610719576040805160e560020a62461bcd02815260206004820152601060248201527f416c726561647920726566756e64656400000000000000000000000000000000604482015290519081900360640190fd5b60028301805462ff00001916620100001790556107353361171d565b909250905080151561078157604080513381526020810184905281517fbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d929181900390910190a16107fe565b60008211156107ca57604080513381526020810184905281517f7417b321c40d8a0c375656909a6cb0a232297cc6addda0f3a386b9f761669d90929181900390910190a16107fe565b6040805133815290517fbb386939e3dac174d852c26db4592a53fb01fde4794fa3124c31416a0c9b9ec09181900360200190a15b600082111561083657604051339083156108fc029084906000818181858888f19350505050158015610834573d6000803e3d6000fd5b505b505050565b60006007544211801561084f575060085442105b905090565b60045490565b600a5442116108b3576040805160e560020a62461bcd02815260206004820152601460248201527f576974686472617720706572696f64206f766572000000000000000000000000604482015290519081900360640190fd5b600054600160a060020a03163314610903576040805160e560020a62461bcd0281526020600482015260096024820152600080516020611b30833981519152604482015290519081900360640190fd5b600d5460408051600160a060020a0390921682523031602083015280517fddc398b321237a8d40ac914388309c2f52a08c134e4dc4ce61e32f57cb7d80f19281900390910190a1600d54604051600160a060020a0390911690303180156108fc02916000818181858888f19350505050158015610984573d6000803e3d6000fd5b50565b600082826040516020018083815260200182805190602001908083835b602083106109c35780518252601f1990920191602091820191016109a4565b6001836020036101000a038019825116818451168082178552505050505050905001925050506040516020818303038152906040526040518082805190602001908083835b60208310610a275780518252601f199092019160209182019101610a08565b5181516020939093036101000a600019018019909116921691909117905260405192018290039091209695505050505050565b60008054600160a060020a03163314610aab576040805160e560020a62461bcd0281526020600482015260096024820152600080516020611b30833981519152604482015290519081900360640190fd5b6008544211610b04576040805160e560020a62461bcd02815260206004820152601860248201527f52657665616c20686173206e6f7420656e646564207965740000000000000000604482015290519081900360640190fd5b5060005b82518110156108365781600160008584815181101515610b2457fe5b90602001906020020151600160a060020a0316600160a060020a0316815260200190815260200160002060020160036101000a81548160ff0219169083151502179055507fdeb65b46e26a4ed6abe63bd73c4613e1262a81bafb7bb35b92f5160e8f1724f08382815181101515610b9757fe5b602090810290910181015160408051600160a060020a039092168252519081900390910190a1600101610b08565b600054600160a060020a03163314610c15576040805160e560020a62461bcd0281526020600482015260096024820152600080516020611b30833981519152604482015290519081900360640190fd5b600b54600010610c6f576040805160e560020a62461bcd02815260206004820152601660248201527f57696e6e696e6720416d6f756e74204e6f742053657400000000000000000000604482015290519081900360640190fd5b428211610cec576040805160e560020a62461bcd02815260206004820152602360248201527f43616e6e6f74207374617274207769746864726177616c20696e20746865207060448201527f6173740000000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6009829055600a819055604080518381526020810183905281517f6efecf63a32b9b51e8ead9e19a996f69916a2753adab220b8bced7624fb50e34929181900390910190a15050565b600054600160a060020a03163314610d85576040805160e560020a62461bcd0281526020600482015260096024820152600080516020611b30833981519152604482015290519081900360640190fd5b6008544211610dde576040805160e560020a62461bcd02815260206004820152601860248201527f52657665616c20686173206e6f7420656e646564207965740000000000000000604482015290519081900360640190fd5b600b8190556040805182815290517f85679d62923fa647b865510f1c0f24c7babbee0a25eca9118b226faf71f42f1c9181900360200190a150565b600054600160a060020a031690565b600080600954421015610e4057506000905080610ec1565b50600a54600b54600160a060020a0384166000908152600160205260409020600301544292909210911180610ebe5750600b54600160a060020a038416600090815260016020526040902060030154148015610ebe5750600160a060020a0383166000908152600160205260409020600201546301000000900460ff165b91505b915091565b60006005544211801561084f575050600654421090565b600554600090421015610f3a576040805160e560020a62461bcd02815260206004820152601760248201527f42696464696e67206e6f74207374617274656420796574000000000000000000604482015290519081900360640190fd5b6006544210610f93576040805160e560020a62461bcd02815260206004820152601160248201527f42696464696e672068617320656e646564000000000000000000000000000000604482015290519081900360640190fd5b3360009081526001602052604090206002015460ff1615610ffe576040805160e560020a62461bcd02815260206004820152601860248201527f6f6e6c79206f6e65206269642070657220616464726573730000000000000000604482015290519081900360640190fd5b600c54341161107c576040805160e560020a62461bcd028152602060048201526024808201527f416d6f756e742073656e74206973206c657373207468616e206d696e696d756d60448201527f2062696400000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b50336000818152600160208181526040928390208581553481840181905560028201805460ff1916909417909355835194855290840191909152828201849052905190917f43dfb6c14d1b44aaeaef2c0487e8423e2843a99f9e09108814b95ce236ed8707919081900360600190a15050565b60006009544211801561084f575050600a54421090565b60016020819052600091825260409091208054918101546002820154600390920154909160ff8082169261010083048216926201000081048316926301000000909104169087565b6000805481908190819081908190600160a060020a031633146111a9576040805160e560020a62461bcd0281526020600482015260096024820152600080516020611b30833981519152604482015290519081900360640190fd5b6009544211611202576040805160e560020a62461bcd02815260206004820152601c60248201527f576974686472617720686173206e6f7420737461727465642079657400000000604482015290519081900360640190fd5b600a54421061125b576040805160e560020a62461bcd02815260206004820152601460248201527f576974686472617720706572696f64206f766572000000000000000000000000604482015290519081900360640190fd5b86519350600094505b8385101561136b57868581518110151561127a57fe5b6020908102909101810151600160a060020a0381166000908152600390925260409091205490935060ff16151561133b576112b483610e28565b90925090508180156112c35750805b1515611319576040805160e560020a62461bcd02815260206004820152600c60248201527f4e6f7420612077696e6e65720000000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a03831660009081526001602052604090206003015495909501945b600160a060020a0383166000908152600360205260409020805460ff191660019081179091559490940193611264565b600d5460408051600160a060020a0390921682526020820188905280517fddc398b321237a8d40ac914388309c2f52a08c134e4dc4ce61e32f57cb7d80f19281900390910190a1600d54604051600160a060020a039091169087156108fc029088906000818181858888f193505050501580156113ec573d6000803e3d6000fd5b5050505050505050565b60008060006007544210151515611457576040805160e560020a62461bcd02815260206004820152601660248201527f52657665616c206e6f7420737461727465642079657400000000000000000000604482015290519081900360640190fd5b60085442106114b0576040805160e560020a62461bcd02815260206004820152601060248201527f52657665616c2068617320656e64656400000000000000000000000000000000604482015290519081900360640190fd5b6114ba8585610987565b336000908152600160205260409020805491945092508314611526576040805160e560020a62461bcd02815260206004820152601360248201527f68617368657320646f206e6f74206d6174636800000000000000000000000000604482015290519081900360640190fd5b6001820154851115611582576040805160e560020a62461bcd02815260206004820152600d60248201527f426964206e6f742076616c696400000000000000000000000000000000000000604482015290519081900360640190fd5b600c548510156115dc576040805160e560020a62461bcd02815260206004820152601d60248201527f42696420776173206c657373207468616e206d696e696d756d20626964000000604482015290519081900360640190fd5b6002820154610100900460ff161561163e576040805160e560020a62461bcd02815260206004820152601460248201527f42696420616c72656164792072657665616c6564000000000000000000000000604482015290519081900360640190fd5b50600381018490556002808201805461ff00191661010017905560008581526020918252604081208054600181810180845592845293909220909101805473ffffffffffffffffffffffffffffffffffffffff191633179055908114156116d557600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018590555b604080513381526020810187905281517f9891bd98143d5a2c5f29be9a7ef8e5d7b723746fb3b75220234af7aa6caeecfe929181900390910190a15050505050565b600c5481565b60008060008061172b611af3565b61173486610e28565b909350915081151561174c5760009450849350611806565b50600160a060020a038516600090815260016020818152604092839020835160e081018552815481529281015491830191909152600281015460ff808216151594840194909452610100810484161515606084018190526201000082048516151560808501526301000000909104909316151560a08301526003015460c0820152906117de5760009450849350611806565b82156117fa578060c00151816020015103600194509450611806565b80602001516000945094505b505050915091565b60008054600160a060020a0316331461185f576040805160e560020a62461bcd0281526020600482015260096024820152600080516020611b30833981519152604482015290519081900360640190fd5b82600160a060020a031663a9059cbb611876610e19565b846040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156118e257600080fd5b505af11580156118f6573d6000803e3d6000fd5b505050506040513d602081101561190c57600080fd5b50519392505050565b600454600090821115611972576040805160e560020a62461bcd02815260206004820152601560248201527f706f736974696f6e206e6f7420696e2061727261790000000000000000000000604482015290519081900360640190fd5b600480548390811061198057fe5b90600052602060002001549050919050565b6002602052816000526040600020818154811015156119ad57fe5b600091825260209091200154600160a060020a03169150829050565b60009081526002602052604090205490565b600054600160a060020a03163314611a2b576040805160e560020a62461bcd0281526020600482015260096024820152600080516020611b30833981519152604482015290519081900360640190fd5b600160a060020a0381161515611a8b576040805160e560020a62461bcd02815260206004820152601a60248201527f4e6f6e2d7a65726f20616464726573732072657175697265642e000000000000604482015290519081900360640190fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152905600466f7262696464656e0000000000000000000000000000000000000000000000a165627a7a72305820a700d8f5bd0eab0f4ad9dae04c222acefb7b69239b732f742fa1d3cfe44f4c7e00296461746573206d757374206265206e6f6e207a65726f00000000000000000000"

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
