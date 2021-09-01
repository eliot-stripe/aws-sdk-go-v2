// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the data for a fleet metric. Requires permission to access the
// UpdateFleetMetric
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) UpdateFleetMetric(ctx context.Context, params *UpdateFleetMetricInput, optFns ...func(*Options)) (*UpdateFleetMetricOutput, error) {
	if params == nil {
		params = &UpdateFleetMetricInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFleetMetric", params, optFns, c.addOperationUpdateFleetMetricMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFleetMetricOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFleetMetricInput struct {

	// The name of the index to search.
	//
	// This member is required.
	IndexName *string

	// The name of the fleet metric to update.
	//
	// This member is required.
	MetricName *string

	// The field to aggregate.
	AggregationField *string

	// The type of the aggregation query.
	AggregationType *types.AggregationType

	// The description of the fleet metric.
	Description *string

	// The expected version of the fleet metric record in the registry.
	ExpectedVersion *int64

	// The time in seconds between fleet metric emissions. Range [60(1 min), 86400(1
	// day)] and must be multiple of 60.
	Period *int32

	// The search query string.
	QueryString *string

	// The version of the query.
	QueryVersion *string

	// Used to support unit transformation such as milliseconds to seconds. The unit
	// must be supported by CW metric
	// (https://docs.aws.amazon.com/https:/docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html).
	Unit types.FleetMetricUnit

	noSmithyDocumentSerde
}

type UpdateFleetMetricOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFleetMetricMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFleetMetric{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFleetMetric{}, middleware.After)
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
	if err = addOpUpdateFleetMetricValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFleetMetric(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFleetMetric(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateFleetMetric",
	}
}
