package main

import (
	"fmt"
	"tddtest/utils"
	"time"
)

// 1. สร้างโครงสร้างข้อมูลชื่อว่า Data
type Data struct {
	Line string
}

// 2. สร้างส่วนการดึงข้อมูลจากอุปกรณ์
type Device struct {
	Host    string
	Timeout time.Duration
}

func (*Device) Pull(d *Data) error {

	return nil
}

// 3. สร้างส่วนการจัดเก็บข้อมูลจากอุปกรณ์
type Storage struct {
	Host    string
	Timeout time.Duration
}

func (*Storage) Store(d *Data) error {

	return nil
}

// 4. สร้างส่วน logic ของการ pull ข้อมูลจากอุปกรณ์จริง ๆ
func pull(d *Device, data []Data) (int, error) {
	for i := range data {
		if err := d.Pull(&data[i]); err != nil {
			return i, err // Found error
		}
	}
	return len(data), nil
}

// 5. สร้างส่วน logic ของการ store ข้อมูล
func store(s *Storage, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err // Found error
		}
	}
	return len(data), nil
}

// 6. สิ่งที่ยังขาดไปคือ ส่วนการทำงานหลัก สำหรับดึงและบันทึกข้อมูล
type System struct {
	Device
	Storage
}

func Copy(sys *System, batch int) error {
	data := make([]Data, batch)
	for {
		i, _ := pull(&sys.Device, data)
		store(&sys.Storage, data[:i])
	}
}

// จากตัวอย่าง code นั้นพบว่า ทำงานได้อย่างดี
// แต่ปัญหาที่น่าสนใจ หรือ Code Smell คือ
// ในตอนนี้ Data กับ Behavior ของ Device และ Storage รวมกันอยู่
// นั่นคือผูกมัดกันอย่างมาก

// ลองคิดดูสิว่า ถ้ามีจำนวน Device และ Storage จำนวนมาก
// ซึ่งมีพฤติกรรมการทำงานที่แตกต่างกัน
// ดังนั้นจึงต้องหาวิธีแก้ไขปัญหา ?

// ขั้นตอนที่ 2 Decoupling with Interface

// แยกพฤติกรรมออกมาจาก Device และ Storage ดังนี้
type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
}

// แก้ไข code ในส่วนของ method ที่ทำงาน pull() และ store() ข้อมูล
// จากเดิมที่ส่ง struct เข้าไป ก็เปลี่ยนเป็น interface ได้ดังนี้

func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := d.Pull(&data[i]); err != nil {
			return i, err // Found error
		}
	}
	return len(data), nil
}

func store(s Storer, data []Data) (int, error) {
	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err // Found error
		}
	}
	return len(data), nil
}

// ขั้นตอนที่ 3 Interface Composition

type PullStorer interface {
	Puller
	Storer
}

func Copy(ps PullStorer, batch int) error {
	data := make([]Data, batch)
	for {
		i, _ := pull(ps, data)
		store(ps, data[:i])
	}
}

// ขั้นตอนที่ 4 Decoupling with Interface composition
// แยกออกมาด้วยการใช้งาน interface composition แทน

type System struct {
	Puller
	Storer
}

func main() {
	sys := System{
		Puller: &Device{},
		Storer: &Storage{},
	}

	Copy(&sys, 3)

	fmt.Println(utils.Reverse("Hello"))
}
