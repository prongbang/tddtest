## Create Document

```
godoc path/package function
```

Example:
```
godoc tddtest/utils Reverse
```

## TDD Test

www.somkiat.cc/golang-basic-to-write-tests/

- การตั้งชื่อสำหรับการทดสอบก็สำคัญนะ

- ชื่อไฟล์ของการทดสอบจะลงท้ายด้วย _test.go เสมอ

- ตัวอย่างเช่นใน package hello มีไฟล์ชื่อว่า sample.go

- ถ้าต้องการทดสอบก็เขียนไว้ในไฟล์ sample_test.go
- ซึ่งอยู่ใน package hello เช่นกัน

- ส่วนชื่อของ test case นั้นต้องขึ้นต้นด้วยคำว่า Test เสมอ
- จากนั้นคำต่อมาต้องขึ้นต้นด้วยตัวพิมพ์ใหญ่เสมอนะ !!
- ยกตัวอย่างเช่น

- Test<ชื่อ function ที่ทดสอบ>ToReturn<ผลที่คาดหวัง><เงื่อนไขที่ต้องการทดสอบ>

- ในการทดสอบก็ไม่ยาก ใช้คำสั่งดังนี้ (ปกติจะทดสอบแบบ parallel อยู่แล้ว)

### ทำการทดสอบทั้งหมด
```
$go test ./…  
```
### ทำการทดสอบทีละ package
```
$go test -p 1  ./… 
``` 

### ทำการทดสอบตามชื่อ test case ที่ต้องการ
```
$go test -run TestName 
```
