// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Replaces the current rules of the specified IP access control group with the
// specified rules.
func (c *Client) UpdateRulesOfIpGroup(ctx context.Context, params *UpdateRulesOfIpGroupInput, optFns ...func(*Options)) (*UpdateRulesOfIpGroupOutput, error) {
	stack := middleware.NewStack("UpdateRulesOfIpGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateRulesOfIpGroupMiddlewares(stack)
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
	addOpUpdateRulesOfIpGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRulesOfIpGroup(options.Region), middleware.Before)

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
			OperationName: "UpdateRulesOfIpGroup",
			Err:           err,
		}
	}
	out := result.(*UpdateRulesOfIpGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRulesOfIpGroupInput struct {
	// One or more rules.
	UserRules []*types.IpRuleItem
	// The identifier of the group.
	GroupId *string
}

type UpdateRulesOfIpGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateRulesOfIpGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateRulesOfIpGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateRulesOfIpGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRulesOfIpGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "UpdateRulesOfIpGroup",
	}
}