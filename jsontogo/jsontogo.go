package jsontogo

import (
	"fmt"

	"github.com/pkg/errors"
	v8 "rogchap.com/v8go"
)

// JSONToGo converts the JSON string `jsonStr` into a go struct with type `typeName`
func JSONToGo(jsonStr, typeName string) (string, error) {
	ctx := v8.NewContext()
	if _, err := ctx.RunScript(jsonToGoJS, "jsonToGo.js"); err != nil {
		return "", errors.Errorf("trouble executing JS: %v", err)
	}
	js := fmt.Sprintf("jsonToGo(%q, %q).go", jsonStr, typeName)
	res, err := ctx.RunScript(js, "jsonToGo.js")
	if err != nil {
		return "", errors.Errorf("trouble executing JS: %v", err)
	}
	return res.String(), nil
}
