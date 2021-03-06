// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// A UsageRecord indicates a quantity of usage for a given product, customer,
// dimension and time. Multiple requests with the same UsageRecords as input will
// be deduplicated to prevent double charges.
type UsageRecord struct {

	// The CustomerIdentifier is obtained through the ResolveCustomer operation and
	// represents an individual buyer in your application.
	//
	// This member is required.
	CustomerIdentifier *string

	// During the process of registering a product on AWS Marketplace, up to eight
	// dimensions are specified. These represent different units of value in your
	// application.
	//
	// This member is required.
	Dimension *string

	// Timestamp, in UTC, for which the usage is being reported. Your application can
	// meter usage for up to one hour in the past. Make sure the timestamp value is not
	// before the start of the software usage.
	//
	// This member is required.
	Timestamp *time.Time

	// The quantity of usage consumed by the customer for the given dimension and time.
	// Defaults to 0 if not specified.
	Quantity *int32
}

// A UsageRecordResult indicates the status of a given UsageRecord processed by
// BatchMeterUsage.
type UsageRecordResult struct {

	// The MeteringRecordId is a unique identifier for this metering event.
	MeteringRecordId *string

	// The UsageRecordResult Status indicates the status of an individual UsageRecord
	// processed by BatchMeterUsage.
	//
	// * Success- The UsageRecord was accepted and
	// honored by BatchMeterUsage.
	//
	// * CustomerNotSubscribed- The CustomerIdentifier
	// specified is not subscribed to your product. The UsageRecord was not honored.
	// Future UsageRecords for this customer will fail until the customer subscribes to
	// your product.
	//
	// * DuplicateRecord- Indicates that the UsageRecord was invalid and
	// not honored. A previously metered UsageRecord had the same customer, dimension,
	// and time, but a different quantity.
	Status UsageRecordResultStatus

	// The UsageRecord that was part of the BatchMeterUsage request.
	UsageRecord *UsageRecord
}
