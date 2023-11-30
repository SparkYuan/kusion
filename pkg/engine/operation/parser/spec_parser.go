package parser

import (
	"fmt"

	"kusionstack.io/kusion/pkg/apis/intent"
	"kusionstack.io/kusion/pkg/apis/status"
	"kusionstack.io/kusion/pkg/engine/operation/graph"
	opsmodels "kusionstack.io/kusion/pkg/engine/operation/models"
	"kusionstack.io/kusion/pkg/util"
	"kusionstack.io/kusion/pkg/util/json"
	"kusionstack.io/kusion/third_party/terraform/dag"
)

type IntentParser struct {
	intent *intent.Intent
}

func NewIntentParser(i *intent.Intent) *IntentParser {
	return &IntentParser{intent: i}
}

var _ Parser = (*IntentParser)(nil)

func (m *IntentParser) Parse(g *dag.AcyclicGraph) (s status.Status) {
	util.CheckNotNil(g, "dag is nil")
	i := m.intent
	util.CheckNotNil(i, "models is nil")
	if i.Resources == nil {
		sprintf := fmt.Sprintf("no resources in models:%s", json.Marshal2String(i))
		return status.NewBaseStatus(status.Warning, status.NotFound, sprintf)
	}

	root, err := g.Root()
	util.CheckNotError(err, "get dag root error")
	util.CheckNotNil(root, fmt.Sprintf("No root in this DAG:%s", json.Marshal2String(g)))
	resourceIndex := i.Resources.Index()
	for key, resource := range resourceIndex {
		rn, s := graph.NewResourceNode(key, resourceIndex[key], opsmodels.Update)
		if status.IsErr(s) {
			return s
		}

		// add this resource to dag at first time
		if !g.HasVertex(rn) {
			g.Add(rn)
			g.Connect(dag.BasicEdge(root, rn))
		} else {
			// always get the latest vertex in this g otherwise you will get subtle mistake in walking this g
			rn = GetVertex(g, rn).(*graph.ResourceNode)
			g.Connect(dag.BasicEdge(root, rn))
		}

		// compute implicit and explicate dependencies
		refNodeKeys, s := updateDependencies(resource)
		if status.IsErr(s) {
			return s
		}

		// linkRefNodes
		s = LinkRefNodes(g, refNodeKeys, resourceIndex, rn, opsmodels.Update, nil)
		if status.IsErr(s) {
			return s
		}
	}

	if err = g.Validate(); err != nil {
		return status.NewErrorStatusWithMsg(status.IllegalManifest, "Found circle dependency in models:"+err.Error())
	}
	g.TransitiveReduction()
	return s
}
