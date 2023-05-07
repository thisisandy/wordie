import { visibleTextDomWalker } from "./Walker";


/**
 * Abstract class representing a task to be performed on a DOM node.
 */
export abstract class Task {
  abstract run(): Promise<void>;
  abstract get finished(): boolean;
  public ret: any;
}

/**
 * Abstract class representing a task that involves walking through a DOM tree.
 */
abstract class NodeWalkerTask extends Task {
  abstract run(): Promise<void>;
  protected walk:TreeWalker;
  protected queue:Node[] = [];

  constructor(private node: Node) {
    super();
    this.walk = visibleTextDomWalker(this.node);
  }

  /**
   * Returns the next node in the queue.
   * If the queue is empty, it enqueues the next node.
   */
  get nextNode() {
    if (this.queue.length === 0) {
      this.enqueueNextNode();
    }
    return this.queue.shift();
  }

  /**
   * Enqueues the next node in the queue.
   */
  private enqueueNextNode(): void {
    if(this.walk.nextNode()) {
      this.queue.push(this.walk.currentNode);
    }
  }

  /**
   * Returns true if the task is finished.
   * If the queue is empty, it enqueues the next node.
   */
  get finished(){
    if (this.queue.length === 0) {
      this.enqueueNextNode();
    }
    return this.queue.length === 0;
  }

}

/**
 * Class representing a task to extract text from a DOM node.
 */
export class TextTask extends NodeWalkerTask {
  public ret = "";

  constructor(node: Node) {
    super(node);
  }

  async run() {
      const currentNode = this.nextNode;
      this.ret += currentNode.textContent;
  }
}

interface Options {
  className?: string;
  style?: string;
  words: string[];
  Tag?: string;
}

/**
 * Class representing a task to wrap words in a DOM node with a specified tag.
 */
export class HrTask extends NodeWalkerTask {
  constructor(node: Node, private options: Options) {
    super(node);
  }

  async run() {
    const currentNode = this.nextNode;
    const nodeValue = currentNode.nodeValue;
    const parentNode = currentNode.parentNode;
  
    // Create a DocumentFragment to hold the new nodes
    const fragment = document.createDocumentFragment();
  
    let lastIndex = 0;
    for (const word of this.options.words) {
      const regex = new RegExp(`\\b${word}\\b`, 'g');
      let match;

      // Iterate over all matches of the word in the text content
      while ((match = regex.exec(nodeValue)) !== null) {
        // Create a text node for the text before the matched word
        const beforeWord = document.createTextNode(nodeValue.substring(lastIndex, match.index));
        fragment.appendChild(beforeWord);
  
        // Create the wrapping element for the matched word
        const wrappedWord = document.createElement(this.options.Tag || "span");
        wrappedWord.className = this.options.className || "";
        wrappedWord.style.cssText = this.options.style || "";
        wrappedWord.textContent = word;
        fragment.appendChild(wrappedWord);
  
        lastIndex = match.index + word.length;
      }
    }
  
    // Create a text node for the remaining text after the last matched word
    const afterWords = document.createTextNode(nodeValue.substring(lastIndex));
    fragment.appendChild(afterWords);
  
    // Replace the original text node with the new nodes in the DocumentFragment
    parentNode.replaceChild(fragment, currentNode);
  }
  
}

export class NewHrTask extends NodeWalkerTask {
  private highlightedRanges: Range[] = [];
  private highlightRegistry: Highlight =  new Highlight()
  constructor(node: Node, private options: Options) {
    super(node);
    CSS.highlights.set("new-word", this.highlightRegistry);
  }
  async run() {
    const currentNode = this.nextNode;
    const nodeValue = currentNode.textContent;
    for (const word of this.options.words) {
      const regex = new RegExp(`\\b${word}\\b`, 'g');
      let match;
      // Iterate over all matches of the word in the text content
      while ((match = regex.exec(nodeValue)) !== null) {
        // Create a range for the matched word
        const range = document.createRange();
        range.setStart(currentNode, match.index);
        range.setEnd(currentNode, match.index + word.length);
        this.highlightedRanges.push(range);
        this.highlightRegistry.add(range);
        // Do something with the range, e.g., replace the word with a wrapped element
      }
    }
    console.log(this.highlightedRanges);
  }
}
  
