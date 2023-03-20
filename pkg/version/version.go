package version

// The following variables are private fields and should be set when compiling with ldflags, for example --ldflags="-X github.com/crc-org/crc/pkg/version.crcVersion=vX.Y.Z
var (
	version = "unset"
)

func GetVersion() string {
	return version
}
