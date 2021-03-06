// This file was auto-generated from YAML files

package v1_1

func init() {
	Profile["/tosca/simple/1.1/js/get_nodes_of_type.js"] = `

// [TOSCA-Simple-Profile-YAML-v1.1] @ 4.7.1

clout.exec('tosca.helpers');

function evaluate(typeName) {
	if (arguments.length !== 1)
		throw 'must have 1 argument';
	names = [];
	for (var name in clout.vertexes) {
		vertex = clout.vertexes[name];
		if (tosca.isTosca(vertex))
			names.push(vertex.properties.name);
	}
	return names;
}
`
}
