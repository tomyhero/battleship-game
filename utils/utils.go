package utils

import (
	"fmt"
	"github.com/nu7hatch/gouuid"
)

func GUID() string {
	u4, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s", u4)

}
