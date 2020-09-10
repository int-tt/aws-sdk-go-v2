// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a proposal for a change to the network that other members of the network
// can vote on, for example, a proposal to add a new member to the network. Any
// member can create a proposal.
func (c *Client) CreateProposal(ctx context.Context, params *CreateProposalInput, optFns ...func(*Options)) (*CreateProposalOutput, error) {
	stack := middleware.NewStack("CreateProposal", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateProposalMiddlewares(stack)
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
	addIdempotencyToken_opCreateProposalMiddleware(stack, options)
	addOpCreateProposalValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProposal(options.Region), middleware.Before)

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
			OperationName: "CreateProposal",
			Err:           err,
		}
	}
	out := result.(*CreateProposalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProposalInput struct {
	// A description for the proposal that is visible to voting members, for example,
	// "Proposal to add Example Corp. as member."
	Description *string
	// The unique identifier of the network for which the proposal is made.
	NetworkId *string
	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the operation. An idempotent operation completes no more than one time. This
	// identifier is required only if you make a service request directly using an HTTP
	// client. It is generated automatically if you use an AWS SDK or the AWS CLI.
	ClientRequestToken *string
	// The unique identifier of the member that is creating the proposal. This
	// identifier is especially useful for identifying the member making the proposal
	// when multiple members exist in a single AWS account.
	MemberId *string
	// The type of actions proposed, such as inviting a member or removing a member.
	// The types of Actions in a proposal are mutually exclusive. For example, a
	// proposal with Invitations actions cannot also contain Removals actions.
	Actions *types.ProposalActions
}

type CreateProposalOutput struct {
	// The unique identifier of the proposal.
	ProposalId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateProposalMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateProposal{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProposal{}, middleware.After)
}

type idempotencyToken_initializeOpCreateProposal struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateProposal) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateProposal) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateProposalInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateProposalInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateProposalMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateProposal{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateProposal(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "CreateProposal",
	}
}