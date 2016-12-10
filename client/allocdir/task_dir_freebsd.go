package allocdir

// MountSpecialDirs mounts the dev and proc file system on the chroot of the
// task. It's a no-op on FreeBSD right now.
func (t *TaskDir) MountSpecialDirs(taskDir string) error {
	return nil
}

// unmountSpecialDirs unmounts the dev and proc file system from the chroot
func (t *TaskDir) unmountSpecialDirs(taskDir string) error {
	return nil
}
