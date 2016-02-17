// Code generated by go-bindata.
// sources:
// templates/default.bra.toml
// DO NOT EDIT!

package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _templatesDefaultBraToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x52\x3d\x6f\xdc\x30\x0c\x9d\xa3\x5f\x41\x28\x1d\xd3\xf4\x3a\xdc\x52\xa0\x43\xd0\x64\x4c\x11\x14\x01\x3a\x1c\x04\xc1\x91\x68\x97\xa8\x3e\x0e\x12\x95\xfa\xfe\x7d\x69\xd9\x87\xeb\xc7\x1c\x2f\x26\x9e\xde\x7b\x7c\x14\x75\x28\x2d\x19\x45\x89\xd8\xba\xe8\x2b\x7c\x86\x83\xba\x3a\xe8\x29\xeb\x1b\xd0\x94\x2a\x0f\x21\x68\x73\x73\xc1\x5e\x1a\x05\xbf\x21\xb7\x1f\xde\xdd\x3d\x3d\xd9\xaf\x77\x8f\x0f\xda\x28\x73\x75\xfe\xae\xe1\x4b\x8e\x71\x48\xe2\x27\xf6\x40\x09\xc4\xa7\xb0\xfa\x35\xb0\xfb\x61\xc5\x51\xda\x70\x69\xb8\x91\xbf\x2f\x30\x2c\x70\x6d\x2f\xef\x3d\x15\x74\x9c\x0b\x61\xdd\x04\x82\xf4\x60\xe6\xec\x7e\x7f\xa1\x00\x67\xe8\xac\x8d\x8b\x33\x77\xae\xbe\x95\xb8\x66\xa5\x3f\xcc\x8c\xa9\x52\x4e\x7f\xb0\x69\x4a\xb9\xe0\xc6\x24\x5e\x26\x4b\xd9\xa3\x8d\xd9\xb7\x80\x55\x1b\xf8\xbf\x0f\xce\x2e\x34\x8f\x30\x96\x1c\x57\x1b\x4a\xd3\xe6\x64\x47\x12\xd9\x25\xe5\x35\x7c\xc3\x09\xe7\x63\x85\x31\x17\xe8\x1c\x21\x43\x3d\xa2\xa3\x91\x1c\xa4\xcc\xf2\x97\x11\xfb\x7d\x5a\x8f\x61\x38\x89\xfa\xe3\x7e\xb7\xdb\xf4\x8f\xb2\x94\x38\x04\xb9\x3d\xc6\xf2\x2a\x85\x24\x78\x2e\x34\x4d\x58\xa0\x8b\x00\x5f\x31\xb1\xea\xe7\xa5\x1d\xd9\x32\xc5\xdc\xb8\xbb\xac\x16\xcf\x14\x71\x1d\x99\x18\x5a\x62\x0a\x4b\x1a\x87\xf0\x93\x42\x50\xbd\xb4\x4b\xf9\xf7\x3a\xee\xf3\x92\x6e\x11\x75\xba\x8c\x2d\x6a\x59\x26\xac\x9b\x09\xa7\x55\xfe\x26\xcf\x45\xc2\xca\x8b\x51\xea\x50\x4f\xc9\x19\x15\xa8\xca\xea\xec\xe0\x7d\x91\x56\xfa\xd3\x7e\xb7\xdf\x69\x55\x30\x66\xc6\x7f\xd0\xdf\x01\x00\x00\xff\xff\xd2\x1b\x16\x28\xc9\x02\x00\x00")

func templatesDefaultBraTomlBytes() ([]byte, error) {
	return bindataRead(
		_templatesDefaultBraToml,
		"templates/default.bra.toml",
	)
}

func templatesDefaultBraToml() (*asset, error) {
	bytes, err := templatesDefaultBraTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/default.bra.toml", size: 713, mode: os.FileMode(420), modTime: time.Unix(1455748370, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"templates/default.bra.toml": templatesDefaultBraToml,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"default.bra.toml": &bintree{templatesDefaultBraToml, map[string]*bintree{
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

