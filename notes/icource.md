# [数据结构与算法笔记](https://www.icourse163.org/learn/ZJU-93001?tid=1459700443#/learn/content?type=detail&id=1235254043&sm=1)

## 注意事项
1. 递归应关注可能出现**内存溢出**；
2.

## 知识点

### 树
1. 特殊树
![](../img/3.2-01.png)

2. 二叉搜索树(`BST`)满足如下性质：
- 左子树所有节点的元素值均小于根的元素值；
- 右子树所有节点的元素值均大于根的元素值。\

2. 1. 二叉搜索树的良好形态：平衡二叉树
    > ![](../img/4.2-01.png)

平衡二叉树的再平衡代码：java
```java
   /**
     * 向以node为根的二分搜索树中插入元素(key, value)，递归算法
     * 时间复杂度 O(log n)
     * @param node
     * @param key
     * @param value
     * @return 返回插入新节点后二分搜索树的根
     */
    private Node add(Node node, K key, V value){

        if(node == null){
            size ++;
            return new Node(key, value);
        }

        if(key.compareTo(node.key) < 0)
            node.left = add(node.left, key, value);
        else if(key.compareTo(node.key) > 0)
            node.right = add(node.right, key, value);
        else // key.compareTo(node.key) == 0
            node.value = value;

        // 更新height
        node.height = 1 + Math.max(getHeight(node.left), getHeight(node.right));

        // 计算平衡因子
        int balanceFactor = getBalanceFactor(node);

        // 平衡维护
        //////////////////////////////////////////////////////
        // LL  T1<Z<T2< X <T3<Y<T4                          //
        //        y                              x          //
        //       / \                           /   \        //
        //      x   T4     向右旋转 (y)        z     y       //
        //     / \       - - - - - - - ->    / \   / \      //
        //    z   T3                        T1 T2 T3 T4     //
        //   / \                                            //
        // T1   T2                                          //
        //////////////////////////////////////////////////////
        if (balanceFactor > 1 && getBalanceFactor(node.left) >= 0)
            return rightRotate(node);
        //////////////////////////////////////////////////////////////////////////////////////////
        //  LR  T1<X<T2< Z <T3<Y<T4                                                             //
        //         y                                y                              z            //
        //        / \                              / \                           /   \          //
        //       x  t4    向左旋转（x）             z   T4      向右旋转（y）       x     y         //
        //      / \     --------------->         / \        --------------->   / \   / \        //
        //     T1  z                            x   T3                        T1  T2 T3 T4      //
        //        / \                          / \                                              //
        //       T2  T3                      T1   T2                                            //
        //////////////////////////////////////////////////////////////////////////////////////////
        if (balanceFactor > 1 && getBalanceFactor(node.left) < 0) {
            node.left = leftRotate(node.left);
            return rightRotate(node);
        }
        //////////////////////////////////////////////////
        // RR: T1<Y<T2< X <T3<Z<T4                      //
        //    y                              x          //
        //  /  \                           /   \        //
        // T1   x      向左旋转 (y)        y     z       //
        //     / \   - - - - - - - ->    / \   / \      //
        //   T2   z                     T1 T2 T3 T4     //
        //       / \                                    //
        //      T3 T4                                   //
        //////////////////////////////////////////////////
        if (balanceFactor < -1 && getBalanceFactor(node.right) <= 0)
            return leftRotate(node);

        //////////////////////////////////////////////////////////////////////////////////////////
        // RL: T1<Y<T2< Z <T3<X<T4                                                              //
        //      y                           y                                       z           //
        //     / \                         / \                                    /   \         //
        //    T1  x       向右旋转（x）     T1  z         向左旋转（y）              y     x        //
        //       / \    - - - - - - ->       / \      - - - - - - - - ->        / \   / \       //
        //      z  T4                       T2  x                              T1 T2 T3 T4      //
        //     / \                             / \                                              //
        //    T2  T3                          T3  T4                                            //
        //////////////////////////////////////////////////////////////////////////////////////////
        if (balanceFactor < -1 && getBalanceFactor(node.right) > 0) {
            node.right = rightRotate(node.right);
            return leftRotate(node);
        }

        return node;
    }
    
    /**
     * 获得节点node的高度
     * @param node
     * @return
     */
    private int getHeight(Node node){
    if(node == null){
        return 0;
    }
    int i = getHeight(node.left);
    int j = getHeight(node.right);
    return (i<j)? j+1:i+1;
    }

    /**
     * 获得节点node的平衡因子
     * @param node
     * @return
     */
    private int getBalanceFactor(Node node){
        if(node == null)
            return 0;
        return getHeight(node.left) - getHeight(node.right);
    }
```

3. [堆](https://blog.csdn.net/qq_37044026/article/details/86714130) ：是计算机科学中一类特殊的数据结构的统称。\
   堆通常是一个可以被看做一棵树的**数组对象**。\
   堆总是满足下列性质：
   - 任意节点小于（或大于）它的所有后裔，最小元（或最大元）在堆的根上（堆序性）。
   - 堆总是一棵**完全树**。即除了最底层，其他层的节点都被元素填满，且最底层尽可能地从左到右填入。
   - 数组0索引值为允许的最大或最小值，起边界作用，同时也起占位作用

4. 哈夫曼树：\
![](../img/5.2-1.png) \
   ![](../img/5.2-2.png)
   哈夫曼树的构建思路为 
   1. 从小到大对给定的数据进行排序；
   2. 取出权值最小的两个数据构成二叉树，两数之和为这两个数据的父节点；
   3. 再将两数之和放入数据中再排序，不断重复1-2-3步骤，直到数据中，所有数据都被处理，就得到了哈夫曼树。

5. 集合 \
路径压缩：
   ![](../img/5.小白专场%5B陈越%5D.png)
   
### 图

1. 图的抽象定义
![](../img/6.1-1.png)
