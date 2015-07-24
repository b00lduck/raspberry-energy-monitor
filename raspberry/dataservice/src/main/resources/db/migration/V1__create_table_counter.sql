
CREATE TABLE `hibernate_sequences` (
  `sequence_name` varchar(255) DEFAULT NULL,
  `sequence_next_hi_value` int(11) DEFAULT NULL
);

CREATE TABLE `counter` (
  `id` bigint(20) NOT NULL,
  `name` varchar(255) NOT NULL,
  `unit` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `counter_event` (
  `id` bigint(20) NOT NULL,
  `counter_id` bigint(20) NOT NULL,
  `type` int(3) NOT NULL,
  `timestamp` TIMESTAMP DEFAULT NOW(),
  `delta` decimal(12,6) NOT NULL,
  `absolute` decimal(12,6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_COUNTER` (`counter_id`),
  CONSTRAINT `FK_COUNTER` FOREIGN KEY (`counter_id`) REFERENCES `counter` (`id`)
);

