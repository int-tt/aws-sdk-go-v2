// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an alias for a fleet. In most situations, you can use an alias ID in
// place of a fleet ID. An alias provides a level of abstraction for a fleet that
// is useful when redirecting player traffic from one fleet to another, such as
// when updating your game build. Amazon GameLift supports two types of routing
// strategies for aliases: simple and terminal. A simple alias points to an active
// fleet. A terminal alias is used to display messaging or link to a URL instead of
// routing players to an active fleet. For example, you might use a terminal alias
// when a game version is no longer supported and you want to direct players to an
// upgrade site. To create a fleet alias, specify an alias name, routing strategy,
// and optional description. Each simple alias can point to only one fleet, but a
// fleet can have multiple aliases. If successful, a new alias record is returned,
// including an alias ID and an ARN. You can reassign an alias to another fleet by
// calling UpdateAlias.
//
//     * CreateAlias ()
//
//     * ListAliases ()
//
//     *
// DescribeAlias ()
//
//     * UpdateAlias ()
//
//     * DeleteAlias ()
//
//     * ResolveAlias
// ()
func (c *Client) CreateAlias(ctx context.Context, params *CreateAliasInput, optFns ...func(*Options)) (*CreateAliasOutput, error) {
	stack := middleware.NewStack("CreateAlias", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateAliasMiddlewares(stack)
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
	addOpCreateAliasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAlias(options.Region), middleware.Before)

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
			OperationName: "CreateAlias",
			Err:           err,
		}
	}
	out := result.(*CreateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type CreateAliasInput struct {
	// A descriptive label that is associated with an alias. Alias names do not need to
	// be unique.
	Name *string
	// A list of labels to assign to the new alias resource. Tags are developer-defined
	// key-value pairs. Tagging AWS resources are useful for resource management,
	// access management and cost allocation. For more information, see  Tagging AWS
	// Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in
	// the AWS General Reference. Once the resource is created, you can use TagResource
	// (), UntagResource (), and ListTagsForResource () to add, remove, and view tags.
	// The maximum tag limit may be lower than stated. See the AWS General Reference
	// for actual tagging limits.
	Tags []*types.Tag
	// A human-readable description of the alias.
	Description *string
	// The routing configuration, including routing type and fleet target, for the
	// alias.
	RoutingStrategy *types.RoutingStrategy
}

// Represents the returned data in response to a request action.
type CreateAliasOutput struct {
	// The newly created alias resource.
	Alias *types.Alias

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateAliasMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAlias{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAlias{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateAlias(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateAlias",
	}
}