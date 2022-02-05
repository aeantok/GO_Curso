package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//fmt.Printf("%v\n", timeZone) // or just fmt.Println(timeZone)
	fmt.Printf("Fecha: %s\nHora: %s", queDiaEs(), queHoraEs())
}

func queDiaEs() string {
	t := time.Now()
	d := t.Local().Day()
	m := t.Local().Month()
	y := t.Local().Year()
	msg := ""

	//Día
	if len(strconv.FormatInt(int64(d), 10)) < 2 {
		//msg += "0" + strconv.FormatInt(int64(d), 10)
		msg += "0" + strconv.Itoa(d)
	} else {
		//msg += strconv.FormatInt(int64(d), 10)
		msg += strconv.Itoa(d)
	}

	//Mes
	if len(strconv.FormatInt(int64(m), 10)) < 2 {
		msg += "/0" + strconv.FormatInt(int64(m), 10)
	} else {
		msg += "/" + strconv.FormatInt(int64(m), 10)
	}

	//Año
	//msg += "/" + strconv.FormatInt(int64(y), 10)
	msg += "/" + strconv.Itoa(y)
	//fmt.Printf("Fecha: %s", msg)
	return msg
}

func queHoraEs() string {
	t := time.Now()
	h := t.Local().Hour()
	m := t.Local().Minute()
	s := t.Local().Second()
	msg := ""

	//Hora
	//if len(strconv.FormatInt(int64(d), 10)) < 2 {
	if len(strconv.Itoa(h)) < 2 {
		//msg += "0" + strconv.FormatInt(int64(d), 10)
		msg += "0" + strconv.Itoa(h)
	} else {
		//msg += strconv.FormatInt(int64(d), 10)
		msg += strconv.Itoa(h)
	}

	//Minutos
	//if len(strconv.FormatInt(int64(m), 10)) < 2 {
	if len(strconv.Itoa(m)) < 2 {
		//msg += ":0" + strconv.FormatInt(int64(m), 10)
		msg += ":0" + strconv.Itoa(m)
	} else {
		//msg += ":" + strconv.FormatInt(int64(m), 10)
		msg += ":" + strconv.Itoa(m)
	}

	//Segundos

	//Minutos
	//if len(strconv.FormatInt(int64(m), 10)) < 2 {
	if len(strconv.Itoa(s)) < 2 {
		//msg += ":0" + strconv.FormatInt(int64(s), 10)
		msg += ":0" + strconv.Itoa(s)
	} else {
		//msg += ":" + strconv.FormatInt(int64(s), 10)
		msg += ":" + strconv.Itoa(s)
	}

	return msg
}
