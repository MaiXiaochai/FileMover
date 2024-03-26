# FileMover
移动符合条件的文件到新目录，并保持部分原文件的目录结构。

##### 1. 使用说明
- 修改`config.toml`中path下的src_dir和dest_dir
- 运行程序：`go run main.go`
- 程序会将src_dir下的目录和文件移动到dest_dir目录下，并保持基于src_dir的路径相对结构
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