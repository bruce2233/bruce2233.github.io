---
title: Python
date: 202310-11
author: 岛石  
categories: 
- "Python"
- "程序"
tag: 
- "并发"
---

# multiprocessing 模块
## 使用方法
```python
from multiprocessing import Pool
worker_num = 10
p = Pool(worker_num)

def workder(worker_id):
    ......

for i in range(workder_num):
    p.apply_async(worker, args=(i,), error_callback=lamba err: print(err))
```