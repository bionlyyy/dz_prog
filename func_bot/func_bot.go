package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"
	"math"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
)

type Function func(float64) float64

func main() {
	bot, err := tgbotapi.NewBotAPI("5721764505:AAE_N8uVY4irFjVW4jONjQL-kHpnQgDjPmo")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Авторизован как %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	u.Limit = 1

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Обработка команды /start
		if update.Message.IsCommand() && update.Message.Command() == "start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				"Отправьте функцию и интервал в формате: функция a b\n"+
					"Пример: x^2 0 5")
			bot.Send(msg)
			continue
		}

		// Обработка математического выражения
		response, imgData, err := processMathExpression(update.Message.Text)
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка: "+err.Error())
			bot.Send(msg)
			continue
		}

		// Отправка графика
		if imgData != nil {
			file := tgbotapi.FileBytes{Name: "graph.png", Bytes: imgData}
			photo := tgbotapi.NewPhoto(update.Message.Chat.ID, file)
			photo.Caption = response
			bot.Send(photo)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
			bot.Send(msg)
		}
	}
}

func processMathExpression(text string) (string, []byte, error) {
	parts := strings.Fields(text)
	if len(parts) != 3 {
		return "", nil, fmt.Errorf("неверный формат. Пример: x^2 0 5")
	}

	expr := parts[0]
	a, err1 := strconv.ParseFloat(parts[1], 64)
	b, err2 := strconv.ParseFloat(parts[2], 64)

	if err1 != nil || err2 != nil {
		return "", nil, fmt.Errorf("неверные числа интервала")
	}

	if a >= b {
		return "", nil, fmt.Errorf("начало интервала должно быть меньше конца")
	}

	fn, err := parseFunction(expr)
	if err != nil {
		return "", nil, fmt.Errorf("ошибка в функции: %v", err)
	}

	area := calculateArea(fn, a, b)
	imgData, err := plotFunction(fn, a, b)
	if err != nil {
		return "", nil, fmt.Errorf("ошибка построения графика: %v", err)
	}

	response := fmt.Sprintf("Функция: %s\nИнтервал: [%.2f, %.2f]\nПлощадь: %.4f", expr, a, b, area)
	return response, imgData, nil
}

func parseFunction(expr string) (Function, error) {
	switch expr {
	case "sin(x)":
		return math.Sin, nil
	case "cos(x)":
		return math.Cos, nil
	case "tan(x)":
		return math.Tan, nil
	case "exp(x)":
		return math.Exp, nil
	case "log(x)":
		return math.Log, nil
	case "sqrt(x)":
		return math.Sqrt, nil
	}

	if strings.HasPrefix(expr, "x^") {
		power, err := strconv.ParseFloat(expr[2:], 64)
		if err != nil {
			return nil, fmt.Errorf("неверная степень")
		}
		return func(x float64) float64 { return math.Pow(x, power) }, nil
	}

	if strings.Contains(expr, "*x") {
		parts := strings.Split(expr, "*x")
		a, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return nil, fmt.Errorf("неверный коэффициент")
		}

		var b float64
		if len(parts) > 1 && parts[1] != "" {
			b, err = strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, fmt.Errorf("неверный свободный член")
			}
		}

		return func(x float64) float64 { return a*x + b }, nil
	}

	return nil, fmt.Errorf("неподдерживаемая функция")
}

func calculateArea(f Function, a, b float64) float64 {
	const n = 1000
	h := (b - a) / n
	sum := 0.0

	for i := 0; i < n; i++ {
		x0 := a + float64(i)*h
		x1 := x0 + h
		sum += (f(x0) + f(x1)) * h / 2
	}

	return sum
}

func plotFunction(f Function, a, b float64) ([]byte, error) {
	p := plot.New()
	p.Title.Text = "График функции"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Создаем точки для графика
	points := make(plotter.XYs, 100)
	for i := range points {
		x := a + (b-a)*float64(i)/float64(len(points)-1)
		points[i].X = x
		points[i].Y = f(x)
	}

	line, err := plotter.NewLine(points)
	if err != nil {
		return nil, err
	}
	p.Add(line)

	// Создаем полигон для заливки
	poly := make(plotter.XYs, len(points)+2)
	copy(poly, points)
	poly[len(points)].X = b
	poly[len(points)].Y = 0
	poly[len(points)+1].X = a
	poly[len(points)+1].Y = 0

	area, err := plotter.NewPolygon(poly)
	if err != nil {
		return nil, err
	}
	area.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	p.Add(area)

	// Сохраняем в PNG
	img := vgimg.New(400, 300)
	dc := draw.New(img)
	p.Draw(dc)

	var buf bytes.Buffer
	imgCanvas := vgimg.PngCanvas{Canvas: img}
	if _, err := imgCanvas.WriteTo(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
