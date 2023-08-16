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

// checks if the PaginatedCollectionResponseListLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedCollectionResponseListLinks{}

// PaginatedCollectionResponseListLinks struct for PaginatedCollectionResponseListLinks
type PaginatedCollectionResponseListLinks struct {
	First NullableString `json:"first,omitempty"`
	Previous NullableString `json:"previous,omitempty"`
	Next NullableString `json:"next,omitempty"`
	Last NullableString `json:"last,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginatedCollectionResponseListLinks PaginatedCollectionResponseListLinks

// NewPaginatedCollectionResponseListLinks instantiates a new PaginatedCollectionResponseListLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedCollectionResponseListLinks() *PaginatedCollectionResponseListLinks {
	this := PaginatedCollectionResponseListLinks{}
	return &this
}

// NewPaginatedCollectionResponseListLinksWithDefaults instantiates a new PaginatedCollectionResponseListLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedCollectionResponseListLinksWithDefaults() *PaginatedCollectionResponseListLinks {
	this := PaginatedCollectionResponseListLinks{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedCollectionResponseListLinks) GetFirst() string {
	if o == nil || IsNil(o.First.Get()) {
		var ret string
		return ret
	}
	return *o.First.Get()
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedCollectionResponseListLinks) GetFirstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.First.Get(), o.First.IsSet()
}

// HasFirst returns a boolean if a field has been set.
func (o *PaginatedCollectionResponseListLinks) HasFirst() bool {
	if o != nil && o.First.IsSet() {
		return true
	}

	return false
}

// SetFirst gets a reference to the given NullableString and assigns it to the First field.
func (o *PaginatedCollectionResponseListLinks) SetFirst(v string) {
	o.First.Set(&v)
}
// SetFirstNil sets the value for First to be an explicit nil
func (o *PaginatedCollectionResponseListLinks) SetFirstNil() {
	o.First.Set(nil)
}

// UnsetFirst ensures that no value is present for First, not even an explicit nil
func (o *PaginatedCollectionResponseListLinks) UnsetFirst() {
	o.First.Unset()
}

// GetPrevious returns the Previous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedCollectionResponseListLinks) GetPrevious() string {
	if o == nil || IsNil(o.Previous.Get()) {
		var ret string
		return ret
	}
	return *o.Previous.Get()
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedCollectionResponseListLinks) GetPreviousOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Previous.Get(), o.Previous.IsSet()
}

// HasPrevious returns a boolean if a field has been set.
func (o *PaginatedCollectionResponseListLinks) HasPrevious() bool {
	if o != nil && o.Previous.IsSet() {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given NullableString and assigns it to the Previous field.
func (o *PaginatedCollectionResponseListLinks) SetPrevious(v string) {
	o.Previous.Set(&v)
}
// SetPreviousNil sets the value for Previous to be an explicit nil
func (o *PaginatedCollectionResponseListLinks) SetPreviousNil() {
	o.Previous.Set(nil)
}

// UnsetPrevious ensures that no value is present for Previous, not even an explicit nil
func (o *PaginatedCollectionResponseListLinks) UnsetPrevious() {
	o.Previous.Unset()
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedCollectionResponseListLinks) GetNext() string {
	if o == nil || IsNil(o.Next.Get()) {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedCollectionResponseListLinks) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginatedCollectionResponseListLinks) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginatedCollectionResponseListLinks) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginatedCollectionResponseListLinks) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginatedCollectionResponseListLinks) UnsetNext() {
	o.Next.Unset()
}

// GetLast returns the Last field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedCollectionResponseListLinks) GetLast() string {
	if o == nil || IsNil(o.Last.Get()) {
		var ret string
		return ret
	}
	return *o.Last.Get()
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedCollectionResponseListLinks) GetLastOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Last.Get(), o.Last.IsSet()
}

// HasLast returns a boolean if a field has been set.
func (o *PaginatedCollectionResponseListLinks) HasLast() bool {
	if o != nil && o.Last.IsSet() {
		return true
	}

	return false
}

// SetLast gets a reference to the given NullableString and assigns it to the Last field.
func (o *PaginatedCollectionResponseListLinks) SetLast(v string) {
	o.Last.Set(&v)
}
// SetLastNil sets the value for Last to be an explicit nil
func (o *PaginatedCollectionResponseListLinks) SetLastNil() {
	o.Last.Set(nil)
}

// UnsetLast ensures that no value is present for Last, not even an explicit nil
func (o *PaginatedCollectionResponseListLinks) UnsetLast() {
	o.Last.Unset()
}

func (o PaginatedCollectionResponseListLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedCollectionResponseListLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.First.IsSet() {
		toSerialize["first"] = o.First.Get()
	}
	if o.Previous.IsSet() {
		toSerialize["previous"] = o.Previous.Get()
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	if o.Last.IsSet() {
		toSerialize["last"] = o.Last.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginatedCollectionResponseListLinks) UnmarshalJSON(bytes []byte) (err error) {
	varPaginatedCollectionResponseListLinks := _PaginatedCollectionResponseListLinks{}

	if err = json.Unmarshal(bytes, &varPaginatedCollectionResponseListLinks); err == nil {
		*o = PaginatedCollectionResponseListLinks(varPaginatedCollectionResponseListLinks)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "first")
		delete(additionalProperties, "previous")
		delete(additionalProperties, "next")
		delete(additionalProperties, "last")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginatedCollectionResponseListLinks struct {
	value *PaginatedCollectionResponseListLinks
	isSet bool
}

func (v NullablePaginatedCollectionResponseListLinks) Get() *PaginatedCollectionResponseListLinks {
	return v.value
}

func (v *NullablePaginatedCollectionResponseListLinks) Set(val *PaginatedCollectionResponseListLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedCollectionResponseListLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedCollectionResponseListLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedCollectionResponseListLinks(val *PaginatedCollectionResponseListLinks) *NullablePaginatedCollectionResponseListLinks {
	return &NullablePaginatedCollectionResponseListLinks{value: val, isSet: true}
}

func (v NullablePaginatedCollectionResponseListLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedCollectionResponseListLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

