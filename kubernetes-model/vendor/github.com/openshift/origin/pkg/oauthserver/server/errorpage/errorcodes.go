/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package errorpage

import "github.com/openshift/origin/pkg/oauthserver/userregistry/identitymapper"

const (
	// error occurred attempting to claim a user
	errorCodeClaim = "mapping_claim_error"
	// error occurred looking up the user
	errorCodeLookup = "mapping_lookup_error"
	// general authentication error
	errorCodeAuthentication = "authentication_error"
	// general grant error
	errorCodeGrant = "grant_error"
)

// AuthenticationErrorCode returns an error code for the given authentication error.
// If the error is not recognized, a generic error code is returned.
func AuthenticationErrorCode(err error) string {
	switch {
	case identitymapper.IsClaimError(err):
		return errorCodeClaim
	case identitymapper.IsLookupError(err):
		return errorCodeLookup
	default:
		return errorCodeAuthentication
	}
}

// AuthenticationErrorMessage returns an error message for the given authentication error code.
// If the error code is not recognized, a generic error message is returned.
func AuthenticationErrorMessage(code string) string {
	switch code {
	case errorCodeClaim:
		return "Could not create user."
	case errorCodeLookup:
		return "Could not find user."
	default:
		return "An authentication error occurred."
	}
}

// GrantErrorCode returns an error code for the given grant error.
// If the error is not recognized, a generic error code is returned.
func GrantErrorCode(err error) string {
	return errorCodeGrant
}

// GrantErrorMessage returns an error message for the given grant error code.
// If the error is not recognized, a generic error message is returned.
func GrantErrorMessage(code string) string {
	return "A grant error occurred."
}
