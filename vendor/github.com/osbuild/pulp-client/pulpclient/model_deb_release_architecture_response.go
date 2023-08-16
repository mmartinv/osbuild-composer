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
	"time"
)

// checks if the DebReleaseArchitectureResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DebReleaseArchitectureResponse{}

// DebReleaseArchitectureResponse A Serializer for ReleaseArchitecture.
type DebReleaseArchitectureResponse struct {
	PulpHref *string `json:"pulp_href,omitempty"`
	// Timestamp of creation.
	PulpCreated *time.Time `json:"pulp_created,omitempty"`
	// Name of the architecture.
	Architecture string `json:"architecture"`
	// Name of the distribution.
	Distribution string `json:"distribution"`
	Codename string `json:"codename"`
	Suite string `json:"suite"`
	AdditionalProperties map[string]interface{}
}

type _DebReleaseArchitectureResponse DebReleaseArchitectureResponse

// NewDebReleaseArchitectureResponse instantiates a new DebReleaseArchitectureResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDebReleaseArchitectureResponse(architecture string, distribution string, codename string, suite string) *DebReleaseArchitectureResponse {
	this := DebReleaseArchitectureResponse{}
	this.Architecture = architecture
	this.Distribution = distribution
	this.Codename = codename
	this.Suite = suite
	return &this
}

// NewDebReleaseArchitectureResponseWithDefaults instantiates a new DebReleaseArchitectureResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDebReleaseArchitectureResponseWithDefaults() *DebReleaseArchitectureResponse {
	this := DebReleaseArchitectureResponse{}
	return &this
}

// GetPulpHref returns the PulpHref field value if set, zero value otherwise.
func (o *DebReleaseArchitectureResponse) GetPulpHref() string {
	if o == nil || IsNil(o.PulpHref) {
		var ret string
		return ret
	}
	return *o.PulpHref
}

// GetPulpHrefOk returns a tuple with the PulpHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebReleaseArchitectureResponse) GetPulpHrefOk() (*string, bool) {
	if o == nil || IsNil(o.PulpHref) {
		return nil, false
	}
	return o.PulpHref, true
}

// HasPulpHref returns a boolean if a field has been set.
func (o *DebReleaseArchitectureResponse) HasPulpHref() bool {
	if o != nil && !IsNil(o.PulpHref) {
		return true
	}

	return false
}

// SetPulpHref gets a reference to the given string and assigns it to the PulpHref field.
func (o *DebReleaseArchitectureResponse) SetPulpHref(v string) {
	o.PulpHref = &v
}

// GetPulpCreated returns the PulpCreated field value if set, zero value otherwise.
func (o *DebReleaseArchitectureResponse) GetPulpCreated() time.Time {
	if o == nil || IsNil(o.PulpCreated) {
		var ret time.Time
		return ret
	}
	return *o.PulpCreated
}

// GetPulpCreatedOk returns a tuple with the PulpCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DebReleaseArchitectureResponse) GetPulpCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PulpCreated) {
		return nil, false
	}
	return o.PulpCreated, true
}

// HasPulpCreated returns a boolean if a field has been set.
func (o *DebReleaseArchitectureResponse) HasPulpCreated() bool {
	if o != nil && !IsNil(o.PulpCreated) {
		return true
	}

	return false
}

// SetPulpCreated gets a reference to the given time.Time and assigns it to the PulpCreated field.
func (o *DebReleaseArchitectureResponse) SetPulpCreated(v time.Time) {
	o.PulpCreated = &v
}

// GetArchitecture returns the Architecture field value
func (o *DebReleaseArchitectureResponse) GetArchitecture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Architecture
}

// GetArchitectureOk returns a tuple with the Architecture field value
// and a boolean to check if the value has been set.
func (o *DebReleaseArchitectureResponse) GetArchitectureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Architecture, true
}

// SetArchitecture sets field value
func (o *DebReleaseArchitectureResponse) SetArchitecture(v string) {
	o.Architecture = v
}

// GetDistribution returns the Distribution field value
func (o *DebReleaseArchitectureResponse) GetDistribution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Distribution
}

// GetDistributionOk returns a tuple with the Distribution field value
// and a boolean to check if the value has been set.
func (o *DebReleaseArchitectureResponse) GetDistributionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Distribution, true
}

// SetDistribution sets field value
func (o *DebReleaseArchitectureResponse) SetDistribution(v string) {
	o.Distribution = v
}

// GetCodename returns the Codename field value
func (o *DebReleaseArchitectureResponse) GetCodename() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Codename
}

// GetCodenameOk returns a tuple with the Codename field value
// and a boolean to check if the value has been set.
func (o *DebReleaseArchitectureResponse) GetCodenameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Codename, true
}

// SetCodename sets field value
func (o *DebReleaseArchitectureResponse) SetCodename(v string) {
	o.Codename = v
}

// GetSuite returns the Suite field value
func (o *DebReleaseArchitectureResponse) GetSuite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Suite
}

// GetSuiteOk returns a tuple with the Suite field value
// and a boolean to check if the value has been set.
func (o *DebReleaseArchitectureResponse) GetSuiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suite, true
}

// SetSuite sets field value
func (o *DebReleaseArchitectureResponse) SetSuite(v string) {
	o.Suite = v
}

func (o DebReleaseArchitectureResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DebReleaseArchitectureResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PulpHref) {
		toSerialize["pulp_href"] = o.PulpHref
	}
	if !IsNil(o.PulpCreated) {
		toSerialize["pulp_created"] = o.PulpCreated
	}
	toSerialize["architecture"] = o.Architecture
	toSerialize["distribution"] = o.Distribution
	toSerialize["codename"] = o.Codename
	toSerialize["suite"] = o.Suite

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DebReleaseArchitectureResponse) UnmarshalJSON(bytes []byte) (err error) {
	varDebReleaseArchitectureResponse := _DebReleaseArchitectureResponse{}

	if err = json.Unmarshal(bytes, &varDebReleaseArchitectureResponse); err == nil {
		*o = DebReleaseArchitectureResponse(varDebReleaseArchitectureResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pulp_href")
		delete(additionalProperties, "pulp_created")
		delete(additionalProperties, "architecture")
		delete(additionalProperties, "distribution")
		delete(additionalProperties, "codename")
		delete(additionalProperties, "suite")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDebReleaseArchitectureResponse struct {
	value *DebReleaseArchitectureResponse
	isSet bool
}

func (v NullableDebReleaseArchitectureResponse) Get() *DebReleaseArchitectureResponse {
	return v.value
}

func (v *NullableDebReleaseArchitectureResponse) Set(val *DebReleaseArchitectureResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDebReleaseArchitectureResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDebReleaseArchitectureResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDebReleaseArchitectureResponse(val *DebReleaseArchitectureResponse) *NullableDebReleaseArchitectureResponse {
	return &NullableDebReleaseArchitectureResponse{value: val, isSet: true}
}

func (v NullableDebReleaseArchitectureResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDebReleaseArchitectureResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


