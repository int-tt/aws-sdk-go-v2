// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// This API action gets a bucket policy for an Amazon S3 on Outposts bucket. To get
// a policy for an S3 bucket, see GetBucketPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicy.html) in
// the Amazon Simple Storage Service API. Returns the policy of a specified
// Outposts bucket. For more information, see Using Amazon S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in the
// Amazon Simple Storage Service Developer Guide. If you are using an identity
// other than the root user of the AWS account that owns the bucket, the calling
// identity must have the GetBucketPolicy permissions on the specified bucket and
// belong to the bucket owner's account in order to use this operation. If you
// don't have s3outposts:GetBucketPolicy permissions, Amazon S3 returns a 403
// Access Denied error. If you have the correct permissions, but you're not using
// an identity that belongs to the bucket owner's account, Amazon S3 returns a 405
// Method Not Allowed error. As a security precaution, the root user of the AWS
// account that owns a bucket can always use this operation, even if the policy
// explicitly denies the root user the ability to perform this action. For more
// information about bucket policies, see Using Bucket Policies and User Policies
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html). All
// Amazon S3 on Outposts REST API requests for this action require an additional
// parameter of outpost-id to be passed with the request and an S3 on Outposts
// endpoint hostname prefix instead of s3-control. For an example of the request
// syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname
// prefix and the outpost-id derived using the access point ARN, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_GetBucketPolicy.html#API_control_GetBucketPolicy_Examples)
// section below. The following actions are related to GetBucketPolicy:
//
// *
// GetObject
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html)
//
// *
// PutBucketPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketPolicy.html)
//
// *
// DeleteBucketPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketPolicy.html)
func (c *Client) GetBucketPolicy(ctx context.Context, params *GetBucketPolicyInput, optFns ...func(*Options)) (*GetBucketPolicyOutput, error) {
	if params == nil {
		params = &GetBucketPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketPolicy", params, optFns, addOperationGetBucketPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketPolicyInput struct {

	// The AWS account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The ARN of the bucket. For Amazon S3 on Outposts specify the ARN of the bucket
	// accessed in the format arn:aws:s3-outposts:::outpost//bucket/. For example, to
	// access the bucket reports through outpost my-outpost owned by account
	// 123456789012 in Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string
}

type GetBucketPolicyOutput struct {

	// The policy of the Outposts bucket.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addEndpointPrefix_opGetBucketPolicyMiddleware(stack)
	addOpGetBucketPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketPolicy(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	return nil
}

type endpointPrefix_opGetBucketPolicyMiddleware struct {
}

func (*endpointPrefix_opGetBucketPolicyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetBucketPolicyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetBucketPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.HostPrefix = prefix.String()

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetBucketPolicyMiddleware(stack *middleware.Stack) (err error) {
	err = stack.Serialize.Insert(&endpointPrefix_opGetBucketPolicyMiddleware{}, `OperationSerializer`, middleware.Before)
	if err != nil {
		return err
	}
	err = stack.Build.Add(&smithyhttp.HostPrefixMiddleware{}, middleware.Before)
	if err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetBucketPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketPolicy",
	}
}
