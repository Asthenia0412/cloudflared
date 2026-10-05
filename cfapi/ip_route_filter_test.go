package cfapi

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

func newFilterCLIContext(t *testing.T, args ...string) *cli.Context {
	t.Helper()
	set := flag.NewFlagSet("test", flag.ContinueOnError)
	set.String(filterSubsetIpRoute.Name, "", "")
	set.String(filterSupersetIpRoute.Name, "", "")
	require.NoError(t, set.Parse(args))
	return cli.NewContext(nil, set, nil)
}

func TestIpRouteFilterFromCLI(t *testing.T) {
	filter, err := NewIpRouteFilterFromCLI(newFilterCLIContext(t,
		"--filter-network-is-subset-of", "10.0.0.0/8"))
	require.NoError(t, err)
	require.Equal(t, "is_deleted=false&network_subset=10.0.0.0%2F8&tun_types=cfd_tunnel", filter.Encode())

	filter, err = NewIpRouteFilterFromCLI(newFilterCLIContext(t,
		"--filter-network-is-superset-of", "10.0.0.0/24"))
	require.NoError(t, err)
	require.Equal(t, "is_deleted=false&network_superset=10.0.0.0%2F24&tun_types=cfd_tunnel", filter.Encode())
}
