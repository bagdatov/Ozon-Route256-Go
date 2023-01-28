package grpc

import (
	"context"
	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/usecase"
	"reflect"
	"testing"
)

func Test_myServer_Submit(t *testing.T) {
	type fields struct {
		UnimplementedChgkServer pb.UnimplementedChgkServer
		game                    usecase.Game
	}
	type args struct {
		ctx context.Context
		m   *pb.Guess
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GuessResponse
		wantErr bool
	}{
		{
			name: "empty question id",
			args: args{
				ctx: context.Background(),
				m: &pb.Guess{
					QuestionID: 0,
				},
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "empty chat id",
			args: args{
				ctx: context.Background(),
				m: &pb.Guess{
					ChatID: 0,
				},
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "empty username",
			args: args{
				ctx: context.Background(),
				m: &pb.Guess{
					Username: "",
				},
			},
			want:    nil,
			wantErr: true,
		},

		{
			name: "empty answer",
			args: args{
				ctx: context.Background(),
				m: &pb.Guess{
					Answer: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := myServer{
				UnimplementedChgkServer: tt.fields.UnimplementedChgkServer,
				game:                    tt.fields.game,
			}
			got, err := s.Submit(tt.args.ctx, tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Submit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Submit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
