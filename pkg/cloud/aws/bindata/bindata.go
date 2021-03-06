// Code generated for package bindata by go-bindata DO NOT EDIT. (@generated)
// sources:
// pkg/cloud/aws/manifests/machineset-aggeragate-clusterrole.yaml
// pkg/cloud/aws/manifests/machineset.yaml
package bindata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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

// Mode return file modify time
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

var _pkgCloudAwsManifestsMachinesetAggeragateClusterroleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: open-cluster-management:submariner-addon-machinesets-aggregate-clusterrole
  labels:
    rbac.authorization.k8s.io/aggregate-to-admin: "true"
rules:
  - apiGroups: ["machine.openshift.io"]
    resources: ["machinesets"]
    verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]
`)

func pkgCloudAwsManifestsMachinesetAggeragateClusterroleYamlBytes() ([]byte, error) {
	return _pkgCloudAwsManifestsMachinesetAggeragateClusterroleYaml, nil
}

func pkgCloudAwsManifestsMachinesetAggeragateClusterroleYaml() (*asset, error) {
	bytes, err := pkgCloudAwsManifestsMachinesetAggeragateClusterroleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/cloud/aws/manifests/machineset-aggeragate-clusterrole.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgCloudAwsManifestsMachinesetYaml = []byte(`apiVersion: machine.openshift.io/v1beta1
kind: MachineSet
metadata:
  labels:
    machine.openshift.io/cluster-api-cluster: {{ .InfraId }}
  name: {{ .InfraId }}-submariner-gw-{{ .AZ }}
  namespace: openshift-machine-api
spec:
  replicas: 1
  selector:
    matchLabels:
      machine.openshift.io/cluster-api-cluster: {{ .InfraId }}
      machine.openshift.io/cluster-api-machineset: {{ .InfraId }}-submariner-gw-{{ .AZ }}
  template:
    metadata:
      labels:
        machine.openshift.io/cluster-api-cluster: {{ .InfraId }}
        machine.openshift.io/cluster-api-machine-role: worker
        machine.openshift.io/cluster-api-machine-type: worker
        machine.openshift.io/cluster-api-machineset: {{ .InfraId }}-submariner-gw-{{ .AZ }}
    spec:
      metadata:
        labels:
          submariner.io/gateway: "true"
      taints:
        - effect: NoSchedule
          key: node-role.submariner.io/gateway
      providerSpec:
        value:
          ami:
            id: {{ .AMIId }}
          apiVersion: awsproviderconfig.openshift.io/v1beta1
          credentialsSecret:
            name: aws-cloud-credentials
          deviceIndex: 0
          iamInstanceProfile:
            id: {{ .InfraId }}-worker-profile
          instanceType: {{ .InstanceType }}
          kind: AWSMachineProviderConfig
          placement:
            availabilityZone: {{ .AZ }}
            region: {{ .Region }}
          securityGroups:
            - filters:
                - name: tag:Name
                  values:
                    - {{ .InfraId }}-worker-sg
                    - {{ .SecurityGroupName }}
          subnet:
            filters:
              - name: tag:Name
                values:
                  - {{ .SubnetName }}
          tags:
            - name: kubernetes.io/cluster/{{ .InfraId }}
              value: owned
            - name: submariner.io
              value: gateway
          userDataSecret:
            name: worker-user-data
          publicIp: true
`)

func pkgCloudAwsManifestsMachinesetYamlBytes() ([]byte, error) {
	return _pkgCloudAwsManifestsMachinesetYaml, nil
}

func pkgCloudAwsManifestsMachinesetYaml() (*asset, error) {
	bytes, err := pkgCloudAwsManifestsMachinesetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/cloud/aws/manifests/machineset.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"pkg/cloud/aws/manifests/machineset-aggeragate-clusterrole.yaml": pkgCloudAwsManifestsMachinesetAggeragateClusterroleYaml,
	"pkg/cloud/aws/manifests/machineset.yaml":                        pkgCloudAwsManifestsMachinesetYaml,
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
	"pkg": {nil, map[string]*bintree{
		"cloud": {nil, map[string]*bintree{
			"aws": {nil, map[string]*bintree{
				"manifests": {nil, map[string]*bintree{
					"machineset-aggeragate-clusterrole.yaml": {pkgCloudAwsManifestsMachinesetAggeragateClusterroleYaml, map[string]*bintree{}},
					"machineset.yaml":                        {pkgCloudAwsManifestsMachinesetYaml, map[string]*bintree{}},
				}},
			}},
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
