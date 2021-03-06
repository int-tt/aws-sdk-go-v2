// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the permissions for your network interfaces.
func (c *Client) DescribeNetworkInterfacePermissions(ctx context.Context, params *DescribeNetworkInterfacePermissionsInput, optFns ...func(*Options)) (*DescribeNetworkInterfacePermissionsOutput, error) {
	if params == nil {
		params = &DescribeNetworkInterfacePermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNetworkInterfacePermissions", params, optFns, addOperationDescribeNetworkInterfacePermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNetworkInterfacePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeNetworkInterfacePermissions.
type DescribeNetworkInterfacePermissionsInput struct {

	// One or more filters.
	//
	// *
	// network-interface-permission.network-interface-permission-id - The ID of the
	// permission.
	//
	// * network-interface-permission.network-interface-id - The ID of the
	// network interface.
	//
	// * network-interface-permission.aws-account-id - The AWS
	// account ID.
	//
	// * network-interface-permission.aws-service - The AWS service.
	//
	// *
	// network-interface-permission.permission - The type of permission
	// (INSTANCE-ATTACH | EIP-ASSOCIATE).
	Filters []*types.Filter

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. If this
	// parameter is not specified, up to 50 results are returned by default.
	MaxResults *int32

	// One or more network interface permission IDs.
	NetworkInterfacePermissionIds []*string

	// The token to request the next page of results.
	NextToken *string
}

// Contains the output for DescribeNetworkInterfacePermissions.
type DescribeNetworkInterfacePermissionsOutput struct {

	// The network interface permissions.
	NetworkInterfacePermissions []*types.NetworkInterfacePermission

	// The token to use to retrieve the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeNetworkInterfacePermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeNetworkInterfacePermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeNetworkInterfacePermissions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNetworkInterfacePermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeNetworkInterfacePermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeNetworkInterfacePermissions",
	}
}
