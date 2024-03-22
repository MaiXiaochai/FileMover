package mover

import (
	"io"
	"os"
	"path/filepath"
)

func MoveFiles(srcDir, destDir string, conditionFunc func(os.FileInfo) bool) error {
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果当前项满足移动条件
		if conditionFunc(info) {
			// filepath.Rel，根据两个文件路径计算出从第一个路径到第二个路径的相对路径
			relativePath, err := filepath.Rel(srcDir, path)
			if err != nil {
				return err
			}

			newPath := filepath.Join(destDir, relativePath)
			// 创建新路径所需的所有目录
			// 用来获取给定文件路径的目录部分，等价于Python中的 os.path.dirname
			if err := os.MkdirAll(filepath.Dir(newPath), 0755); err != nil {
				return err
			}
			// 移动文件
			return moveFile(path, newPath)
		}
		return nil
	})

	return err

}

// 文件移动
func moveFile(src, dest string) error {
	input, err := os.Open(src)
	if err != nil {
		return err
	}

	defer input.Close()

	output, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer output.Close()

	_, err = io.Copy(output, input)
	if err != nil {
		return err
	}
	return os.Remove(src)
}
