package hdf5

import (
	"os"
	"testing"
)

func TestChunkWrite(t *testing.T) {
	DisplayErrors(true)
	defer DisplayErrors(false)
	defer os.Remove(fname)

	fdims := []uint{12, 4, 6}
	fspace, err := CreateSimpleDataspace(fdims, nil)
	if err != nil {
		t.Fatal(err)
	}
	mdims := []uint{2, 6}
	mspace, err := CreateSimpleDataspace(mdims, nil)
	if err != nil {
		t.Fatal(err)
	}

	f, err := CreateFile(fname, F_ACC_TRUNC)
	if err != nil {
		t.Fatalf("CreateFile failed: %s\n", err)
	}
	defer f.Close()

	dset, err := f.CreateDataset("dset", T_NATIVE_USHORT, fspace)
	if err != nil {
		t.Fatal(err)
	}

	offset := []uint{6, 0, 0}
	stride := []uint{3, 1, 1}
	count := []uint{mdims[0], 1, mdims[1]}
	if err = fspace.SelectHyperslab(offset, stride, count, nil); err != nil {
		t.Fatal(err)
	}

	data := make([]uint16, mdims[0]*mdims[1])

	createDataset
	WriteDirect
}
