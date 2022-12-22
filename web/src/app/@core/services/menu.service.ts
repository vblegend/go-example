import { Injectable, Injector } from "@angular/core";
import { RouteConfigure } from "../models/route.configure";
import { RestfulService } from "./restful.service";


@Injectable({
    providedIn: 'root',
})
export class MenuService extends RestfulService {

    public isCollapsed: boolean = false;

    /**
     *
     */
    protected onInit(): void {
        this.isCollapsed = false;
    }

    public toggleCollapsed(): void {
        this.isCollapsed = !this.isCollapsed;
    }


    public async load(): Promise<void> {
        const res = await this.get<RouteConfigure[]>("/api/menutree");
        this.menus = res.data
    }



    public menus: RouteConfigure[] = [
        {
            title: '首页',
            icon: 'grace-home',
            path: 'welcome/1',
        },
        {
            title: '主控面板',
            icon: 'grace-hudun',
            path: 'dashboard',
        },
        {
            title: '健康监控',
            icon: 'grace-jiankang2',
            path: 'examples',
        },
        {
            title: '服务进程',
            icon: 'grace-memory',
            path: 'ifame',
        },
        {
            title: '服务管理',
            icon: 'grace-jiqunguanli',
            path: 'ifame',
        },
        {
            title: '计划任务',
            icon: 'grace-renwujihua',
            path: 'tasks',
        },
        {
            title: '日志管理',
            icon: 'grace-order',
            opened: false,
            selected: false,
            children: [
                {
                    title: '服务日志',
                    icon: 'grace-wj-rz',
                    path: 'welcome/6',
                },
                {
                    title: '聊天日志',
                    icon: 'grace-order',
                    path: 'welcome/7',
                }
            ]
        },
        {
            title: '数据管理',
            icon: 'grace-shujuku',
            opened: false,
            selected: false,
            children: [
                {
                    title: '数据库',
                    icon: 'grace-shujuku',
                    path: 'welcome/9',
                },
                {
                    title: '账号管理',
                    icon: 'grace-iconfontme',
                    path: 'welcome/78489',
                },
                {
                    title: '角色管理',
                    icon: 'grace-people',
                    path: 'welcome/45456',
                },
                {
                    title: '功能管理',
                    icon: 'grace-BIMfuneng',
                    path: 'welcome/775',
                },
                {
                    title: '文件管理',
                    icon: 'grace-zu',
                    path: 'welcome/454',
                }
            ]
        },
        {
            title: '工具箱',
            icon: 'grace-Tools',
            opened: false,
            selected: false,
            children: [
                {
                    title: '邮件',
                    icon: 'grace-youjian1',
                    path: 'welcome/45342',
                },
                {
                    title: '在线消息',
                    icon: 'grace-iconfontzaixiankefu1',
                    path: 'welcome/4542',
                },
                {
                    title: '广播消息',
                    icon: 'grace-news-fill',
                    path: 'welcome/78524',
                },
                {
                    title: '在线终端',
                    icon: 'grace-terminal-fill',
                    path: 'welcome/545',
                },
            ]
        },
        {
            title: '部署',
            icon: 'grace-bushu',
            opened: false,
            selected: false,
            children: [
                {
                    title: '服务部署',
                    icon: 'grace-iconfonticon-xitong',
                    path: 'welcome/3455',
                }
            ]
        },
        {
            title: '关于我们',
            icon: 'grace-about',
            path: 'welcome/4343213',
        }
    ];


}