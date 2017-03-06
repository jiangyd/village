-- 添加联合字段约束
ALTER TABLE dz ADD UNIQUE KEY (type,type_id,uid_id);
ALTER TABLE collection ADD UNIQUE KEY (type,type_id,uid_id);