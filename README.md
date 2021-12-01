# malshare-downloader

![](https://img.shields.io/github/workflow/status/fissssssh/malshare-downloader/Build)
![](https://img.shields.io/github/v/release/fissssssh/malshare-downloader?display_name=tag&include_prereleases)

malshare-downloader

[简体中文](/README.md)|[English](/docs/README_en-us.md)

## 用法

- 下载 hash 文件

  ```shell
  $ ./spider -start=<start_unix_milli_timestamp> [-end=<end_unix_milli_timestamp>] [-o=<output_dir>]
  ```

  | 参数名  | 是否必要 | 备注                                         |
  | ------- | -------- | -------------------------------------------- |
  | `start` | ✅        | hash 文件的开始时间                          |
  | `end`   |          | hash 文件的结束时间<br>默认是 `现在`         |
  | `o`     |          | hash 文件的保存文件夹<br>默认是 `hash_files` |

- 下载 mal 文件

  ```shell
  $ ./downloader -keys_file=<keys_file> [-source=<hash_files_dir>] [-type=<mal_file_tyle>] [-yara=<mal_file_yara>] [-o=<output_dir>]
  ```

  | 参数名      | 是否必要 | 备注                                       |
  | ----------- | -------- | ------------------------------------------ |
  | `keys_file` | ✅        | 调用 Malshare 的 api keys文件              |
  | `source`    |          | hash 文件的文件夹<br>默认是 `hash_files`   |
  | `type`      |          | 你想要下载的 mal 文件类型                  |
  | `yara`      |          | 你想要下载的 mal 文件的Yarahits            |
  | `o`         |          | mal 文件的保存文件夹<br>默认是 `mal_files` |
