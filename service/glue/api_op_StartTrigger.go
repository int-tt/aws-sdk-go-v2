// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an existing trigger. See Triggering Jobs
// (https://docs.aws.amazon.com/glue/latest/dg/trigger-job.html) for information
// about how different types of trigger are started.
func (c *Client) StartTrigger(ctx context.Context, params *StartTriggerInput, optFns ...func(*Options)) (*StartTriggerOutput, error) {
	stack := middleware.NewStack("StartTrigger", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartTriggerMiddlewares(stack)
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
	addOpStartTriggerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartTrigger(options.Region), middleware.Before)

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
			OperationName: "StartTrigger",
			Err:           err,
		}
	}
	out := result.(*StartTriggerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTriggerInput struct {
	// The name of the trigger to start.
	Name *string
}

type StartTriggerOutput struct {
	// The name of the trigger that was started.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartTriggerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartTrigger{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartTrigger{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartTrigger(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartTrigger",
	}
}