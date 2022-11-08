-- dashboard
INSERT INTO menu VALUES (1000,'Admin',0, 'Dashboard','api-server',  '/dashboard',  0,   1,1,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);

-- 定时任务
INSERT INTO menu VALUES (2000,'Admin',0, '定时任务','api-server',  '/tasks',   0,   2,1,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);

-- 日志管理
INSERT INTO menu VALUES (3000,'Admin',0, '日志管理','api-server',  '/logs',   0,   3,1,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (3100,'Admin',0, '服务日志','api-server',  '/servicelog',3000,   1,1,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (3200,'Admin',0, '登录日志','api-server',  '/loginlog',3000,   2,1,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);

-- 系统设置
INSERT INTO menu VALUES (4000,'Admin',0, '系统设置','api-server',  '/setting',   0,   4,1,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (4100,'Admin',0, '菜单管理','api-server',  '/menu',4000,   2,1,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (4200,'Admin',0, '参数设置','api-server',  '/options',4000,   1,1,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);


-- INSERT INTO sys_job VALUES (1, '接口测试', 'DEFAULT', 1, '0/5 * * * * ', 'http://localhost:8000', '', 1, 1, 1, 0, '2021-05-13 19:56:37.914', '2021-06-14 20:59:55.417', NULL, 1, 1);
-- INSERT INTO sys_job VALUES (2, '函数测试', 'DEFAULT', 2, '0/5 * * * * ', 'ExamplesOne', '参数', 1, 1, 1, 0, '2021-05-13 19:56:37.914', '2021-05-31 23:55:37.221', NULL, 1, 1);


INSERT INTO user VALUES (1, 'admin', '$2a$10$.3KjCzbE./bI6nj5ayCL9.LkQL6pO.TmNCsbkjt1nmI4t0lGderZW', '管理员', '13333333333', 1, '', '1', '123@qq.com', '', '2', '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);









