package mover

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func MoveFiles(srcDir, destDir string, conditionFunc func(os.FileInfo) bool) error {
	// 遍历路径总数，成功移动总数
	var pathCount, movedCount int
	err := filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		pathCount++
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
			if err = moveFile(path, newPath); err == nil {
				movedCount++
				if remainder := movedCount % 100; remainder == 0 && movedCount != 0 {
					fmt.Printf("[ %d/%d: %.1f%% ][ %s ]", movedCount, pathCount, float64(movedCount)/float64(pathCount)*100, filepath.Base(path))
				}
			}
			return err
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
	err = os.Remove(src)
	if err != nil {
		return err
	}
	return nil
}
