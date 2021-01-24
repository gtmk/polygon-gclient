package polygonio

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

const (
	MaxConnectionAttempts = 3
)

var (
	once   sync.Once
	stream *Stream
)

type credentials struct {
	apiKey         string
	streamEndpoint string
}

type Stream struct {
	sync.Mutex
	sync.Once
	conn                  *websocket.Conn
	authenticated, closed atomic.Value
	credentials           credentials

	MessageC chan []byte
	ErrorC   chan error
}

func GetStream(apiKey, streamEndpoint string) (*Stream, error) {
	once.Do(func() {
		stream = &Stream{
			authenticated: atomic.Value{},
			MessageC:      make(chan []byte, 100),
			ErrorC:        make(chan error, 100),
			credentials: credentials{apiKey: apiKey,
				streamEndpoint: streamEndpoint},
		}
		stream.authenticated.Store(false)
		stream.closed.Store(false)
	})
	err := stream.register()
	if err != nil {
		return nil, err
	}
	return stream, nil
}

func (s *Stream) register() error {
	var err error
	if s.conn == nil {
		s.conn, err = s.openSocket()
		if err != nil {
			return err
		}
	}
	if err = s.auth(); err != nil {
		return err
	}
	return nil
}

func (s *Stream) Subscribe(channel string) error {
	s.Do(func() {
		go s.start()
	})
	if err := s.sub(channel); err != nil {
		return err
	}
	return nil
}

func (s *Stream) Unsubscribe(channel string) error {
	var err error
	if s.conn == nil {
		return errors.New("connection has not been initialized")
	}
	if err = s.auth(); err != nil {
		return err
	}
	if err = s.unsub(channel); err != nil {
		return err
	}
	return nil
}

func (s *Stream) Close() error {
	s.Lock()
	defer s.Unlock()

	if s.conn == nil {
		return nil
	}

	if err := s.conn.WriteMessage(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
	); err != nil {
		return err
	}
	s.closed.Store(true)
	close(s.MessageC)
	close(s.ErrorC)
	return s.conn.Close()
}

func (s *Stream) openSocket() (*websocket.Conn, error) {
	connectionAttempts := 0
	for connectionAttempts < MaxConnectionAttempts {
		conn, _, err := websocket.DefaultDialer.Dial(s.credentials.streamEndpoint, nil)
		if err != nil {
			if connectionAttempts == MaxConnectionAttempts {
				return nil, err
			}
		} else {
			msg := []PolgyonServerMsg{}
			if err = conn.ReadJSON(&msg); err == nil {
				return conn, err
			}
		}
		time.Sleep(1 * time.Second)
		connectionAttempts++
	}
	return nil, fmt.Errorf("Error: Could not open Polygon stream (max retries exceeded).")
}

func (s *Stream) isAuthenticated() bool {
	return s.authenticated.Load().(bool)
}

func (s *Stream) auth() error {
	s.Lock()
	defer s.Unlock()

	if s.isAuthenticated() {
		return nil
	}

	authRequest := PolygonClientMsg{
		Action: "auth",
		Params: s.credentials.apiKey,
	}

	if err := s.conn.WriteJSON(authRequest); err != nil {
		return err
	}
	msg := []PolygonAuthMsg{}
	// ensure the auth response comes in a timely manner
	s.conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	defer s.conn.SetReadDeadline(time.Time{})

	if err := s.conn.ReadJSON(&msg); err != nil {
		return err
	}

	if !strings.EqualFold(msg[0].Status, "auth_success") {
		return fmt.Errorf("failed to authorize Polygon stream")
	}

	s.authenticated.Store(true)

	return nil
}

func (s *Stream) start() {
	for {
		code, bts, err := s.conn.ReadMessage()
		if err != nil {
			if s.closed.Load().(bool) {
				return
			} else if websocket.IsCloseError(err, code) {
				err := s.reconnect()
				if err != nil {
					s.ErrorC <- err
					return
				}
				continue
			} else {
				s.ErrorC <- err
				continue
			}
		}
		s.MessageC <- bts
	}
}

func (s *Stream) sub(channel string) error {
	s.Lock()
	defer s.Unlock()

	subReq := PolygonClientMsg{
		Action: "subscribe",
		Params: channel,
	}
	if err := s.conn.WriteJSON(subReq); err != nil {
		return err
	}
	return nil
}

func (s *Stream) unsub(channel string) error {
	s.Lock()
	defer s.Unlock()
	subReq := PolygonClientMsg{
		Action: "unsubscribe",
		Params: channel,
	}
	err := s.conn.WriteJSON(subReq)
	return err
}

func (s *Stream) reconnect() error {
	var err error
	s.authenticated.Store(false)
	if s.conn, err = s.openSocket(); err != nil {
		return err
	}
	if err = s.auth(); err != nil {
		return err
	}
	return nil
}
