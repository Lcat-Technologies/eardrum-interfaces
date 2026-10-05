package transaction

import (
	"time"
)



// Transaction represents the core properties of a financial transaction.
type Transaction interface {
	GetTransactionID() string // GetTransactionID returns the unique identifier of the transaction.
	GetUserName() string // GetUserName returns the unique username of the user who made the transaction.
	GetMerchantName() string // GetMerchantName returns the unique username of the merchant where the transaction occurred.
	GetCreatedAt() time.Time // GetCreatedAt returns the creation timestamp of the transaction.
	GetUpdatedAt() time.Time // GetUpdatedAt returns the last update timestamp of the transaction.
	GetDeletedAt() time.Time // GetDeletedAt returns the deletion timestamp. The timestamp is a zero value if the transaction has not been deleted.
	GetTotalAmountInCents() uint // GetTotalAmountInCents returns the total amount spent in the transaction, in cents.
	GetTransactionCostInCents() uint // GetTransactionCostInCents returns the transaction processing cost in cents.
	GetScanLog() string //GetScanLog fetches the image of the scan that authorized transaction
	GetTransactionDeviceID() (string, error) //GetTransactionDeviceID returns the device ID of the device that processed the transaction 
	GetTransactionDeviceModel() *string //GetTransactionDeviceModel returns the model of the device that processed the transaction
}

// TransactionAuthorization provides the data required for a user to authorize a transaction.
// It is intended for embedding in more specific transaction types to avoid code duplication.
type TransactionAuthorization interface {
	GetUUID() string // GetUUID returns the unique UUID string provided by the user to authorize the transaction.
	GetFacialEmbedding() string // GetFacialEmbedding returns the FacialEmbedding provided by the user for authorization.
}

// NewTransaction represents a new transaction with a single total amount, not tied to specific products.
// It embeds the TransactionAuthorization interface to include common authorization data.
type NewTransaction interface {
	TransactionAuthorization
	GetTotalAmountInCents() uint // GetTotalAmountInCents returns the total amount intended to be transacted.
}


// NewOfflineTransaction represents a transaction initiated offline.
// It inherits authorization and amount details from NewTransaction, adding specific offline metadata.
type NewOfflineTransaction interface {
	NewTransaction
	GetPhoneNumber() string       // GetPhoneNumber returns the phone number associated with the offline transaction.
	GetOfflineTimestamp() time.Time // GetOfflineTimestamp returns the actual time the transaction occurred offline.
	GetScanLog()string // GetScanLog fetches the image of the scan that authorized transaction
	GetOfflineTransactionID() string //GetOfflineTransactionID returns the offline transaction id which includes the deviceIMEI and the Unix timestamp.
}