package grpc

import (
	"context"
	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/usecase"
	"reflect"
	"testing"
)

func Test_myServer_ReadAnswer(t *testing.T) {
	type fields struct {
		UnimplementedChgkServer pb.UnimplementedChgkServer
		game                    usecase.Game
	}
	type args struct {
		ctx     context.Context
		message *pb.AnswerRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Answer
		wantErr bool
	}{
		{
			name: "empty question id",
			args: args{
				ctx: context.Background(),
				message: &pb.AnswerRequest{
					QuestionID: 0,
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
			got, err := s.ReadAnswer(tt.args.ctx, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAnswer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAnswer() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myServer_ReadQuestion(t *testing.T) {
	type fields struct {
		UnimplementedChgkServer pb.UnimplementedChgkServer
		game                    usecase.Game
	}
	type args struct {
		ctx     context.Context
		message *pb.QuestionRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Question
		wantErr bool
	}{
		{
			name: "empty question id",
			args: args{
				ctx: context.Background(),
				message: &pb.QuestionRequest{
					QuestionID: 0,
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
			got, err := s.ReadQuestion(tt.args.ctx, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadQuestion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myServer_ReadScore(t *testing.T) {
	type fields struct {
		UnimplementedChgkServer pb.UnimplementedChgkServer
		game                    usecase.Game
	}
	type args struct {
		ctx     context.Context
		message *pb.ScoreRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Score
		wantErr bool
	}{
		{
			name: "empty chat id",
			args: args{
				ctx: context.Background(),
				message: &pb.ScoreRequest{
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
			got, err := s.ReadScore(tt.args.ctx, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadScore() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myServer_ReadTournament(t *testing.T) {
	type fields struct {
		UnimplementedChgkServer pb.UnimplementedChgkServer
		game                    usecase.Game
	}
	type args struct {
		ctx     context.Context
		message *pb.TournamentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.Tournament
		wantErr bool
	}{
		{
			name: "empty tournament key",
			args: args{
				ctx: context.Background(),
				message: &pb.TournamentRequest{
					Name: "",
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
			got, err := s.ReadTournament(tt.args.ctx, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadTournament() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadTournament() got = %v, want %v", got, tt.want)
			}
		})
	}
}
