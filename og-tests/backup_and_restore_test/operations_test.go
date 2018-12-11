package backupandrestore_test

import (
	"testing"

	"og/helpers"
)

const testDirectory = "operations/backup-and-restore"

var backupAndRestoreTests = map[string]helpers.OpsFileTestParams{
	"enable-backup-restore-azure.yml": {
		Ops:       []string{"enable-backup-restore.yml", "../use-external-blobstore.yml", "../use-azure-storage-blobstore.yml", "enable-backup-restore-azure.yml"},
		VarsFiles: []string{"../example-vars-files/vars-use-azure-storage-blobstore.yml"},
	},
	"enable-backup-restore-nfs-broker.yml": {
		Ops:  []string{"enable-backup-restore.yml", "enable-backup-restore-nfs-broker.yml"},
		Vars: []string{"nfs-broker-database-password=i_am_a_password"},
	},
	"enable-backup-restore-s3-unversioned.yml": {
		Ops:       []string{"enable-backup-restore.yml", "../use-external-blobstore.yml", "../use-s3-blobstore.yml", "enable-backup-restore-s3-unversioned.yml"},
		VarsFiles: []string{"../example-vars-files/vars-use-s3-blobstore.yml", "example-vars-files/vars-enable-backup-restore-s3-unversioned.yml"},
	},
	"enable-backup-restore-s3-versioned.yml": {
		Ops:       []string{"enable-backup-restore.yml", "../use-external-blobstore.yml", "../use-s3-blobstore.yml", "enable-backup-restore-s3-versioned.yml"},
		VarsFiles: []string{"../example-vars-files/vars-use-s3-blobstore.yml"},
	},
	"enable-backup-restore-smb-broker.yml": {
		Ops:  []string{"enable-backup-restore.yml", "enable-backup-restore-smb-broker.yml"},
		Vars: []string{"azurefile-broker-database-password=i_am_a_password"},
	},
	"enable-backup-restore.yml": {},
	"enable-restore-azure-clone.yml": {
		Ops:       []string{"enable-backup-restore.yml", "../use-external-blobstore.yml", "../use-azure-storage-blobstore.yml", "enable-restore-azure-clone.yml"},
		VarsFiles: []string{"../example-vars-files/vars-use-azure-storage-blobstore.yml", "example-vars-files/vars-enable-restore-azure-clone.yml"},
	},
}

func TestBackupAndRestore(t *testing.T) {
	cfDeploymentHome, err := helpers.SetPath()
	if err != nil {
		t.Fatalf("setup: %v", err)
	}

	suite := helpers.NewSuiteTest(cfDeploymentHome, testDirectory, backupAndRestoreTests)
	suite.EnsureTestCoverage(t)
	suite.ReadmeTest(t)
	suite.InterpolateTest(t)
}
