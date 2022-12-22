import { verifyDocument } from "@hmi/configuration/global.default.configure";
import { Exception } from "./exception";




export class WebWorker {
    private LoadData(data: number): void {
        const err = Exception.build("aaa","bbbbb",100);
        console.log(err);

        let v = 0;
        for (let i = 0; i < 10000; i++) {
            v+= i;
        }
        self.postMessage(data+v);
    }



    public Test(): void {
        this.Invoke(this.LoadData);
    }


    public Invoke(workFunc: (input: any) => void): void {
        const funcSrc = this.GetFunctionLmd(workFunc);
        const userScript = `${funcSrc}\nself.onmessage = function (event) { _MAIN_(event.data); };`;
        const url = `data:application/javascript,${encodeURIComponent(userScript)}`;
        const worker = new Worker(url, { type: 'module' });

        worker.onmessage = (event) => {
            // 数据打印
            console.log(event.data);
        };
        worker.onerror = (err) => {
            console.error(err);
        }
        // 向 worker 线程发送数据
        worker.postMessage(555555);
        // worker.terminate()
    }

    private GetFunctionLmd(workFunc: (input: any) => void): string {
        const origin = workFunc.toString()
        const codeSegment = origin.substring(origin.indexOf("{") + 1, origin.lastIndexOf("}"));
        const args = origin.substring(origin.indexOf("(") + 1, origin.indexOf(")"));
        return `function _MAIN_(${args}) {${codeSegment}}`;
    }
}
