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

func TestYamlToJson(t *testing.T) {
	cases := []struct {
		name string
		yaml string
		want string
	}{
		{
			name: "OnePair",
			yaml: "MyKey: MyValue\n",
			want: "{\"MyKey\":\"MyValue\"}",
		},
		{
			name: "MultiPair",
			yaml: "MyFirstKey: MyFirstValue\nMySecondKey: 0\n",
			want: "{\"MyFirstKey\":\"MyFirstValue\",\"MySecondKey\":0}",
		},
		{
			name: "Array",
			yaml: "MyArray:\n    - 1\n    - 2\n    - 3\n",
			want: "{\"MyArray\":[1,2,3]}",
		},
		{
			name: "NestedObject",
			yaml: "MyKey:\n    MyNestedKey: true\n",
			want: "{\"MyKey\":{\"MyNestedKey\":true}}",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := YamlToJson([]byte(tc.yaml))
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
