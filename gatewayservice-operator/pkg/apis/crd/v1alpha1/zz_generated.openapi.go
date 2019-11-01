// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/crd/v1alpha1.GatewayService":       schema_pkg_apis_crd_v1alpha1_GatewayService(ref),
		"./pkg/apis/crd/v1alpha1.GatewayServiceSpec":   schema_pkg_apis_crd_v1alpha1_GatewayServiceSpec(ref),
		"./pkg/apis/crd/v1alpha1.GatewayServiceStatus": schema_pkg_apis_crd_v1alpha1_GatewayServiceStatus(ref),
	}
}

func schema_pkg_apis_crd_v1alpha1_GatewayService(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GatewayService is the Schema for the gatewayservice API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/crd/v1alpha1.GatewayServiceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/crd/v1alpha1.GatewayServiceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/crd/v1alpha1.GatewayServiceSpec", "./pkg/apis/crd/v1alpha1.GatewayServiceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_crd_v1alpha1_GatewayServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GatewayServiceSpec defines the desired state of GatewayService",
				Properties: map[string]spec.Schema{
					"hosts": {
						SchemaProps: spec.SchemaProps{
							Description: "List of Servers > map of list of hosts and port",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"port": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
					"mode": {
						SchemaProps: spec.SchemaProps{
							Description: "Options: SIMPLE|PASSTHROUGH|MUTUAL|ISTIO_MUTUAL|AUTO_PASSTHROUGH",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"protocol": {
						SchemaProps: spec.SchemaProps{
							Description: "Options: HTTP|HTTPS|GRPC|HTTP2|MONGO|TCP|TLS",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"trafficType": {
						SchemaProps: spec.SchemaProps{
							Description: "Options: \"ingress\" or \"egress\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tlsOptions": {
						SchemaProps: spec.SchemaProps{
							Description: "Options: TLSSecret|TLSSecretRef|TLSSecretPath Supports either creating the secret, referencing the secret, or explicitly referencing the mount path in the pod.",
							Ref:         ref("./pkg/apis/crd/v1alpha1.TLSOptions"),
						},
					},
				},
				Required: []string{"hosts", "port", "mode", "protocol", "trafficType"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/crd/v1alpha1.TLSOptions"},
	}
}

func schema_pkg_apis_crd_v1alpha1_GatewayServiceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "GatewayServiceStatus defines the observed state of GatewayService",
				Properties: map[string]spec.Schema{
					"condition": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html",
							Ref:         ref("./pkg/apis/crd/v1alpha1.Condition"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/crd/v1alpha1.Condition"},
	}
}
