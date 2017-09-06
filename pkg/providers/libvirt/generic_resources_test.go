package libvirt

import (
	"testing"
)

func TestNewStoragePool(t *testing.T) {
	if newSP := newStoragePool(); newSP == nil {
		t.Logf("Expected new storage pool definition to contain a structure")
	}
}

func TestNewNetwork(t *testing.T) {
	if newNet := newNetwork(); newNet == nil {
		t.Logf("Expected new network definition to contain a structure")
	}
}
