// Code generated by smithy-go-codegen DO NOT EDIT.

package pricing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pricing/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of attribute values. Attibutes are similar to the details in a
// Price List API offer file. For a list of available attributes, see Offer File
// Definitions
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/reading-an-offer.html#pps-defs)
// in the AWS Billing and Cost Management User Guide
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-what-is.html).
func (c *Client) GetAttributeValues(ctx context.Context, params *GetAttributeValuesInput, optFns ...func(*Options)) (*GetAttributeValuesOutput, error) {
	if params == nil {
		params = &GetAttributeValuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAttributeValues", params, optFns, addOperationGetAttributeValuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAttributeValuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAttributeValuesInput struct {

	// The name of the attribute that you want to retrieve the values for, such as
	// volumeType.
	//
	// This member is required.
	AttributeName *string

	// The service code for the service whose attributes you want to retrieve. For
	// example, if you want the retrieve an EC2 attribute, use AmazonEC2.
	//
	// This member is required.
	ServiceCode *string

	// The maximum number of results to return in response.
	MaxResults *int32

	// The pagination token that indicates the next set of results that you want to
	// retrieve.
	NextToken *string
}

type GetAttributeValuesOutput struct {

	// The list of values for an attribute. For example, Throughput Optimized HDD and
	// Provisioned IOPS are two available values for the AmazonEC2volumeType.
	AttributeValues []*types.AttributeValue

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAttributeValuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAttributeValues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAttributeValues{}, middleware.After)
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
	if err = addOpGetAttributeValuesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAttributeValues(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAttributeValues(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "pricing",
		OperationName: "GetAttributeValues",
	}
}
