create TABLE `blog` (
	`id`         int(11)   NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
	`updated_at` TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE now(),
	`title`   varchar(53)  NOT NULL COMMENT '标题',
	`cover`   varchar(511) NOT NULL COMMENT '封面',
	`summary` varchar(255) NOT NULL COMMENT '简介',
	`path`    varchar(127) NOT NULL COMMENT '文件路径',
	`date`    datetime     NOT NULL COMMENT '时间',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='博客文章';
