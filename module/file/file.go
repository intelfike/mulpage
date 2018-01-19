package file

import "os"

func Exists(filename string) bool {
	f, err := os.Stat(filename)
	return err == nil && f != nil
}
func ExistsFile(filename string) bool {
	f, err := os.Stat(filename)
	return err == nil && f != nil && !f.IsDir()
}
func ExistsDir(dirname string) bool {
	f, err := os.Stat(dirname)
	return err == nil && f != nil && f.IsDir()
}
