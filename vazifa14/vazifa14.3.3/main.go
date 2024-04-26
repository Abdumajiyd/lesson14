package main

import (
	"fmt"
	"time"
)
type Tadbir struct {
	Nomi      string
	BoshlanishVaqt time.Time
}

type Foydalanuvchi struct {
	Ismi   string
	Tadbirlar []Tadbir
}

func (f *Foydalanuvchi) TadbirQoshish(tadbirNomi string, boshlanishVaqt time.Time) {
	tadbir := Tadbir{Nomi: tadbirNomi, BoshlanishVaqt: boshlanishVaqt}
	f.Tadbirlar = append(f.Tadbirlar, tadbir)
}

func (f *Foydalanuvchi) TadbirOchirish(tadbirNomi string) {
	var yangilanganTadbirlar []Tadbir
	for _, t := range f.Tadbirlar {
		if t.Nomi != tadbirNomi {
			yangilanganTadbirlar = append(yangilanganTadbirlar, t)
		}
	}
	f.Tadbirlar = yangilanganTadbirlar
}

func (f *Foydalanuvchi) TadbirlarniChiqarish() {
	fmt.Printf("Foydalanuvchi: %s\n", f.Ismi)
	fmt.Println("Tadbirlar:")
	for _, t := range f.Tadbirlar {
		fmt.Printf("- %s, Boshlanish vaqti: %s\n", t.Nomi, t.BoshlanishVaqt.Format("2006-01-02 15:04:05"))
	}
}

func main() {
	foydalanuvchi := Foydalanuvchi{Ismi: "Husan"}
	bugun := time.Now()
	foydalanuvchi.TadbirQoshish("Tadbir 1", bugun)
	foydalanuvchi.TadbirQoshish("Dars 1", bugun.AddDate(0, 0, 1))

	foydalanuvchi.TadbirlarniChiqarish()

	foydalanuvchi.TadbirOchirish("Tadbir 1")
	fmt.Println("Tadbirlar o'chirildi:")

	foydalanuvchi.TadbirQoshish("Tadbir 2", bugun.AddDate(0, 0, 2))
	foydalanuvchi.TadbirQoshish("Dars 2", bugun.AddDate(0, 0, 3))

	foydalanuvchi.TadbirlarniChiqarish()
}
