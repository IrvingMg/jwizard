package formatter

import (
	"testing"
)

func TestBeautify(t *testing.T) {
	cases := []struct {
		name string
		json string
		want string
	}{
		{
			name: "OnePair",
			json: "{\"MyKey\":\"MyValue\"}",
			want: "{\n\t\"MyKey\": \"MyValue\"\n}",
		},
		{
			name: "MultiPair",
			json: "{\"MyFirstKey\":\"MyFirstValue\",\"MySecondKey\":0}",
			want: "{\n\t\"MyFirstKey\": \"MyFirstValue\",\n\t\"MySecondKey\": 0\n}",
		},
		{
			name: "Array",
			json: "{\"MyArray\":[1,2,3]}",
			want: "{\n\t\"MyArray\": [\n\t\t1,\n\t\t2,\n\t\t3\n\t]\n}",
		},
		{
			name: "NestedObject",
			json: "{\"MyKey\":{\"MyNestedKey\":true}}",
			want: "{\n\t\"MyKey\": {\n\t\t\"MyNestedKey\": true\n\t}\n}",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Beautify([]byte(tc.json))
			if err != nil {
				t.Errorf("error: %v", err)
				return
			}

			if string(got) != tc.want {
				t.Errorf("expected: %v, got: %v", tc.want, string(got))
			}
		})
	}
}

func TestMinify(t *testing.T) {
	cases := []struct {
		name string
		json string
		want string
	}{
		{
			name: "OnePair",
			want: "{\"MyKey\":\"MyValue\"}",
			json: "{\n\t\"MyKey\": \"MyValue\"\n}",
		},
		{
			name: "MultiPair",
			json: "{\n\t\"MyFirstKey\": \"MyFirstValue\",\n\t\"MySecondKey\": 0\n}",
			want: "{\"MyFirstKey\":\"MyFirstValue\",\"MySecondKey\":0}",
		},
		{
			name: "Array",
			json: "{\n\t\"MyArray\": [\n\t\t1,\n\t\t2,\n\t\t3\n\t]\n}",
			want: "{\"MyArray\":[1,2,3]}",
		},
		{
			name: "NestedObject",
			json: "{\n\t\"MyKey\": {\n\t\t\"MyNestedKey\": true\n\t}\n}",
			want: "{\"MyKey\":{\"MyNestedKey\":true}}",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Minify([]byte(tc.json))
			if err != nil {
				t.Errorf("error: %v", err)
				return
			}

			if string(got) != tc.want {
				t.Errorf("expected: %v, got: %v", tc.want, string(got))
			}
		})
	}
}
