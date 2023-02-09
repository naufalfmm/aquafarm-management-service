create table `ponds` (
	`id` bigint unsigned not null auto_increment,
	`farm_id` bigint unsigned not null,
	`code` varchar(25) not null,
	`description` varchar(255) not null default '',
	`wide` float not null,
	`long` float not null,
	`depth` float not null,
	`created_at` datetime not null default current_timestamp,
	`updated_at` datetime not null default current_timestamp,
	`deleted_at` datetime default null,
	`created_by` varchar(255) not null,
	`updated_by` varchar(255) not null,
	`deleted_by` varchar(255) null,
	`deleted_unix` int not null default 0,
	
	primary key (`id`),
	unique key `uq_pond_code_farm` (`farm_id`, `code`, `deleted_unix`)
);