// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts processing a stream processor. You create a stream processor by calling
// CreateStreamProcessor (). To tell StartStreamProcessor which stream processor to
// start, use the value of the Name field specified in the call to
// CreateStreamProcessor.
func (c *Client) StartStreamProcessor(ctx context.Context, params *StartStreamProcessorInput, optFns ...func(*Options)) (*StartStreamProcessorOutput, error) {
	stack := middleware.NewStack("StartStreamProcessor", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartStreamProcessorMiddlewares(stack)
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
	addOpStartStreamProcessorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartStreamProcessor(options.Region), middleware.Before)

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
			OperationName: "StartStreamProcessor",
			Err:           err,
		}
	}
	out := result.(*StartStreamProcessorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartStreamProcessorInput struct {
	// The name of the stream processor to start processing.
	Name *string
}

type StartStreamProcessorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartStreamProcessorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartStreamProcessor{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartStreamProcessor{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartStreamProcessor(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartStreamProcessor",
	}
}