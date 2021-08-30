package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	mr "github.com/MichaelMure/go-term-markdown"

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

func OutputPost(post Post, short, verbose, veryVerbose bool) {
	converter := md.NewConverter("", true, nil)
	switch {
	case short:
		fmt.Print(post.ID, " ")
	case verbose:
		fmt.Println(Segmenter)
		fmt.Println("Post ID:    ", post.ID)
		fmt.Println("Post title: ", post.Title)
		fmt.Println("Private:    ", post.Private)
		fmt.Println("Column(s):    ", post.Columns)
		fmt.Print("Content: ")
		if !(strings.Contains(post.Content, "</iframe>") || strings.Contains(post.Content, "<img")) {
			markdown, err := converter.ConvertString(post.Content)
			cobra.CheckErr(err)
			content := mr.Render(markdown, 120, 0)
			strContent := string(content)
			if len(strContent) > 200 {
				fmt.Println(strContent[:200])
			} else {
				fmt.Println(strContent)
			}
		} else {
			fmt.Println("Post Content contains iframes or images which cannot printed in the terminal")
			flogURL, err := GetFlogURL()
			cobra.CheckErr(err)
			fmt.Printf("Please visit %s%s to view it\n", flogURL, post.Self)
		}
	case veryVerbose:
		fmt.Println(Segmenter)
		fmt.Println("Post ID:    ", post.ID)
		fmt.Println("Post title: ", post.Title)
		fmt.Println("Private:    ", post.Private)
		fmt.Println("Columns:    ", post.Columns)
		fmt.Println("Comments:   ", post.Comments)
		fmt.Println("URL:        ", post.Self)
		fmt.Print("Content: ")
		if !(strings.Contains(post.Content, "</iframe>") || strings.Contains(post.Content, "<img")) {
			markdown, err := converter.ConvertString(post.Content)
			cobra.CheckErr(err)
			content := mr.Render(markdown, 120, 0)
			fmt.Println(string(content))
		} else {
			fmt.Println("Post Content contains iframes or images which cannot printed in the terminal")
			flogURL, err := GetFlogURL()
			cobra.CheckErr(err)
			fmt.Printf("Please visit %s%s to view it\n", flogURL, post.Self)
		}
	default:
		fmt.Println(Segmenter)
		fmt.Println("Post ID:    ", post.ID)
		fmt.Println("Post title: ", post.Title)
		fmt.Println("Private:    ", post.Private)
	}
}
