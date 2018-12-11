package experimental_test

import (
	"testing"

	"og/helpers"
)

const testDirectory = "operations/experimental"

var experimentalTests = map[string]helpers.OpsFileTestParams{
	"add-cflinuxfs3.yml": {},
	"add-credhub-lb.yml": {},
	"add-deployment-updater.yml": {},
	"add-deployment-updater-postgres.yml": {
		Ops: []string{"add-deployment-updater.yml", "add-deployment-updater-postgres.yml"},
	},
	"bits-service-local.yml": {
		Ops: []string{"bits-service.yml", "bits-service-local.yml"},
	},
	"bits-service-s3.yml": {
		Ops:       []string{"../use-external-blobstore.yml", "../use-s3-blobstore.yml", "bits-service.yml", "bits-service-s3.yml"},
		VarsFiles: []string{"../example-vars-files/vars-use-s3-blobstore.yml"},
	},
	"bits-service-webdav.yml": {
		Ops: []string{"bits-service.yml", "bits-service-webdav.yml"},
	},
	"bits-service.yml": {},
	"disable-consul-bosh-lite.yml": {
		Ops: []string{"disable-consul.yml", "../bosh-lite.yml", "disable-consul-bosh-lite.yml"},
	},
	"disable-consul-windows.yml": {
		Ops: []string{"../windows2012R2-cell.yml", "disable-consul.yml", "disable-consul-windows.yml"},
	},
	"disable-consul-windows1803.yml": {
		Ops: []string{"windows1803-cell.yml", "disable-consul.yml", "disable-consul-windows1803.yml"},
	},
	"disable-consul-windows2016.yml": {
		Ops: []string{"../windows2016-cell.yml", "disable-consul.yml", "disable-consul-windows2016.yml"},
	},
	"disable-consul.yml":                       {},
	"disable-interpolate-service-bindings.yml": {},
	"enable-bits-service-consul.yml": {
		Ops: []string{"bits-service.yml", "bits-service-local.yml", "enable-bits-service-consul.yml"},
	},
	"enable-bpm-garden.yml":                    {},
	"enable-iptables-logger.yml":               {},
	"enable-mysql-tls.yml":                     {},
	"enable-nfs-volume-service-credhub.yml":    {},
	"enable-oci-phase-1.yml":                   {},
	"enable-routing-integrity.yml":             {},
	"enable-smb-volume-service.yml":            {},
	"enable-suspect-actual-lrp-generation.yml": {},
	"enable-tls-cloud-controller-postgres.yml": {
		Ops: []string{"../use-postgres.yml", "enable-tls-cloud-controller-postgres.yml"},
	},
	"enable-traffic-to-internal-networks.yml":  {},
	"fast-deploy-with-downtime-and-danger.yml": {},
	"infrastructure-metrics.yml":               {},
	"migrate-cf-mysql-to-pxc.yml":              {},
	"migrate-nfsbroker-mysql-to-credhub.yml": {
		Ops:       []string{"../enable-nfs-volume-service.yml", "migrate-nfsbroker-mysql-to-credhub.yml"},
		VarsFiles: []string{"../example-vars-files/vars-migrate-nfsbroker-mysql-to-credhub.yml"},
	},
	"perm-service-with-pxc-release.yml": {
		Ops:  []string{"perm-service.yml", "use-pxc.yml", "perm-service-with-pxc-release.yml"},
		Vars: []string{"perm_uaa_clients_cc_perm_secret=perm_secret", "perm_uaa_clients_perm_monitor_secret=perm_monitor_secret"},
	},
	"perm-service.yml": {
		Ops:  []string{"enable-mysql-tls.yml", "perm-service.yml"},
		Vars: []string{"perm_uaa_clients_cc_perm_secret=perm_secret", "perm_uaa_clients_perm_monitor_secret=perm_monitor_secret"},
	},
	"rootless-containers.yml": {},
	"set-cpu-weight.yml":      {},
	"use-compiled-releases-windows.yml": {
		Ops: []string{"../use-compiled-releases.yml", "../windows-cell.yml", "use-compiled-releases-windows.yml"},
	},
	"use-compiled-releases-xenial-stemcell.yml": {
		Ops: []string{"use-xenial-stemcell.yml", "use-compiled-releases-xenial-stemcell.yml"},
	},
	"use-garden-containerd.yml":                       {},
	"use-logcache-for-cloud-controller-app-stats.yml": {},
	"use-pxc-for-smb-volume-service.yml": {
		Ops: []string{"enable-smb-volume-service.yml", "use-pxc.yml", "use-pxc-for-smb-volume-service.yml"},
	},
	"use-pxc.yml":             {},
	"use-xenial-stemcell.yml": {},
	"windows-component-syslog-ca.yml": {
		Ops:       []string{"windows-enable-component-syslog.yml", "windows-component-syslog-ca.yml"},
		VarsFiles: []string{"../addons/example-vars-files/vars-enable-component-syslog.yml"},
	},
	"windows-enable-component-syslog.yml": {
		Ops:       []string{"windows-enable-component-syslog.yml"},
		VarsFiles: []string{"../addons/example-vars-files/vars-enable-component-syslog.yml"},
	},
	"windows1803-cell.yml": {},
}

func TestExperimental(t *testing.T) {
	cfDeploymentHome, err := helpers.SetPath()
	if err != nil {
		t.Fatalf("setup: %v", err)
	}

	suite := helpers.NewSuiteTest(cfDeploymentHome, testDirectory, experimentalTests)
	suite.EnsureTestCoverage(t)
	suite.ReadmeTest(t)
	suite.InterpolateTest(t)
}
