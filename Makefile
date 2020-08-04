INIT_DIR=../kafka-rest-sdk-go/
GENERATOR_DIR=../openapi-generator
OPENAPI_SPEC_DIR=../kafka-rest/api/v3/kafka-rest-v3-spec.yaml
SDK_OUTPUT_DIR=../kafka-rest-sdk-go/kafkarestv3


.PHONY: deps
deps:
	# https://github.com/travisjeffery/mocker
	go install github.com/travisjeffery/mocker/cmd/mocker

	# https://github.com/confluentinc/openapi-generator
	# cd .. && git clone https://github.com/confluentinc/openapi-generator.git
	# cd ../openapi-generator && git checkout ath-add-interfaces && mvn clean install


.PHONY: build
build:
	cd $(GENERATOR_DIR) && java -jar modules/openapi-generator-cli/target/openapi-generator-cli.jar generate \
		-i $(OPENAPI_SPEC_DIR) \
		-g go \
		-o $(SDK_OUTPUT_DIR) \
  		-p enumClassPrefix=true,packageName=kafkarestv3,generateInterfaces=true

.PHONY: mock
mock:
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_acl.go --pkg mock $(SDK_OUTPUT_DIR)/api_acl.go ACLApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_broker.go --pkg mock $(SDK_OUTPUT_DIR)/api_broker.go BrokerApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_cluster.go --pkg mock $(SDK_OUTPUT_DIR)/api_cluster.go ClusterApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_configs.go --pkg mock $(SDK_OUTPUT_DIR)/api_configs.go ConfigsApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_consumer_group.go --pkg mock $(SDK_OUTPUT_DIR)/api_consumer_group.go ConsumerGroupApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_partition.go --pkg mock $(SDK_OUTPUT_DIR)/api_partition.go PartitionApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_replica.go --pkg mock $(SDK_OUTPUT_DIR)/api_replica.go ReplicaApi
	mocker --prefix "" --dst $(SDK_OUTPUT_DIR)/mock/api_topic.go --pkg mock $(SDK_OUTPUT_DIR)/api_topic.go TopicApi