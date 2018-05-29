# nem.sh
nemのトランザクションを使用したシェルスクリプトホストCLIツール。  

<img src="https://i.imgur.com/i5MzBpt.png" data-canonical-src="https://i.imgur.com/i5MzBpt.png" width="200"/>

５月は重厚な作業が多かったので...。ただトランザクション読むだけのやつをやりたくなった...。

## インストール

```
go get github.com/scrpgil/nem.sh
```

## 実行

```
# 実行(Txハッシュから)
nem.sh run --hash b37685ca16474b6897550f51f008c11b1e24e93e51b5543d066d9266d4e35008

# 実行(登録したエイリアスから)
nem.sh run --alias hello

# 実行(ソースコードの表示)
nem.sh run --alias hello --view
```


## エイリアスの設定

```
# 設定(txハッシュと名前のヒモ付)
nem.sh set-alias --hash b37685ca16474b6897550f51f008c11b1e24e93e51b5543d066d9266d4e35008 --name hello
```

## コンフィグファイルについて

Configファイルのパスは「$HOME/.nem.sh.json」です。

```
{
  "address": "NCNUJTJ7HAYF6PV3ZPJIAYVJFTR4VT4FG4C4FRA5",
  "alias": {
    "hello": "b37685ca16474b6897550f51f008c11b1e24e93e51b5543d066d9266d4e35008"
  },
  "node": "san.nem.ninja",
  "port": "7890",
  "protocol": "http"
}
```

## 新しいコマンドの追加

新しいコマンドを追加する場合は、そのコマンドをトランザクションのメッセージに書き込みます。
