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


-- 菜单数据
INSERT INTO `village`.`menu` (`key`, `title`) VALUES ('content', '内容管理');
INSERT INTO `village`.`menu` (`key`, `title`) VALUES ('system', '系统管理');

-- 子菜单数据
INSERT INTO `village`.`sub_menu` (`key`, `title`, `url`, `parent_id`) VALUES ('category', '分类', '/admin/categorymanagelist', 'system');
INSERT INTO `village`.`sub_menu` (`key`, `title`, `url`, `parent_id`) VALUES ('document', '文档', '/document', 'content');
INSERT INTO `village`.`sub_menu` (`key`, `title`, `url`, `parent_id`) VALUES ('menu', '菜单', '/admin/menumanagelist', 'system');
INSERT INTO `village`.`sub_menu` (`key`, `title`, `url`, `parent_id`) VALUES ('reply', '评论', '/admin/replymanagelist', 'content');
INSERT INTO `village`.`sub_menu` (`key`, `title`, `url`, `parent_id`) VALUES ('submenu', '子菜单', '/admin/submenumanagelist', 'system');
INSERT INTO `village`.`sub_menu` (`key`, `title`, `url`, `parent_id`) VALUES ('topic', '帖子', '/admin/topicmanagelist', 'content');
INSERT INTO `village`.`sub_menu` (`key`, `title`, `url`, `parent_id`) VALUES ('user', '用户', '/admin/usermanagelist', 'system');
INSERT INTO `village`.`sub_menu` (`key`, `title`, `url`, `parent_id`) VALUES ('site', '站点', '/admin/sitemanagelist', 'content');




--权限
insert into permission (`name`,`description`,`pname`) values ('content','内容管理','');
insert into permission (`name`,`description`,`pname`) values ('document','文档','content');
insert into permission (`name`,`description`,`pname`) values ('topic','帖子','content');
insert into permission (`name`,`description`,`pname`) values ('site','站点','content');
insert into permission (`name`,`description`,`pname`) values ('reply','评论','content');

insert into permission (`name`,`description`,`pname`) values ('system','系统管理','');
insert into permission (`name`,`description`,`pname`) values ('menu','菜单','system');	
insert into permission (`name`,`description`,`pname`) values ('role','角色','system');
insert into permission (`name`,`description`,`pname`) values ('category','分类','system');
insert into permission (`name`,`description`,`pname`) values ('submenu','子菜单','system');
insert into permission (`name`,`description`,`pname`) values ('user','用户','system');	
