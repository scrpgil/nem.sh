
[日本語のREADME.mdはこちら](https://github.com/scrpgil/nem.sh/blob/master/README.ja.md)

# nem.sh
Shell Script Host CLI Tool Using nem's Transaction.

<img src="https://i.imgur.com/i5MzBpt.png" data-canonical-src="https://i.imgur.com/i5MzBpt.png" width="200"/>

## Install

```
go get github.com/scrpgil/nem.sh
```

## Run

```
# run
nem.sh run --hash b37685ca16474b6897550f51f008c11b1e24e93e51b5543d066d9266d4e35008

# run(use alias)
nem.sh run --alias hello

# run(view shell scripts)
nem.sh run --alias hello --view
```


## SetAlias

```
# set-alias
nem.sh set-alias --hash b37685ca16474b6897550f51f008c11b1e24e93e51b5543d066d9266d4e35008 --name hello
```

## Edit Config

The path of the Config file is ($HOME/.nem.sh.json).

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

## Add New Command

If you want to add a new command, write it in the message of the transaction.
