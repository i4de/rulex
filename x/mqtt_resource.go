package x

import (
	"fmt"
	"time"

	"github.com/ngaut/log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//
const DEFAULT_CLIENTID string = "X_IN_END_CLIENT"
const DEFAULT_USERNAME string = "X_IN_END"
const DEFAULT_PASSWORD string = "X_IN_END"
const DEFAULT_TOPIC string = "$X_IN_END"

//
type MqttInEndResource struct {
	enabled bool
	inEndId string
	client  mqtt.Client
}

func NewMqttInEndResource(inEndId string) *MqttInEndResource {
	return &MqttInEndResource{
		inEndId: inEndId,
	}
}

func (mm *MqttInEndResource) Start(e *RuleEngine) error {

	var messageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		log.Infof("Received message: [%s] from topic: [%s]\n", msg.Payload(), msg.Topic())
		if mm.enabled {
			e.Work(e.GetInEnd(mm.inEndId), string(msg.Payload()))
		}
	}
	//
	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		log.Infof("Mqtt InEnd Connected Success")
		// TODO support multipul topics
		client.Subscribe(DEFAULT_TOPIC, 1, nil)
		e.GetInEnd(mm.inEndId).State = 1
	}

	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		log.Infof("Connect lost: %v\n", err)
		e.GetInEnd(mm.inEndId).State = 0
	}
	config := e.GetInEnd(mm.inEndId).Config
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%v", (*config)["server"], (*config)["port"].(int)))
	if (*config)["clientId"] != nil {
		opts.SetClientID("x-client-main-" + (*config)["clientId"].(string))
	} else {
		opts.SetPassword(DEFAULT_CLIENTID)
	}
	if (*config)["username"] != nil {
		opts.SetUsername((*config)["username"].(string))
	} else {
		opts.SetPassword(DEFAULT_USERNAME)
	}
	if (*config)["password"] != nil {
		opts.SetPassword((*config)["password"].(string))
	} else {
		opts.SetPassword(DEFAULT_PASSWORD)
	}
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	opts.SetDefaultPublishHandler(messageHandler)
	opts.SetPingTimeout(10 * time.Second)
	opts.SetAutoReconnect(true)
	opts.OnReconnecting = func(mqtt.Client, *mqtt.ClientOptions) {
		log.Warn("Try to reconnect")
	}
	opts.SetMaxReconnectInterval(5 * time.Second)
	mm.client = mqtt.NewClient(opts)
	mm.enabled = true
	if token := mm.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	} else {
		return nil
	}

}
func (mm *MqttInEndResource) Stop() {

}
func (mm *MqttInEndResource) Reload() {

}
func (mm *MqttInEndResource) Pause() {

}
func (mm *MqttInEndResource) Status(e *RuleEngine) TargetState {
	return e.GetInEnd(mm.inEndId).State
}

func (mm *MqttInEndResource) Register(inEndId string) error {

	return nil
}

func (mm *MqttInEndResource) Test(inEndId string) bool {
	return true
}

func (mm *MqttInEndResource) Enabled() bool {
	return mm.enabled
}