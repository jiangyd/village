-- 添加联合字段约束
ALTER TABLE dz ADD UNIQUE KEY (type,type_id,uid_id);
ALTER TABLE collection ADD UNIQUE KEY (type,type_id,uid_id);
ALTER TABLE firend ADD UNIQUE KEY (user_a_id,user_b_id);

--创建session表
CREATE TABLE `session` (
	`session_key` char(64) NOT NULL,
	`session_data` blob,
	`session_expiry` int(11) unsigned NOT NULL,
	PRIMARY KEY (`session_key`)
	) ENGINE=MyISAM DEFAULT CHARSET=utf8;