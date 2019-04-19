# Go Map to Struct

---
### English

This package is simple Map to Struct function.

Can reflect :
    int     int8        int16       int32       int64
    uint    uint8       uint16      uint32      uint64
    bool    float32     float64     string

This function can doing:

1. Map to struct.
2. Map slice to struct slice.
3. Map to struct can custom map key to struct key.
    
    EX.
``` 
    type A struct{
        number string
    }

    Map[string]string{numberCode:"12345678"}
```

    then struct key not map key,
    you can do  1. fuck off ,manual input data.
                2. add config MapToSturctName list.

4. Map to struct can add convert function.
    EX. a:10.010 -> a:10.01.

#### You can see Test folder Example.

---

### Chinese

這是一著簡單的Map轉換Struct的功能。

可以反射：
    int     int8        int16       int32       int64
    uint    uint8       uint16      uint32      uint64
    bool    float32     float64     string

這個function能做到以下功能：

1. Map 轉換成 Struct。
2. Map 數組 轉換成 Struct 數組。
3. Map 轉換成 Struct 時能自訂Key的轉換。
   
   例如：
```
   type A struct{
        number string
    }

    Map[string]string{numberCode:"12345678"}
```

    這時候結構的Key跟map的key不一樣
    這時可以做以下  1. 去你的，老子想手動輸入。
                   1. 使用在config資料中的表單。 

4. Map 轉換成 Struct 時能給予轉換功能。
    例如：a:10.010 -> a:10.01.

#### 可以看Test資料夾的資料。

-----------------------------------------------------------
