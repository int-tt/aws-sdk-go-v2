// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the events and the corresponding Lambda functions associated with an
// identity pool. This API can only be called with developer credentials. You
// cannot call this API with the temporary user credentials provided by Cognito
// Identity.
func (c *Client) GetCognitoEvents(ctx context.Context, params *GetCognitoEventsInput, optFns ...func(*Options)) (*GetCognitoEventsOutput, error) {
	stack := middleware.NewStack("GetCognitoEvents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetCognitoEventsMiddlewares(stack)
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
	addOpGetCognitoEventsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCognitoEvents(options.Region), middleware.Before)

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
			OperationName: "GetCognitoEvents",
			Err:           err,
		}
	}
	out := result.(*GetCognitoEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for a list of the configured Cognito Events
type GetCognitoEventsInput struct {
	// The Cognito Identity Pool ID for the request
	IdentityPoolId *string
}

// The response from the GetCognitoEvents request
type GetCognitoEventsOutput struct {
	// The Cognito Events returned from the GetCognitoEvents request
	Events map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetCognitoEventsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetCognitoEvents{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCognitoEvents{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCognitoEvents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "GetCognitoEvents",
	}
}