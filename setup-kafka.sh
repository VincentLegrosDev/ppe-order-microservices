docker compose -f bitnami-kafka.yml up -d

docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --create --topic Order-received
docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic Order-received

docker exec -it kafka kafka-configs.sh --bootstrap-server kafka:9092 --entity-type topics --entity-name  Order-received  --alter --add-config retention.ms=259200000

docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --create --topic Order-confirmed
docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic Order-confirmed

docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --create --topic Order-picked 
docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic Order-picked

docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --create --topic Order-packed 
docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic Order-packed

docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --create --topic Notifications
docker exec -it kafka kafka-topics.sh --bootstrap-server kafka:9092 --describe --topic Notifications

