package allocdir

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

const (
	// secretDirTmpfsSize is the size of the tmpfs per task in MBs
	secretDirTmpfsSize = 1
)

// linkDir bind mounts src to dst as Linux doesn't support hardlinking
// directories.
func linkDir(src, dst string) error {
	if err := os.MkdirAll(dst, 0777); err != nil {
		return err
	}

	return syscall.Mount(src, dst, "", syscall.MS_BIND, "")
}

// unlinkDir unmounts a bind mounted directory as Linux doesn't support
// hardlinking directories.
func unlinkDir(dir string) error {
	return syscall.Unmount(dir, 0)
}

// createSecretDir creates the secrets dir folder at the given path using a
// tmpfs
func createSecretDir(dir string) error {
	// Only mount the tmpfs if we are root
	if unix.Geteuid() == 0 {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}

		var flags uintptr
		flags = syscall.MS_NOEXEC
		options := fmt.Sprintf("size=%dm", secretDirTmpfsSize)
		err := syscall.Mount("tmpfs", dir, "tmpfs", flags, options)
		return os.NewSyscallError("mount", err)
	}

	return os.MkdirAll(dir, 0777)
}

// createSecretDir removes the secrets dir folder
func removeSecretDir(dir string) error {
	if unix.Geteuid() == 0 {
		if err := syscall.Unmount(dir, 0); err != nil {
			return os.NewSyscallError("unmount", err)
		}
	}

	return os.RemoveAll(dir)
}
