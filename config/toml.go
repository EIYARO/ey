package config

import (
	"path"

	cmn "github.com/tendermint/tmlibs/common"
)

/****** these are for production settings ***********/
func EnsureRoot(rootDir string, network string) {
	cmn.EnsureDir(rootDir, 0700)
	cmn.EnsureDir(rootDir+"/data", 0700)

	configFilePath := path.Join(rootDir, "config.toml")

	// Write default config file if missing.
	if !cmn.FileExists(configFilePath) {
		cmn.MustWriteFile(configFilePath, []byte(selectNetwork(network)), 0644)
	}
}

var defaultConfigTmpl = `# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml
fast_sync = true
db_backend = "leveldb"
api_addr = "0.0.0.0:9888"
node_alias = ""
`

var mainNetConfigTmpl = `chain_id = "mainnet"
[p2p]
laddr = "tcp://0.0.0.0:46657"
seeds = "154.44.8.233:46657,154.44.8.32:46657,154.44.8.62:46657,154.40.59.31:46657,27.25.156.254:46657,24.233.3.133:46657,103.115.46.201:46657,217.194.133.61:46657"
keep_dial = "154.44.8.233:46657,154.44.8.32:46657,154.44.8.62:46657,154.40.59.31:46657,27.25.156.254:46657,24.233.3.133:46657,103.115.46.201:46657,217.194.133.61:46657"
max_num_peers = 100
`

var testNetConfigTmpl = `chain_id = "wisdom"
[p2p]
laddr = "tcp://0.0.0.0:46656"
seeds = "154.3.0.182:46656,154.3.0.179:46656,38.46.31.72:46656"
keep_dial = "154.3.0.182:46656,154.3.0.179:46656,38.46.31.72:46656"
max_num_peers = 10
`

var soloNetConfigTmpl = `chain_id = "solonet"
[p2p]
laddr = "tcp://0.0.0.0:46658"
seeds = ""
`

// Select network seeds to merge a new string.
func selectNetwork(network string) string {
	switch network {
	case "mainnet":
		return defaultConfigTmpl + mainNetConfigTmpl
	case "testnet":
		return defaultConfigTmpl + testNetConfigTmpl
	default:
		return defaultConfigTmpl + soloNetConfigTmpl
	}
}
