
CREATE TABLE `counter` (
  `id` bigint(20) NOT NULL,
  `name` varchar(255) NOT NULL,
  `unit` varchar(255) NOT NULL,
  `reading` decimal(20,6) DEFAULT 0,
  `last_tick` DATETIME(3),
  PRIMARY KEY (`id`)
);
