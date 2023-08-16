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

// checks if the PatchedRole type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedRole{}

// PatchedRole Serializer for Role.
type PatchedRole struct {
	// The name of this role.
	Name *string `json:"name,omitempty"`
	// An optional description.
	Description NullableString `json:"description,omitempty"`
	// List of permissions defining the role.
	Permissions []string `json:"permissions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedRole PatchedRole

// NewPatchedRole instantiates a new PatchedRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedRole() *PatchedRole {
	this := PatchedRole{}
	return &this
}

// NewPatchedRoleWithDefaults instantiates a new PatchedRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedRoleWithDefaults() *PatchedRole {
	this := PatchedRole{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedRole) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRole) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedRole) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedRole) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchedRole) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchedRole) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedRole) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PatchedRole) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PatchedRole) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PatchedRole) UnsetDescription() {
	o.Description.Unset()
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *PatchedRole) GetPermissions() []string {
	if o == nil || IsNil(o.Permissions) {
		var ret []string
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedRole) GetPermissionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *PatchedRole) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []string and assigns it to the Permissions field.
func (o *PatchedRole) SetPermissions(v []string) {
	o.Permissions = v
}

func (o PatchedRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PatchedRole) UnmarshalJSON(bytes []byte) (err error) {
	varPatchedRole := _PatchedRole{}

	if err = json.Unmarshal(bytes, &varPatchedRole); err == nil {
		*o = PatchedRole(varPatchedRole)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "permissions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedRole struct {
	value *PatchedRole
	isSet bool
}

func (v NullablePatchedRole) Get() *PatchedRole {
	return v.value
}

func (v *NullablePatchedRole) Set(val *PatchedRole) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedRole) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedRole(val *PatchedRole) *NullablePatchedRole {
	return &NullablePatchedRole{value: val, isSet: true}
}

func (v NullablePatchedRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


