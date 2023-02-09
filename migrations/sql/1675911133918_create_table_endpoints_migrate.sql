create table `endpoints` (
	`id` bigint unsigned not null auto_increment,
	`method` varchar(15) not null,
	`path` varchar(255) not null,
	`created_at` datetime not null default current_timestamp,
	`updated_at` datetime not null default current_timestamp,
	`deleted_at` datetime default null,
	`created_by` varchar(255) not null,
	`updated_by` varchar(255) not null,
	`deleted_by` varchar(255) null,
	
	primary key (`id`)
);