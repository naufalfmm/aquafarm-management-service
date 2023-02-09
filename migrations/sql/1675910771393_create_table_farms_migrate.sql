create table `farms` (
	`id` bigint unsigned not null auto_increment,
	`code` varchar(25) not null,
	`description` varchar(255) not null default '',
	`address` varchar(255) not null,
	`village` varchar(255) not null,
	`district` varchar(255) not null,
	`city` varchar(255) not null,
	`province` varchar(255) not null,
	`postal_code` varchar(255) not null,
	`latitude` float null,
	`longitude` float null,
	`created_at` datetime not null default current_timestamp,
	`updated_at` datetime not null default current_timestamp,
	`deleted_at` datetime default null,
	`created_by` varchar(255) not null,
	`updated_by` varchar(255) not null,
	`deleted_by` varchar(255) null,
	`deleted_unix` int not null default 0,
	
	primary key (`id`),
	unique key `uq_farm_code` (`code`, `deleted_unix`)
);