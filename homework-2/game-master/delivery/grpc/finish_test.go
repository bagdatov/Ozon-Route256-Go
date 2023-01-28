package grpc

import (
	"context"
	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/usecase"
	"reflect"
	"testing"
)

func Test_myServer_Finish(t *testing.T) {
	type fields struct {
		UnimplementedChgkServer pb.UnimplementedChgkServer
		game                    usecase.Game
	}
	type args struct {
		ctx     context.Context
		message *pb.StopRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.StopResponse
		wantErr bool
	}{
		{
			name: "empty chat id",
			args: args{
				ctx: context.Background(),
				message: &pb.StopRequest{
					ChatID: 0,
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
			got, err := s.Finish(tt.args.ctx, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Finish() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Finish() got = %v, want %v", got, tt.want)
			}
		})
	}
}
