-- 开始初始化数据 ;

INSERT INTO sys_menu VALUES (2, 'Admin', '系统管理', 'api-server', '/admin', '/0/2', 'M', '无', 0, 1, '', 'Layout', 90, '0', '1', 0, 1, '2021-05-20 21:58:45.679+00:00', '2022-06-09 10:48:50.6825313+08:00', NULL);
INSERT INTO sys_menu VALUES (58, 'Dict', '字典管理', 'education', '/admin/dict', '/0/2/58', 'C', '无', 2, 0, '', '/admin/dict/index', 60, '0', '1', 0, 1, '2021-05-20 22:08:44.526', '2021-06-17 11:48:40.703', '2022-07-05 13:01:37.1791604+08:00');
INSERT INTO sys_menu VALUES (459, 'ScheduleManage', '定时任务', 'time-range', '/schedule', '/0/459', 'C', '无',  0, 0, '', '/schedule/index', 45, '0', '1', 1, 1, '2020-08-03 09:17:37.000', '2021-06-05 22:15:03.465', NULL);
INSERT INTO sys_menu VALUES (537, 'SysTools', '系统工具', 'system-tools', '/sys-tools', '', 'M', '',  0, 0, '', 'Layout', 30, '0', '1', 1, 1, '2021-05-21 11:13:32.166', '2021-06-16 21:26:12.446', '2022-07-05 13:02:18.9628099+08:00');
INSERT INTO sys_menu VALUES (540, 'SysConfigSet', '参数设置', 'system-tools', '/admin/sys-config/set', '', 'C', '', 2, 0, '', '/admin/sys-config/set', 0, '0', '1', 1, 1, '2021-05-25 16:06:52.560', '2021-06-17 11:48:40.703', NULL);
INSERT INTO sys_menu VALUES (554, 'Index11', '本机监控', 'people', '/1', '', 'C', '', 0, 0, '', '/dashboard/index', 5, '0', '1', 1, 1, '2022-07-05 15:20:02.3556927+08:00', '2022-07-05 15:26:35.433933+08:00', '2022-07-05 15:38:30.534543+08:00');
INSERT INTO sys_menu VALUES (557, 'LogManage', '日志管理', 'log', '', '', 'M', '', 0, 0, '', 'Layout', 50, '0', '1', 1, 1, '2022-07-05 16:28:01.9283957+08:00', '2022-07-05 16:54:05.0357054+08:00', NULL);
INSERT INTO sys_menu VALUES (561, 'ServiceLog', '服务日志', 'druid', '/log/service', '', 'C', '', 557, 0, '', '/log/service', 0, '0', '1', 1, 0, '2022-07-05 16:50:31.2189392+08:00', '2022-07-05 16:50:31.2189392+08:00', NULL);

-- INSERT INTO sys_job VALUES (1, '接口测试', 'DEFAULT', 1, '0/5 * * * * ', 'http://localhost:8000', '', 1, 1, 1, 0, '2021-05-13 19:56:37.914', '2021-06-14 20:59:55.417', NULL, 1, 1);
-- INSERT INTO sys_job VALUES (2, '函数测试', 'DEFAULT', 2, '0/5 * * * * ', 'ExamplesOne', '参数', 1, 1, 1, 0, '2021-05-13 19:56:37.914', '2021-05-31 23:55:37.221', NULL, 1, 1);


INSERT INTO sys_user VALUES (1, 'admin', '$2a$10$.3KjCzbE./bI6nj5ayCL9.LkQL6pO.TmNCsbkjt1nmI4t0lGderZW', '管理员', '13333333333', 1, '', '1', '123@qq.com', '', '2', 1, 1, '2021-05-13 19:56:37.914', '2021-05-13 19:56:40.205', NULL);
