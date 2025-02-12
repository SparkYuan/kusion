package models

import (
	"reflect"
	"testing"

	"kusionstack.io/kusion/pkg/engine/operation/types"

	"kusionstack.io/kusion/pkg/engine/models"

	"kusionstack.io/kusion/pkg/projectstack"
	"kusionstack.io/kusion/pkg/util/pretty"
)

var (
	TestChangeStepOpCreate   = NewChangeStep("id", types.Create, nil, nil, nil)
	TestChangeStepOpDelete   = NewChangeStep("id", types.Delete, nil, nil, nil)
	TestChangeStepOpUpdate   = NewChangeStep("id", types.Update, nil, nil, nil)
	TestChangeStepOpUnChange = NewChangeStep("id", types.UnChange, nil, nil, nil)
	TestStepKeys             = []string{"test-key-1", "test-key-2", "test-key-3", "test-key-4"}
	TestChangeSteps          = map[string]*ChangeStep{
		"test-key-1": TestChangeStepOpCreate,
		"test-key-2": TestChangeStepOpDelete,
		"test-key-3": TestChangeStepOpUpdate,
		"test-key-4": TestChangeStepOpUnChange,
	}
)

func TestOpType_Ing(t *testing.T) {
	tests := []struct {
		name string
		op   types.ActionType
		want string
	}{
		{
			name: "t1",
			op:   types.Create,
			want: "Creating",
		},
		{
			name: "t2",
			op:   types.Delete,
			want: "Deleting",
		},
		{
			name: "t3",
			op:   types.Update,
			want: "Updating",
		},
		{
			name: "t4",
			op:   types.UnChange,
			want: "Unchanged",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.op.Ing(); got != tt.want {
				t.Errorf("ActionType.Ing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpType_PrettyString(t *testing.T) {
	tests := []struct {
		name string
		op   types.ActionType
		want string
	}{
		{
			name: "t1",
			op:   types.Create,
			want: pretty.Green(types.Create.Ing()),
		},
		{
			name: "t2",
			op:   types.Delete,
			want: pretty.Red(types.Delete.Ing()),
		},
		{
			name: "t3",
			op:   types.Update,
			want: pretty.Blue(types.Update.Ing()),
		},
		{
			name: "t4",
			op:   types.UnChange,
			want: pretty.Gray(types.UnChange.Ing()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.op.PrettyString(); got != tt.want {
				t.Errorf("ActionType.PrettyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChangeStep_Diff(t *testing.T) {
	type fields struct {
		ID  string
		Op  types.ActionType
		Old interface{}
		New interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "t1",
			fields: fields{
				ID:  "id",
				Op:  types.Create,
				Old: nil,
				New: nil,
			},
			want: `[32;1m[32;1mID: [0m[0m[32mid[0m
[32m[0m[32;1m[32;1mPlan: [0m[0m[32mCreating[0m
[32;1m[32;1mDiff: [0m[0m

`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := &ChangeStep{
				ID:       tt.fields.ID,
				Action:   tt.fields.Op,
				Original: tt.fields.Old,
				Modified: tt.fields.New,
			}
			got, err := cs.Diff()
			if (err != nil) != tt.wantErr {
				t.Errorf("ChangeStep.Diff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ChangeStep.Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChanges_Get(t *testing.T) {
	type fields struct {
		order   *ChangeOrder
		project *projectstack.Project
		stack   *projectstack.Stack
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ChangeStep
	}{
		{
			name: "t1",
			fields: fields{
				order: &ChangeOrder{
					ChangeSteps: map[string]*ChangeStep{
						"test-key": TestChangeStepOpCreate,
					},
				},
			},
			args: args{
				key: "test-key",
			},
			want: TestChangeStepOpCreate,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Changes{
				ChangeOrder: tt.fields.order,
				project:     tt.fields.project,
				stack:       tt.fields.stack,
			}
			if got := p.Get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Changes.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChanges_Values(t *testing.T) {
	type fields struct {
		order   *ChangeOrder
		project *projectstack.Project
		stack   *projectstack.Stack
	}
	type args struct {
		filters []ChangeStepFilterFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*ChangeStep
	}{
		{
			name: "filter-opcreate",
			fields: fields{
				order: &ChangeOrder{
					StepKeys:    TestStepKeys,
					ChangeSteps: TestChangeSteps,
				},
			},
			args: args{
				filters: []ChangeStepFilterFunc{CreateChangeStepFilter},
			},
			want: []*ChangeStep{TestChangeStepOpCreate},
		},
		{
			name: "filter-opdelete",
			fields: fields{
				order: &ChangeOrder{
					StepKeys:    TestStepKeys,
					ChangeSteps: TestChangeSteps,
				},
			},
			args: args{
				filters: []ChangeStepFilterFunc{DeleteChangeStepFilter},
			},
			want: []*ChangeStep{TestChangeStepOpDelete},
		},
		{
			name: "filter-opupdate",
			fields: fields{
				order: &ChangeOrder{
					StepKeys:    TestStepKeys,
					ChangeSteps: TestChangeSteps,
				},
			},
			args: args{
				filters: []ChangeStepFilterFunc{UpdateChangeStepFilter},
			},
			want: []*ChangeStep{TestChangeStepOpUpdate},
		},
		{
			name: "filter-opunchange",
			fields: fields{
				order: &ChangeOrder{
					StepKeys:    TestStepKeys,
					ChangeSteps: TestChangeSteps,
				},
			},
			args: args{
				filters: []ChangeStepFilterFunc{UnChangeChangeStepFilter},
			},
			want: []*ChangeStep{TestChangeStepOpUnChange},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Changes{
				ChangeOrder: tt.fields.order,
				project:     tt.fields.project,
				stack:       tt.fields.stack,
			}
			if got := p.Values(tt.args.filters...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Changes.Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChanges_Stack(t *testing.T) {
	type fields struct {
		order   *ChangeOrder
		project *projectstack.Project
		stack   *projectstack.Stack
	}
	tests := []struct {
		name   string
		fields fields
		want   *projectstack.Stack
	}{
		{
			name: "t1",
			fields: fields{
				order:   &ChangeOrder{StepKeys: []string{}, ChangeSteps: map[string]*ChangeStep{}},
				project: &projectstack.Project{},
				stack: &projectstack.Stack{
					StackConfiguration: projectstack.StackConfiguration{
						Name: "test-name",
					},
					Path: "test-path",
				},
			},
			want: &projectstack.Stack{
				StackConfiguration: projectstack.StackConfiguration{
					Name: "test-name",
				},
				Path: "test-path",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Changes{
				ChangeOrder: tt.fields.order,
				project:     tt.fields.project,
				stack:       tt.fields.stack,
			}
			if got := p.Stack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Changes.Stack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChanges_Project(t *testing.T) {
	type fields struct {
		order   *ChangeOrder
		project *projectstack.Project
		stack   *projectstack.Stack
	}
	tests := []struct {
		name   string
		fields fields
		want   *projectstack.Project
	}{
		{
			name: "t1",
			fields: fields{
				project: &projectstack.Project{
					ProjectConfiguration: projectstack.ProjectConfiguration{
						Name:   "test-name",
						Tenant: "test-tenant",
					},
					Path: "test-path",
				},
			},
			want: &projectstack.Project{
				ProjectConfiguration: projectstack.ProjectConfiguration{
					Name:   "test-name",
					Tenant: "test-tenant",
				},
				Path: "test-path",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewChanges(tt.fields.project, tt.fields.stack, tt.fields.order)
			if got := p.Project(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Changes.Project() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChanges_Diffs(t *testing.T) {
	type fields struct {
		order   *ChangeOrder
		project *projectstack.Project
		stack   *projectstack.Stack
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "t1",
			fields: fields{
				order: &ChangeOrder{
					StepKeys: []string{"test-key"},
					ChangeSteps: map[string]*ChangeStep{
						"test-key": TestChangeStepOpCreate,
					},
				},
			},
			want: `[32;1m[32;1mID: [0m[0m[32mid[0m
[32m[0m[32;1m[32;1mPlan: [0m[0m[32mCreating[0m
[32;1m[32;1mDiff: [0m[0m

`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Changes{
				ChangeOrder: tt.fields.order,
				project:     tt.fields.project,
				stack:       tt.fields.stack,
			}
			if got := p.Diffs(); got != tt.want {
				t.Errorf("Changes.Diffs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChanges_Preview(t *testing.T) {
	type fields struct {
		order   *ChangeOrder
		project *projectstack.Project
		stack   *projectstack.Stack
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "t1",
			fields: fields{
				order: &ChangeOrder{
					StepKeys: []string{"test-key"},
					ChangeSteps: map[string]*ChangeStep{
						"test-key": TestChangeStepOpCreate,
					},
				},
				stack: &projectstack.Stack{
					StackConfiguration: projectstack.StackConfiguration{
						Name: "test-name",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Changes{
				ChangeOrder: tt.fields.order,
				project:     tt.fields.project,
				stack:       tt.fields.stack,
			}
			p.Summary()
		})
	}
}

func Test_buildResourceStateMap(t *testing.T) {
	type args struct {
		rs []*models.Resource
	}
	tests := []struct {
		name string
		args args
		want map[string]*models.Resource
	}{
		{
			name: "t1",
			args: args{
				rs: []*models.Resource{},
			},
			want: map[string]*models.Resource{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildResourceStateMap(tt.args.rs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildResourceStateMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
