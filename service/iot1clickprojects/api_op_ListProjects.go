// Code generated by smithy-go-codegen DO NOT EDIT.

package iot1clickprojects

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot1clickprojects/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the AWS IoT 1-Click project(s) associated with your AWS account and
// region.
func (c *Client) ListProjects(ctx context.Context, params *ListProjectsInput, optFns ...func(*Options)) (*ListProjectsOutput, error) {
	stack := middleware.NewStack("ListProjects", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListProjectsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListProjects(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ListProjects",
			Err:           err,
		}
	}
	out := result.(*ListProjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProjectsInput struct {
	// The maximum number of results to return per request. If not set, a default value
	// of 100 is used.
	MaxResults *int32
	// The token to retrieve the next set of results.
	NextToken *string
}

type ListProjectsOutput struct {
	// The token used to retrieve the next set of results - will be effectively empty
	// if there are no further results.
	NextToken *string
	// An object containing the list of projects.
	Projects []*types.ProjectSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListProjectsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListProjects{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListProjects{}, middleware.After)
}

func newServiceMetadataMiddleware_opListProjects(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iot1click",
		OperationName: "ListProjects",
	}
}