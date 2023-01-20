class Node {
  constructor(value, left = null, right = null) {
    this.value = value;
    this.left = left;
    this.right = right;
  }
}

function treeByLevels(rootNode) {
  if (rootNode === null) {
    return [];
  }
  let list = [];
  let node = rootNode;
  list.push(node.value);
  processNode([node], list);

  return list;
}

function processNode(nodes, list) {
  let children = [];
  for (const node of nodes) {
    if (node.left !== null) {
      children.push(node.left);
      list.push(node.left.value);
    }
    if (node.right !== null) {
      children.push(node.right);
      list.push(node.right.value);
    }
  }
  if (children.length > 0) {
    processNode(children, list);
  }
}

// [2,8,9,1,3,4,5]
const treeOne = new Node(
  2,
  new Node(8, new Node(1), new Node(3)),
  new Node(9, new Node(4), new Node(5))
);

// [1,8,4,3,5,7]
const treeTwo = new Node(
  1,
  new Node(8, null, new Node(3)),
  new Node(4, null, new Node(5, null, new Node(7)))
);

console.log(treeByLevels(treeOne));
console.log(treeByLevels(treeTwo));
