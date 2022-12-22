// import { Injectable, TemplateRef } from "@angular/core";

// @Injectable({
//     providedIn: 'root',
// })
// export class ContextMenuService {

//     constructor() {

//         document.oncontextmenu = (e: MouseEvent) => {
//             const element = e.target as HTMLElement;
//             if (element instanceof HTMLInputElement) return;
//             const selection = document.getSelection().toString();
//             if (selection === '') e.preventDefault();
//         }
//     }
// }