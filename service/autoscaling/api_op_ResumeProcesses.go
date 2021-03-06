// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resumes the specified suspended automatic scaling processes, or all suspended
// process, for the specified Auto Scaling group. For more information, see
// Suspending and Resuming Scaling Processes
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-suspend-resume-processes.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) ResumeProcesses(ctx context.Context, params *ResumeProcessesInput, optFns ...func(*Options)) (*ResumeProcessesOutput, error) {
	if params == nil {
		params = &ResumeProcessesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResumeProcesses", params, optFns, addOperationResumeProcessesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResumeProcessesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResumeProcessesInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// One or more of the following processes:
	//
	// * Launch
	//
	// * Terminate
	//
	// *
	// AddToLoadBalancer
	//
	// * AlarmNotification
	//
	// * AZRebalance
	//
	// * HealthCheck
	//
	// *
	// InstanceRefresh
	//
	// * ReplaceUnhealthy
	//
	// * ScheduledActions
	//
	// If you omit this
	// parameter, all processes are specified.
	ScalingProcesses []*string
}

type ResumeProcessesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationResumeProcessesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpResumeProcesses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpResumeProcesses{}, middleware.After)
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
	if err = addOpResumeProcessesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResumeProcesses(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResumeProcesses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "ResumeProcesses",
	}
}
