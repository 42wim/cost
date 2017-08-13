package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
	"text/tabwriter"
)

var tw *tabwriter.Writer

func genCoin(pk *btcec.PrivateKey, PubKeyHashAddrID, PrivateKeyID byte, name string) {
	net := &chaincfg.MainNetParams
	net.PubKeyHashAddrID = PubKeyHashAddrID
	net.PrivateKeyID = PrivateKeyID
	wif, _ := btcutil.NewWIF(pk, net, false)
	addr, _ := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), net)
	fmt.Fprintf(tw, "%s\t%s\t%s\n", name, addr.EncodeAddress(), wif.String())
}

func genEth(privkey *btcec.PrivateKey) {
	address := crypto.PubkeyToAddress(privkey.PublicKey)
	fmt.Fprintf(tw, "%s\t%s\t%s\n", "Ethereum", address.String(), hex.EncodeToString(crypto.FromECDSA(privkey.ToECDSA())))
}

func main() {
	var pkstr string
	var genpk bool
	var pk *btcec.PrivateKey
	flag.StringVar(&pkstr, "usepk", "", "private key in hex format")
	flag.BoolVar(&genpk, "genpk", false, "generate a private key in hex format and exit")
	flag.Parse()
	if genpk {
		pk, _ = btcec.NewPrivateKey(btcec.S256())
		fmt.Println("privkey:", hex.EncodeToString(pk.Serialize()))
		return
	}
	if pkstr != "" {
		pkb, _ := hex.DecodeString(pkstr)
		pk, _ = btcec.PrivKeyFromBytes(btcec.S256(), pkb)
	} else {
		pk, _ = btcec.NewPrivateKey(btcec.S256())
	}
	tw = tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	//	fmt.Println("pubkey uncompressed:", hex.EncodeToString(pk.PubKey().SerializeUncompressed()))
	//	fmt.Println("pubkey compressed:", hex.EncodeToString(pk.PubKey().SerializeCompressed()))
	fmt.Fprintln(tw, "coin\taddress\twif / private key")
	genEth(pk)
	for _, coin := range altcoins {
		genCoin(pk, coin.PubKeyHashAddrID, coin.PrivateKeyID, coin.Name)
	}
	tw.Flush()
	fmt.Println("privkey:", hex.EncodeToString(pk.Serialize()))
}
