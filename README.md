# FileMover
移动符合条件的文件到新目录，并保持原文件的目录结构。

##### 1. 使用说明
- 修改`config.toml`中`path`下的`src_dir`和`dest_dir`
- 运行程序：`go run main.go`
- 程序会将`src_dir`下的目录和文件移动到`dest_dir`目录下，并保持原来基于`src_dir`的路径相对结构
    - 如：
    ```
    src_dir
     |- a.txt
     |- b_dir
        |- b.txt
    ```
    移动到dest_dir后
    ```
    dest_dir
     |- a.txt
     |- b_dir
        |- b.txt
    ```
---