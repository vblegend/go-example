import { Rectangle, Vector2 } from "@hmi/editor/core/common";


export class HmiMath {

    private static readonly chars: string[] = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'];
    /**
     * 合并两个矩形区域返回一个新的矩形区域
     * @param rect1 
     * @param rect2 
     * @returns 
     */
    public static extendsRectangle(rect1: Rectangle | null, rect2: Rectangle | null): Rectangle {
        const left = Math.min(rect1!.left!, rect2!.left!);
        const top = Math.min(rect1!.top!, rect2!.top!);
        const width = Math.max(rect1!.left! + rect1!.width, rect2!.left! + rect2!.width) - left;
        const height = Math.max(rect1!.top! + rect1!.height, rect2!.top! + rect2!.height) - top;
        return { left, top, width, height };
    }

    public static checkRectangleCross(rect1: Rectangle | null, rect2: Rectangle | null): boolean {
        return rect1!.left! + rect1!.width >= rect2!.left! &&
            rect1!.left! <= rect2!.left! + rect2!.width &&
            rect1!.top! + rect1!.height >= rect2!.top! &&
            rect1!.top! <= rect2!.top! + rect2!.height;
    }


    public static randomString(count: number): string {
        let result = "";
        for (let i = 0; i < count; i++) {
            const index = Math.floor(Math.random() * HmiMath.chars.length);
            result += HmiMath.chars[index];
        }
        return result;
    }


    public static rotateVector2(point: Vector2, center: Vector2, a: number): Vector2 {
        const ang = a / 180 * Math.PI;
        const x = (point.x - center.x) * Math.cos(ang) - (point.y - center.y) * Math.sin(ang) + center.x;
        const y = (point.x - center.x) * Math.sin(ang) + (point.y - center.y) * Math.cos(ang) + center.y;
        return { x, y };
    }

}