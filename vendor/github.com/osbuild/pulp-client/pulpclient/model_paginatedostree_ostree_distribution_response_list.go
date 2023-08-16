/*
Pulp 3 API

Fetch, Upload, Organize, and Distribute Software Packages

API version: v3
Contact: pulp-list@redhat.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pulpclient

import (
	"encoding/json"
)

// checks if the PaginatedostreeOstreeDistributionResponseList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedostreeOstreeDistributionResponseList{}

// PaginatedostreeOstreeDistributionResponseList struct for PaginatedostreeOstreeDistributionResponseList
type PaginatedostreeOstreeDistributionResponseList struct {
	Count *int32 `json:"count,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Results []OstreeOstreeDistributionResponse `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedostreeOstreeDistributionResponseList PaginatedostreeOstreeDistributionResponseList

// NewPaginatedostreeOstreeDistributionResponseList instantiates a new PaginatedostreeOstreeDistributionResponseList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedostreeOstreeDistributionResponseList() *PaginatedostreeOstreeDistributionResponseList {
	this := PaginatedostreeOstreeDistributionResponseList{}
	return &this
}

// NewPaginatedostreeOstreeDistributionResponseListWithDefaults instantiates a new PaginatedostreeOstreeDistributionResponseList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedostreeOstreeDistributionResponseListWithDefaults() *PaginatedostreeOstreeDistributionResponseList {
	this := PaginatedostreeOstreeDistributionResponseList{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PaginatedostreeOstreeDistributionResponseList) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedostreeOstreeDistributionResponseList) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PaginatedostreeOstreeDistributionResponseList) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PaginatedostreeOstreeDistributionResponseList) SetCount(v int32) {
	o.Count = &v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedostreeOstreeDistributionResponseList) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedostreeOstreeDistributionResponseList) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedostreeOstreeDistributionResponseList) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedostreeOstreeDistributionResponseList) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedostreeOstreeDistributionResponseList) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedostreeOstreeDistributionResponseList) UnsetNext() {
	o.Next.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedostreeOstreeDistributionResponseList) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedostreeOstreeDistributionResponseList) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedostreeOstreeDistributionResponseList) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedostreeOstreeDistributionResponseList) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedostreeOstreeDistributionResponseList) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedostreeOstreeDistributionResponseList) UnsetPrevious() {
	o.Previous.Unset()
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PaginatedostreeOstreeDistributionResponseList) GetResults() []OstreeOstreeDistributionResponse {
	if o == nil || IsNil(o.Results) {
		var ret []OstreeOstreeDistributionResponse
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginatedostreeOstreeDistributionResponseList) GetResultsOk() ([]OstreeOstreeDistributionResponse, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PaginatedostreeOstreeDistributionResponseList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OstreeOstreeDistributionResponse and assigns it to the Results field.
func (o *PaginatedostreeOstreeDistributionResponseList) SetResults(v []OstreeOstreeDistributionResponse) {
	o.Results = v
}

func (o PaginatedostreeOstreeDistributionResponseList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedostreeOstreeDistributionResponseList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedostreeOstreeDistributionResponseList) UnmarshalJSON(bytes []byte) (err error) {
	varPaginatedostreeOstreeDistributionResponseList := _PaginatedostreeOstreeDistributionResponseList{}

	if err = json.Unmarshal(bytes, &varPaginatedostreeOstreeDistributionResponseList); err == nil {
		*o = PaginatedostreeOstreeDistributionResponseList(varPaginatedostreeOstreeDistributionResponseList)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "next")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedostreeOstreeDistributionResponseList struct {
	value *PaginatedostreeOstreeDistributionResponseList
	isSet bool
}

func (v NullablePaginatedostreeOstreeDistributionResponseList) Get() *PaginatedostreeOstreeDistributionResponseList {
	return v.value
}

func (v *NullablePaginatedostreeOstreeDistributionResponseList) Set(val *PaginatedostreeOstreeDistributionResponseList) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedostreeOstreeDistributionResponseList) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedostreeOstreeDistributionResponseList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedostreeOstreeDistributionResponseList(val *PaginatedostreeOstreeDistributionResponseList) *NullablePaginatedostreeOstreeDistributionResponseList {
	return &NullablePaginatedostreeOstreeDistributionResponseList{value: val, isSet: true}
}

func (v NullablePaginatedostreeOstreeDistributionResponseList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedostreeOstreeDistributionResponseList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


