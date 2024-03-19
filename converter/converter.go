package converter

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func JsonToYaml(data []byte) ([]byte, error) {
	v := make(map[string]any)

	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	return yaml.Marshal(&v)
}
