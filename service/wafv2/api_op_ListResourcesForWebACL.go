// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// Retrieves an array of the Amazon Resource Names (ARNs) for the regional
// resources that are associated with the specified web ACL. If you want the list
// of AWS CloudFront resources, use the AWS CloudFront call
// ListDistributionsByWebACLId.
func (c *Client) ListResourcesForWebACL(ctx context.Context, params *ListResourcesForWebACLInput, optFns ...func(*Options)) (*ListResourcesForWebACLOutput, error) {
	if params == nil {
		params = &ListResourcesForWebACLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResourcesForWebACL", params, optFns, addOperationListResourcesForWebACLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourcesForWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourcesForWebACLInput struct {

	// The Amazon Resource Name (ARN) of the Web ACL.
	//
	// This member is required.
	WebACLArn *string

	// Used for web ACLs that are scoped for regional applications. A regional
	// application can be an Application Load Balancer (ALB), an API Gateway REST API,
	// or an AppSync GraphQL API.
	ResourceType types.ResourceType
}

type ListResourcesForWebACLOutput struct {

	// The array of Amazon Resource Names (ARNs) of the associated resources.
	ResourceArns []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListResourcesForWebACLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResourcesForWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResourcesForWebACL{}, middleware.After)
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
	if err = addOpListResourcesForWebACLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResourcesForWebACL(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListResourcesForWebACL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "ListResourcesForWebACL",
	}
}
