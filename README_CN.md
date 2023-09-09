<div align="center" width="100%">
    <img src="./assets/icon.svg" width="50%" alt="">
</div>

# Munager

[![Go Report Card](https://goreportcard.com/badge/github.com/TurboHsu/munager)](https://goreportcard.com/report/github.com/TurboHsu/munager)

## 0x01 这啥

Munager 是一个有以下功能的音乐库管理工具：

- 自动化的为您的本地乐库抓取歌词
- 跨设备同步乐库，支持自定义转码规则

## 0x02 咋装

如果你喜欢自己编译&运行一些东西，直接使用 `go build .`就可以了。

如果你想要直接用这个, 可以在 `Releases` 中找到发行版，或者你可以在 `Github Actions Artifacts` 中找到CI构建的版本。

## 0x03 咋用

_Munager 的帮助信息足够解释自己。 尝试添加 `--help` 来查看详细参数以及使用方法。_

### 歌词

敲 `munager lyric` 来获取帮助。 它有很多子指令，分别是：

- `query` 可以以关键字搜索指定歌词，并且丢到标准输出中。
- `fetch` 可以为指定目录下的所有歌曲抓取歌词。

### 同步

敲 `munager sync` 来获取帮助。 它有很多子指令，分别是：

- `client` 将会启动客户端。在局域网内，它一般能够自动发现服务器，并且同步本地乐库（增量更新）。您也可以使用 `--transcode` 参数来指定转码规则。（这将会用到FFmpeg。您可以指定FFmpeg可执行文件的路径。）
- `server` 将会启动服务器。它会向局域网内广播自己的存在，并且提供本地乐库的服务，就如你想要的那样。

## 0x04 例子

### 我要同步一些歌

想象你有一个乐库，它里面并没有歌词，就像这样：

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

如果你想给他们抓取歌词，只需要运行：

```bash
$ munager lyric fetch --path="./THE BOOK/"
2023-07-14 09:43:53  [I]  Found 9 songs without lyrics
 100% |███████████████████████████████████████████████████████████████████████████████████████████| (9/9, 7 it/s)        
2023-07-14 09:43:55  [I]  Done!

```

然后，你应该可以在你的乐库中看到一些 `.lrc` 文件，就像这样：

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

### 乐库同步

想象你有两台设备，他们的乐库并不同步。你想要同步他们，但是你又不想要把所有的音乐都复制一遍。这时候，你可以使用 Munager 来帮助你。这里把同步源作为服务器，被同步源作为客户端，他们的乐库长这样：

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

你可以看见有几首没有同步。

在服务器那边，运行 `munager sync server`，在客户端那边，运行 `munager sync client`。他们会自动做他们应该做的事情，然后你就可以看见他们的乐库同步了。

如果你想要在客户端处转码，你只需要添加`--transcode`参数，它包含一些预制的压制参数，运行之后就会变成这样：

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

## 0xFF 感谢

- 网易云API是加密的，我从 [mxget](https://github.com/winterssy/mxget) 抄了一些代码。谢谢你兄弟。
