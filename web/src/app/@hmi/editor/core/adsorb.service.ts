import { Injectable } from "@angular/core";
import { Vector2 } from "@hmi/editor/core/common";
import { HmiEditorComponent } from "../hmi.editor.component";

@Injectable({
    providedIn: 'root'
})
/**
 * 编辑器的移动吸附服务
 */
export class AdsorbService {

    public version: number;
    /**
     *
     */
    constructor(private editor: HmiEditorComponent) {
        this.version = Math.random() * Number.MAX_VALUE;

        this.axisX = [];
        this.axisY = [];
        this.enabled = true;

    }



    /**
     * 水平坐标
     */
    private axisX: number[];


    /**
     * 垂直坐标
     */
    private axisY: number[];

    /**
     * 
     */
    public enabled: boolean;







    /**
     * 捕获所有未选中对象的坐标点
     */
    public captureAnchors(): void {
        if (!this.enabled) return;
        this.axisX.length = 0;
        this.axisY.length = 0;
        const selection = this.editor.selection;
        for (const comp of this.editor.canvas.children) {
            if (!selection.contains(comp)) {
                this.axisX.push(comp.instance.left!);
                this.axisX.push(Math.floor(comp.instance.left! + comp.instance.width! / 2));
                this.axisX.push(comp.instance.left! + comp.instance.width!);
                this.axisY.push(comp.instance.top!);
                this.axisY.push(Math.floor(comp.instance.top! + comp.instance.height! / 2));
                this.axisY.push(comp.instance.top! + comp.instance.height!);
            }
        }
        this.axisX = Array.from(new Set(this.axisX)).sort((a, b) => a - b);
        this.axisY = Array.from(new Set(this.axisY)).sort((a, b) => a - b);
    }



    /**
     * 
     */
    public clear(): void {
        this.axisX.length = 0;
        this.axisY.length = 0;
    }

    /**
     * 二分查找近似值
     * @param datas 数据源
     * @param value 查找的近似值
     */
    private binarySearch(array: number[], targetNum: number): number {
        let min = 0;
        let max = array.length - 1;
        while (min != max) {
            const midIndex = Math.round((max + min) / 2);
            const mid = max - min;
            if (targetNum === array[midIndex]) {
                return array[midIndex];
            }
            if (targetNum > array[midIndex]) {
                min = midIndex;
            } else {
                max = midIndex;
            }
            if (mid < 2) {
                break;
            }
        }
        if (Math.abs(array[max] - targetNum) >= Math.abs(array[min] - targetNum)) {
            return array[min];
        }
        else {
            return array[max];
        }
    }


    /**
     * 匹配X轴坐标点查找离该点最近点\
     * 查询成功返回离得最近的x轴坐标点，查询失败返回null
     * @param x 提供一个X轴坐标点
     * @param threshold 查询阈值
     * @returns x轴点
     */
    public matchXAxis(x: number, threshold = 5): number | null {
        const result = this.binarySearch(this.axisX, x);
        if (result && Math.abs(result - x) < threshold) {
            return result;
        }
        return null;
    }

    /**
     * 匹配Y轴坐标点查找离该点最近点
     * 查询成功返回离得最近的y轴坐标点，查询失败返回null
     * @param y 提供一个y轴坐标点
     * @param threshold 查询阈值
     * @returns y轴点
     */
    public matchYAxis(y: number, threshold = 5): number | null {
        const result = this.binarySearch(this.axisY, y);
        if (result && Math.abs(result - y) < threshold) {
            return result;
        }
        return null;
    }


    /**
     * 吸附一个近似值   
     * @param in_out_Point  要查找的坐标，最后返回
     * @param threshold     阈值
     * @param return        Vector2 (x,y) 返回 x,y轴修正的范围 当坐标值为null时 未吸附到值
     */
    public match(in_out_Point: Vector2, threshold = 15): Vector2 {
        const result: Vector2 = { x: 0, y: 0 };
        // lessValue *= this.designer.res;
        const x = this.binarySearch(this.axisX, in_out_Point.x);
        const y = this.binarySearch(this.axisY, in_out_Point.y);
        const x_dis = Math.abs(x - in_out_Point.x);
        if (x != null && x_dis <= threshold) {
            in_out_Point.x = x;
            result.x = x_dis;
        }
        const y_dis = Math.abs(y - in_out_Point.y);
        if (y != null && y_dis <= threshold) {
            in_out_Point.y = y;
            result.y = y_dis;
        }
        return result;
    }


}