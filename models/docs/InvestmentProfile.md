# InvestmentProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnualIncome** | Pointer to **string** |  | [optional] 
**InvestmentExperience** | Pointer to [**InvestmentExperience**](InvestmentExperience.md) |  | [optional] 
**InvestmentObjective** | Pointer to [**InvestmentObjective**](InvestmentObjective.md) |  | [optional] 
**LiquidNetWorth** | Pointer to **string** |  | [optional] 
**LiquidityNeeds** | Pointer to [**LiquidityNeeds**](LiquidityNeeds.md) |  | [optional] 
**RiskTolerance** | Pointer to [**RiskTolerance**](RiskTolerance.md) |  | [optional] 
**SourceOfFunds** | Pointer to [**SourceOfFunds**](SourceOfFunds.md) |  | [optional] 
**SuitabilityVerified** | Pointer to **bool** |  | [optional] 
**TaxBracket** | Pointer to [**TaxBracket**](TaxBracket.md) |  | [optional] 
**TimeHorizon** | Pointer to [**TimeHorizon**](TimeHorizon.md) |  | [optional] 
**TotalNetWorth** | Pointer to [**TotalNetWorth**](TotalNetWorth.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewInvestmentProfile

`func NewInvestmentProfile() *InvestmentProfile`

NewInvestmentProfile instantiates a new InvestmentProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvestmentProfileWithDefaults

`func NewInvestmentProfileWithDefaults() *InvestmentProfile`

NewInvestmentProfileWithDefaults instantiates a new InvestmentProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnualIncome

`func (o *InvestmentProfile) GetAnnualIncome() string`

GetAnnualIncome returns the AnnualIncome field if non-nil, zero value otherwise.

### GetAnnualIncomeOk

`func (o *InvestmentProfile) GetAnnualIncomeOk() (*string, bool)`

GetAnnualIncomeOk returns a tuple with the AnnualIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualIncome

`func (o *InvestmentProfile) SetAnnualIncome(v string)`

SetAnnualIncome sets AnnualIncome field to given value.

### HasAnnualIncome

`func (o *InvestmentProfile) HasAnnualIncome() bool`

HasAnnualIncome returns a boolean if a field has been set.

### GetInvestmentExperience

`func (o *InvestmentProfile) GetInvestmentExperience() InvestmentExperience`

GetInvestmentExperience returns the InvestmentExperience field if non-nil, zero value otherwise.

### GetInvestmentExperienceOk

`func (o *InvestmentProfile) GetInvestmentExperienceOk() (*InvestmentExperience, bool)`

GetInvestmentExperienceOk returns a tuple with the InvestmentExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentExperience

`func (o *InvestmentProfile) SetInvestmentExperience(v InvestmentExperience)`

SetInvestmentExperience sets InvestmentExperience field to given value.

### HasInvestmentExperience

`func (o *InvestmentProfile) HasInvestmentExperience() bool`

HasInvestmentExperience returns a boolean if a field has been set.

### GetInvestmentObjective

`func (o *InvestmentProfile) GetInvestmentObjective() InvestmentObjective`

GetInvestmentObjective returns the InvestmentObjective field if non-nil, zero value otherwise.

### GetInvestmentObjectiveOk

`func (o *InvestmentProfile) GetInvestmentObjectiveOk() (*InvestmentObjective, bool)`

GetInvestmentObjectiveOk returns a tuple with the InvestmentObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvestmentObjective

`func (o *InvestmentProfile) SetInvestmentObjective(v InvestmentObjective)`

SetInvestmentObjective sets InvestmentObjective field to given value.

### HasInvestmentObjective

`func (o *InvestmentProfile) HasInvestmentObjective() bool`

HasInvestmentObjective returns a boolean if a field has been set.

### GetLiquidNetWorth

`func (o *InvestmentProfile) GetLiquidNetWorth() string`

GetLiquidNetWorth returns the LiquidNetWorth field if non-nil, zero value otherwise.

### GetLiquidNetWorthOk

`func (o *InvestmentProfile) GetLiquidNetWorthOk() (*string, bool)`

GetLiquidNetWorthOk returns a tuple with the LiquidNetWorth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidNetWorth

`func (o *InvestmentProfile) SetLiquidNetWorth(v string)`

SetLiquidNetWorth sets LiquidNetWorth field to given value.

### HasLiquidNetWorth

`func (o *InvestmentProfile) HasLiquidNetWorth() bool`

HasLiquidNetWorth returns a boolean if a field has been set.

### GetLiquidityNeeds

`func (o *InvestmentProfile) GetLiquidityNeeds() LiquidityNeeds`

GetLiquidityNeeds returns the LiquidityNeeds field if non-nil, zero value otherwise.

### GetLiquidityNeedsOk

`func (o *InvestmentProfile) GetLiquidityNeedsOk() (*LiquidityNeeds, bool)`

GetLiquidityNeedsOk returns a tuple with the LiquidityNeeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidityNeeds

`func (o *InvestmentProfile) SetLiquidityNeeds(v LiquidityNeeds)`

SetLiquidityNeeds sets LiquidityNeeds field to given value.

### HasLiquidityNeeds

`func (o *InvestmentProfile) HasLiquidityNeeds() bool`

HasLiquidityNeeds returns a boolean if a field has been set.

### GetRiskTolerance

`func (o *InvestmentProfile) GetRiskTolerance() RiskTolerance`

GetRiskTolerance returns the RiskTolerance field if non-nil, zero value otherwise.

### GetRiskToleranceOk

`func (o *InvestmentProfile) GetRiskToleranceOk() (*RiskTolerance, bool)`

GetRiskToleranceOk returns a tuple with the RiskTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskTolerance

`func (o *InvestmentProfile) SetRiskTolerance(v RiskTolerance)`

SetRiskTolerance sets RiskTolerance field to given value.

### HasRiskTolerance

`func (o *InvestmentProfile) HasRiskTolerance() bool`

HasRiskTolerance returns a boolean if a field has been set.

### GetSourceOfFunds

`func (o *InvestmentProfile) GetSourceOfFunds() SourceOfFunds`

GetSourceOfFunds returns the SourceOfFunds field if non-nil, zero value otherwise.

### GetSourceOfFundsOk

`func (o *InvestmentProfile) GetSourceOfFundsOk() (*SourceOfFunds, bool)`

GetSourceOfFundsOk returns a tuple with the SourceOfFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfFunds

`func (o *InvestmentProfile) SetSourceOfFunds(v SourceOfFunds)`

SetSourceOfFunds sets SourceOfFunds field to given value.

### HasSourceOfFunds

`func (o *InvestmentProfile) HasSourceOfFunds() bool`

HasSourceOfFunds returns a boolean if a field has been set.

### GetSuitabilityVerified

`func (o *InvestmentProfile) GetSuitabilityVerified() bool`

GetSuitabilityVerified returns the SuitabilityVerified field if non-nil, zero value otherwise.

### GetSuitabilityVerifiedOk

`func (o *InvestmentProfile) GetSuitabilityVerifiedOk() (*bool, bool)`

GetSuitabilityVerifiedOk returns a tuple with the SuitabilityVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuitabilityVerified

`func (o *InvestmentProfile) SetSuitabilityVerified(v bool)`

SetSuitabilityVerified sets SuitabilityVerified field to given value.

### HasSuitabilityVerified

`func (o *InvestmentProfile) HasSuitabilityVerified() bool`

HasSuitabilityVerified returns a boolean if a field has been set.

### GetTaxBracket

`func (o *InvestmentProfile) GetTaxBracket() TaxBracket`

GetTaxBracket returns the TaxBracket field if non-nil, zero value otherwise.

### GetTaxBracketOk

`func (o *InvestmentProfile) GetTaxBracketOk() (*TaxBracket, bool)`

GetTaxBracketOk returns a tuple with the TaxBracket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxBracket

`func (o *InvestmentProfile) SetTaxBracket(v TaxBracket)`

SetTaxBracket sets TaxBracket field to given value.

### HasTaxBracket

`func (o *InvestmentProfile) HasTaxBracket() bool`

HasTaxBracket returns a boolean if a field has been set.

### GetTimeHorizon

`func (o *InvestmentProfile) GetTimeHorizon() TimeHorizon`

GetTimeHorizon returns the TimeHorizon field if non-nil, zero value otherwise.

### GetTimeHorizonOk

`func (o *InvestmentProfile) GetTimeHorizonOk() (*TimeHorizon, bool)`

GetTimeHorizonOk returns a tuple with the TimeHorizon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeHorizon

`func (o *InvestmentProfile) SetTimeHorizon(v TimeHorizon)`

SetTimeHorizon sets TimeHorizon field to given value.

### HasTimeHorizon

`func (o *InvestmentProfile) HasTimeHorizon() bool`

HasTimeHorizon returns a boolean if a field has been set.

### GetTotalNetWorth

`func (o *InvestmentProfile) GetTotalNetWorth() TotalNetWorth`

GetTotalNetWorth returns the TotalNetWorth field if non-nil, zero value otherwise.

### GetTotalNetWorthOk

`func (o *InvestmentProfile) GetTotalNetWorthOk() (*TotalNetWorth, bool)`

GetTotalNetWorthOk returns a tuple with the TotalNetWorth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNetWorth

`func (o *InvestmentProfile) SetTotalNetWorth(v TotalNetWorth)`

SetTotalNetWorth sets TotalNetWorth field to given value.

### HasTotalNetWorth

`func (o *InvestmentProfile) HasTotalNetWorth() bool`

HasTotalNetWorth returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InvestmentProfile) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InvestmentProfile) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InvestmentProfile) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InvestmentProfile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *InvestmentProfile) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *InvestmentProfile) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *InvestmentProfile) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *InvestmentProfile) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


