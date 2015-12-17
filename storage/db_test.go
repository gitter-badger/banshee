// Copyright 2015 Eleme Inc. All rights reserved.

package storage

import (
	"fmt"
	"github.com/eleme/banshee/util"
	"github.com/eleme/banshee/util/assert"
	"os"
	"path"
	"testing"
)

func TestOpen(t *testing.T) {
	fileName := "storage_test"
	options := &Options{NumGrid: 288, GridLen: 300}
	db, err := Open(fileName, options)
	assert.Ok(t, err == nil)
	assert.Ok(t, db != nil)
	defer db.Close()
	defer os.RemoveAll(fileName)
	assert.Ok(t, util.IsFileExist(path.Join(fileName, adbFileName)))
	assert.Ok(t, util.IsFileExist(path.Join(fileName, mdbFileName)))
	sFileName := fmt.Sprintf("%s-%dx%d", sdbFileName, options.NumGrid, options.GridLen)
	assert.Ok(t, util.IsFileExist(path.Join(fileName, sFileName)))
}