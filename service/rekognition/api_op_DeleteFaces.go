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

// Deletes faces from a collection. You specify a collection ID and an array of
// face IDs to remove from the collection. This operation requires permissions to
// perform the rekognition:DeleteFaces action.
func (c *Client) DeleteFaces(ctx context.Context, params *DeleteFacesInput, optFns ...func(*Options)) (*DeleteFacesOutput, error) {
	stack := middleware.NewStack("DeleteFaces", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteFacesMiddlewares(stack)
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
	addOpDeleteFacesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFaces(options.Region), middleware.Before)

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
			OperationName: "DeleteFaces",
			Err:           err,
		}
	}
	out := result.(*DeleteFacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteFacesInput struct {
	// Collection from which to remove the specific faces.
	CollectionId *string
	// An array of face IDs to delete.
	FaceIds []*string
}

type DeleteFacesOutput struct {
	// An array of strings (face IDs) of the faces that were deleted.
	DeletedFaces []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteFacesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteFaces{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteFaces{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteFaces(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DeleteFaces",
	}
}