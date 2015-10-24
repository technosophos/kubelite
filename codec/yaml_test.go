package codec

import (
	"io/ioutil"
	"path"
	"testing"
)

const testdata = "../testdata"

func TestYamlDecoderOne(t *testing.T) {
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

func TestYamlDecoderAll(t *testing.T) {
	d, err := ioutil.ReadFile(path.Join(testdata, "three-pods.yaml"))
	if err != nil {
		t.Error(err)
	}

	ms, err := YAML(d).All()
	if err != nil {
		t.Error(err)
	}

	if len(ms) != 3 {
		t.Errorf("Expected 3 pods, got %d", len(ms))
	}

	ref, err := ms[2].Ref()
	if err != nil {
		t.Errorf("Expected a reference for pod[2]: %s", err)
	}

	if ref.Kind != "Pod" {
		t.Errorf("Expected Pod, got %s", ref.Kind)
	}
}
