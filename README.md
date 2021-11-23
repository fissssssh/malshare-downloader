# malshare-downloader

malshare-downloader

[简体中文](/README.md)|[English](/docs/README_en-us.md)

## 用法

- 下载 hash 文件

  ```shell
  $ ./spider -start=<start_unix_milli_timestamp> [-end=<end_unix_milli_timestamp>] [-o=<output_dir>]
  ```

  | 参数名  | 是否必要 | 备注                                         |
  | ------- | -------- | -------------------------------------------- |
  | `start` | ✅       | hash 文件的开始时间                          |
  | `end`   |          | hash 文件的结束时间<br>默认是 `现在`         |
  | `o`     |          | hash 文件的保存文件夹<br>默认是 `hash_files` |

- 下载 mal 文件

  ```shell
  $ ./downloader -api=<api_key> [-source=<hash_files_dir>] [-t=<mal_file_tyle>] [-o=<output_dir>]
  ```

  | 参数名   | 是否必要 | 备注                                                 |
  | -------- | -------- | ---------------------------------------------------- |
  | `api`    | ✅       | 调用 Malshare 的 api key                             |
  | `source` |          | hash 文件的文件夹<br>默认是 `hash_files`             |
  | `t`      |          | 你想要下载的 mal 文件类型                            |
  | `o`      |          | mal 文件的保存文件夹<br>默认是 `mal_files/hash_file` |
