package producer

import (
	"log"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/twmb/franz-go/pkg/kgo"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var (
	wsReadLimit int64 = 1048576 // 1024 Kb
)

func NewWebsocketProducer(name, url string, seedBrokers []string) gen.ServerBehavior {
	return &WebsocketProducer{
		writeWS:     make(chan interface{}),
		name:        name,
		url:         url,
		seedBrokers: seedBrokers,
	}
}

type WebsocketProducer struct {
	gen.Server
	wsConn      *websocket.Conn
	kClient     *kgo.Client
	writeWS     chan interface{}
	name        string
	url         string
	seedBrokers []string
}

func (p *WebsocketProducer) Init(process *gen.ServerProcess, args ...etf.Term) error {
	wsConn, _, err := websocket.Dial(process.Context(), p.url, nil)
	if err != nil {
		return err
	}
	wsConn.SetReadLimit(wsReadLimit)
	p.wsConn = wsConn

	kClient, err := kgo.NewClient(
		kgo.SeedBrokers(p.seedBrokers...),
	)
	if err != nil {
		return err
	}
	p.kClient = kClient

	go p.readPump(process)
	go p.writePump(process)

	return nil
}

func (p *WebsocketProducer) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	switch m := message.(type) {
	case []interface{}:
		wsMsgType := m[0].(websocket.MessageType)
		wsMsg := m[1].([]byte)
		log.Printf("HandleCast type=%v, msg=%s, producer=%s\n", wsMsgType, wsMsg, p.name)
	default:
		log.Printf("HandleCast unknown message=%v, producer=%s\n", message, p.name)
	}

	return gen.ServerStatusOK
}

func (p *WebsocketProducer) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	log.Printf("HandleCall message=%v, producer=%s\n", message, p.name)
	return process.State, gen.ServerStatusOK
}

func (p *WebsocketProducer) Terminate(process *gen.ServerProcess, reason string) {
	defer func() {
		p.kClient.Close()
		err := p.wsConn.Close(websocket.StatusAbnormalClosure, reason)
		if err != nil {
			log.Printf("Terminate wsConn.Close error=%s, producer=%s", err, p.name)
		}
	}()
	log.Printf("Terminate reason=%s, producer=%s\n", reason, p.name)
}

func (p *WebsocketProducer) readPump(process *gen.ServerProcess) {
	for {
		msgType, msg, err := p.wsConn.Read(process.Context())
		if err != nil {
			log.Printf("readPump error=%v, producer=%s\n", err, p.name)
			break
		}
		process.Cast(process.Self(), []interface{}{msgType, msg})
	}
}

func (p *WebsocketProducer) writePump(process *gen.ServerProcess) {
	for {
		select {
		case <-process.Context().Done():
			break
		case msg := <-p.writeWS:
			err := wsjson.Write(process.Context(), p.wsConn, msg)
			if err != nil {
				log.Printf("writePump error=%s, producer=%s\n", err, p.name)
				break
			}
		}
	}
}
