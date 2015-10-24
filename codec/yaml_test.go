package codec

import (
	"io/ioutil"
	"path"
	"testing"
)

const testdata = "../testdata"

func TestYamlDecoder(t *testing.T) {
	d, err := ioutil.ReadFile(path.Join(testdata, "pod.yaml"))
	if err != nil {
		t.Error(err)
	}

	m, err := YAML(d).One()
	if err != nil {
		t.Error(err)
	}

	ref, err := m.Ref()
	if err != nil {
		t.Errorf("Could not get reference: %s", err)
	}
	if ref.Kind != "Pod" {
		t.Errorf("Expected a pod, got a %s", ref.Kind)
	}
	if ref.APIVersion != "v1" {
		t.Errorf("Expected v1, got %s", ref.APIVersion)
	}
}
