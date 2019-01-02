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

// GetAppVersionPackageReader is a Reader for the GetAppVersionPackage structure.
type GetAppVersionPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppVersionPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAppVersionPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAppVersionPackageOK creates a GetAppVersionPackageOK with default headers values
func NewGetAppVersionPackageOK() *GetAppVersionPackageOK {
	return &GetAppVersionPackageOK{}
}

/*GetAppVersionPackageOK handles this case with default header values.

A successful response.
*/
type GetAppVersionPackageOK struct {
	Payload *models.OpenpitrixGetAppVersionPackageResponse
}

func (o *GetAppVersionPackageOK) Error() string {
	return fmt.Sprintf("[GET /api/AppManager.GetAppVersionPackage][%d] getAppVersionPackageOK  %+v", 200, o.Payload)
}

func (o *GetAppVersionPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixGetAppVersionPackageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
