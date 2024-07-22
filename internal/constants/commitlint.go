package constants

var COMMITLINT_CONFIG = map[string]string{
	"npm":  "#!/bin/sh\n. \"$(dirname \"$0\")/_/husky.sh\"\nnpx --no-install commitlint --edit \"$1\"\n",
	"yarn": "#!/bin/sh\n. \"$(dirname \"$0\")/_/husky.sh\"\nyarn commitlint --edit \"$1\"\n",
}
