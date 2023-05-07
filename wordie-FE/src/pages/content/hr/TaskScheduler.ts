import { Task } from "./Task"


export class TaskScheduler {
  async parseWhenIdle(task: Task): Promise<string> {
    // request idle time and parse the document content and extract all of the words
    // and then send the words to the server to get the word frequency
    return new Promise((resolve, reject) => {
      window.requestIdleCallback(async (deadline) => {
        try {
          while (deadline.timeRemaining() > 0) {
            if (!task.finished) {
              await task.run()
            } else {
              resolve("success")
            }
          }
          if (!task.finished) {
            await this.parseWhenIdle(task)
            resolve("success")
          }
        } catch (e) {
          reject(e)
        }
      })
    })
  }
}

