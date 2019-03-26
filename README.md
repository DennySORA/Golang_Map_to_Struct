# Go Map to Struct

---
### English

This package is simple Map to Struct function.

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

#### Bug:

- Input Only map[string]string and struct all data is string.
  Because this function is design to SQL get data to GraphQL struct.
  SQL get data is all string.

#### You can see Test folder Example.

---

### Chinese

這是一著簡單的Map轉換Struct的功能。

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

#### Bug:

- 只能使用在map[string]string 跟 結構中輸入的要是『字串』。
  因為這原本是設計時是為了SQL取完數據之後轉到GraphQL結構用的。
  而SQL取出來都是『字串』.

#### 可以看Test資料夾的資料。

-----------------------------------------------------------
