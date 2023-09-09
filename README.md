<div align="center" width="100%">
    <img src="./assets/icon.svg" width="50%" alt="">
</div>

# Munager

[![Go Report Card](https://goreportcard.com/badge/github.com/TurboHsu/munager)](https://goreportcard.com/report/github.com/TurboHsu/munager)

[中文文档](./README_CN.md)

## 0x00 Intro

Munager is a music library management helper that have several functions:

- Automatically search for lyric files on online services like Netease Music, etc.
- Syncs your music library between devices with customizable transcoding rules.

## 0x01 Install

If you wanna compile yourself, just clone this project and run `go build .`.

If you just want to run the binary, find it in `Releases` or `Github Actions Artifacts` for CI Build.

## 0x02 Usage

_Munager is quite self-explained. Use `--help` for explainations._

### Lyric

Type `munager lyric` for help. It have several subcommand which does:

- `query` Querys lyric for given keyword.
- `fetch` Fetches lyrics for every songs in provided path.

### Sync

Type `munager sync` for help. It have several subcommands which does:

- `client` Starts a client. It can automatically discover server side and syncs local music library with server side. You can also transcode your received songs to another format (FFmpeg) using `--transcode` flag.
- `server` Starts a server. It broadcasts its existance to local network, and serves local music library as you expected.

## 0x03 Example

### Fetching lyrics

Now you have an folder which contains some songs:

```bash
THE BOOK
├── 01. Epilogue.flac
├── 02. アンコール.flac
├── 03. ハルジオン.flac
├── 04. あの夢をなぞって.flac
├── 05. たぶん.flac
├── 06. 群青.flac
├── 07. ハルカ.flac
├── 08. 夜に駆ける.flac
├── 09. Prologue.flac
└── cover.png

1 directory, 10 files
```

To fetch lyric for it, just run

```bash
$ munager lyric fetch --path="./THE BOOK/"
2023-07-14 09:43:53  [I]  Found 9 songs without lyrics
 100% |███████████████████████████████████████████████████████████████████████████████████████████| (9/9, 7 it/s)        
2023-07-14 09:43:55  [I]  Done!

```

Then you should see some `.lrc` files in your sweet folder.

```bash
THE BOOK
├── 01. Epilogue.flac
├── 01. Epilogue.lrc
├── 02. アンコール.flac
├── 02. アンコール.lrc
├── 03. ハルジオン.flac
├── 03. ハルジオン.lrc
├── 04. あの夢をなぞって.flac
├── 04. あの夢をなぞって.lrc
├── 05. たぶん.flac
├── 05. たぶん.lrc
├── 06. 群青.flac
├── 06. 群青.lrc
├── 07. ハルカ.flac
├── 07. ハルカ.lrc
├── 08. 夜に駆ける.flac
├── 08. 夜に駆ける.lrc
├── 09. Prologue.flac
├── 09. Prologue.lrc
└── cover.png

1 directory, 19 files
```

### Syncing music library

Imagine you have two devices which contains music library that is not synced.

```bash
server
└── THE BOOK
    ├── 01. Epilogue.flac
    ├── 01. Epilogue.lrc
    ├── 02. アンコール.flac
    ├── 02. アンコール.lrc
    ├── 03. ハルジオン.flac
    ├── 03. ハルジオン.lrc
    ├── 04. あの夢をなぞって.flac
    ├── 04. あの夢をなぞって.lrc
    ├── 05. たぶん.flac
    ├── 05. たぶん.lrc
    ├── 06. 群青.flac
    ├── 06. 群青.lrc
    ├── 07. ハルカ.flac
    ├── 07. ハルカ.lrc
>   ├── 08. 夜に駆ける.flac
>   ├── 08. 夜に駆ける.lrc
>   ├── 09. Prologue.flac
>   ├── 09. Prologue.lrc
    └── cover.png   
client
└── THE BOOK
    ├── 01. Epilogue.flac
    ├── 01. Epilogue.lrc
    ├── 02. アンコール.flac
    ├── 02. アンコール.lrc
    ├── 03. ハルジオン.flac
    ├── 03. ハルジオン.lrc
    ├── 04. あの夢をなぞって.flac
    ├── 04. あの夢をなぞって.lrc
    ├── 05. たぶん.flac
    ├── 05. たぶん.lrc
    ├── 06. 群青.flac
    ├── 06. 群青.lrc
    ├── 07. ハルカ.flac
    ├── 07. ハルカ.lrc
    └── cover.png

4 directories, 34 files
```

On server side, just run `munager sync server`, and on client side, just run `munager sync client`. After they do their stuff, the music library should be synced.

Adding a `--transcode` flag can also achieve this:

```bash
server
└── THE BOOK
    ├── 01. Epilogue.flac
    ├── 02. アンコール.flac
    ├── 03. ハルジオン.flac
    ├── 04. あの夢をなぞって.flac
    ├── 05. たぶん.flac
    ├── 06. 群青.flac
    ├── 07. ハルカ.flac
    ├── 08. 夜に駆ける.flac
    ├── 09. Prologue.flac
    └── cover.png
client
└── THE BOOK
    ├── 01. Epilogue.opus
    ├── 02. アンコール.opus
    ├── 03. ハルジオン.opus
    ├── 04. あの夢をなぞって.opus
    ├── 05. たぶん.opus
    ├── 06. 群青.opus
    ├── 07. ハルカ.opus
    ├── 08. 夜に駆ける.opus
    ├── 09. Prologue.opus
    └── cover.png

4 directories, 20 files

```

## 0xFF Credits

- Netease API is somehow encrypted, I took some code from [mxget](https://github.com/winterssy/mxget) for reference.
