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

/*
Signature: herr_t H5DOread_chunk( hid_t dset_id, hid_t dxpl_id, const hsize_t *offset, uint32_t *filter_mask, void *buf )
Parameters:
hid_t dset_id				IN: Identifier for the dataset to be read
hid_t dxpl_id				IN: Transfer property list identifier for this I/O operation
const hsize_t *offset		IN: Logical position of the chunk’s first element in the dataspace
uint32_t * filter_mask   	IN: Mask for identifying the filters used with the chunk
void *buf					IN: Buffer containing the chunk read from the dataset

*/
// ReadChunk
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

	// 3. Read chunk into go slice
	rc := C.H5DOread_chunk(s.id, C.H5P_DEFAULT, c_offset, c_filters, addr)
	err = h5err(C.herr_t(rc))
	if err != nil {
		return data, filters, err
	}

	return data, filters, err
}

// WriteSubset(data interface{}, memspace, filespace *Dataspace) error
// rc := C.H5Dwrite(s.id, dtype.id, memspace_id, filespace_id, 0, addr)

/*
Extend dataset: https://support.hdfgroup.org/HDF5/doc/RM/RM_H5D.html#Dataset-SetExtent
Set chunk size: https://support.hdfgroup.org/HDF5/doc/RM/RM_H5P.html#Property-SetChunk

Documentation on chunks: https://support.hdfgroup.org/HDF5/doc/Advanced/Chunking/Chunking_Tutorial_EOS13_2009.pdf
*/
