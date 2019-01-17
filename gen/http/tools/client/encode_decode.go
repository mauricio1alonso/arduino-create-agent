// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// tools HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/arduino/arduino-create-agent/design -debug

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	tools "github.com/arduino/arduino-create-agent/gen/tools"
	toolsviews "github.com/arduino/arduino-create-agent/gen/tools/views"
	goahttp "goa.design/goa/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "tools" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListToolsPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("tools", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListResponse returns a decoder for responses returned by the tools
// list endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("tools", "list", err)
			}
			p := NewListToolOK(&body)
			view := "default"
			vres := &toolsviews.Tool{p, view}
			if err = toolsviews.ValidateTool(vres); err != nil {
				return nil, goahttp.ErrValidationError("tools", "list", err)
			}
			res := tools.NewTool(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("tools", "list", resp.StatusCode, string(body))
		}
	}
}
