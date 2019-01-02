// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DescribeActiveAppVersionsReader is a Reader for the DescribeActiveAppVersions structure.
type DescribeActiveAppVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeActiveAppVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeActiveAppVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeActiveAppVersionsOK creates a DescribeActiveAppVersionsOK with default headers values
func NewDescribeActiveAppVersionsOK() *DescribeActiveAppVersionsOK {
	return &DescribeActiveAppVersionsOK{}
}

/*DescribeActiveAppVersionsOK handles this case with default header values.

A successful response.
*/
type DescribeActiveAppVersionsOK struct {
	Payload *models.OpenpitrixDescribeAppVersionsResponse
}

func (o *DescribeActiveAppVersionsOK) Error() string {
	return fmt.Sprintf("[GET /api/AppManager.DescribeActiveAppVersions][%d] describeActiveAppVersionsOK  %+v", 200, o.Payload)
}

func (o *DescribeActiveAppVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeAppVersionsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
