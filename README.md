
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
