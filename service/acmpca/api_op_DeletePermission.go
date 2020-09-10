// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Revokes permissions that a private CA assigned to a designated AWS service.
// Permissions can be created with the CreatePermission () action and listed with
// the ListPermissions () action.
func (c *Client) DeletePermission(ctx context.Context, params *DeletePermissionInput, optFns ...func(*Options)) (*DeletePermissionOutput, error) {
	stack := middleware.NewStack("DeletePermission", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeletePermissionMiddlewares(stack)
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
	addOpDeletePermissionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePermission(options.Region), middleware.Before)

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
			OperationName: "DeletePermission",
			Err:           err,
		}
	}
	out := result.(*DeletePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePermissionInput struct {
	// The Amazon Resource Number (ARN) of the private CA that issued the permissions.
	// You can find the CA's ARN by calling the ListCertificateAuthorities () action.
	// This must have the following form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	CertificateAuthorityArn *string
	// The AWS account that calls this action.
	SourceAccount *string
	// The AWS service or identity that will have its CA permissions revoked. At this
	// time, the only valid service principal is acm.amazonaws.com
	Principal *string
}

type DeletePermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeletePermissionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePermission{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePermission{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeletePermission(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "DeletePermission",
	}
}