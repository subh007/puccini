package v1_1

import (
	"github.com/tliron/puccini/tosca"
	"github.com/tliron/puccini/tosca/normal"
)

//
// WorkflowDefinition
//
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.7.7
//

type WorkflowDefinition struct {
	*Entity `name:"workflow definition"`
	Name    string `namespace:""`

	Metadata                Metadata                          `read:"metadata,Metadata"`
	Description             *string                           `read:"description"`
	InputDefinitions        PropertyDefinitions               `read:"inputs,PropertyDefinition"`
	PreconditionDefinitions []*WorkflowPreconditionDefinition `read:"preconditions,WorkflowPreconditionDefinition"`
	StepDefinitions         WorkflowStepDefinitions           `read:"steps,WorkflowStepDefinition"`
}

func NewWorkflowDefinition(context *tosca.Context) *WorkflowDefinition {
	return &WorkflowDefinition{
		Entity:           NewEntity(context),
		Name:             context.Name,
		InputDefinitions: make(PropertyDefinitions),
		StepDefinitions:  make(WorkflowStepDefinitions),
	}
}

// tosca.Reader signature
func ReadWorkflowDefinition(context *tosca.Context) interface{} {
	self := NewWorkflowDefinition(context)
	context.ValidateUnsupportedFields(context.ReadFields(self, Readers))
	return self
}

// tosca.Mappable interface
func (self *WorkflowDefinition) GetKey() string {
	return self.Name
}

// tosca.Renderable interface
func (self *WorkflowDefinition) Render() {
	log.Infof("{render} workflow definition: %s", self.Name)

	self.StepDefinitions.Render()
}

func (self *WorkflowDefinition) Normalize(s *normal.ServiceTemplate) *normal.Workflow {
	log.Infof("{normalize} workflow: %s", self.Name)

	w := s.NewWorkflow(self.Name)

	if self.Description != nil {
		w.Description = *self.Description
	}

	// TODO: support property definitions
	//self.InputDefinitions.Normalize(w.Inputs)

	sts := make(normal.WorkflowSteps)
	for name, step := range self.StepDefinitions {
		sts[name] = step.Normalize(w, s)
	}

	for name, step := range self.StepDefinitions {
		step.NormalizeNext(sts[name], w)
	}

	return w
}

//
// WorkflowDefinitions
//

type WorkflowDefinitions map[string]*WorkflowDefinition
