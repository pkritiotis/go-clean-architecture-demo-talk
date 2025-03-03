package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/pkritiotis/go-clean-architecture-example/internal/domain/runner"
	"github.com/stretchr/testify/assert"
)

func TestNewRepo(t *testing.T) {
	tests := []struct {
		name string
		want runner.Repository
	}{
		{
			name: "Should create an inmemory memory",
			want: Repo{
				runners: make(map[uuid.UUID]runner.Runner),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewRepo()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_inMemoryRepo_AddRunner(t *testing.T) {
	type fields struct {
		runners map[uuid.UUID]runner.Runner
	}
	type args struct {
		runner runner.Runner
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "should add runner",
			fields: fields{
				runners: make(map[uuid.UUID]runner.Runner),
			},
			args: args{
				runner: func() runner.Runner {
					r, _ := runner.NewRunner("John Doe", "johndoe@email.com")
					return r
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Repo{
				runners: tt.fields.runners,
			}
			err := m.Add(tt.args.runner)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Contains(t, m.runners, tt.args.runner.ID())
		})
	}
}
