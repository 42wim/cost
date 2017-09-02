# cost

bitcoin, litecoin, ethereum, altcoin address (pubkey/privkey) generator   

can create bip44 receive addresses based on bip39 mnemonic seed words

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
cost

Use -genpk to generate a private key. Use this output as input for -pk
this will create a coin address for each altcoin with the compressed private key
eg ./cost -usepk "$(./cost -genpk)"

Use -genseedwords to create 24 bip39 seed words. Use this as input for -seedwords
This will create bip44 Hierarchy with 20 receive addresses for bitcoin/litecoin and ether

Usage of ./cost:
  -bip39
        handle specified key (with -pk) as bip39 seed key
  -count int
        create this number of addresses for each account (default 20)
  -genpk
        generate a private key in hex format and exit
  -genseedwords
        generate 24 seed words (mnemonic) and exit
  -pk string
        use private key in hex format
  -seedwords string
        use bip39 seed words (mnemonic)
```

# Running

### Generate a random 24 word mnemonic seed (bip39) and use it to create deterministic accountnts

```
$ ./cost -genseedwords
hire warfare legal voice dice meat machine equip anger income scare hazard favorite poet current evidence rain gown injury thunder grace collect foam flight

$ ./cost -seedwords "hire warfare legal voice dice meat machine equip anger income scare hazard favorite poet current evidence rain gown injury thunder grace collect foam flight"

bip39 seed words: hire warfare legal voice dice meat machine equip anger income scare hazard favorite poet current evidence rain gown injury thunder grace collect foam flight
bip39 seed key: 117b06c79833d2843a1586736cb415f4f771e612e126882188c48a45481f3f2c4b5b2bbf9db57567d99f063f7b74eff6aea4bbb6105ee2675453c0c632274799
bip32 root key: xprv9s21ZrQH143K4Ve7hQDQx6mDbYoN3K6JQtrjEGuGwddbN9CXbEg8rQjyhdKNgw3JkWejvWQmJb8aKTsdCtDw6PvyeQvaZ4pjW691grR9tYn

Bitcoin
account extended key: xprv9yFZjLzbibfBVFxAkJX549rQVoNMz8JaTCsAAnb9k9zq1cSHTkoC5ZJDuUk13iJaY8Db6coHF6dBXL1vfKpbpwQn8UHUVvVZgeb7FGM2Mjp
external (m/44'/type'/0'/0) extended key: xprvA1z5xBPBhKtBgmkicNpCbVAKMXLWE2ij2JWqe3YRrAb2SVfuxmrdhgbnULjTpQU6adt4m2qhFf9kL6QNgYeNywL3dDa5ZWgKn6p9xaY77RA
0  |1CW3oddC37pUzFdGeuuJeX36ZUAoFpVhJc |L58bhhRH8qmVBrqNNhe7jbpnbC7A6r6EPmFgHLvzSufJ1jBQWTaD |0216e11316fc42c6a1164b6970eb38b9b7a7071d33a75a6bbad10beaf076c40156
1  |1QBVadvswk9hQTmajJTZ6yc5UwZSNiPhmn |L3FhubUtLp2vLsAWiF9gg2Dtfk2jx5Mm7gHsyQ8KWFQTJqXJPTNc |0209b819f9fe2dd5d9b8621fa6ccb27c1a54fe4100b3e1b928dd909f8ffa51fa56
2  |115uPqkQ8yyYJ7Yg3ejHhgFbEwm7jhbRHG |L1XaYLXGd1L2Ye92xnyJMQRnYdrB66oj4wQh7Hu362xaCozuQpZh |03c643963daa8bee73f3227cbe231d784fc90ec4e1bc9477de94b91841050fe0c9
3  |1Avfs2rJdfNvEkPApuD78sQPMBRPfAEguR |KzXMjARVwegPbbXkgwJu4BTDhZZEEAQJ4ezoT5eJVrwG4AcGu1WP |03e244c5a0604370a8bb224e0fc99b3ecd425cda75b2c97643330fb4481674813d
4  |1ZnaXAz7oB8qLYQSqYDzRXFC3LLbJnpQK  |L3xasQT599eJFchUUQ7vfBa1Wn8UZuNLWkigLKdoH4mDrPjdh7M4 |02bef5c01842cc0c8a0bf1b9e6b5fac628694885d44817001b1e1614074c6dc53e
5  |14nv4dUsTNdPREMJRXa4mbScb64eESdNu9 |L2z2HYWUEHDaiLwFxvJ3BM8LXBAyXB4ndp8ubdzwYaYPTtV2Gju1 |030f4ffcef2ba4b26ec680a8363d33e380a3d7fd0bb4c3cc09d6c05abf87e039c8
6  |16x3VvPQvHES7qh9uVwjXoYZrbukYU2Bbi |L1w1Vd5f2i2MFNi8LR8hvpYXoRMFfHEhaDPkCxyJeKWehq5TWNeL |03ad4d3b0aab5fe9b4c2a0cb6e8eeb03f935e142084ed14cbd046b6b8a51c3494c
7  |1LPCaBKCwhtTzHd6fCd3T6R1cc5bBYVpUq |KzZZm4HaZbgmymRb2936ppAPriFeV1xpNLLXAWcFezwqpmDLWUx9 |031a49e11932215b94f5f387807f34e1ad7bdc03394613394f4c79188f77ad21f9
8  |12xSn9imFPp2YrA7LkzjqKnB1nNnf9Tgki |KzsgtRs4umBxUcakPBvs9dUnHsAkqLDCnvCdTKWH7nZdqRHeRe3i |02177a9963ca661445c2c3aa67716545904272ff071230423e316c40ff837d32ae
9  |1NuEcz6L69XWKSaRvBUKEbN3L3VcVbhFG1 |L2feFDSmnaoVAPhCTwCPhy1YjJuUaiyx1BUoSP9LniYcjMT6XCXc |03b7a5c77c7a1a1774887c720e0bdd121fb065c540deac4c312221391c3aa9294d
10 |1puRL9tt1nDou27RccNVhzxG3SavkNzV8  |Kxdy3T7AHz3jtFTvfhD9TGuRmdSNSDh3TLYwqzN73BVcPGRn4odp |039b23d5398afff3c05df934c919fdd1e43a7b8ea4867fff031933a87ae1684275
11 |1A9r5JaUvfstGfd4ApxpXHAakPuY6FAXZ2 |L3Zbnx1YR773MAX8xMmNt1MYggN6b2CNRfmfgdGCtwtwdt64USDY |0220c976836fb8dff75324d87facc41985ea8490f3a465c97652ee95754498e724
12 |1PUKfEZp7c9dqvJardxuxymj5A4eva6VqC |KyiDoURjqyUV5fxxbcMLD91X3PwciWGiziNCoERVwhTxpTnXBqYR |0361018a9630b92b10949dc68fae7161aaa75bcdd54d810050ce9c05e8584561e6
13 |1FpZtPRtERNr2CknRbwbZJUUFbfWKZZwet |L4nLGRBUd7rqGJEUBKkR21ahzdzQ56tnZjkrgy5g2uw9hQb7gHsf |0314ef837199a2ce01e8ab28df3d97e9fb68f5560cf0a0e2b83625f98a4c601d2f
14 |1PpfaYzLGqvHNW5a8wEQJg5iEtQbK2yGCn |L5NPc98XT4mDHRatFw1o4WFCovqr1PRWUbDv3vBwLodaGX6S83jH |02cb2575610a8335e24be2d22695867a3ceeaf4d41baddbee06ea006970d8e3c41
15 |1NMyrbLfNxAHhXo8S6juNe47GHWoPZfACs |L1eANMX4JGxN1LH1SvqV2n269gy8ivSZ8VE5pAHT8ehzjMPSGNLT |03d6336e47d0ed00dd8c55cb24a8467c69527b99c6316b03fcafac130f1ca6aeb5
16 |15DZfgqLUi5ddFXbeEFUeYy84mqJ4N8VjC |Ky3RM1WokdVB9WW8msBBfvsy93fvybc1nfe5v4K1ZoWgufNRV7Bj |02de42fbc70e8551f99901d193d9e365b7aa4414b0329592f6baf058b71bd9aa31
17 |1Q2tXL6Dz72dMdm9gnSkEpt2FwRCf5WxvB |KwmTERdbxFWCodw3rV7scxZfbsGaEbCRaaLYFcmGZKGNNT1Hx7mW |0276585c544894e1efe48e701994d91149dbb4a670dbe60331723ae53793efbc77
18 |157hYoD77PRq7ghjjLC1LWxqvPnTrpyuDo |Kwpc5XAzdq5jvh3QCsSXmL8M2WvMeBT3TYCKmrvCWaDAwPo8sBHZ |02d1adce5754d89c5d85a8738cccf1f52d283c8cd08b274997ffbdef9a3af91e8c
19 |1JJQ758QRvMGoEADKppjH3hfMvioDxHfMR |L3Z15DinPRXWQBqcawD1ZCEBRqZtMr2jtYHT4ePiq7VhZRQDJNQM |024fbb8afc5ad8f4f4017c3e90a023bccf41239bf1d754d3f72fa59a9c27768b23

Litecoin
account extended key: xprv9yc3vS8v7eHHYdxVzN42FjchCiLGPKsaNaae5C7DkYVeco8DRM98TVywEK5WK5tRMuS4XMJKiqi8FbHRDT5SrBjdDEAaPUQs43TEynSG2uo
external (m/44'/type'/0'/0) extended key: xprvA2RMVBJkSKmn59LzfQeWfoT1mhw8dw4N7hpoiVy5QZ7T8an6dx1GcDphCbxag5w8j4JstdajfUX1yWRhHEFXMBdMDHjUvaYM9xPXxVJRtN3
0  |LLdu8g6y4X9vBtqYt3jiGerrMsTYhN3Hva |T3sEaVGGDjDGnN1qmMZub4ksMR6PMCSLTUWtbGpdGtRMyDTExcpV |021ae96e93f336383a6da40a97501932109396f08904aa33a7e58401db36562793
1  |LZk2sqQztFkVLnj61s32tZaGZnN49jHEdF |T8sixM5tGe4SxqhQmRQWqEWBgdH9HaPnAc2nwEqkVkJU8oMdhqsZ |02711696d14dde29dc917015f0a210015e79e9eb1eba3223ce06b42b348ecde7d0
2  |LdPLk2afVc96PUHTso5UC5H2CwASV7LwdB |TB4mUT3JmJQMfAhgzjieM7KgHuqNJdqgHDxceskVZcjWdAucECxz |034fbf2c279430a08edbd41253746efdabb5305b141eb418924a04170eb7121f3f
3  |LYkDRMz1gptYAXjPj4Mmih27rANJr2sYmH |T9HwC2uGuHqq8TZriBv1UfqfXD2tnHdyXRy8cJhquQuJBo4v7AMT |03730d21981a4a0c6fa451b24e622f53e666d3d4900962516065d266365d1f7e06
4  |LeoMZEw3rM74zSRfH95LMmfVqpVs8oQZeD |T7scoHTvKUoLR86Bf9sogugnZqRDoCEdzB2geHCTRoNQ3K38rqin |0243f726faa22a0f0eda8d5c92ce00ee0b9d2ea9e130f2823735ed7e5be894b127
5  |LajzExfhEwvRadyLx2W6J8sjdBK9HsZsvv |T8ZhgHiNtiVo7jzSanKA5gTrbVvUfTixTLwTaCC8mTPiUAW6CKeA |02574266afdb5c924f75fca47a192edf71d841f51ff19297a26e496567b5154f9b
6  |LPyrdetmFnJroUZxoC5Gk7Yk2BxaA2cTvP |TAHPM45AMJUXeGGCboadayxuDDQP6pUjMCJBfHjAXbzK9eKeWni3 |038f0fdd3c8eaf507fa378af695f714f5a850023d774d0a9c79ec7b2a7a98c5737
7  |LLQjCktKpw6Yu47TCfYnfrjrUEW1ttGpRM |T8MKz2xm4Yy3k4uHbtFQLFaFQNWKTSTwvsY2HqAqRKfsXQAtfiSJ |03d4919f98ed9d2fd95e6e14586fbe136b7c876b53c712ebef8e109ef3471e8bfd
8  |LR4V9jmP2pRcL1fCCiutseGeV1TjjuAo6n |T75SNwH7L2Xm8m1GZdMLie1Akrh18qF9PnuCHxAc8Gh4fE34eznQ |03c2e87bef4f012815eeb1a284c99ce58f0968e48144bcd0f2ec79f6d8787c18df
9  |LWV1QmxXx8qdVkNqJoCyu2ZNAe2bNhffuH |T82opY1Xiuf3zTXHNNvSpKnnwBzTryWsgSVQX1YvFhSxydD4gDFL |03e0e1c85c7df40acbe8ed89fb7531c9cfd897c5ff8af04552bfa24885358ee21b
10 |LabvgGAvTTyW6yyhazXToGYgKozax6hecx |T3RzkZeaMB1pMsZEQEPPJrumZTZVfE16W4pgFZsqwFHBWSXagTsk |03534bdefa581ac80917ebe5d7767b9d4eaeef08aa772fe053d9787110a3f608f4
11 |LReCGSkgHkrxxmQbtGtk3cP96LzbyKmAMw |T8Xgh3eZQ3krABRgEbs9jE5CQH5xYoqe9rQocAQ6E6hdz2RrzWTE |034f65347647c0c09ac3a2ce9a8ba4a9c6db1f7758f34888cbab44c737cf943e4e
12 |Le25eaokiFUuYLj9C5PrKNVxvSrsDJ98Ws |T7Q1Wiumn7J2zGLqc4oK5e3fUtkoLZok8xCCpPEqm4ka3SfsGTfz |022d5b74e15ba7b47ce9e72edc3c4bcbde1801831b1aab5e9b99160286badb3242
13 |LVCeAQggEPxx88rsxXJWK5hrJjygjyrHL4 |T7mGbMY6yBbZaCjis6TUsUT7VsB9zjGBZ9Z3DgAvKcZkGLfmMm8o |0266ba272f4eb707f2565b1bc670b5612ce8dee09409b3c5dba95dfea91f9acb54
14 |LhWp7bLShMThTPcQ7T9rAxbMV9YeoQUP8V |T43ZrWgSsb8HqXZS6FXSbvqTWjPLGJWjnT9efKytcdbiEcSfoT7Z |031e870fc445a1699d135754f490b8647d6f354de81bd1a9484c654bcd8b68e2ed
15 |LVJxcETrS1N9UXTn92EMVnQVi34juY1ueA |T5G1Lq5GrSV29HzaroJUxgiLGQnB9asdq6ouiQhKmrwbvqv6GnAy |02b8c81216c881fe14fabc47f10c4ef59f078035c07bd3d948f0750c024a554ac1
16 |Ldu8sN12mDkkTZqnQKxhJnyZgMd8niKdnW |T4HoYZhyvzKCNH3TPS9JYMqCbJBeBKdk8GZ37uNYRQz7nkaumhck |038ad676283a1c8d6060ff388d6898296ac79f0592e15b79b0461f3aa016418dc5
17 |LamcEfQBCGhDeetNpT2L6P6DMB2pEazi9r |TAGRhRYBNsoUZpR1tqVJWrWnEgqPU2eyrqxCyF62jqGfrHzAzeEx |03d7eb67d58b6a0e7d61c7d5a857b499914946e67bc721e30d821ad50b75f058ab
18 |LMF8GjHEPDPXi1UyWMg6DNUVXMem7M6qwn |T8Ey3munHfQcftzPt3RRi19cpBkj5frMSZK118uucvijcg39birN |0200360433c86a70e222d2beac51a5e996386dd6fac127352aba5730fb879273db
19 |Lg7bNaTUw2TCub5RbfoSpTm8Fxsa6KRUAD |T7pmG5chiCLrvWhxG1zXocwiK97WxzvBnhRNvfJyuVnxV4nF99iL |025de41030b49b960e98681e956537fb1fe3664086c9aed994ca077f3398acefe4

Ether
account extended key: xprv9zXBePh3mvjY7qJxQzpFCxLxpYVgDBo9sw6u5QzaMCkSwx3e1RuJR372fX3R1D7ARKqfRzi3WHrT6tNVae5kyNqfXHiaiuiMAoF5YdcURXj
external (m/44'/type'/0'/0) extended key: xprvA14eHBPxLa1pFFrnvYxworBn4uuxwADBBs4PbfSf4QT8u9RzkW4a9NMYBAohqURuTcwrBy5cJ4DT7yQ3Ap9wahNQeBr9M26mcHNZyN9q6rn
0  |0x7115227980EFB6acF699812F1682c646eC2CadbC |0x0f743bcc48cd28aefdff899e32b2bda3598d2833e8c68eb54221027cd8f8cf19 |0x032d5c25d94287fb22eafe51aad50eb96fef6ee7f2ff218beaf895d12af75ae1d0
1  |0xB016A46CF299dFF589fb2Cfa4c3d83D95Dc88ddD |0x467f2354eccf3180dadc4a1d2abae2bde8bede75a80e0524b0d3dd94abe38028 |0x0253e219d4264f27275e0991ae8008a05c6a389440147d9e2976004512dd508b31
2  |0x1145cdb04b99374C2dd5Da29C989eaE6BdCd6074 |0xd40475afdfd4862d83f721fd3d7c6db099709e5676645f464e850838ba3b70b0 |0x03aecfd7c5416f478b4324716c12b39b645382e8442a447887329e48b3a8760f36
3  |0xa9268ea38529b908b6Dcb2cF8d547d6463738f79 |0x52d97253479a1594e05afc7db33958249bf89ced91f0d52a11fb75ac71655c3d |0x03f09b3884509e276c256febd01e1bc69bb63ca763960b042512ebaa8e585b46f3
4  |0xABC73696064075263B27C636865Cc9D0c48362b1 |0x3bad99d22c2945c8ad43b3545bd3037ee4bc1cd7ba75c3344336ebebdd57f4b7 |0x021ec6bf7b4272cdcb28ece9be7f79476e47053a07a33d353e6ba0791b0dbed541
5  |0xEFC7a410f9cE0E54a20Fb92516709D705F850d2d |0x99ca22bbb86360c6ee02b87cff99016fb07f618a407278cdfef962c5f0532324 |0x0352032ba6a2321d39adbd78bc8497c73fe082a626e2acc46341d20caa1427f286
6  |0x4353F9a6475A8360E98F67411d07607Af6b1e495 |0x568b7eaa790d04758137f861d0f55e95f84e0486627d973491f35f040c63b27c |0x03cc674c61d11b9d267272e010644ae31b36e7d3f6b4c919a7b0831385da56cb51
7  |0x95e92365824f2A625b491cC567b1915669b6b238 |0xfa2736c95747b623ac10f296bee1d06d0039d2ea6ff14ef6eeb459d208781d4c |0x03640ab1202c62c34054c3e1c94b4bae803aeab9fd6ec34a647ce9571be083e094
8  |0x16e550470cCF8b17F13aa1ec6Bc16bfcA3Ad0136 |0xed8fd170fbc32fee804307a6b4011bdd603d4a85370b02db0902dc484e43d449 |0x022764871545412b9d666d28052c468a7eae3245bff2b6504168404c2ae43cfbe9
9  |0x3C8350D7A590D20Cb5b1e794941E9F9d2108e0aF |0xee59c28f3e78c77bb7ef86f66d7736907680c1c5769e81dd4a9c70412d260b95 |0x03e013f523a8b4c7d99dea8f52995d365f97bea6fd5519ac9eb48e11a8c376dc02
10 |0xEdb390c6B0B6001A4D99C9fF8e3A909D12e5e1C4 |0xea3014e1954298f6e41c5942754dcec5cfcc892eca3e237e11bef13f1673470f |0x028ee6a1f742b053da0a332ec0bb2249d49594455df988161784dad36e915a5947
11 |0x1D99fb0D6166912ce357B2a3225d2eabAb60f25B |0xe841c1bb9ea44d1a208c50e78e4d51684de0ebcd6371e9465129408fa9dffd32 |0x034c048778022f9ab3cc12e2c413313276757f177f77631c91a5e936f48057942c
12 |0xBd9a0c51240BD7aa4736074FD4614Fad2aDA4C61 |0x7567e73278ef3c0ba71d776a97fc7efd14710c34166dbca70d1f67109c1fe957 |0x024901776a788bc4a848eeccfefe2064dae809edab26c4906d54b9ceb4e228e996
13 |0x5B182c9aD0F4CE3a2055fCbee0ee49bC93405B89 |0x66e1e32c2bae150d1dbc6cf3fc4462aedea6b39a9312364c9e1427571b90e24d |0x036eb30457462dde7705e13e918815eeece2f7ad9cfe0740b500fe472fc1c54206
14 |0xfc0427798c19bB1511ded3861634F9FfDc950bd8 |0x883b209061d081668536097633c8dd62849f5dec210a2fec5a64bd9b4f5fde01 |0x02a79a90829426ac3f56d099da81bd9ecf2b70578c5f378f56dc00aebe6b398d66
15 |0xF1Da49368710c9166A54A5FF0490bD81FDC97897 |0xa33b42fcfc4d3ce063ee472e6815445e1ccf429ec8f484b4082a401c753b135a |0x02ce4778fe14d0fd0ed077712aacb67ce48b9e7cd249d93d0fd7fbeda819e848ae
16 |0x7f7879CB12D97bf9B3b9D16DD2C903B761681d5a |0x08cc21112ed24f5448c5d6c5a806e3af4007e42bcbdcac22452827f5fdb6902b |0x0225494e7d47b11c91de797d287d567a843cf30963e6759718452a8f8f3f068ad5
17 |0xeE37a7D9D266DB61408e8e58c4a689e77c7A11c5 |0xd32fb1fb797813eed3a1640356e756cc1c2efcee91bce7577ee7a32be2bedd2b |0x03f9dc2e8311f5f0beeb17fc071805a82768706a868cca116671ee91c630f698fe
18 |0xB88722b047319048c200CA49a1aa06f2BfDD5F77 |0x1a522383a105145bfe543081f086d0f6582ffece9d1472c8ca021c93d43816ac |0x03b85967484cfce41ccb19474c099e4a3bb482dd0881dab805d5996f34dfd6e20e
19 |0x4C11A54D3216C2123f6A5E54a747A04069150279 |0xf4756cdcf652951f571a54131b4bac929cd1bd8ea1ebc9eef24b6ebed01c9191 |0x0382c6c7411926740aa55535f9cfc6d33aacbe9a6677c77240b036dec1604c5dfd
```

### Generate a random private key and use it for all the coins
```
$ ./cost -genpk
c22de4bc71db1c83f10ede9657c8773757bad57862a64d714afb1974f11e3631

$ ./cost -pk c22de4bc71db1c83f10ede9657c8773757bad57862a64d714afb1974f11e3631
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
