/*
Cisco Intersight

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.11-2025062323
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// checks if the IamUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IamUser{}

// IamUser The Intersight account user.
type IamUser struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// IP address from which the user last logged in to Intersight.
	ClientIpAddress *string `json:"ClientIpAddress,omitempty"`
	// Email of the user. Remote users are added to Intersight using the email configured in the IdP.
	Email *string "json:\"Email,omitempty\" validate:\"regexp=^$|^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$\""
	// First name of the user. For remote users, this field is populated from the IdP attributes received after authentication.
	FirstName *string `json:"FirstName,omitempty"`
	// Last successful login time for user.
	LastLoginTime *time.Time `json:"LastLoginTime,omitempty"`
	// Last name of the user. For remote users, this field is populated from the IdP attributes received after authentication.
	LastName *string `json:"LastName,omitempty"`
	// Last role modification time for user.
	LastRoleModifiedTime *time.Time `json:"LastRoleModifiedTime,omitempty"`
	// Time until which the user account will be locked out.
	LockedUntil *time.Time `json:"LockedUntil,omitempty"`
	// Name of the user. For remote users, it is the value as configured in the IdP.
	Name *string `json:"Name,omitempty"`
	// UserID or email of the user. For remote users, it is the value as configured in the IDP.
	UserIdOrEmail *string `json:"UserIdOrEmail,omitempty"`
	// Type of the User. If a user is added manually by specifying the email address, or has logged in using groups, based on the IdP attributes received during authentication. If added manually, the user type will be static, otherwise dynamic.
	UserType *string `json:"UserType,omitempty"`
	// Unique id of the user used by the identity provider to store the user.
	UserUniqueIdentifier *string `json:"UserUniqueIdentifier,omitempty"`
	// An array of relationships to iamApiKey resources.
	ApiKeys []IamApiKeyRelationship `json:"ApiKeys,omitempty"`
	// An array of relationships to iamAppRegistration resources.
	AppRegistrations  []IamAppRegistrationRelationship         `json:"AppRegistrations,omitempty"`
	Idp               NullableIamIdpRelationship               `json:"Idp,omitempty"`
	Idpreference      NullableIamIdpReferenceRelationship      `json:"Idpreference,omitempty"`
	LocalUserPassword NullableIamLocalUserPasswordRelationship `json:"LocalUserPassword,omitempty"`
	// An array of relationships to iamOAuthToken resources.
	OauthTokens []IamOAuthTokenRelationship `json:"OauthTokens,omitempty"`
	// An array of relationships to iamPermission resources.
	Permissions []IamPermissionRelationship `json:"Permissions,omitempty"`
	// An array of relationships to iamSession resources.
	Sessions             []IamSessionRelationship `json:"Sessions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamUser IamUser

// NewIamUser instantiates a new IamUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamUser(classId string, objectType string) *IamUser {
	this := IamUser{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamUserWithDefaults instantiates a new IamUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamUserWithDefaults() *IamUser {
	this := IamUser{}
	var classId string = "iam.User"
	this.ClassId = classId
	var objectType string = "iam.User"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamUser) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamUser) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamUser) SetClassId(v string) {
	o.ClassId = v
}

// GetDefaultClassId returns the default value "iam.User" of the ClassId field.
func (o *IamUser) GetDefaultClassId() interface{} {
	return "iam.User"
}

// GetObjectType returns the ObjectType field value
func (o *IamUser) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamUser) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamUser) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDefaultObjectType returns the default value "iam.User" of the ObjectType field.
func (o *IamUser) GetDefaultObjectType() interface{} {
	return "iam.User"
}

// GetClientIpAddress returns the ClientIpAddress field value if set, zero value otherwise.
func (o *IamUser) GetClientIpAddress() string {
	if o == nil || IsNil(o.ClientIpAddress) {
		var ret string
		return ret
	}
	return *o.ClientIpAddress
}

// GetClientIpAddressOk returns a tuple with the ClientIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUser) GetClientIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ClientIpAddress) {
		return nil, false
	}
	return o.ClientIpAddress, true
}

// HasClientIpAddress returns a boolean if a field has been set.
func (o *IamUser) HasClientIpAddress() bool {
	if o != nil && !IsNil(o.ClientIpAddress) {
		return true
	}

	return false
}

// SetClientIpAddress gets a reference to the given string and assigns it to the ClientIpAddress field.
func (o *IamUser) SetClientIpAddress(v string) {
	o.ClientIpAddress = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *IamUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *IamUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *IamUser) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *IamUser) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUser) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *IamUser) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *IamUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastLoginTime returns the LastLoginTime field value if set, zero value otherwise.
func (o *IamUser) GetLastLoginTime() time.Time {
	if o == nil || IsNil(o.LastLoginTime) {
		var ret time.Time
		return ret
	}
	return *o.LastLoginTime
}

// GetLastLoginTimeOk returns a tuple with the LastLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUser) GetLastLoginTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastLoginTime) {
		return nil, false
	}
	return o.LastLoginTime, true
}

// HasLastLoginTime returns a boolean if a field has been set.
func (o *IamUser) HasLastLoginTime() bool {
	if o != nil && !IsNil(o.LastLoginTime) {
		return true
	}

	return false
}

// SetLastLoginTime gets a reference to the given time.Time and assigns it to the LastLoginTime field.
func (o *IamUser) SetLastLoginTime(v time.Time) {
	o.LastLoginTime = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *IamUser) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUser) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *IamUser) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *IamUser) SetLastName(v string) {
	o.LastName = &v
}

// GetLastRoleModifiedTime returns the LastRoleModifiedTime field value if set, zero value otherwise.
func (o *IamUser) GetLastRoleModifiedTime() time.Time {
	if o == nil || IsNil(o.LastRoleModifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastRoleModifiedTime
}

// GetLastRoleModifiedTimeOk returns a tuple with the LastRoleModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUser) GetLastRoleModifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastRoleModifiedTime) {
		return nil, false
	}
	return o.LastRoleModifiedTime, true
}

// HasLastRoleModifiedTime returns a boolean if a field has been set.
func (o *IamUser) HasLastRoleModifiedTime() bool {
	if o != nil && !IsNil(o.LastRoleModifiedTime) {
		return true
	}

	return false
}

// SetLastRoleModifiedTime gets a reference to the given time.Time and assigns it to the LastRoleModifiedTime field.
func (o *IamUser) SetLastRoleModifiedTime(v time.Time) {
	o.LastRoleModifiedTime = &v
}

// GetLockedUntil returns the LockedUntil field value if set, zero value otherwise.
func (o *IamUser) GetLockedUntil() time.Time {
	if o == nil || IsNil(o.LockedUntil) {
		var ret time.Time
		return ret
	}
	return *o.LockedUntil
}

// GetLockedUntilOk returns a tuple with the LockedUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUser) GetLockedUntilOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LockedUntil) {
		return nil, false
	}
	return o.LockedUntil, true
}

// HasLockedUntil returns a boolean if a field has been set.
func (o *IamUser) HasLockedUntil() bool {
	if o != nil && !IsNil(o.LockedUntil) {
		return true
	}

	return false
}

// SetLockedUntil gets a reference to the given time.Time and assigns it to the LockedUntil field.
func (o *IamUser) SetLockedUntil(v time.Time) {
	o.LockedUntil = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamUser) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUser) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamUser) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamUser) SetName(v string) {
	o.Name = &v
}

// GetUserIdOrEmail returns the UserIdOrEmail field value if set, zero value otherwise.
func (o *IamUser) GetUserIdOrEmail() string {
	if o == nil || IsNil(o.UserIdOrEmail) {
		var ret string
		return ret
	}
	return *o.UserIdOrEmail
}

// GetUserIdOrEmailOk returns a tuple with the UserIdOrEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUser) GetUserIdOrEmailOk() (*string, bool) {
	if o == nil || IsNil(o.UserIdOrEmail) {
		return nil, false
	}
	return o.UserIdOrEmail, true
}

// HasUserIdOrEmail returns a boolean if a field has been set.
func (o *IamUser) HasUserIdOrEmail() bool {
	if o != nil && !IsNil(o.UserIdOrEmail) {
		return true
	}

	return false
}

// SetUserIdOrEmail gets a reference to the given string and assigns it to the UserIdOrEmail field.
func (o *IamUser) SetUserIdOrEmail(v string) {
	o.UserIdOrEmail = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *IamUser) GetUserType() string {
	if o == nil || IsNil(o.UserType) {
		var ret string
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUser) GetUserTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UserType) {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *IamUser) HasUserType() bool {
	if o != nil && !IsNil(o.UserType) {
		return true
	}

	return false
}

// SetUserType gets a reference to the given string and assigns it to the UserType field.
func (o *IamUser) SetUserType(v string) {
	o.UserType = &v
}

// GetUserUniqueIdentifier returns the UserUniqueIdentifier field value if set, zero value otherwise.
func (o *IamUser) GetUserUniqueIdentifier() string {
	if o == nil || IsNil(o.UserUniqueIdentifier) {
		var ret string
		return ret
	}
	return *o.UserUniqueIdentifier
}

// GetUserUniqueIdentifierOk returns a tuple with the UserUniqueIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUser) GetUserUniqueIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.UserUniqueIdentifier) {
		return nil, false
	}
	return o.UserUniqueIdentifier, true
}

// HasUserUniqueIdentifier returns a boolean if a field has been set.
func (o *IamUser) HasUserUniqueIdentifier() bool {
	if o != nil && !IsNil(o.UserUniqueIdentifier) {
		return true
	}

	return false
}

// SetUserUniqueIdentifier gets a reference to the given string and assigns it to the UserUniqueIdentifier field.
func (o *IamUser) SetUserUniqueIdentifier(v string) {
	o.UserUniqueIdentifier = &v
}

// GetApiKeys returns the ApiKeys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUser) GetApiKeys() []IamApiKeyRelationship {
	if o == nil {
		var ret []IamApiKeyRelationship
		return ret
	}
	return o.ApiKeys
}

// GetApiKeysOk returns a tuple with the ApiKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUser) GetApiKeysOk() ([]IamApiKeyRelationship, bool) {
	if o == nil || IsNil(o.ApiKeys) {
		return nil, false
	}
	return o.ApiKeys, true
}

// HasApiKeys returns a boolean if a field has been set.
func (o *IamUser) HasApiKeys() bool {
	if o != nil && !IsNil(o.ApiKeys) {
		return true
	}

	return false
}

// SetApiKeys gets a reference to the given []IamApiKeyRelationship and assigns it to the ApiKeys field.
func (o *IamUser) SetApiKeys(v []IamApiKeyRelationship) {
	o.ApiKeys = v
}

// GetAppRegistrations returns the AppRegistrations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUser) GetAppRegistrations() []IamAppRegistrationRelationship {
	if o == nil {
		var ret []IamAppRegistrationRelationship
		return ret
	}
	return o.AppRegistrations
}

// GetAppRegistrationsOk returns a tuple with the AppRegistrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUser) GetAppRegistrationsOk() ([]IamAppRegistrationRelationship, bool) {
	if o == nil || IsNil(o.AppRegistrations) {
		return nil, false
	}
	return o.AppRegistrations, true
}

// HasAppRegistrations returns a boolean if a field has been set.
func (o *IamUser) HasAppRegistrations() bool {
	if o != nil && !IsNil(o.AppRegistrations) {
		return true
	}

	return false
}

// SetAppRegistrations gets a reference to the given []IamAppRegistrationRelationship and assigns it to the AppRegistrations field.
func (o *IamUser) SetAppRegistrations(v []IamAppRegistrationRelationship) {
	o.AppRegistrations = v
}

// GetIdp returns the Idp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUser) GetIdp() IamIdpRelationship {
	if o == nil || IsNil(o.Idp.Get()) {
		var ret IamIdpRelationship
		return ret
	}
	return *o.Idp.Get()
}

// GetIdpOk returns a tuple with the Idp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUser) GetIdpOk() (*IamIdpRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Idp.Get(), o.Idp.IsSet()
}

// HasIdp returns a boolean if a field has been set.
func (o *IamUser) HasIdp() bool {
	if o != nil && o.Idp.IsSet() {
		return true
	}

	return false
}

// SetIdp gets a reference to the given NullableIamIdpRelationship and assigns it to the Idp field.
func (o *IamUser) SetIdp(v IamIdpRelationship) {
	o.Idp.Set(&v)
}

// SetIdpNil sets the value for Idp to be an explicit nil
func (o *IamUser) SetIdpNil() {
	o.Idp.Set(nil)
}

// UnsetIdp ensures that no value is present for Idp, not even an explicit nil
func (o *IamUser) UnsetIdp() {
	o.Idp.Unset()
}

// GetIdpreference returns the Idpreference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUser) GetIdpreference() IamIdpReferenceRelationship {
	if o == nil || IsNil(o.Idpreference.Get()) {
		var ret IamIdpReferenceRelationship
		return ret
	}
	return *o.Idpreference.Get()
}

// GetIdpreferenceOk returns a tuple with the Idpreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUser) GetIdpreferenceOk() (*IamIdpReferenceRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Idpreference.Get(), o.Idpreference.IsSet()
}

// HasIdpreference returns a boolean if a field has been set.
func (o *IamUser) HasIdpreference() bool {
	if o != nil && o.Idpreference.IsSet() {
		return true
	}

	return false
}

// SetIdpreference gets a reference to the given NullableIamIdpReferenceRelationship and assigns it to the Idpreference field.
func (o *IamUser) SetIdpreference(v IamIdpReferenceRelationship) {
	o.Idpreference.Set(&v)
}

// SetIdpreferenceNil sets the value for Idpreference to be an explicit nil
func (o *IamUser) SetIdpreferenceNil() {
	o.Idpreference.Set(nil)
}

// UnsetIdpreference ensures that no value is present for Idpreference, not even an explicit nil
func (o *IamUser) UnsetIdpreference() {
	o.Idpreference.Unset()
}

// GetLocalUserPassword returns the LocalUserPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUser) GetLocalUserPassword() IamLocalUserPasswordRelationship {
	if o == nil || IsNil(o.LocalUserPassword.Get()) {
		var ret IamLocalUserPasswordRelationship
		return ret
	}
	return *o.LocalUserPassword.Get()
}

// GetLocalUserPasswordOk returns a tuple with the LocalUserPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUser) GetLocalUserPasswordOk() (*IamLocalUserPasswordRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalUserPassword.Get(), o.LocalUserPassword.IsSet()
}

// HasLocalUserPassword returns a boolean if a field has been set.
func (o *IamUser) HasLocalUserPassword() bool {
	if o != nil && o.LocalUserPassword.IsSet() {
		return true
	}

	return false
}

// SetLocalUserPassword gets a reference to the given NullableIamLocalUserPasswordRelationship and assigns it to the LocalUserPassword field.
func (o *IamUser) SetLocalUserPassword(v IamLocalUserPasswordRelationship) {
	o.LocalUserPassword.Set(&v)
}

// SetLocalUserPasswordNil sets the value for LocalUserPassword to be an explicit nil
func (o *IamUser) SetLocalUserPasswordNil() {
	o.LocalUserPassword.Set(nil)
}

// UnsetLocalUserPassword ensures that no value is present for LocalUserPassword, not even an explicit nil
func (o *IamUser) UnsetLocalUserPassword() {
	o.LocalUserPassword.Unset()
}

// GetOauthTokens returns the OauthTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUser) GetOauthTokens() []IamOAuthTokenRelationship {
	if o == nil {
		var ret []IamOAuthTokenRelationship
		return ret
	}
	return o.OauthTokens
}

// GetOauthTokensOk returns a tuple with the OauthTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUser) GetOauthTokensOk() ([]IamOAuthTokenRelationship, bool) {
	if o == nil || IsNil(o.OauthTokens) {
		return nil, false
	}
	return o.OauthTokens, true
}

// HasOauthTokens returns a boolean if a field has been set.
func (o *IamUser) HasOauthTokens() bool {
	if o != nil && !IsNil(o.OauthTokens) {
		return true
	}

	return false
}

// SetOauthTokens gets a reference to the given []IamOAuthTokenRelationship and assigns it to the OauthTokens field.
func (o *IamUser) SetOauthTokens(v []IamOAuthTokenRelationship) {
	o.OauthTokens = v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUser) GetPermissions() []IamPermissionRelationship {
	if o == nil {
		var ret []IamPermissionRelationship
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUser) GetPermissionsOk() ([]IamPermissionRelationship, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *IamUser) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []IamPermissionRelationship and assigns it to the Permissions field.
func (o *IamUser) SetPermissions(v []IamPermissionRelationship) {
	o.Permissions = v
}

// GetSessions returns the Sessions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUser) GetSessions() []IamSessionRelationship {
	if o == nil {
		var ret []IamSessionRelationship
		return ret
	}
	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUser) GetSessionsOk() ([]IamSessionRelationship, bool) {
	if o == nil || IsNil(o.Sessions) {
		return nil, false
	}
	return o.Sessions, true
}

// HasSessions returns a boolean if a field has been set.
func (o *IamUser) HasSessions() bool {
	if o != nil && !IsNil(o.Sessions) {
		return true
	}

	return false
}

// SetSessions gets a reference to the given []IamSessionRelationship and assigns it to the Sessions field.
func (o *IamUser) SetSessions(v []IamSessionRelationship) {
	o.Sessions = v
}

func (o IamUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IamUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return map[string]interface{}{}, errMoBaseMo
	}
	if _, exists := toSerialize["ClassId"]; !exists {
		toSerialize["ClassId"] = o.GetDefaultClassId()
	}
	toSerialize["ClassId"] = o.ClassId
	if _, exists := toSerialize["ObjectType"]; !exists {
		toSerialize["ObjectType"] = o.GetDefaultObjectType()
	}
	toSerialize["ObjectType"] = o.ObjectType
	if !IsNil(o.ClientIpAddress) {
		toSerialize["ClientIpAddress"] = o.ClientIpAddress
	}
	if !IsNil(o.Email) {
		toSerialize["Email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["FirstName"] = o.FirstName
	}
	if !IsNil(o.LastLoginTime) {
		toSerialize["LastLoginTime"] = o.LastLoginTime
	}
	if !IsNil(o.LastName) {
		toSerialize["LastName"] = o.LastName
	}
	if !IsNil(o.LastRoleModifiedTime) {
		toSerialize["LastRoleModifiedTime"] = o.LastRoleModifiedTime
	}
	if !IsNil(o.LockedUntil) {
		toSerialize["LockedUntil"] = o.LockedUntil
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.UserIdOrEmail) {
		toSerialize["UserIdOrEmail"] = o.UserIdOrEmail
	}
	if !IsNil(o.UserType) {
		toSerialize["UserType"] = o.UserType
	}
	if !IsNil(o.UserUniqueIdentifier) {
		toSerialize["UserUniqueIdentifier"] = o.UserUniqueIdentifier
	}
	if o.ApiKeys != nil {
		toSerialize["ApiKeys"] = o.ApiKeys
	}
	if o.AppRegistrations != nil {
		toSerialize["AppRegistrations"] = o.AppRegistrations
	}
	if o.Idp.IsSet() {
		toSerialize["Idp"] = o.Idp.Get()
	}
	if o.Idpreference.IsSet() {
		toSerialize["Idpreference"] = o.Idpreference.Get()
	}
	if o.LocalUserPassword.IsSet() {
		toSerialize["LocalUserPassword"] = o.LocalUserPassword.Get()
	}
	if o.OauthTokens != nil {
		toSerialize["OauthTokens"] = o.OauthTokens
	}
	if o.Permissions != nil {
		toSerialize["Permissions"] = o.Permissions
	}
	if o.Sessions != nil {
		toSerialize["Sessions"] = o.Sessions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IamUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ClassId",
		"ObjectType",
	}

	// defaultValueFuncMap captures the default values for required properties.
	// These values are used when required properties are missing from the payload.
	defaultValueFuncMap := map[string]func() interface{}{
		"ClassId":    o.GetDefaultClassId,
		"ObjectType": o.GetDefaultObjectType,
	}
	var defaultValueApplied bool
	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			if _, ok := defaultValueFuncMap[requiredProperty]; ok {
				allProperties[requiredProperty] = defaultValueFuncMap[requiredProperty]()
				defaultValueApplied = true
			}
		}
		if value, exists := allProperties[requiredProperty]; !exists || value == "" {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	if defaultValueApplied {
		data, err = json.Marshal(allProperties)
		if err != nil {
			return err
		}
	}
	type IamUserWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// IP address from which the user last logged in to Intersight.
		ClientIpAddress *string `json:"ClientIpAddress,omitempty"`
		// Email of the user. Remote users are added to Intersight using the email configured in the IdP.
		Email *string "json:\"Email,omitempty\" validate:\"regexp=^$|^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$\""
		// First name of the user. For remote users, this field is populated from the IdP attributes received after authentication.
		FirstName *string `json:"FirstName,omitempty"`
		// Last successful login time for user.
		LastLoginTime *time.Time `json:"LastLoginTime,omitempty"`
		// Last name of the user. For remote users, this field is populated from the IdP attributes received after authentication.
		LastName *string `json:"LastName,omitempty"`
		// Last role modification time for user.
		LastRoleModifiedTime *time.Time `json:"LastRoleModifiedTime,omitempty"`
		// Time until which the user account will be locked out.
		LockedUntil *time.Time `json:"LockedUntil,omitempty"`
		// Name of the user. For remote users, it is the value as configured in the IdP.
		Name *string `json:"Name,omitempty"`
		// UserID or email of the user. For remote users, it is the value as configured in the IDP.
		UserIdOrEmail *string `json:"UserIdOrEmail,omitempty"`
		// Type of the User. If a user is added manually by specifying the email address, or has logged in using groups, based on the IdP attributes received during authentication. If added manually, the user type will be static, otherwise dynamic.
		UserType *string `json:"UserType,omitempty"`
		// Unique id of the user used by the identity provider to store the user.
		UserUniqueIdentifier *string `json:"UserUniqueIdentifier,omitempty"`
		// An array of relationships to iamApiKey resources.
		ApiKeys []IamApiKeyRelationship `json:"ApiKeys,omitempty"`
		// An array of relationships to iamAppRegistration resources.
		AppRegistrations  []IamAppRegistrationRelationship         `json:"AppRegistrations,omitempty"`
		Idp               NullableIamIdpRelationship               `json:"Idp,omitempty"`
		Idpreference      NullableIamIdpReferenceRelationship      `json:"Idpreference,omitempty"`
		LocalUserPassword NullableIamLocalUserPasswordRelationship `json:"LocalUserPassword,omitempty"`
		// An array of relationships to iamOAuthToken resources.
		OauthTokens []IamOAuthTokenRelationship `json:"OauthTokens,omitempty"`
		// An array of relationships to iamPermission resources.
		Permissions []IamPermissionRelationship `json:"Permissions,omitempty"`
		// An array of relationships to iamSession resources.
		Sessions []IamSessionRelationship `json:"Sessions,omitempty"`
	}

	varIamUserWithoutEmbeddedStruct := IamUserWithoutEmbeddedStruct{}

	err = json.Unmarshal(data, &varIamUserWithoutEmbeddedStruct)
	if err == nil {
		varIamUser := _IamUser{}
		varIamUser.ClassId = varIamUserWithoutEmbeddedStruct.ClassId
		varIamUser.ObjectType = varIamUserWithoutEmbeddedStruct.ObjectType
		varIamUser.ClientIpAddress = varIamUserWithoutEmbeddedStruct.ClientIpAddress
		varIamUser.Email = varIamUserWithoutEmbeddedStruct.Email
		varIamUser.FirstName = varIamUserWithoutEmbeddedStruct.FirstName
		varIamUser.LastLoginTime = varIamUserWithoutEmbeddedStruct.LastLoginTime
		varIamUser.LastName = varIamUserWithoutEmbeddedStruct.LastName
		varIamUser.LastRoleModifiedTime = varIamUserWithoutEmbeddedStruct.LastRoleModifiedTime
		varIamUser.LockedUntil = varIamUserWithoutEmbeddedStruct.LockedUntil
		varIamUser.Name = varIamUserWithoutEmbeddedStruct.Name
		varIamUser.UserIdOrEmail = varIamUserWithoutEmbeddedStruct.UserIdOrEmail
		varIamUser.UserType = varIamUserWithoutEmbeddedStruct.UserType
		varIamUser.UserUniqueIdentifier = varIamUserWithoutEmbeddedStruct.UserUniqueIdentifier
		varIamUser.ApiKeys = varIamUserWithoutEmbeddedStruct.ApiKeys
		varIamUser.AppRegistrations = varIamUserWithoutEmbeddedStruct.AppRegistrations
		varIamUser.Idp = varIamUserWithoutEmbeddedStruct.Idp
		varIamUser.Idpreference = varIamUserWithoutEmbeddedStruct.Idpreference
		varIamUser.LocalUserPassword = varIamUserWithoutEmbeddedStruct.LocalUserPassword
		varIamUser.OauthTokens = varIamUserWithoutEmbeddedStruct.OauthTokens
		varIamUser.Permissions = varIamUserWithoutEmbeddedStruct.Permissions
		varIamUser.Sessions = varIamUserWithoutEmbeddedStruct.Sessions
		*o = IamUser(varIamUser)
	} else {
		return err
	}

	varIamUser := _IamUser{}

	err = json.Unmarshal(data, &varIamUser)
	if err == nil {
		o.MoBaseMo = varIamUser.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ClientIpAddress")
		delete(additionalProperties, "Email")
		delete(additionalProperties, "FirstName")
		delete(additionalProperties, "LastLoginTime")
		delete(additionalProperties, "LastName")
		delete(additionalProperties, "LastRoleModifiedTime")
		delete(additionalProperties, "LockedUntil")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "UserIdOrEmail")
		delete(additionalProperties, "UserType")
		delete(additionalProperties, "UserUniqueIdentifier")
		delete(additionalProperties, "ApiKeys")
		delete(additionalProperties, "AppRegistrations")
		delete(additionalProperties, "Idp")
		delete(additionalProperties, "Idpreference")
		delete(additionalProperties, "LocalUserPassword")
		delete(additionalProperties, "OauthTokens")
		delete(additionalProperties, "Permissions")
		delete(additionalProperties, "Sessions")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamUser struct {
	value *IamUser
	isSet bool
}

func (v NullableIamUser) Get() *IamUser {
	return v.value
}

func (v *NullableIamUser) Set(val *IamUser) {
	v.value = val
	v.isSet = true
}

func (v NullableIamUser) IsSet() bool {
	return v.isSet
}

func (v *NullableIamUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamUser(val *IamUser) *NullableIamUser {
	return &NullableIamUser{value: val, isSet: true}
}

func (v NullableIamUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
