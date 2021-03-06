// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/synthetics/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of runs for a specified canary.
func (c *Client) GetCanaryRuns(ctx context.Context, params *GetCanaryRunsInput, optFns ...func(*Options)) (*GetCanaryRunsOutput, error) {
	if params == nil {
		params = &GetCanaryRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCanaryRuns", params, optFns, addOperationGetCanaryRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCanaryRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCanaryRunsInput struct {

	// The name of the canary that you want to see runs for.
	//
	// This member is required.
	Name *string

	// Specify this parameter to limit how many runs are returned each time you use the
	// GetCanaryRuns operation. If you omit this parameter, the default of 100 is used.
	MaxResults *int32

	// A token that indicates that there is more data available. You can use this token
	// in a subsequent GetCanaryRuns operation to retrieve the next set of results.
	NextToken *string
}

type GetCanaryRunsOutput struct {

	// An array of structures. Each structure contains the details of one of the
	// retrieved canary runs.
	CanaryRuns []*types.CanaryRun

	// A token that indicates that there is more data available. You can use this token
	// in a subsequent GetCanaryRuns operation to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCanaryRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCanaryRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCanaryRuns{}, middleware.After)
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
	if err = addOpGetCanaryRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCanaryRuns(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCanaryRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "synthetics",
		OperationName: "GetCanaryRuns",
	}
}
