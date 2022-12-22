import { Component, ElementRef, HostListener, Injector, Input, TemplateRef, ViewChild } from '@angular/core';
import { TimerTask, TimeState } from '@core/common/timer.task';
import { nzSelectItem } from '@core/common/types';
import { GenericComponent } from '@core/components/basic/generic.component';
import { FileUtil } from '@core/util/file.util';
import { GraphicConfigure, GraphicOptions } from '@hmi/configuration/graphic.configure';
import { HmiZoomMode } from '@hmi/editor/core/common';
import { BatchCommand } from '@hmi/editor/commands/batch.command';
import { GenericAttributeCommand } from '@hmi/editor/commands/generic.attribute.command';
import { HmiEditorComponent } from '@hmi/editor/hmi.editor.component';
import { HmiEditorService } from '@hmi/services/hmi.editor.service';
import { Subject } from 'rxjs';

@Component({
  selector: 'hmi-editor-toolbar',
  styleUrls: ['./toolbar.component.less'],
  templateUrl: './toolbar.component.html',
})
export class EditorToolbarComponent extends GenericComponent {
  public toolbarState: boolean = false;
  public settingVisible: boolean = false;
  public readonly graphicOptions: GraphicOptions;
  public readonly scaleModels: nzSelectItem[] = [];
  public readonly onSave: Subject<GraphicConfigure>;

  constructor(injector: Injector, public editor: HmiEditorComponent, private hmiEditorService: HmiEditorService) {
    super(injector);
    this.onSave = new Subject<GraphicConfigure>();
    this.scaleModels.push({ label: '不缩放', value: 0 });
    this.scaleModels.push({ label: '保持比例', value: 1 });
    this.scaleModels.push({ label: '拉伸填充', value: 2 });
    this.graphicOptions = { width: 0, height: 0, zoomMode: HmiZoomMode.None };
  }


  public fullscreen_click(): void {
    if (this.isFullScreen) {
      document.exitFullscreen();
    } else {
      document.documentElement.requestFullscreen();
    }
  }

  public get fullScreenIcon(): string {
    return document.body.scrollHeight === window.screen.height &&
      document.body.scrollWidth === window.screen.width ?
      'icon-hide-sidebar' : 'icon-full-screen';
  }


  public get isFullScreen(): boolean {
    return document.body.scrollHeight === window.screen.height && document.body.scrollWidth === window.screen.width;
  }


  @HostListener('mouseenter', ['$event'])
  public onMouseEnter(event: MouseEvent): void {
    if (event.buttons != 0) return;

    this.toolbarState = true;
    if (this.animateTimer && this.animateTimer.state == TimeState.Runing) {
      this.animateTimer.cancel();
    }
  }

  @HostListener('mouseleave@outside', ['$event'])
  public onMouseLeave(event: MouseEvent): void {
    this.animateTimer = this.timeout(0.3, () => {
      this.run(() => { this.toolbarState = false; });
    });
  }
  private animateTimer!: TimerTask;

  public save_click(): void {
    const json = this.editor.toJson();
    this.editor.onSave.next(json);
  }




  public async setting_click(): Promise<void> {
    this.graphicOptions.width = this.editor.width;
    this.graphicOptions.height = this.editor.height;
    this.graphicOptions.zoomMode = this.editor.zoomMode;
    this.settingVisible = true;
  }


  public saveSetting(): void {
    this.editor.execute(new BatchCommand(this.editor,
      new GenericAttributeCommand(this.editor, [this.editor], 'width', [this.graphicOptions.width]),
      new GenericAttributeCommand(this.editor, [this.editor], 'height', [this.graphicOptions.height]),
      new GenericAttributeCommand(this.editor, [this.editor], 'zoomMode', [this.graphicOptions.zoomMode])
    ));
    this.settingVisible = false;
  }


  public closeSetting(): void {
    this.settingVisible = false;
  }

  /**
   * 增量导入
   */
  public async addimport_click(): Promise<void> {
    const file = await FileUtil.selectLocalFile('application/json');
    if (file && file.length == 1) {
      const json = await FileUtil.loadJsonFromFile<GraphicConfigure>(file[0]);
      try {
        this.editor.appendFromJson(json);
        this.notification.success('提示', '导入成功');
      } catch (ex) {
        this.notification.error('导入失败', `${ex}`);
      }
    }
  }

  /**
   * 覆盖导入
   */
  public async import_click(): Promise<void> {
    const file = await FileUtil.selectLocalFile('application/json');
    if (file && file.length == 1) {
      const json = await FileUtil.loadJsonFromFile<GraphicConfigure>(file[0]);
      try {
        this.editor.loadFromJson(json);
        this.notification.success('提示', '导入成功');
      } catch (ex) {
        this.notification.error('导入失败', `${ex}`);
      }
    }
  }

  public export_click(): void {
    const json = this.editor.toJson();
    FileUtil.download(json, 'widgets.json');
  }

  public close_click(): void {
    this.hmiEditorService.close();
  }




  public retracted_click(): void {
    this.toolbarState = !this.toolbarState;
  }


}
