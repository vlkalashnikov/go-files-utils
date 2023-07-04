package files

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"reflect"
)

// JSONEqual compares the JSON from two Readers.
func JSONEqual(a, b io.Reader) (bool, error) {
	var j, j2 interface{}
	d := json.NewDecoder(a)
	if err := d.Decode(&j); err != nil {
		return false, err
	}
	d = json.NewDecoder(b)
	if err := d.Decode(&j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

// JSONBytesEqual compares the JSON in two byte slices.
func JSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

// UnmarshalFile unmarshals a JSON file into a struct.
func UnmarshalJSON(path string, respStruct interface{}) error {
	byteValue, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(byteValue), respStruct)
}

// SaveJSON saves a struct as a JSON file.
func SaveJSON(path string, data any) error {
	bt, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, bt, 0644)
}
