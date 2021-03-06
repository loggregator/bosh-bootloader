// Code generated by go-bindata.
// sources:
// templates/cf_dns.tf
// templates/cf_lb.tf
// templates/concourse_lb.tf
// templates/network.tf
// templates/network_security_group.tf
// templates/output.tf
// templates/resource_group.tf
// templates/storage.tf
// templates/tls.tf
// templates/vars.tf
// DO NOT EDIT!

package azure

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

var _templatesCf_dnsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x92\xc1\x6a\xc3\x30\x10\x44\xef\xfe\x8a\x45\xe4\x54\xb0\x08\xf4\x9c\x2f\x09\x65\x59\x4b\xeb\x54\x60\x4b\x42\x92\x53\x12\xe3\x7f\x2f\x52\x70\xd3\x28\x34\x2d\xe4\x54\xdf\x8c\x67\xd6\x6f\x66\x57\x53\x22\x10\x74\x9e\x02\x87\x11\xfd\xd4\x0d\x46\xa1\xf1\x02\x84\xea\xdb\xa1\x13\x30\x37\x00\x96\x46\x86\xea\xd9\x81\xd8\xcc\x47\x0a\x92\xed\x11\x8d\x5e\xda\xa2\x6f\x8d\x17\x0d\x40\xe0\xe8\xa6\xa0\x18\x0f\xc1\x4d\x1e\x8b\xbf\x18\xd6\x1f\xdd\x0a\x64\xe7\xe2\xbb\xcc\xaa\x25\xbb\x35\x7b\xb6\x3a\xa2\xb3\xb0\x83\xfd\x17\x1c\x79\x3f\x18\x45\xc9\x38\x8b\x07\x4a\xfc\x41\x27\xa9\x7a\xf1\xd6\x2c\x4d\xb3\xce\xbb\x46\xd1\x36\xe2\xd9\x59\x2e\x49\xfe\x1e\x43\x5e\xde\xe2\x29\x26\x1e\x51\xbb\x91\x8c\x5d\x9e\x8d\xf4\x23\x20\x61\x60\xe5\x82\xfe\x1d\xf2\x25\x33\xe4\x3c\x58\x09\x6e\x18\xd6\xd0\x52\xf5\xd7\x3a\x9f\x5b\x46\x4a\x43\x8d\x93\xdd\xaf\xdb\xed\x65\x76\xc6\x8f\xd5\xd7\xbd\xd8\xcc\xf9\xb0\xe4\xdd\x5d\xc9\x72\x25\xd2\x78\x24\xad\x03\xc7\xb8\x3c\xda\xdf\xb7\x7a\x32\xd4\xc3\x82\x8a\xe0\x9f\x75\x74\x5f\x4f\x99\x5b\xb7\xf3\x19\x00\x00\xff\xff\x30\xb2\xa5\x8e\xa4\x03\x00\x00")

func templatesCf_dnsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_dnsTf,
		"templates/cf_dns.tf",
	)
}

func templatesCf_dnsTf() (*asset, error) {
	bytes, err := templatesCf_dnsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_dns.tf", size: 932, mode: os.FileMode(480), modTime: time.Unix(1521569332, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xcd\x6e\xe3\x36\x10\xbe\xeb\x29\x08\xa1\x87\xb6\x58\x39\x4e\x62\x14\xb9\xe8\xd0\xa2\x87\xf6\xbc\xbd\x13\x94\x34\xb2\x89\xd0\x24\xcb\x1f\x67\xdd\x85\xdf\xbd\x20\x29\xc9\x16\x4d\xd9\x72\x8a\xdd\xcd\xa2\xf1\xcd\x9a\x1f\x8e\xbe\xf9\x66\x38\xa3\x1d\x51\x94\x54\x0c\x50\xae\xf7\xda\xc0\x16\x37\x62\x4b\x28\xcf\xd1\xe7\x43\x96\x1d\x85\xb2\xfd\x84\x6b\x50\x06\x57\x44\xc3\x2f\xab\x94\x58\x12\xad\x5f\x84\x6a\x82\x4c\x81\x16\x56\xd5\x80\x72\xf2\x8f\x55\xa0\xb6\x58\xdb\x8a\x83\xc9\x51\x5e\xb7\x85\x76\x07\x64\x08\x71\xb2\x05\x14\xff\x4a\x94\xff\xf0\x79\x47\xd4\x02\xf8\x0e\xd3\xe6\x50\x04\x83\x0c\x21\xd2\x34\x0a\xb4\xc6\x52\x41\x4b\x3f\x9d\xaa\xd7\xb4\x51\xe1\x80\x1f\x9d\x25\x07\xf3\x22\xd4\x33\x76\x8f\x3f\xa0\xa7\x0f\xe8\xfe\xa7\x83\x73\xd0\x47\x85\xd7\x4a\x58\x89\xc3\xf1\xde\x41\x1f\xe5\x58\x63\x51\x09\xbd\x59\x38\x35\x6f\xbe\xa3\xca\x58\xc2\x70\xef\xde\xdb\x8f\xcc\x23\x8d\x91\x7d\x12\x95\xde\x95\x86\xda\x2a\x6a\xf6\xe1\x5c\x8f\xd2\x34\x44\x09\x84\x5c\x78\x4c\xd4\xc4\x50\xc1\x93\xaa\x0a\xd6\x54\xf0\x49\x14\xe6\x82\x90\x21\x64\xc8\x5a\xfb\xd0\x10\x02\xbe\xa3\x4a\xf0\x2d\x70\x73\x16\x94\x3b\xe9\x30\xf3\xa5\x95\x65\x10\x98\xb1\x31\x46\x5e\xe0\xc6\x05\x8a\x78\xcb\x0c\x21\xa9\xa8\x70\x4e\xd3\x76\x0f\xcb\xfb\x0c\xa1\x86\x2a\xa8\x63\xa8\x8e\xbe\xff\xe4\x95\xb0\xbc\xf1\x94\xab\x6b\xd0\x7a\x32\x8a\x5f\x19\x13\x2f\xe1\x54\x61\x44\x2d\xd8\x84\xde\x5f\xb5\x8f\xad\x83\x55\x0a\x65\xb0\x22\x7c\x0d\x63\xad\x9f\x9d\x4e\x03\xda\x50\xee\x13\x79\xa6\x58\xa2\xfc\x69\x79\xe2\x68\xaa\x20\xce\x1c\xc5\x8a\xbd\x4e\xb2\x20\x4e\x41\x9e\x55\x17\x69\x12\x27\x88\x95\x56\x5c\xd4\xed\x6d\x35\x32\xa6\x8b\x7e\x3d\x5f\xf4\x1c\xc2\x3c\x7c\xdf\x84\x59\xad\x1e\xdf\x19\x73\x64\x0c\x13\xeb\xd7\xf1\xc5\x19\xce\x60\xcb\xe3\xf7\xce\x96\xff\x21\x5d\xa4\xad\x18\xad\x31\xbd\x76\xef\x5e\xe6\x47\x55\x50\x39\x75\x0d\x9f\x5b\x5e\xb9\x8f\x5f\x01\xd2\xf0\x16\x43\x32\x08\x1b\x62\x29\x51\xde\xec\x39\xd9\xd2\x7a\x02\x03\x22\x25\xa3\x41\x19\xaf\x89\x81\x17\xb2\xbf\x75\x0a\x21\x52\x16\xbd\xe9\x7f\x1c\x33\xe6\x0e\x33\x8e\xa9\xcf\xb6\x9b\x46\x86\x20\x4b\x94\x7f\x34\x84\x37\x44\x35\xf8\xe3\x96\x30\x96\x7b\xb9\xa1\xa0\x62\x79\x90\xd4\x44\x92\xda\x15\x75\x89\x1e\xfc\xd8\x12\x4a\xae\x82\xd8\xf3\x38\x98\x0d\x10\x66\x36\x85\xd7\x0c\x8e\x52\x75\x5a\xa2\xfc\x8f\x6e\x36\x41\x48\x12\xb3\x49\x78\xba\x0b\xd2\x8d\xd0\x26\x21\x25\x92\x2e\x46\x50\x77\xff\x46\x23\xfb\x21\xb8\xa0\xdc\x80\xda\x91\x28\x82\xc7\x65\x87\xc0\x16\x84\x35\x28\x29\xb4\x3c\xbc\xcf\x1e\x9b\x8d\x02\xbd\x11\xac\x71\xc2\x1e\x8f\x2e\xb3\x8e\x5f\xb5\xe0\x2d\x5d\x5b\x15\x12\x14\x43\x94\x2a\x8e\xce\xb8\xa0\xb2\x18\x19\x87\x98\xc3\xe4\x8e\x69\x33\x63\x98\xa6\xcd\xe1\x2e\xe8\xeb\xbb\xa3\x6a\x78\xb2\xf0\x9b\xc2\x91\x43\x3e\xee\x56\x09\x6e\x80\x37\xbe\xdb\x9d\x06\x5b\xa2\xbc\x97\x39\xd1\x30\x0f\x20\xe4\x35\x4b\xb4\x5a\x3d\xbe\xc6\xc9\xc8\xc7\xd3\xf2\x56\x17\x4c\xac\xe3\x30\x12\x71\x5c\xcd\xc2\xa5\x42\xad\xdb\xa2\x77\x34\x91\x91\xf3\x5e\x12\x27\x67\xd0\x70\xcd\x75\x18\xf7\x33\x84\x2a\x52\x3f\xbb\x08\x87\x1b\x41\x08\x16\xbd\xee\x59\x34\x9d\x4d\xd1\xd9\x14\xce\xe6\xcc\xa1\xc3\x16\x6b\x30\x86\xf2\x61\x01\x49\x77\xe9\x99\x2b\x59\x51\x41\xb1\x31\xda\x74\x3d\x40\x88\x67\x0a\x7e\xc1\x6d\x30\x69\x5b\xca\x43\x43\xc8\x7f\xa7\xda\x6d\xb9\xcd\x49\x52\x12\x27\x3e\x2d\x27\x1b\x40\xd4\x02\x14\xfc\x6d\x41\x1b\x3c\x2e\xc6\x12\xdd\x0f\x1e\x2a\xc0\xf1\x8b\x25\xba\x8d\x07\x47\x6b\xe6\x17\x73\xda\xba\xe6\x7d\xd6\xaf\x4a\x94\x6b\xcd\x0a\xa7\x11\x0e\x6f\x88\x21\x63\x52\x44\xab\xfd\xa1\xef\x53\x61\x9b\x1f\xeb\xf5\x4f\x8f\xd9\xf6\x49\x61\x54\x1b\xe0\xa0\x2e\x26\xe5\xe6\xec\x38\xd7\x05\xd3\xa6\xa3\xe4\x24\xf5\xf1\x24\xad\xae\x90\x7c\x54\x91\x67\x88\x5f\x28\xed\x0b\xc3\xd8\x38\xd9\x5f\x03\x22\xfd\xc6\x30\xd2\xb7\x80\xd4\x29\x47\x2c\x8e\xcf\x89\x58\xfc\xa5\x51\x75\x1d\xf8\x0d\x81\x7a\x72\x21\x7c\x61\x4c\xfb\xce\xa4\x84\x75\x4d\xd6\x6f\x4d\xd7\xa1\xbd\xa9\x9e\xfd\x22\x16\xda\xa0\x65\x80\xcd\x5e\xa6\xfb\xf7\x6f\x44\xbb\x51\xd5\xfd\x1b\x65\xfa\xc2\x06\x31\xbf\x95\xa4\xae\xa8\xd3\x8f\x81\x33\x6f\xa7\x89\xab\xe9\xa6\xcf\x82\xa7\x77\xd0\x57\xca\x81\xfe\x86\x49\xd0\xef\x59\xe8\x1b\xcc\xb7\x49\x42\xd4\xdb\xde\x6c\x0e\x0e\x59\x26\xac\x91\xd6\xb8\x1d\xd4\x2d\xa7\xfd\x52\xea\x7d\x86\xa5\x74\x47\x98\x8d\xdc\x27\xb6\xd8\xf1\xf2\x7f\xe2\x34\xfe\xe4\x3e\xe1\x72\xc6\x27\x85\x7f\x03\x00\x00\xff\xff\xb0\xbc\xae\x58\x45\x19\x00\x00")

func templatesCf_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_lbTf,
		"templates/cf_lb.tf",
	)
}

func templatesCf_lbTf() (*asset, error) {
	bytes, err := templatesCf_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_lb.tf", size: 6469, mode: os.FileMode(480), modTime: time.Unix(1521569246, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConcourse_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\x3d\x6f\xdb\x30\x10\xdd\xf9\x2b\x08\xa2\x53\x01\x09\x46\xe3\xc1\x1d\x3c\x14\x9d\xba\x75\xe8\x2e\x50\x14\xe3\x10\x61\x78\xc4\x91\x74\xda\x06\xfe\xef\x05\x25\x4b\x91\x2c\x4a\x96\x9d\x02\xad\xd1\x7a\x7e\xf7\xc8\x7b\x1f\x96\x84\xd2\x41\x40\x21\x29\xe3\x3f\x03\x4a\x7c\x2a\x6c\x28\xb5\x12\x85\xb2\x8c\x32\x01\x46\x40\x40\x27\x19\x7d\x21\x94\x1a\xfe\x24\xe9\xd4\x6f\x4b\xd9\xbb\x97\x3d\xc7\x5c\x9a\x7d\xa1\xaa\x43\xd6\x0d\x67\xba\x64\x84\x52\x0d\x82\x7b\x05\x66\x7e\x1a\xe5\x4e\x81\x39\xc4\x81\xf6\x6e\xc5\x0e\x21\xd8\x62\x78\x7a\x3d\xd0\xde\x79\x88\xcc\x4b\x70\x0f\x79\x84\xd7\x34\xdd\x42\x05\xaf\x2a\x94\xce\x15\x5c\x77\x77\xd9\x52\xe6\x3c\xf7\x4a\x30\x42\x28\xf5\x7c\xe7\xea\x4d\x29\x95\x66\xaf\x10\xcc\x93\x34\x7e\xb4\x5a\x64\x3d\x90\x03\x21\x63\xf5\x74\xb9\x4c\xb6\xb3\x6a\xa5\x96\x5f\xbe\x73\x4a\xeb\xb1\xc4\x84\xd2\x7b\x04\xe3\xa5\xa9\xa2\x3c\x02\xcc\xbd\xda\x05\x6c\x26\x1b\x19\x92\x96\xcf\x5c\xbe\xe5\xcb\x94\xcd\x06\x7c\xac\xa6\x1b\x7b\xa1\xaa\xe1\x5a\x1d\x22\xef\x48\xf3\x73\x92\x17\x18\xb4\xec\xeb\x9e\x3d\x78\x6f\xdd\x55\xea\x37\x93\xbf\xc1\x00\x5e\x95\x5c\x73\x23\x24\xc6\x1d\x47\x91\xd5\xe5\xe9\x82\x73\x6e\xf4\x8e\xbf\x42\x77\x8b\xe0\x41\x80\x9e\x6e\xee\xb7\xcf\x5f\x59\xff\x7c\x0b\xe8\x53\xc0\xf5\xfa\x8e\x50\x5a\x72\xf1\x38\x8d\x3a\xc2\x7a\xb8\xd6\x6b\x0b\xa0\x47\x86\xeb\xb2\x48\xe1\xc6\xf6\x5b\x84\x52\xb6\x5a\x9e\x9a\xd9\x63\xab\x71\xf9\x89\xa3\x0d\xc9\x54\x7e\xea\x91\x9b\x0f\x50\xda\xe8\x57\x77\x53\x76\x35\x5e\x5d\xd4\xab\xab\x55\xf9\xdf\xaa\x89\x56\x6d\x56\x4b\x4a\xb5\x59\xfd\x5d\x9d\xba\xa6\x52\xb7\x95\x9d\xeb\x0a\xb5\x59\xa5\x45\x31\xd2\x3f\x03\x3e\x16\x4e\x8a\x80\xca\xff\xb8\xb4\x5d\x17\x28\x65\x51\x41\x3c\x22\x3d\xfe\x61\xf5\x91\x50\x5a\x29\x94\x62\xe2\x7d\x6c\x4b\xd9\x17\x53\x42\x30\x55\x64\xe3\x42\x48\xe7\x26\x2f\xf3\x49\x6b\x78\x3e\x57\x88\x28\x9a\xa8\xef\x76\xb4\x28\x6a\x57\x20\x37\x3b\x39\x44\xbd\x8f\x98\x4a\x3a\xaf\x4c\xd3\xd1\x53\xe0\x96\xb2\xcd\xaa\x47\xd4\x05\x1c\xe5\xbd\xfa\x3e\x43\x74\x0a\x6c\x31\x73\xaf\x99\xcb\xe3\x35\x72\x77\x2a\xa4\x69\xe0\x80\xed\x0d\xf1\x99\x79\x66\x2d\xcc\x8f\x5b\x12\xa0\xcd\x6d\x07\x68\xbd\xbe\xfb\x07\x13\x34\xf1\x54\x78\xf3\x07\xcb\x91\x34\xab\xc9\xfe\xc4\xff\xf4\x81\x10\x08\xde\x06\xdf\xdb\x24\x6e\x1b\x39\x9b\x8d\xf6\x5c\x07\x39\x43\xf3\x2a\x5b\x92\x48\xd9\x49\x9a\xe4\xf7\x4a\xf7\x79\x53\x73\xfe\x0a\x00\x00\xff\xff\x37\xac\x0e\xab\x64\x0f\x00\x00")

func templatesConcourse_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesConcourse_lbTf,
		"templates/concourse_lb.tf",
	)
}

func templatesConcourse_lbTf() (*asset, error) {
	bytes, err := templatesConcourse_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/concourse_lb.tf", size: 3940, mode: os.FileMode(480), modTime: time.Unix(1519854618, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetworkTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\xcf\x4a\xc4\x30\x10\xc6\xef\x79\x8a\x61\xf0\xa0\xb0\x5b\x3c\x7a\xf1\x49\x44\x42\x36\x19\xd7\xe0\x36\x29\x93\x3f\x8a\x25\xef\x2e\xa9\xad\xd8\xd8\xe2\xe6\xfc\x9b\x99\xef\xf7\x85\x29\xf8\xc4\x9a\x00\xd5\x67\x62\xe2\x5e\x66\xcb\x31\xa9\x8b\x74\x14\xdf\x3d\xbf\x21\xe0\xc9\x87\x57\x84\x51\x00\x38\xd5\x13\x34\xef\x11\xf0\x66\xcc\x8a\x3b\x72\x59\x5a\x53\x8e\x15\x3f\x66\x87\x02\x40\x19\xc3\x14\x82\x0c\x83\xd2\xf4\xc3\x3f\xcd\x03\xf3\x05\xa9\xad\xe1\x82\xcf\x02\xe0\xe2\xb5\x8a\xd6\xbb\xcd\xfd\x4c\x67\xeb\x5d\xa9\x7b\x97\xd4\xf2\xcc\x3e\x0d\x72\x8a\x35\x71\x8b\xc4\x1a\xe8\x6a\xa4\xae\x52\x05\x45\x11\xe2\xaf\x74\x48\x27\x47\xf1\x5f\xd7\x1d\xd9\xb0\x92\x1d\x98\x5e\xec\xc7\xef\x81\x2a\xf8\x7d\xe1\xb6\xf5\x3e\xc0\xc3\x01\xee\xef\x76\xad\xae\xd6\x02\x68\x3e\x6e\xa3\x95\x86\x68\x6a\xf9\x0a\x00\x00\xff\xff\xdb\x9a\xf5\x1a\x0b\x02\x00\x00")

func templatesNetworkTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetworkTf,
		"templates/network.tf",
	)
}

func templatesNetworkTf() (*asset, error) {
	bytes, err := templatesNetworkTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network.tf", size: 523, mode: os.FileMode(480), modTime: time.Unix(1519854618, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesNetwork_security_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x94\xcf\x8e\x9b\x30\x10\x87\xef\x7e\x8a\x91\xd5\xd3\x4a\x1b\x6d\xd9\xb0\xe2\x92\x43\x8f\xbd\xf7\x8e\x1c\x7b\x4a\xac\x12\x0f\x1a\x9b\xa4\x6d\xc4\xbb\x57\x86\x92\x86\x84\x44\x14\x55\xaa\xb2\x0a\xe7\x6f\x7e\xf3\xc7\x9f\x60\xf4\x54\xb3\x46\x90\xea\x67\xcd\xc8\xdb\xdc\x61\xd8\x13\x7f\xcb\x3d\xea\x9a\x6d\xf8\x91\x17\x4c\x75\x25\x41\xae\xc9\x6f\x24\x1c\x04\x80\x53\x5b\x84\xb3\x6f\x05\xf2\xc3\x61\xa7\x78\x81\x6e\x97\x5b\xd3\x3c\xb7\xb8\x00\x28\x49\xab\x60\xc9\x8d\xc2\x8c\x85\x25\xd7\x44\xae\x9f\xa4\xeb\x97\xb7\x3d\x5a\xae\x1f\x6c\x08\x2c\x62\xfe\x22\x52\x8d\x14\x02\x20\xa8\xc2\xb7\xc3\x01\xa0\xdb\x59\x26\xb7\x45\x17\x2e\xc6\x8a\x9d\x1a\xd1\x08\x31\x61\x71\xae\x4b\x94\x20\xfd\xad\xb5\xaf\xae\xef\xbb\xed\x2b\xb6\x14\xc3\xc6\x6b\x92\x97\x17\x01\x60\x2c\xa3\x3e\x3f\xd1\x9f\xdc\xcf\x6e\x4d\xb5\x33\x31\x4d\x69\x8d\xde\x5f\x9d\xe0\x53\x59\xd2\xbe\xeb\x4a\x81\x34\x95\x57\xb8\x2f\xba\x8a\xd4\xef\x73\x56\xc4\x21\x67\xe5\x0a\x1c\x52\x4f\x91\x31\xe8\x83\x75\xed\x03\x5e\x80\x2b\x90\x49\x72\x12\xa4\x8c\x61\xf4\x3e\xaf\x18\xbf\xda\xef\x37\x82\xce\xc1\x9e\x19\x53\x60\x70\xe0\x09\x2a\x00\x8c\x0b\x3c\x22\xd4\x38\x38\x48\xfb\x2b\x51\x62\xe1\xb3\x2a\xd0\x85\x19\xbe\x9c\x14\x4f\xd0\xe6\xe3\x7d\x6b\xf3\x96\xbd\x65\x0f\x71\x86\xe2\x74\xcf\x49\x3c\xd7\x9d\x63\xfd\x04\x7d\x92\xfb\xd6\x27\x49\xd3\x34\x7d\xf8\x73\xf4\xc7\x38\x3f\xc3\x9a\x58\x35\xc1\x95\xd7\xff\xe1\xca\xd3\x3f\x32\x25\x7d\x7d\x68\x72\xd4\x44\x33\x9a\x4d\xbd\x9e\xa1\x4a\x5f\x39\x41\x97\xe5\x7d\xff\x5a\xb2\x6c\xb9\x7c\xef\xca\xfc\x0a\x00\x00\xff\xff\xea\xda\x59\x3f\xf4\x0b\x00\x00")

func templatesNetwork_security_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesNetwork_security_groupTf,
		"templates/network_security_group.tf",
	)
}

func templatesNetwork_security_groupTf() (*asset, error) {
	bytes, err := templatesNetwork_security_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/network_security_group.tf", size: 3060, mode: os.FileMode(480), modTime: time.Unix(1519854618, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesOutputTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x94\xd1\x8e\xda\x30\x10\x45\xdf\xf3\x15\x56\xd4\x87\x45\xa2\x6c\x8b\x94\xaa\x42\xda\x6f\xb1\x4c\x32\x1b\xdc\x75\x6c\x6b\x66\x1c\x76\xbb\xe2\xdf\xab\xc4\x04\xc5\x01\x17\xda\x3c\x7a\xee\x3d\x73\x3d\x63\x70\x81\x7d\x60\x51\xf6\x16\x58\x5a\xd5\x41\x29\x3e\x0b\x21\x7a\x65\x02\x88\x17\x51\x7e\xf9\x54\xbf\x03\x02\x76\xb2\xd7\xc8\x41\x19\x69\x81\x8f\x0e\xdf\x36\x7b\x47\x87\xcd\xe0\x38\x95\xc5\xa9\x28\x26\x10\x85\xfd\x5d\x54\xd4\xe4\x08\x08\xe4\x02\xd6\x20\x5b\x74\xc1\xff\x9d\x94\x6a\xb3\x99\xd8\xa1\x6a\x41\xaa\xba\x76\xc1\xde\x0b\x97\x8a\x73\xcc\x06\x5e\x55\x30\x2c\x09\xea\x80\x9a\x3f\x62\x82\x2c\xf5\x3c\xb5\x85\x3c\x07\x87\x77\x06\xb4\xca\x48\x9d\x27\xfa\xb0\x37\xba\x96\xfa\x0c\xd1\x5e\xaa\xa6\x41\x20\x5a\xe4\xd4\x08\x35\x3b\x9c\xaa\x0b\xde\x81\xd9\xd3\xee\xf9\xf9\x11\xee\x6e\x5b\x55\x55\x95\xd0\x3d\xea\x5e\x31\xc8\x37\xf8\x98\x83\x87\x6f\x0c\xcb\x86\xe4\x4c\x33\x22\x65\xdf\xd1\x66\x76\x28\x3d\x74\xa7\xb2\x10\x82\xc0\x92\x66\xdd\x0f\xc1\x18\x03\x24\x8d\x62\xaa\x7f\xef\x73\xf1\x49\xe7\xc1\x12\x1d\xae\x5a\xbd\x2a\x43\x49\xaf\x5f\xa1\xf3\x7b\xf7\x2e\x03\x9a\xff\x98\xfe\x6e\xbb\x4d\x46\x34\x6d\xbe\xd6\x0d\x5e\xe1\x7a\x85\x9b\xb9\x20\xee\xce\xb8\x5a\x19\x1a\xb5\x97\xf5\x0d\x8f\x44\x9c\xbf\x17\x51\x0e\x5d\xbf\x46\x3f\xd8\x5e\xea\x66\xbc\x96\xb6\xe7\x77\x33\xb0\x66\xea\x28\x4c\xaa\xa9\xbe\x3d\x0a\x91\xe8\x07\xc9\xc1\x11\x3f\x8d\x59\x52\xeb\x5a\x7c\x5f\x8d\xf6\x69\x50\x97\xaa\xf6\x0f\xd9\xab\x68\xbf\xdc\x6d\xee\x7f\xc0\xfe\x63\x95\x79\xe2\x37\x7f\xd7\x11\x91\x68\x52\x7b\x42\xcf\xd8\x97\xa3\xbb\xfe\xc3\xbb\x69\x1e\x0e\x63\xfd\x69\xb9\xe9\xb5\xf8\xb9\x16\xdf\x56\x99\x28\xed\xf1\x5e\x90\xf6\x98\x5a\xa7\x5d\xcc\x87\x99\x61\xdc\x58\x5b\x66\xa0\x0f\xc0\x6e\x2d\x71\xa4\xfd\x09\x00\x00\xff\xff\x40\xfd\xdb\x23\x5a\x06\x00\x00")

func templatesOutputTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesOutputTf,
		"templates/output.tf",
	)
}

func templatesOutputTf() (*asset, error) {
	bytes, err := templatesOutputTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/output.tf", size: 1626, mode: os.FileMode(480), modTime: time.Unix(1519854618, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResource_groupTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x90\x41\x8a\xc4\x20\x10\x45\xf7\x9e\xe2\x23\xb3\x9d\xdc\x60\xce\x22\x15\x53\x64\x84\x44\x43\xa9\x59\x4c\xf0\xee\x83\x42\xa7\x3b\x24\xdd\x0d\x5d\x4b\xf9\xf5\xeb\xf9\x84\x63\xc8\x62\x19\x9a\xfe\xb2\xb0\xcc\xe6\xf6\x62\x46\x09\x79\xd1\xd0\x7d\x88\xbf\x1a\x9b\x02\x3c\xcd\x8c\x3a\x3f\xd0\x5f\xdb\x4a\xd2\xb1\x5f\x8d\x1b\xca\x77\xcb\x28\x60\x0a\x96\x92\x0b\xfe\x9e\x10\x1e\x5d\xf0\x45\x2b\x05\x24\x1a\x63\x2b\x02\xd8\xaf\x4e\x82\x9f\xd9\xa7\x53\x5b\x2d\x2a\xaa\x28\x75\x86\x5b\x72\x3f\x39\x6b\xdc\x13\xae\xab\x79\xcf\xfa\x72\x6b\xe7\x07\x8e\x66\xcc\xf1\x6a\x5b\xb8\x76\xd8\xd5\x8b\x5d\x8d\xb7\x9a\xfd\x0f\x86\x86\x41\x38\x46\x43\xd3\xa3\xb7\x98\x28\x39\xfb\x89\xb0\xff\x00\x00\x00\xff\xff\x58\xec\xbb\xb4\xcd\x01\x00\x00")

func templatesResource_groupTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesResource_groupTf,
		"templates/resource_group.tf",
	)
}

func templatesResource_groupTf() (*asset, error) {
	bytes, err := templatesResource_groupTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/resource_group.tf", size: 461, mode: os.FileMode(480), modTime: time.Unix(1519854618, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesStorageTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\xcd\x6a\xc3\x30\x0c\xc7\xef\x7e\x0a\x61\x76\xee\x1b\xec\xbc\xfb\xfa\x00\x46\x71\x44\x66\x70\xa4\x20\x2b\x81\xad\xe4\xdd\x47\xbc\x36\x6d\x0d\xa5\x3d\x2e\xd7\xfc\xf4\xff\xb2\x52\x91\x59\x23\x81\xc7\x9f\x59\x49\xc7\x50\x4c\x14\x07\x0a\x18\xa3\xcc\x6c\x1e\x7c\x27\xe5\xcb\xc3\xc9\x01\x30\x8e\x04\xcd\xf7\x0e\xfe\xed\xb4\xa0\x1e\x4a\x1a\xa7\x4c\x81\x78\x09\xa9\x5f\xbd\x03\xb8\x88\x87\x41\x65\x9e\x42\xbd\xae\xf8\xc5\xeb\x1e\x38\x6c\x46\x87\x8d\x5a\xbd\x73\x00\x59\x22\x5a\x12\x6e\x1d\xaf\x96\x4a\x43\x12\xae\x5e\xe7\xb8\xc1\x12\x69\x0b\x1f\x0d\xb9\x47\xed\x6f\x39\xa5\x29\xa7\x3f\xfd\x60\xdf\x53\x0d\xf6\xf1\x79\xac\xc6\x86\x43\xa9\x7d\x01\x88\x97\xa4\xc2\x23\xb1\x5d\x6d\x6f\x2a\xae\x6e\x75\xee\xf1\x88\x51\xd8\x30\x31\xe9\xd3\x19\x6b\xd0\x8a\x3c\x18\x0e\x5e\x9e\x0e\xa0\x79\xc3\xb3\xc0\xdd\x7d\x83\x34\x02\x7b\xee\xed\x3f\x95\xb2\x4f\x34\x69\x5a\xd0\xc8\xbf\x5e\xbb\x18\x8d\x91\x72\x7e\x52\x7d\xc7\xfe\x75\xfd\x2e\x4b\xb7\x75\xff\x0d\x00\x00\xff\xff\x23\xe4\x2a\x58\x37\x03\x00\x00")

func templatesStorageTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesStorageTf,
		"templates/storage.tf",
	)
}

func templatesStorageTf() (*asset, error) {
	bytes, err := templatesStorageTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/storage.tf", size: 823, mode: os.FileMode(480), modTime: time.Unix(1519854618, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesTlsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\x41\x0a\x02\x31\x0c\x05\xd0\x7d\x4e\xf1\xc9\x09\x5c\x88\xe0\xa2\x0b\xaf\xa0\x07\x08\xad\x04\x5b\x6c\xa9\x24\xb1\x30\x0c\x73\xf7\x79\xa6\x3e\xff\xf6\x56\x70\x74\x97\x9f\xb5\x95\x43\xe5\xab\x1b\x83\xcb\xf4\x2a\x6b\x38\x63\x27\x20\xf7\xcf\xb4\x16\x75\x20\x81\x9f\xaf\x07\x13\x60\x9e\xa5\xb4\x70\x20\xe1\x7a\xb9\xdf\xe8\xa0\x33\x00\x00\xff\xff\x52\x4d\xac\xad\x51\x00\x00\x00")

func templatesTlsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesTlsTf,
		"templates/tls.tf",
	)
}

func templatesTlsTf() (*asset, error) {
	bytes, err := templatesTlsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/tls.tf", size: 81, mode: os.FileMode(480), modTime: time.Unix(1519854618, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x90\xc1\x6a\xc4\x20\x10\x40\xef\x7e\xc5\x20\x3d\xa7\xcd\xa5\xb7\x7e\x4b\x30\x3a\x2d\x43\xcd\x18\x46\x63\xa1\x21\xff\x5e\x12\xc1\xee\x8a\x2c\x9b\xdc\xe6\xbd\x07\xce\x64\x23\x64\x66\x8f\xa0\x91\xf3\x44\x4e\xc3\x7e\x28\xf5\x3f\x15\xfc\xa2\xc0\xed\x34\xd2\xb2\x7a\x9c\xfa\x49\xdc\xe6\x68\x85\xd6\x44\x81\x3b\x38\x21\x1b\x4e\x1d\x60\x3d\xe1\x23\x10\xd1\x0a\xa6\x16\x32\xa6\x9f\x20\xdf\x93\x25\x27\x1a\x76\x05\xe0\xf0\xd3\x6c\x3e\xc1\x07\xe8\xf1\x6d\xb8\xfe\xd7\xf1\x5d\xab\xbb\x8c\x38\xa1\xb0\xf1\xcf\x75\xab\x84\x4c\x0e\x05\xb4\xf9\xdd\x04\x65\x29\x45\xb3\xe9\x59\xbe\xec\xd9\xc8\xd0\x80\x43\x2b\x80\xba\x37\x94\xaf\xca\x15\x5c\x5a\xbd\x42\xab\x55\x70\xab\x95\x9b\x74\xb4\x02\x8e\xf3\xf5\x7f\x01\x00\x00\xff\xff\x59\x69\xa5\xda\xe3\x01\x00\x00")

func templatesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVarsTf,
		"templates/vars.tf",
	)
}

func templatesVarsTf() (*asset, error) {
	bytes, err := templatesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars.tf", size: 483, mode: os.FileMode(480), modTime: time.Unix(1519854618, 0)}
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
	"templates/cf_dns.tf": templatesCf_dnsTf,
	"templates/cf_lb.tf": templatesCf_lbTf,
	"templates/concourse_lb.tf": templatesConcourse_lbTf,
	"templates/network.tf": templatesNetworkTf,
	"templates/network_security_group.tf": templatesNetwork_security_groupTf,
	"templates/output.tf": templatesOutputTf,
	"templates/resource_group.tf": templatesResource_groupTf,
	"templates/storage.tf": templatesStorageTf,
	"templates/tls.tf": templatesTlsTf,
	"templates/vars.tf": templatesVarsTf,
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
	"templates": &bintree{nil, map[string]*bintree{
		"cf_dns.tf": &bintree{templatesCf_dnsTf, map[string]*bintree{}},
		"cf_lb.tf": &bintree{templatesCf_lbTf, map[string]*bintree{}},
		"concourse_lb.tf": &bintree{templatesConcourse_lbTf, map[string]*bintree{}},
		"network.tf": &bintree{templatesNetworkTf, map[string]*bintree{}},
		"network_security_group.tf": &bintree{templatesNetwork_security_groupTf, map[string]*bintree{}},
		"output.tf": &bintree{templatesOutputTf, map[string]*bintree{}},
		"resource_group.tf": &bintree{templatesResource_groupTf, map[string]*bintree{}},
		"storage.tf": &bintree{templatesStorageTf, map[string]*bintree{}},
		"tls.tf": &bintree{templatesTlsTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
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

