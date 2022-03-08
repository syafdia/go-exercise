package auth

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/syafdia/xaam/internal/domain/entity"
	e "github.com/syafdia/xaam/internal/domain/entity"
	ae "github.com/syafdia/xaam/internal/domain/entity/auth"
	"github.com/syafdia/xaam/internal/domain/repo"
	mock_repo "github.com/syafdia/xaam/internal/mock/domain/repo"
)

func Test_findResourcesByComplianceUC_Execute(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockPolicyRepo := mock_repo.NewMockPolicyRepo(ctl)
	mockResourceRepo := mock_repo.NewMockResourceRepo(ctl)

	type fields struct {
		PolicyRepo   repo.PolicyRepo
		ResourceRepo repo.ResourceRepo
	}
	type args struct {
		ctx     context.Context
		request ae.FindResourcesByComplianceRequest
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		beforeTest func()
		want       map[e.Resource][]e.Action
		wantErr    bool
	}{
		{
			name: "on failed FindMultipleByIDandTargetType for industries",
			fields: fields{
				PolicyRepo:   mockPolicyRepo,
				ResourceRepo: mockResourceRepo,
			},
			args: args{
				ctx: context.TODO(),
				request: ae.FindResourcesByComplianceRequest{
					IndustryID:    1,
					LegalEntityID: 2,
				},
			},
			beforeTest: func() {
				mockPolicyRepo.EXPECT().
					FindMultipleByIDandTargetType(context.TODO(), int64(1), entity.TargetTypeIndustry).
					Return(nil, errors.New("whoops, error"))

			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "on failed FindMultipleByIDandTargetType for legal entity",
			fields: fields{
				PolicyRepo:   mockPolicyRepo,
				ResourceRepo: mockResourceRepo,
			},
			args: args{
				ctx: context.TODO(),
				request: ae.FindResourcesByComplianceRequest{
					IndustryID:    1,
					LegalEntityID: 2,
				},
			},
			beforeTest: func() {

				mockPolicyRepo.EXPECT().
					FindMultipleByIDandTargetType(context.TODO(), int64(1), entity.TargetTypeIndustry).
					Return([]e.Action{
						{ID: 1, Name: "compliant", ResourceID: 1},
					}, nil)

				mockPolicyRepo.EXPECT().
					FindMultipleByIDandTargetType(context.TODO(), int64(2), entity.TargetTypeLegalEntity).
					Return(nil, errors.New("whoops, error"))

			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "on success",
			fields: fields{
				PolicyRepo:   mockPolicyRepo,
				ResourceRepo: mockResourceRepo,
			},
			args: args{
				ctx: context.TODO(),
				request: ae.FindResourcesByComplianceRequest{
					IndustryID:    1,
					LegalEntityID: 2,
					Resources:     []string{"VA"},
				},
			},
			beforeTest: func() {

				mockPolicyRepo.EXPECT().
					FindMultipleByIDandTargetType(context.TODO(), int64(1), entity.TargetTypeIndustry).
					Return([]e.Action{
						{ID: 1, Name: "compliant", ResourceID: 10},
					}, nil)

				mockPolicyRepo.EXPECT().
					FindMultipleByIDandTargetType(context.TODO(), int64(2), entity.TargetTypeLegalEntity).
					Return([]e.Action{
						{ID: 1, Name: "compliant", ResourceID: 10},
						{ID: 3, Name: "compliant", ResourceID: 11},
					}, nil)

				mockResourceRepo.EXPECT().
					FindMultipleByIDs(context.TODO(), []int64{10}).
					Return(map[int64]e.Resource{
						10: {ID: 10, Name: "VA"},
					}, nil)

			},
			want: map[e.Resource][]e.Action{
				{ID: 10, Name: "VA"}: {
					{ID: 1, Name: "compliant", ResourceID: 10},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &findResourcesByComplianceUC{
				PolicyRepo:   tt.fields.PolicyRepo,
				ResourceRepo: tt.fields.ResourceRepo,
			}

			if tt.beforeTest != nil {
				tt.beforeTest()
			}

			got, err := f.Execute(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("findResourcesByComplianceUC.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findResourcesByComplianceUC.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
