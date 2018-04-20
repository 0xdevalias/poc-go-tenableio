package api

import (
	"fmt"

	"github.com/go-resty/resty"
)

func errorResponseFormatter(resp *resty.Response) error {
	return fmt.Errorf("api returned an error response (%v): %s", resp.StatusCode(), string(resp.Body()))
}
