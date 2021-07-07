package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os/exec"

	"github.com/spf13/cobra"
)

func SetArgs(args ...string) *exec.Cmd {
	argsCopy := args
	args = append([]string{"run", "../../main.go"}, argsCopy...)
	fmt.Println(args)
	return exec.Command("go", args...)
}

func URLFor(pattern string, args ...interface{}) string {
	flogURL, err := GetFlogURL()
	cobra.CheckErr(err)
	fmt.Println(fmt.Sprintf(flogURL+pattern, args...))
	return fmt.Sprintf(flogURL+pattern, args...)
}

func CheckStatusCode(resp *http.Response, expected int) (err error) {
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
			err = fmt.Errorf("expected %d, actual %d", expected, resp.StatusCode)
		}
		return
	}
	return
}
