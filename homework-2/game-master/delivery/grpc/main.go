package grpc

import (
	"fmt"
	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/usecase"
	"reflect"
)

// myServer -
type myServer struct {
	pb.UnimplementedChgkServer
	game usecase.Game
}

// New is a delivery grpc constructor
func New(p usecase.Game) (*myServer, error) {

	if isNil(p) {
		return nil, fmt.Errorf("usecase is nil")
	}

	return &myServer{
		game: p,
	}, nil
}

// isNil is using reflect
// package to determine rather input is nil
func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
