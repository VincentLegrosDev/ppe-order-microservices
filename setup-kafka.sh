docker compose -f bitnami-kafka.yml up -d

docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --create  --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderReceived 
docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic OrderReceived

docker exec -it kafka kafka-configs.sh --bootstrap-server kafka:9092 --entity-type topics --entity-name  OrderReceived  --alter --add-config retention.ms=259200000

docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --create --replication-factor 1 --partitions 1 --config retention.ms=10800000  --topic OrderConfirmed
docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic OrderConfirmed

docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --create --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderPickedAndPacked
docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic OrderPickedAndPacked

docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --create --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic Notification
docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic Notification

docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --create --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic Error
docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic Error

docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --create --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic DeadQueueLetter
docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic DeadQueueLetter

docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --create --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderCountMetric
docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic  OrderCountMetric 

docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --create --replication-factor 1 --partitions 1 --config retention.ms=10800000 --topic OrderTimeMetric
docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic  OrderTimeMetric 

