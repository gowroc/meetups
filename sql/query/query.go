// Code generated by go-bindata.
// sources:
// sql/queries/users.sql
// DO NOT EDIT!

package query

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _sqlQueriesUsersSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x4f\x83\x30\x14\x87\xef\xfd\x2b\x7e\x07\x0e\x60\x3a\x92\xe9\xcd\xc5\x03\x93\x37\xac\x41\x58\x28\xb8\x78\x5a\x88\xbe\xe0\x0c\x32\x42\xdb\x83\xff\xbd\x19\x6c\xba\x68\x4c\xbc\xb6\x5f\xbf\xf7\xf5\xcd\x66\xe8\xea\x77\xbe\x86\xe6\x96\x9f\x6d\xd4\xb6\x42\x53\x4a\xb7\x25\x9c\xe1\x61\xbb\x7b\x91\xe3\xbd\x44\xdd\x30\x56\x45\xfe\x80\x4a\x53\xa1\x17\x42\x7c\xbd\x54\x9d\xe1\xc1\x0a\x95\x69\x2a\x4a\xa8\xac\xcc\x27\xc6\xff\x6d\x08\xf0\x18\xa5\x15\x69\xf8\xde\x5c\xc2\xbb\x94\xf0\xae\x82\x73\x57\xcc\x2d\x5b\x3e\x54\xc4\x94\x52\x49\x7f\x8c\x4c\xd8\x2e\x3f\x54\xfc\xaf\x54\x6c\xee\xa8\xa0\x13\x83\x1b\x78\xf3\x1f\xaa\xca\xf0\xb0\xd9\xd9\xd7\xf5\xde\x58\x73\x74\x0a\x17\x5e\x48\xbc\x99\x7d\xb7\xad\x9b\xc6\x77\x7d\x80\x48\xa3\x1f\x89\xd1\x7d\xf0\x19\x38\x91\xd2\xaa\xc4\x7d\xae\xb2\x69\xc2\x48\xc0\xf5\xa8\xb4\xca\x12\x9c\x76\x10\x88\x63\x45\x78\xde\x21\x92\x22\xaf\xd6\x58\x3e\x7d\x9f\x4b\xb8\x70\xfa\x85\x0b\xeb\x86\x17\xe2\x33\x00\x00\xff\xff\xe0\x15\x7c\x15\xa1\x01\x00\x00")

func sqlQueriesUsersSqlBytes() ([]byte, error) {
	return bindataRead(
		_sqlQueriesUsersSql,
		"sql/queries/users.sql",
	)
}

func sqlQueriesUsersSql() (*asset, error) {
	bytes, err := sqlQueriesUsersSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql/queries/users.sql", size: 417, mode: os.FileMode(420), modTime: time.Unix(1494362710, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"sql/queries/users.sql": sqlQueriesUsersSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"sql": &bintree{nil, map[string]*bintree{
		"queries": &bintree{nil, map[string]*bintree{
			"users.sql": &bintree{sqlQueriesUsersSql, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
