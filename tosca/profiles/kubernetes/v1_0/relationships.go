// This file was auto-generated from YAML files

package v1_0

func init() {
	Profile["/tosca/kubernetes/1.0/relationships.yaml"] = `
tosca_definitions_version: tosca_simple_yaml_1_1

imports:
- capabilities.yaml

relationship_types:

  kubernetes.Route:
    derived_from: tosca.relationships.Root
    valid_target_types: [ kubernetes.Service ] # capability
`
}
