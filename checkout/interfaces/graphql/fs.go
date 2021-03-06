// Package graphql Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// schema.graphql
package graphql

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x51\x6b\xdb\x30\x10\x7e\xf7\xaf\x50\xe8\xcb\xb6\x87\xfd\x00\x43\x1f\x3a\xdb\x1b\x66\x6b\x62\xec\xb5\x30\xca\x08\x57\xe9\xda\x88\xc9\x52\x38\x9d\xd7\x85\x91\xff\x3e\x2c\xab\x89\xed\x79\x61\x86\x80\xa4\xbb\xfb\xbe\xfb\xbe\x93\xc2\x87\x3d\x8a\xcc\xb5\x2d\x92\xc4\x6d\xb6\x43\xf9\xc3\x75\xbc\xad\x0c\x48\xdc\x90\x42\xaa\xd1\x77\x86\xc5\xef\x44\x08\x21\xae\x44\xdc\x4a\xb0\xe2\x11\x85\xb3\xfd\x4f\xf0\x0e\xc5\xde\x79\xaf\x1f\x0d\x0a\xcf\xc0\x9d\x0f\xe9\xc3\x32\x5d\xc0\x6f\x18\x88\xcf\x24\xfd\x9e\x3b\xbf\x8a\x24\xe5\x53\x2c\x15\xd7\xd7\xa2\x5c\xdf\xdf\x7c\x29\x73\x21\x81\xf8\x1e\x8c\x56\xc0\xda\xd9\xd7\x3e\x9c\x65\xd0\xd6\x87\x16\x14\x32\x68\x33\x50\x2f\x65\x8f\x1b\x01\xe2\xed\x3c\x1e\xd9\xf5\x98\x7d\x53\xe7\x45\xdd\xdc\x65\x59\xd1\x34\x21\xee\xfa\x7e\x9b\x4e\x4a\xf4\x3e\x07\x86\x25\x75\x9b\x59\xce\x00\xbc\x87\x43\x8b\x96\x3f\x1a\xf7\x92\x86\x13\x24\x72\xb4\x54\x5f\xf4\x81\xe4\x98\x24\x68\xbb\xf6\xbf\xdd\x8b\x33\x8a\x7e\x85\x75\x51\xd7\x9b\x3a\xac\xfe\xd2\x51\xdd\x7c\xbb\x2d\xd6\x5f\xab\x62\x9d\x97\xeb\x4f\x3d\xd9\x3f\xae\xc2\x5c\x4c\xa4\x89\x6a\x4a\xfb\xe4\x7c\x2a\xe2\xf7\x70\xe9\x26\x55\xe7\x8a\xd5\xf7\x01\xa3\x8f\xa9\x10\x3c\xe3\x3c\x4c\x87\x54\x4d\x73\x62\x25\xb6\xa0\xcd\x89\x36\x7e\x0d\x93\xb6\xcf\xab\x11\x74\x8e\xd2\x11\x30\xaa\x1e\x6a\xe4\xf4\xe4\x7c\x75\x41\x7c\x98\x44\x54\xac\x7d\x54\x50\x0c\x73\xfb\xe0\x9c\x41\xb0\xab\xf3\x2c\x3f\xe3\x21\x9d\xf5\xf2\x0a\x7d\xf1\x8d\x8d\x9c\x89\x5c\xcf\xc0\xf8\x02\x87\x74\x59\xdb\x90\x5e\x91\xfb\xa9\x15\x52\x3a\x09\xb6\xc8\x3b\xa7\xd2\x65\x57\xae\x24\xa1\xd2\x9c\x01\xa9\x40\x26\xde\x05\x9f\xc2\x9d\x7e\x9f\x4d\x62\x21\x1f\x5a\xd7\x59\x1e\x81\x9d\x54\x54\xa4\x25\x0e\xa0\xac\xd9\x60\xba\x30\x87\x63\x92\x24\xf8\x8b\xd1\x2a\x11\x3c\xb8\xed\x38\xbc\xb6\xa8\xf1\x92\x23\x6f\x08\xb9\x23\x7b\x47\xe6\xa4\xee\xed\xd2\x4b\x99\xff\x4f\x25\xc7\xe4\x4f\x00\x00\x00\xff\xff\xc5\x4d\x07\x78\xd2\x04\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"schema.graphql": schemaGraphql,
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
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
