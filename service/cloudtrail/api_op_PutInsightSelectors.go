// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lets you enable Insights event logging by specifying the Insights selectors that
// you want to enable on an existing trail. You also use PutInsightSelectors to
// turn off Insights event logging, by passing an empty list of insight types. The
// valid Insights event type in this release is ApiCallRateInsight.
func (c *Client) PutInsightSelectors(ctx context.Context, params *PutInsightSelectorsInput, optFns ...func(*Options)) (*PutInsightSelectorsOutput, error) {
	if params == nil {
		params = &PutInsightSelectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutInsightSelectors", params, optFns, c.addOperationPutInsightSelectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutInsightSelectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutInsightSelectorsInput struct {

	// A JSON string that contains the Insights types that you want to log on a trail.
	// The valid Insights type in this release is ApiCallRateInsight.
	//
	// This member is required.
	InsightSelectors []types.InsightSelector

	// The name of the CloudTrail trail for which you want to change or add Insights
	// selectors.
	//
	// This member is required.
	TrailName *string

	noSmithyDocumentSerde
}

type PutInsightSelectorsOutput struct {

	// A JSON string that contains the Insights event types that you want to log on a
	// trail. The valid Insights type in this release is ApiCallRateInsight.
	InsightSelectors []types.InsightSelector

	// The Amazon Resource Name (ARN) of a trail for which you want to change or add
	// Insights selectors.
	TrailARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutInsightSelectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutInsightSelectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutInsightSelectors{}, middleware.After)
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
	if err = addOpPutInsightSelectorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutInsightSelectors(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutInsightSelectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "PutInsightSelectors",
	}
}
