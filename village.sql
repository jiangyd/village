-- 添加联合字段约束
ALTER TABLE dz ADD UNIQUE KEY (type,type_id,uid_id);
ALTER TABLE collection ADD UNIQUE KEY (type,type_id,uid_id);
ALTER TABLE firend ADD UNIQUE KEY (user_a_id,user_b_id);