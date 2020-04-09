// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DOCS: https://support.hdfgroup.org/HDF5/doc/HL/RM_HDF5Optimized.html

package hdf5

// #include "hdf5.h"
// #include "hdf5_hl.h"
import "C"

// Regular write
// herr_t H5Dwrite( hid_t dataset_id, hid_t mem_type_id, hid_t mem_space_id, hid_t file_space_id, hid_t xfer_plist_id, const void * buf
// Chunk write
// herr_t H5DOwrite_chunk( hid_t dset_id, hid_t dxpl_id, uint32_t filter_mask, hsize_t *offset, size_t data_size, const void *buf )

// Reads a raw data chunk directly from a dataset in a file into a buffer.
/*
hid_t dset_id	        IN: Identifier for the dataset to write to
hid_t dxpl_id	        IN: Transfer property list identifier for this I/O operation
uint32_t filter_mask   	IN: Mask for identifying the filters in use
hsize_t *offset	        IN: Logical position of the chunk’s first element in the dataspace
size_t data_size	    IN: Size of the actual data to be written in bytes
const void *buf	        IN: Buffer containing data to be written to the chunk

*/
func (s *Dataset) WriteChunk(data interface{}, memspace, filespace *Dataspace) error {

	C.H5DOwrite_chunk(s.id, C.H5P_DEFAULT, 0, <offset>, <data-size>, <data/buffer>)
	return nil
}

// WriteSubset(data interface{}, memspace, filespace *Dataspace) error
// rc := C.H5Dwrite(s.id, dtype.id, memspace_id, filespace_id, 0, addr)
