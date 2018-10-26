package main

import (
	"chainspace.io/prototype/internal/log"

	"github.com/tav/golly/optparse"
)

const logo = `          [0;1;34;94m█[0m               [0;1;31;91m▀[0m
    [0;1;32;92m▄▄[0;1;36;96m▄[0m   [0;1;35;95m█[0m [0;1;31;91m▄▄[0m    [0;1;36;96m▄▄[0;1;34;94m▄[0m   [0;1;31;91m▄▄[0;1;33;93m▄[0m    [0;1;36;96m▄[0m [0;1;34;94m▄[0;1;35;95m▄[0m    [0;1;33;93m▄[0;1;32;92m▄▄[0m   [0;1;34;94m▄[0;1;35;95m▄▄[0;1;31;91m▄[0m    [0;1;32;92m▄[0;1;36;96m▄▄[0m    [0;1;31;91m▄▄[0;1;33;93m▄[0m    [0;1;36;96m▄[0;1;34;94m▄▄[0m
   [0;1;32;92m█[0;1;36;96m▀[0m  [0;1;34;94m▀[0m  [0;1;31;91m█▀[0m  [0;1;32;92m█[0m  [0;1;36;96m▀[0m   [0;1;35;95m█[0m    [0;1;32;92m█[0m    [0;1;34;94m█[0;1;35;95m▀[0m  [0;1;31;91m█[0m  [0;1;32;92m█[0m   [0;1;34;94m▀[0m  [0;1;35;95m█[0;1;31;91m▀[0m [0;1;33;93m▀█[0m  [0;1;36;96m▀[0m   [0;1;35;95m█[0m  [0;1;31;91m█[0;1;33;93m▀[0m  [0;1;32;92m▀[0m  [0;1;34;94m█▀[0m  [0;1;31;91m█[0m
   [0;1;36;96m█[0m      [0;1;33;93m█[0m   [0;1;36;96m█[0m  [0;1;34;94m▄[0;1;35;95m▀▀[0;1;31;91m▀█[0m    [0;1;36;96m█[0m    [0;1;35;95m█[0m   [0;1;33;93m█[0m   [0;1;36;96m▀[0;1;34;94m▀▀[0;1;35;95m▄[0m  [0;1;31;91m█[0m   [0;1;32;92m█[0m  [0;1;34;94m▄▀[0;1;35;95m▀▀[0;1;31;91m█[0m  [0;1;33;93m█[0m      [0;1;35;95m█▀[0;1;31;91m▀▀[0;1;33;93m▀[0m
   [0;1;34;94m▀[0;1;35;95m█▄[0;1;31;91m▄▀[0m  [0;1;32;92m█[0m   [0;1;34;94m█[0m  [0;1;35;95m▀[0;1;31;91m▄▄[0;1;33;93m▀█[0m  [0;1;36;96m▄▄[0;1;34;94m█▄[0;1;35;95m▄[0m  [0;1;31;91m█[0m   [0;1;32;92m█[0m  [0;1;34;94m▀▄[0;1;35;95m▄▄[0;1;31;91m▀[0m  [0;1;33;93m█[0;1;32;92m█▄[0;1;36;96m█▀[0m  [0;1;35;95m▀▄[0;1;31;91m▄▀[0;1;33;93m█[0m  [0;1;32;92m▀[0;1;36;96m█▄[0;1;34;94m▄▀[0m  [0;1;31;91m▀█[0;1;33;93m▄▄[0;1;32;92m▀[0m
                                             [0;1;32;92m█[0m
                                             [0;1;36;96m▀[0m`

func main() {
	log.ToConsole(log.DebugLevel)
	cmds := map[string]func([]string, string){
		"genkeys":   cmdGenKeys,
		"genload":   cmdGenLoad,
		"init":      cmdInit,
		"interpret": cmdInterpret,
		"run":       cmdRun,
		"sbac":      cmdSBAC,
		"contracts": cmdContracts,
	}
	info := map[string]string{
		"genkeys":   "Generate new keys for a node",
		"genload":   "Run a node with transaction load",
		"init":      "Initialise a new Chainspace network",
		"interpret": "Interpret a node's block graph",
		"run":       "Run a node in a Chainspace network",
		"sbac":      "Send transactions in a chainspace network",
		"contracts": "Manage chainspace contracts",
	}
	optparse.Commands("chainspace", "0.0.1", cmds, info, logo)
}
