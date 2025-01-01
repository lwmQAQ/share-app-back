-- 创建数据库
CREATE DATABASE TaskDB;

-- 使用数据库
USE TaskDB;

-- 创建任务表
CREATE TABLE tasks (
    taskid INT AUTO_INCREMENT PRIMARY KEY, -- 主键，自动递增
    user_id BIGINT(20) UNSIGNED NOT NULL,     --用户id
    source_url VARCHAR(255) NOT NULL,       -- 来源 URL，不允许为空
    file_name VARCHAR(255) NOT NULL,   --文件名
    mono_url VARCHAR(255),                  -- 单声道 URL，可以为空
    dual_url VARCHAR(255),                   -- 双声道 URL，可以为空
    status INT NOT NULL DEFAULT 0,         -- 任务状态 0 -未完成 -1 已完成  -2失败
    INDEX idx_userid (user_id)              -- 给userid添加索引
);
