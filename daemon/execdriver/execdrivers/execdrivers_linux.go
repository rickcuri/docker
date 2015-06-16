// +build linux

package execdrivers

import (
	"fmt"
	"path"

	"github.com/docker/docker/daemon/execdriver"
	"github.com/docker/docker/daemon/execdriver/lxc"
	"github.com/docker/docker/daemon/execdriver/native"
	"github.com/docker/docker/pkg/sysinfo"
)

<<<<<<< HEAD:daemon/execdriver/execdrivers/execdrivers.go
func NewDriver(name, root, libPath, initPath string, sysInfo *sysinfo.SysInfo) (execdriver.Driver, error) {
=======
func NewDriver(name string, options []string, root, libPath, initPath string, sysInfo *sysinfo.SysInfo) (execdriver.Driver, error) {
>>>>>>> 9dffd032f43dca88e8a66685aceb21b0e00cc232:daemon/execdriver/execdrivers/execdrivers_linux.go
	switch name {
	case "lxc":
		// we want to give the lxc driver the full docker root because it needs
		// to access and write config and template files in /var/lib/docker/containers/*
		// to be backwards compatible
		return lxc.NewDriver(root, libPath, initPath, sysInfo.AppArmor)
	case "native":
		return native.NewDriver(path.Join(root, "execdriver", "native"), initPath, options)
	}
	return nil, fmt.Errorf("unknown exec driver %s", name)
}
