// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modify the version-specific settings of a Lambda function. When you update a
// function, Lambda provisions an instance of the function and its supporting
// resources. If your function connects to a VPC, this process can take a minute.
// During this time, you can't modify the function, but you can still invoke it.
// The LastUpdateStatus, LastUpdateStatusReason, and LastUpdateStatusReasonCode
// fields in the response from GetFunctionConfiguration indicate when the update is
// complete and the function is processing events with the new configuration. For
// more information, see Function States
// (https://docs.aws.amazon.com/lambda/latest/dg/functions-states.html). These
// settings can vary between versions of a function and are locked when you publish
// a version. You can't modify the configuration of a published version, only the
// unpublished version. To configure function concurrency, use
// PutFunctionConcurrency. To grant invoke permissions to an account or Amazon Web
// Services service, use AddPermission.
func (c *Client) UpdateFunctionConfiguration(ctx context.Context, params *UpdateFunctionConfigurationInput, optFns ...func(*Options)) (*UpdateFunctionConfigurationOutput, error) {
	if params == nil {
		params = &UpdateFunctionConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFunctionConfiguration", params, optFns, c.addOperationUpdateFunctionConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFunctionConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFunctionConfigurationInput struct {

	// The name of the Lambda function. Name formats
	//
	// * Function name - my-function.
	//
	// *
	// Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	// *
	// Partial ARN - 123456789012:function:my-function.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it is limited to 64
	// characters in length.
	//
	// This member is required.
	FunctionName *string

	// A dead letter queue configuration that specifies the queue or topic where Lambda
	// sends asynchronous events when they fail processing. For more information, see
	// Dead Letter Queues
	// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq).
	DeadLetterConfig *types.DeadLetterConfig

	// A description of the function.
	Description *string

	// Environment variables that are accessible from function code during execution.
	Environment *types.Environment

	// Connection settings for an Amazon EFS file system.
	FileSystemConfigs []types.FileSystemConfig

	// The name of the method within your code that Lambda calls to execute your
	// function. The format includes the file name. It can also include namespaces and
	// other qualifiers, depending on the runtime. For more information, see
	// Programming Model
	// (https://docs.aws.amazon.com/lambda/latest/dg/programming-model-v2.html).
	Handler *string

	// Container image configuration values
	// (https://docs.aws.amazon.com/lambda/latest/dg/images-parms.html) that override
	// the values in the container image Dockerfile.
	ImageConfig *types.ImageConfig

	// The ARN of the Amazon Web Services Key Management Service (KMS) key that's used
	// to encrypt your function's environment variables. If it's not provided, Lambda
	// uses a default service key.
	KMSKeyArn *string

	// A list of function layers
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html) to add
	// to the function's execution environment. Specify each layer by its ARN,
	// including the version.
	Layers []string

	// The amount of memory available to the function
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-memory.html) at
	// runtime. Increasing the function memory also increases its CPU allocation. The
	// default value is 128 MB. The value can be any multiple of 1 MB.
	MemorySize *int32

	// Only update the function if the revision ID matches the ID that's specified. Use
	// this option to avoid modifying a function that has changed since you last read
	// it.
	RevisionId *string

	// The Amazon Resource Name (ARN) of the function's execution role.
	Role *string

	// The identifier of the function's runtime
	// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html).
	Runtime types.Runtime

	// The amount of time that Lambda allows a function to run before stopping it. The
	// default is 3 seconds. The maximum allowed value is 900 seconds. For additional
	// information, see Lambda execution environment
	// (https://docs.aws.amazon.com/lambda/latest/dg/runtimes-context.html).
	Timeout *int32

	// Set Mode to Active to sample and trace a subset of incoming requests with X-Ray
	// (https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html).
	TracingConfig *types.TracingConfig

	// For network connectivity to Amazon Web Services resources in a VPC, specify a
	// list of security groups and subnets in the VPC. When you connect a function to a
	// VPC, it can only access resources and the internet through that VPC. For more
	// information, see VPC Settings
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

// Details about a function's configuration.
type UpdateFunctionConfigurationOutput struct {

	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string

	// The size of the function's deployment package, in bytes.
	CodeSize int64

	// The function's dead letter queue.
	DeadLetterConfig *types.DeadLetterConfig

	// The function's description.
	Description *string

	// The function's environment variables
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html).
	Environment *types.EnvironmentResponse

	// Connection settings for an Amazon EFS file system
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-filesystem.html).
	FileSystemConfigs []types.FileSystemConfig

	// The function's Amazon Resource Name (ARN).
	FunctionArn *string

	// The name of the function.
	FunctionName *string

	// The function that Lambda calls to begin executing your function.
	Handler *string

	// The function's image configuration values.
	ImageConfigResponse *types.ImageConfigResponse

	// The KMS key that's used to encrypt the function's environment variables. This
	// key is only returned if you've configured a customer managed CMK.
	KMSKeyArn *string

	// The date and time that the function was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string

	// The status of the last update that was performed on the function. This is first
	// set to Successful after function creation completes.
	LastUpdateStatus types.LastUpdateStatus

	// The reason for the last update that was performed on the function.
	LastUpdateStatusReason *string

	// The reason code for the last update that was performed on the function.
	LastUpdateStatusReasonCode types.LastUpdateStatusReasonCode

	// The function's  layers
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	Layers []types.Layer

	// For Lambda@Edge functions, the ARN of the master function.
	MasterArn *string

	// The amount of memory available to the function at runtime.
	MemorySize *int32

	// The type of deployment package. Set to Image for container image and set Zip for
	// .zip file archive.
	PackageType types.PackageType

	// The latest updated revision of the function or alias.
	RevisionId *string

	// The function's execution role.
	Role *string

	// The runtime environment for the Lambda function.
	Runtime types.Runtime

	// The ARN of the signing job.
	SigningJobArn *string

	// The ARN of the signing profile version.
	SigningProfileVersionArn *string

	// The current state of the function. When the state is Inactive, you can
	// reactivate the function by invoking it.
	State types.State

	// The reason for the function's current state.
	StateReason *string

	// The reason code for the function's current state. When the code is Creating, you
	// can't invoke or modify the function.
	StateReasonCode types.StateReasonCode

	// The amount of time in seconds that Lambda allows a function to run before
	// stopping it.
	Timeout *int32

	// The function's X-Ray tracing configuration.
	TracingConfig *types.TracingConfigResponse

	// The version of the Lambda function.
	Version *string

	// The function's networking configuration.
	VpcConfig *types.VpcConfigResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFunctionConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFunctionConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFunctionConfiguration{}, middleware.After)
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
	if err = addOpUpdateFunctionConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFunctionConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFunctionConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "UpdateFunctionConfiguration",
	}
}
