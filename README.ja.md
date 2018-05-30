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

# run(引数の追加)
nem.sh run --hash 666ade37fe3f09882b72645d262a015d68bc1230dc1d0f742bb098bc6567d682 20

# 実行(登録したエイリアスから)
nem.sh run --alias hello

# 実行(ソースコードの表示)
nem.sh run --alias hello --view
Y='\033[0;33m'
B='\033[1;34m'
G='\033[1;36m'
N='\033[0m'
printf "Hello, nem.sh!\n"
printf "       ${Y}sssssssssssss${N}\n"
printf "  ${Y}ssssssssssssssss${N}  ${B}sssssssssss${N}\n"
printf "${Y}ssssssssssssssss${N}   ${B}sssssssssssss${N}\n"
printf "${Y}sssssssssssssss${N}  ${B}sssssssssssssss${N}\n"
printf " ${Y}ssssssssssssss${N}  ${B}ssssssssssssss${N}\n"
printf "  ${Y}sssssssssssss${N}   ${B}sssssssssssss${N}\n"
printf "    ${Y}'''''''''${N}         ${B}ssssssss${N}\n"
printf "   ${G}hssssssssssssssss${N}    ${B}sssss${N}\n"
printf "    ${G}sssssssssssssssssss${N}  ${B}sss${N}\n"
printf "      ${G}sssssssssssssssss${N}  ${B}ss${N}\n"
printf "        ${G}sssssssssssssss${N}  \n"
printf "          ${G}ssssssssssss${N}\n"
printf "             ${G}sssssss${N}\n"
printf "               ${G}sss${N}\n"
```


## エイリアスの設定

```
# 設定(txハッシュと名前のヒモ付)
nem.sh set-alias --hash b37685ca16474b6897550f51f008c11b1e24e93e51b5543d066d9266d4e35008 --name hello
```

## コンフィグファイルについて

Configファイルのパスは「$HOME/.nem.sh/config.json」です。

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
