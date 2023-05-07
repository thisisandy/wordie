function visibleTextDomWalker(node: Node) {
  return document.createTreeWalker(node, NodeFilter.SHOW_TEXT, {
    acceptNode: (node) => {
      const { display, visibility } = window.getComputedStyle(node.parentElement);
      const { width, height } = node.parentElement.getBoundingClientRect();
      if (display === "none" || visibility === "hidden" || width === 0 || height === 0) {
        return NodeFilter.FILTER_REJECT;
      }
      return NodeFilter.FILTER_ACCEPT;
    }
  });
}
export { visibleTextDomWalker };
