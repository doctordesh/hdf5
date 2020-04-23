// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DOCS: https://support.hdfgroup.org/HDF5/doc/HL/RM_HDF5Optimized.html

package hdf5

// #include "hdf5.h"
// #include "hdf5_hl.h"
import "C"
import (
	"reflect"
	"unsafe"
)

// WriteChunk writes a raw data chunk from a buffer directly to a dataset in a file.
// It requires that the dataset is chunked
func (s *Dataset) WriteChunk(data *[]byte, offset []uint) error {
	// Offset and size should be quite easy
	c_offset := (*C.hsize_t)(unsafe.Pointer(&offset[0]))
	c_size := C.size_t(len(*data))

	// Make C of data
	v := reflect.Indirect(reflect.ValueOf(data))
	slice := (*reflect.SliceHeader)(unsafe.Pointer(v.UnsafeAddr()))
	addr := unsafe.Pointer(slice.Data)

	rc := C.H5DOwrite_chunk(s.id, C.H5P_DEFAULT, 0, c_offset, c_size, addr)

	return h5err(C.herr_t(rc))
}

// ReadChunk reads a raw data chunk directly from a dataset in a file into a buffer.
// It requires that the dataset is chunked.
func (s *Dataset) ReadChunk(offset []uint) ([]byte, uint32, error) {

	// Three step process
	// 1. Figure out number of bytes in chunk
	// 2. Prepare Go slice to fit chunk
	// 3. Read chunk into go slice

	var err error
	var data []byte
	var filters uint32

	// Number of bytes in chunk

	nbytes, err := s.GetChunkStorageSize(offset)
	if err != nil {
		return data, filters, err
	}

	// Prepare Go slice (and other data)
	c_offset := (*C.hsize_t)(unsafe.Pointer(&offset[0]))
	data = make([]byte, nbytes)
	v := reflect.Indirect(reflect.ValueOf(&data))
	slice := (*reflect.SliceHeader)(unsafe.Pointer(v.UnsafeAddr()))
	addr := unsafe.Pointer(slice.Data)
	c_filters := (*C.uint32_t)(unsafe.Pointer(&filters))

	// Read chunk into go slice
	rc := C.H5DOread_chunk(s.id, C.H5P_DEFAULT, c_offset, c_filters, addr)
	err = h5err(C.herr_t(rc))
	if err != nil {
		return data, filters, err
	}

	return data, filters, err
}

/*
Extend dataset: https://support.hdfgroup.org/HDF5/doc/RM/RM_H5D.html#Dataset-SetExtent
*/
