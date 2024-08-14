package store

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Ethernal-Tech/ethgo"
	"github.com/Ethernal-Tech/ethgo/abi"
	"github.com/stretchr/testify/require"
)

var StateSyncEventABI = abi.MustNewEvent("event StateSynced(uint256 indexed id, " +
	"address indexed sender, address indexed receiver, bytes data)")

func CreateTestLogForStateSyncEvent(t *testing.T, blockNumber, logIndex uint64) *ethgo.Log {
	t.Helper()

	topics := make([]ethgo.Hash, 3)
	topics[0] = StateSyncEventABI.ID()
	topics[1] = ethgo.BytesToHash(ethgo.ZeroAddress.Bytes())
	topics[2] = ethgo.BytesToHash(ethgo.ZeroAddress.Bytes())
	encodedData, err := abi.MustNewType("tuple(string a)").Encode([]string{"data"})
	require.NoError(t, err)

	return &ethgo.Log{
		BlockNumber: blockNumber,
		LogIndex:    logIndex,
		Address:     ethgo.ZeroAddress,
		Topics:      topics,
		Data:        encodedData,
	}
}

// NewTestTrackerStore creates new instance of state used by tests.
func NewTestTrackerStore(tb testing.TB) *BBoltDBEventTrackerStore {
	tb.Helper()

	dir, err := os.MkdirTemp("", "even-tracker-temp")
	if err != nil {
		tb.Fatal(err)
	}

	tb.Cleanup(func() {
		if err := os.RemoveAll(dir); err != nil {
			tb.Fatal(err)
		}
	})

	store, err := NewBoltDBEventTrackerStore(filepath.Join(dir, "tracker.db"))
	if err != nil {
		tb.Fatal(err)
	}

	return store
}
