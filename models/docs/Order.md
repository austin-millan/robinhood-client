# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Action** | Pointer to [**OrderAction**](OrderAction.md) |  | [optional] 
**AveragePrice** | Pointer to **string** |  | [optional] 
**CancelUrl** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CumulativeQuantity** | Pointer to **string** |  | [optional] 
**DollarBasedAmount** | Pointer to **string** |  | [optional] 
**Executions** | Pointer to [**[]Execution**](Execution.md) |  | [optional] 
**TotalNotional** | Pointer to [**OrderTotalNotional**](OrderTotalNotional.md) |  | [optional] 
**ExecutedNotional** | Pointer to [**OrderTotalNotional**](OrderTotalNotional.md) |  | [optional] 
**ExtendedHours** | Pointer to **bool** |  | [optional] 
**Fees** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InvestmentScheduleId** | Pointer to **string** |  | [optional] 
**Instrument** | Pointer to **string** |  | [optional] 
**LastTrailPrice** | Pointer to **string** |  | [optional] 
**LastTrailPriceUpdatedAt** | Pointer to **string** |  | [optional] 
**LastTransactionAt** | Pointer to **string** |  | [optional] 
**OverrideDayTradeChecks** | Pointer to **bool** |  | [optional] 
**OverrideDtbpChecks** | Pointer to **bool** |  | [optional] 
**Position** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**RejectReason** | Pointer to **string** |  | [optional] 
**Side** | Pointer to [**Side**](Side.md) |  | [optional] 
**State** | Pointer to [**OrderState**](OrderState.md) |  | [optional] 
**StopPrice** | Pointer to **string** |  | [optional] 
**StopTriggeredAt** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**TimeInForce** | Pointer to [**TimeInForce**](TimeInForce.md) |  | [optional] 
**Trigger** | Pointer to [**Trigger**](Trigger.md) |  | [optional] 
**Type** | Pointer to [**ExecutionType**](ExecutionType.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Order) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Order) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Order) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Order) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAction

`func (o *Order) GetAction() OrderAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Order) GetActionOk() (*OrderAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Order) SetAction(v OrderAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *Order) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAveragePrice

`func (o *Order) GetAveragePrice() string`

GetAveragePrice returns the AveragePrice field if non-nil, zero value otherwise.

### GetAveragePriceOk

`func (o *Order) GetAveragePriceOk() (*string, bool)`

GetAveragePriceOk returns a tuple with the AveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePrice

`func (o *Order) SetAveragePrice(v string)`

SetAveragePrice sets AveragePrice field to given value.

### HasAveragePrice

`func (o *Order) HasAveragePrice() bool`

HasAveragePrice returns a boolean if a field has been set.

### GetCancelUrl

`func (o *Order) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *Order) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *Order) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.

### HasCancelUrl

`func (o *Order) HasCancelUrl() bool`

HasCancelUrl returns a boolean if a field has been set.

### GetClientId

`func (o *Order) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Order) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Order) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Order) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Order) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Order) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Order) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Order) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCumulativeQuantity

`func (o *Order) GetCumulativeQuantity() string`

GetCumulativeQuantity returns the CumulativeQuantity field if non-nil, zero value otherwise.

### GetCumulativeQuantityOk

`func (o *Order) GetCumulativeQuantityOk() (*string, bool)`

GetCumulativeQuantityOk returns a tuple with the CumulativeQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCumulativeQuantity

`func (o *Order) SetCumulativeQuantity(v string)`

SetCumulativeQuantity sets CumulativeQuantity field to given value.

### HasCumulativeQuantity

`func (o *Order) HasCumulativeQuantity() bool`

HasCumulativeQuantity returns a boolean if a field has been set.

### GetDollarBasedAmount

`func (o *Order) GetDollarBasedAmount() string`

GetDollarBasedAmount returns the DollarBasedAmount field if non-nil, zero value otherwise.

### GetDollarBasedAmountOk

`func (o *Order) GetDollarBasedAmountOk() (*string, bool)`

GetDollarBasedAmountOk returns a tuple with the DollarBasedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDollarBasedAmount

`func (o *Order) SetDollarBasedAmount(v string)`

SetDollarBasedAmount sets DollarBasedAmount field to given value.

### HasDollarBasedAmount

`func (o *Order) HasDollarBasedAmount() bool`

HasDollarBasedAmount returns a boolean if a field has been set.

### GetExecutions

`func (o *Order) GetExecutions() []Execution`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *Order) GetExecutionsOk() (*[]Execution, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *Order) SetExecutions(v []Execution)`

SetExecutions sets Executions field to given value.

### HasExecutions

`func (o *Order) HasExecutions() bool`

HasExecutions returns a boolean if a field has been set.

### GetTotalNotional

`func (o *Order) GetTotalNotional() OrderTotalNotional`

GetTotalNotional returns the TotalNotional field if non-nil, zero value otherwise.

### GetTotalNotionalOk

`func (o *Order) GetTotalNotionalOk() (*OrderTotalNotional, bool)`

GetTotalNotionalOk returns a tuple with the TotalNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNotional

`func (o *Order) SetTotalNotional(v OrderTotalNotional)`

SetTotalNotional sets TotalNotional field to given value.

### HasTotalNotional

`func (o *Order) HasTotalNotional() bool`

HasTotalNotional returns a boolean if a field has been set.

### GetExecutedNotional

`func (o *Order) GetExecutedNotional() OrderTotalNotional`

GetExecutedNotional returns the ExecutedNotional field if non-nil, zero value otherwise.

### GetExecutedNotionalOk

`func (o *Order) GetExecutedNotionalOk() (*OrderTotalNotional, bool)`

GetExecutedNotionalOk returns a tuple with the ExecutedNotional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedNotional

`func (o *Order) SetExecutedNotional(v OrderTotalNotional)`

SetExecutedNotional sets ExecutedNotional field to given value.

### HasExecutedNotional

`func (o *Order) HasExecutedNotional() bool`

HasExecutedNotional returns a boolean if a field has been set.

### GetExtendedHours

`func (o *Order) GetExtendedHours() bool`

GetExtendedHours returns the ExtendedHours field if non-nil, zero value otherwise.

### GetExtendedHoursOk

`func (o *Order) GetExtendedHoursOk() (*bool, bool)`

GetExtendedHoursOk returns a tuple with the ExtendedHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedHours

`func (o *Order) SetExtendedHours(v bool)`

SetExtendedHours sets ExtendedHours field to given value.

### HasExtendedHours

`func (o *Order) HasExtendedHours() bool`

HasExtendedHours returns a boolean if a field has been set.

### GetFees

`func (o *Order) GetFees() string`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *Order) GetFeesOk() (*string, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *Order) SetFees(v string)`

SetFees sets Fees field to given value.

### HasFees

`func (o *Order) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetId

`func (o *Order) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Order) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvestmentScheduleId

`func (o *Order) GetInvestmentScheduleId() string`

GetInvestmentScheduleId returns the InvestmentScheduleId field if non-nil, zero value otherwise.

### GetInvestmentScheduleIdOk

`func (o *Order) GetInvestmentScheduleIdOk() (*string, bool)`

GetInvestmentScheduleIdOk returns a tuple with the InvestmentScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentScheduleId

`func (o *Order) SetInvestmentScheduleId(v string)`

SetInvestmentScheduleId sets InvestmentScheduleId field to given value.

### HasInvestmentScheduleId

`func (o *Order) HasInvestmentScheduleId() bool`

HasInvestmentScheduleId returns a boolean if a field has been set.

### GetInstrument

`func (o *Order) GetInstrument() string`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *Order) GetInstrumentOk() (*string, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *Order) SetInstrument(v string)`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *Order) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### GetLastTrailPrice

`func (o *Order) GetLastTrailPrice() string`

GetLastTrailPrice returns the LastTrailPrice field if non-nil, zero value otherwise.

### GetLastTrailPriceOk

`func (o *Order) GetLastTrailPriceOk() (*string, bool)`

GetLastTrailPriceOk returns a tuple with the LastTrailPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrailPrice

`func (o *Order) SetLastTrailPrice(v string)`

SetLastTrailPrice sets LastTrailPrice field to given value.

### HasLastTrailPrice

`func (o *Order) HasLastTrailPrice() bool`

HasLastTrailPrice returns a boolean if a field has been set.

### GetLastTrailPriceUpdatedAt

`func (o *Order) GetLastTrailPriceUpdatedAt() string`

GetLastTrailPriceUpdatedAt returns the LastTrailPriceUpdatedAt field if non-nil, zero value otherwise.

### GetLastTrailPriceUpdatedAtOk

`func (o *Order) GetLastTrailPriceUpdatedAtOk() (*string, bool)`

GetLastTrailPriceUpdatedAtOk returns a tuple with the LastTrailPriceUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTrailPriceUpdatedAt

`func (o *Order) SetLastTrailPriceUpdatedAt(v string)`

SetLastTrailPriceUpdatedAt sets LastTrailPriceUpdatedAt field to given value.

### HasLastTrailPriceUpdatedAt

`func (o *Order) HasLastTrailPriceUpdatedAt() bool`

HasLastTrailPriceUpdatedAt returns a boolean if a field has been set.

### GetLastTransactionAt

`func (o *Order) GetLastTransactionAt() string`

GetLastTransactionAt returns the LastTransactionAt field if non-nil, zero value otherwise.

### GetLastTransactionAtOk

`func (o *Order) GetLastTransactionAtOk() (*string, bool)`

GetLastTransactionAtOk returns a tuple with the LastTransactionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransactionAt

`func (o *Order) SetLastTransactionAt(v string)`

SetLastTransactionAt sets LastTransactionAt field to given value.

### HasLastTransactionAt

`func (o *Order) HasLastTransactionAt() bool`

HasLastTransactionAt returns a boolean if a field has been set.

### GetOverrideDayTradeChecks

`func (o *Order) GetOverrideDayTradeChecks() bool`

GetOverrideDayTradeChecks returns the OverrideDayTradeChecks field if non-nil, zero value otherwise.

### GetOverrideDayTradeChecksOk

`func (o *Order) GetOverrideDayTradeChecksOk() (*bool, bool)`

GetOverrideDayTradeChecksOk returns a tuple with the OverrideDayTradeChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDayTradeChecks

`func (o *Order) SetOverrideDayTradeChecks(v bool)`

SetOverrideDayTradeChecks sets OverrideDayTradeChecks field to given value.

### HasOverrideDayTradeChecks

`func (o *Order) HasOverrideDayTradeChecks() bool`

HasOverrideDayTradeChecks returns a boolean if a field has been set.

### GetOverrideDtbpChecks

`func (o *Order) GetOverrideDtbpChecks() bool`

GetOverrideDtbpChecks returns the OverrideDtbpChecks field if non-nil, zero value otherwise.

### GetOverrideDtbpChecksOk

`func (o *Order) GetOverrideDtbpChecksOk() (*bool, bool)`

GetOverrideDtbpChecksOk returns a tuple with the OverrideDtbpChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideDtbpChecks

`func (o *Order) SetOverrideDtbpChecks(v bool)`

SetOverrideDtbpChecks sets OverrideDtbpChecks field to given value.

### HasOverrideDtbpChecks

`func (o *Order) HasOverrideDtbpChecks() bool`

HasOverrideDtbpChecks returns a boolean if a field has been set.

### GetPosition

`func (o *Order) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Order) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Order) SetPosition(v string)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Order) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPrice

`func (o *Order) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Order) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Order) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Order) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *Order) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Order) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Order) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Order) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRefId

`func (o *Order) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *Order) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *Order) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *Order) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRejectReason

`func (o *Order) GetRejectReason() string`

GetRejectReason returns the RejectReason field if non-nil, zero value otherwise.

### GetRejectReasonOk

`func (o *Order) GetRejectReasonOk() (*string, bool)`

GetRejectReasonOk returns a tuple with the RejectReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectReason

`func (o *Order) SetRejectReason(v string)`

SetRejectReason sets RejectReason field to given value.

### HasRejectReason

`func (o *Order) HasRejectReason() bool`

HasRejectReason returns a boolean if a field has been set.

### GetSide

`func (o *Order) GetSide() Side`

GetSide returns the Side field if non-nil, zero value otherwise.

### GetSideOk

`func (o *Order) GetSideOk() (*Side, bool)`

GetSideOk returns a tuple with the Side field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSide

`func (o *Order) SetSide(v Side)`

SetSide sets Side field to given value.

### HasSide

`func (o *Order) HasSide() bool`

HasSide returns a boolean if a field has been set.

### GetState

`func (o *Order) GetState() OrderState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Order) GetStateOk() (*OrderState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Order) SetState(v OrderState)`

SetState sets State field to given value.

### HasState

`func (o *Order) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStopPrice

`func (o *Order) GetStopPrice() string`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *Order) GetStopPriceOk() (*string, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *Order) SetStopPrice(v string)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *Order) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStopTriggeredAt

`func (o *Order) GetStopTriggeredAt() string`

GetStopTriggeredAt returns the StopTriggeredAt field if non-nil, zero value otherwise.

### GetStopTriggeredAtOk

`func (o *Order) GetStopTriggeredAtOk() (*string, bool)`

GetStopTriggeredAtOk returns a tuple with the StopTriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopTriggeredAt

`func (o *Order) SetStopTriggeredAt(v string)`

SetStopTriggeredAt sets StopTriggeredAt field to given value.

### HasStopTriggeredAt

`func (o *Order) HasStopTriggeredAt() bool`

HasStopTriggeredAt returns a boolean if a field has been set.

### GetSymbol

`func (o *Order) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Order) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Order) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Order) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTimeInForce

`func (o *Order) GetTimeInForce() TimeInForce`

GetTimeInForce returns the TimeInForce field if non-nil, zero value otherwise.

### GetTimeInForceOk

`func (o *Order) GetTimeInForceOk() (*TimeInForce, bool)`

GetTimeInForceOk returns a tuple with the TimeInForce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInForce

`func (o *Order) SetTimeInForce(v TimeInForce)`

SetTimeInForce sets TimeInForce field to given value.

### HasTimeInForce

`func (o *Order) HasTimeInForce() bool`

HasTimeInForce returns a boolean if a field has been set.

### GetTrigger

`func (o *Order) GetTrigger() Trigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *Order) GetTriggerOk() (*Trigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *Order) SetTrigger(v Trigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *Order) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetType

`func (o *Order) GetType() ExecutionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Order) GetTypeOk() (*ExecutionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Order) SetType(v ExecutionType)`

SetType sets Type field to given value.

### HasType

`func (o *Order) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Order) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Order) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Order) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Order) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *Order) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Order) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Order) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Order) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


