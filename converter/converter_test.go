package converter

import (
	"strconv"
	"testing"
)

func TestJsonToYaml(t *testing.T) {
	cases := []struct {
		name string
		json string
		want string
	}{
		{
			name: "OnePair",
			json: "{\"MyKey\":\"MyValue\"}",
			want: "MyKey: MyValue\n",
		},
		{
			name: "MultiPair",
			json: "{\"MyFirstKey\":\"MyFirstValue\",\"MySecondKey\":0}",
			want: "MyFirstKey: MyFirstValue\nMySecondKey: 0\n",
		},
		{
			name: "Array",
			json: "{\"MyArray\":[1,2,3]}",
			want: "MyArray:\n    - 1\n    - 2\n    - 3\n",
		},
		{
			name: "NestedObject",
			json: "{\"MyKey\":{\"MyNestedKey\":true}}",
			want: "MyKey:\n    MyNestedKey: true\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := JsonToYaml([]byte(tc.json))
			if err != nil {
				t.Errorf("error: %v", err)
				return
			}

			print(strconv.Quote(string(got)))
			if string(got) != tc.want {
				t.Errorf("expected: %v, got: %v", tc.want, string(got))
			}
		})
	}
}
