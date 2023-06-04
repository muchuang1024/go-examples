### 查看net/http 库所有的对外库函数

go doc net/http | grep "^func"

## 查看库提供的所有 struct

go doc net/http | grep "^type"|grep struct