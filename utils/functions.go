package utils

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

func URLFor(pattern string, args...interface{}) string {
	flogURL, err := GetFlogURL()
	cobra.CheckErr(err)
	return fmt.Sprintf(flogURL+pattern, args...)
}

func CheckStatusCode(resp *http.Response, expected int) (err error){
	if resp.StatusCode != expected {
		switch resp.StatusCode {
		case 400:
			err = errors.New("input invalid")
		case 401:
			err = errors.New("unauthorized")
		case 403:
			err = errors.New("you don't have the permission to access this")
		case 404:
			err = errors.New("not found")
		default:
			err = errors.New(fmt.Sprintf("expected %d, actual %d", expected, resp.StatusCode))
		}
		return
	}
	return
}
