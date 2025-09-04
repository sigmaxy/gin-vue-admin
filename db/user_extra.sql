CREATE TABLE `user_extra` (
  `id` int NOT NULL AUTO_INCREMENT,
  `sys_user_id` int DEFAULT NULL,
  `agent_code` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `agent_name` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `service_type` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `ucc_id` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;