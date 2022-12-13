module platgit.mihoyo.com/easyai/easyai-core

go 1.19

// TODO: we need to update k8s.io and volcano API when  `kubeflow/training-operator` is updated, And
// this will bring a lot of changes once the k8s library version updated
// for now. as we know,  kubeflow/training-operator v1.4.0 released at 2022.03.1:
// 	1. works with k8s libaray v0.19.9 and volcano v1.2.0-k8s1.19.6
//  2. works with k8s cluster < 1.22
require (
	k8s.io/api v0.19.9
	k8s.io/apimachinery v0.19.9 // indirect
)

// for web backend
require (
	github.com/BurntSushi/toml v1.2.0
	github.com/Masterminds/sprig/v3 v3.2.2
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/gin-gonic/gin v1.8.1
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-resty/resty/v2 v2.7.0
	github.com/google/uuid v1.3.0
	github.com/jinzhu/copier v0.3.5
	github.com/json-iterator/go v1.1.12
	github.com/pkg/errors v0.9.1
	github.com/sony/sonyflake v1.0.0
	github.com/spf13/cast v1.5.0
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-logr/logr v0.3.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/huandu/xstrings v1.3.1 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/onsi/ginkgo v1.14.1 // indirect
	github.com/onsi/gomega v1.10.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	golang.org/x/net v0.0.0-20220520000938-2e3eb7b945c2 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/xerrors v0.0.0-20220517211312-f3a8303e98df // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/klog/v2 v2.2.0 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.0.1 // indirect
)

//replace (
//	github.com/golang/protobuf v1.5.2 => github.com/golang/protobuf v1.4.3
//	google.golang.org/grpc v1.42.0 => google.golang.org/grpc v1.40.0
//	google.golang.org/protobuf v1.27.1 => google.golang.org/protobuf v1.25.0
//)

replace (
	github.com/xen0n/go-workwx => github.com/shuaiyy/go-workwx v0.0.7
	k8s.io/api => k8s.io/api v0.19.9
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.9
	k8s.io/client-go => k8s.io/client-go v0.19.9
	k8s.io/code-generator => k8s.io/code-generator v0.19.9
)
