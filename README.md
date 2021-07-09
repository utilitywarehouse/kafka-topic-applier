# kafka-topic-applier (KTA)

Your trusty neighbourhood topic `orchestrator`

Supply a topics.yaml file via a configMap and KTA will ensure all topics are created, and any marked for deletion removed. Who could ask for anything more?

Use the client to create your YAML file

```
client list --yaml
```

## Topic File

the topics.yaml file can be used to set up the topic with some Kafka settings:

```yaml
topics:
  - name:
    partitions:
    replication_factor:
    topic_compression: [uncompressed|gzip|lz4|snappy]
    max_message_bytes:
    max_retention_bytes:
    max_rentention_time:
    cleanup_policy: [delete|compact|delete, compact]
    should_be_removed: [true|false]
```

More details about these settings can be found on the [Kafka documentation](https://kafka.apache.org/documentation/#topicconfigs).
