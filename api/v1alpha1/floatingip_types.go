/*
Copyright 2023 Peter Kurfer.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net"
)

// FloatingIPProtocolVersion is the IP protocol version
// +kubebuilder:validation:Enum=v4;v6
type FloatingIPProtocolVersion hcloud.FloatingIPType

const (
	FloatingIPProtocolVersionV4 FloatingIPProtocolVersion = "ipv4"
	FloatingIPProtocolVersionV6 FloatingIPProtocolVersion = "ipv6"
)

// FloatingIPSpec defines the desired state of FloatingIP
type FloatingIPSpec struct {

	// ProtocolVersion is either v4 or v6 and determines the IP protocol version for the floating IP
	// +kubebuilder:default=v4
	ProtocolVersion FloatingIPProtocolVersion `json:"protocolVersion"`

	// Location is the Hetzner cloud location where
	Location string `json:"location,omitempty"`

	// Description that will be added to the floating IP
	Description *string `json:"description,omitempty"`
	// Labels that will be applied to the floating IP
	Labels map[string]string `json:"labels,omitempty"`
}

// FloatingIPStatus defines the observed state of FloatingIP
type FloatingIPStatus struct {
	// +listType=map
	// +listMapKey=type
	// +patchStrategy=merge
	// +patchMergeKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`

	ProviderID int64  `json:"providerID,omitempty"`
	IP         net.IP `json:"ip,omitempty"`

	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:path=floatingIPs,scope=Cluster

// FloatingIP is the Schema for the floatingips API
type FloatingIP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FloatingIPSpec   `json:"spec,omitempty"`
	Status FloatingIPStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// FloatingIPList contains a list of FloatingIP
type FloatingIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FloatingIP `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FloatingIP{}, &FloatingIPList{})
}
