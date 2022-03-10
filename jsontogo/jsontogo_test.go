package jsontogo

import (
	"testing"
)

func TestJSONToGo(t *testing.T) {
	var tests = []struct {
		name              string
		jsonStr, typeName string
		want              string
	}{
		{
			name:     "empty",
			jsonStr:  "",
			typeName: "Foo",
			want:     "",
		},
		{
			name:     "simple",
			jsonStr:  `{"query": "some string"}`,
			typeName: "Foo",
			want: `type Foo struct {
	Query string ` + "`" + `json:"query"` + "`" + `
}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := JSONToGo(test.jsonStr, test.typeName)
			if err != nil {
				t.Fatalf("JSONToGo: %v", err)
			}
			if test.want != got {
				t.Errorf("JSONToGo(%q,%q): want %v, got %v", test.jsonStr, test.typeName, test.want, got)
			}
		})
	}
}
