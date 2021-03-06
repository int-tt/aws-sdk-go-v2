// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays a list of all reservations that have been purchased by this account in
// the current AWS Region. This list includes all reservations in all states (such
// as active and expired).
func (c *Client) ListReservations(ctx context.Context, params *ListReservationsInput, optFns ...func(*Options)) (*ListReservationsOutput, error) {
	if params == nil {
		params = &ListReservationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReservations", params, optFns, addOperationListReservationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReservationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReservationsInput struct {

	// The maximum number of results to return per API request. For example, you submit
	// a ListReservations request with MaxResults set at 5. Although 20 items match
	// your request, the service returns no more than the first 5 items. (The service
	// also returns a NextToken value that you can use to fetch the next batch of
	// results.) The service might return fewer results than the MaxResults value. If
	// MaxResults is not included in the request, the service defaults to pagination
	// with a maximum of 10 results per page.
	MaxResults *int32

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListReservations request with MaxResults set at 5. The
	// service returns the first batch of results (up to 5) and a NextToken value. To
	// see the next batch of results, you can submit the ListOfferings request a second
	// time and specify the NextToken value.
	NextToken *string
}

type ListReservationsOutput struct {

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListReservations request with MaxResults set at 5. The
	// service returns the first batch of results (up to 5) and a NextToken value. To
	// see the next batch of results, you can submit the ListReservations request a
	// second time and specify the NextToken value.
	NextToken *string

	// A list of all reservations that have been purchased by this account in the
	// current AWS Region.
	Reservations []*types.Reservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListReservationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReservations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReservations{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReservations(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListReservations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "ListReservations",
	}
}
