DROP TABLE IF EXISTS gva.user_extra;
CREATE TABLE gva.user_extra (
	id INT auto_increment NOT NULL,
	sys_user_id INT NULL,
	agent_code varchar(100) NULL,
	agent_name varchar(100) NULL,
	service_type varchar(100) NULL,
	ucc_id varchar(100) NULL,
	CONSTRAINT user_extra_pk PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
