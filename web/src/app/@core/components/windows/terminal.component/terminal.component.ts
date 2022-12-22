import { Component, ElementRef, OnDestroy, OnInit, ViewChild, ComponentFactoryResolver, Injector } from '@angular/core';
import { Exception } from '@core/common/exception';
import { GenericComponent } from '@core/components/basic/generic.component';
import { NzModalRef, NzModalService } from 'ng-zorro-antd/modal'
import { Terminal } from 'xterm';
import { AttachAddon } from 'xterm-addon-attach';





@Component({
  selector: 'ngx-terminal-component',
  styleUrls: ['./terminal.component.less'],
  templateUrl: './terminal.component.html'
})
export class TerminalComponent extends GenericComponent {
  @ViewChild('TerminalParent', { static: true }) 
  public terminalDiv?: ElementRef;
  public term?: Terminal;
  private command?: string;
  constructor(injector: Injector, private modal: NzModalRef) {
    super(injector)

  }

  handleCancel(): void {
    this.modal.close(true);
  }

  protected onInit(): void {
    this.term = new Terminal({
      fontFamily: '"Cascadia Code", Menlo, monospace',
      cursorStyle: 'underline', // 光标样式
      cursorBlink: true, // 光标闪烁
      convertEol: true, // 启用时，光标将设置为下一行的开头
      disableStdin: false, // 是否应禁用输入。
      theme: {
        foreground: '#F8F8F8',
        background: '#2D2E2C',
        // selection: '#5DA5D533',
        black: '#1E1E1D',
        brightBlack: '#262625',
        red: '#CE5C5C',
        brightRed: '#FF7272',
        green: '#5BCC5B',
        brightGreen: '#72FF72',
        yellow: '#CCCC5B',
        brightYellow: '#FFFF72',
        blue: '#5D5DD3',
        brightBlue: '#7279FF',
        magenta: '#BC5ED1',
        brightMagenta: '#E572FF',
        cyan: '#5DA5D5',
        brightCyan: '#72F0FF',
        white: '#F8F8F8',
        brightWhite: '#FFFFFF',
        // foreground: 'yellow', //字体
        // background: '#060101', //背景色
        cursor: 'help' // 设置光标
      }
    });
    this.command = '';
    this.term.open(this.terminalDiv!.nativeElement);

    this.term.writeln('\x1b[35;1mPowered By \x1b[33;1mHanks.');
    this.term.writeln([
      ' ┌ \x1b[1mFeatures\x1b[0m ──────────────────────────────────────────────────────────────────┐',
      ' │                                                                            │',
      ' │  \x1b[31;1mApps just work                         \x1b[32mPerformance\x1b[0m                        │',
      ' │   Xterm.js works with most terminal      Xterm.js is fast and includes an  │',
      ' │   apps like bash, vim and tmux           optional \x1b[3mWebGL renderer\x1b[0m           │',
      ' │                                                                            │',
      ' │  \x1b[33;1mAccessible                             \x1b[34mSelf-contained\x1b[0m                     │',
      ' │   A screen reader mode is available      Zero external dependencies        │',
      ' │                                                                            │',
      ' │  \x1b[35;1mUnicode support                        \x1b[36mAnd much more...\x1b[0m                   │',
      ' │   Supports CJK 語 and emoji \u2764\ufe0f            \x1b[3mLinks\x1b[0m, \x1b[3mthemes\x1b[0m, \x1b[3maddons\x1b[0m, \x1b[3mtyped API\x1b[0m  │',
      ' │                                            ^ Try clicking italic text      │',
      ' │                                                                            │',
      ' └────────────────────────────────────────────────────────────────────────────┘',
      ''
    ].join('\n\r'));

    const socket = new WebSocket('wss://docker.example.com/containers/mycontainerid/attach/ws');
    socket.onerror = (e) => {
      throw Exception.fromCatch('Terminal does not work ', e, `Unable to connect to server：${socket.url}`);
    };
    const attachAddon = new AttachAddon(socket);
    // Attach the socket to term
    this.term.loadAddon(attachAddon);
    this.term.onData(this.onTermData.bind(this));
    this.prompt(this.term);
  }


  private onTermData(e: string): void {
    switch (e) {
      case '\u0003': // Ctrl+C
        this.term!.write('^C');
        this.prompt(this.term!);
        break;
      case '\r': // Enter
        this.runCommand(this.term!, this.command!);
        this.command = '';
        break;
      case '\u007F': // Backspace (DEL)
        // Do not delete the prompt
        if (this.term!.buffer.active.cursorX > 2) {
          this.term!.write('\b \b');
          if (this.command!.length > 0) {
            this.command = this.command!.substr(0, this.command!.length - 1);
          }
        }
        break;
      default: // Print all other characters for demo
        if (e >= String.fromCharCode(0x20) && e <= String.fromCharCode(0x7B) || e >= '\u00a0') {
          this.command += e;
          this.term!.write(e);
        }
    }
  }

  public runCommand(term: Terminal, text: string): void {
    const command = text.trim().split(' ')[0];
    if (command.length > 0) {
      term.writeln('');
      // if (command in commands) {
      //   commands[command].f();
      //   return;
      // }
      term.writeln(`${command}: command not found`);
    }
    this.prompt(term);
  }





  private prompt(term: Terminal): void {
    this.command = '';
    term.write('\r\n$ ');
  }

  public onDestroy(): void {

  }


}
