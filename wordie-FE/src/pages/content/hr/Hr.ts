import { NewHrTask, TextTask } from "./Task";
import { TaskScheduler } from './TaskScheduler';

export class Hr {
  constructor(private el: Element) {
    this.fetchNewWords();
  }
  private async fetchNewWords() {
    const extractTextTask = new TextTask(this.el);
    const scheduler = new TaskScheduler()
    await scheduler.parseWhenIdle(extractTextTask);
    const text = extractTextTask.ret;
    console.log("extracted", text);
    const highlightTask = new NewHrTask(this.el, {
      words: ["the", "a"]
    });
    await scheduler.parseWhenIdle(highlightTask);
  }

}
