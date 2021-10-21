package {{.ContextName}}

func New{{.ProducerName | Title}}WebsocketProducer(name string, brokerSeeds []string) *{{.ProducerName | Title}}WebsocketProducer {
  return &{{.ProducerName | Title}}WebsocketProducer{
    WebsocketProducer{
      name: "{{.ContextName}}/{{.ProducerName | Title}}WebsocketProducer",
      brokerSeeds: brokerSeeds,
    },
  }
}

type {{.ProducerName | Title}}WebsocketProducer struct {
  WebsocketProducer
}
