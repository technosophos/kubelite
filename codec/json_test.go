package codec

import (
	"io/ioutil"
	"path"
	"testing"
)

func TestJsonDecoderOne(t *testing.T) {
	d, err := ioutil.ReadFile(path.Join(testdata, "policy.json"))
	if err != nil {
		t.Error(err)
	}

	m, err := JSON.Decode(d).One()
	if err != nil {
		t.Error(err)
	}

	ref, err := m.Ref()
	if err != nil {
		t.Errorf("Could not get reference: %s", err)
	}
	if ref.Kind != "Policy" {
		t.Errorf("Expected a Policy, got a %s", ref.Kind)
	}
	if ref.APIVersion != "v1" {
		t.Errorf("Expected v1, got %s", ref.APIVersion)
	}
}

func TestJsonDecoderAll(t *testing.T) {
	data := `{"one": "hello"}
{"two": "world"}`

	ms, err := JSON.Decode([]byte(data)).All()
	if err != nil {
		t.Errorf("Failed to parse multiple JSON entries: %s", err)
	}
	if len(ms) != 2 {
		t.Errorf("Expected 2 JSON items, got %d", len(ms))
	}
}
