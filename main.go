package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/diegosz/go-bip39"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"os"
	"text/tabwriter"
)

var tw *tabwriter.Writer
var pkstr, seedwords string

func genCoin(pk *btcec.PrivateKey, PubKeyHashAddrID, PrivateKeyID byte, name string) {
	net := &chaincfg.MainNetParams
	net.PubKeyHashAddrID = PubKeyHashAddrID
	net.PrivateKeyID = PrivateKeyID
	wif, _ := btcutil.NewWIF(pk, net, true)
	addr, _ := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), net)
	fmt.Fprintf(tw, "%s\t%s\t%s\n", name, addr.EncodeAddress(), wif.String())
}

func genEth(privkey *btcec.PrivateKey) {
	address := crypto.PubkeyToAddress(privkey.PublicKey)
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "Ethereum", address.String(), hex.EncodeToString(crypto.FromECDSA(privkey.ToECDSA())))
}

type Address struct {
	Address string
	Pubkey  string
	Privkey string
}

type Account struct {
	Coin      string
	CoinType  int
	Key       *hdkeychain.ExtendedKey // account extended key (m/44'/cointype'/0')
	External  *hdkeychain.ExtendedKey // external extended key (m/44'/cointype'/0'/0)
	Addresses []Address
}

type Wallet struct {
	Seed      []byte                  // bip39 64byte (512bits) seed
	Seedwords string                  // bip39 mnemonic
	Masterkey *hdkeychain.ExtendedKey // bip32 master (root) key
	Accounts  []Account               // different coin accounts
}

// bip44 creates a bip44 account with count receive addresses, based on pkb (bip39 seed key)
// m/44'/cointype'/0'/0/i
func bip44(name string, count int, pkb []byte) Account {
	account := Account{Coin: name}
	net := &chaincfg.MainNetParams

	// cointype as specified in https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	altcointype := 0
	for _, coin := range altcoins {
		if coin.Name == name {
			altcointype = coin.CoinType
			net.PubKeyHashAddrID = coin.PubKeyHashAddrID
			net.PrivateKeyID = coin.PrivateKeyID
			account.CoinType = coin.CoinType
		}
	}

	// See https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki
	ext, err := hdkeychain.NewMaster(pkb, net)
	if err != nil {
		log.Fatalln("Bip32 root key generation failed", err)
	}

	// m/44'

	// Child returns a derived child extended key at the given index.  When this
	// extended key is a private extended key (as determined by the IsPrivate
	// function), a private extended key will be derived.  Otherwise, the derived
	// extended key will be also be a public extended key.
	purpose, err := ext.Child(44 + hdkeychain.HardenedKeyStart)
	if err != nil {
		log.Fatalln("Bip44 purpose (m/44') failed", err)
	}

	// m/44'/altcointype'
	coinType, err := purpose.Child(uint32(altcointype) + hdkeychain.HardenedKeyStart)
	if err != nil {
		log.Fatalln("Bip44 coin_type (m/44'/coin_type') failed", err)
	}

	// m/44'/altcointype'/0'
	acct0, err := coinType.Child(0 + hdkeychain.HardenedKeyStart)
	if err != nil {
		log.Fatalln("Bip44 first account (m/44'/coin_type'/0') failed", err)
	}

	// Account extended private key (eg to import in electrum)
	account.Key = acct0

	// m/44'/altcointype'/0'/0
	// 0 = external accounts for receive addresses
	// 1 = internal accounts for change
	acct0External, err := acct0.Child(0)
	if err != nil {
		log.Fatalln(err)
	}
	account.External = acct0External

	for i := 0; i < count; i++ {
		receive, err := acct0External.Child(uint32(i))
		if err != nil {
			log.Fatalf("Failed to create %v receive address: %s", i, err)
		}
		// ECPrivKey converts the extended key to a btcec private key and returns it.
		// As you might imagine this is only possible if the extended key is a private
		// extended key (as determined by the IsPrivate function).  The ErrNotPrivExtKey
		// error will be returned if this function is called on a public extended key.
		privk, err := receive.ECPrivKey()
		if err != nil {
			log.Fatalln("ECPrivKey() failed", err)
		}

		// ECPubKey converts the extended key to a btcec public key and returns it.
		pubk, err := receive.ECPubKey()
		if err != nil {
			log.Fatalln("ECPubKey() failed", err)
		}

		// Ether addresses are handle differently (TODO link)
		if name == "Ether" {
			// create our address from the publickey
			address := crypto.PubkeyToAddress(*pubk.ToECDSA())

			// add the address to our addresses, with the pub and privkey as a string (compressed)
			account.Addresses = append(account.Addresses,
				Address{address.String(),
					"0x" + hex.EncodeToString(pubk.SerializeCompressed()),
					"0x" + hex.EncodeToString(privk.Serialize())})
		} else {
			// Address converts the extended key to a standard bitcoin pay-to-pubkey-hash
			// address for the passed network.
			address, err := receive.Address(net)
			if err != nil {
				log.Fatalln("Address conversion failed:", err)
			}

			// NewWIF creates a new WIF structure to export an address and its private key
			// as a string encoded in the Wallet Import Format.  The compress argument
			// specifies whether the address intended to be imported or exported was created
			// by serializing the public key compressed rather than uncompressed.
			wif, err := btcutil.NewWIF(privk, net, true)
			if err != nil {
				log.Fatalln("WIF creation failed:", err)
			}
			account.Addresses = append(account.Addresses,
				Address{address.String(), hex.EncodeToString(pubk.SerializeCompressed()), wif.String()})
		}
	}
	return account
}

func main() {
	var genpk, genseedwords, bip39seed bool
	var addrcount int
	var pk *btcec.PrivateKey
	var err error
	flag.StringVar(&pkstr, "pk", "", "use private key in hex format")
	flag.StringVar(&seedwords, "seedwords", "", "use bip39 seed words (mnemonic)")
	flag.BoolVar(&bip39seed, "bip39", false, "handle specified key (with -pk) as bip39 seed key")
	flag.BoolVar(&genpk, "genpk", false, "generate a private key in hex format and exit")
	flag.BoolVar(&genseedwords, "genseedwords", false, "generate 24 seed words (mnemonic) and exit")
	flag.IntVar(&addrcount, "count", 20, "create this number of addresses for each account")
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("cost")
		fmt.Println()
		fmt.Println("Use -genpk to generate a private key. Use this output as input for -pk")
		fmt.Println("this will create a coin address for each altcoin with the compressed private key")
		fmt.Println("eg ./cost -usepk \"$(./cost -genpk)\"")
		fmt.Println()
		fmt.Println("Use -genseedwords to create 24 bip39 seed words. Use this as input for -seedwords")
		fmt.Println("This will create bip44 Hierarchy with 20 receive addresses for bitcoin/litecoin and ether")
		fmt.Println()
		flag.Usage()
		return
	}
	wallet := &Wallet{}

	// create a 256bit private key
	if genpk {
		// S256 returns a Curve which implements secp256k1.
		curve := btcec.S256()

		// NewPrivateKey is a wrapper for ecdsa.GenerateKey(curve,rand.Reader) that returns a PrivateKey
		// instead of the normal ecdsa.PrivateKey.
		pk, err = btcec.NewPrivateKey(curve)
		if err != nil {
			log.Fatalln("Failed to generate a private key:", err)
		}

		// Serialize returns the private key number d as a big-endian binary-encoded
		// number, padded to a length of 32 bytes. (256 bit)
		fmt.Println(hex.EncodeToString(pk.Serialize()))
		return
	}

	// generate seed words (mnemonic)
	if genseedwords {
		// NewEntropy will create random entropy bytes
		// Return non-zero first byte, unless all random zeros occurs.
		entropy, err := bip39.NewEntropy(256)
		if err != nil {
			log.Fatalln("Failed to generate entropy:", err)
		}

		// generate (english) seed words based on the entropy
		mnemonic, err := bip39.NewMnemonic(entropy)
		if err != nil {
			log.Fatal("Failed to generate mnemonic:", err)
		}
		fmt.Println(mnemonic)
		return
	}

	// use the seed words specified by the user
	if seedwords != "" {
		wallet.Seedwords = seedwords
		// https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki
		// to create binary seed from the mnemonic, we use the PBKDF2 function with a mnemonic sentence (in UTF-8 NFKD)
		// used as the password and the string "mnemonic" + passphrase (again in UTF-8 NFKD) used as the salt. i
		//The iteration count is set to 2048 and HMAC-SHA512 is used as the pseudo-random function.
		// The length of the derived key is 512 bits (= 64 bytes).
		// don't use an extra password for the seed
		wallet.Seed, err = bip39.NewSeedWithErrorChecking(seedwords, "")
		if err != nil {
			log.Fatalln("Seed words incorrect:", err)
		}

		// hexencode the seed and give it to pkstr
		pkstr = hex.EncodeToString(wallet.Seed)
	}

	tw = tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)

	// we have a user specified 256bits private key in hex OR we have a 512bits bip39 seed key in hex
	if pkstr != "" {
		// set the key back to bytes
		pkb, err := hex.DecodeString(pkstr)
		if err != nil {
			log.Fatalln("Couldn't hex decode the key:", err)
		}

		// PrivKeyFromBytes returns a private and public key for `curve' based on the
		// private key passed as an argument as a byte slice.
		pk, _ = btcec.PrivKeyFromBytes(btcec.S256(), pkb)

		// if we're using seed words or we have specified that the key is a bip39 seed key
		if seedwords != "" || bip39seed {
			// we have to specify a net (use the mainnet (bitcoin) for now)
			net := &chaincfg.MainNetParams

			// NewMaster creates a new master node (bip32 root key) for use in creating a hierarchical
			// deterministic key chain.
			// see https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki#Master_key_generation
			// pkb is the bip39 seed key
			ext, err := hdkeychain.NewMaster(pkb, net)
			if err != nil {
				log.Fatalln("Bip32 root key generation failed", err)
			}

			fmt.Println("bip39 seed words:", wallet.Seedwords)
			fmt.Println("bip39 seed key:", hex.EncodeToString(pk.Serialize()))
			fmt.Println("bip32 root key:", ext.String())

			// create a litecoin account with 20 receive addresses.
			// pass the bip39 seed key because we create a new master in the bip44 function (redundant)
			wallet.Accounts = append(wallet.Accounts, bip44("Bitcoin", addrcount, pkb))
			wallet.Accounts = append(wallet.Accounts, bip44("Litecoin", addrcount, pkb))
			wallet.Accounts = append(wallet.Accounts, bip44("Ether", addrcount, pkb))

			fmt.Println()
			// print out our wallet
			for _, account := range wallet.Accounts {
				fmt.Fprintln(tw, account.Coin)
				fmt.Fprintln(tw, "account extended key:", account.Key.String())
				fmt.Fprintln(tw, "external (m/44'/type'/0'/0) extended key:", account.External.String())
				for i, address := range account.Addresses {
					fmt.Fprintf(tw, "%d\t%s\t%s\t%s\n", i, address.Address, address.Privkey, address.Pubkey)
				}
				fmt.Fprintln(tw)
			}
			tw.Flush()
		}
	} else {
		pk, _ = btcec.NewPrivateKey(btcec.S256())
	}

	// if we specified a normal 256bits key and not a bip39 seed
	if pkstr != "" && !bip39seed && seedwords == "" {
		fmt.Fprintln(tw, "coin\taddress\twif / private key")
		genEth(pk)
		for _, coin := range altcoins {
			if coin.Name != "Ether" {
				genCoin(pk, coin.PubKeyHashAddrID, coin.PrivateKeyID, coin.Name)
			}
		}
		tw.Flush()
		fmt.Println("privkey:", hex.EncodeToString(pk.Serialize()))
	}
}
