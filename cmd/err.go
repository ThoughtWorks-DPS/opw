package cmd

import (
	"log"
	"fmt"
	"errors"
	"github.com/1Password/connect-sdk-go/onepassword"
)

func exitOnError(err error) bool {
	if err != nil {
		var opErr *onepassword.Error
    if errors.As(err, &opErr){
        fmt.Printf("message=%s, status code=%d\n",
            opErr.Message,
            opErr.StatusCode,
        )
    }
		log.Fatal(err)
	}
	return true
}
