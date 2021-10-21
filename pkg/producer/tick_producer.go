package producer

import (
	"log"

	"github.com/ergo-services/ergo/etf"
	"github.com/ergo-services/ergo/gen"
	"github.com/twmb/franz-go/pkg/kgo"
)

func NewTickProducer(name string, seedBrokers []string) gen.ServerBehavior {
	return &TickProducer{
		writeWS:     make(chan interface{}),
		name:        name,
		seedBrokers: seedBrokers,
	}
}

type TickProducer struct {
	gen.Server
	kClient     *kgo.Client
	writeWS     chan interface{}
	name        string
	seedBrokers []string
}

func (p *TickProducer) Init(process *gen.ServerProcess, args ...etf.Term) error {
	kClient, err := kgo.NewClient(
		kgo.SeedBrokers(p.seedBrokers...),
	)
	if err != nil {
		return err
	}
	p.kClient = kClient

	return nil
}

func (p *TickProducer) HandleCast(process *gen.ServerProcess, message etf.Term) gen.ServerStatus {
	log.Printf("HandleCast unknown message=%v, producer=%s\n", message, p.name)
	return gen.ServerStatusOK
}

func (p *TickProducer) HandleCall(process *gen.ServerProcess, from gen.ServerFrom, message etf.Term) (etf.Term, gen.ServerStatus) {
	log.Printf("HandleCall message=%v, producer=%s\n", message, p.name)
	return process.State, gen.ServerStatusOK
}

func (p *TickProducer) Terminate(process *gen.ServerProcess, reason string) {
	defer p.kClient.Close()
	log.Printf("Terminate reason=%s, producer=%s\n", reason, p.name)
}
