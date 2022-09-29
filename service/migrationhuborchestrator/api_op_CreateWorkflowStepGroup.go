// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhuborchestrator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/migrationhuborchestrator/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Create a step group in a migration workflow.
func (c *Client) CreateWorkflowStepGroup(ctx context.Context, params *CreateWorkflowStepGroupInput, optFns ...func(*Options)) (*CreateWorkflowStepGroupOutput, error) {
	if params == nil {
		params = &CreateWorkflowStepGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkflowStepGroup", params, optFns, c.addOperationCreateWorkflowStepGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkflowStepGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkflowStepGroupInput struct {

	// The name of the step group.
	//
	// This member is required.
	Name *string

	// The ID of the migration workflow that will contain the step group.
	//
	// This member is required.
	WorkflowId *string

	// The description of the step group.
	Description *string

	// The next step group.
	Next []string

	// The previous step group.
	Previous []string

	noSmithyDocumentSerde
}

type CreateWorkflowStepGroupOutput struct {

	// The time at which the step group is created.
	CreationTime *time.Time

	// The description of the step group.
	Description *string

	// The ID of the step group.
	Id *string

	// The name of the step group.
	Name *string

	// The next step group.
	Next []string

	// The previous step group.
	Previous []string

	// List of AWS services utilized in a migration workflow.
	Tools []types.Tool

	// The ID of the migration workflow that contains the step group.
	WorkflowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkflowStepGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorkflowStepGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorkflowStepGroup{}, middleware.After)
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
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
	if err = addOpCreateWorkflowStepGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkflowStepGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWorkflowStepGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "migrationhub-orchestrator",
		OperationName: "CreateWorkflowStepGroup",
	}
}
