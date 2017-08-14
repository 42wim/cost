# cost

bitcoin, ethereum, altcoin address (pubkey/privkey) generator

:warning: This is just an experiment to learn more about cryptocurrencies.
Do not use in production :)

## Building
Go 1.7+ is required. Make sure you have [Go](https://golang.org/doc/install) properly installed, including setting up your [GOPATH] (https://golang.org/doc/code.html#GOPATH)

```
cd $GOPATH
go get -tags nocgo github.com/42wim/cost
```

You should now have cost binary in the bin directory:

```
$ ls bin/
cost
```

# Usage
```
Usage of ./cost:
  -genpk
        generate a private key in hex format and exit
  -usepk string
        private key in hex format
```

# Running

### Generate a random private key and use it for all the coins
```
$ ./cost
coin             |address                                    |wif / private key
Ethereum         |0x7FB55b8280281043FB36bE444EBac2c1C12eea94 |c22de4bc71db1c83f10ede9657c8773757bad57862a64d714afb1974f11e3631
Bitcoin          |1PL3RW8nWESBufESk3P51nqyFbH3wUW694         |5KHodWrz7nRkhVCZGBK8jCd75TqK96d7zs4AEQDpjdYTyMUJ1AA
42coin           |4c2sJNX6Bg9CT8M8wQ3cto2GHdLbeuHfjJ         |5aqm3NTEwBqZn3rCWpH8MmRJQNm4LjN8Hyd2BbizgsbZvLsSbh6
Acoin            |Ae6v4zzQqP6LidSSJh3QAg74jCCkKRPrjC         |8jZjdAeF3BQrTAUVjJd41un6V6eZSvaSG4H97gPG8UE2JnZof9A
Animecoin        |Ae6v4zzQqP6LidSSJh3QAg74jCCkKRPrjC         |65zvgrBiiZyxoMZn7Wb7faQv95zq6J1G6Zgdqr3SbhhPKevTJPS
Anoncoin         |Ae6v4zzQqP6LidSSJh3QAg74jCCkKRPrjC         |65zvgrBiiZyxoMZn7Wb7faQv95zq6J1G6Zgdqr3SbhhPKevTJPS
Auroracoin       |Ae6v4zzQqP6LidSSJh3QAg74jCCkKRPrjC         |65zvgrBiiZyxoMZn7Wb7faQv95zq6J1G6Zgdqr3SbhhPKevTJPS
Aquariuscoin     |Ae6v4zzQqP6LidSSJh3QAg74jCCkKRPrjC         |65zvgrBiiZyxoMZn7Wb7faQv95zq6J1G6Zgdqr3SbhhPKevTJPS
BBQcoin          |bb4K7jUGrZpfRW5ophhAESyqkUBELeyfPv         |8AX5dE1hYwv1ADbWKT54oTavvABMPxFYtC5Jy8h2EL7RvBuWy8s
BitcoinDark      |RXcEW22574EkyfbeDDNC7KBB1rjeapjTvg         |7KvUDRnuFK1Lni4sexZ5xSba2JnQ9MBfDDJcsPVcNwwjae5U86k
Blackcoin        |BSn83DazFk26MVicMXi38vedzCidkD78F1         |69tR3K5nRRLR4zjBvfq7aDcTyKEGPhh1fqprL9REb1iAJsVVTN6
BunnyCoin        |Br7j2KtGxvUyAvrhNx3Md3vRchyaQLXPY2         |6BqADYXpGr1eCpJtqFT7XYDEtRqz3QXtTUtxaJ78afiYoZVezft
CanadaeCoin      |CenvzYUrPHQioo8sRnhzbJTzsiVTocP3kp         |6Fiea1RsyhN6UTUJeQh7SBQnif5RLpDe2m3B4bUvZyjKnmwpTBq
CannabisCoin     |CenvzYUrPHQioo8sRnhzbJTzsiVTocP3kp         |6Fiea1RsyhN6UTUJeQh7SBQnif5RLpDe2m3B4bUvZyjKnmwpTBq
Capricoin        |CenvzYUrPHQioo8sRnhzbJTzsiVTocP3kp         |6Fiea1RsyhN6UTUJeQh7SBQnif5RLpDe2m3B4bUvZyjKnmwpTBq
CashCoin         |F4pYuCGaeNByiPyPaJhuW46iek388ysXgW         |6TP7dQ856FRTHNxY5tR7A71SDLmjF3GtkbTpXVcJXumekT5kBLV
Catcoin          |9qRi6nPqR2Ab5mAGFrNmCRZVUBgrsxR8o1         |627SLPHf1idWXiQNJMM7kwDNJrmPntKWXHYRMYfecPgcLN1RrEq
Corgicoin        |CenvzYUrPHQioo8sRnhzbJTzsiVTocP3kp         |6Fiea1RsyhN6UTUJeQh7SBQnif5RLpDe2m3B4bUvZyjKnmwpTBq
Cryptoescudo     |CenvzYUrPHQioo8sRnhzbJTzsiVTocP3kp         |6Fiea1RsyhN6UTUJeQh7SBQnif5RLpDe2m3B4bUvZyjKnmwpTBq
Dash             |Xy1tFkngTwen4bq2bvhHsKXm5vrjwYdFEd         |7s2P38yQt7pxwqNAAEV5DaBxg8dtYcffoSSLmnVxHS3wUeABecF
Devcoin          |1PL3RW8nWESBufESk3P51nqyFbH3wUW694         |5KHodWrz7nRkhVCZGBK8jCd75TqK96d7zs4AEQDpjdYTyMUJ1AA
DigiByte         |DTU8xm5RoeLUSfR3UdNdZZ1a8j1MDLC4D6         |6Kc8vUKwgYiYk6diTZw7LpcLYtJreDuPc3BPYtriZHk6n1LiE1n
Digitalcoin      |DTU8xm5RoeLUSfR3UdNdZZ1a8j1MDLC4D6         |6Kc8vUKwgYiYk6diTZw7LpcLYtJreDuPc3BPYtriZHk6n1LiE1n
Dogecoin         |DTU8xm5RoeLUSfR3UdNdZZ1a8j1MDLC4D6         |6Kc8vUKwgYiYk6diTZw7LpcLYtJreDuPc3BPYtriZHk6n1LiE1n
Emerald          |F4pYuCGaeNByiPyPaJhuW46iek388ysXgW         |6TP7dQ856FRTHNxY5tR7A71SDLmjF3GtkbTpXVcJXumekT5kBLV
Emercoin         |EfUwv5yHwBj6txqJYtNb1vpw2EnBQAxe5A         |6RSNTAg3EpkE9ZNqBJo7CnQfJEA1bLS1xxPiHLvQYFmGFrLDUKx
EnergyCoin       |eQRY1VZHpq4o9Y4Qze2PdKtM9zyqLcpDLf         |8Q8Hrr9vWveb6xfSfWR4UhnMKxVNwt9gPfa4gBWJBvA9Nd2MUW1
Espers           |EfUwv5yHwBj6txqJYtNb1vpw2EnBQAxe5A         |6RSNTAg3EpkE9ZNqBJo7CnQfJEA1bLS1xxPiHLvQYFmGFrLDUKx
Fastcoin         |g1mwwvkSfYvJRGcm6KMfZpyVg21cD7M9aK         |8XuGZmx3vdMVeEzGHpu4HzBSzQxFYhXBYDrVenFtAYBhM2Dgc13
Feathercoin      |724VD2JpSkvTMjBf5v3XoYez4etFtKoS6K         |5nWE6m9S3jtvayLRxJ185h1wu4TNExRP1p3feVrNeodtt3MAk6v
Fluttercoin      |FUA9tJZsMYerXq7Ubj3DzBNWHFJ4pz7bae         |6VKroda6wg6gRCYEzU377RcD8TPStk7mYEXvmeJCXZn3F5weD6s
Freicoin         |1PL3RW8nWESBufESk3P51nqyFbH3wUW694         |5KHodWrz7nRkhVCZGBK8jCd75TqK96d7zs4AEQDpjdYTyMUJ1AA
Fuelcoin         |FsVksQsA4j7jMGFZd9NYUJeHukZ1b2fG45         |5KHodWrz7nRkhVCZGBK8jCd75TqK96d7zs4AEQDpjdYTyMUJ1AA
Fujicoin         |FsVksQsA4j7jMGFZd9NYUJeHukZ1b2fG45         |6XGbys28o6muZ27wu3f74kCz3a1AYSxeKsc31nz6XDnRjo7Nj3t
GlobalBoost      |GgAxqdTjV63Uz8Xjfz3BSZBsAm4tzxHoed         |6bA6LKvCVx8MpfHMiCu6yPQXsoEbqrePu9kFW6MtWXoCixPDDDM
GridcoinResearch |SLHSUEceXRAWcXspG42q5ZikGsFY2ds3D5         |7PoxZtgxxAMo4MEHU7o5s5o7rY1qSksQnVSqMgsQNFxWZtVGBcW
Gulden           |GgAxqdTjV63Uz8Xjfz3BSZBsAm4tzxHoed         |6bA6LKvCVx8MpfHMiCu6yPQXsoEbqrePu9kFW6MtWXoCixPDDDM
Guncoin          |H5WZpjm2CGWMoZfphQNVvgTeoGKqgWk4Ud         |6d6qWZNEMNoaxUs4cnX6vi1JnurKVZVGgnpMkF3nWBobDb7resE
HTML5Coin        |HUrAor4JuSyEczouiphpQojSRmanKxZ2E1         |6f3agnpGCoUp6JSmXN96t2c5i2U39GL9URtTzPjgVqoyiGug4Hi
HyperStake       |pTscdD1VaLehbNYac8MMkTh1tcRR8fypGe         |9EiuGeNipZZFUUC5Kzw3KimiDotLCVDa4eLkmvhi3JKqi8FyukR
IncaKoin         |NiF1cFw48nzdFdd33H2xiSGfcKw3Zv3Hjj         |76KFyoegHLGkqxzwJuD6HCQ9cWUNbRHXhjosALgLRMu28GnNKtK
Influxcoin       |iRoZraYAvdhZKsTHEqMaUacDT3ZGQx2SbM         |8jZjdAeF3BQrTAUVjJd41un6V6eZSvaSG4H97gPG8UE2JnZof9A
iXcoin           |xuyHJVGYV8P6mUUQ7wM3w6QY7CqE5hzamn         |5KHodWrz7nRkhVCZGBK8jCd75TqK96d7zs4AEQDpjdYTyMUJ1AA
Lanacoin         |LhYzgiScatgFATvbvBNNHoujToeL57768u         |6vbY6eQX2CtdAs6Qn176WbQH2wPnLu59mYTKwbErT5s5fDusX2f
Litecoin         |LhYzgiScatgFATvbvBNNHoujToeL57768u         |6vbY6eQX2CtdAs6Qn176WbQH2wPnLu59mYTKwbErT5s5fDusX2f
LiteDoge         |dbkL3GxiQU93WfnEwoMkf5LmtzTwxd9YTM         |6kspDVAMm5WVVmBtF716jzQQTNKC6NrmqM6niqoNUoq9C6L6y95
Magicoin         |9S677g6YhqhiGL2BES3SiJHhqgRv71qYYb         |5zAhA9qdAHxHPtpfPmj7occbPk9g9BUdjeUK7PykcjgDqpyjqNH
Marscoin         |MWECew3C1FbzoLCmy231G4TJipADWFZV4q         |6zV2T7Jaj4F5SWFpbAM6REbpsAdDeJkuLpbYRtceSPsreVJ3uqN
MarteXcoin       |MWECew3C1FbzoLCmy231G4TJipADWFZV4q         |6zV2T7Jaj4F5SWFpbAM6REbpsAdDeJkuLpbYRtceSPsreVJ3uqN
Mazacoin         |MWECew3C1FbzoLCmy231G4TJipADWFZV4q         |8XuGZmx3vdMVeEzGHpu4HzBSzQxFYhXBYDrVenFtAYBhM2Dgc13
Megacoin         |MWECew3C1FbzoLCmy231G4TJipADWFZV4q         |6zV2T7Jaj4F5SWFpbAM6REbpsAdDeJkuLpbYRtceSPsreVJ3uqN
MintCoin         |MuZoe3LUiS4scmLrzSNKkBj6MKRACzSVLX         |72RmdLkcaUvJaKqXVjy6NZCbnHEwJ1bn8Tfeg3JYS3tF94TieaY
MonetaryUnit     |7RQ6C8c79wPLBAKk7LNrHfvmhA9Ca581Xv         |5pSyGzbTuAa9inv8rsd831cipB55tfGFoT7mteYGeTeHNf5VUbp
MoonCoin         |2bLrNq2edmppMxehpJP2UAfL973t4ySxgA         |5R83ADD5g4TS6wwfyvB8bARRpogU6D9kMnGUxrHWibZdTEfB9J9
Myriadcoin       |MWECew3C1FbzoLCmy231G4TJipADWFZV4q         |6zV2T7Jaj4F5SWFpbAM6REbpsAdDeJkuLpbYRtceSPsreVJ3uqN
NameCoin         |NJuQd9dmRcXkSCUx1rheEJzsypg6wYnmbx         |5KHodWrz7nRkhVCZGBK8jCd75TqK96d7zs4AEQDpjdYTyMUJ1AA
Neoscoin         |NiF1cFw48nzdFdd33H2xiSGfcKw3Zv3Hjj         |6xYHGsrYsdZrJgg7gaj6Tv13x41Vzbv2ZBXSBjvkSjsU9rfpR8C
Nevacoin         |NiF1cFw48nzdFdd33H2xiSGfcKw3Zv3Hjj         |6xYHGsrYsdZrJgg7gaj6Tv13x41Vzbv2ZBXSBjvkSjsU9rfpR8C
Novacoin         |4c2sJNX6Bg9CT8M8wQ3cto2GHdLbeuHfjJ         |5aqm3NTEwBqZn3rCWpH8MmRJQNm4LjN8Hyd2BbizgsbZvLsSbh6
Nubits           |BSn83DazFk26MVicMXi38vedzCidkD78F1         |7Rkhk88zob32CAozNhR5pQPtmedZ6TiHa8WwbqZJMuxu4YdLtEM
Omnicoin         |ofCQezQv9yiwxWGQZHginD9SdbuXe3AbTE         |9AqQvBUf7iCoCq2fWqh3R5aAPaetu5XpVNCYHdKv3zK4isPEMQg
Pandacoin        |PWvDaUXdZ9vNtVuD67hbggpEsLSw65kZuL         |7ACkLGYjzBdD7cAM84T6BqbhSjhotpyHH1x5ee48Qfuo7XVeJ7P
ParkByte         |PWvDaUXdZ9vNtVuD67hbggpEsLSw65kZuL         |7ACkLGYjzBdD7cAM84T6BqbhSjhotpyHH1x5ee48Qfuo7XVeJ7P
Pesetacoin       |LJDPhc9KsiDNM2nWtm33ogdwqJPPMQ9Erm         |6tenvQxVAnDQ33WhsRV6ZGoW7pn4hCEGyuPDhSYxTRrhAdDi9K8
PhoenixCoin      |PvFpZapvGLPFhw3J7Y2vAp62VqhsoQfTmp         |7C9VWVzmqcJSFRk42e569ACUMrKXYXpA4f2Btnk2QKvBcA34k2p
Peercoin         |PWvDaUXdZ9vNtVuD67hbggpEsLSw65kZuL         |7ACkLGYjzBdD7cAM84T6BqbhSjhotpyHH1x5ee48Qfuo7XVeJ7P
Potcoin          |PWvDaUXdZ9vNtVuD67hbggpEsLSw65kZuL         |7ACkLGYjzBdD7cAM84T6BqbhSjhotpyHH1x5ee48Qfuo7XVeJ7P
Primecoin        |Ae6v4zzQqP6LidSSJh3QAg74jCCkKRPrjC         |65zvgrBiiZyxoMZn7Wb7faQv95zq6J1G6Zgdqr3SbhhPKevTJPS
Quark            |Qiw2XoRVghK1LoKUANhZ94dbkrDm8GVYJi         |7G2yrxtqYTetX4uTqoK63oQ2C5YxqwVudwAQP67pPdvxbQbizsv
Reddcoin         |RvwqV8KMpEhdo6jjEdhWbSSxeMzbKLWf2G         |7MsDPfEw6jgZvXeaZYB5umCLwRQ7o42XzrNj7YBWNbx85G7nkmX
Riecoin          |RXcEW22574EkyfbeDDNC7KBB1rjeapjTvg         |5KHodWrz7nRkhVCZGBK8jCd75TqK96d7zs4AEQDpjdYTyMUJ1AA
Rimbit           |RXcEW22574EkyfbeDDNC7KBB1rjeapjTvg         |7KvUDRnuFK1Lni4sexZ5xSba2JnQ9MBfDDJcsPVcNwwjae5U86k
Rubycoin         |RXcEW22574EkyfbeDDNC7KBB1rjeapjTvg         |7KvUDRnuFK1Lni4sexZ5xSba2JnQ9MBfDDJcsPVcNwwjae5U86k
SecKCoin         |Sjd3TLuwEbdPRy1uHUN9ZgzXuNWUm4XfC6         |7Rkhk88zob32CAozNhR5pQPtmedZ6TiHa8WwbqZJMuxu4YdLtEM
SibCoin          |Sjd3TLuwEbdPRy1uHUN9ZgzXuNWUm4XfC6         |5KHodWrz7nRkhVCZGBK8jCd75TqK96d7zs4AEQDpjdYTyMUJ1AA
SongCoin         |Sjd3TLuwEbdPRy1uHUN9ZgzXuNWUm4XfC6         |7Rkhk88zob32CAozNhR5pQPtmedZ6TiHa8WwbqZJMuxu4YdLtEM
SpreadCoin       |Sjd3TLuwEbdPRy1uHUN9ZgzXuNWUm4XfC6         |7Rkhk88zob32CAozNhR5pQPtmedZ6TiHa8WwbqZJMuxu4YdLtEM
StealthCoin      |SLHSUEceXRAWcXspG42q5ZikGsFY2ds3D5         |7PoxZtgxxAMo4MEHU7o5s5o7rY1qSksQnVSqMgsQNFxWZtVGBcW
Syscoin          |Sjd3TLuwEbdPRy1uHUN9ZgzXuNWUm4XfC6         |7Rkhk88zob32CAozNhR5pQPtmedZ6TiHa8WwbqZJMuxu4YdLtEM
Tajcoin          |TYJFRZWWexZ94qJ5LK2nXwY7AP2NCFPxG9         |4kF9daESdYvuQYKZrKm9WkRwWXN768JEczrL5rXaqVRsapfFyaQ
Titcoin          |1PL3RW8nWESBufESk3P51nqyFbH3wUW694         |5KHodWrz7nRkhVCZGBK8jCd75TqK96d7zs4AEQDpjdYTyMUJ1AA
TittieCoin       |TYJFRZWWexZ94qJ5LK2nXwY7AP2NCFPxG9         |7VeC6b34WSPUToyQBrf5j3bSbsrzPsQ39QfA68w6MDyg3p6NyBm
TransferCoin     |TwdrQfooN921tGSAMjN724otntHJowENVJ         |69tR3K5nRRLR4zjBvfq7aDcTyKEGPhh1fqprL9REb1iAJsVVTN6
Unobtanium       |uhGTRctEogg6E1MhvagW46EF5AmgKeFp8h         |8XuGZmx3vdMVeEzGHpu4HzBSzQxFYhXBYDrVenFtAYBhM2Dgc13
Vcash            |VxKsLDJEv3LPyS8bUq2hShApwQa2P423Zg         |7hJf9yjFczSqGjTddLP5SyC66ZZJJ6THsF5oZ34UKA211YVJ22K
Vertcoin         |VxKsLDJEv3LPyS8bUq2hShApwQa2P423Zg         |7hJf9yjFczSqGjTddLP5SyC66ZZJJ6THsF5oZ34UKA211YVJ22K
Viacoin          |VxKsLDJEv3LPyS8bUq2hShApwQa2P423Zg         |7hJf9yjFczSqGjTddLP5SyC66ZZJJ6THsF5oZ34UKA211YVJ22K
WeAreSatoshiCoin |whxUMANgMazUKB493gM6UibBDh4PuQkn6c         |65zvgrBiiZyxoMZn7Wb7faQv95zq6J1G6Zgdqr3SbhhPKevTJPS
WorldCoin        |Wm15JRtpLQG9cJQmXfhLQwiQCR5uoXiQcE         |7mC9WSdKKqoHYNd3SVd5McPdvnnjbW93SXE23LSGJU2mznTnF19
Zetacoin         |ZaNJCByqJfWHLLPNhc2ZopcubwtWnDY2xr         |8XuGZmx3vdMVeEzGHpu4HzBSzQxFYhXBYDrVenFtAYBhM2Dgc13
privkey: c22de4bc71db1c83f10ede9657c8773757bad57862a64d714afb1974f11e3631
```

### Use specified private key to generate all the addresses

```
./cost -usepk c22de4bc71db1c83f10ede9657c8773757bad57862a64d714afb1974f11e3631
```
