
CREATE TABLE `counter_event` (
  `id` bigint(20) NOT NULL,
  `counter_id` bigint(20) NOT NULL,
  `type` int(3) NOT NULL,
  `timestamp` TIMESTAMP DEFAULT NOW(),
  `value` decimal(20,6) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_COUNTER` (`counter_id`),
  CONSTRAINT `FK_COUNTER` FOREIGN KEY (`counter_id`) REFERENCES `counter` (`id`)
);

