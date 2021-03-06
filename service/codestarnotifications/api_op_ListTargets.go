// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the notification rule targets for an AWS account.
func (c *Client) ListTargets(ctx context.Context, params *ListTargetsInput, optFns ...func(*Options)) (*ListTargetsOutput, error) {
	if params == nil {
		params = &ListTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTargets", params, optFns, addOperationListTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTargetsInput struct {

	// The filters to use to return information by service or resource type. Valid
	// filters include target type, target address, and target status. A filter with
	// the same name can appear more than once when used with OR statements. Filters
	// with different names should be applied with AND statements.
	Filters []*types.ListTargetsFilter

	// A non-negative integer used to limit the number of returned results. The maximum
	// number of results that can be returned is 100.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string
}

type ListTargetsOutput struct {

	// An enumeration token that can be used in a request to return the next batch of
	// results.
	NextToken *string

	// The list of notification rule targets.
	Targets []*types.TargetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTargets{}, middleware.After)
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
	if err = addOpListTargetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTargets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-notifications",
		OperationName: "ListTargets",
	}
}
