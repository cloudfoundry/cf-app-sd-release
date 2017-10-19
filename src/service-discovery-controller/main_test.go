package main_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"fmt"
	"time"

	"github.com/nats-io/gnatsd/server"
	"github.com/nats-io/nats"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {

	var (
		session      *gexec.Session
		pathToConfig string
		configJson   string
		natsServer   *server.Server
		routeEmitter *nats.Conn
	)

	BeforeEach(func() {
		natsServer = RunNatsServerOnPort(8080)
		config, err := ioutil.TempFile(os.TempDir(), "sd_config")
		Expect(err).ToNot(HaveOccurred())
		pathToConfig = config.Name()
		configJson = `{
			"address":"127.0.0.1",
			"port":"8055",
			"nats":[
				{
					"host":"localhost",
					"port":8080,
					"user":"",
					"pass":""
				}
			]
		}`

		err = ioutil.WriteFile(pathToConfig, []byte(configJson), os.ModePerm)
		Expect(err).ToNot(HaveOccurred())

		startCmd := exec.Command(pathToServer, "-c", pathToConfig)
		session, err = gexec.Start(startCmd, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())

		time.Sleep(1 * time.Second) // TODO : get rid of me

		routeEmitter = newFakeRouteEmitter("nats://" + natsServer.Addr().String())

		register(routeEmitter, "192.168.0.1", "app-id.internal.local.")
		register(routeEmitter, "192.168.0.2", "app-id.internal.local.")
		register(routeEmitter, "192.168.0.1", "large-id.internal.local.")
		register(routeEmitter, "192.168.0.2", "large-id.internal.local.")
		register(routeEmitter, "192.168.0.3", "large-id.internal.local.")
		register(routeEmitter, "192.168.0.4", "large-id.internal.local.")
		register(routeEmitter, "192.168.0.5", "large-id.internal.local.")
		register(routeEmitter, "192.168.0.6", "large-id.internal.local.")
		register(routeEmitter, "192.168.0.7", "large-id.internal.local.")
		register(routeEmitter, "192.168.0.8", "large-id.internal.local.")
		register(routeEmitter, "192.168.0.9", "large-id.internal.local.")
		register(routeEmitter, "192.168.0.10", "large-id.internal.local.")
		register(routeEmitter, "192.168.0.11", "large-id.internal.local.")
		register(routeEmitter, "192.168.0.12", "large-id.internal.local.")
		register(routeEmitter, "192.168.0.13", "large-id.internal.local.")
		Expect(routeEmitter.Flush()).ToNot(HaveOccurred()) //TODO: necessary?
	})

	AfterEach(func() {
		session.Kill()
		os.Remove(pathToConfig)
		routeEmitter.Close()
		natsServer.Shutdown()
	})

	It("accepts interrupt signals and shuts down", func() {
		Eventually(session).Should(gbytes.Say("Server Started"))
		session.Signal(os.Interrupt)

		Eventually(session).Should(gexec.Exit())
		Eventually(session).Should(gbytes.Say("Shutting service-discovery-controller down"))
	})

	PIt("should not return ips for unregistered domains", func() {
		Fail("")
	})

	It("should return a http app json", func() {
		Eventually(session).Should(gbytes.Say("Server Started"))

		req, err := http.NewRequest("GET", "http://localhost:8055/v1/registration/app-id.internal.local.", nil)
		Expect(err).ToNot(HaveOccurred())
		resp, err := http.DefaultClient.Do(req)
		Expect(err).ToNot(HaveOccurred())
		respBody, err := ioutil.ReadAll(resp.Body)
		Expect(err).ToNot(HaveOccurred())

		Expect(respBody).To(MatchJSON(`{
			"env": "",
			"hosts": [
			{
				"ip_address": "192.168.0.1",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			},
			{
				"ip_address": "192.168.0.2",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			}],
			"service": ""
		}`))
	})

	It("should return a http large json", func() {
		Eventually(session).Should(gbytes.Say("Server Started"))

		req, err := http.NewRequest("GET", "http://localhost:8055/v1/registration/large-id.internal.local.", nil)
		Expect(err).ToNot(HaveOccurred())
		resp, err := http.DefaultClient.Do(req)
		Expect(err).ToNot(HaveOccurred())
		respBody, err := ioutil.ReadAll(resp.Body)
		Expect(err).ToNot(HaveOccurred())

		Expect(respBody).To(MatchJSON(`{
			"env": "",
			"hosts": [
			{
				"ip_address": "192.168.0.1",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			},
			{
				"ip_address": "192.168.0.2",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			},
			{
				"ip_address": "192.168.0.3",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			},
			{
				"ip_address": "192.168.0.4",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			},
			{
				"ip_address": "192.168.0.5",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			},
			{
				"ip_address": "192.168.0.6",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			},
			{
				"ip_address": "192.168.0.7",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			},
			{
				"ip_address": "192.168.0.8",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			},
			{
				"ip_address": "192.168.0.9",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			},
			{
				"ip_address": "192.168.0.10",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			},
			{
				"ip_address": "192.168.0.11",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			},
			{
				"ip_address": "192.168.0.12",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			},
						{
				"ip_address": "192.168.0.13",
				"last_check_in": "",
				"port": 0,
				"revision": "",
				"service": "",
				"service_repo_name": "",
				"tags": {}
			}
			],
			"service": ""
		}`))
	})
})

func register(routeEmitter *nats.Conn, ip string, url string) {
	natsRegistryMsg := nats.Msg{
		Subject: "service-discovery.register",
		Data:    []byte(fmt.Sprintf(`{"host": "%s","uris":["%s"]}`, ip, url)),
	}

	Expect(routeEmitter.PublishMsg(&natsRegistryMsg)).ToNot(HaveOccurred())
}

func newFakeRouteEmitter(natsUrl string) *nats.Conn {
	natsClient, err := nats.Connect(natsUrl, nats.ReconnectWait(1*time.Nanosecond))
	Expect(err).NotTo(HaveOccurred())
	return natsClient
}
