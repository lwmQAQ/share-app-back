CREATE TABLE apps (
    id CHAR(36) PRIMARY KEY,               -- 唯一标识，使用 CHAR(36) 存储 UUID
    `name` VARCHAR(255) NOT NULL,             -- app 名字
    `type` INT NOT NULL,                     -- 类型 1 网页应用，2 桌面应用，3 博客
   `version` VARCHAR(50) NOT NULL,          -- 版本
    author VARCHAR(255) NOT NULL,          -- 作者
    released DATETIME NOT NULL,            -- 发布时间
    readme VARCHAR(255),                   -- 说明文档（存储 URL 地址）
    `description` TEXT,                      -- 描述
    icon VARCHAR(255),                     -- 图标 URL
    url VARCHAR(255)                       -- 下载地址
);

CREATE INDEX idx_type ON apps (TYPE);