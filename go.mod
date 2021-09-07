module github.com/erda-project/erda-proto-go

go 1.14

require (
	github.com/erda-project/erda-infra v0.0.0-20210901113044-1130cdf38f06
	github.com/golang/protobuf v1.5.2
	github.com/mwitkow/go-proto-validators v0.3.2
	google.golang.org/genproto v0.0.0-20210820002220-43fce44e7af1
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
	k8s.io/api => github.com/kubernetes/api v0.18.3
	k8s.io/apiextensions-apiserver => github.com/kubernetes/apiextensions-apiserver v0.18.3
	k8s.io/apimachinery => github.com/kubernetes/apimachinery v0.18.3
	k8s.io/apiserver => github.com/kubernetes/apiserver v0.18.3
	k8s.io/client-go => github.com/kubernetes/client-go v0.18.3
	k8s.io/component-base => github.com/kubernetes/component-base v0.18.3
	k8s.io/klog => github.com/kubernetes/klog v1.0.0
	k8s.io/kube-scheduler => github.com/kubernetes/kube-scheduler v0.18.3
	k8s.io/kubectl => github.com/kubernetes/kubectl v0.18.3
	k8s.io/kubernetes => github.com/kubernetes/kubernetes v1.13.5
)
