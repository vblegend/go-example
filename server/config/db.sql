
-- INSERT INTO sys_job VALUES (1, '接口测试', 'DEFAULT', 1, '0/5 * * * * ', 'http://localhost:8000', '', 1, 1, 1, 0, '2021-05-13 19:56:37.914', '2021-06-14 20:59:55.417', NULL, 1, 1);
-- INSERT INTO sys_job VALUES (2, '函数测试', 'DEFAULT', 2, '0/5 * * * * ', 'ExamplesOne', '参数', 1, 1, 1, 0, '2021-05-13 19:56:37.914', '2021-05-31 23:55:37.221', NULL, 1, 1);


INSERT INTO user VALUES (1, 'admin', '$2a$10$.3KjCzbE./bI6nj5ayCL9.LkQL6pO.TmNCsbkjt1nmI4t0lGderZW', '管理员', '13333333333', 1, '', '1', '123@qq.com', '', '2', '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);





INSERT INTO menu VALUES (1000000 ,'Admin',0, '首页'    ,'grace-home',                  '/pages/dashboard',          0,  1,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);

INSERT INTO menu VALUES (2000000 ,'Admin',0, '主控面板','grace-hudun',                 '/pages/welcome/1',          0,  2,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (3000000 ,'Admin',0, '健康监控','grace-jiankang2',             '/pages/examples',           0,  3,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (4000000 ,'Admin',0, '服务进程','grace-memory',                '/pages/services',            0,  4,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (5000000 ,'Admin',0, '服务管理','grace-jiqunguanli',           '/pages/welcome/2',          0,  5,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (6000000 ,'Admin',0, '计划任务','grace-renwujihua',            '/pages/tasks',              0,  6,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);

INSERT INTO menu VALUES (7000000 ,'Admin',0, '数据管理','grace-shujuku',               '/pages/tasks',              0,  7,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (7001000 ,'Admin',0, '数据库'  ,'grace-shujuku',               '/pages/welcome/3',    7000000,  1,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (7002000 ,'Admin',0, '账号管理','grace-iconfontme',            '/pages/welcome/4',    7000000,  2,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (7003000 ,'Admin',0, '角色管理','grace-people',                '/pages/welcome/5',    7000000,  3,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (7004000 ,'Admin',0, '功能管理','grace-BIMfuneng',             '/pages/welcome/6',    7000000,  4,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (7005000 ,'Admin',0, '文件管理','grace-zu',                    '/pages/welcome/7',    7000000,  5,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);

INSERT INTO menu VALUES (8000000 ,'Admin',0, '工具箱'  ,'grace-Tools',                 '/pages/tasks',              0,  8,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (8001000 ,'Admin',0, '邮件'    ,'grace-youjian1',              '/pages/welcome/8',    8000000,  1,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (8002000 ,'Admin',0, '在线消息','grace-iconfontzaixiankefu1',  '/pages/welcome/9',    8000000,  2,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (8003000 ,'Admin',0, '广播消息','grace-news-fill',             '/pages/welcome/10',   8000000,  3,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (8004000 ,'Admin',0, '在线终端','grace-terminal-fill',         '/pages/welcome/11',   8000000,  4,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);

INSERT INTO menu VALUES (9000000 ,'Admin',0, '部署'    ,'grace-bushu',                 '/pages/logs',               0,  9,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);
INSERT INTO menu VALUES (9001000 ,'Admin',0, '服务部署','grace-iconfonticon-xitong',   '/pages/welcome/12',   9000000,  1,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);

INSERT INTO menu VALUES (10000000,'Admin',0, '关于我们','grace-about',                 '/pages/welcome/13',         0, 10,  1, 0,'2021-05-20 21:58:45.679+00:00','2022-06-09 10:48:50.6825313+08:00',NULL);




