package {{.ContextName}}

import (
  "context"
  "fmt"
  "log"
	"os"
	"os/signal"
  "sync"
  "time"

  "github.com/twmb/franz-go/pkg/kgo"
)

func New{{.ProducerName | Title}}Producer(seeds []string) *{{.ProducerName | Title}}Producer {
  return &{{.ProducerName | Title}}Producer{
    name: "{{.ContextName}}/{{.ProducerName | Title}}Producer",
    seeds: seeds,
  }
}

type {{.ProducerName | Title}}Producer struct {
  name  string
  seeds []string
}

func (s *{{.ProducerName | Title}}Producer) Start() error {
  log.Printf("starting producer %s\n", s.name)

  cl, err := kgo.NewClient(
    kgo.SeedBrokers(s.seeds...),
  )
  if err != nil {
    return err
  }
  defer cl.Close()

  signalChan := make(chan os.Signal, 1)
  signal.Notify(signalChan, os.Interrupt)
  ctx := context.Background()
  wg := sync.WaitGroup{}
  wg.Add(1)
  go func() {
    select {
      case <- signalChan:
        wg.Done()
        return
      case t := <-time.After(10 * time.Second):
        topic := "{{.ContextName | ToLower}}.{{.ProducerName | ToLower}}"
        tickMsg := fmt.Sprintf("{{.ContextName}}/{{.ProducerName}} time=%s", t.String())
        record := &kgo.Record{Topic: topic, Value: []byte(tickMsg)}
        cl.Produce(ctx, record, func(_ *kgo.Record, err error) {
          if err != nil {
            log.Printf("record had a produce error: %v\n", err)
          }
        })
    }
  }()
  wg.Wait()

  return nil
}
