// This file was auto-generated from YAML files

package v1_0

func init() {
	Profile["/tosca/simple-for-nfv/1.0/relationships.yaml"] = `
# Modified from a file that was distributed with this NOTICE:
#
#   Apache AriaTosca
#   Copyright 2016-2017 The Apache Software Foundation
#
#   This product includes software developed at
#   The Apache Software Foundation (http://www.apache.org/).

tosca_definitions_version: tosca_simple_yaml_1_1

imports:
- capabilities.yaml

relationship_types:

  tosca.relationships.nfv.VirtualBindsTo:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-NFV-v1.0]'
      citation_location: 5.7.1
    description: >-
      This relationship type represents an association relationship between VDU and CP node types.
    derived_from: tosca.relationships.DependsOn
    valid_target_types: [ tosca.capabilities.nfv.VirtualBindable ]

  # ERRATUM: csd04 lacks the definition of tosca.relationships.nfv.Monitor (the derived_from and
  # valid_target_types), so we are using the definition in csd03 section 8.4.2.
  tosca.relationships.nfv.Monitor:
    metadata:
      normative: 'true'
      citation: '[TOSCA-Simple-Profile-NFV-v1.0]'
      citation_location: 5.7.2
    description: >-
      This relationship type represents an association relationship to the Metric capability of VDU
      node types.
    derived_from: tosca.relationships.ConnectsTo
    valid_target_types: [ tosca.capabilities.nfv.Metric ]
`
}
