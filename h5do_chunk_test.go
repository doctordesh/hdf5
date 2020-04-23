package hdf5

import (
	"encoding/binary"
	"fmt"
	"os"
	"testing"
)

func TestWriteChunkBytes(t *testing.T) {
	filename := "chunk_bytes.h5"
	DisplayErrors(true)
	defer DisplayErrors(false)
	defer os.Remove(filename)

	fileDims := []uint{5, 3, 10}
	fspace, err := CreateSimpleDataspace(fileDims, nil)
	if err != nil {
		t.Fatal(err)
	}

	f, err := CreateFile(filename, F_ACC_TRUNC)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	dtype, err := NewDatatypeFromValue(byte(0))
	if err != nil {
		t.Fatal(err)
	}

	chunkProps, err := NewPropList(P_DATASET_CREATE)
	if err != nil {
		t.Fatal(err)
	}
	defer chunkProps.Close()

	chunkDims := []uint{1, fileDims[1], fileDims[2]}
	err = chunkProps.SetChunk(chunkDims)
	if err != nil {
		t.Fatal(err)
	}

	dset, err := f.CreateDatasetWith("dset", dtype, fspace, chunkProps)
	if err != nil {
		t.Fatal(err)
	}

	indexedWrite := func(k uint) {
		offset := []uint{k, 0, 0}

		data := make([]byte, fileDims[1]*fileDims[2])
		for i := 0; i < len(data); i++ {
			if i > len(data)/2 {
				break
			}
			data[i] = byte(k + (k * k))
		}

		err := dset.WriteChunk(&data, offset)
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < int(fileDims[0]); i++ {
		indexedWrite(uint(i))
	}
}

func TestWriteChunkUInt32(t *testing.T) {
	filename := "chunk_uint32.h5"
	DisplayErrors(true)
	defer DisplayErrors(false)
	defer os.Remove(filename)

	fileDims := []uint{5, 3, 10}
	fspace, err := CreateSimpleDataspace(fileDims, nil)
	if err != nil {
		t.Fatal(err)
	}

	f, err := CreateFile(filename, F_ACC_TRUNC)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	dtype, err := NewDatatypeFromValue(uint32(0))
	if err != nil {
		t.Fatal(err)
	}

	chunkProps, err := NewPropList(P_DATASET_CREATE)
	if err != nil {
		t.Fatal(err)
	}
	defer chunkProps.Close()

	chunkDims := []uint{1, fileDims[1], fileDims[2]}
	err = chunkProps.SetChunk(chunkDims)
	if err != nil {
		t.Fatal(err)
	}

	dset, err := f.CreateDatasetWith("dset", dtype, fspace, chunkProps)
	if err != nil {
		t.Fatal(err)
	}

	indexedWrite := func(k uint) {
		offset := []uint{k, 0, 0}

		data := make([]byte, fileDims[1]*fileDims[2]*4)
		b := make([]byte, 4)
		for i := 0; i < len(data); i = i + 4 {
			binary.LittleEndian.PutUint32(b, uint32(k+1)+uint32(i))
			data[i+0] = b[0]
			data[i+1] = b[1]
			data[i+2] = b[2]
			data[i+3] = b[3]
		}

		err := dset.WriteChunk(&data, offset)
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < int(fileDims[0]); i++ {
		indexedWrite(uint(i))
	}
}
