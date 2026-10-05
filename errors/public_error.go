package errors

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type SystemCode string

const (
	EARInternalError    SystemCode = "EAR-001"
	EARInvalidInput     SystemCode = "EAR-002"
	EARUserNotFound     SystemCode = "EAR-003" // Legacy / Generic
	EARDuplicateUser    SystemCode = "EAR-004" // Legacy / Generic
	EARUnauthenticated  SystemCode = "EAR-005"
	EARPermissionDenied SystemCode = "EAR-006"
	EARFileNotFound     SystemCode = "EAR-007"

	// ==========================================
	// MERCHANT SPECIFIC ERRORS (EAR-008 to EAR-019)
	// ==========================================
	EARMerchantInvalidInput             SystemCode = "EAR-008"
	EARMerchantNotFoundByPhone          SystemCode = "EAR-009"
	EARMerchantNotFoundByUsername       SystemCode = "EAR-010"
	EARMerchantPhoneExistsVerified      SystemCode = "EAR-011"
	EARMerchantPhoneExistsUnverified    SystemCode = "EAR-012"
	EARMerchantUsernameExistsVerified   SystemCode = "EAR-013"
	EARMerchantUsernameExistsUnverified SystemCode = "EAR-014"
	EARMerchantUnauthenticated          SystemCode = "EAR-015"
	EARMerchantPermissionDenied         SystemCode = "EAR-016"
	EARMerchantLookupFailedByPhone      SystemCode = "EAR-017"
	EARMerchantLookupFailedByUsername   SystemCode = "EAR-018"
	EARMerchantListRetrievalFailed      SystemCode = "EAR-019"

	// ==========================================
	// USER SPECIFIC ERRORS (EAR-020 to EAR-031)
	// ==========================================
	EARUserInvalidInput SystemCode = "EAR-020"

	// Specific User Not Found Scenarios
	EARUserNotFoundByPhone    SystemCode = "EAR-021"
	EARUserNotFoundByUsername SystemCode = "EAR-022"

	// Specific User Conflict Scenarios
	EARUserPhoneExistsVerified      SystemCode = "EAR-023"
	EARUserPhoneExistsUnverified    SystemCode = "EAR-024"
	EARUserUsernameExistsVerified   SystemCode = "EAR-025"
	EARUserUsernameExistsUnverified SystemCode = "EAR-026"

	EARUserUnauthenticated  SystemCode = "EAR-027"
	EARUserPermissionDenied SystemCode = "EAR-028"

	// User Database Lookup Failures (HTTP 500)
	EARUserLookupFailedByPhone    SystemCode = "EAR-029"
	EARUserLookupFailedByUsername SystemCode = "EAR-030"
	EARUserListRetrievalFailed    SystemCode = "EAR-031"

	// ==========================================
	// TRANSACTION SPECIFIC ERRORS (EAR-032 onwards)
	// ==========================================
	EARTxMerchantAccountNotFound     SystemCode = "EAR-032"
	EARTxUserAccountNotFound         SystemCode = "EAR-033"
	EARTxInsufficientBalance         SystemCode = "EAR-034"
	EARTxAmountMustBeGreaterThanZero SystemCode = "EAR-035"
	EARTxInvalidAuthentication       SystemCode = "EAR-036"

	// ==========================================
	// VALIDATION & EMBEDDING SPECIFIC ERRORS (EAR-037 onwards)
	// ==========================================
	EAREmbeddingInvalidBase64         SystemCode = "EAR-037"
	EAREmbeddingInvalidLength         SystemCode = "EAR-038"
	EAREmbeddingInvalidFloatData      SystemCode = "EAR-039"
	EAREmbeddingNotNormalized         SystemCode = "EAR-040"
	EARValidationInvalidPhone         SystemCode = "EAR-041"
	EARValidationInvalidId            SystemCode = "EAR-042"
	EARValidationInvalidOtpLength     SystemCode = "EAR-043"
	EARValidationInvalidName          SystemCode = "EAR-044"
	EARValidationInvalidPassword      SystemCode = "EAR-045"
	EARValidationInvalidDateFormat    SystemCode = "EAR-046"
	EARValidationFutureDateNotAllowed SystemCode = "EAR-047"
	// ==========================================
	// OTP & SMS DELIVERY SPECIFIC ERRORS (EAR-048 onwards)
	// ==========================================
	EAROtpDeliveryFailed          SystemCode = "EAR-048"
	EAROtpVerificationSystemError SystemCode = "EAR-049"
	EAROtpIncorrectCode           SystemCode = "EAR-050"
	// ==========================================
	// AUTHENTICATION CONTEXT SPECIFIC ERRORS (EAR-051 onwards)
	// ==========================================
	EARAuthTokenInvalid   SystemCode = "EAR-051"
	EARAuthContextMissing SystemCode = "EAR-052"
	// ==========================================
	// JWT LIFECYCLE SPECIFIC ERRORS (EAR-053 onwards)
	// ==========================================
	EARJwtSigningFailed        SystemCode = "EAR-053"
	EARJwtParsingFailed        SystemCode = "EAR-054"
	EARJwtConfigurationMissing SystemCode = "EAR-055"
	// ==========================================
	// PASSWORD SPECIFIC ERRORS (EAR-056 onwards)
	// ==========================================
	EARAuthPasswordMismatch      SystemCode = "EAR-056"
	EARAuthPasswordHashingFailed SystemCode = "EAR-057"
	// ==========================================
	// Offline Transaction ID Validation Scenarios
	// ==========================================
	EARValidationInvalidOfflineTxId SystemCode = "EAR-058"
	EARValidationInvalidDeviceID        SystemCode = "EAR-059"
	EARValidationInvalidUnixTimestamp   SystemCode = "EAR-060"
)

type errorDefinition struct {
	HttpStatus int
	Message    string
}

var registry = map[SystemCode]errorDefinition{
	EARInternalError: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "An unexpected internal server error occurred. Please try again later.",
	},
	EARInvalidInput: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The provided input parameters are invalid or malformed.",
	},
	EARUserNotFound: {
		HttpStatus: http.StatusNotFound,
		Message:    "The requested user record could not be found in the system.",
	},
	EARDuplicateUser: {
		HttpStatus: http.StatusConflict,
		Message:    "A user account with those details already exists.",
	},
	EARUnauthenticated: {
		HttpStatus: http.StatusUnauthorized,
		Message:    "Authentication is required to access this resource.",
	},
	EARPermissionDenied: {
		HttpStatus: http.StatusForbidden,
		Message:    "You do not have the required permissions to perform this action.",
	},
	EARFileNotFound: {
		HttpStatus: http.StatusNotFound,
		Message:    "The requested file or asset could not be found.",
	},

	// ------------------------------------------
	// Merchant Registrations
	// ------------------------------------------
	EARMerchantInvalidInput: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The provided merchant registration or input details are invalid.",
	},
	EARMerchantNotFoundByPhone: {
		HttpStatus: http.StatusNotFound,
		Message:    "No merchant account matches the provided phone number.",
	},
	EARMerchantNotFoundByUsername: {
		HttpStatus: http.StatusNotFound,
		Message:    "No merchant account matches the provided username.",
	},
	EARMerchantPhoneExistsVerified: {
		HttpStatus: http.StatusConflict,
		Message:    "This phone number is already registered to a verified merchant account.",
	},
	EARMerchantPhoneExistsUnverified: {
		HttpStatus: http.StatusConflict,
		Message:    "This phone number belongs to an unverified merchant. Please complete your verification process.",
	},
	EARMerchantUsernameExistsVerified: {
		HttpStatus: http.StatusConflict,
		Message:    "This username is already taken by a verified merchant.",
	},
	EARMerchantUsernameExistsUnverified: {
		HttpStatus: http.StatusConflict,
		Message:    "This username is reserved by an unverified merchant profile.",
	},
	EARMerchantUnauthenticated: {
		HttpStatus: http.StatusUnauthorized,
		Message:    "Merchant authentication failed or credentials are missing.",
	},
	EARMerchantPermissionDenied: {
		HttpStatus: http.StatusForbidden,
		Message:    "This merchant account does not have permission to access this resource.",
	},
	EARMerchantLookupFailedByPhone: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "The system encountered an error while scanning for the merchant phone number.",
	},
	EARMerchantLookupFailedByUsername: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "The system encountered an error while scanning for the merchant username.",
	},
	EARMerchantListRetrievalFailed: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "The system encountered an error while attempting to compile the merchant list.",
	},

	// ------------------------------------------
	// User Registrations
	// ------------------------------------------
	EARUserInvalidInput: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The provided user input or profile details are invalid.",
	},
	EARUserNotFoundByPhone: {
		HttpStatus: http.StatusNotFound,
		Message:    "No user account matches the provided phone number.",
	},
	EARUserNotFoundByUsername: {
		HttpStatus: http.StatusNotFound,
		Message:    "No user account matches the provided username.",
	},
	EARUserPhoneExistsVerified: {
		HttpStatus: http.StatusConflict,
		Message:    "This phone number is already registered to a verified user account.",
	},
	EARUserPhoneExistsUnverified: {
		HttpStatus: http.StatusConflict,
		Message:    "This phone number belongs to an unverified user. Please complete your verification process.",
	},
	EARUserUsernameExistsVerified: {
		HttpStatus: http.StatusConflict,
		Message:    "This username is already taken by a verified user.",
	},
	EARUserUsernameExistsUnverified: {
		HttpStatus: http.StatusConflict,
		Message:    "This username is reserved by an unverified user profile.",
	},
	EARUserUnauthenticated: {
		HttpStatus: http.StatusUnauthorized,
		Message:    "User authentication failed or credentials are missing.",
	},
	EARUserPermissionDenied: {
		HttpStatus: http.StatusForbidden,
		Message:    "This user account does not have permission to access this resource.",
	},
	EARUserLookupFailedByPhone: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "The system encountered an error while scanning for the user phone number.",
	},
	EARUserLookupFailedByUsername: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "The system encountered an error while scanning for the user username.",
	},
	EARUserListRetrievalFailed: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "The system encountered an error while attempting to compile the user list.",
	},

	// ------------------------------------------
	// Transactions
	// ------------------------------------------
	EARTxMerchantAccountNotFound: {
		HttpStatus: http.StatusNotFound,
		Message:    "The merchant account associated with this transaction could not be found.",
	},
	EARTxUserAccountNotFound: {
		HttpStatus: http.StatusNotFound,
		Message:    "The user account associated with this transaction could not be found.",
	},
	EARTxInsufficientBalance: {
		HttpStatus: http.StatusPaymentRequired,
		Message:    "Transaction declined. The account balance is insufficient to complete this payment.",
	},
	EARTxAmountMustBeGreaterThanZero: {
		HttpStatus: http.StatusBadRequest,
		Message:    "Invalid transaction details. The transaction amount must be greater than zero.",
	},
	EARTxInvalidAuthentication: {
		HttpStatus: http.StatusUnauthorized,
		Message:    "Transaction authentication failed or is invalid.",
	},
	// ------------------------------------------
	// Validation & Embeddings
	// ------------------------------------------
	EAREmbeddingInvalidBase64: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The provided face embedding is not a valid base64 encoded string.",
	},
	EAREmbeddingInvalidLength: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The face embedding size is incorrect. It must be exactly 2048 bytes (512 float32 values).",
	},
	EAREmbeddingInvalidFloatData: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The face embedding contains invalid data vectors (NaN or Infinity).",
	},
	EAREmbeddingNotNormalized: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The face embedding structure is invalid because it is not L2 normalized.",
	},
	EARValidationInvalidPhone: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The provided phone number format is invalid.",
	},
	EARValidationInvalidId: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The provided identifier must be greater than zero.",
	},
	EARValidationInvalidOtpLength: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The one-time password (OTP) must be exactly 6 digits.",
	},
	EARValidationInvalidName: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The name field cannot be empty.",
	},
	EARValidationInvalidPassword: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The password is too short. It must be at least 8 characters long.",
	},
	EARValidationInvalidDateFormat: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The provided date format is invalid. Please use YYYY-MM-DD.",
	},
	EARValidationFutureDateNotAllowed: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The date provided cannot be in the future.",
	},
	// ------------------------------------------
	// OTP & SMS Lifecycle
	// ------------------------------------------
	EAROtpDeliveryFailed: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "Failed to send the verification code. Please check the phone number format and try again.",
	},
	EAROtpVerificationSystemError: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "The system encountered an error while validating the verification code.",
	},
	EAROtpIncorrectCode: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The verification code entered is incorrect or has expired.",
	},
	// ------------------------------------------
	// Authentication Context
	// ------------------------------------------
	EARAuthTokenInvalid: {
		HttpStatus: http.StatusUnauthorized,
		Message:    "The provided authentication token is invalid or expired.",
	},
	EARAuthContextMissing: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "Authentication credentials could not be retrieved from the request context.",
	},
	// ------------------------------------------
	// JWT Lifecycle
	// ------------------------------------------
	EARJwtSigningFailed: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "The system encountered an error while signing the authentication token.",
	},
	EARJwtParsingFailed: {
		HttpStatus: http.StatusUnauthorized,
		Message:    "The provided authentication token is malformed, expired, or digitally invalid.",
	},
	EARJwtConfigurationMissing: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "Critical server configuration is missing. Token cryptographic keys are uninitialized.",
	},
	// ------------------------------------------
	// Password Verification
	// ------------------------------------------
	EARAuthPasswordMismatch: {
		HttpStatus: http.StatusUnauthorized,
		Message:    "The password provided does not match our records.",
	},
	// ------------------------------------------
	// Password Encryption
	// ------------------------------------------
	EARAuthPasswordHashingFailed: {
		HttpStatus: http.StatusInternalServerError,
		Message:    "The system encountered an unexpected error while securing the password.",
	},
    // ==========================================
	// Offline Transaction ID Validation Scenarios
	// ==========================================
    EARValidationInvalidOfflineTxId: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The offline transaction ID format is invalid. Expected format: 'IMEI_MillisecondTimestamp'.",
	},
	EARValidationInvalidDeviceID: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The device ID in the offline transaction ID is invalid.",
	},
	EARValidationInvalidUnixTimestamp: {
		HttpStatus: http.StatusBadRequest,
		Message:    "The timestamp in the offline transaction ID must be a valid 13-digit millisecond Unix timestamp.",
	},
}

// PublicError holds the structured details sent back to clients/GraphQL.
type PublicError struct {
	SystemCode SystemCode `json:"code"`
	HttpStatus int        `json:"http_status"`
	Message    string     `json:"message"`
	Err        error      `json:"-"`
}

func (e *PublicError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] (HTTP %d) %s (wrapped: %s)", e.SystemCode, e.HttpStatus, e.Message, e.Err.Error())
	}
	return fmt.Sprintf("[%s] (HTTP %d) %s", e.SystemCode, e.HttpStatus, e.Message)
}

func (e *PublicError) Unwrap() error {
	return e.Err
}

func New(code SystemCode, err error) *PublicError {
	defn, exists := registry[code]
	if !exists {
		defn = registry[EARInternalError]
		code = EARInternalError
	}

	return &PublicError{
		SystemCode: code,
		HttpStatus: defn.HttpStatus,
		Message:    defn.Message,
		Err:        err,
	}
}

func (e *PublicError) Log() {
	if e.HttpStatus >= 500 {
		log.Error().
			Str("code", string(e.SystemCode)).
			Int("status", e.HttpStatus).
			Err(e.Err).
			Msg(e.Message)
	} else {
		log.Warn().
			Str("code", string(e.SystemCode)).
			Int("status", e.HttpStatus).
			Msg(e.Message)
	}
}
