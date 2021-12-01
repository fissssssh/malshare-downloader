# malshare-downloader

![](https://img.shields.io/github/workflow/status/fissssssh/malshare-downloader/Build)
![](https://img.shields.io/github/v/release/fissssssh/malshare-downloader?display_name=tag&include_prereleases)

malshare-downloader

[简体中文](/README.md)|[English](/docs/README_en-us.md)

## Usage

- Download hash files

  ```shell
  $ ./spider -start=<start_unix_milli_timestamp> [-end=<end_unix_milli_timestamp>] [-o=<output_dir>]
  ```

  | Parameters | Required | Remark                                                        |
  | ---------- | -------- | ------------------------------------------------------------- |
  | `start`    | ✅        | the start time of hash files                                  |
  | `end`      |          | the end time of hash files<br>default is `now`                |
  | `o`        |          | the output directory of hash files<br>default is `hash_files` |

- Download mal files

  ```shell
  $ ./downloader -keys_file=<keys_file> [-source=<hash_files_dir>] [-type=<mal_file_tyle>] [-yara=<mal_file_yara>] [-o=<output_dir>]
  ```

  | Parameters  | Required | Remark                                                                |
  | ----------- | -------- | --------------------------------------------------------------------- |
  | `keys_file` | ✅        | the api keys file of Malshare                                         |
  | `source`    |          | the directory of hash files<br>default is `hash_files`                |
  | `type`      |          | the type of mal files that you want                                   |
  | `yara`      |          | the yara of mal files that you want                                   |
  | `o`         |          | the output directory of mal files<br>default is `mal_files/hash_file` |
