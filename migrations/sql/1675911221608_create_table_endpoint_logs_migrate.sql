create table `endpoint_logs` (
	`id` bigint unsigned not null auto_increment,
	`endpoint_id` bigint not null,
	`request_id` varchar(40) not null,
	`uri` varchar(255) not null,
	`query` varchar(255) not null,
	`user_agent` varchar(255) not null,
	`ip_address` varchar(255) not null,
	`requested_by` varchar(255) null,
	`response_status_code` int not null,
	`start_at` bigint not null default 0,
	`end_at` bigint null,
	`created_at` datetime not null default current_timestamp,
	`updated_at` datetime not null default current_timestamp,
	`created_by` varchar(255) not null,
	`updated_by` varchar(255) not null,
	
	primary key (`id`)
);