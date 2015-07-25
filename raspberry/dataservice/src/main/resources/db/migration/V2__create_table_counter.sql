
CREATE TABLE `counter` (
  `id` bigint(20) NOT NULL,
  `name` varchar(255) NOT NULL,
  `unit` varchar(255) NOT NULL,
  `value` varchar(255) DEFAULT 0,
  `date_created` TIMESTAMP NULL,
  `date_changed` TIMESTAMP NULL,
  `last_tick` TIMESTAMP NULL,
  PRIMARY KEY (`id`)
);
