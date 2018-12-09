package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/instahell/hells"
	//"math/rand"
	"sync"
	//"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	"golang.org/x/net/context"
)

type SessionManager struct {
	mu       sync.RWMutex
	hell	 map[hells.HellsId]*hells.HellsId
}
func NewSessionManager() *SessionManager {
	return &SessionManager{
		mu:       sync.RWMutex{},
		sessions: map[hells.HellsId]*hells.HellsId{},
	}
}
// самое красивое что тут только мжно было сделать
func (sm *SessionManager) Make(ctx context.Context, in *hells.HellsId) (*hells.HellsId, error) {
	fmt.Println("call Create", in)
	id := &hells.HellsId{uuid.New().String()}
	sm.mu.Lock()
	sm.sessions[*id] = in
	sm.mu.Unlock()
	return id, nil
}

/*
func (sm *SessionManager) Check(ctx context.Context, in *hells.HellsId) (*hells.HellsId, error) {
	fmt.Println("call Check", in)
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	if sess, ok := sm.sessions[*in]; ok {
		return sess, nil
	}
	return nil, grpc.Errorf(codes.NotFound, "hells not found")
}

func (sm *SessionManager) Delete(ctx context.Context, in *hells.HellsId) (*hells.Nothing, error) {
	fmt.Println("call Delete", in)
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.sessions, *in)
	return &hells.Nothing{Dummy: true}, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
*/
