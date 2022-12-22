



export class Guid {
    private static chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'.split('');


    /**
     * 生成一个标准长度的guid \
     * xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
     * @returns 
     */
    public static new(): string {
        return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
            const r = Math.random() * 16 | 0;
            const v = c == 'x' ? r : (r & 0x3 | 0x8);
            return v.toString(16);
        });
    }


    /**
     * 生成一个指定长度指定基数的唯一ID
     * @param len 生成的长度
     * @param base 进制基数  2、10、16、62
     * @returns 
     */
    public static custom(len: number, base: number): string {
        const uuid = [];
        base = base || this.chars.length;
        for (let i = 0; i < len; i++) {
            uuid[i] = this.chars[0 | Math.random() * base];
        }
        return uuid.join('');
    }




}
