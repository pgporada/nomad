package allocdir

// currently a noop on darwin
func (d *TaskDir) mountSpecialDirs() error {
	return nil
}

// currently a noop on darwin
func (d *TaskDir) unmountSpecialDirs() error {
	return nil
}
