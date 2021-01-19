package dbtest

import (
	"os"
	"path/filepath"
	"testing"

	fixtures "github.com/aquasecurity/bolt-fixtures"
	"github.com/aquasecurity/trivy-db/pkg/db"
	"github.com/stretchr/testify/require"
)

// InitDB loads fixtures and initializes the database
func InitDB(t *testing.T, fixtureFiles []string) {
	// Create a temp dir
	dir := t.TempDir()

	dbPath := db.Path(dir)
	dbDir := filepath.Dir(dbPath)
	err := os.MkdirAll(dbDir, 0700)
	require.NoError(t, err)

	// Load testdata into BoltDB
	loader, err := fixtures.New(dbPath, fixtureFiles)
	require.NoError(t, err)
	require.NoError(t, loader.Load())
	require.NoError(t, loader.Close())

	require.NoError(t, db.Init(dir))
}
