module k8s.io/autoscaler/cluster-autoscaler

go 1.19

require (
	cloud.google.com/go/compute v1.7.0
	github.com/Azure/azure-sdk-for-go v66.0.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.28
	github.com/Azure/go-autorest/autorest/adal v0.9.21
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.11
	github.com/Azure/go-autorest/autorest/date v0.3.0
	github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/Azure/skewer v0.0.16
	github.com/aws/aws-sdk-go v1.44.105
	github.com/digitalocean/godo v1.86.0
	github.com/ghodss/yaml v1.0.0
	github.com/gofrs/uuid v4.3.0+incompatible
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.9
	github.com/google/go-querystring v1.1.0
	github.com/google/uuid v1.3.0
	github.com/jmespath/go-jmespath v0.4.0
	github.com/json-iterator/go v1.1.12
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.13.0
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.8.0
	golang.org/x/crypto v0.0.0-20220924013350-4ba4fb4dd9e7
	golang.org/x/net v0.0.0-20220909164309-bea034e7d591
	golang.org/x/oauth2 v0.0.0-20220909003341-f21342109be1
	google.golang.org/api v0.96.0
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/gcfg.v1 v1.2.3
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.25.2
	k8s.io/apimachinery v0.25.2
	k8s.io/apiserver v0.25.2
	k8s.io/client-go v0.25.2
	k8s.io/cloud-provider v0.25.2
	k8s.io/component-base v0.25.2
	k8s.io/component-helpers v0.25.2
	k8s.io/klog/v2 v2.80.1
	k8s.io/kubelet v0.25.2
	k8s.io/kubernetes v1.25.1
	k8s.io/legacy-cloud-providers v0.25.2
	k8s.io/utils v0.0.0-20220823124924-e9cbc92d1a73
	sigs.k8s.io/cloud-provider-azure v1.25.0
)

require (
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.6 // indirect
	github.com/Azure/go-autorest/autorest/mocks v0.4.2 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/GoogleCloudPlatform/k8s-cloud-provider v1.20.0 // indirect
	github.com/JeffAshton/win_pdh v0.0.0-20161109143554-76bb4ee9f0ab // indirect
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/Microsoft/hcsshim v0.9.4 // indirect
	github.com/NYTimes/gziphandler v1.1.1 // indirect
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/checkpoint-restore/go-criu/v5 v5.3.0 // indirect
	github.com/cilium/ebpf v0.9.1 // indirect
	github.com/container-storage-interface/spec v1.6.0 // indirect
	github.com/containerd/cgroups v1.0.4 // indirect
	github.com/containerd/console v1.0.3 // indirect
	github.com/containerd/ttrpc v1.1.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.4.0 // indirect
	github.com/cyphar/filepath-securejoin v0.2.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/docker/distribution v2.8.1+incompatible // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/emicklei/go-restful/v3 v3.9.0 // indirect
	github.com/euank/go-kmsg-parser v2.0.0+incompatible // indirect
	github.com/evanphx/json-patch v5.6.0+incompatible // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/form3tech-oss/jwt-go v3.2.5+incompatible // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible // indirect
	github.com/godbus/dbus/v5 v5.1.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/cadvisor v0.45.0 // indirect
	github.com/google/gnostic v0.6.9 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.1.0 // indirect
	github.com/googleapis/gax-go/v2 v2.5.1 // indirect
	github.com/gophercloud/gophercloud v1.0.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/heketi/heketi v10.3.0+incompatible // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/karrick/godirwalk v1.17.0 // indirect
	github.com/libopenstorage/openstorage v1.0.0 // indirect
	github.com/lithammer/dedent v1.1.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mindprince/gonvml v0.0.0-20211002210717-ac0b66419a41 // indirect
	github.com/mistifyio/go-zfs v2.1.2-0.20190413222219-f784269be439+incompatible // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/moby/ipvs v1.0.2 // indirect
	github.com/moby/spdystream v0.2.0 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/mrunalp/fileutils v0.5.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/runc v1.1.4 // indirect
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417 // indirect
	github.com/opencontainers/selinux v1.10.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/rubiojr/go-vhd v0.0.0-20200706105327-02e210299021 // indirect
	github.com/seccomp/libseccomp-golang v0.10.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/cobra v1.5.0 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635 // indirect
	github.com/vishvananda/netlink v1.1.1-0.20201029203352-d40f9887b852 // indirect
	github.com/vishvananda/netns v0.0.0-20220913150850-18c4f4234207 // indirect
	github.com/vmware/govmomi v0.29.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.5 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.5 // indirect
	go.etcd.io/etcd/client/v3 v3.5.5 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.opentelemetry.io/contrib v0.20.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful v0.20.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.20.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.20.0 // indirect
	go.opentelemetry.io/otel v0.20.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp v0.20.0 // indirect
	go.opentelemetry.io/otel/metric v0.20.0 // indirect
	go.opentelemetry.io/otel/sdk v0.20.0 // indirect
	go.opentelemetry.io/otel/sdk/export/metric v0.20.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v0.20.0 // indirect
	go.opentelemetry.io/otel/trace v0.20.0 // indirect
	go.opentelemetry.io/proto/otlp v0.7.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	golang.org/x/sync v0.0.0-20220907140024-f12130a52804 // indirect
	golang.org/x/sys v0.0.0-20220915200043-7b5979e65e41 // indirect
	golang.org/x/term v0.0.0-20220722155259-a9ba230a4035 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220722155302-e5dcc9cfc0b9 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220916172020-2692e8806bfa // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/cri-api v0.25.2 // indirect
	k8s.io/csi-translation-lib v0.25.2 // indirect
	k8s.io/kube-openapi v0.0.0-20220803164354-a70c9af30aea // indirect
	k8s.io/kube-proxy v0.25.2 // indirect
	k8s.io/kube-scheduler v0.25.2 // indirect
	k8s.io/kubectl v0.25.2 // indirect
	k8s.io/mount-utils v0.25.2 // indirect
	sigs.k8s.io/apiserver-network-proxy/konnectivity-client v0.0.32 // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace (
	github.com/aws/aws-sdk-go/service/eks => github.com/aws/aws-sdk-go/service/eks v1.38.49

	github.com/digitalocean/godo => github.com/digitalocean/godo v1.27.0
	github.com/libopenstorage/openstorage => github.com/libopenstorage/openstorage v1.0.0

	github.com/rancher/go-rancher => github.com/rancher/go-rancher v0.1.0
	go.opentelemetry.io/contrib => go.opentelemetry.io/contrib v0.20.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful => go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful v0.20.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc => go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.20.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp => go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.20.0 // indirect
	go.opentelemetry.io/otel => go.opentelemetry.io/otel v0.20.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp => go.opentelemetry.io/otel/exporters/otlp v0.20.0 // indirect
	go.opentelemetry.io/otel/metric => go.opentelemetry.io/otel/metric v0.20.0 // indirect
	go.opentelemetry.io/otel/sdk => go.opentelemetry.io/otel/sdk v0.20.0 // indirect
	go.opentelemetry.io/otel/sdk/export/metric => go.opentelemetry.io/otel/sdk/export/metric v0.20.0 // indirect
	go.opentelemetry.io/otel/sdk/metric => go.opentelemetry.io/otel/sdk/metric v0.20.0 // indirect
	go.opentelemetry.io/otel/trace => go.opentelemetry.io/otel/trace v0.20.0 // indirect
	go.opentelemetry.io/proto/otlp => go.opentelemetry.io/proto/otlp v0.7.0 // indirect

	k8s.io/api => k8s.io/api v0.25.2

	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.25.2

	k8s.io/apimachinery => k8s.io/apimachinery v0.25.2

	k8s.io/apiserver => k8s.io/apiserver v0.25.2

	k8s.io/cli-runtime => k8s.io/cli-runtime v0.25.2

	k8s.io/client-go => k8s.io/client-go v0.25.2

	k8s.io/cloud-provider => k8s.io/cloud-provider v0.25.2

	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.25.2

	k8s.io/code-generator => k8s.io/code-generator v0.25.2

	k8s.io/component-base => k8s.io/component-base v0.25.2

	k8s.io/component-helpers => k8s.io/component-helpers v0.25.2

	k8s.io/controller-manager => k8s.io/controller-manager v0.25.2

	k8s.io/cri-api => k8s.io/cri-api v0.25.2

	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.25.2

	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.25.2

	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.25.2

	k8s.io/kube-proxy => k8s.io/kube-proxy v0.25.2

	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.25.2

	k8s.io/kubectl => k8s.io/kubectl v0.25.2

	k8s.io/kubelet => k8s.io/kubelet v0.25.2

	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.25.2

	k8s.io/metrics => k8s.io/metrics v0.25.2

	k8s.io/mount-utils => k8s.io/mount-utils v0.25.2

	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.25.2

	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.25.2

	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.25.2

	k8s.io/sample-controller => k8s.io/sample-controller v0.25.2
)
