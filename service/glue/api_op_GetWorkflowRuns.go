// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves metadata for all runs of a given workflow.
func (c *Client) GetWorkflowRuns(ctx context.Context, params *GetWorkflowRunsInput, optFns ...func(*Options)) (*GetWorkflowRunsOutput, error) {
	if params == nil {
		params = &GetWorkflowRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkflowRuns", params, optFns, addOperationGetWorkflowRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkflowRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkflowRunsInput struct {

	// Name of the workflow whose metadata of runs should be returned.
	//
	// This member is required.
	Name *string

	// Specifies whether to include the workflow graph in response or not.
	IncludeGraph *bool

	// The maximum number of workflow runs to be included in the response.
	MaxResults *int32

	// The maximum size of the response.
	NextToken *string
}

type GetWorkflowRunsOutput struct {

	// A continuation token, if not all requested workflow runs have been returned.
	NextToken *string

	// A list of workflow run metadata objects.
	Runs []*types.WorkflowRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetWorkflowRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetWorkflowRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetWorkflowRuns{}, middleware.After)
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
	if err = addOpGetWorkflowRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkflowRuns(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWorkflowRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetWorkflowRuns",
	}
}
